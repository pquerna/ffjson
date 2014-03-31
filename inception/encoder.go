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
	"github.com/pquerna/ffjson/pills"
	"reflect"
)

func typeInInception(ic *Inception, typ reflect.Type) bool {
	for _, v := range ic.objs {
		if v.Typ == typ {
			return true
		}
	}

	return false
}

func getOmitEmpty(ic *Inception, sf *StructField) string {
	switch sf.Typ.Kind() {

	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return "if len(mj." + sf.Name + ") != 0 {" + "\n"

	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64,
		reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr,
		reflect.Float32,
		reflect.Float64:
		return "if mj." + sf.Name + " != 0 {" + "\n"

	case reflect.Bool:
		return "if mj." + sf.Name + " != false {" + "\n"

	case reflect.Interface, reflect.Ptr:
		// TODO(pquerna): pointers. oops.
		return "if mj." + sf.Name + " != nil {" + "\n"

	default:
		// TODO(pquerna): fix types
		return "if true {" + "\n"
	}
}

func getGetInnerValue(ic *Inception, name string, typ reflect.Type) string {
	var out = ""
	if typ.Implements(marshalerBufType) || typeInInception(ic, typ) {
		out += "err = " + name + ".MarshalJSONBuf(buf)" + "\n"
		out += "if err != nil {" + "\n"
		out += "  return err" + "\n"
		out += "}" + "\n"
		return out
	}

	if typ.Implements(marshalerType) {
		out += "obj, err = " + name + ".MarshalJSON()" + "\n"
		out += "if err != nil {" + "\n"
		out += "  return err" + "\n"
		out += "}" + "\n"
		out += "buf.Write(obj)" + "\n"
		return out
	}

	switch typ.Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		ic.OutputPills[pills.Pill_FormatBits] = true
		out += "ffjson_FormatBits(buf, uint64(" + name + "), 10, " + name + " < 0)" + "\n"
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:
		ic.OutputPills[pills.Pill_FormatBits] = true
		out += "ffjson_FormatBits(buf, uint64(" + name + "), 10, false)" + "\n"
	case reflect.Float32,
		reflect.Float64:
		ic.OutputImports[`"strconv"`] = true
		out += "buf.Write(strconv.AppendFloat([]byte{}, float64(" + name + "), 10))" + "\n"
	case reflect.Array,
		reflect.Slice:
		out += "if " + name + "!= nil {" + "\n"
		out += "buf.WriteString(`[`)" + "\n"
		out += "for _, v := range " + name + "{" + "\n"
		out += getGetInnerValue(ic, "v", typ.Elem())
		out += "}" + "\n"
		out += "buf.WriteString(`]`)" + "\n"
		out += "} else {" + "\n"
		out += "buf.WriteString(`null`)" + "\n"
		out += "}" + "\n"
	case reflect.String:
		ic.OutputPills[pills.Pill_WriteJsonString] = true
		out += "ffjson_WriteJsonString(buf, " + name + ")" + "\n"
	case reflect.Ptr,
		reflect.Interface:
		out += "if " + name + "!= nil {" + "\n"
		out += getGetInnerValue(ic, "v", typ.Elem())
		out += "} else {" + "\n"
		out += "buf.WriteString(`null`)" + "\n"
		out += "}" + "\n"
	case reflect.Bool:
		out += "if " + name + " {" + "\n"
		out += "buf.WriteString(`true`)" + "\n"
		out += "} else {" + "\n"
		out += "buf.WriteString(`false`)" + "\n"
		out += "}" + "\n"
	default:
		ic.OutputImports[`"encoding/json"`] = true
		out += fmt.Sprintf("/* Falling back. type=%v kind=%v */\n", typ, typ.Kind())
		out += "obj, err = json.Marshal(" + name + ")" + "\n"
		out += "if err != nil {" + "\n"
		out += "  return err" + "\n"
		out += "}" + "\n"
		out += "buf.Write(obj)" + "\n"
	}
	return out
}

func getValue(ic *Inception, sf *StructField) string {
	return getGetInnerValue(ic, "mj."+sf.Name, sf.Typ)
}

func CreateMarshalJSON(ic *Inception, si *StructInfo) error {
	var out = ""

	ic.OutputImports[`"bytes"`] = true

	out += `func (mj *` + si.Name + `) MarshalJSON() ([]byte, error) {` + "\n"
	out += `var buf bytes.Buffer` + "\n"
	out += "buf.Grow(1024)" + "\n" // TOOD(pquerna): automatically calc a good size!
	out += `err := mj.MarshalJSONBuf(&buf)` + "\n"
	out += `if err != nil {` + "\n"
	out += "  return nil, err" + "\n"
	out += `}` + "\n"
	out += `return buf.Bytes(), nil` + "\n"
	out += `}` + "\n"

	out += `func (mj *` + si.Name + `) MarshalJSONBuf(buf *bytes.Buffer) (error) {` + "\n"
	out += `var err error` + "\n"
	out += `var obj []byte` + "\n"
	out += `var first bool = true` + "\n"
	out += `_ = obj` + "\n"
	out += `_ = err` + "\n"
	out += `_ = first` + "\n"
	out += "buf.WriteString(`{`)" + "\n"

	for _, f := range si.Fields {
		if f.JsonName == "-" {
			continue
		}

		if f.OmitEmpty {
			out += getOmitEmpty(ic, f)
		}

		out += "if first == true {" + "\n"
		out += "first = false" + "\n"
		out += "} else {" + "\n"
		out += "buf.WriteString(`,`)" + "\n"
		out += "}" + "\n"

		// JsonName is already escaped and quoted.
		out += "buf.WriteString(`" + f.JsonName + ":`)" + "\n"
		out += getValue(ic, f)
		if f.OmitEmpty {
			out += "}" + "\n"
		}
	}

	out += "buf.WriteString(`}`)" + "\n"
	out += `return nil` + "\n"
	out += `}` + "\n"
	ic.OutputFuncs = append(ic.OutputFuncs, out)
	return nil
}
