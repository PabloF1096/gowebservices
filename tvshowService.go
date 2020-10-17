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
    Age        string `json:"Age"`
    RottenT    string `json:"Rotten Tomatoes"`
    Netflix    string `json:"Netflix"`
    Hulu       string `json:"Hulu"`
    PrimeV     string `json:"Prime Video"`
    Disney     string `json:"Disney+"`
    Type       string `json:"type"`
    IdYear     string `json:"idYear"`
    IdRate     string `json:"idRate"`
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
            Age:       record[2],
            RottenT:   record[3],
            Netflix:   record[4],
            Hulu:      record[5],
            PrimeV:    record[6],
            Disney:    record[7],
            Type:      record[8],
            IdYear:    record[9],
            IdRate:    record[10]}
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
        record := []string{show.IdShow,show.Title,show.Age,show.RottenT,show.Netflix,show.Hulu,show.PrimeV,show.Disney,show.Type,show.IdYear,show.IdRate}
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
    GetYearByTVshowId(ctx context.Context, id string) (interface{}, error)
    GetRateByTVshowId(ctx context.Context, id string) (interface{}, error)
    UpdateYearWithTVshowId(ctx context.Context, tvshow TVshow, id string) (string, error)
    DeleteYearWithTVshow(ctx context.Context, idShow string, idRate string) (string, error)
    UpdateRateWithTVshowId(ctx context.Context, tvshow TVshow, id string) (string, error)
    DeleteRateWithTVshow(ctx context.Context, idShow string, idRate string) (string, error)
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

func (s tvshowservice) GetYearByTVshowId(ctx context.Context, id string) (interface{}, error) {
    var err error
    var year interface{}
    var empty interface{}
    i := findTVshow(id)
    if i == -1 {
        return empty, err
    }
    var idYear = shows[i].IdYear

    i2 := findYear(idYear)
    if i2 == -1 {
        return empty, err
    }
    year = years[i2]

    return year, nil
}

func (s tvshowservice) GetRateByTVshowId(ctx context.Context, id string) (interface{}, error) {
    var err error
    var rate interface{}
    var empty interface{}
    i := findTVshow(id)
    if i == -1 {
        return empty, err
    }
    
    var idRate = shows[i].IdRate
    i2 := findRate(idRate)
    if i2 == -1 {
        return empty, err
    }
    rate = rates[i2]

    return rate, nil
}

func (s tvshowservice) UpdateYearWithTVshowId(ctx context.Context, tvshow TVshow, id string) (string, error) {
    var empty = ""
    var err error
    var msg = "success"
    i := findTVshow(id)
    if i == -1 {
        return empty, err
    }
    shows[i].IdYear = tvshow.IdYear
    return msg, nil
}

func (s tvshowservice) DeleteYearWithTVshow(ctx context.Context, idShow string, idYear string) (string, error) {
    var err error
    msg := ""
    i := findTVshow(idShow)
    if i == -1 {
        return "", err
    }
    shows[i].IdYear = ""
    return msg, nil
}

func (s tvshowservice) UpdateRateWithTVshowId(ctx context.Context, tvshow TVshow, id string) (string, error) {
    var empty = ""
    var err error
    var msg = "success"
    i := findTVshow(id)
    if i == -1 {
        return empty, err
    }
    shows[i].IdRate = tvshow.IdRate
    return msg, nil
}

func (s tvshowservice) DeleteRateWithTVshow(ctx context.Context, idShow string, idRate string) (string, error) {
    var err error
    msg := ""
    i := findTVshow(idShow)
    if i == -1 {
        return "", err
    }
    shows[i].IdRate = ""
    return msg, nil
}