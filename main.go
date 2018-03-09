package main
import (
	// Standard
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"aco/workspace/routes"
)

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, string("Is up and running successfully!"))
	})

	n := negroni.Classic() // Includes some default middlewares

	n.Use(cors.New(cors.Options{
		AllowedOrigins   : []string{"*"},
		AllowedMethods   : []string{"HEAD","GET","POST","PUT","DELETE","PATCH","OPTIONS"},
		AllowedHeaders   : []string{"Origin","Authorization","X-Requested-With","Content-Type","Accept","Signature"},
		ExposedHeaders   : []string{"Content-Length"},
		AllowCredentials : true,
	}))

	n.UseHandler(r)

	//Routes
	apiRoutes(r)

	//Run server
	fmt.Println("Its works")
	http.ListenAndServe( ":3000", n)
}
func apiRoutes(r *mux.Router) {
	routes.ACOApi(r)
}