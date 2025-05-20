package handlers

import (
	"encoding/json"
	"imagecolors/internal/app/services"
	"imagecolors/internal/models"
	"io"
	"net/http"
	"strconv"
	"strings"
)

func ProcImage(res http.ResponseWriter, req *http.Request) {
	var input models.InputImage

	body, err := io.ReadAll(req.Body)

	if err != nil {

	}

	defer req.Body.Close()

	json.Unmarshal(body, &input)

	colors := make(map[string]string)

	for _, pixel := range input.Pixels {
		XYpoints := strings.Split(pixel, ",")
		X, _ := strconv.Atoi(XYpoints[0])
		Y, _ := strconv.Atoi(XYpoints[1])
		colors[pixel] = services.GetPixelColor(input.Base64Image, X, Y)
	}

	if len(colors) == 0 {
		res.Header().Add("Content-Type", "application/json")
		res.WriteHeader(500)
		res.Write([]byte("Server error!"))
		return
	}

	response := models.PixelesImageResponse{
		Colors: colors,
	}

	r, _ := json.Marshal(&response)
	res.Header().Add("Content-Type", "application/json")
	res.WriteHeader(200)
	res.Write(r)
}
