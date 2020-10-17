package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/go-kit/kit/endpoint"
    "github.com/gorilla/mux"
    "net/http"
)

func makeCreateRateEndpoint(s RateService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(CreateRateRequest)
        msg, err := s.CreateRate(ctx, req.rate)
        return CreateRateResponse{Msg: msg, Err: err}, nil
    }
}
func makeGetRateByIdEndpoint(s RateService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetRateByIdRequest)
        rateDetails, err := s.GetRateById(ctx, req.Id)
        if err != nil {
            return GetRateByIdResponse{Rate: rateDetails, Err: "Id not found"}, nil
        }
        return GetRateByIdResponse{Rate: rateDetails, Err: ""}, nil
    }
}
func makeDeleteRateEndpoint(s RateService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(DeleteRateRequest)
        msg, err := s.DeleteRate(ctx, req.Rateid)
        if err != nil {
            return DeleteRateResponse{Msg: msg, Err: err}, nil
        }
        return DeleteRateResponse{Msg: msg, Err: nil}, nil
    }
}
func makeUpdateRateendpoint(s RateService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UpdateRateRequest)
        msg, err := s.UpdateRate(ctx, req.rate)
        return msg, err
    }
}

func decodeCreateRateRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req CreateRateRequest
    fmt.Println("-------->>>>into Decoding")
    if err := json.NewDecoder(r.Body).Decode(&req.rate); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeGetRateByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetRateByIdRequest
    fmt.Println("-------->>>>into GetById Decoding")
    vars := mux.Vars(r)
    req = GetRateByIdRequest{
        Id: vars["rateid"],
    }
    return req, nil
}
func decodeDeleteRateRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Delete Decoding")
    var req DeleteRateRequest
    vars := mux.Vars(r)
    req = DeleteRateRequest{
        Rateid: vars["rateid"],
    }
    return req, nil
}
func decodeUpdateRateRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Update Decoding")
    var req UpdateRateRequest
    if err := json.NewDecoder(r.Body).Decode(&req.rate); err != nil {
        return nil, err
    }
    return req, nil
}

func encodeRateResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    fmt.Println("into Encoding <<<<<<----------------")
    return json.NewEncoder(w).Encode(response)
}

type (
    CreateRateRequest struct {
        rate Rate
    }
    CreateRateResponse struct {
        Msg string `json:"msg"`
        Err error  `json:"error,omitempty"`
    }
    GetRateByIdRequest struct {
        Id string `json:"rateid"`
    }
    GetRateByIdResponse struct {
        Rate interface{} `json:"rate,omitempty"`
        Err  string      `json:"error,omitempty"`
    }

    DeleteRateRequest struct {
        Rateid string `json:"rateid"`
    }

    DeleteRateResponse struct {
        Msg string `json:"response"`
        Err error  `json:"error,omitempty"`
    }
    UpdateRateRequest struct {
        rate Rate
    }
    UpdateRateResponse struct {
        Msg string `json:"status,omitempty"`
        Err error  `json:"error,omitempty"`
    }
)