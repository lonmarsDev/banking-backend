package route

import (
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	// Add your routes here
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	return mux
}
