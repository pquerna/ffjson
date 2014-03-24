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
	"encoding/json"
	"reflect"
)

type StructField struct {
	Name             string
	JsonName         string
	Typ              reflect.Type
	OmitEmpty        bool
	ForceString      bool
	HasMarshalJSON   bool
	HasUnmarshalJSON bool
}

type StructInfo struct {
	Name   string
	Obj    interface{}
	Typ    reflect.Type
	Fields []*StructField
}

func NewStructInfo(obj interface{}) *StructInfo {
	t := reflect.TypeOf(obj)
	return &StructInfo{
		Obj:    obj,
		Name:   t.Name(),
		Typ:    t,
		Fields: extractFields(obj),
	}
}

var marshalerType = reflect.TypeOf(new(json.Marshaler)).Elem()
var unmarshalerType = reflect.TypeOf(new(json.Unmarshaler)).Elem()

func extractFields(obj interface{}) []*StructField {
	rv := make([]*StructField, 0)
	typ := reflect.TypeOf(obj)
	for i := 0; i < typ.NumField(); i++ {
		f := typ.Field(i)

		jsonName := f.Name
		omitEmpty := false
		forceString := false

		tag := f.Tag.Get("json")

		if tag != "" {
			tagName, opts := parseTag(tag)
			if tagName != "" {
				jsonName = tagName
			}
			omitEmpty = opts.Contains("omitempty")
			forceString = opts.Contains("string")
		}

		sf := &StructField{
			Name:             f.Name,
			JsonName:         jsonName,
			Typ:              f.Type,
			HasMarshalJSON:   f.Type.Implements(marshalerType),
			HasUnmarshalJSON: f.Type.Implements(unmarshalerType),
			OmitEmpty:        omitEmpty,
			ForceString:      forceString,
		}
		rv = append(rv, sf)
	}
	return rv
}
