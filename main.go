package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jasonlvhit/gocron"
)

func main() {
	log.Println("---> Iniciando Cron Job...")
	s := gocron.NewScheduler()
	s.Every(1).Day().At("12:00").Do(call)
	s.Every(1).Day().At("18:00").Do(call)
	<-s.Start()
}

func call() {
	log.Println("---> Iniciando requisição...")
	res, err := http.Get("http://localhost:3000/")
	if err != nil {
		log.Printf("---> A requisição retornou erro %s\n", err)
	}
	result, _ := ioutil.ReadAll(res.Body)
	log.Println("--->", string(result))
}
