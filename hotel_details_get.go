package webrezpro

import (
	"net/http"
	"net/url"

	"github.com/omniboost/go-webrezpro/utils"
)

func (c *Client) NewGetHotelDetailsRequest() GetHotelDetailsRequest {
	return GetHotelDetailsRequest{
		client:      c,
		queryParams: c.NewGetHotelDetailsQueryParams(),
		pathParams:  c.NewGetHotelDetailsPathParams(),
		method:      http.MethodGet,
		headers:     http.Header{},
		requestBody: c.NewGetHotelDetailsRequestBody(),
	}
}

type GetHotelDetailsRequest struct {
	client      *Client
	queryParams *GetHotelDetailsQueryParams
	pathParams  *GetHotelDetailsPathParams
	method      string
	headers     http.Header
	requestBody GetHotelDetailsRequestBody
}

func (c *Client) NewGetHotelDetailsQueryParams() *GetHotelDetailsQueryParams {
	return &GetHotelDetailsQueryParams{}
}

type GetHotelDetailsQueryParams struct{}

func (p GetHotelDetailsQueryParams) ToURLValues() (url.Values, error) {
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

func (r *GetHotelDetailsRequest) QueryParams() *GetHotelDetailsQueryParams {
	return r.queryParams
}

func (c *Client) NewGetHotelDetailsPathParams() *GetHotelDetailsPathParams {
	return &GetHotelDetailsPathParams{}
}

type GetHotelDetailsPathParams struct {
}

func (p *GetHotelDetailsPathParams) Params() map[string]string {
	return map[string]string{}
}

func (r *GetHotelDetailsRequest) PathParams() *GetHotelDetailsPathParams {
	return r.pathParams
}

func (r *GetHotelDetailsRequest) PathParamsInterface() PathParams {
	return r.pathParams
}

func (r *GetHotelDetailsRequest) SetMethod(method string) {
	r.method = method
}

func (r *GetHotelDetailsRequest) Method() string {
	return r.method
}

func (s *Client) NewGetHotelDetailsRequestBody() GetHotelDetailsRequestBody {
	return GetHotelDetailsRequestBody{}
}

type GetHotelDetailsRequestBody struct {
}

func (r *GetHotelDetailsRequest) RequestBody() *GetHotelDetailsRequestBody {
	return nil
}

func (r *GetHotelDetailsRequest) RequestBodyInterface() interface{} {
	return nil
}

func (r *GetHotelDetailsRequest) SetRequestBody(body GetHotelDetailsRequestBody) {
	r.requestBody = body
}

func (r *GetHotelDetailsRequest) NewResponseBody() *GetHotelDetailsResponseBody {
	return &GetHotelDetailsResponseBody{}
}

type GetHotelDetailsResponseBody HotelDetails

func (r *GetHotelDetailsRequest) URL() *url.URL {
	u := r.client.GetEndpointURL("hotel_details", r.PathParams())
	return &u
}

func (r *GetHotelDetailsRequest) Do() (GetHotelDetailsResponseBody, error) {
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
