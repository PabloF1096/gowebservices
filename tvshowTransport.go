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
func makGetYearByTVshowIdEndpoint(s TVshowService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetYearbyTVshowIdRequest)
        resultYear, err := s.GetYearByTVshowId(ctx, req.Id)
        if err != nil {
            return GetYearbyTVshowIdResponse{Year: resultYear, Err: "Id not found"}, nil
        }
        return GetYearbyTVshowIdResponse{Year: resultYear, Err: ""}, nil
    }
}
func makGetRateByTVshowIdEndpoint(s TVshowService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetRateByTVshowIdRequest)
        resultRate, err := s.GetRateByTVshowId(ctx, req.Id)
        if err != nil {
            return GetRateByTVshowIdResponse{Rate: resultRate, Err: "Id not found"}, nil
        }
        return GetRateByTVshowIdResponse{Rate: resultRate, Err: ""}, nil
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
func makUpdateYearWithTVshowIdendpoint(s TVshowService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UpdateYearWithTVshowIdRequest)
        msg, err := s.UpdateYearWithTVshowId(ctx, req.tvshow,req.Id)
        return msg, err
    }
}

func makeDeleteYearWithTVshowEndpoint(s TVshowService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(DeleteYearWithTVshowRequest)
        msg, err := s.DeleteYearWithTVshow(ctx, req.TVshowid, req.IdYear)
        if err != nil {
            return DeleteYearWithTVshowResponse{Msg: msg, Err: err}, nil
        }
        return DeleteYearWithTVshowResponse{Msg: msg, Err: nil}, nil
    }
}
func makeUpdateRateWithTVshowIdendpoint(s TVshowService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UpdateRateWithTVshowIdRequest)
        msg, err := s.UpdateRateWithTVshowId(ctx, req.tvshow,req.Id)
        return msg, err
    }
}

func makeDeleteRateWithTVshowEndpoint(s TVshowService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(DeleteRateWithTVshowRequest)
        msg, err := s.DeleteRateWithTVshow(ctx, req.TVshowid, req.IdRate)
        if err != nil {
            return DeleteRateWithTVshowResponse{Msg: msg, Err: err}, nil
        }
        return DeleteRateWithTVshowResponse{Msg: msg, Err: nil}, nil
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
func decodeGetYearByTVshowIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetYearbyTVshowIdRequest
    fmt.Println("-------->>>>into GetYearbyTVshowId Decoding")
    vars := mux.Vars(r)
    req = GetYearbyTVshowIdRequest{
        Id: vars["tvshowid"],
    }
    return req, nil
}
func decodeGetRateByTVshowIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetRateByTVshowIdRequest
    fmt.Println("-------->>>>into GetRateByTVshowId Decoding")
    vars := mux.Vars(r)
    req = GetRateByTVshowIdRequest{
        Id: vars["tvshowid"],
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
func decodeUpdateYearWithTVshowIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Update Year With TVshowId Decoding")
    var req UpdateYearWithTVshowIdRequest
    vars := mux.Vars(r)
    req = UpdateYearWithTVshowIdRequest{
        Id: vars["tvshowid"],
    }
    if err := json.NewDecoder(r.Body).Decode(&req.tvshow); err != nil {
        return nil, err
    }
    return req, nil
}
func decodeDeleteYearWithTVshowRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Delete Year With TVshow Decoding")
    var req DeleteYearWithTVshowRequest
    vars := mux.Vars(r)
    req = DeleteYearWithTVshowRequest{
        TVshowid: vars["tvshowid"],
        IdYear:   vars["yearid"],
    }
    return req, nil
}
func decodeUpdateRateWithTVshowIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Update Rate With TVshowId Decoding")
    var req UpdateRateWithTVshowIdRequest
    vars := mux.Vars(r)
    req = UpdateRateWithTVshowIdRequest{
        Id: vars["tvshowid"],
    }
    if err := json.NewDecoder(r.Body).Decode(&req.tvshow); err != nil {
        return nil, err
    }
    return req, nil
}
func decodeDeleteRateWithTVshowRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Delete Rate With TVshow Decoding")
    var req DeleteRateWithTVshowRequest
    vars := mux.Vars(r)
    req = DeleteRateWithTVshowRequest{
        TVshowid: vars["tvshowid"],
        IdRate:   vars["rateid"],
    }
    return req, nil
}

func encodeTVshowResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    fmt.Println("into Encoding <<<<<<----------------")
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

    GetYearbyTVshowIdRequest struct {
        Id string `json:"tvshowid"`
    }
    GetYearbyTVshowIdResponse struct {
        Year interface{} `json:"year,omitempty"`
        Err  string      `json:"error,omitempty"`
    }
    GetRateByTVshowIdRequest struct {
        Id string `json:"tvshowid"`
    }
    GetRateByTVshowIdResponse struct {
        Rate interface{} `json:"rate,omitempty"`
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

    UpdateYearWithTVshowIdRequest struct {
        tvshow TVshow
        Id string `json:"tvshowid"`
    }
    UpdateYearWithTVshowIdResponse struct {
        Msg string `json:"status,omitempty"`
        Err error  `json:"error,omitempty"`
    }
    DeleteYearWithTVshowRequest struct {
        TVshowid string `json:"tvshowid"`
        IdYear string `json:"yearid"`
    }

    DeleteYearWithTVshowResponse struct {
        Msg string `json:"response"`
        Err error  `json:"error,omitempty"`
    }
    UpdateRateWithTVshowIdRequest struct {
        tvshow TVshow
        Id string `json:"tvshowid"`
    }
    UpdateRateWithTVshowIdResponse struct {
        Msg string `json:"status,omitempty"`
        Err error  `json:"error,omitempty"`
    }
    DeleteRateWithTVshowRequest struct {
        TVshowid string `json:"tvshowid"`
        IdRate string `json:"rateid"`
    }

    DeleteRateWithTVshowResponse struct {
        Msg string `json:"response"`
        Err error  `json:"error,omitempty"`
    }
)