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
	ic.OutputImports[`fflib "github.com/pquerna/ffjson/fflib/v1"`] = true
	ic.OutputImports[`"bytes"`] = true
	ic.OutputImports[`"fmt"`] = true

	out += tplStr(decodeTpl["header"], header{
		IC: ic,
		SI: si,
	})

	out += tplStr(decodeTpl["ujFunc"], ujFunc{
		SI:          si,
		IC:          ic,
		ValidValues: validValues,
	})

	ic.OutputFuncs = append(ic.OutputFuncs, out)

	return nil
}

func handleField(ic *Inception, name string, typ reflect.Type) string {
	return handleFieldAddr(ic, name, false, typ)
}

func handleFieldAddr(ic *Inception, name string, takeAddr bool, typ reflect.Type) string {
	out := ""
	out += fmt.Sprintf("/* handler: %s type=%v kind=%v */\n", name, typ, typ.Kind())

	out += tplStr(decodeTpl["handleUnmarshaler"], handleUnmarshaler{
		IC:                   ic,
		Name:                 name,
		Type:                 typ,
		Ptr:                  reflect.Ptr,
		UnmarshalJSONFFLexer: typ.Implements(unmarshalFasterType) || typeInInception(ic, typ),
		Unmarshaler:          typ.Implements(unmarshalerType) || reflect.PtrTo(typ).Implements(unmarshalerType),
	})

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
		out += tplStr(decodeTpl["handleBool"], handleBool{
			Name: name,
			Typ:  typ,
		})

	case reflect.Ptr,
		reflect.Interface:
		out += tplStr(decodeTpl["handlePtr"], handlePtr{
			IC:   ic,
			Name: name,
			Typ:  typ,
		})

	case reflect.Array,
		reflect.Slice:
		out += tplStr(decodeTpl["handleArray"], handleArray{
			IC:   ic,
			Name: name,
			Typ:  typ,
			Ptr:  reflect.Ptr,
		})

	case reflect.String:
		out += tplStr(decodeTpl["handleString"], handleString{
			Name:     name,
			Typ:      typ,
			TakeAddr: takeAddr,
		})
	default:
		// TODO(pquerna): layering. let templates declare their needed modules?
		ic.OutputImports[`"encoding/json"`] = true
		out += tplStr(decodeTpl["handleFallback"], handleFallback{
			Name: name,
			Typ:  typ,
			Kind: typ.Kind(),
		})
	}

	return out
}

func getAllowTokens(name string, tokens ...string) string {
	return tplStr(decodeTpl["allowTokens"], allowTokens{
		Name:   name,
		Tokens: tokens,
	})
}

func getNumberHandler(ic *Inception, name string, takeAddr bool, typ reflect.Type, parsefunc string) string {
	return tplStr(decodeTpl["handlerNumeric"], handlerNumeric{
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
