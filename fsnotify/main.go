package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fsnotify/fsnotify"
)


func main() {
	filename := "./test"
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatalln(err)
	}
	defer watcher.Close()
	err = waitUntilFind(filename)
	if err != nil {
		log.Fatalln(err)
	}
	err = watcher.Add(filename)
	if err != nil {
		log.Fatalln(err)
	}

	renameCh := make(chan bool)
	removeCh := make(chan bool)
	errCh := make(chan error)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				switch {
				case event.Op&fsnotify.Write == fsnotify.Write:
					fmt.Printf("Write:  %s: %s\n", event.Op, event.Name)
				case event.Op&fsnotify.Create == fsnotify.Create:
					fmt.Printf("Create: %s: %s\n", event.Op, event.Name)
				case event.Op&fsnotify.Remove == fsnotify.Remove:
					fmt.Printf("Remove: %s: %s\n", event.Op, event.Name)
					removeCh <- true
				case event.Op&fsnotify.Rename == fsnotify.Rename:
					fmt.Printf("Rename: %s: %s\n", event.Op, event.Name)
					renameCh <- true
				case event.Op&fsnotify.Chmod == fsnotify.Chmod:
					fmt.Printf("Chmod:  %s: %s\n", event.Op, event.Name)
				}
			case err := <-watcher.Errors:
				errCh <- err
			}
		}
	}()
	go func() {
		for {
			select {
			case <-renameCh:
				fmt.Println("rn wait",filename)
				err := waitUntilFind(filename)
				if err != nil {
					log.Fatalln(err)
				}
				err = watcher.Add(filename)
				if err != nil {
					log.Fatalln(err)
				}
			case <-removeCh:
				fmt.Println("rv wait",filename)
				err := waitUntilFind(filename)
				if err != nil {
					log.Fatalln(err)
				}
				err = watcher.Add(filename)
				if err != nil {
					log.Fatalln(err)
				}
			}
		}
	}()
	<-errCh
}
func waitUntilFind(filename string) error {
	for {
		time.Sleep(1 * time.Second)
		_, err := os.Stat(filename)
		if err != nil {
			if os.IsNotExist(err) {
				continue
			} else {
				return err
			}
		}
		break
	}
	return nil
}