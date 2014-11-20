/**
 *  Copyright 2014 Paul Querna
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package ffjsoninception

import (
	"bytes"
	"reflect"
	"text/template"
)

var handlerNumericTpl *template.Template
var allowTokensTpl *template.Template
var handleFallbackTpl *template.Template
var handleArrayTpl *template.Template
var handleStringTpl *template.Template
var handleBoolTpl *template.Template
var handlePtrTpl *template.Template
var headerTpl *template.Template
var ujFuncTpl *template.Template
var handleUnmarshalerTpl *template.Template

func init() {
	var tplFuncs = template.FuncMap{
		"getAllowTokens":  getAllowTokens,
		"getNumberSize":   getNumberSize,
		"getNumberCast":   getNumberCast,
		"handleField":     handleField,
		"handleFieldAddr": handleFieldAddr,
	}

	handlerNumericTpl = template.Must(template.New("handlerNumeric").Funcs(tplFuncs).Parse(handlerNumericTxt))
	allowTokensTpl = template.Must(template.New("allowTokens").Parse(allowTokensTxt))
	handleFallbackTpl = template.Must(template.New("handleFallback").Funcs(tplFuncs).Parse(handleFallbackTxt))
	handleStringTpl = template.Must(template.New("handleString").Funcs(tplFuncs).Parse(handleStringTxt))
	handleArrayTpl = template.Must(template.New("handleArray").Funcs(tplFuncs).Parse(handleArrayTxt))
	handleBoolTpl = template.Must(template.New("handleBool").Funcs(tplFuncs).Parse(handleBoolTxt))
	handlePtrTpl = template.Must(template.New("handlePtr").Funcs(tplFuncs).Parse(handlePtrTxt))
	headerTpl = template.Must(template.New("header").Funcs(tplFuncs).Parse(headerTxt))
	ujFuncTpl = template.Must(template.New("ujFunc").Funcs(tplFuncs).Parse(ujFuncTxt))
	handleUnmarshalerTpl = template.Must(template.New("handleUnmarshaler").Funcs(tplFuncs).Parse(handleUnmarshalerTxt))
}

func tplStr(t *template.Template, data interface{}) string {
	buf := bytes.Buffer{}
	err := t.Execute(&buf, data)
	if err != nil {
		panic(err)
	}
	return buf.String()
}

type handlerNumeric struct {
	Name      string
	ParseFunc string
	Typ       reflect.Type
	TakeAddr  bool
}

var handlerNumericTxt = `
{
	{{if eq .ParseFunc "ParseFloat" }}
	tval, err := ffjson_pills.{{ .ParseFunc}}(fs.Output.Bytes(), {{getNumberSize .Typ}})
	{{else}}
	tval, err := ffjson_pills.{{ .ParseFunc}}(fs.Output.Bytes(), 10, {{getNumberSize .Typ}})
	{{end}}

	if err != nil {
		goto wraperr
	}
	{{if eq .TakeAddr true}}
	ttypval := {{getNumberCast .Name .Typ }}(tval)
	{{.Name}} = &ttypval
	{{else}}
	{{.Name}} = {{getNumberCast .Name .Typ}}(tval)
	{{end}}
}
`

type allowTokens struct {
	Name   string
	Tokens []string
}

var allowTokensTxt = `
{
	if {{range $index, $element := .Tokens}}{{if ne $index 0 }}&&{{end}} tok != ffjson_scanner.{{$element}}{{end}} {
		return fs.WrapErr(fmt.Errorf("cannot unmarshal %s into Go value for {{.Name}}", tok))
	}
}
`

type handleFallback struct {
	Name string
	Typ  reflect.Type
	Kind reflect.Kind
}

var handleFallbackTxt = `
{
	/* Falling back. type={{printf "%v" .Typ}} kind={{printf "%v" .Kind}} */
	tbuf, err := fs.CaptureField(tok)
	if err != nil {
		return fs.WrapErr(err)
	}

	err = json.Unmarshal(tbuf, &{{.Name}})
	if err != nil {
		return fs.WrapErr(err)
	}
}
`

type handleString struct {
	Name     string
	Typ      reflect.Type
	TakeAddr bool
}

var handleStringTxt = `
{
	{{getAllowTokens .Typ.Name "FFTok_string" "FFTok_string_with_escapes"}}
	{{if eq .TakeAddr true}}
	var tval string
	if tok == ffjson_scanner.FFTok_string_with_escapes {
		// TODO: decoding escapes.
		tval = fs.Output.String()
	} else {
		tval = fs.Output.String()
	}
	{{.Name}} = &tval
	{{else}}
	if tok == ffjson_scanner.FFTok_string_with_escapes {
		// TODO: decoding escapes.
		{{.Name}} = fs.Output.String()
	} else {
		{{.Name}} = fs.Output.String()
	}
	{{end}}
}
`

type handleArray struct {
	IC   *Inception
	Name string
	Typ  reflect.Type
	Ptr  reflect.Kind
}

var handleArrayTxt = `
{
	{{getAllowTokens .Typ.Name "FFTok_left_brace" "FFTok_null"}}
	if tok == ffjson_scanner.FFTok_null {
		{{.Name}} = nil
	} else {
	{{if eq .Typ.Elem.Kind .Ptr }}
		{{.Name}} = make([]*{{.Typ.Elem.Elem.Name}}, 0)
	{{else}}
		{{.Name}} = make([]{{.Typ.Elem.Name}}, 0)
	{{end}}
	}

	for {
	{{if eq .Typ.Elem.Kind .Ptr }}
		var v *{{.Typ.Elem.Elem.Name}}
	{{else}}
		var v {{.Typ.Elem.Name}}
	{{end}}

		tok = fs.Scan()
		if tok == ffjson_scanner.FFTok_error {
			goto tokerror
		}
		if tok == ffjson_scanner.FFTok_right_brace {
			break
		}
		// TODO(pquerna): this allows invalid json like [,,,,]
		if tok == ffjson_scanner.FFTok_comma {
			continue
		}
		{{handleField .IC "v" .Typ.Elem}}
		{{.Name}} = append({{.Name}}, v)
	}
}
`

type handleBool struct {
	Name     string
	Typ      reflect.Type
	TakeAddr bool
}

var handleBoolTxt = `
{
	{{getAllowTokens .Typ.Name "FFTok_bool" "FFTok_null"}}

	tmpb := fs.Output.Bytes()

	{{if eq .TakeAddr true}}
	var tval bool
	{{end}}
	if bytes.Compare([]byte{'t', 'r', 'u', 'e'}, tmpb) == 0 {
	{{if eq .TakeAddr true}}
		tval = true
	{{else}}
		{{.Name}} = true
	{{end}}
	} else if bytes.Compare([]byte{'f', 'a', 'l', 's', 'e'}, tmpb) == 0 {
	{{if eq .TakeAddr true}}
		tval = false
	{{else}}
		{{.Name}} = false
	{{end}}
	} else {
		err = errors.New("unexpected bytes for true/false value")
		goto wraperr
	}
	{{if eq .TakeAddr true}}
	{{.Name}} = &tval
	{{end}}
}
`

type handlePtr struct {
	IC   *Inception
	Name string
	Typ  reflect.Type
}

var handlePtrTxt = `
{
	if tok == ffjson_scanner.FFTok_null {
		{{.Name}} = nil
	} else {
		if {{.Name}} == nil {
			{{.Name}} = new({{.Typ.Elem.Name}})
		}

		{{handleFieldAddr .IC .Name true .Typ.Elem}}

	}
}
`

type header struct {
	IC *Inception
	SI *StructInfo
}

var headerTxt = `
const (
	ffj_t_{{.SI.Name}}base = iota
	ffj_t_{{.SI.Name}}no_such_key
	{{with $si := .SI}}
		{{range $index, $field := $si.Fields}}
			{{if ne $field.JsonName "-"}}
		ffj_t_{{$si.Name}}_{{$field.Name}}
			{{end}}
		{{end}}
	{{end}}
)

