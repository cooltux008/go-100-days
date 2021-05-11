package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var (
		Book1 Books
		Book2 Books
	)

	Book1.title = "Go 语言"
	Book1.author = "www.runoob.com"
	Book1.subject = "Go 语言教程"
	Book1.book_id = 6495407

	Book2.title = "Python 语言"
	Book2.author = "www.runoob.com"
	Book2.subject = "Python 语言教程"
	Book2.book_id = 6495700

	printBook(Book1)
	printBook(Book2)
}

func printBook(book Books) {
	fmt.Printf("book title: %s\n", book.title)
	fmt.Printf("book author: %s\n", book.author)
	fmt.Printf("book subject: %s\n", book.subject)
	fmt.Printf("book book_id: %d\n", book.book_id)

}
