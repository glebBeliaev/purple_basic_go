package main

import "fmt"

// Создать приложение, которое сначала выдаёт меню:
// Посмотреть закладки
// Добавить закладку
// Удалить закладку
// Выход
// При 1 - Выводит закладки
// При 2 - 2 поля ввода названия и адреса и после добавление
// При 3 - Ввод названия и удаление по нему
// При 4 - Завершение

type bookmarkMap = map[string]string

func main() {
	bookmarks := map[string]string{
		"vk": "https://vk.com",
		"fb": "https://facebook.com",
	}

Menu:
	for {
		action := getMenu()
		switch action {
		case 1:
			viewBookmarks(bookmarks)
		case 2:
			bookmarks = addBookmark(bookmarks)
		case 3:
			bookmarks = deleteBookmark(bookmarks)
		default:
			break Menu
		}
	}
}

func getMenu() int {
	var action int
	fmt.Println("")
	fmt.Println("1 - Посмотреть закладки")
	fmt.Println("2 - Добавить закладку")
	fmt.Println("3 - Удалить закладку")
	fmt.Println("4 - Выход")
	fmt.Print("Выберите действие:")

	fmt.Scan(&action)
	return action
}

func viewBookmarks(m bookmarkMap) {
	fmt.Println("Закладки:")
	for k, v := range m {
		fmt.Println(k, v)
	}
}

func addBookmark(m bookmarkMap) bookmarkMap {
	var name, url string
	fmt.Println("Введите название закладки:")
	fmt.Scan(&name)
	fmt.Println("Введите URL закладки:")
	fmt.Scan(&url)
	m[name] = url
	fmt.Println("Закладка добавлена")
	return m
}

func deleteBookmark(m bookmarkMap) bookmarkMap {
	var name string
	fmt.Println("Введите название закладки:")
	fmt.Scan(&name)
	delete(m, name)
	fmt.Println("Закладка удалена")
	return m
}
