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
var constKeysTpl *template.Template

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
	constKeysTpl = template.Must(template.New("constKeys").Funcs(tplFuncs).Parse(constKeysTxt))
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

type constKeys struct {
	IC *Inception
	SI *StructInfo
}

var constKeysTxt = `
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
`
