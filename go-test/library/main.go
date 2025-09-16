package main

// ----- Ошибки -----
// ErrBookNotFound
// ErrBookNotAvailable
// ErrUserNotFound
//
// ----- Структуры -----
// Book ID, Title, Author, IsAvailable
// User ID, Name, BorrowedBooks ([]int)
// Library books (map[int]Book), users (map[int]User)
//
// ----- Интерфейс -----
// LibraryManager
// - AddBook(title, author string)
// - AddUser(name string)
// - BorrowBook(userID, bookID int) error
// - ReturnBook(userID, bookID int) error
// - ShowBooks()
// - ShowUsers()
//
// ----- Методы -----
// AddBook создать Book с новым ID, добавить в map
// AddUser создать User с новым ID, добавить в map
// BorrowBook проверить ошибки, снять книгу с полки, отдать пользователю
// ReturnBook проверить, вернуть книгу на полку
// ShowBooks вывести все книги
// ShowUsers вывести всех пользователей и их книги
//
// ----- main -----
// Цикл меню
// 1. Добавить книгу
// 2. Добавить пользователя
// 3. Взять книгу
// 4. Вернуть книгу
// 5. Показать книги
// 6. Показать пользователей
// 7. Выход
func main() {
	//инициализация библиотеки
	//цикл меню и вызов методов
}
