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

package generator

func getOmitEmpty(sf *StructField) string {
	// TODO(pquerna): non-nil checks, look at isEmptyValue()
	//	return "if mj." + sf.Name + " != nil {" + "\n"
	return "if true {" + "\n"
}

func getValue(sf *StructField) string {
	var out = ""
	// TODO(pquerna): non-nil checks, look at isEmptyValue()
	switch sf.Type {
	case "uint", "uint8", "uint16", "uint32", "uint64":
		out += "buf.Write(strconv.AppendUint([]byte{}, uint64(mj." + sf.Name + "), 10))" + "\n"
	case "int", "int8", "int16", "int32", "int64":
		out += "buf.Write(strconv.AppendInt([]byte{}, int64(mj." + sf.Name + "), 10))" + "\n"
	case "string":
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

func CreateMarshalJSON(si *StructInfo) (string, error) {
	var out = ""

	out += `func (mj *` + si.Name + `) MarshalJSON() ([]byte, error) {` + "\n"
	out += `var buf bytes.Buffer` + "\n"
	out += `var err error` + "\n"
	out += `var obj []byte` + "\n"
	out += `_ = obj` + "\n"
	out += `_ = err` + "\n"
	out += "buf.WriteString(`{`)" + "\n"

	var first = true

	for _, f := range si.Fields {
		if f.JsonName == "-" {
			continue
		}

		if f.OmitEmpty {
			out += getOmitEmpty(&f)
		}
		if !first {
			out += "buf.WriteString(`,\"`)" + "\n"
		} else {
			out += "buf.WriteString(`\"`)" + "\n"
			first = false
		}
		out += "buf.WriteString(`" + f.JsonName + "`)" + "\n"
		out += "buf.WriteString(`\":`)" + "\n"
		out += getValue(&f)
		if f.OmitEmpty {
			out += "}" + "\n"
		}
	}

	out += "buf.WriteString(`}`)" + "\n"
	//	out += "println(string(buf.Bytes()))" + "\n"
	out += `return buf.Bytes(), nil` + "\n"
	out += `}` + "\n"
	return out, nil
}
