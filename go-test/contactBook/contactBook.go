package main

import "fmt"

type ContactManager interface {
	Add(name, phone string)
	Show()
	Delete(name string)
	Search(name string)
}

type Contact struct {
	Name  string
	Phone string
}
type ContactBook struct {
	Contact map[string]Contact
}

const separatorLn string = "_______________________________________________"

func (cb *ContactBook) Add(name, phone string) {
	cb.Contact[name] = Contact{
		Name:  name,
		Phone: phone,
	}
}

func (cb *ContactBook) Show() {
	if len(cb.Contact) == 0 {
		fmt.Println("Список контактов пуст")
		return
	}

	for _, val := range cb.Contact {
		fmt.Println(separatorLn)
		fmt.Println("Имя: " + val.Name)
		fmt.Println("Телефон: " + val.Phone)
		fmt.Println(separatorLn)
	}
}

func (cb *ContactBook) Search(name string) {
	if contact, ok := cb.Contact[name]; ok {
		fmt.Println(separatorLn)
		fmt.Println("Имя: " + contact.Name)
		fmt.Println("Телефон: " + contact.Phone)
		fmt.Println(separatorLn)
	} else {
		fmt.Println(separatorLn)
		fmt.Println("Контакт не найден")
		fmt.Println(separatorLn)
		return
	}
}

func (cb *ContactBook) Delete(name string) {
	if _, ok := cb.Contact[name]; ok {
		delete(cb.Contact, name)
		fmt.Println("Контакт удален")
	} else {
		fmt.Println("Контакт не найден")
		return
	}
}

func readInput(promt string) string {
	var res string
	fmt.Println(promt)
	fmt.Scan(&res)
	return res
}

func main() {
	var manager ContactManager = &ContactBook{Contact: make(map[string]Contact)}
	for {
		fmt.Println("\n1. Добавить контакт")
		fmt.Println("2. Показать контакты")
		fmt.Println("3. Найти контакт")
		fmt.Println("4. Удалить контакт")
		fmt.Println("5. Выход")

		var choice int
		fmt.Print("Ваш выбор: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			name := readInput("Введите имя:")
			phone := readInput("Введите телефон:")
			manager.Add(name, phone)
		case 2:
			manager.Show()
		case 3:
			name := readInput("Введите имя:")
			manager.Search(name)
		case 4:
			name := readInput("Введите имя:")
			manager.Delete(name)
		case 5:
			fmt.Println("Выход...")
			return
		}
	}
}
