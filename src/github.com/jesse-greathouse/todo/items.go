package todo

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Items struct {
	Env *Environment
	Db  *Database
}

type Item struct {
	Id          int
	Description string
	Priority    int
	Completed   int
}

func (i Items) ItemsHandler(w http.ResponseWriter, r *http.Request) {

	items := i.Db.GetAllItems()

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(items)

	if err != nil {
		log.Println(err)
	}
}

func (i Items) ItemHandler(w http.ResponseWriter, r *http.Request) {

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
	}

	id, err := strconv.Atoi(ids[0])
	if err != nil {
		log.Println("Url Param 'id' unable to convert to int")
	}

	item := i.Db.GetItem(id)

	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w).Encode(item)

	if e != nil {
		log.Println(err)
	}
}

func (i Items) CreateHandler(w http.ResponseWriter, r *http.Request) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	newItem := i.Db.CreateItem(item)

	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w).Encode(newItem)

	if e != nil {
		log.Println(err)
	}
}

func (i Items) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	var item Item

	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	newItem := i.Db.UpdateItem(item)

	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w).Encode(newItem)

	if e != nil {
		log.Println(err)
	}
}

func (i Items) DeleteHandler(w http.ResponseWriter, r *http.Request) {

	ids, ok := r.URL.Query()["id"]
	if !ok || len(ids[0]) < 1 {
		log.Println("Url Param 'id' is missing")
	}

	id, err := strconv.Atoi(ids[0])
	if err != nil {
		log.Println("Url Param 'id' unable to convert to int")
	}

	response := i.Db.DeleteItem(id)

	w.Header().Set("Content-Type", "application/json")
	e := json.NewEncoder(w).Encode(response)

	if e != nil {
		log.Println(err)
	}
}
