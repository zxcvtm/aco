package routes

import (
	// Standard library packages
	"net/http"
	_"encoding/json"

	// Third party packages
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	"encoding/json"
	"aco/workspace/controllers"
	"github.com/kr/pretty"
)

// *****************************************************************************
// API Routes
// *****************************************************************************

func ACOApi(r *mux.Router) {

	r.Handle("/aco",negroni.New(
		negroni.Wrap(ACO),
	)).Methods("POST")

}

var ACO = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	type Request struct {
		Graph [][]float64 `json:"graph"`
	}
	req := Request{}
	json.NewDecoder(r.Body).Decode(&req)
	pretty.Println(req)
	controllers.AcoAlgorithm(req.Graph).JsonResponse(w)
	return
})