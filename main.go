package main

import ("fmt"
        "net/http")

func main()  {
  // первый способ создания объекта, название любое, наследуем от User и заполняем данными
  //var bob User = ...

  // второй способ
  //bob := User {name: 'Bob', age: 23, money: -123, avg_grades: 4,6, happiness: 0.8}

  // третий способ
  //bob := User{"Bob", 25, -123. 4.6, 0.8}

  handleRequest()
}

func handleRequest() {
  http.HandleFunc("/", home_page) // при перехое на главную страницу переходим в функцию
  http.HandleFunc("/contacts/", contacts_page)
  http.ListenAndServe(":8080", nil) // создаем локальный сервер
}

type User struct {
  name string // название переменной и ее тип
  age uint16 // целое число которое не может быть отрицательным (u)
  money int16
  avg_grades, happiness float64 // средняя оценка и уровень счастья, число с точкой

}

// метод для структуры(объекта) - возвращаем тип (строку)
func (u User) getAllInfo() string {
  return fmt.Sprintf("User name is: %s. He is %d and " +
    "he has money equal: %d", u.name, u.age, u.money) // в строку подставляем данные из структуры
}

func (u *User) setNewName(newName string) { // ссылка на сам объект, ссылочный тип данных
  u.name = newName
}

func home_page(w http.ResponseWriter, r *http.Request) { // параметр, и запрос (вывод на экран)
  bob := User{"Bob", 25, -123, 4.6, 0.8}
  //bob.name = "Alex" // переопределяем переменную
  bob.setNewName("Alex") // переопределяем

  fmt.Fprintf(w, bob.getAllInfo()) // обращаемся к методу структуры(объекта)
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "This is contacts page")
}
