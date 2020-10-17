package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/go-kit/kit/endpoint"
    "github.com/gorilla/mux"
    "net/http"
)

func makeCreateYearEndpoint(s YearService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(CreateYearRequest)
        msg, err := s.CreateYear(ctx, req.year)
        return CreateYearResponse{Msg: msg, Err: err}, nil
    }
}
func makeGetYearByIdEndpoint(s YearService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetYearByIdRequest)
        yearDetails, err := s.GetYearById(ctx, req.Id)
        if err != nil {
            return GetYearByIdResponse{Year: yearDetails, Err: "Id not found"}, nil
        }
        return GetYearByIdResponse{Year: yearDetails, Err: ""}, nil
    }
}
func makeDeleteYearEndpoint(s YearService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(DeleteYearRequest)
        msg, err := s.DeleteYear(ctx, req.Yearid)
        if err != nil {
            return DeleteYearResponse{Msg: msg, Err: err}, nil
        }
        return DeleteYearResponse{Msg: msg, Err: nil}, nil
    }
}
func makeUpdateYearendpoint(s YearService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UpdateYearRequest)
        msg, err := s.UpdateYear(ctx, req.year)
        return msg, err
    }
}

func decodeCreateYearRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req CreateYearRequest
    fmt.Println("-------->>>>into Decoding")
    if err := json.NewDecoder(r.Body).Decode(&req.year); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeGetYearByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetYearByIdRequest
    fmt.Println("-------->>>>into GetById Decoding")
    vars := mux.Vars(r)
    req = GetYearByIdRequest{
        Id: vars["yearid"],
    }
    return req, nil
}
func decodeDeleteYearRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Delete Decoding")
    var req DeleteYearRequest
    vars := mux.Vars(r)
    req = DeleteYearRequest{
        Yearid: vars["yearid"],
    }
    return req, nil
}
func decodeUpdateYearRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Update Decoding")
    var req UpdateYearRequest
    if err := json.NewDecoder(r.Body).Decode(&req.year); err != nil {
        return nil, err
    }
    return req, nil
}

func encodeYearResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    fmt.Println("into Encoding <<<<<<----------------")
    return json.NewEncoder(w).Encode(response)
}

type (
    CreateYearRequest struct {
        year Year
    }
    CreateYearResponse struct {
        Msg string `json:"msg"`
        Err error  `json:"error,omitempty"`
    }
    GetYearByIdRequest struct {
        Id string `json:"yearid"`
    }
    GetYearByIdResponse struct {
        Year interface{} `json:"year,omitempty"`
        Err  string      `json:"error,omitempty"`
    }

    DeleteYearRequest struct {
        Yearid string `json:"yearid"`
    }

    DeleteYearResponse struct {
        Msg string `json:"response"`
        Err error  `json:"error,omitempty"`
    }
    UpdateYearRequest struct {
        year Year
    }
    UpdateYearResponse struct {
        Msg string `json:"status,omitempty"`
        Err error  `json:"error,omitempty"`
    }
)