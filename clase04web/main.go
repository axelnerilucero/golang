package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"clase_04_web/rutas"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
)

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", rutas.Home)
	mux.HandleFunc("/nosotros", rutas.Nosotros)
	mux.HandleFunc("/params/{id}/{slug}", rutas.Parametros)
	mux.HandleFunc("/parametros-querystring", rutas.ParametrosQueryString)

	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
		return
	}

	server := &http.Server{
		Addr:         "127.0.0.1:" + os.Getenv("PORT"),
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	fmt.Println("Corriendo servidor desde http://localhost:" + os.Getenv("PORT"))
	log.Fatal(server.ListenAndServe())

}

/*func main() {
	//mux := http.NewServeMux()
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Println(response, "Hola mundo")
	})
	fmt.Println("Corriendo en el servidor 8081")
	log.Fatal(http.ListenAndServe("Localhost:8081", nil))
}
*/
