package main

import (
    "context"
    "encoding/json"
    "fmt"
    "github.com/go-kit/kit/endpoint"
    "github.com/gorilla/mux"
    "net/http"
)

func makeCreatePublisherEndpoint(s PublisherService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(CreatePublisherRequest)
        msg, err := s.CreatePublisher(ctx, req.publisher)
        return CreatePublisherResponse{Msg: msg, Err: err}, nil
    }
}
func makeGetPublisherByIdEndpoint(s PublisherService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(GetPublisherByIdRequest)
        publisherDetails, err := s.GetPublisherById(ctx, req.Id)
        if err != nil {
            return GetPublisherByIdResponse{Publisher: publisherDetails, Err: "Id not found"}, nil
        }
        return GetPublisherByIdResponse{Publisher: publisherDetails, Err: ""}, nil
    }
}
func makeDeletePublisherEndpoint(s PublisherService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(DeletePublisherRequest)
        msg, err := s.DeletePublisher(ctx, req.Publisherid)
        if err != nil {
            return DeletePublisherResponse{Msg: msg, Err: err}, nil
        }
        return DeletePublisherResponse{Msg: msg, Err: nil}, nil
    }
}
func makeUpdatePublisherendpoint(s PublisherService) endpoint.Endpoint {
    return func(ctx context.Context, request interface{}) (interface{}, error) {
        req := request.(UpdatePublisherRequest)
        msg, err := s.UpdatePublisher(ctx, req.publisher)
        return msg, err
    }
}

func decodeCreatePublisherRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req CreatePublisherRequest
    fmt.Println("-------->>>>into Decoding")
    if err := json.NewDecoder(r.Body).Decode(&req.publisher); err != nil {
        return nil, err
    }
    return req, nil
}

func decodeGetPublisherByIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
    var req GetPublisherByIdRequest
    fmt.Println("-------->>>>into GetById Decoding")
    vars := mux.Vars(r)
    req = GetPublisherByIdRequest{
        Id: vars["publisherid"],
    }
    return req, nil
}
func decodeDeletePublisherRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Delete Decoding")
    var req DeletePublisherRequest
    vars := mux.Vars(r)
    req = DeletePublisherRequest{
        Publisherid: vars["publisherid"],
    }
    return req, nil
}
func decodeUpdatePublisherRequest(_ context.Context, r *http.Request) (interface{}, error) {
    fmt.Println("-------->>>> Into Update Decoding")
    var req UpdatePublisherRequest
    if err := json.NewDecoder(r.Body).Decode(&req.publisher); err != nil {
        return nil, err
    }
    return req, nil
}

func encodePublisherResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
    w.Header().Set("Content-Type", "application/json; charset=utf-8")
    fmt.Println("into Encoding <<<<<<----------------")
    return json.NewEncoder(w).Encode(response)
}

type (
    CreatePublisherRequest struct {
        publisher Publisher
    }
    CreatePublisherResponse struct {
        Msg string `json:"msg"`
        Err error  `json:"error,omitempty"`
    }
    GetPublisherByIdRequest struct {
        Id string `json:"publisherid"`
    }
    GetPublisherByIdResponse struct {
        Publisher interface{} `json:"publisher,omitempty"`
        Err  string      `json:"error,omitempty"`
    }

    DeletePublisherRequest struct {
        Publisherid string `json:"publisherid"`
    }

    DeletePublisherResponse struct {
        Msg string `json:"response"`
        Err error  `json:"error,omitempty"`
    }
    UpdatePublisherRequest struct {
        publisher Publisher
    }
    UpdatePublisherResponse struct {
        Msg string `json:"status,omitempty"`
        Err error  `json:"error,omitempty"`
    }
)