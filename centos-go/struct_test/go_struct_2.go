package main

import "fmt"

type Books struct {
    title string
    author string
    subject string
    book_id int
}

func changeBook(book *Books) {
    book.title = "book1_change"
}

func main() {
    var book1 Books
    book1.title = "book1"
    book1.author = "zuozhe"
    book1.book_id = 1
    changeBook(&book1)
    fmt.Println(book1)
}

/*
结构体是作为参数的值传递
如果想在函数里面改结构体(本体)数据内容，需要传入指针
*/
