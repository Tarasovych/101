package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
)

type App struct {
	hostname string
}

func (app *App) handler(w http.ResponseWriter, r *http.Request)  {
	bytes, err := fmt.Fprintf(w, "hello, world! I'm %s\n", app.hostname)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%d bytes written", bytes)
}

func main() {
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
	log.Fatalln(http.ListenAndServe(":8090", nil))
}
