package runner

import (
	"imagecolors/internal/app/handlers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/image/proc", handlers.ProcImage)
}
