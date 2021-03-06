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

// LogsClient is the client of the 'logs' resource.
//
// Manages a collection of logs.
type LogsClient struct {
	transport http.RoundTripper
	path      string
}

// NewLogsClient creates a new client for the 'logs'
// resource using the given transport to sned the requests and receive the
// responses.
func NewLogsClient(transport http.RoundTripper, path string) *LogsClient {
	client := new(LogsClient)
	client.transport = transport
	client.path = path
	return client
}

// List creates a request for the 'list' method.
//
// Retrieves the list of clusters.
func (c *LogsClient) List() *LogsListRequest {
	request := new(LogsListRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// Log returns the target 'log' resource for the given identifier.
//
// Retursn a reference to the service that manages an specific log.
func (c *LogsClient) Log(id string) *LogClient {
	return NewLogClient(c.transport, path.Join(c.path, id))
}

// LogsListRequest is the request for the 'list' method.
type LogsListRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
}

// Context sets the context that will be used to send the request.
func (r *LogsListRequest) Context(value context.Context) *LogsListRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *LogsListRequest) Timeout(value time.Duration) *LogsListRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *LogsListRequest) Deadline(value time.Time) *LogsListRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *LogsListRequest) Parameter(name string, value interface{}) *LogsListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *LogsListRequest) Header(name string, value interface{}) *LogsListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *LogsListRequest) Send() (result *LogsListResponse, err error) {
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
	result = new(LogsListResponse)
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

// LogsListResponse is the response for the 'list' method.
type LogsListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *LogList
}

// Status returns the response status code.
func (r *LogsListResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *LogsListResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *LogsListResponse) Error() *errors.Error {
	return r.err
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
func (r *LogsListResponse) Page() int {
	if r.page != nil {
		return *r.page
	}
	return 0
}

// Size returns the value of the 'size' parameter.
//
// Number of items contained in the returned page.
func (r *LogsListResponse) Size() int {
	if r.size != nil {
		return *r.size
	}
	return 0
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection.
func (r *LogsListResponse) Total() int {
	if r.total != nil {
		return *r.total
	}
	return 0
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of logs.
func (r *LogsListResponse) Items() *LogList {
	return r.items
}

// unmarshal is the method used internally to unmarshal responses to the
// 'list' method.
func (r *LogsListResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(logsListResponseData)
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

// logsListResponseData is the structure used internally to unmarshal
// the response of the 'list' method.
type logsListResponseData struct {
	Page  *int        "json:\"page,omitempty\""
	Size  *int        "json:\"size,omitempty\""
	Total *int        "json:\"total,omitempty\""
	Items logListData "json:\"items,omitempty\""
}
