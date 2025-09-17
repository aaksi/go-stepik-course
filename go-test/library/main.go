package main

import "fmt"

// ----- Ошибки -----
// ErrBookNotFound
// ErrBookNotAvailable
// ErrUserNotFound
//
// ----- Структуры -----
type Book struct {
	ID          int
	Title       string
	Author      string
	isAvailable bool
}

type User struct {
	ID            int
	Name          string
	BorrowedBooks []int
}

type LibraryBooks struct {
	Books map[int]Book
	Users map[int]User
}

type LibraryManager interface {
	AddBook(title, author string)
	AddUser(Name string)
	BorrowBook(userID, bookID int) error
	ReturnBook(userID, bookID int) error
	ShowBooks()
	ShowUsers()
}

// ----- Методы -----
// AddBook создать Book с новым ID, добавить в map
// AddUser создать User с новым ID, добавить в map
// BorrowBook проверить ошибки, снять книгу с полки, отдать пользователю
// ReturnBook проверить, вернуть книгу на полку
// ShowBooks вывести все книги
// ShowUsers вывести всех пользователей и их книги

func main() {

	var library LibraryManager = &LibraryBooks{}

	//инициализация библиотеки
	//цикл меню и вызов методов
	var chair string

	fmt.Println("Добро пожаловать в библиотеку!")
	fmt.Println("1. Добавить книгу")
	fmt.Println("2. Добавить пользователя")
	fmt.Println("3. Взять книгу")
	fmt.Println("4. Вернуть книгу")
	fmt.Println("5. Показать книги")
	fmt.Println("6. Показать пользователей")
	fmt.Println("7. Выход")

	fmt.Scan(&chair)

	switch chair {
	case "1":
		fmt.Println(chair)
	case "2":
		fmt.Println(chair)
	case "3":
		fmt.Println(chair)
	case "4":
		fmt.Println(chair)
	case "5":
		fmt.Println(chair)
	case "6":
		fmt.Println(chair)
	case "7":
		fmt.Println(chair)
	}

}
