package main

import (
	"fmt"
	"net/http"

	"github.com/jakskal/simpleCRUD-go/userreview"
	"github.com/julienschmidt/httprouter"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)

	r.GET("/review", userreview.Index)
	r.POST("/review", userreview.Create)
	r.PATCH("/review", userreview.Update)
	r.DELETE("/review/:id", userreview.Delete)
	r.GET("/review/:id", userreview.Find)

	http.ListenAndServe(":3000", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// w.Write("welcome to crud preparing")
	fmt.Fprintf(w, "welcome to simple crud userreview")
}
