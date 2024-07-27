package main

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"path/filepath"
)

func listFiles() ([]string, error) {
	const storePath = "../store/"
	var storeArray []string
	err := filepath.Walk(storePath, func(storePath string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			storeArray = append(storeArray, filepath.Base(storePath))
		}
		return nil
	})

	return storeArray, err
}

func listFileHandler(w http.ResponseWriter, r *http.Request) {
	list, err := listFiles()

	if err != nil {
		http.Error(w, "Error Occured While Parsing the list", http.StatusInternalServerError)
		log.Println("Error listing files:", err)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, "Error Occured While Parsing the html template with list", http.StatusInternalServerError)
		log.Println("Error Parsing files:", err)
		return
	}

	err = tmpl.Execute(w, list)
	if err != nil {
		http.Error(w, "Error Occured While Embedding the list in template", http.StatusInternalServerError)
		log.Println("Error Executing template:", err)
		return
	}

}

func main() {
	const PORT = ":8080"
	http.HandleFunc("/", listFileHandler)

	log.Printf("Server is running at %v ", PORT)
	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		log.Printf("Cannot Start server %v", err)
	}
}
