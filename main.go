package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

const HOMEPAGE = "https://bioinf.shenwei.me/todo"

var dbPath = "./db"

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

		r.Route(`/{itemID:\d+}`, func(r chi.Router) {
			r.Get("/", getItem)
			r.Put("/", updateItem)
			r.Delete("/", deleteItem)
		})

		r.Get("/search", searchItems)
	})

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	go func() {
		closed := false
		for range signalChan {
			log.Println("received an interrupt")

			if closed {
				os.Exit(0)
			}

			log.Printf("closing db: %s", dbPath)
			err = db.Close()
			if err != nil {
				log.Fatalf("fail to close db")
				return
			}
			log.Printf("db closed: %s", dbPath)

			closed = true
			break
		}
	}()

	http.ListenAndServe(":8080", r)
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
	w.Write([]byte(item.String()))
}

func getItem(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, http.StatusText(422), 422)
		return
	}
	item, err := db.GetItem([]byte(id))
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	w.Write([]byte(item.String()))
}

func getItems(w http.ResponseWriter, r *http.Request) {

}

func updateItem(w http.ResponseWriter, r *http.Request) {

}

func deleteItem(w http.ResponseWriter, r *http.Request) {

}

func searchItems(w http.ResponseWriter, r *http.Request) {

}
