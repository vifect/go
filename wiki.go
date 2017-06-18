package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)
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
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
