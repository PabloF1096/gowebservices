package main

import (
    "github.com/go-kit/kit/log"
    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/gorilla/mux"
    "net/http"
    "os"
)

func main() {
    logger := log.NewLogfmtLogger(os.Stderr)

    r := mux.NewRouter()

    var svc BookService
	svc = NewService(logger)
	
	var svcAuthor AuthorService
	svcAuthor = NewServiceAuthor(logger)
	
	var svcPublisher PublisherService
    svcPublisher = NewServicePublisher(logger)

    // svc = loggingMiddleware{logger, svc}
    // svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	//Book
    CreateBookHandler := httptransport.NewServer(
        makeCreateBookEndpoint(svc),
        decodeCreateBookRequest,
        encodeResponse,
    )
    GetByBookIdHandler := httptransport.NewServer(
        makeGetBookByIdEndpoint(svc),
        decodeGetBookByIdRequest,
        encodeResponse,
    )
    DeleteBookHandler := httptransport.NewServer(
        makeDeleteBookEndpoint(svc),
        decodeDeleteBookRequest,
        encodeResponse,
    )
    UpdateBookHandler := httptransport.NewServer(
        makeUpdateBookendpoint(svc),
        decodeUpdateBookRequest,
        encodeResponse,
	)
	
	//Author
	CreateAuthorHandler := httptransport.NewServer(
        makeCreateAuthorEndpoint(svcAuthor),
        decodeCreateAuthorRequest,
        encodeAuthorResponse,
    )
    GetByAuthorIdHandler := httptransport.NewServer(
        makeGetAuthorByIdEndpoint(svcAuthor),
        decodeGetAuthorByIdRequest,
        encodeAuthorResponse,
    )
    DeleteAuthorHandler := httptransport.NewServer(
        makeDeleteAuthorEndpoint(svcAuthor),
        decodeDeleteAuthorRequest,
        encodeAuthorResponse,
    )
    UpdateAuthorHandler := httptransport.NewServer(
        makeUpdateAuthorendpoint(svcAuthor),
        decodeUpdateAuthorRequest,
        encodeAuthorResponse,
	)

	//Publisher
	CreatePublisherHandler := httptransport.NewServer(
        makeCreatePublisherEndpoint(svcPublisher),
        decodeCreatePublisherRequest,
        encodePublisherResponse,
    )
    GetByPublisherIdHandler := httptransport.NewServer(
        makeGetPublisherByIdEndpoint(svcPublisher),
        decodeGetPublisherByIdRequest,
        encodePublisherResponse,
    )
    DeletePublisherHandler := httptransport.NewServer(
        makeDeletePublisherEndpoint(svcPublisher),
        decodeDeletePublisherRequest,
        encodePublisherResponse,
    )
    UpdatePublisherHandler := httptransport.NewServer(
        makeUpdatePublisherendpoint(svcPublisher),
        decodeUpdatePublisherRequest,
        encodePublisherResponse,
	)


    http.Handle("/", r)
    http.Handle("/book", CreateBookHandler)
    http.Handle("/book/update", UpdateBookHandler)
    r.Handle("/book/{bookid}", GetByBookIdHandler).Methods("GET")
	r.Handle("/book/{bookid}", DeleteBookHandler).Methods("DELETE")
	//Author
	http.Handle("/author", CreateAuthorHandler)
    http.Handle("/author/update", UpdateAuthorHandler)
    r.Handle("/author/{authorid}", GetByAuthorIdHandler).Methods("GET")
	r.Handle("/author/{authorid}", DeleteAuthorHandler).Methods("DELETE")
	//Publisher
	http.Handle("/publisher", CreatePublisherHandler)
    http.Handle("/publisher/update", UpdatePublisherHandler)
    r.Handle("/publisher/{publisherid}", GetByPublisherIdHandler).Methods("GET")
	r.Handle("/publisher/{publisherid}", DeletePublisherHandler).Methods("DELETE")

    // http.Handle("/metrics", promhttp.Handler())
    logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
    logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
}