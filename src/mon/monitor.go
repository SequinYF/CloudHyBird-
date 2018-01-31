package m

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
)

const filepath  =  "/Users/sequin_yf/Git/CloudHyBird-/"

func client_monitor() {
	fmt.Println("mon run")
	watch, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watch.Close()

	done := make(chan bool)


	go func() {
		for {
			select {
			case ev := <-watch.Events:
				{
					go sendToserver(ev.name, ev.Op)
					if ev.Op & fsnotify.Create == fsnotify.Create  {
						log.Println("create file", ev.Name)
					}
					if ev.Op & fsnotify.Write == fsnotify.Write {
						log.Println("write file", ev.Name)
					}
					if ev.Op & fsnotify.Remove == fsnotify.Remove {
						log.Println("remove file", ev.Name)
					}
					if ev.Op & fsnotify.Rename == fsnotify.Rename {
						log.Println("rename file", ev.Name)
					}
					if ev.Op & fsnotify.Chmod == fsnotify.Chmod {
						log.Println("change chomd", ev.Name)
					}
				}
			case err := <- watch.Errors:
				{
					log.Println("errror", err)
					return
				}
			}
		}
	}()

	err = watch.Add(filepath)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}
