//Author: Kevin Delassus
//Problem Sheet 2
//This problem set is for you to learn the fundamentals of creating a web application in Go. 
//Create a single Git repository as your submission, complete with README and gitignore files. 
//NB: after completing each exercise, commit your code - there should be at least one commit per exercise. 
//dYou be will required to submit a URL to the repository and the use of GitHub is recommended for this purpose. 
//All code should be fully commented, and the README should explain how to clone your repository and run the code.

package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"bytes"
)
type Page struct {
	Title string
	Body  []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Guessing Game</h1>")

	body := new(bytes.Buffer)
	body.ReadFrom(r.Body)
	fmt.Fprintln(w, "r.Body:             ",  body.String())
}
func editHandler(w http.ResponseWriter, r *http.Request) {
    title := r.URL.Path[len("/bootstrapTemplate/"):]
    p, err := loadPage(title)
    if err != nil {
        p = &Page{Title: title}
    }
    t, _ := template.ParseFiles("bootstrapTemplate.html")
    t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/edit/", editHandler)
	http.ListenAndServe(":8080", nil)
}