{{with $si := .SI}}
	{{range $index, $field := $si.Fields}}
		{{if ne $field.JsonName "-"}}
var ffj_key_{{$si.Name}}_{{$field.Name}} = []byte({{$field.JsonName}})
		{{end}}
	{{end}}
{{end}}

`

type ujFunc struct {
	IC          *Inception
	SI          *StructInfo
	ValidValues []string
}

var ujFuncTxt = `
{{$si := .SI}}
{{$ic := .IC}}

func (uj *{{.SI.Name}}) XUnmarshalJSON(input []byte) error {
	fs := ffjson_scanner.NewFFLexer(input)
    return uj.UnmarshalJSONFFLexer(fs, ffjson_scanner.FFParse_map_start)
}

func (uj *{{.SI.Name}}) UnmarshalJSONFFLexer(fs *ffjson_scanner.FFLexer, state ffjson_scanner.FFParseState) error {
	var err error = nil
	currentKey := ffj_t_{{.SI.Name}}base
	_ = currentKey
	tok := ffjson_scanner.FFTok_init
	wantedTok := ffjson_scanner.FFTok_init

mainparse:
	for {
		tok = fs.Scan()
		//	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))
		if tok == ffjson_scanner.FFTok_error {
			goto tokerror
		}

		switch state {

		case ffjson_scanner.FFParse_map_start:
			if tok != ffjson_scanner.FFTok_left_bracket {
				wantedTok = ffjson_scanner.FFTok_left_bracket
				goto wrongtokenerror
			}
			state = ffjson_scanner.FFParse_want_key
			continue

		case ffjson_scanner.FFParse_after_value:
			if tok == ffjson_scanner.FFTok_comma {
				state = ffjson_scanner.FFParse_want_key
			} else if tok == ffjson_scanner.FFTok_right_bracket {
				goto done
			} else {
				wantedTok = ffjson_scanner.FFTok_comma
				goto wrongtokenerror
			}

		case ffjson_scanner.FFParse_want_key:
			// json {} ended. goto exit. woo.
			if tok == ffjson_scanner.FFTok_right_bracket {
				goto done
			}
			if tok != ffjson_scanner.FFTok_string {
				wantedTok = ffjson_scanner.FFTok_string
				goto wrongtokenerror
			}

			kn := fs.Output.Bytes()

			{{range $index, $field := $si.Fields}}
			{{if ne $index 0 }}} else if {{else}}if {{end}} bytes.Equal(ffj_key_{{$si.Name}}_{{$field.Name}}, kn) {
				currentKey = ffj_t_{{$si.Name}}_{{$field.Name}}
				state = ffjson_scanner.FFParse_want_colon
				continue
			{{end}}} else {
				currentKey = ffj_t_{{.SI.Name}}no_such_key
				state = ffjson_scanner.FFParse_want_colon
				continue
			}

		case ffjson_scanner.FFParse_want_colon:
			if tok != ffjson_scanner.FFTok_colon {
				wantedTok = ffjson_scanner.FFTok_colon
				goto wrongtokenerror
			}
			state = ffjson_scanner.FFParse_want_value
			continue
		case ffjson_scanner.FFParse_want_value:

			if {{range $index, $v := .ValidValues}}{{if ne $index 0 }}||{{end}}tok == ffjson_scanner.{{$v}}{{end}} {
				switch currentKey {
				{{range $index, $field := $si.Fields}}
				case ffj_t_{{$si.Name}}_{{$field.Name}}:
					goto handle_{{$field.Name}}
				{{end}}
				case ffj_t_{{$si.Name}}no_such_key:
					err = fs.SkipField(tok)
					if err != nil {
						return fs.WrapErr(err)
					}
					state = ffjson_scanner.FFParse_after_value
					goto mainparse
				}
			} else {
				goto wantedvalue
			}
		}
	}

