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

package v1 // github.com/openshift-online/ocm-sdk-go/accountsmgmt/v1

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	jsoniter "github.com/json-iterator/go"
	"github.com/openshift-online/ocm-sdk-go/errors"
	"github.com/openshift-online/ocm-sdk-go/helpers"
)

// PullSecretsClient is the client of the 'pull_secrets' resource.
//
// Manages pull secrets.
type PullSecretsClient struct {
	transport http.RoundTripper
	path      string
	metric    string
}

// NewPullSecretsClient creates a new client for the 'pull_secrets'
// resource using the given transport to send the requests and receive the
// responses.
func NewPullSecretsClient(transport http.RoundTripper, path string, metric string) *PullSecretsClient {
	return &PullSecretsClient{
		transport: transport,
		path:      path,
		metric:    metric,
	}
}

// Delete creates a request for the 'delete' method.
//
// Deletes the pull secret
func (c *PullSecretsClient) Delete() *PullSecretsDeleteRequest {
	return &PullSecretsDeleteRequest{
		transport: c.transport,
		path:      c.path,
		metric:    c.metric,
	}
}

// Post creates a request for the 'post' method.
//
// Returns access token generated from registries in docker format.
func (c *PullSecretsClient) Post() *PullSecretsPostRequest {
	return &PullSecretsPostRequest{
		transport: c.transport,
		path:      c.path,
		metric:    c.metric,
	}
}

// PullSecretsDeleteRequest is the request for the 'delete' method.
type PullSecretsDeleteRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
}

// Parameter adds a query parameter.
func (r *PullSecretsDeleteRequest) Parameter(name string, value interface{}) *PullSecretsDeleteRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *PullSecretsDeleteRequest) Header(name string, value interface{}) *PullSecretsDeleteRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *PullSecretsDeleteRequest) Send() (result *PullSecretsDeleteResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *PullSecretsDeleteRequest) SendContext(ctx context.Context) (result *PullSecretsDeleteResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.SetHeader(r.header, r.metric)
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "DELETE",
		URL:    uri,
		Header: header,
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &PullSecretsDeleteResponse{}
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

// PullSecretsDeleteResponse is the response for the 'delete' method.
type PullSecretsDeleteResponse struct {
	status int
	header http.Header
	err    *errors.Error
}

// Status returns the response status code.
func (r *PullSecretsDeleteResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *PullSecretsDeleteResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *PullSecretsDeleteResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// PullSecretsPostRequest is the request for the 'post' method.
type PullSecretsPostRequest struct {
	transport http.RoundTripper
	path      string
	metric    string
	query     url.Values
	header    http.Header
	request   *PullSecretsRequest
}

// Parameter adds a query parameter.
func (r *PullSecretsPostRequest) Parameter(name string, value interface{}) *PullSecretsPostRequest {
	helpers.AddValue(&r.query, name, value)
	return r
}

// Header adds a request header.
func (r *PullSecretsPostRequest) Header(name string, value interface{}) *PullSecretsPostRequest {
	helpers.AddHeader(&r.header, name, value)
	return r
}

// Request sets the value of the 'request' parameter.
//
//
func (r *PullSecretsPostRequest) Request(value *PullSecretsRequest) *PullSecretsPostRequest {
	r.request = value
	return r
}

// Send sends this request, waits for the response, and returns it.
//
// This is a potentially lengthy operation, as it requires network communication.
// Consider using a context and the SendContext method.
func (r *PullSecretsPostRequest) Send() (result *PullSecretsPostResponse, err error) {
	return r.SendContext(context.Background())
}

// SendContext sends this request, waits for the response, and returns it.
func (r *PullSecretsPostRequest) SendContext(ctx context.Context) (result *PullSecretsPostResponse, err error) {
	query := helpers.CopyQuery(r.query)
	header := helpers.SetHeader(r.header, r.metric)
	buffer := &bytes.Buffer{}
	err = writePullSecretsPostRequest(r, buffer)
	if err != nil {
		return
	}
	uri := &url.URL{
		Path:     r.path,
		RawQuery: query.Encode(),
	}
	request := &http.Request{
		Method: "POST",
		URL:    uri,
		Header: header,
		Body:   ioutil.NopCloser(buffer),
	}
	if ctx != nil {
		request = request.WithContext(ctx)
	}
	response, err := r.transport.RoundTrip(request)
	if err != nil {
		return
	}
	defer response.Body.Close()
	result = &PullSecretsPostResponse{}
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
	err = readPullSecretsPostResponse(result, response.Body)
	if err != nil {
		return
	}
	return
}

// marshall is the method used internally to marshal requests for the
// 'post' method.
func (r *PullSecretsPostRequest) marshal(writer io.Writer) error {
	stream := helpers.NewStream(writer)
	r.stream(stream)
	return stream.Error
}
func (r *PullSecretsPostRequest) stream(stream *jsoniter.Stream) {
}

// PullSecretsPostResponse is the response for the 'post' method.
type PullSecretsPostResponse struct {
	status int
	header http.Header
	err    *errors.Error
	body   *AccessToken
}

// Status returns the response status code.
func (r *PullSecretsPostResponse) Status() int {
	if r == nil {
		return 0
	}
	return r.status
}

// Header returns header of the response.
func (r *PullSecretsPostResponse) Header() http.Header {
	if r == nil {
		return nil
	}
	return r.header
}

// Error returns the response error.
func (r *PullSecretsPostResponse) Error() *errors.Error {
	if r == nil {
		return nil
	}
	return r.err
}

// Body returns the value of the 'body' parameter.
//
//
func (r *PullSecretsPostResponse) Body() *AccessToken {
	if r == nil {
		return nil
	}
	return r.body
}

// GetBody returns the value of the 'body' parameter and
// a flag indicating if the parameter has a value.
//
//
func (r *PullSecretsPostResponse) GetBody() (value *AccessToken, ok bool) {
	ok = r != nil && r.body != nil
	if ok {
		value = r.body
	}
	return
}
