package main

import ("fmt"
        "net/http")

func main()  {
  handleRequest()
}

func handleRequest() {
  http.HandleFunc("/", home_page) // при перехое на главную страницу переходим в функцию
  http.HandleFunc("/contacts/", contacts_page)
  http.ListenAndServe(":8080", nil) // создаем локальный сервер
}

func home_page(w http.ResponseWriter, r *http.Request) { // параметр, и запрос (вывод на экран)
  fmt.Fprintf(w, "Go is super easy!")
}

func contacts_page(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "This is contacts page")
}
