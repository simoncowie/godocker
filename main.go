package main

import (
	"fmt"
	"net/http"
	"time"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {

	locTime := time.Now()
	nzLoc, _ := time.LoadLocation("Pacific/Auckland")
	nzTime := locTime.In(nzLoc)

	fmt.Fprintf(w, "Hello world, it's %s. That's %s New Zealand time!", locTime.Format("2006-01-02 3:04PM"), nzTime.Format("2006-01-02 3:04PM"))

}

func main() {
	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":8080", nil)
}
