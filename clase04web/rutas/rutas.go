package rutas

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Hola mundo")
}

func Nosotros(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Pagina nosotros con fresh")
}

func Parametros(response http.ResponseWriter, request *http.Request) {
	// recibir argumentos
	vars := mux.Vars(request)
	fmt.Fprintln(response, "ID = "+vars["id"]+" | Slug = "+vars["slug"])
}

func ParametrosQueryString(response http.ResponseWriter, request *http.Request) {
	// recibir argumentos
	fmt.Fprintln(response, request.URL)
	fmt.Fprintln(response, request.URL.RawQuery)
	fmt.Fprintln(response, request.URL.Query())
	fmt.Fprintln(response, request.URL.Query().Get("id"))
	fmt.Fprintln(response, request.URL.Query().Get("slug"))
	id := request.URL.Query().Get("id")
	slug := request.URL.Query().Get("slug")
	fmt.Fprintln(response, "-------------------------------")
	fmt.Printf("ID = %s | slug = %s", id, slug)

}
