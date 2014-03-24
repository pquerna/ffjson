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
	"reflect"
)

func getOmitEmpty(ic *InceptionContext, sf *StructField) string {
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

func getValue(ic *InceptionContext, sf *StructField) string {
	var out = ""

	if sf.HasMarshalJSON {
		out += "obj, err = mj." + sf.Name + ".MarshalJSON()" + "\n"
		out += "if err != nil {" + "\n"
		out += "  return nil, err" + "\n"
		out += "}" + "\n"
		out += "buf.Write(obj)" + "\n"
		return out
	}

	switch sf.Typ.Kind() {
	case reflect.Int,
		reflect.Int8,
		reflect.Int16,
		reflect.Int32,
		reflect.Int64:
		out += "buf.Write(strconv.AppendInt([]byte{}, int64(mj." + sf.Name + "), 10))" + "\n"
	case reflect.Uint,
		reflect.Uint8,
		reflect.Uint16,
		reflect.Uint32,
		reflect.Uint64,
		reflect.Uintptr:
		out += "buf.Write(strconv.AppendUint([]byte{}, uint64(mj." + sf.Name + "), 10))" + "\n"
	case reflect.Float32,
		reflect.Float64:
		out += "buf.Write(strconv.AppendFloat([]byte{}, float64(mj." + sf.Name + "), 10))" + "\n"
	case reflect.String:
		out += "buf.WriteString(`\"`)" + "\n"
		out += "buf.WriteString(mj." + sf.Name + ")" + "\n"
		out += "buf.WriteString(`\"`)" + "\n"
	default:
		// println(sf.Type)
		out += "obj, err = json.Marshal(mj." + sf.Name + ")" + "\n"
		out += "if err != nil {" + "\n"
		out += "  return nil, err" + "\n"
		out += "}" + "\n"
		out += "buf.Write(obj)" + "\n"
	}
	return out
}

func CreateMarshalJSON(ic *InceptionContext, si *StructInfo) error {
	var out = ""

	out += `func (mj *` + si.Name + `) MarshalJSON() ([]byte, error) {` + "\n"
	out += `var buf bytes.Buffer` + "\n"
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
		out += "buf.WriteString(`\"`)" + "\n"
		out += "} else {" + "\n"
		out += "buf.WriteString(`,\"`)" + "\n"
		out += "}" + "\n"

		out += "buf.WriteString(`" + f.JsonName + "`)" + "\n"
		out += "buf.WriteString(`\":`)" + "\n"
		out += getValue(ic, f)
		if f.OmitEmpty {
			out += "}" + "\n"
		}
	}

	out += "buf.WriteString(`}`)" + "\n"
	// out += "println(string(buf.Bytes()))" + "\n"
	out += `return buf.Bytes(), nil` + "\n"
	out += `}` + "\n"
	ic.OutputFuncs = append(ic.OutputFuncs, out)
	return nil
}
