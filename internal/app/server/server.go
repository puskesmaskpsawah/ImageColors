package server

import "net/http"

func SetupServer(serverAddress string) http.Server {
	return http.Server{
		Addr: serverAddress,
	}
}
