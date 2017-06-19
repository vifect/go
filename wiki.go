package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/view/", viewHanlder)
	http.HandleFunc("/edit/", viewHanlder)
	http.ListenAndServe(":8080", nil)
	//p2, _ := loadPage("TestPage")
	//fmt.Println(string(p2.Body))
	//p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}

}

// data sructures
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filname := title + ".txt"
	body, err := ioutil.ReadFile(filname)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

// net/http
func handler(w http.ResponseWriter, r *http.Request) {
	//p2, _ := loadPage("TestPage")
	//fmt.Fprintf(w, p2.Body, r.URL.Path[1:])
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func viewHanlder(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, err := loadPage(title)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	renderTemplate(w, "view", p)
}

func editHanlder(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func renderTemplate(w http.ResponseWriter, templName string, p *Page) {
	templ, _ := template.ParseFiles(templName + ".html")
	templ.Execute(w, p)
}
