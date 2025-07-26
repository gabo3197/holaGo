package handler

import (
	"fmt"
	"net/http"
)

func HolaHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "¡Hola mundo!")
}
