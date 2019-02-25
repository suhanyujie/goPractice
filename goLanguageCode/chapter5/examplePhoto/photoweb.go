package main

import (
	"fmt"
	"html/template"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

const UPLOAD_DIR = "goLanguageCode/chapter5/examplePhoto/uploads"

func viewHandler(w http.ResponseWriter, r *http.Request) {
	imageId := r.FormValue("id")
	imagePath := UPLOAD_DIR + "/" + imageId
	if exists := fileExists(imagePath); !exists {
		http.NotFound(w, r)
	}
	w.Header().Set("Content-type", "image")
	http.ServeFile(w, r, imagePath)
}

func main() {
	const PORT = 8001
	http.HandleFunc("/upload", safeHandler(uploadHandler))
	http.HandleFunc("/view", safeHandler(viewHandler))
	http.HandleFunc("/list", safeHandler(listHandler))
	fmt.Println("http server in " + strconv.Itoa(PORT) + ".")
	err := http.ListenAndServe(":"+strconv.Itoa(PORT), nil)
	if err != nil {
		fmt.Println(err)
	}
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		renderHtml(w, "goLanguageCode/chapter5/examplePhoto/views/upload.html", nil)
		return
	} else if r.Method == "POST" {
		f, h, err := r.FormFile("image")
		if err != nil {
			fmt.Println(err)
			return
		}
		filename := h.Filename
		defer f.Close()
		t, err := os.Create(UPLOAD_DIR + "/" + filename)
		defer t.Close()
		if _, err := io.Copy(t, f); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/view?id="+filename, http.StatusFound)
	}
}

func listHandler(w http.ResponseWriter, r *http.Request) {
	fileInfoArr, err := ioutil.ReadDir(UPLOAD_DIR)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var locals = make(map[string]interface{})
	images := []string{}
	for _, fileInfo := range fileInfoArr {
		imageId := fileInfo.Name()
		images = append(images, imageId)
	}
	locals["images"] = images
	err = renderHtml(w, "goLanguageCode/chapter5/examplePhoto/views/list.html", locals)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func renderHtml(w http.ResponseWriter, path string, locals map[string]interface{}) error {
	t, err := template.ParseFiles(path)
	if err != nil {
		return err
	}
	return t.Execute(w, locals)
}

func safeHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err, ok := recover().(error); ok {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				log.Println("WARN:panic in %v-%v", fn, err)
			}
		}()
		fn(w, r)
	}
}

func checkError(err error) {
	if err != nil {
		panic(err)
		return
	}
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return os.IsExist(err)
}
