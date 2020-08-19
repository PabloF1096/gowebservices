package main

import (
    "context"
    pb "github.com/PabloF1096/gowebservices/booksapp"
    "github.com/PabloF1096/gowebservices/booksapp"
    "google.golang.org/grpc"
    "log"
    "os"
    "time"
    "encoding/csv"
    "fmt"
)

func main() {
    address := os.Getenv("ADDRESS")
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()
    c := pb.NewBookInfoClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    //Read Books Data
    readData("books.csv")
    //Disyplay
    disyplayCsvData()
    //Add books to web service
    addBooks(c, ctx)
    //Get All books from the web service
    getAll(c, ctx)
    //Update book from the web service
    updateBook(c, ctx)
    //delete book from the web service
    deleteBook(c, ctx,  bookIDs[0])
    
}

type Book struct {
    Id        string `json:"id"`
    Title     string `json:"title"`
    Edition   string `json:"edition"`
    Copyright string `json:"copyright"`
    Language  string `json:"language"`
    Pages     string `json:"pages"`
    Author    string `json:"author"`
    Publisher string `json:"publisher"`
}

var books []Book
var bookIDs []string

func checkError(message string, err error) {
    if err != nil {
        log.Fatal(message, err)
    }
}

func readData(filePath string) {
    fmt.Println("Read CSV")
    file, err1 := os.Open(filePath)
    checkError("Unable to read input file "+filePath, err1)
    defer file.Close()

    csvReader := csv.NewReader(file)
    records, err2 := csvReader.ReadAll()
    checkError("Unable to parse file as CSV for "+filePath, err2)
    defer file.Close()

    books = []Book{}

    for _, record := range records {
        book := Book{
            Id:        record[0],
            Title:     record[1],
            Edition:   record[2],
            Copyright: record[3],
            Language:  record[4],
            Pages:     record[5],
            Author:    record[6],
            Publisher: record[7]}
        books = append(books, book)
    }
    file.Close()
}

func disyplayCsvData() {
    fmt.Println("Display data from CSV")
    for _, book := range books {
        fmt.Println(book)
	}
}

func addBooks(c booksapp.BookInfoClient, ctx context.Context) {
    fmt.Println("Add books")
    for _, book := range books {
        r, err := c.AddBook(ctx, &pb.Book{
            Id:        book.Id,
            Title:     book.Title,
            Edition:   book.Edition,
            Copyright: book.Copyright,
            Language:  book.Language,
            Pages:     book.Pages,
            Author:    book.Author,
            Publisher: book.Publisher})
        if err != nil {
        log.Fatalf("Could not add book: %v", err)
        }
        bookIDs = append(bookIDs, r.Value)
        log.Printf("Book ID: %s added successfully", r.Value)
    }
}

func getAll(c booksapp.BookInfoClient, ctx context.Context) {
    fmt.Println("Get all books")
    for _, id := range bookIDs {
		book, err := c.GetBook(ctx, &pb.BookID{Value: id})
		if err != nil {
			log.Fatalf("Could not delete the book: %v", err)
		}
		fmt.Println(book.String())
	}

}

func getBook(c booksapp.BookInfoClient, ctx context.Context, id string) {

	book, err := c.GetBook(ctx, &pb.BookID{Value: id})
	if err != nil {
		log.Fatalf("Could not get book: %v", err)
	}
	fmt.Println(book.String())
}

func updateBook(c booksapp.BookInfoClient, ctx context.Context) {
    fmt.Println("Update book")
    bookId := bookIDs[0]

    fmt.Println("original book data")
    getBook(c, ctx, bookId)

    bookInfo := &pb.Book{
		Id:        "1",
		Title:     "Cyberpunk",
		Edition:   "22th",
		Copyright: "2077",
		Language:  "Letones",
		Pages:     "600",
		Author:    "Pablo",
        Publisher: "Pablo"}
    idUpdate := &pb.BookID{Value: bookId}
    bookUpdate := &pb.BookUpdate{Update: bookInfo, Id: idUpdate}

    _, err := c.UpdateBook(ctx, bookUpdate)
    if err != nil {
        log.Fatalf("Could not update book: %v", err)
    }
    fmt.Println("update book data")
    getBook(c, ctx,bookId)
}

func deleteBook(c booksapp.BookInfoClient, ctx context.Context, id string) {
    fmt.Println("delete book")
    status, err := c.DeleteBook(ctx, &pb.BookID{Value: id})
    if err != nil{
        log.Fatalf("Could not delete book: %v", err)
    }
    fmt.Println("Book deleted: "+id)
    log.Printf("Book deleted: ", status.String())
}