package main

import ("fmt"; "net/http"; "html/template")

func main()  {
  handleRequest()
}

func handleRequest() {
  http.HandleFunc("/", home_page) // при перехое на главную страницу переходим в функцию
  http.HandleFunc("/contacts/", contacts_page)
  http.ListenAndServe(":8080", nil) // создаем локальный сервер
}

type User struct {
  Name string // название переменной и ее тип
  Age uint16 // целое число которое не может быть отрицательным (u)
  Money int16
  Avg_grades, Happiness float64 // средняя оценка и уровень счастья, число с точкой

}

func home_page(w http.ResponseWriter, r *http.Request) { // параметр, и запрос (вывод на экран)
  bob := User{"Bob", 25, -123, 4.6, 0.8}
  tmpl, _ := template.ParseFiles("templates/home_page.html") // две переменных, вторая не важна от того и такое название
  tmpl.Execute(w, bob) // вывод на страницу (куда писать и какие значения)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "This is contacts page")
}