{{range $index, $field := $si.Fields}}
handle_{{$field.Name}}:
	{{with $fieldName := $field.Name | printf "uj.%s"}}
		{{handleField $ic $fieldName $field.Typ}}
		state = ffjson_scanner.FFParse_after_value
		goto mainparse
	{{end}}
{{end}}

wraperr:
	return fs.WrapErr(err)
wantedvalue:
	return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))
wrongtokenerror:
	return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))
tokerror:
	if fs.BigError != nil {
		return fs.BigError
	}
	err = fs.Error.ToError()
	if err != nil {
		return fs.WrapErr(err)
	}
	panic("ffjson-generated: unreachable, please report bug.")
done:
	return nil
}

`

type handleUnmarshaler struct {
	IC                   *Inception
	Name                 string
	Type                 reflect.Type
	Ptr                  reflect.Kind
	UnmarshalJSONFFLexer bool
	Unmarshaler          bool
}

var handleUnmarshalerTxt = `
	{{if eq .UnmarshalJSONFFLexer true}}
	{
		{{if eq .Type.Kind .Ptr }}
			if {{.Name}} == nil {
				{{.Name}} = new({{.Type.Elem.Name}})
			}
		{{end}}
		err = {{.Name}}.UnmarshalJSONFFLexer(fs, ffjson_scanner.FFParse_want_key)
		if err != nil {
			return err
		}
		state = ffjson_scanner.FFParse_after_value
		goto mainparse
	}
	{{end}}
	{{if eq .Unmarshaler true}}
	{
		tbuf, err := fs.CaptureField(tok)
		if err != nil {
			return fs.WrapErr(err)
		}

		err = {{.Name}}.UnmarshalJSON(tbuf)
		if err != nil {
			return fs.WrapErr(err)
		}
		state = ffjson_scanner.FFParse_after_value
		goto mainparse
	}
	{{end}}
`
