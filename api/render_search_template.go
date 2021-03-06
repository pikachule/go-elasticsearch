// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package api

import (
	"net/http"

	"github.com/elastic/go-elasticsearch/transport"
	"github.com/elastic/go-elasticsearch/util"
)

// RenderSearchTemplateOption is a non-required RenderSearchTemplate option that gets applied to an HTTP request.
type RenderSearchTemplateOption func(r *transport.Request)

// WithRenderSearchTemplateID - the id of the stored search template.
func WithRenderSearchTemplateID(id string) RenderSearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithRenderSearchTemplateBody - the search definition template and its params.
func WithRenderSearchTemplateBody(body map[string]interface{}) RenderSearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithRenderSearchTemplateErrorTrace - include the stack trace of returned errors.
func WithRenderSearchTemplateErrorTrace(errorTrace bool) RenderSearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithRenderSearchTemplateFilterPath - a comma-separated list of filters used to reduce the respone.
func WithRenderSearchTemplateFilterPath(filterPath []string) RenderSearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithRenderSearchTemplateHuman - return human readable values for statistics.
func WithRenderSearchTemplateHuman(human bool) RenderSearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithRenderSearchTemplateIgnore - ignores the specified HTTP status codes.
func WithRenderSearchTemplateIgnore(ignore []int) RenderSearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithRenderSearchTemplatePretty - pretty format the returned JSON response.
func WithRenderSearchTemplatePretty(pretty bool) RenderSearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// WithRenderSearchTemplateSourceParam - the URL-encoded request definition. Useful for libraries that do not accept a request body for non-POST requests.
func WithRenderSearchTemplateSourceParam(sourceParam string) RenderSearchTemplateOption {
	return func(r *transport.Request) {
	}
}

// RenderSearchTemplate - see http://www.elasticsearch.org/guide/en/elasticsearch/reference/5.x/search-template.html for more info.
//
// options: optional parameters.
func (a *API) RenderSearchTemplate(options ...RenderSearchTemplateOption) (*RenderSearchTemplateResponse, error) {
	req := a.transport.NewRequest("GET")
	for _, option := range options {
		option(req)
	}
	resp, err := a.transport.Do(req)
	return &RenderSearchTemplateResponse{resp}, err
}

// RenderSearchTemplateResponse is the response for RenderSearchTemplate.
type RenderSearchTemplateResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}

// DecodeBody decodes the JSON body of the HTTP response.
func (r *RenderSearchTemplateResponse) DecodeBody() (util.MapStr, error) {
	return transport.DecodeResponseBody(r.Response)
}
