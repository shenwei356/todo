package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/valve"
)

const HOMEPAGE = "https://bioinf.shenwei.me/todo"

var dbPath = "db/todo.db"

var db *ItemDB

func main() {
	var err error
	db, err = Connect(dbPath)
	if err != nil {
		log.Fatalf("fail to connect db: %s", dbPath)
		return
	}
	log.Printf("db connected: %s", dbPath)

	defer func() {
		log.Printf("closing db: %s", dbPath)
		err = db.Close()
		if err != nil {
			log.Fatalf("fail to close db")
			return
		}
		log.Printf("db closed: %s", dbPath)
	}()

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Throttle(2))

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, HOMEPAGE, http.StatusSeeOther)
	})

	r.Route("/items", func(r chi.Router) {
		r.Post("/", putItem)
		r.Get("/", getItems)

		r.Route(`/{id:\d+}`, func(r chi.Router) {
			r.Get("/", getItem)
			r.Put("/", updateItem)
			r.Delete("/", deleteItem)
		})

		r.Get("/search", searchItems)
	})

	baseCtx := valve.New().Context()
	server := http.Server{Addr: ":8080", Handler: chi.ServerBaseContext(baseCtx, r)}

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		select {
		case <-signalChan:
			log.Print("closing server")
			server.Shutdown(baseCtx)
			log.Print("server closed")

			return
		}
	}()
	log.Print("server started")
	server.ListenAndServe()
}

func putItem(w http.ResponseWriter, r *http.Request) {
	task := r.FormValue("task")
	if task == "" {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	item, err := db.PutItem([]byte(task))
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write([]byte(item.String() + "\n"))
}

func getItem(w http.ResponseWriter, r *http.Request) {
	idS := chi.URLParam(r, "id")
	if idS == "" {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	id, err := strconv.Atoi(idS)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	item, err := db.GetItem(id)
	if err != nil {
		http.Error(w, http.StatusText(404), 404)
		return
	}
	w.Write([]byte(item.String() + "\n"))
}

func getItems(w http.ResponseWriter, r *http.Request) {

}

func updateItem(w http.ResponseWriter, r *http.Request) {

}

func deleteItem(w http.ResponseWriter, r *http.Request) {

}

func searchItems(w http.ResponseWriter, r *http.Request) {

}
