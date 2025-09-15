package main

import "fmt"

type Contact struct {
	Name  string
	Phone string
}
type ContactBook struct {
	Contact map[string]Contact
}

const separatorLn string = "_______________________________________________"

func addContacts(contacts map[string]Contact) {
	var name string
	var phone string
	fmt.Println("Добавьте имя")
	fmt.Scan(&name)
	fmt.Println("Добавьте телефон")
	fmt.Scan(&phone)

	contacts[name] = Contact{
		Name:  name,
		Phone: phone,
	}
}

func showContacts(contacts map[string]Contact) {
	if len(contacts) == 0 {
		fmt.Println("Список контактов пуст")
		return
	}

	for _, val := range contacts {
		fmt.Println(separatorLn)
		fmt.Println("Имя: " + val.Name)
		fmt.Println("Телефон: " + val.Phone)
		fmt.Println(separatorLn)
	}
}

func searchContact(contacts map[string]Contact) {
	var name string
	var contact Contact
	var completeSearch bool = false
	fmt.Println("Введите имя для поиска контакта")
	fmt.Scan(&name)

	for key, val := range contacts {
		if name == key {
			contact = val
			completeSearch = true
			break
		}
	}

	if !completeSearch {
		fmt.Println(separatorLn)
		fmt.Println("Контакт не найден")
		fmt.Println(separatorLn)
		return
	}
	fmt.Println(separatorLn)
	fmt.Println("Имя: " + contact.Name)
	fmt.Println("Телефон: " + contact.Phone)
	fmt.Println(separatorLn)

}

func deleteContact(contacts map[string]Contact) {
	var name string
	var completeSearch bool = false
	fmt.Println("Введите имя для поиска контакта")
	fmt.Scan(&name)

	for key, _ := range contacts {
		if key == name {
			delete(contacts, key)
			completeSearch = true
			break
		}
	}

	if !completeSearch {
		fmt.Println("Контакт не найден")
		return
	}
	fmt.Println("Контакт удален")
}

func main() {
	// тут храни все контакты в виде map[string]Contact
	contacts := ContactBook{
		Contact: make(map[string]Contact),
	}
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
			addContacts(contacts.Contact)
		case 2:
			showContacts(contacts.Contact)
		case 3:
			searchContact(contacts.Contact)
		case 4:
			deleteContact(contacts.Contact)
		case 5:
			fmt.Println("Выход...")
			return
		}
	}
}
