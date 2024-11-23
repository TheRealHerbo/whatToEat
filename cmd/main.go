package main

import (
	"log"
	"net/http"

	"github.com/HergyoBotond/whatToEat/view"
	"github.com/HergyoBotond/whatToEat/view/layout"
	"github.com/HergyoBotond/whatToEat/view/partial"

	"github.com/a-h/templ"
)

func main() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))

    c := layout.Base(view.Index())
    http.Handle("/", templ.Handler(c))

    http.Handle("/recipies", templ.Handler(partial.Foo()))

    log.Fatal(http.ListenAndServe(":8080", nil))
}
