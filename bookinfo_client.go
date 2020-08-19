package main

import (
    "context"
    pb "github.com/PabloF1096/gowebservices/booksapp"
    "google.golang.org/grpc"
    "log"
    "os"
    "time"
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
    //Add
    r, err := c.AddBook(ctx, &pb.Book{
        Id:        "1",
        Title:     "Operating System Concepts",
        Edition:   "9th",
        Copyright: "2012",
        Language:  "ENGLISH",
        Pages:     "976",
        Author:    "Abraham Silberschatz",
        Publisher: "John Wiley & Sons"})
    if err != nil {
        log.Fatalf("Could not add book: %v", err)
    }

    log.Printf("Book ID: %s added successfully", r.Value)
    //Get
    book, err := c.GetBook(ctx, &pb.BookID{Value: r.Value})
    if err != nil {
        log.Fatalf("Could not get book: %v", err)
    }
    log.Printf("Book: ", book.String())
    // Update
    bookInfo := &pb.Book{
		Id:        "1",
		Title:     "Cyberpunk",
		Edition:   "22th",
		Copyright: "2077",
		Language:  "Letones",
		Pages:     "600",
		Author:    "Pablo",
        Publisher: "Pablo"}
    idUpdate := &pb.BookID{Value: r.Value}
    bookUpdate := &pb.BookUpdate{Update: bookInfo, Id: idUpdate}
    book1, err := c.UpdateBook(ctx,bookUpdate)
    if err != nil {
        log.Fatalf("Could not update book: %v", err)
    }
    log.Printf("Book Updated: ", book1.String())
    //Delete
    status, err := c.DeleteBook(ctx, &pb.BookID{Value: r.Value})
    if err != nil{
        log.Fatalf("Could not delete book: %v", err)
    }
    log.Printf("Book deleted: ", status.String())
}