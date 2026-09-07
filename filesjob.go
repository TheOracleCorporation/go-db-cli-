package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)
//структура где хранятся наши данные 
type Person struct {
	Name string
	Age int
	City string
}
//функция для просмотра людей в базе
func showBase(people []Person) {
	for _, p := range people {
		p.FullInfo()
		fmt.Println("-------------")
	}
}
//функция для добавления людей в базу
func addBase(people []Person) []Person {
	var number int
	fmt.Println("Хотите добавить нового человека?")
	fmt.Println("1.Да")
	fmt.Println("2.Нет")
	fmt.Scanln(&number)

	switch number {
	case 1:
		var newName string
		var newAge int
		var newCity string

		fmt.Println("-----Добавление нового человека-----")
		fmt.Println("Как зовут?")
		fmt.Scanln(&newName)
		fmt.Println("Какой возраст?")
		fmt.Scanln(&newAge)
		fmt.Println("Из какого города?")
		fmt.Scanln(&newCity)

		newPerson := Person {
			Name: newName,
			Age: newAge,
			City: newCity,
		}

		people = append(people, newPerson)

		savePeople(people)

		case 2:
			break
	}
	return people
}
//функция для сохранения в файле
func savePeople(people []Person) {
	text := ""

	for _, p := range people {
		text += fmt.Sprintf("%s,%d,%s\n",p.Name,p.Age,p.City)
	}
	err := os.WriteFile("people.txt", []byte(text), 0644)
	if err != nil {
		fmt.Println("Ошибка сохранения файла:", err)
		return
	}
	fmt.Println("База успешно перезаписана")
}
//функция для поиска людей в базе
func findBase(people []Person) {
	var tempname string
	fmt.Println("кого хотите найти?")
	fmt.Scanln(&tempname)

	flag := false
	for _, p := range people {
		if tempname == p.Name {
			p.FullInfo()
			flag = true
		}
	}
	if !flag {
		fmt.Println("Такого человека нет")
	}
}
//функция для выдачи полной информации о людях
func (p Person) FullInfo() {
	fmt.Print("%s,%d,%s\n", p.Name, p.Age, p.City)
}
//функция для поиска самого молодого и возрастного человека + чредний возраст
func staticBase(people []Person) {
	if len(people) == 0 {
		fmt.Println("База пуста")
		return
	} else {
		young := people[0]
		for _, p := range people {
			if young.Age > p.Age {
				young = p
			}
		}
		fmt.Println("Самый молодой человек в базе:",young.Name,"-",young.Age,"из",young.City)

		old := people[0]
		for _, p := range people {
			if old.Age < p.Age {
				old = p
			}
		}
		fmt.Println("Самый возрастной человек в базе:",old.Name,"-",old.Age,"из",old.City)

		sum := 0
		value := 0
		for _, p := range people {
			sum += p.Age
			value++
		}
		fmt.Println("Средний возраст людей в базе:",sum / value)
	}
}
//поиск людей по городу
func searchCity(people []Person) {
	var foundMCity string
	fmt.Println("Какой город ищите?")
	fmt.Scanln(&foundMCity)

	flag := false
	fmt.Println("Все люди из",foundMCity,":")
	for _, p := range people {
		if p.City == foundMCity {
			fmt.Println(p.Name,"-",p.Age)
			flag = true
		}
	}
	if !flag {
		fmt.Println("Таких людей нет")
	}
}


//основная функция
func main() {
	//флаг для бесконечного цикла
	flag := false

	// Читаем весь файл
	data, err := os.ReadFile("people.txt")
	//если файл не откроется выдаем ошибку
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	// Превращаем байты в строку и разбиваем по переносу строки
	lines := strings.Split(string(data), "\n")
	//создаем нашу структуру
	people := []Person{}
	//скипаем пробелы
	for _, line := range lines {
		if line == "" {
			continue
		}
		//разделяем строку на три кусочка(имя , возраст и город)
		parts := strings.Split(line, ",")
		//скипаем если нет трех столбцов с критериями(имя , возраст и город) , чтоб программа не ушла в панику
		if len(parts) != 3 {
			continue
		}
		//преобразуем int в string
		age, err := strconv.Atoi(parts[1])
		//если нет возраста скипаем
		if err != nil {
			continue
		}
		//инициализируем нашу струтуру
		person := Person {
			Name: parts[0],
			Age: age,
			City: parts[2],
		}
		//добавялем наших людей из файла в структуру 
		people = append(people, person)
	}
	//запускаем интерактивчик
	for !flag {
		fmt.Println("1.Показать базу")
		fmt.Println("2.Добавить")
		fmt.Println("3.Найти")
		fmt.Println("4.Статистика")
		fmt.Println("5.Поиск по городам")
		fmt.Println("6.Выход")
		//выбор действия
		var count int
		fmt.Println("Выберите действие")
		fmt.Scanln(&count)
		//уходим в функции в зависимости от выбора действия 
		switch count {
			case 1:
				showBase(people)
			case 2:
				people = addBase(people)
			case 3:
				findBase(people)
			case 4:
				staticBase(people)
			case 5:
				searchCity(people)
			case 6:
				fmt.Println("Всего хорошего , до свидания")
				flag = true

		}
	}
}