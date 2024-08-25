package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

const storePath = "../store"

var likeCounter int32 = 0

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
	tmpl := template.Must(template.ParseFiles("templates/audio-player.html"))
	tmpl.Execute(w, struct {
		File string
	}{
		File: "/store/" + item,
	})

}

func trackLikesHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Invalid form data", http.StatusBadRequest)
			return
		}

		likeCount, _ := strconv.Atoi(r.FormValue("likeCount"))
		dislikeCount, _ := strconv.Atoi(r.FormValue("dislikeCount"))
		fileName := r.FormValue("fileName")

		if likeCount > 0 {
			likeCounter += 1
		}
		if dislikeCount > 0 && likeCounter > 0 {
			likeCounter -= 1
		}

		fmt.Printf("Total likes for %s : %d \n", fileName, likeCounter)
		fmt.Fprintf(w, "Updated like counter: %d", likeCounter)
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	const PORT = ":8080"
	http.HandleFunc("/", listFileHandler)
	http.HandleFunc("/play", playFileHandler)
	http.Handle("/store/", http.StripPrefix("/store/", http.FileServer(http.Dir("../store"))))
	http.HandleFunc("/like", trackLikesHandler)
	err := http.ListenAndServe(PORT, nil)

	if err != nil {
		errorMessage := fmt.Errorf("error founds during server initialization: %v", err)
		fmt.Println(errorMessage)
	} else {
		fmt.Printf("Server is running at %v\n", PORT)
	}
}
