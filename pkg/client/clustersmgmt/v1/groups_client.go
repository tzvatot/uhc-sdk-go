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
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"path"
	time "time"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// GroupsClient is the client of the 'groups' resource.
//
// Manages the collection of groups of a cluster.
type GroupsClient struct {
	transport http.RoundTripper
	path      string
}

// NewGroupsClient creates a new client for the 'groups'
// resource using the given transport to sned the requests and receive the
// responses.
func NewGroupsClient(transport http.RoundTripper, path string) *GroupsClient {
	client := new(GroupsClient)
	client.transport = transport
	client.path = path
	return client
}

// List creates a request for the 'list' method.
//
// Retrieves the list of groups.
func (c *GroupsClient) List() *GroupsListRequest {
	request := new(GroupsListRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// Group returns the target 'group' resource for the given identifier.
//
// Reference to the service that manages an specific group.
func (c *GroupsClient) Group(id string) *GroupClient {
	return NewGroupClient(c.transport, path.Join(c.path, id))
}

// GroupsListRequest is the request for the 'list' method.
type GroupsListRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
}

// Context sets the context that will be used to send the request.
func (r *GroupsListRequest) Context(value context.Context) *GroupsListRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *GroupsListRequest) Timeout(value time.Duration) *GroupsListRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *GroupsListRequest) Deadline(value time.Time) *GroupsListRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *GroupsListRequest) Parameter(name string, value interface{}) *GroupsListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *GroupsListRequest) Header(name string, value interface{}) *GroupsListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *GroupsListRequest) Send() (result *GroupsListResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodGet,
		URL:    uri,
		Header: header,
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(GroupsListResponse)
	result.status = response.StatusCode
	result.header = response.Header
	if result.status >= 400 {
		result.err, err = errors.UnmarshalError(response.Body)
		if err != nil {
			return
		}
		err = result.err
		return
	}
	err = result.unmarshal(response.Body)
	if err != nil {
		return
	}
	return
}

// GroupsListResponse is the response for the 'list' method.
type GroupsListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *GroupList
}

// Status returns the response status code.
func (r *GroupsListResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *GroupsListResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *GroupsListResponse) Error() *errors.Error {
	return r.err
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *GroupsListResponse) Page() int {
	if r.page != nil {
		return *r.page
	}
	return 0
}

// Size returns the value of the 'size' parameter.
//
// Number of items contained in the returned page.
func (r *GroupsListResponse) Size() int {
	if r.size != nil {
		return *r.size
	}
	return 0
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection.
func (r *GroupsListResponse) Total() int {
	if r.total != nil {
		return *r.total
	}
	return 0
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of groups.
func (r *GroupsListResponse) Items() *GroupList {
	return r.items
}

// unmarshal is the method used internally to unmarshal responses to the
// 'list' method.
func (r *GroupsListResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(groupsListResponseData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.page = data.Page
	r.size = data.Size
	r.total = data.Total
	r.items, err = data.Items.unwrap()
	if err != nil {
		return err
	}
	return err
}

// groupsListResponseData is the structure used internally to unmarshal
// the response of the 'list' method.
type groupsListResponseData struct {
	Page  *int          "json:\"page,omitempty\""
	Size  *int          "json:\"size,omitempty\""
	Total *int          "json:\"total,omitempty\""
	Items groupListData "json:\"items,omitempty\""
}
