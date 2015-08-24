package http

import (
	"net/http"

	"github.com/niean/anteye/proc"
)

func configProcHttpRoutes() {
	// counter
	http.HandleFunc("/counter/all", func(w http.ResponseWriter, r *http.Request) {
		RenderDataJson(w, proc.GetAll())
	})
}
