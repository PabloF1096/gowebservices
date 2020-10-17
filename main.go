package main

import (
    "github.com/go-kit/kit/log"
    httptransport "github.com/go-kit/kit/transport/http"
    "github.com/gorilla/mux"
    "net/http"
    "os"
)

func handler(writer http.ResponseWriter, request *http.Request) {
    var err error
    // readData("tv_shows.csv")
    writeDataTvShowCSV("tv_shows.csv")
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }

    writeDataToRateCSV("rate.csv")
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
    writeDataToYearCSV("year.csv")
    if err != nil {
        http.Error(writer, err.Error(), http.StatusInternalServerError)
        return
    }
}

func main() {
    readDataFromTvShowCSV("tv_shows.csv")
    readDataFromRateCSV("rate.csv")
    readDataFromYearCSV("year.csv")
    // disyplayCsvData()
    

    logger := log.NewLogfmtLogger(os.Stderr)

    r := mux.NewRouter()

    var svcTV TVshowService
    svcTV = NewServiceTVshow(logger)

    var svRate RateService
    svRate = NewServiceRate(logger)
    
    var svYear YearService
	svYear = NewServiceYear(logger)

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
    GetYearByTVshowId := httptransport.NewServer(
        makGetYearByTVshowIdEndpoint(svcTV),
        decodeGetYearByTVshowIdRequest,
        encodeTVshowResponse,
    )
    GetRateByTVshowId := httptransport.NewServer(
        makGetRateByTVshowIdEndpoint(svcTV),
        decodeGetRateByTVshowIdRequest,
        encodeTVshowResponse,
    )
    UpdateYearWithTVshowId := httptransport.NewServer(
        makUpdateYearWithTVshowIdendpoint(svcTV),
        decodeUpdateYearWithTVshowIdRequest,
        encodeTVshowResponse,
    )
    DelteYearWithTVshowId := httptransport.NewServer(
        makeDeleteYearWithTVshowEndpoint(svcTV),
        decodeDeleteYearWithTVshowRequest,
        encodeTVshowResponse,
    )
    UpdateRateWithTVshowId := httptransport.NewServer(
        makeUpdateRateWithTVshowIdendpoint(svcTV),
        decodeUpdateRateWithTVshowIdRequest,
        encodeTVshowResponse,
    )
    DelteRateWithTVshowId := httptransport.NewServer(
        makeDeleteRateWithTVshowEndpoint(svcTV),
        decodeDeleteRateWithTVshowRequest,
        encodeTVshowResponse,
    )

    //Rate
    CreateRateHandler := httptransport.NewServer(
        makeCreateRateEndpoint(svRate),
        decodeCreateRateRequest,
        encodeRateResponse,
    )
    GetByRateIdHandler := httptransport.NewServer(
        makeGetRateByIdEndpoint(svRate),
        decodeGetRateByIdRequest,
        encodeRateResponse,
    )
    DeleteRateHandler := httptransport.NewServer(
        makeDeleteRateEndpoint(svRate),
        decodeDeleteRateRequest,
        encodeRateResponse,
    )
    UpdateRateHandler := httptransport.NewServer(
        makeUpdateRateendpoint(svRate),
        decodeUpdateRateRequest,
        encodeRateResponse,
    )
    
    //Year
    CreateYearHandler := httptransport.NewServer(
        makeCreateYearEndpoint(svYear),
        decodeCreateYearRequest,
        encodeYearResponse,
    )
    GetByYearIdHandler := httptransport.NewServer(
        makeGetYearByIdEndpoint(svYear),
        decodeGetYearByIdRequest,
        encodeYearResponse,
    )
    DeleteYearHandler := httptransport.NewServer(
        makeDeleteYearEndpoint(svYear),
        decodeDeleteYearRequest,
        encodeYearResponse,
    )
    UpdateYearHandler := httptransport.NewServer(
        makeUpdateYearendpoint(svYear),
        decodeUpdateYearRequest,
        encodeYearResponse,
    )

    
    http.Handle("/", r)
    http.HandleFunc("/write", handler)
    //Tv show
    http.Handle("/show", CreateTVshowHandler)
    http.Handle("/show/update", UpdateTVshowHandler)
    r.Handle("/show/{tvshowid}", GetByTVshowIdHandler).Methods("GET")
    r.Handle("/show/{tvshowid}", DeleteTVshowHandler).Methods("DELETE")

    r.Handle("/show/{tvshowid}/year", GetYearByTVshowId).Methods("GET")
    r.Handle("/show/{tvshowid}/year/{yearid}", GetYearByTVshowId).Methods("GET")
    r.Handle("/show/{tvshowid}/year", UpdateYearWithTVshowId).Methods("PUT")
    r.Handle("/show/{tvshowid}/year/{yearid}", DelteYearWithTVshowId).Methods("DELETE")

    r.Handle("/show/{tvshowid}/rate", GetRateByTVshowId).Methods("GET")
    r.Handle("/show/{tvshowid}/rate/{rateid}", GetRateByTVshowId).Methods("GET")
    r.Handle("/show/{tvshowid}/rate", UpdateRateWithTVshowId).Methods("PUT")
    r.Handle("/show/{tvshowid}/rate/{rateid}", DelteRateWithTVshowId).Methods("DELETE")
    //Rate
    http.Handle("/rate", CreateRateHandler)
    http.Handle("/rate/update", UpdateRateHandler)
    r.Handle("/rate/{rateid}", GetByRateIdHandler).Methods("GET")
    r.Handle("/rate/{rateid}", DeleteRateHandler).Methods("DELETE")
    //Year
    http.Handle("/year", CreateYearHandler)
    http.Handle("/year/update", UpdateYearHandler)
    r.Handle("/year/{yearid}", GetByYearIdHandler).Methods("GET")
    r.Handle("/year/{yearid}", DeleteYearHandler).Methods("DELETE")
    
    logger.Log("msg", "HTTP", "addr", ":"+os.Getenv("PORT"))
    logger.Log("err", http.ListenAndServe(":"+os.Getenv("PORT"), nil))
    
}