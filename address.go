package apidq

import (
	"context"
	"net/http"

	"github.com/nikitaksv/apidq-client-go/dto/address"
)

type AddressService service

func (s AddressService) prepareCtx(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKeyService, "address")
}

// Clean Стандартизация адреса
func (s AddressService) Clean(ctx context.Context, req *address.CleanRequest) (*address.CleanResponse, *http.Response, error) {
	ctx = s.prepareCtx(ctx)
	r, err := s.client.newRequest(ctx, http.MethodPost, "/v1/clean/address", contentTypeJSON, req)
	if err != nil {
		return nil, nil, err
	}

	rsp := &address.CleanResponse{}
	httpRsp, err := s.client.do(ctx, r, rsp)
	if err != nil {
		return nil, httpRsp, err
	}
	return rsp, httpRsp, nil
}

// CleanIqdq Стандартизация адреса в формате старого API
func (s AddressService) CleanIqdq(ctx context.Context, req *address.CleanRequest) (*address.CleanIqdqResponse, *http.Response, error) {
	ctx = s.prepareCtx(ctx)
	r, err := s.client.newRequest(ctx, http.MethodPost, "/v1/clean/address/iqdq", contentTypeJSON, req)
	if err != nil {
		return nil, nil, err
	}

	rsp := &address.CleanIqdqResponse{}
	httpRsp, err := s.client.do(ctx, r, rsp)
	if err != nil {
		return nil, httpRsp, err
	}
	return rsp, httpRsp, nil
}

// Suggest Подсказки адреса
func (s AddressService) Suggest(ctx context.Context, req *address.SuggestRequest) (*address.SuggestResponse, *http.Response, error) {
	ctx = s.prepareCtx(ctx)
	r, err := s.client.newRequest(ctx, http.MethodPost, "/v1/suggest/address", contentTypeJSON, req)
	if err != nil {
		return nil, nil, err
	}

	rsp := &address.SuggestResponse{}
	httpRsp, err := s.client.do(ctx, r, rsp)
	if err != nil {
		return nil, httpRsp, err
	}
	return rsp, httpRsp, nil
}
