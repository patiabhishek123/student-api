package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/patiabhishek123/students-api/internal/config"
	"github.com/patiabhishek123/students-api/internal/http/handler/student"
	"github.com/patiabhishek123/students-api/internal/storage/sqlite"
)



func main(){
	//Steps
		//load config
		cfg :=config.MustLoad()

		//database setup
		storage,err:=sqlite.New(cfg)

		if err!=nil{
			log.Fatal(err)
		}
		slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
		//setup routers

		router:=http.NewServeMux()
		router.HandleFunc("POST /api/students", student.New(storage))
		//setup server

		server :=http.Server{
			Addr: cfg.Addr,
			Handler: router,
		}
		fmt.Println("Server started ",cfg.HTTPServer.Addr)
		done :=make(chan os.Signal,1)

		signal.Notify(done,os.Interrupt,syscall.SIGINT,syscall.SIGTERM)
		go func ()  {
			
			err :=server.ListenAndServe()
			if err !=nil {
				log.Fatal(" Failed to start server")
			}
		}()
		<-done

		slog.Info("shutting down the server")

		ctx,cancel :=context.WithTimeout(context.Background(),5*time.Second)
		defer cancel()
		if err:=server.Shutdown(ctx); err !=nil{
			slog.Error("failed to shutdoen server",slog.String("error",err.Error()))

		}
	
		slog.Info("server shutdown successfully")


		


}
