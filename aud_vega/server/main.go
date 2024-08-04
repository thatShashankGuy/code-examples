package main

import (
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const storePath = "../store/"

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
func playFileHandler(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")

	filePath := filepath.Join(storePath, item)

	file, err := os.Open(filePath)

	if err != nil {
		http.Error(w, "Error Occured While Playing the audio file", http.StatusInternalServerError)
		log.Println("Error Playing File:", err)
		return
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		http.Error(w, "File stat error.", 500)
		return
	}

	w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Disposition", "inline")
	http.ServeContent(w, r, item, fileInfo.ModTime(), file)
}

func main() {
	const PORT = ":8080"
	http.HandleFunc("/", listFileHandler)
	http.HandleFunc("/play", playFileHandler)

	log.Printf("Server is running at %v ", PORT)
	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		log.Printf("Cannot Start server %v", err)
	}
}
