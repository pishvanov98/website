package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"net/http"
)

type User struct {
	Name                  string
	Age                   uint16
	Money                 int16
	Avg_grades, Happiness float64
	Hobbies               []string
}
type UserDb struct {
	Name string `json:"name"`
}

func home_page(w http.ResponseWriter, r *http.Request) {
	bob := User{"Bob", 25, -50, 4.2, 0.8, []string{"Football", "Skate", "Dance"}}

	tmpl, _ := template.ParseFiles("templates/home_page.html")
	tmpl.Execute(w, bob)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Contacts page!")
}

func handleRequest() {
	http.HandleFunc("/", home_page)
	http.HandleFunc("/contacts/", contacts_page)
	http.ListenAndServe(":8080", nil)
}

func main() {

	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/golang")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	////Установка данных
	//insert, err := db.Query("INSERT INTO `users` (`name`) VALUES('Bobbys')")
	//
	//if err != nil {
	//	panic(err)
	//}
	//defer insert.Close()

	fmt.Println("Подключено к MySQL")

	res, err := db.Query("SELECT `name` FROM `users`")

	if err != nil {
		panic(err)
	}

	for res.Next() {
		var user UserDb
		err = res.Scan(&user.Name) //запись
		if err != nil {
			panic(err)
		}
		fmt.Println("Имя: " + user.Name)
	}

	//handleRequest()

}
