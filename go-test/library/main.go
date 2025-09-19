package main

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
)

var bookCounterID int = 1
var userCounterID int = 1

const separator string = "======================"

var (
	ErrBookNotFound     = errors.New("Книга не найдена по такому ID")
	ErrBookNotAvailable = errors.New("Книги нет в наличии")
	ErrUserNotFound     = errors.New("Пользователь не найден по такому ID")
)

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
	Books map[int]*Book
	Users map[int]*User
}

type LibraryManager interface {
	AddBook(title, author string)
	AddUser(name string)
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

func (lb *LibraryBooks) AddBook(title, author string) {
	lb.Books[bookCounterID] = &Book{
		ID:          bookCounterID,
		Title:       title,
		Author:      author,
		isAvailable: false,
	}
	bookCounterID++
}

func (lb *LibraryBooks) AddUser(name string) {
	lb.Users[userCounterID] = &User{
		ID:            userCounterID,
		Name:          name,
		BorrowedBooks: []int{},
	}
	userCounterID++
}

func (lb *LibraryBooks) ShowUsers() {
	for _, user := range lb.Users {
		fmt.Println("ID Пользователя: " + strconv.Itoa(user.ID))
		fmt.Println("Имя пользователя: " + user.Name)
		if len(user.BorrowedBooks) == 0 {
			fmt.Println("У пользователя нет книг на руках")
		} else {
			fmt.Println("Книги на руках: ")
			for _, book := range user.BorrowedBooks {
				fmt.Println("ID книги: " + strconv.Itoa(book))
				fmt.Println("Название книги: " + lb.Books[book].Title)
				fmt.Println("Автор книги: " + lb.Books[book].Author)
			}
		}
		fmt.Println(separator)
	}
}

func (lb *LibraryBooks) ShowBooks() {

	for _, item := range lb.Books {
		fmt.Println("ID книги: " + strconv.Itoa(item.ID))
		fmt.Println("Название книги: " + item.Title)
		fmt.Println("Автор книги: " + item.Author)
		if item.isAvailable {
			fmt.Println("Книга у читателя")
		} else {
			fmt.Println("Книга в библиотеке")
		}

		fmt.Println(separator)
	}

}

func (lb *LibraryBooks) BorrowBook(userID, bookID int) error {

	_, ok := lb.Users[userID]
	if !ok {
		return ErrUserNotFound
	} else {
		_, ok = lb.Books[bookID]
		if !ok {
			return ErrBookNotFound
		}
	}

	if lb.Books[bookID].isAvailable {
		return ErrBookNotAvailable
	} else {
		lb.Users[userID].BorrowedBooks = append(lb.Users[userID].BorrowedBooks, bookID)
		lb.Books[bookID].isAvailable = true
	}

	return nil
}
func (lb *LibraryBooks) ReturnBook(userID, bookID int) error {
	_, ok := lb.Users[userID]

	if !ok {
		return ErrUserNotFound
	} else {
		_, ok = lb.Books[bookID]
		if !ok {
			return ErrBookNotFound
		}
	}

	bookInUserIdx := slices.Index(lb.Users[userID].BorrowedBooks, bookID)

	if bookInUserIdx == -1 {
		return errors.New("Пользователь не брал книгу с таким ID")
	} else {
		lb.Users[userID].BorrowedBooks = slices.Delete(lb.Users[userID].BorrowedBooks, bookInUserIdx, bookInUserIdx+1)
		lb.Books[bookID].isAvailable = false
	}

	return nil
}

func main() {
	var library LibraryManager = &LibraryBooks{
		Books: make(map[int]*Book),
		Users: make(map[int]*User),
	}

	for {
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
			var name string
			var author string
			fmt.Println("Введите название книги")
			fmt.Scan(&name)
			fmt.Println("Введите автора")
			fmt.Scan(&author)
			library.AddBook(name, author)
		case "2":
			var name string
			fmt.Println("Введите имя пользователя")
			fmt.Scan(&name)
			library.AddUser(name)
		case "3":
			var userID, bookID int
			fmt.Println("Введите ID пользователя")
			fmt.Scan(&userID)
			fmt.Println("Введите ID книги")
			fmt.Scan(&bookID)
			library.BorrowBook(userID, bookID)
		case "4":
			var userID, bookID int
			fmt.Println("Введите ID пользователя")
			fmt.Scan(&userID)
			fmt.Println("Введите ID книги")
			fmt.Scan(&bookID)
			library.ReturnBook(userID, bookID)
		case "5":
			library.ShowBooks()
		case "6":
			library.ShowUsers()
		case "7":
			return
		}
	}
}
