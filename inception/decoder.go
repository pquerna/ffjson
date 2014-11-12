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
	"fmt"
	"reflect"
)

var validValues []string = []string{
	"FFTok_left_brace",
	"FFTok_left_bracket",
	"FFTok_integer",
	"FFTok_double",
	"FFTok_string",
	"FFTok_string_with_escapes",
	"FFTok_bool",
	"FFTok_null",
}

func CreateUnmarshalJSON(ic *Inception, si *StructInfo) error {
	out := ""
	ic.OutputImports[`ffjson_scanner "github.com/pquerna/ffjson/scanner"`] = true
	ic.OutputImports[`ffjson_pills "github.com/pquerna/ffjson/pills"`] = true
	ic.OutputImports[`"bytes"`] = true
	ic.OutputImports[`"fmt"`] = true

	out += tplStr(constKeysTpl, constKeys{
		IC: ic,
		SI: si,
	})

	for _, f := range si.Fields {
		if f.JsonName == "-" {
			continue
		}
		out += `var ffj_key_` + si.Name + `_` + f.Name + ` = []byte(` + f.JsonName + `)` + "\n"
	}

	out += `func (uj *` + si.Name + `) XUnmarshalJSON(input []byte) error {` + "\n"
	out += `	fs := ffjson_scanner.NewFFLexer(input)` + "\n"
	out += `    return uj.UnmarshalJSONFFLexer(fs, ffjson_scanner.FFParse_map_start)` + "\n"
	out += `}` + "\n"

	out += `func (uj *` + si.Name + `) UnmarshalJSONFFLexer(fs *ffjson_scanner.FFLexer, state ffjson_scanner.FFParseState) error {` + "\n"
	out += `var err error = nil` + "\n"
	out += `currentKey := ffj_t_` + si.Name + `base` + "\n"
	out += `_ = currentKey` + "\n"
	out += `tok := ffjson_scanner.FFTok_init` + "\n"
	out += `wantedTok := ffjson_scanner.FFTok_init` + "\n"
	out += `mainparse:` + "\n"
	out += `for {` + "\n"
	out += `	tok = fs.Scan()` + "\n"
	// out += `	println(fmt.Sprintf("debug: tok: %v  state: %v", tok, state))` + "\n"
	out += `	if tok == ffjson_scanner.FFTok_error {` + "\n"
	out += `		goto tokerror` + "\n"
	out += `	}` + "\n"
	out += `	switch state {` + "\n"

	out += `		case ffjson_scanner.FFParse_map_start:` + "\n"
	out += `if tok != ffjson_scanner.FFTok_left_bracket {` + "\n"
	out += `	wantedTok = ffjson_scanner.FFTok_left_bracket` + "\n"
	out += `	goto wrongtokenerror` + "\n"
	out += `}` + "\n"
	out += `state = ffjson_scanner.FFParse_want_key` + "\n"
	out += `continue` + "\n"

	out += `		case ffjson_scanner.FFParse_after_value:` + "\n"
	out += `if tok == ffjson_scanner.FFTok_comma {` + "\n"
	out += `	state = ffjson_scanner.FFParse_want_key` + "\n"
	out += `} else if tok == ffjson_scanner.FFTok_right_bracket {` + "\n"
	out += `	goto done` + "\n"
	out += `} else {` + "\n"
	out += `	wantedTok = ffjson_scanner.FFTok_comma` + "\n"
	out += `	goto wrongtokenerror` + "\n"
	out += `}` + "\n"

	out += `		case ffjson_scanner.FFParse_want_key:` + "\n"
	out += `		` + "\n"
	// json {} ended. goto exit. woo.
	out += `if tok == ffjson_scanner.FFTok_right_bracket {` + "\n"
	out += `	goto done` + "\n"
	out += `}` + "\n"
	out += `if tok != ffjson_scanner.FFTok_string {` + "\n"
	out += `	wantedTok = ffjson_scanner.FFTok_string` + "\n"
	out += `	goto wrongtokenerror` + "\n"
	out += `}` + "\n"
	// TODO(pquerna): convert keynames to bytes at generation time.
	out += `kn := fs.Output.Bytes()` + "\n"

	out += `if false {` + "\n"
	for _, f := range si.Fields {
		if f.JsonName == "-" {
			continue
		}
		out += `} else if bytes.Equal(ffj_key_` + si.Name + `_` + f.Name + `, kn) {` + "\n"
		out += `currentKey = ffj_t_` + si.Name + `_` + f.Name + "\n"
		out += `state = ffjson_scanner.FFParse_want_colon` + "\n"
		out += `continue` + "\n"
	}
	// a JSON name we didn't know about.
	// TOOD(pquerna): suck whole value.
	out += "} else {"
	out += `	currentKey = ffj_t_` + si.Name + `no_such_key` + "\n"
	out += `	state = ffjson_scanner.FFParse_want_colon` + "\n"
	out += `	continue` + "\n"
	out += `}` + "\n"

	out += `		case ffjson_scanner.FFParse_want_colon:` + "\n"
	out += `if tok != ffjson_scanner.FFTok_colon {` + "\n"
	out += `	wantedTok = ffjson_scanner.FFTok_colon` + "\n"
	out += `	goto wrongtokenerror` + "\n"
	out += `}` + "\n"
	out += `state = ffjson_scanner.FFParse_want_value` + "\n"
	out += `continue` + "\n"

	out += `case ffjson_scanner.FFParse_want_value:` + "\n"

	out += `if false `
	for _, v := range validValues {
		out += " || tok == ffjson_scanner." + v
	}
	out += ` {` + "\n"
	{
		out += `switch currentKey {` + "\n"
		for _, f := range si.Fields {
			if f.JsonName == "-" {
				continue
			}
			out += `case ffj_t_` + si.Name + `_` + f.Name + `:` + "\n"
			out += `goto handle_` + f.Name + "\n"
		}

		out += `case ffj_t_` + si.Name + `no_such_key:` + "\n"
		// TODO don't capture, skip
		out += `err = fs.SkipField(tok)` + "\n"
		out += "if err != nil {" + "\n"
		out += "  return fs.WrapErr(err)" + "\n"
		out += "}" + "\n"
		out += `state = ffjson_scanner.FFParse_after_value` + "\n"
		out += `goto mainparse` + "\n"
		out += `}` + "\n"
	}
	out += `} else {` + "\n"
	out += `	goto wantedvalue` + "\n"
	out += `}` + "\n"

	out += `	}` + "\n"
	out += `}` + "\n"

	for _, f := range si.Fields {
		if f.JsonName == "-" {
			continue
		}

		out += `handle_` + f.Name + `:` + "\n"
		// TODO: write handler.
		//out += `println("got: ` + f.Name + `")` + "\n"
		out += handleField(ic, "uj."+f.Name, f.Typ)
		out += `state = ffjson_scanner.FFParse_after_value` + "\n"
		out += `goto mainparse` + "\n"
	}

	out += "wraperr:" + "\n"
	// TODO: include line / byte offsets / field name
	// TODO: dont wrap all errors?
	out += `return fs.WrapErr(err)` + "\n"

	out += "wantedvalue:" + "\n"
	out += `return fs.WrapErr(fmt.Errorf("wanted value token, but got token: %v", tok))` + "\n"

	out += "wrongtokenerror:" + "\n"
	out += `return fs.WrapErr(fmt.Errorf("ffjson: wanted token: %v, but got token: %v output=%s", wantedTok, tok, fs.Output.String()))` + "\n"

	out += "tokerror:" + "\n"
	out += `if fs.BigError != nil {` + "\n"
	out += `return fs.BigError` + "\n"
	out += `}` + "\n"
	out += `err = fs.Error.ToError()` + "\n"
	out += `if err != nil {` + "\n"
	out += `return fs.WrapErr(err)` + "\n"
	out += `}` + "\n"
	out += `panic("ffjson-generated: unreachable, please report bug.")` + "\n"

	out += `done:` + "\n"
	out += `return nil` + "\n"
	out += `}` + "\n"

	ic.OutputFuncs = append(ic.OutputFuncs, out)

	return nil
}

