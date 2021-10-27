package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	hostname string
}

func (app *App) handler(w http.ResponseWriter, r *http.Request)  {
	bytes, err := fmt.Fprintf(w, "hello, world! I'm %s\n", app.hostname)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%d bytes written from host %s", bytes, app.hostname)
}

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- true
	}()

	file, err := os.Open("/etc/hostname")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	data, _, err := reader.ReadLine()
	if err != nil {
		log.Fatalln(err)
	}
	app := &App{hostname: string(data)}

	http.HandleFunc("/", app.handler)
	err = http.ListenAndServe(":8090", nil)
	if err != nil {
		log.Fatalln(err)
	}

	<-done
}
