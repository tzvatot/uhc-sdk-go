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
	time "time"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// UserClient is the client of the 'user' resource.
//
// Manages a specific user.
type UserClient struct {
	transport http.RoundTripper
	path      string
}

// NewUserClient creates a new client for the 'user'
// resource using the given transport to sned the requests and receive the
// responses.
func NewUserClient(transport http.RoundTripper, path string) *UserClient {
	client := new(UserClient)
	client.transport = transport
	client.path = path
	return client
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of the user.
func (c *UserClient) Get() *UserGetRequest {
	request := new(UserGetRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// Delete creates a request for the 'delete' method.
//
// Deletes the user.
func (c *UserClient) Delete() *UserDeleteRequest {
	request := new(UserDeleteRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// UserGetRequest is the request for the 'get' method.
type UserGetRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
}

// Context sets the context that will be used to send the request.
func (r *UserGetRequest) Context(value context.Context) *UserGetRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *UserGetRequest) Timeout(value time.Duration) *UserGetRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *UserGetRequest) Deadline(value time.Time) *UserGetRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *UserGetRequest) Parameter(name string, value interface{}) *UserGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *UserGetRequest) Header(name string, value interface{}) *UserGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *UserGetRequest) Send() (result *UserGetResponse, err error) {
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
	result = new(UserGetResponse)
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

// UserGetResponse is the response for the 'get' method.
type UserGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *User
}

// Status returns the response status code.
func (r *UserGetResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *UserGetResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *UserGetResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *UserGetResponse) Body() *User {
	return r.body
}

// unmarshal is the method used internally to unmarshal responses to the
// 'get' method.
func (r *UserGetResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(userData)
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

// UserDeleteRequest is the request for the 'delete' method.
type UserDeleteRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
}

// Context sets the context that will be used to send the request.
func (r *UserDeleteRequest) Context(value context.Context) *UserDeleteRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *UserDeleteRequest) Timeout(value time.Duration) *UserDeleteRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *UserDeleteRequest) Deadline(value time.Time) *UserDeleteRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *UserDeleteRequest) Parameter(name string, value interface{}) *UserDeleteRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *UserDeleteRequest) Header(name string, value interface{}) *UserDeleteRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *UserDeleteRequest) Send() (result *UserDeleteResponse, err error) {
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
	result = new(UserDeleteResponse)
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

// UserDeleteResponse is the response for the 'delete' method.
type UserDeleteResponse struct {
	status int
	header http.Header
	err    *errors.Error
}

// Status returns the response status code.
func (r *UserDeleteResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *UserDeleteResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *UserDeleteResponse) Error() *errors.Error {
	return r.err
}
