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
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified file:", event.Name)
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
