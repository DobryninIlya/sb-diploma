package main

import (
	"fmt"
	"github.com/go-chi/chi"
	"log"
	"main/internal/ping"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const gate = "http://127.0.0.1:8383"

type App struct {
	router *chi.Mux
	done   chan os.Signal
	//store  map[int]city.City
}

func NewApp() *App {
	ret := &App{
		router: chi.NewRouter(),
		done:   make(chan os.Signal, 1),
	}
	signal.Notify(ret.done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	return ret
}

func (a *App) run() {
	a.router.Route("/api", func(r chi.Router) {
		r.Get("/", ping.New())
	})
	go func() {
		fmt.Println("Starting worker")
		log.Fatal(http.ListenAndServe(":8282", a.router))

	}()
	<-a.done
	fmt.Println("Exiting")
}
func main() {
	var app = NewApp()
	app.run()
}
