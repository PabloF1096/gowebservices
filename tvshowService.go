package main

import (
    "os"
    "encoding/csv"
    "fmt"
    "context"
    "github.com/go-kit/kit/log"
)

type TVshow struct {
    IdShow     string `json:"idShow"`
    Title      string `json:"Title"`
    Year       string `json:"Year"`
    Age        string `json:"Age"`
    IMDb       string `json:"IMDb"`
    RottenT    string `json:"Rotten Tomatoes"`
    Netflix    string `json:"Netflix"`
    Hulu       string `json:"Hulu"`
    PrimeV     string `json:"Prime Video"`
    Disney     string `json:"Disney+"`
    Type       string `json:"type"`
}

type tvshowservice struct {
    logger log.Logger
}

var shows []TVshow

func checkError(message string, err error) {
    if err != nil {
        // log.Fatal(message, err)
        fmt.Println(message, err)
    }
}

func readDataFromTvShowCSV(filePath string) {
    fmt.Println("Read CSV")
    file, err1 := os.Open(filePath)
    checkError("Unable to read input file "+filePath, err1)
    defer file.Close()

    csvReader := csv.NewReader(file)
    records, err2 := csvReader.ReadAll()
    checkError("Unable to parse file as CSV for "+filePath, err2)
    defer file.Close()

    shows = []TVshow{}

    for _, record := range records {
        show := TVshow{
            IdShow:    record[0],
            Title:     record[1],
            Year:      record[2],
            Age:       record[3],
            IMDb:      record[4],
            RottenT:   record[5],
            Netflix:   record[6],
            Hulu:      record[7],
            PrimeV:    record[8],
            Disney:    record[9],
            Type:      record[10]}
		shows = append(shows, show)
    }
    file.Close()
}

func disyplayCsvData() {
    fmt.Println("Display data from CSV")
    for _, show := range shows {
        fmt.Println(show)
	}
}

func writeDataTvShowCSV(filePath string) {
    fmt.Println("Write CSV File")
    file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
    checkError("Cannot create file", err)
    defer file.Close()

    file.Seek(0, 0)
    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, show := range shows {
        record := []string{show.IdShow,show.Title,show.Year,show.Age,show.IMDb,show.RottenT,show.Netflix,show.Hulu,show.PrimeV,show.Disney,show.Type}
        err := writer.Write(record)
        checkError("Cannot write to file", err)
    }
    writer.Flush()
    file.Close()
}

type TVshowService interface {
    CreateTVshow(ctx context.Context, tvshow TVshow) (string, error)
    GetTVshowById(ctx context.Context, id string) (interface{}, error)
    UpdateTVshow(ctx context.Context, tvshow TVshow) (string, error)
    DeleteTVshow(ctx context.Context, id string) (string, error)
    GetYears(ctx context.Context, id string) (interface{}, error)
    GetRates(ctx context.Context, id string) (interface{}, error)
    GetAges(ctx context.Context, id string) (interface{}, error)
}

func findTVshow(x string) int {
	for i, show := range shows {
		if x == show.IdShow {
			return i
		}
    }
	return -1
}

func NewServiceTVshow(logger log.Logger) TVshowService {
    return &tvshowservice{
        logger: logger,
    }
}

func (s tvshowservice) CreateTVshow(ctx context.Context, tvshow TVshow) (string, error) {
    var msg = "success"
    shows = append(shows, tvshow)
    return msg, nil
}

func (s tvshowservice) GetTVshowById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var tvshow interface{}
    var empty interface{}
    i := findTVshow(id)
    if i == -1 {
        return empty, err
    }

    tvshow = shows[i]
    return tvshow, nil
}
func (s tvshowservice) DeleteTVshow(ctx context.Context, id string) (string, error) {
    var err error
    msg := ""
    i := findTVshow(id)
    if i == -1 {
        return "", err
    }
    copy(shows[i:], shows[i+1:])
    shows[len(shows)-1] = TVshow{}
    shows = shows[:len(shows)-1]
    return msg, nil
}
func (s tvshowservice) UpdateTVshow(ctx context.Context, tvshow TVshow) (string, error) {
    var empty = ""
    var err error
    var msg = "success"
    i := findTVshow(tvshow.IdShow)
    if i == -1 {
        return empty, err
    }
    shows[i] = tvshow
    return msg, nil
}

func (s tvshowservice) GetYears(ctx context.Context, id string) (interface{}, error) {
    var tvshow = []TVshow{}
    var err error
    var empty interface{}
    for _, show := range shows {
		if id == show.Year {
            tvshow = append(tvshow, show)
		}
    }
    if tvshow == nil {
        return empty, err
    }
	return tvshow, nil
}
func (s tvshowservice) GetRates(ctx context.Context, id string) (interface{}, error) {
    var tvshow = []TVshow{}
    var err error
    var empty interface{}
    for _, show := range shows {
		if id == show.IMDb {
            tvshow = append(tvshow, show)
		}
    }
    if tvshow == nil {
        return empty, err
    }
	return tvshow, nil
}
func (s tvshowservice) GetAges(ctx context.Context, id string) (interface{}, error) {
    var tvshow = []TVshow{}
    var err error
    var empty interface{}
    for _, show := range shows {
		if id == show.Age {
            tvshow = append(tvshow, show)
		}
    }
    if tvshow == nil {
        return empty, err
    }
	return tvshow, nil
}