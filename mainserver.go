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

//Page title
type Page struct {
	Title string
}

//------------------------------Global Variables-------------------------------------//
//Compile templates on start
var templ = ParseTemplates()

//ParseTemplates use for multi directory html parsing
func ParseTemplates() *template.Template {
	t := template.New("Base")
	err := filepath.Walk("./", func(path string, info os.FileInfo, err error) error {
		if strings.Contains(path, ".html") {
			_, err = t.ParseFiles(path)
			fmt.Println(path)
			if err != nil {
				logger.Println(err)
				fmt.Println(err)
			}
		}

		return err
	})

	if err != nil {
		panic(err)
	}

	return t
}

//logging
var errorlog *os.File
var logger *log.Logger

//---------------------------------------Page Handlers----------------------------------//
//Handler for homepage
func homepageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home")
	err := templ.ExecuteTemplate(w, "index", &Page{Title: "Welcome to TL;DR"})
	if err != nil {
		fmt.Println(err)
		logger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Handler for about page
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("About")
	err := templ.ExecuteTemplate(w, "about", &Page{Title: "About TL;DR"})
	if err != nil {
		fmt.Println(err)
		logger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

//Server log to file
func init() {
	errorlog, err := os.OpenFile("serverlog.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Printf("Error opening file: %v", err)
		os.Exit(1)
	}
	log.SetOutput(errorlog)
	logger = log.New(errorlog, "\r\nTDLR : ", log.Lshortfile|log.LstdFlags)
}

func main() {
	//--------------------------------------Routers-------------------------------------//
	http.HandleFunc("/", homepageHandler)
	http.HandleFunc("/index", homepageHandler)
	http.HandleFunc("/about", aboutHandler)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	//---------------------------------------------------------------------------------//
	//start server
	fmt.Println("Starting server on port 9090")
	logger.Println("Starting server on port 9090")
	logger.Fatal(http.ListenAndServe(":9090", nil))

}
