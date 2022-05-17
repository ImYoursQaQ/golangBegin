package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

type Page struct {
	Title string
	Body  []byte
}

type IPAddr [4]byte

func (ip IPAddr) String() string {
	var strs []string
	for i, num := range ip {
		strs[i] = string(num)
	}
	return strings.Join(strs, ".")
}

var a uint = 100
var templates = template.Must(template.ParseFiles("tmpl/edit.html", "tmpl/view.html"))

var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")

var dataPath = "data/"

func (p *Page) save() error {
	filename := dataPath + p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := dataPath + title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

//demo handler
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[2:])
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	err := templates.ExecuteTemplate(w, tmpl+".html", p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}

func Pic(dx, dy int) [][]uint8 {

	px := make([][]uint8, dy)
	for i := range px {
		py := make([]uint8, dx)
		for j := range py {
			py[j] = uint8(i * j)
		}
		px[i] = py
	}
	return px
}

func WordCount(s string) map[string]int {
	strs := strings.Split(s, " ")
	m := make(map[string]int)
	for _, e := range strs {
		elem, _ := m[e]
		m[e] = elem + 1
	}
	return m
}

func main() {
	//http.HandleFunc("/", handler)
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	http.HandleFunc("/save/", makeHandler(saveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}

//func main() {
//	p1 := &Page{Title: "TestPage", Body: []byte("This is a sample Page.")}
//	p1.save()
//	p2, _ := loadPage("TestPage")
//	fmt.Println(string(p2.Body))
//}
