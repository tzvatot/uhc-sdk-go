/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

import (
	"fmt"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// flavourListData is type used internally to marshal and unmarshal lists of objects
// of type 'flavour'.
type flavourListData []*flavourData

// UnmarshalFlavourList reads a list of values of the 'flavour'
// from the given source, which can be a slice of bytes, a string, an io.Reader or a
// json.Decoder.
func UnmarshalFlavourList(source interface{}) (list *FlavourList, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	var data flavourListData
	err = decoder.Decode(&data)
	if err != nil {
		return
	}
	list, err = data.unwrap()
	return
}

// wrap is the method used internally to convert a list of values of the
// 'flavour' value to a JSON document.
func (l *FlavourList) wrap() (data flavourListData, err error) {
	if l == nil {
		return
	}
	data = make(flavourListData, len(l.items))
	for i, item := range l.items {
		data[i], err = item.wrap()
		if err != nil {
			return
		}
	}
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// list of values of the 'flavour' type.
func (d flavourListData) unwrap() (list *FlavourList, err error) {
	if d == nil {
		return
	}
	items := make([]*Flavour, len(d))
	for i, item := range d {
		items[i], err = item.unwrap()
		if err != nil {
			return
		}
	}
	list = new(FlavourList)
	list.items = items
	return
}

// flavourListLinkData is type used internally to marshal and unmarshal links
// to lists of objects of type 'flavour'.
type flavourListLinkData struct {
	Kind  *string        "json:\"kind,omitempty\""
	HREF  *string        "json:\"href,omitempty\""
	Items []*flavourData "json:\"items,omitempty\""
}

// wrapLink is the method used internally to convert a list of values of the
// 'flavour' value to a link.
func (l *FlavourList) wrapLink() (data *flavourListLinkData, err error) {
	if l == nil {
		return
	}
	items := make([]*flavourData, len(l.items))
	for i, item := range l.items {
		items[i], err = item.wrap()
		if err != nil {
			return
		}
	}
	data = new(flavourListLinkData)
	data.Items = items
	data.HREF = l.href
	data.Kind = new(string)
	if l.link {
		*data.Kind = FlavourListLinkKind
	} else {
		*data.Kind = FlavourListKind
	}
	return
}

// unwrapLink is the function used internally to convert a JSON link to a list
// of values of the 'flavour' type to a list.
func (d *flavourListLinkData) unwrapLink() (list *FlavourList, err error) {
	if d == nil {
		return
	}
	items := make([]*Flavour, len(d.Items))
	for i, item := range d.Items {
		items[i], err = item.unwrap()
		if err != nil {
			return
		}
	}
	list = new(FlavourList)
	list.items = items
	list.href = d.HREF
	if d.Kind != nil {
		switch *d.Kind {
		case FlavourListKind:
			list.link = false
		case FlavourListLinkKind:
			list.link = true
		default:
			err = fmt.Errorf(
				"expected kind '%s' or '%s' but got '%s'",
				FlavourListKind,
				FlavourListLinkKind,
				*d.Kind,
			)
			return
		}
	}
	return
}
