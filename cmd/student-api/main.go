package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/patiabhishek123/students-api/internal/config"
)



func main(){
	//Steps
		//load config
		cfg :=config.MustLoad()

		//database setup
		//setup routers

		router:=http.NewServeMux()
		router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("welcome to student api "))
		})
		//setup server

		server :=http.Server{
			Addr: cfg.Addr,
			Handler: router,
		}
		fmt.Println("Server started ",cfg.HTTPServer.Addr)
		err :=server.ListenAndServe()
		if err !=nil {
			log.Fatal(" Failed to start server")
		}

		


}
