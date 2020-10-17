package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/go-kit/kit/endpoint"
    "github.com/gorilla/mux"
    "net/http"
)

func makeCreateTVshowEndpoint(s TVshowService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreateTVshowRequest)
        msg, err := s.CreateTVshow(ctx, req.tvshow)
        return CreateTVshowResponse{Msg: msg, Err: err}, nil
    }
}
func makeGetTVshowByIdEndpoint(s TVshowService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetTVshowByIdRequest)
        tvshowDetails, err := s.GetTVshowById(ctx, req.Id)
        if err != nil {
            return GetTVshowByIdResponse{TVshow: tvshowDetails, Err: "Id not found"}, nil
        }
        return GetTVshowByIdResponse{TVshow: tvshowDetails, Err: ""}, nil
    }
}
func makeGetYearsEndpoint(s TVshowService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetYearsRequest)
        tvshowDetails, err := s.GetYears(ctx, req.Id)
        if err != nil {
            return GetYearsResponse{TVshow: tvshowDetails, Err: "Id not found"}, nil
        }
        return GetYearsResponse{TVshow: tvshowDetails, Err: ""}, nil
    }
}
func makeGetRatesEndpoint(s TVshowService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRatesRequest)
        tvshowDetails, err := s.GetRates(ctx, req.Id)
        if err != nil {
            return GetRatesResponse{TVshow: tvshowDetails, Err: "Id not found"}, nil
        }
        return GetRatesResponse{TVshow: tvshowDetails, Err: ""}, nil
    }
}
func makeGetAgesEndpoint(s TVshowService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetAgesRequest)
        tvshowDetails, err := s.GetAges(ctx, req.Id)
        if err != nil {
            return GetAgesResponse{TVshow: tvshowDetails, Err: "Id not found"}, nil
        }
        return GetAgesResponse{TVshow: tvshowDetails, Err: ""}, nil
    }
}

func makeDeleteTVshowEndpoint(s TVshowService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(DeleteTVshowRequest)
        msg, err := s.DeleteTVshow(ctx, req.TVshowid)
        if err != nil {
            return DeleteTVshowResponse{Msg: msg, Err: err}, nil
        }
        return DeleteTVshowResponse{Msg: msg, Err: nil}, nil
    }
}
func makeUpdateTVshowendpoint(s TVshowService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UpdateTVshowRequest)
        msg, err := s.UpdateTVshow(ctx, req.tvshow)
        return msg, err
    }
}

func decodeCreateTVshowRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req CreateTVshowRequest
    fmt.Println("-------->>>>into Decoding")
    if err := json.NewDecoder(r.Body).Decode(&req.tvshow); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeGetTVshowByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetTVshowByIdRequest
    fmt.Println("-------->>>>into GetById Decoding")
    vars := mux.Vars(r)
    req = GetTVshowByIdRequest{
        Id: vars["tvshowid"],
    }
    return req, nil
}
func decodeGetYearsRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetYearsRequest
    fmt.Println("-------->>>>into Get Years Decoding")
    vars := mux.Vars(r)
    req = GetYearsRequest{
        Id: vars["yearid"],
    }
    return req, nil
}
func decodeGetRatesRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetRatesRequest
    fmt.Println("-------->>>>into Get Rates Decoding")
    vars := mux.Vars(r)
    req = GetRatesRequest{
        Id: vars["rateid"],
    }
    return req, nil
}
func decodeGetAgesRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetAgesRequest
    fmt.Println("-------->>>>into Get Ages Decoding")
    vars := mux.Vars(r)
    req = GetAgesRequest{
        Id: vars["ageid"],
    }
    return req, nil
}

func decodeDeleteTVshowRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Delete Decoding")
    var req DeleteTVshowRequest
    vars := mux.Vars(r)
    req = DeleteTVshowRequest{
        TVshowid: vars["tvshowid"],
    }
    return req, nil
}
func decodeUpdateTVshowRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Update Decoding")
    var req UpdateTVshowRequest
    if err := json.NewDecoder(r.Body).Decode(&req.tvshow); err != nil {
        return nil, err
    }
    return req, nil
}

func encodeTVshowResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    fmt.Println("into Encoding <<<<<<----------------")
    writeDataTvShowCSV("tv_shows.csv")
    return json.NewEncoder(w).Encode(response)
}

type (
    CreateTVshowRequest struct {
        tvshow TVshow
    }
    CreateTVshowResponse struct {
        Msg string `json:"msg"`
        Err error  `json:"error,omitempty"`
    }
    GetTVshowByIdRequest struct {
        Id string `json:"tvshowid"`
    }
    GetTVshowByIdResponse struct {
        TVshow interface{} `json:"tvshow,omitempty"`
        Err  string      `json:"error,omitempty"`
    }
    GetYearsRequest struct {
        Id string `json:"yearid"`
    }
    GetYearsResponse struct {
        TVshow interface{} `json:"tvshow,omitempty"`
        Err  string      `json:"error,omitempty"`
    }
    GetRatesRequest struct {
        Id string `json:"rateid"`
    }
    GetRatesResponse struct {
        TVshow interface{} `json:"tvshow,omitempty"`
        Err  string      `json:"error,omitempty"`
    }
    GetAgesRequest struct {
        Id string `json:"ageid"`
    }
    GetAgesResponse struct {
        TVshow interface{} `json:"tvshow,omitempty"`
        Err  string      `json:"error,omitempty"`
    }

    DeleteTVshowRequest struct {
        TVshowid string `json:"tvshowid"`
    }

    DeleteTVshowResponse struct {
        Msg string `json:"response"`
        Err error  `json:"error,omitempty"`
    }
    UpdateTVshowRequest struct {
        tvshow TVshow
    }
    UpdateTVshowResponse struct {
        Msg string `json:"status,omitempty"`
        Err error  `json:"error,omitempty"`
    }
)