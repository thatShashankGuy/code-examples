package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
)

const storePath = "../store"

func listFiles() ([]string, error) {

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
		fmt.Printf("Error listing files:%v", err)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		http.Error(w, "Error Occured While Parsing the html template with list", http.StatusInternalServerError)
		fmt.Printf("Error Parsing files: %v", err)
		return
	}

	err = tmpl.Execute(w, list)
	if err != nil {
		http.Error(w, "Error Occured While Embedding the list in template", http.StatusInternalServerError)
		fmt.Printf("Error Executing template: %v", err)
		return
	}

}
func playFileHandler(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	cwd, err := os.Getwd()

	if err != nil {
		http.Error(w, "Error Occured While Playing the audio file", http.StatusInternalServerError)
		fmt.Printf("Error Playing File: %v", err)
		return
	}
	filePath := filepath.Join(cwd, storePath, item)

	file, err := os.Open(filePath)

	if err != nil {
		http.Error(w, "Error Occured While Playing the audio file", http.StatusInternalServerError)
		fmt.Printf("Error Playing File: %v", err)
		return
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "File stat error.", 500)
		return
	}
	fmt.Printf("Playing file: %s, Size: %d bytes\n", fileInfo.Name(), fileInfo.Size())

	http.ServeFile(w, r, filePath)

}

func main() {
	const PORT = ":8080"
	http.HandleFunc("/", listFileHandler)
	http.HandleFunc("/play", playFileHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Printf("Server is running at %v ", PORT)
	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		fmt.Printf("Cannot Start server %v", err)
	}
}
