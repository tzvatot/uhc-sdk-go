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
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	time "time"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/errors"
	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// OrganizationClient is the client of the 'organization' resource.
//
// Manages a specific organization.
type OrganizationClient struct {
	transport http.RoundTripper
	path      string
}

// NewOrganizationClient creates a new client for the 'organization'
// resource using the given transport to sned the requests and receive the
// responses.
func NewOrganizationClient(transport http.RoundTripper, path string) *OrganizationClient {
	client := new(OrganizationClient)
	client.transport = transport
	client.path = path
	return client
}

// Get creates a request for the 'get' method.
//
// Retrieves the details of the organization.
func (c *OrganizationClient) Get() *OrganizationGetRequest {
	request := new(OrganizationGetRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// Update creates a request for the 'update' method.
//
// Updates the organization.
func (c *OrganizationClient) Update() *OrganizationUpdateRequest {
	request := new(OrganizationUpdateRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// ResourceQuota returns the target 'resource_quotas' resource.
//
// Reference to the service that manages the resource quotas for this
// organization.
func (c *OrganizationClient) ResourceQuota() *ResourceQuotasClient {
	return NewResourceQuotasClient(c.transport, path.Join(c.path, "resource_quota"))
}

// OrganizationGetRequest is the request for the 'get' method.
type OrganizationGetRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
}

// Context sets the context that will be used to send the request.
func (r *OrganizationGetRequest) Context(value context.Context) *OrganizationGetRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *OrganizationGetRequest) Timeout(value time.Duration) *OrganizationGetRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *OrganizationGetRequest) Deadline(value time.Time) *OrganizationGetRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *OrganizationGetRequest) Parameter(name string, value interface{}) *OrganizationGetRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *OrganizationGetRequest) Header(name string, value interface{}) *OrganizationGetRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *OrganizationGetRequest) Send() (result *OrganizationGetResponse, err error) {
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
	result = new(OrganizationGetResponse)
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

// OrganizationGetResponse is the response for the 'get' method.
type OrganizationGetResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Organization
}

// Status returns the response status code.
func (r *OrganizationGetResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *OrganizationGetResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *OrganizationGetResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *OrganizationGetResponse) Body() *Organization {
	return r.body
}

// unmarshal is the method used internally to unmarshal responses to the
// 'get' method.
func (r *OrganizationGetResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(organizationData)
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

// OrganizationUpdateRequest is the request for the 'update' method.
type OrganizationUpdateRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
	body      *Organization
}

// Context sets the context that will be used to send the request.
func (r *OrganizationUpdateRequest) Context(value context.Context) *OrganizationUpdateRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *OrganizationUpdateRequest) Timeout(value time.Duration) *OrganizationUpdateRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *OrganizationUpdateRequest) Deadline(value time.Time) *OrganizationUpdateRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *OrganizationUpdateRequest) Parameter(name string, value interface{}) *OrganizationUpdateRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *OrganizationUpdateRequest) Header(name string, value interface{}) *OrganizationUpdateRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Body sets the value of the 'body' parameter.
//
//
func (r *OrganizationUpdateRequest) Body(value *Organization) *OrganizationUpdateRequest {
	r.body = value
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *OrganizationUpdateRequest) Send() (result *OrganizationUpdateResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.CopyHeader(r.header)
	buffer := new(bytes.Buffer)
	err = r.marshal(buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: http.MethodPatch,
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = new(OrganizationUpdateResponse)
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

// marshall is the method used internally to marshal requests for the
// 'update' method.
func (r *OrganizationUpdateRequest) marshal(writer io.Writer) error {
	var err error
	encoder := json.NewEncoder(writer)
	data, err := r.body.wrap()
	if err != nil {
		return err
	}
	err = encoder.Encode(data)
	return err
}

// OrganizationUpdateResponse is the response for the 'update' method.
type OrganizationUpdateResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *Organization
}

// Status returns the response status code.
func (r *OrganizationUpdateResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *OrganizationUpdateResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *OrganizationUpdateResponse) Error() *errors.Error {
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *OrganizationUpdateResponse) Body() *Organization {
	return r.body
}

// unmarshal is the method used internally to unmarshal responses to the
// 'update' method.
func (r *OrganizationUpdateResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(organizationData)
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
