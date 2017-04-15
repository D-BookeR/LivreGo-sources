package main

import (
	"net/http"

	"github.com/D-BookeR/LivreGo-sources/etude_6/elections/controller"
)

func main() {
	http.HandleFunc("/winner", controller.Winner)
	http.HandleFunc("/count", controller.Count)
	http.HandleFunc("/registerVote", controller.RegisterVote)
	http.ListenAndServe(":8080", nil)
}
