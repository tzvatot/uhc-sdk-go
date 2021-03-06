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

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/accountsmgmt/v1

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	time "time"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// RoleBindingClient is the client of the 'role_binding' resource.
//
// Manages a specific role binding.
type RoleBindingClient struct {
	transport http.RoundTripper
	path      string
}

// NewRoleBindingClient creates a new client for the 'role_binding'
// resource using the given transport to sned the requests and receive the
// responses.
func NewRoleBindingClient(transport http.RoundTripper, path string) *RoleBindingClient {
	client := new(RoleBindingClient)
	client.transport = transport
	client.path = path
	return client
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of the role binding.
func (c *RoleBindingClient) Get() *RoleBindingGetRequest {
	request := new(RoleBindingGetRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// Delete creates a request for the 'delete' method.
//
// Deletes the role binding.
func (c *RoleBindingClient) Delete() *RoleBindingDeleteRequest {
	request := new(RoleBindingDeleteRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// RoleBindingGetRequest is the request for the 'get' method.
type RoleBindingGetRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
}

// Context sets the context that will be used to send the request.
func (r *RoleBindingGetRequest) Context(value context.Context) *RoleBindingGetRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *RoleBindingGetRequest) Timeout(value time.Duration) *RoleBindingGetRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *RoleBindingGetRequest) Deadline(value time.Time) *RoleBindingGetRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *RoleBindingGetRequest) Parameter(name string, value interface{}) *RoleBindingGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *RoleBindingGetRequest) Header(name string, value interface{}) *RoleBindingGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *RoleBindingGetRequest) Send() (result *RoleBindingGetResponse, err error) {
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
	result = new(RoleBindingGetResponse)
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

// RoleBindingGetResponse is the response for the 'get' method.
type RoleBindingGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *RoleBinding
}

// Status returns the response status code.
func (r *RoleBindingGetResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *RoleBindingGetResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *RoleBindingGetResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *RoleBindingGetResponse) Body() *RoleBinding {
	return r.body
}

// unmarshal is the method used internally to unmarshal responses to the
// 'get' method.
func (r *RoleBindingGetResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(roleBindingData)
	err = decoder.Decode(data)
	if err != nil {
		return err
	}
	r.body, err = data.unwrap()
	if err != nil {
		return err
	}
	return err
}

// RoleBindingDeleteRequest is the request for the 'delete' method.
type RoleBindingDeleteRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
}

// Context sets the context that will be used to send the request.
func (r *RoleBindingDeleteRequest) Context(value context.Context) *RoleBindingDeleteRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *RoleBindingDeleteRequest) Timeout(value time.Duration) *RoleBindingDeleteRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *RoleBindingDeleteRequest) Deadline(value time.Time) *RoleBindingDeleteRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *RoleBindingDeleteRequest) Parameter(name string, value interface{}) *RoleBindingDeleteRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *RoleBindingDeleteRequest) Header(name string, value interface{}) *RoleBindingDeleteRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *RoleBindingDeleteRequest) Send() (result *RoleBindingDeleteResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodDelete,
		URL:    uri,
		Header: header,
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(RoleBindingDeleteResponse)
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
	return
}

// RoleBindingDeleteResponse is the response for the 'delete' method.
type RoleBindingDeleteResponse struct {
	status int
	header http.Header
	err    *errors.Error
}

// Status returns the response status code.
func (r *RoleBindingDeleteResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *RoleBindingDeleteResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *RoleBindingDeleteResponse) Error() *errors.Error {
	return r.err
}
