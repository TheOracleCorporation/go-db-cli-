package main

import "fmt"

type Bibliotheca struct {
    Title  string
    Author string
    Year   int
    Taken  bool
}

func findIdx(book []Bibliotheca, tempname string) int {
	idxBook := -1

	for i,b := range book {
		if b.Title == tempname {
			idxBook = i
			break
		}
	}
	return idxBook
}

func addBook(book *[]Bibliotheca) {
	
	var title string
	var author string
	var year int

	fmt.Print("Название добавляемой книги: ")
	fmt.Scan(&title)
	fmt.Print("Кто автор этой книги: ")
	fmt.Scan(&author)
	fmt.Print("Год выпуска этой книги: ")
	fmt.Scan(&year)

	newBook := Bibliotheca {
		Title:  title,
    	Author: author,
    	Year:   year,
    	Taken:  false,
	}

	*book = append(*book, newBook)
	fmt.Println("Книга успешно добавлена")
}

func showBook(book []Bibliotheca) {
	if len(book) == 0 {
		fmt.Println("----------")
		fmt.Println("Библиотека еще пуста")
		fmt.Println("----------")
		return
	} else {
		 for _, b := range book {
			fmt.Println("----------")
			fmt.Println("Книга:", b.Title )
			fmt.Println("Автор:", b.Author )
			fmt.Println("Год выпуска книги:", b.Year)

			if b.Taken {
				fmt.Println("Такой книги нет или ее кто то взял")
				fmt.Println("----------")
			} else {
				fmt.Println("Книга в наличии")
				fmt.Println("----------")
			}
		}
	}
}

func findBook(book []Bibliotheca) {
	var tempname string
	fmt.Println("Какую книгу ищите?")
	fmt.Scanln(&tempname)

	idxBook := findIdx(book, tempname)

	if idxBook != -1 {
		foundBook := book[idxBook]
		fmt.Println("----------")
		fmt.Println("Книга:", foundBook.Title )
		fmt.Println("Автор:", foundBook.Author )
		fmt.Println("Год выпуска книги:", foundBook.Year)

		if foundBook.Taken {
			fmt.Println("Такой книги нет")
			fmt.Println("----------")
		} else {
			fmt.Println("Книга в наличии")
			fmt.Println("----------")
		}
	} else {
		fmt.Println("Такой книги нет в бибилотеке")
	}
}

func takeBook(book []Bibliotheca) {
	var tempname string
	fmt.Println("Какую книгу вы хотите взять?")
	fmt.Scanln(&tempname)

	idxBook := findIdx(book, tempname)

	if idxBook != -1 {
		if book[idxBook].Taken {
			fmt.Println("Эту книгу уже кто то взял")
		} else {
			fmt.Println("Вы успешно взяли книгу",book[idxBook].Title)
			book[idxBook].Taken = true
		}
	} else {
		fmt.Println("Такой книги нет")
	}
}

func main() {
	book := []Bibliotheca{}

	flag := false

	for !flag {
		var value int
		fmt.Println("1.Добавить книгу")
		fmt.Println("2.Показать все книги")
		fmt.Println("3.Найти книгу")
		fmt.Println("4.Взять книгу")
		// fmt.Println("5.Вернуть книгу")
		// fmt.Println("6.Удалить книгу")
		fmt.Println("7.Выйти")

		fmt.Println("Что хотите выбрать?")
		fmt.Scanln(&value)

		switch value {
			case 1:
				addBook(&book)
			case 2:
				showBook(book)
			case 3:
				findBook(book)
			case 4:
				takeBook(book)
		// 	case 5:
		// 	case 6:
			case 7:
				fmt.Println("Выход из библиотеки")
				flag = true
			default:
				fmt.Println("Такого пункта нет, попробуйте еще раз")

		}
	}
	
}