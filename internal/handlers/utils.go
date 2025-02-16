package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
)

func RespondWithError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
	slog.Error(message)
}

func RespondWithMessage(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	w.Write([]byte(message))
}

func RespondWithJSON(w http.ResponseWriter, code int, data []byte) {
	w.WriteHeader(code)
	w.Write(data)
}

func FormatTemperature(temp int) string {
	return strconv.Itoa(temp) + "°C"
}

func FormatPercentage(value int) string {
	return strconv.Itoa(value) + "%"
}

func FormatWind(speed int, direction string) string {
	return strconv.Itoa(speed) + " km/h " + direction
}
