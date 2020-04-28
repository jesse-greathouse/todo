// +build ignore

package main

import (
	"log"
	"net/http"

	"github.com/jesse-greathouse/todo"
)

func main() {
	// Create and initialize the runtime environment
	env := &todo.Environment{}
	env.Init()

	// Create and initialize the database
	db := &todo.Database{Env: env}
	db.Init()

	// Set up todo.Items handlers
	items := todo.Items{Env: env, Db: db}

	http.HandleFunc("/item/", items.ItemHandler)
	http.HandleFunc("/item/create/", items.CreateHandler)
	http.HandleFunc("/item/update/", items.UpdateHandler)
	http.HandleFunc("/item/delete/", items.DeleteHandler)
	http.HandleFunc("/items/", items.ItemsHandler)

	// Set up static file server
	http.Handle("/", http.FileServer(http.Dir(env.STATIC)))

	// Some endpoints are meant for the frontend only
	// All frontend endpoints should serve index.html
	// The Angular router will handle these views
	http.HandleFunc("/todo", frontendHandler)
	http.HandleFunc("/dashboard", frontendHandler)
	http.HandleFunc("/detail/", frontendHandler)

	// run server
	log.Println("Serving static files from " + env.STATIC)
	log.Println("Listening on :" + env.PORT + "...")
	log.Fatal(http.ListenAndServe(":"+env.PORT, nil))
}

func frontendHandler(w http.ResponseWriter, r *http.Request) {
	env := &todo.Environment{}
	env.Init()
	http.ServeFile(w, r, env.STATIC+"/index.html")
}
