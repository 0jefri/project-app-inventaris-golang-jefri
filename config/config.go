package config

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/spf13/viper"
)

func RunServer(routes *chi.Mux) {
	port := viper.GetString("PORT")

	if port == "" {
		port = "4000"
	}

	log.Println("Starting server on :" + port)
	if err := http.ListenAndServe(":"+port, routes); err != nil {
		log.Fatal(err)
	}
}

func Router() *chi.Mux {
	r := chi.NewRouter()

	// // middleware
	// r.Use(func(next http.Handler) http.Handler {
	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		ctx := context.WithValue(r.Context(), "username", "jefri")
	// 		next.ServeHTTP(w, r.WithContext(ctx))
	// 	})
	// })

	return r
}
