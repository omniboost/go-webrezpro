package webrezpro

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-webrezpro/utils"
)

func (c *Client) NewGetTrialBalanceRequest() GetTrialBalanceRequest {
	return GetTrialBalanceRequest{
		client:      c,
		queryParams: c.NewGetTrialBalanceQueryParams(),
		pathParams:  c.NewGetTrialBalancePathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetTrialBalanceRequestBody(),
	}
}

type GetTrialBalanceRequest struct {
	client      *Client
	queryParams *GetTrialBalanceQueryParams
	pathParams  *GetTrialBalancePathParams
	method      string
	headers     http.Header
	requestBody GetTrialBalanceRequestBody
}

func (c *Client) NewGetTrialBalanceQueryParams() *GetTrialBalanceQueryParams {
	return &GetTrialBalanceQueryParams{}
}

type GetTrialBalanceQueryParams struct {
	Date             Date   `schema:"date"`
	DepartmentNumber string `schema:"department_number,omitempty"`
}

func (p GetTrialBalanceQueryParams) ToURLValues() (url.Values, error) {
	encoder := utils.NewSchemaEncoder()
	encoder.RegisterEncoder(Date{}, utils.EncodeSchemaMarshaler)
	encoder.RegisterEncoder(DateTime{}, utils.EncodeSchemaMarshaler)
	params := url.Values{}

	err := encoder.Encode(p, params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func (r *GetTrialBalanceRequest) QueryParams() *GetTrialBalanceQueryParams {
	return r.queryParams
}

func (c *Client) NewGetTrialBalancePathParams() *GetTrialBalancePathParams {
	return &GetTrialBalancePathParams{}
}

type GetTrialBalancePathParams struct {
}

func (p *GetTrialBalancePathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetTrialBalanceRequest) PathParams() *GetTrialBalancePathParams {
	return r.pathParams
}

func (r *GetTrialBalanceRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetTrialBalanceRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetTrialBalanceRequest) Method() string {
	return r.method
}

func (s *Client) NewGetTrialBalanceRequestBody() GetTrialBalanceRequestBody {
	return GetTrialBalanceRequestBody{}
}

type GetTrialBalanceRequestBody struct {
}

func (r *GetTrialBalanceRequest) RequestBody() *GetTrialBalanceRequestBody {
	return nil
}

func (r *GetTrialBalanceRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetTrialBalanceRequest) SetRequestBody(body GetTrialBalanceRequestBody) {
	r.requestBody = body
}

func (r *GetTrialBalanceRequest) NewResponseBody() *GetTrialBalanceResponseBody {
	return &GetTrialBalanceResponseBody{}
}

type GetTrialBalanceResponseBody TrialBalanceResponse

func (r *GetTrialBalanceRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("trial_balance", r.PathParams())
	return &u
}

func (r *GetTrialBalanceRequest) Do() (GetTrialBalanceResponseBody, error) {
	// Create http request
	req, err := r.client.NewRequest(nil, r)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	// Process query parameters
	err = utils.AddQueryParamsToRequest(r.QueryParams(), req, true)
	if err != nil {
		return *r.NewResponseBody(), err
	}

	responseBody := r.NewResponseBody()
	_, err = r.client.Do(req, responseBody)
	return *responseBody, err
}
