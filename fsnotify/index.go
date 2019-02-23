package main

import (
	"github.com/fsnotify/fsnotify"
	"log"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer watcher.Close()
	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				//log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("Write file:", event.Name)
				}
				if event.Op&fsnotify.Create == fsnotify.Create {
					log.Println("Create file:", event.Name)
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					log.Println("Remove file:", event.Name)
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					log.Println("Rename file:", event.Name)
				}
				if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					log.Println("Chmod file:", event.Name)
				}
			case wErr := <-watcher.Errors:
				log.Println("error:", wErr)
			}
		}
	}()
	err = watcher.Add("/www/2017/go/src/practice/fsnotify")
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
