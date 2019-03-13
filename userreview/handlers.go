package userreview

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/julienschmidt/httprouter"
)

var wg sync.WaitGroup

//Index function so show all tags
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	queue := make(chan []review, 10)
	wg.Add(1)
	go func() {
		_, err := allReview(queue)
		if err != nil {
			http.Error(w, http.StatusText(204), http.StatusNoContent)
			log.Panicln(err)
			return

		}
		wg.Done()
	}()
	wg.Wait()
	close(queue)
	response, err := json.Marshal(<-queue)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		log.Panicln(err)
		return
	}
	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}

//Create New single tag
func Create(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	createdReview, err := createReview(r)
	if err != nil {
		http.Error(w, http.StatusText(422), http.StatusUnprocessableEntity)
		log.Panicln(err)
		return
	}

	response, err := json.Marshal(createdReview)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		log.Panicln(err)
		return
	}

	w.Header().Set("Content-Type", "application/vnd.api+json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s", response)
}

//Update for update review
func Update(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	updatedReview, err := updateReview(r)

	if err != nil {
		http.Error(w, http.StatusText(204), http.StatusNoContent)
		log.Panicln(err)
		return
	}

	response, err := json.Marshal(updatedReview)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		log.Panicln(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "\n%s", response)

}

func Delete(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	deletedReview, err := deleteReview(p)

	if err != nil {
		http.Error(w, http.StatusText(204), http.StatusNoContent)
		log.Panicln(err)
		return
	}

	response, err := json.Marshal(deletedReview)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		log.Panicln(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", response)

}

func Find(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	Tag, err := findReview(p)

	if err != nil {
		http.Error(w, http.StatusText(204), http.StatusNoContent)
		log.Panicln(err)
		return
	}

	response, err := json.Marshal(Tag)
	if err != nil {
		http.Error(w, http.StatusText(406), http.StatusNotAcceptable)
		log.Panicln(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s", response)

}
