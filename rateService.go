package main

import (
    "os"
    "encoding/csv"
    "fmt"
    "context"
    "github.com/go-kit/kit/log"
)

type Rate struct {
	IdRate  string `json:"idRate"`
	Imdb  string `json:"IMDb"`
}

type rateservice struct {
	logger log.Logger
}

var rates []Rate

func readDataFromRateCSV(filePath string) {
    fmt.Println("Read Rate CSV")
    file, err1 := os.Open(filePath)
    checkError("Unable to read input file "+filePath, err1)
    defer file.Close()

    csvReader := csv.NewReader(file)
    records, err2 := csvReader.ReadAll()
    checkError("Unable to parse file as CSV for "+filePath, err2)
    defer file.Close()

    rates = []Rate{}

    for _, record := range records {
        rate := Rate{
            IdRate:    record[0],
            Imdb:      record[1]}
		rates = append(rates, rate)
    }
    file.Close()
}

func writeDataToRateCSV(filePath string) {
    fmt.Println("Write to Rate CSV File")
    file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
    checkError("Cannot create file", err)
    defer file.Close()

    file.Seek(0, 0)
    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, rate := range rates {
        record := []string{rate.IdRate,rate.Imdb}
        err := writer.Write(record)
        checkError("Cannot write to file", err)
    }
    writer.Flush()
    file.Close()
}

type RateService interface {
    CreateRate(ctx context.Context, rate Rate) (string, error)
    GetRateById(ctx context.Context, id string) (interface{}, error)
    UpdateRate(ctx context.Context, rate Rate) (string, error)
    DeleteRate(ctx context.Context, id string) (string, error)
}

func findRate(x string) int {
	for i, rate := range rates {
    // fmt.Println(show)
		if x == rate.IdRate {
			return i
		}
    }
	return -1
}

func NewServiceRate(logger log.Logger) RateService {
    return &rateservice{
        logger: logger,
    }
}

func (s rateservice) CreateRate(ctx context.Context, rate Rate) (string, error) {
    var msg = "success"
    rates = append(rates, rate)
    return msg, nil
}

func (s rateservice) GetRateById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var rate interface{}
    var empty interface{}
    i := findRate(id)
    if i == -1 {
        return empty, err
    }
    rate = rates[i]
    return rate, nil
}
func (s rateservice) DeleteRate(ctx context.Context, id string) (string, error) {
    var err error
    msg := ""
    i := findRate(id)
    if i == -1 {
        return "", err
    }
    copy(rates[i:], rates[i+1:])
    rates[len(rates)-1] = Rate{}
    rates = rates[:len(rates)-1]
    return msg, nil
}
func (s rateservice) UpdateRate(ctx context.Context, rate Rate) (string, error) {
    var empty = ""
    var err error
    var msg = "success"
    i := findRate(rate.IdRate)
    if i == -1 {
        return empty, err
    }
    rates[i] = rate
    return msg, nil
}