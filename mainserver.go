package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Page struct {
	Title string
}

//-------------------------------------------------------------------------------------//
//Compile templates on start
var templ = func() *template.Template {
	t := template.New("")
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			fmt.Println(path)
			_, err = t.ParseFiles(path)
			if err != nil {
				fmt.Println(err)
			}
		}
		return err
	})

	if err != nil {
		panic(err)
	}
	return t
}()

//---------------------------------------Page Handlers----------------------------------//
//Handler for homepage
func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	err := templ.ExecuteTemplate(w, "index", &Page{Title: "Welcome to TL;DR"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Handler for about page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About")
	err := templ.ExecuteTemplate(w, "about", &Page{Title: "About TL;DR"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Handler for test Page
func testHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("Test")
		testT, _ := template.ParseFiles("static/test.html")
		testT.Execute(w, nil)
	}
}

func main() {
	//--------------------------------------Routers-------------------------------------//
	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/test", testHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//---------------------------------------------------------------------------------//

	//log to file
	f, err := os.OpenFile("serverlog.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	logger := log.New(f, "Logged : ", log.LstdFlags)
	log.SetOutput(f)

	//start server
	logger.Println("Starting server on port 9090")
	logger.Fatal(http.ListenAndServe(":9090", nil))

}
