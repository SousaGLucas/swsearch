package rest

// package responsible for controlling access to the app by REST API

import (
	"net/http"

	"github.com/SousaGLucas/swsearch/app/gateway/controller" // responsible for accessing the domain methods
	"github.com/SousaGLucas/swsearch/log"                    // package responsible for logging system errors
	"github.com/go-chi/chi"                                  // package responsible for manage app routes
)

func Execute() {
	r := chi.NewRouter()

	r.Get("/", helpMessage)         // main route
	r.Get("/help", helpMessage)     // help route
	r.Get("/search/{term}", search) // search data route

	r.Route("/cache", func(r chi.Router) {
		r.Get("/", getCache)      // cache data route
		r.Delete("/", clearCache) // clear cache route
	})

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.SetLog(err)
	}
}

// responsible for serving the help message
func helpMessage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(helpMessageText))
}

// responsible for serving the search data
func search(w http.ResponseWriter, r *http.Request) {
	searchTerm := chi.URLParam(r, "term")

	data, err := controller.Search(searchTerm) // call search data function

	if err != nil {
		switch errString := err.Error(); errString {
		case "'" + searchTerm + "'" + " already searched":
			http.Error(w, err.Error(), 403)
		default:
			http.Error(w, err.Error(), 500)
		}

		return
	}

	w.Write([]byte(data))
	return
}

// responsible for serving the cache data
func getCache(w http.ResponseWriter, r *http.Request) {
	cache, err := controller.GetCache() // call get cache function

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte(cache))
	return
}

// responsible for clear cache data
func clearCache(w http.ResponseWriter, r *http.Request) {
	err := controller.ClearChache() // call clear cache function

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.Write([]byte("Cache deleted!"))
	return
}

// help message
var helpMessageText string = `
	SWSEARCH
	App for searching data from the Star Wars universe.
	
	Version: development

	Routes:
		- GET /search/:term		returns data for the search term 
		- GET /cache			returns cache data
		- GET /help				returns help message
		- DELETE /cache			clear cache data
`
