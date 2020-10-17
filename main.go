package main

import (
    "github.com/go-kit/kit/log"
    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/gorilla/mux"
    "net/http"
    "os"
)

func main() {
    readDataFromTvShowCSV("tv_shows.csv")

    logger := log.NewLogfmtLogger(os.Stderr)

    r := mux.NewRouter()

    var svcTV TVshowService
    svcTV = NewServiceTVshow(logger)

    //TV show
    CreateTVshowHandler := httptransport.NewServer(
        makeCreateTVshowEndpoint(svcTV),
        decodeCreateTVshowRequest,
        encodeTVshowResponse,
    )
    GetByTVshowIdHandler := httptransport.NewServer(
        makeGetTVshowByIdEndpoint(svcTV),
        decodeGetTVshowByIdRequest,
        encodeTVshowResponse,
    )
    DeleteTVshowHandler := httptransport.NewServer(
        makeDeleteTVshowEndpoint(svcTV),
        decodeDeleteTVshowRequest,
        encodeTVshowResponse,
    )
    UpdateTVshowHandler := httptransport.NewServer(
        makeUpdateTVshowendpoint(svcTV),
        decodeUpdateTVshowRequest,
        encodeTVshowResponse,
    )
    GetYearsHandler := httptransport.NewServer(
        makeGetYearsEndpoint(svcTV),
        decodeGetYearsRequest,
        encodeTVshowResponse,
    )
    GetRatesHandler := httptransport.NewServer(
        makeGetRatesEndpoint(svcTV),
        decodeGetRatesRequest,
        encodeTVshowResponse,
    )
    GetAgesHandler := httptransport.NewServer(
        makeGetAgesEndpoint(svcTV),
        decodeGetAgesRequest,
        encodeTVshowResponse,
    )
    
    
    http.Handle("/", r)
    //Tv show
    http.Handle("/show", CreateTVshowHandler)
    http.Handle("/show/update", UpdateTVshowHandler)
    r.Handle("/show/{tvshowid}", GetByTVshowIdHandler).Methods("GET")
    r.Handle("/show/{tvshowid}", DeleteTVshowHandler).Methods("DELETE")
    r.Handle("/years/{yearid}", GetYearsHandler).Methods("GET")
    r.Handle("/rates/{rateid}", GetRatesHandler).Methods("GET")
    r.Handle("/ages/{ageid}", GetAgesHandler).Methods("GET")
    
    logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
    logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
    
}