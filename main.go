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
  port := ":8089"

  //router := mux.NewRouter().StrictSlash(true)

  http.HandleFunc("/", homePage)
  http.HandleFunc("/coeurdalene", serveCDA)
  http.HandleFunc("/postfalls", servePF)
  http.HandleFunc("/lakefernan", serveLF)
  http.HandleFunc("/haydenlake", serveHL)
  http.HandleFunc("/rathdrum", serveRD)
  http.HandleFunc("/sandpoint", serveSP)
  http.HandleFunc("/about", serveAboutPage)

  http.Handle("/resources/", http.StripPrefix("/resources/", http.FileServer(http.Dir("resources"))))

  log.Fatal(http.ListenAndServe(port, nil))
}

func Index(w http.ResponseWriter, r *http.Request){
  fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func homePage(w http.ResponseWriter, r *http.Request){
  p := &Page{Title: "North Idaho"}

  t, _ := template.ParseFiles("views/index.gohtml")
  t.Execute(w, p)
}

func serveCDA(w http.ResponseWriter, r *http.Request){
  p := &Page{Title: "Coeur D'Alene"}
  t, _ := template.ParseFiles("views/CDA.gohtml")
  t.Execute(w, p)
}
func servePF(w http.ResponseWriter, r *http.Request){
  p := &Page{Title: "Post Falls"}
  t, _ := template.ParseFiles("views/postfalls.gohtml")
  t.Execute(w, p)
}
func serveLF(w http.ResponseWriter, r *http.Request){
  p := &Page{Title: "Lake Fernan"}
  t, _ := template.ParseFiles("views/lakefernan.gohtml")
  t.Execute(w, p)
}
func serveHL(w http.ResponseWriter, r *http.Request){
  p := &Page{Title: "Hayden Lake"}
  t, _ := template.ParseFiles("views/haydenlake.gohtml")
  t.Execute(w, p)
}
func serveRD(w http.ResponseWriter, r *http.Request){
  p := &Page{Title: "Rathdrum"}
  t, _ := template.ParseFiles("views/rathdrum.gohtml")
  t.Execute(w, p)
}
func serveSP(w http.ResponseWriter, r *http.Request){
  p := &Page{Title: "Sandpoint"}
  t, _ := template.ParseFiles("views/sandpoint.gohtml")
  t.Execute(w, p)
}

func serveAboutPage(w http.ResponseWriter, r *http.Request){
  p := &Page{Title: "About"}
  t, _ := template.ParseFiles("views/about.gohtml")
  t.Execute(w, p)
}
