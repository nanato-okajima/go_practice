package main

import (
	"bufio"
	"errors"
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"

	"gopractice/post"
	"gopractice/utils"
)

func main() {
	http.HandleFunc("/post", postHandler)
	http.HandleFunc("/new", newHandler)
	http.HandleFunc("/post/create", createHandler)
	_, err := os.Stdout.Write([]byte("Server Start..."))
	utils.Check(err)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	signatures := getStrings("docs/signatures.txt")
	fmt.Printf("%#v\n", signatures)
	html, err := template.ParseFiles("views/view.html")
	utils.Check(err)
	getPost := post.New(len(signatures), signatures)
	err = html.Execute(w, getPost)
	utils.Check(err)
}

func newHandler(w http.ResponseWriter, r *http.Request) {
	html, err := template.ParseFiles("views/new.html")
	utils.Check(err)
	err = html.Execute(w, nil)
	utils.Check(err)
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	post := r.FormValue("post")
	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("docs/signatures.txt", options, os.FileMode(0600))
	utils.Check(err)
	_, err = fmt.Fprintln(file, post)
	utils.Check(err)
	err = file.Close()
	utils.Check(err)
	http.Redirect(w, r, "/post", http.StatusFound)
}

func getStrings(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	if errors.Is(err, fs.ErrNotExist) {
		return nil
	}
	utils.Check(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	utils.Check(scanner.Err())
	return lines
}
