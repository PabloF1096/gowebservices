package main

import (
    "os"
    "encoding/csv"
    "fmt"
    "context"
    "github.com/go-kit/kit/log"
)

type Year struct {
	IdYear  string `json:"idYear"`
	Year    string `json:"Year"`
}

type yearservice struct {
	logger log.Logger
}

var years []Year

func readDataFromYearCSV(filePath string) {
    fmt.Println("Read Year CSV")
    file, err1 := os.Open(filePath)
    checkError("Unable to read input file "+filePath, err1)
    defer file.Close()

    csvReader := csv.NewReader(file)
    records, err2 := csvReader.ReadAll()
    checkError("Unable to parse file as CSV for "+filePath, err2)
    defer file.Close()

    years = []Year{}

    for _, record := range records {
        year := Year{
            IdYear:    record[0],
            Year:      record[1]}
		years = append(years, year)
    }
    file.Close()
}

func writeDataToYearCSV(filePath string) {
    fmt.Println("Write to Year CSV File")
    file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 0644)
    checkError("Cannot create file", err)
    defer file.Close()

    file.Seek(0, 0)
    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, year := range years {
        record := []string{year.IdYear,year.Year}
        err := writer.Write(record)
        checkError("Cannot write to file", err)
    }
    writer.Flush()
    file.Close()
}

type YearService interface {
    CreateYear(ctx context.Context, year Year) (string, error)
    GetYearById(ctx context.Context, id string) (interface{}, error)
    UpdateYear(ctx context.Context, year Year) (string, error)
    DeleteYear(ctx context.Context, id string) (string, error)
}

func findYear(x string) int {
	for i, year := range years {
    // fmt.Println(show)
		if x == year.IdYear {
			return i
		}
    }
	return -1
}

func NewServiceYear(logger log.Logger) YearService {
    return &yearservice{
        logger: logger,
    }
}

func (s yearservice) CreateYear(ctx context.Context, year Year) (string, error) {
    var msg = "success"
    years = append(years, year)
    return msg, nil
}

func (s yearservice) GetYearById(ctx context.Context, id string) (interface{}, error) {
    var err error
    var year interface{}
    var empty interface{}
    i := findYear(id)
    if i == -1 {
        return empty, err
    }
    year = years[i]
    return year, nil
}
func (s yearservice) DeleteYear(ctx context.Context, id string) (string, error) {
    var err error
    msg := ""
    i := findYear(id)
    if i == -1 {
        return "", err
    }
    copy(years[i:], years[i+1:])
    years[len(years)-1] = Year{}
    years = years[:len(years)-1]
    return msg, nil
}
func (s yearservice) UpdateYear(ctx context.Context, year Year) (string, error) {
    var empty = ""
    var err error
    var msg = "success"
    i := findYear(year.IdYear)
    if i == -1 {
        return empty, err
    }
    years[i] = year
    return msg, nil
}