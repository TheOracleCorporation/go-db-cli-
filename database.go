package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var db *sql.DB

func addPeople() {
	var name string
	var age int
	var city string

	fmt.Println("Ваше имя:")
	fmt.Scanln(&name)
	fmt.Println("Ваш возраст:")
	fmt.Scanln(&age)
	fmt.Println("Ваш город:")
	fmt.Scanln(&city)

	_, err := db.Exec(
		"INSERT INTO people(name, age, city) VALUES(?, ?, ?)",
		name,
		age,
		city,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Пользователь успешно добавлен")

}

func showPeople() {
	rows, err := db.Query("SELECT * FROM people")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	fmt.Println("\n-----Список людей-----")
	for rows.Next() {
		var id int
		var name string
		var age int
		var city string

		err := rows.Scan(&id, &name, &age, &city)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d | Имя: %s | Возраст: %d | Город: %s\n", id, name, age, city)
	}	
}

func findPeople() {
	var value int
	fmt.Println("По какому параметру будем искать?")
	fmt.Println("1.По имени")
	fmt.Println("2.По возрасту")
	fmt.Println("3.По городу")
	fmt.Scanln(&value)

	switch value {
		case 1:
			findName()
		case 2:
			findAge()
		case 3:
			findCity()
	}
}

func findName() {
	var tempName string
	fmt.Println("Как зовут пользователя?")
	fmt.Scanln(&tempName)

	rows, err := db.Query("SELECT name, age, city FROM people")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	flag := false

	var foundName string
	var foundAge int
	var foundCity string

	for rows.Next() {
		var name string
		var age int
		var city string

		err := rows.Scan(&name, &age, &city)
		if err != nil {
			log.Fatal(err)
		}

		if tempName == name {
			flag = true
			foundName = name
			foundAge = age
			foundCity = city
			break
		}
	}

	if flag == true {
		fmt.Printf("Пользователь найден! Имя: %s, Возраст: %d, Город: %s\n", foundName, foundAge, foundCity)
	} else {
		fmt.Println("Такого человека нет в базе данных.")
	}

}

func findAge() {
	var tempAge int
	fmt.Println("Сколько лет?")
	fmt.Scanln(&tempAge)

	rows, err := db.Query("SELECT name, age, city FROM people")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	flag := false

	var foundName string
	var foundAge int
	var foundCity string

	for rows.Next() {
		var name string
		var age int
		var city string

		err := rows.Scan(&name, &age, &city)
		if err != nil {
			log.Fatal(err)
		}

		if tempAge == age {
			flag = true
			foundName = name
			foundAge = age
			foundCity = city
			break
		}
	}

	if flag == true {
		fmt.Printf("Пользователь найден! Имя: %s, Возраст: %d, Город: %s\n", foundName, foundAge, foundCity)
	} else {
		fmt.Println("Такого человека нет в базе данных.")
	}
}

func findCity() {
	var tempCity string
	fmt.Println("Какой город?")
	fmt.Scanln(&tempCity)

	rows, err := db.Query("SELECT name, age, city FROM people")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	flag := false

	var foundName string
	var foundAge int
	var foundCity string

	for rows.Next() {
		var name string
		var age int
		var city string

		err := rows.Scan(&name, &age, &city)
		if err != nil {
			log.Fatal(err)
		}

		if tempCity == city {
			flag = true
			foundName = name
			foundAge = age
			foundCity = city
			break
		}
	}

	if flag == true {
		fmt.Printf("Пользователь найден! Имя: %s, Возраст: %d, Город: %s\n", foundName, foundAge, foundCity)
	} else {
		fmt.Println("Такого человека нет в базе данных.")
	}
}

func changePeople() {
	var number int 
	fmt.Println("Что хотите изменить?")
	fmt.Println("----------")
	fmt.Println("1.Имя")
	fmt.Println("2.Возраст")
	fmt.Println("3.Город")
	fmt.Scanln(&number)

	switch number {
		case 1:
			changeName()
		case 2:
			changeAge()
		case 3:
			changeCity()
	}
}

func changeName() {
	var oldName string
    var newName string

    fmt.Println("Чье имя вы хотите изменить (текущее имя)?")
    fmt.Scanln(&oldName)

    fmt.Println("Введите новое имя:")
    fmt.Scanln(&newName)

	_, err := db.Exec(
		"UPDATE people SET name = ? WHERE name = ?",
		newName,
		oldName,
	)

	if err != nil {
        fmt.Println("Ошибка при обновлении имени:", err)
        return
    }

    fmt.Println("Имя успешно изменено!")
}

func changeAge() {
	var name string
    var age int

    fmt.Println("Чей возраст хотите заменить (укажите текущее имя)?")
    fmt.Scanln(&name)

    fmt.Println("Введите новый возраст:")
    fmt.Scanln(&age)

	_, err := db.Exec(
		"UPDATE people SET age = ? WHERE name = ?",
		age,
		name,
	)

	if err != nil {
        fmt.Println("Ошибка при обновлении возраста:", err)
        return
    }

    fmt.Println("Возраст успешно изменен!")
}

func changeCity() {
	var name string
    var city string

    fmt.Println("Чей город хотите заменить (укажите текущее имя)?")
    fmt.Scanln(&name)

    fmt.Println("Введите новый город:")
    fmt.Scanln(&city)

	_, err := db.Exec(
		"UPDATE people SET city = ? WHERE name = ?",
		city,
		name,
	)

	if err != nil {
        fmt.Println("Ошибка при обновлении города:", err)
        return
    }

    fmt.Println("Город успешно изменен!")
}

func deletePeople() {
	var name string
	fmt.Println("Кого хотите удалить?")
	fmt.Scanln(&name)
	
	_, err := db.Exec(
		"DELETE FROM people WHERE name = ?",
		name,
	)
	if err != nil {
		fmt.Println("Ошибка в удалении пользователя", err)
	}

	fmt.Println("Пользователь успешно удален")
}

func statisticsPeople() {
	for {
		var number int
		fmt.Println("Что Вас интреисует?")
		fmt.Println("1.Сколько людей в базе")
		fmt.Println("2.Средний возраст людей в базе")
		fmt.Println("3.Максимальный возраст человека в базе")
		fmt.Println("4.Минимальный возраст человека в базе")
		fmt.Println("5.Выход из статистики")
		fmt.Scanln(&number)

		switch number {
			case 1:
				howmanyPeople()
			case 2:
				avgPeople()
			case 3:
				maxPeople()
			case 4:
				minPeople()
			case 5:
				fmt.Println("Вы вышли из статистики")
				return
		}
	}
}

func howmanyPeople() {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM people").Scan(&count)
	if err != nil {
		log.Println("Ошибка при получении количества:", err)
		return
	}
	fmt.Printf("Всего людей в базе: %d\n\n", count)
}

func avgPeople() {
	var avg float64
	err := db.QueryRow("SELECT AVG(age) FROM people").Scan(&avg)
	if err != nil {
		log.Println("Ошибка при расчете среднего возраста:", err)
		return
	}
	fmt.Printf("Средний возраст людей: %.1f\n\n", avg)
}

func maxPeople() {
	var max int
	err := db.QueryRow("SELECT MAX(age) FROM people").Scan(&max)
	if err != nil {
		log.Println("Ошибка при поиске максимального возраста:", err)
		return
	}
	fmt.Printf("Максимальный возраст в базе: %d\n\n", max)
}

func minPeople() {
	var min int
	err := db.QueryRow("SELECT MIN(age) FROM people").Scan(&min)
	if err != nil {
		log.Println("Ошибка при поиске минимального возраста:", err)
		return
	}
	fmt.Printf("Минимальный возраст в базе: %d\n\n", min)
}


func main() {
	var err error 
	db, err = sql.Open("sqlite", "people.db")
	if err != nil {
		log.Fatal(err)
	}

	query := `
	CREATE TABLE IF NOT EXISTS people(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		age INTEGER,
		city TEXT
	);
	`

	_, err = db.Exec(query) 
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Таблица успешно создана!")

	infinity := false

	for !infinity {
		var number int

		fmt.Println("1.Добавить человека")
		fmt.Println("2.Показать всех людей")
		fmt.Println("3.Найти человека")
		fmt.Println("4.Изменить параментр человека")
		fmt.Println("5.Удалить человека")
		fmt.Println("6.Показать статистику")
		fmt.Println("7.Выход")

		fmt.Scanln(&number)

		switch number {
			case 1:
				addPeople()
			case 2:
				showPeople()
			case 3:
				findPeople()
			case 4:
				changePeople()
			case 5:
				deletePeople()
			case 6:
				statisticsPeople()
			case 7:
				fmt.Println("Выход из программы")
				infinity = true
		}
	}
}