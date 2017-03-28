package main

import(
  "fmt"
  "html"
  "log"
  "net/http"


  "html/template"
)

type Page struct{
  Title string
  Body []byte
}

func main(){

  //router := mux.NewRouter().StrictSlash(true)

  http.HandleFunc("/", homePage)
  http.HandleFunc("/coeurdalene", serveCDA)

  http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

  log.Fatal(http.ListenAndServe(":8080", nil))
}

func Index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func homePage(w http.ResponseWriter, r *http.Request){
  p := &Page{Title: "Home"}

  t, _ := template.ParseFiles("views/index.gohtml")
  t.Execute(w, p)
}

func serveCDA(w http.ResponseWriter, r *http.Request){
  p := &Page{Title: "Coeur D'Alene"}

  t, _ := template.ParseFiles("views/CDA.gohtml")
  t.Execute(w, p)
}
