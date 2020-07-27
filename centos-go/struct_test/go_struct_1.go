package main
import "fmt"

type Books struct{
  title string
  author string
  subject string
  book_id int
}

func main(){
  //创建一个新的结构体
  fmt.Println(Books{"go 语言","www.runoob.com","go 语言教程",123456})

  //可以使用健值对创建
  fmt.Println(Books{title:"go 语言", author:"abc", subject:"go 语言教程", book_id:9527})

  //忽略的字段为 0 或 空
  fmt.Println(Books{title:"go langue", author:"abc"})

  /*==============test 2=============*/
  fun_1()
}

func fun_1(){
   var Book1 Books        /* Declare Book1 of type Book */
   var Book2 Books        /* Declare Book2 of type Book */

   /* book 1 描述 */
   Book1.title = "Go 语言"
   Book1.author = "www.runoob.com"
   Book1.subject = "Go 语言教程"
   Book1.book_id = 6495407

   /* book 2 描述 */
   Book2.title = "Python 教程"
   Book2.author = "www.runoob.com"
   Book2.subject = "Python 语言教程"
   Book2.book_id = 6495700

   /* 打印 Book1 信息 */
   printBook(&Book1)

   /* 打印 Book2 信息 */
   printBook(&Book2)
}

func printBook( book *Books ) {
   fmt.Println("----------------------------")
   fmt.Println(*book)
   fmt.Printf( "Book title : %s\n", book.title)
   fmt.Printf( "Book author : %s\n", book.author)
   fmt.Printf( "Book subject : %s\n", book.subject)
   fmt.Printf( "Book book_id : %d\n", book.book_id)
}
