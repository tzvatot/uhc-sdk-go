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

// DashboardsClient is the client of the 'dashboards' resource.
//
// Manages the collection of dashboards.
type DashboardsClient struct {
	transport http.RoundTripper
	path      string
}

// NewDashboardsClient creates a new client for the 'dashboards'
// resource using the given transport to sned the requests and receive the
// responses.
func NewDashboardsClient(transport http.RoundTripper, path string) *DashboardsClient {
	client := new(DashboardsClient)
	client.transport = transport
	client.path = path
	return client
}

// List creates a request for the 'list' method.
//
// Retrieves a list of dashboards.
func (c *DashboardsClient) List() *DashboardsListRequest {
	request := new(DashboardsListRequest)
	request.transport = c.transport
	request.path = c.path
	return request
}

// Dashboard returns the target 'dashboard' resource for the given identifier.
//
// Reference to the resource that manages a specific dashboard.
func (c *DashboardsClient) Dashboard(id string) *DashboardClient {
	return NewDashboardClient(c.transport, path.Join(c.path, id))
}

// DashboardsListRequest is the request for the 'list' method.
type DashboardsListRequest struct {
	transport http.RoundTripper
	path      string
	context   context.Context
	cancel    context.CancelFunc
	query     url.Values
	header    http.Header
	page      *int
	size      *int
	search    *string
	total     *int
}

// Context sets the context that will be used to send the request.
func (r *DashboardsListRequest) Context(value context.Context) *DashboardsListRequest {
	r.context = value
	return r
}

// Timeout sets a timeout for the completete request.
func (r *DashboardsListRequest) Timeout(value time.Duration) *DashboardsListRequest {
	helpers.SetTimeout(&r.context, &r.cancel, value)
	return r
}

// Deadline sets a deadline for the completete request.
func (r *DashboardsListRequest) Deadline(value time.Time) *DashboardsListRequest {
	helpers.SetDeadline(&r.context, &r.cancel, value)
	return r
}

// Parameter adds a query parameter.
func (r *DashboardsListRequest) Parameter(name string, value interface{}) *DashboardsListRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *DashboardsListRequest) Header(name string, value interface{}) *DashboardsListRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Page sets the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *DashboardsListRequest) Page(value int) *DashboardsListRequest {
	r.page = &value
	return r
}

// Size sets the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *DashboardsListRequest) Size(value int) *DashboardsListRequest {
	r.size = &value
	return r
}

// Search sets the value of the 'search' parameter.
//
// Search criteria.
//
// The syntax of this parameter is similar to the syntax of the
// _where_ clause of an SQL statement, but using the names of
// the attributes of the dashboard instead of the names of the
// columns of a table. For example, in order to retrieve all the
// dashboards with a name starting with `my` the value should be:
//
// [source,sql]
// ----
// name like 'my%'
// ----
//
// If the parameter isn't provided, or if the value is empty,
// then all the dashboards that the user has permission to see
// will be returned.
func (r *DashboardsListRequest) Search(value string) *DashboardsListRequest {
	r.search = &value
	return r
}

// Total sets the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *DashboardsListRequest) Total(value int) *DashboardsListRequest {
	r.total = &value
	return r
}

// Send sends this request, waits for the response, and returns it.
func (r *DashboardsListRequest) Send() (result *DashboardsListResponse, err error) {
	query := helpers.CopyQuery(r.query)
	if r.page != nil {
		helpers.AddValue(&query, "page", *r.page)
	}
	if r.size != nil {
		helpers.AddValue(&query, "size", *r.size)
	}
	if r.search != nil {
		helpers.AddValue(&query, "search", *r.search)
	}
	if r.total != nil {
		helpers.AddValue(&query, "total", *r.total)
	}
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
	result = new(DashboardsListResponse)
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

// DashboardsListResponse is the response for the 'list' method.
type DashboardsListResponse struct {
	status int
	header http.Header
	err    *errors.Error
	page   *int
	size   *int
	total  *int
	items  *DashboardList
}

// Status returns the response status code.
func (r *DashboardsListResponse) Status() int {
	return r.status
}

// Header returns header of the response.
func (r *DashboardsListResponse) Header() http.Header {
	return r.header
}

// Error returns the response error.
func (r *DashboardsListResponse) Error() *errors.Error {
	return r.err
}

// Page returns the value of the 'page' parameter.
//
// Index of the requested page, where one corresponds to the first page.
//
// Default value is `1`.
func (r *DashboardsListResponse) Page() int {
	if r.page != nil {
		return *r.page
	}
	return 0
}

// Size returns the value of the 'size' parameter.
//
// Maximum number of items that will be contained in the returned page.
//
// Default value is `100`.
func (r *DashboardsListResponse) Size() int {
	if r.size != nil {
		return *r.size
	}
	return 0
}

// Total returns the value of the 'total' parameter.
//
// Total number of items of the collection that match the search criteria,
// regardless of the size of the page.
func (r *DashboardsListResponse) Total() int {
	if r.total != nil {
		return *r.total
	}
	return 0
}

// Items returns the value of the 'items' parameter.
//
// Retrieved list of dashboards.
func (r *DashboardsListResponse) Items() *DashboardList {
	return r.items
}

// unmarshal is the method used internally to unmarshal responses to the
// 'list' method.
func (r *DashboardsListResponse) unmarshal(reader io.Reader) error {
	var err error
	decoder := json.NewDecoder(reader)
	data := new(dashboardsListResponseData)
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

// dashboardsListResponseData is the structure used internally to unmarshal
// the response of the 'list' method.
type dashboardsListResponseData struct {
	Page  *int              "json:\"page,omitempty\""
	Size  *int              "json:\"size,omitempty\""
	Total *int              "json:\"total,omitempty\""
	Items dashboardListData "json:\"items,omitempty\""
}
