package main

import (
	"fmt"
	"strconv"
)

// ----- Структуры -----

// Account
// Подсказка: хранит баланс пользователя
// Поля:
// - Balance: сумма на счету (float64)

type Account struct {
	Balance float64
}

// User
// Подсказка: хранит данные пользователя
// Поля:
// - ID: уникальный идентификатор (int)
// - Name: имя пользователя (string)
// - Account: счет пользователя (Account)

type User struct {
	ID      int
	Name    string
	Account Account
}

// ----- Глобальные переменные -----

// users
// Подсказка: map, ключ = ID (int), значение = User
var users = make(map[int]User)

// nextID
// Подсказка: глобальный счетчик для уникальных ID
var nextID = 1

// ----- Функции -----

// addUser
// Подсказка:
// Получает: имя пользователя (ввод внутри)
// Возвращает: ничего
// Действия:
// - спросить имя пользователя
// - создать нового User с уникальным ID
// - положить его в map users
// - увеличить nextID
func addUser() {
	var name string

	fmt.Println("Введите имя пользователя: ")
	fmt.Scanln(&name)

	user := User{
		ID:      nextID,
		Name:    name,
		Account: Account{Balance: 0.0},
	}

	fmt.Printf("Пользователь %s с ID %d добавлен \n", user.Name, user.ID)

	users[nextID] = user
	nextID++
}

// deposit
// Подсказка:
// Получает: ID пользователя и сумму (ввод внутри)
// Возвращает: ничего
// Действия:
// - спросить ID и сумму
// - проверить, есть ли пользователь
// - если есть — увеличить баланс
// - если нет — сообщение об ошибке
func deposit() {
	var idStr string

	fmt.Println("Введите ID пользователя чтоб увеличить депозит")
	fmt.Scan(&idStr)

	id, err := strconv.Atoi(idStr)
	if err != nil {
		fmt.Println("Некорректный ID")
		return
	}

	if user, ok := users[id]; ok {
		var amountStr string

		fmt.Println("Пользователь найден", user.Name)
		fmt.Println("Баланс", user.Account)
		fmt.Println("Введите сумму на которую хотите увеличить баланс")
		fmt.Scan(&amountStr)

		amount, err := strconv.ParseFloat(amountStr, 64)
		if err != nil {
			fmt.Println("Некоректная сумма")
			return
		}

		user.Account.Balance = amount
		users[id] = user

		fmt.Println("Баланс пополнен")
		fmt.Printf("Баланс аккаунта: %f", user.Account.Balance)

	} else {
		fmt.Println("Пользователь с таким ID не найден")
	}

}

// withdraw
// Подсказка:
// Получает: ID пользователя и сумму (ввод внутри)
// Возвращает: ничего
// Действия:
// - спросить ID и сумму
// - проверить пользователя и баланс
// - уменьшить баланс, если хватает денег
// - если нет — сообщение об ошибке
func withdraw() {
	var idStr string
	var amountStr string

	fmt.Println("Введите ID, чтоб снять деньги")
	fmt.Scanln(&idStr)

	id, err := strToInt(idStr)

	if err != nil {
		fmt.Println("Некорректный ID")
		return
	}

	fmt.Println("Введите сумму для снятия")
	fmt.Scanln(&amountStr)

	amount, errAmount := strToFloat(amountStr)

	if errAmount != nil {
		fmt.Println("Некорректно введено значение")
		return
	}

	if user, ok := users[id]; ok {
		if user.Account.Balance > amount {
			user.Account.Balance -= amount
			users[id] = user
		} else {
			fmt.Println("Недостаточно средств")
		}
	}

}

func strToInt(str string) (int, error) {
	res, err := strconv.Atoi(str)

	return res, err
}

func strToFloat(str string) (float64, error) {
	res, err := strconv.ParseFloat(str, 64)
	return res, err
}

// transfer
// Подсказка:
// Получает: ID отправителя, ID получателя, сумму (ввод внутри)
// Возвращает: ничего
// Действия:
// - спросить ID отправителя, ID получателя, сумму
// - проверить, существуют ли оба пользователя
// - проверить баланс отправителя
// - если всё ок — уменьшить баланс отправителя и увеличить баланс получателя
// - если нет — сообщение об ошибке
func transfer() {
	var idSenderStr string
	var idRecipientStr string
	var amountStr string

	fmt.Println("Введите ID отправителя")
	fmt.Scan(&idSenderStr)
	fmt.Println("Введите ID получателя")
	fmt.Scan(&idRecipientStr)
	fmt.Println("Введите сумму")
	fmt.Scan(&amountStr)

	idSender, _ := strToInt(idSenderStr)
	idRecipient, _ := strToInt(idRecipientStr)
	amount, _ := strToFloat(amountStr)

	userSender, okSender := users[idSender]
	userRecipient, okRecipient := users[idRecipient]

	if (okRecipient && okSender) && (users[idSender].Account.Balance >= amount) {
		userSender.Account.Balance = userSender.Account.Balance - amount
		users[idSender] = userSender
		userRecipient.Account.Balance = userRecipient.Account.Balance + amount
		users[idRecipient] = userRecipient
	} else {
		fmt.Println("Ошибка попробуйте еще раз")
	}

}

// showBalance
// Подсказка:
// Получает: ID пользователя (ввод внутри)
// Возвращает: ничего
// Действия:
// - спросить ID
// - проверить, есть ли пользователь
// - если есть — вывести баланс
// - если нет — сообщение
func showBalance() {
	var idstr string
	fmt.Println("Введите ID пользователя чтобы узнать баланс")
	fmt.Scan(&idstr)

	id, err := strToInt(idstr)
	if err != nil {
		fmt.Println("Пользователя с таким ID не существует")
		return
	}

	if user, ok := users[id]; ok {
		fmt.Printf("Баланс пользователя %s с ID - %d = %v", user.Name, user.ID, user.Account.Balance)
	}
}

// ----- main функция -----

func main() {
	for {
		fmt.Println("\n1. Добавить пользователя")
		fmt.Println("2. Пополнить счет")
		fmt.Println("3. Снять деньги")
		fmt.Println("4. Перевести деньги")
		fmt.Println("5. Показать баланс")
		fmt.Println("6. Выход")

		var choice int
		fmt.Print("Ваш выбор: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			addUser()
		case 2:
			deposit()
		case 3:
			withdraw()
		case 4:
			transfer()
		case 5:
			showBalance()
		case 6:
			fmt.Println("Выход...")
			return
		}
	}
}
