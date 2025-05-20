package models

type AppConfig struct {
	Server Server `json:"server"`
}

type Server struct {
	Host string `json:"host"`
	Port string `json:"port"`
}