func handleField(ic *Inception, name string, typ reflect.Type) string {
	return handleFieldAddr(ic, name, false, typ)
}

func handleFieldAddr(ic *Inception, name string, takeAddr bool, typ reflect.Type) string {
	out := ""
	out += fmt.Sprintf("/* handler: %s type=%v kind=%v */\n", name, typ, typ.Kind())

	if typ.Implements(unmarshalFasterType) || typeInInception(ic, typ) {
		if typ.Kind() == reflect.Ptr {
			out += `if ` + name + ` == nil {` + "\n"
			out += `   ` + name + ` = new(` + typ.Elem().Name() + `)` + "\n"
			out += `}` + "\n"
		}

		out += "err = " + name + ".UnmarshalJSONFFLexer(fs, ffjson_scanner.FFParse_want_key)" + "\n"
		out += "if err != nil {" + "\n"
		out += "  return err" + "\n"
		out += "}" + "\n"
		return out
	}

	if typ.Implements(unmarshalerType) || reflect.PtrTo(typ).Implements(unmarshalerType) {
		out += `{` + "\n"
		out += `tbuf, err := fs.CaptureField(tok)` + "\n"
		out += "if err != nil {" + "\n"
		out += `  return fs.WrapErr(err)` + "\n"
		out += `}` + "\n"
		out += `err = ` + name + `.UnmarshalJSON(tbuf)` + "\n"
		out += `if err != nil {` + "\n"
		out += `  return fs.WrapErr(err)` + "\n"
		out += "}" + "\n"
		out += `}` + "\n"
		return out
	}

	// TODO(pquerna): generic handling of token type mismatching struct type

	switch typ.Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		out += getAllowTokens(typ.Name(), "FFTok_integer", "FFTok_null")
		out += getNumberHandler(ic, name, takeAddr, typ, "ParseInt")

	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64:
		out += getAllowTokens(typ.Name(), "FFTok_integer", "FFTok_null")
		out += getNumberHandler(ic, name, takeAddr, typ, "ParseUint")

	case reflect.Float32,
		reflect.Float64:
		out += getAllowTokens(typ.Name(), "FFTok_double", "FFTok_null")
		out += getNumberHandler(ic, name, takeAddr, typ, "ParseFloat")

	case reflect.Bool:
		ic.OutputImports[`"bytes"`] = true
		ic.OutputImports[`"errors"`] = true
		out += tplStr(handleBoolTpl, handleBool{
			Name: name,
			Typ:  typ,
		})

	case reflect.Ptr,
		reflect.Interface:
		out += tplStr(handlePtrTpl, handlePtr{
			IC:   ic,
			Name: name,
			Typ:  typ,
		})

	case reflect.Array,
		reflect.Slice:
		out += tplStr(handleArrayTpl, handleArray{
			IC:   ic,
			Name: name,
			Typ:  typ,
			Ptr:  reflect.Ptr,
		})

	case reflect.String:
		out += tplStr(handleStringTpl, handleString{
			Name:     name,
			Typ:      typ,
			TakeAddr: takeAddr,
		})
	default:
		// TODO(pquerna): layering. let templates declare their needed modules?
		ic.OutputImports[`"encoding/json"`] = true
		out += tplStr(handleFallbackTpl, handleFallback{
			Name: name,
			Typ:  typ,
			Kind: typ.Kind(),
		})
	}

	return out
}

func getAllowTokens(name string, tokens ...string) string {
	return tplStr(allowTokensTpl, allowTokens{
		Name:   name,
		Tokens: tokens,
	})
}

func getNumberHandler(ic *Inception, name string, takeAddr bool, typ reflect.Type, parsefunc string) string {
	return tplStr(handlerNumericTpl, handlerNumeric{
		Name:      name,
		ParseFunc: parsefunc,
		TakeAddr:  takeAddr,
		Typ:       typ,
	})
}

func getNumberSize(typ reflect.Type) string {
	return fmt.Sprintf("%d", typ.Bits())
}

func getNumberCast(name string, typ reflect.Type) string {
	s := typ.Name()
	if s == "" {
		panic("non-numeric type passed in w/o name: " + name)
	}
	return s
}
