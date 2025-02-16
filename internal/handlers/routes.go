package handlers

import (
	"assignment2/internal/models"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"log/slog"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var (
	books      []*models.Book
	bookID     = 3
	weatherAPI string
)

const (
	homeURL = "http://api.weatherstack.com/current?access_key="
)

func Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /books", getBooksHandler)
	mux.HandleFunc("POST /books", postBooksHandler)
	mux.HandleFunc("PUT /book/{id}", updateBookHandler)
	mux.HandleFunc("DELETE /book/{id}", deleteBookHandler)
	mux.HandleFunc("GET /weather/{city}", weatherHandler)

	return mux
}

type APIServer struct {
	address string
	mux     *http.ServeMux
}

func NewAPIServer(address string) *APIServer {
	return &APIServer{
		address: address,
		mux:     Routes(),
	}
}

func (s *APIServer) Run() {
	slog.Info("API server listening on " + s.address)

	books = append(books, []*models.Book{
		{ID: 0, Title: "Grokking Algorithms", Author: "Aditya Bhargava", Year: 2016, Genre: "Science"},
		{ID: 1, Title: "No Longer Human", Author: "Osamu Dazai", Year: 1948, Genre: "Novel"},
		{ID: 2, Title: "Learning Go", Author: "Jon Bodner", Year: 2021, Genre: "Science"},
	}...)

	weatherAPI = os.Getenv("WEATHER_API")

	log.Fatal(http.ListenAndServe(s.address, s.mux))
}

func weatherHandler(w http.ResponseWriter, r *http.Request) {
	city := r.PathValue("city")

	// slog.Debug(city)

	if len(city) == 0 {
		RespondWithError(w, http.StatusBadRequest, "City name is not correct")
		return
	}

	apiURL := fmt.Sprintf("%s%s&query=%s", homeURL, weatherAPI, city)

	// slog.Debug(apiURL)

	response, err := http.Get(apiURL)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to fetch data")
		return
	}
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Failed to read response")
		return
	}

	fmt.Println("Raw API Response:", string(data))

	var apiResponse models.WeatherResponse
	if err := json.Unmarshal(data, &apiResponse); err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Error parsing response: "+err.Error())
		return
	}

	if apiResponse.Location.Name == "" {
		RespondWithError(w, http.StatusNotFound, "City not found or invalid request")
		return
	}

	condition := "Unknown"
	if len(apiResponse.Current.WeatherDescriptions) > 0 {
		condition = apiResponse.Current.WeatherDescriptions[0]
	}

	userResponse := models.UserFriendlyWeather{
		City:        apiResponse.Location.Name,
		Temperature: FormatTemperature(apiResponse.Current.Temperature),
		Condition:   condition,
		FeelsLike:   FormatTemperature(apiResponse.Current.FeelsLike),
		Humidity:    FormatPercentage(apiResponse.Current.Humidity),
		Wind:        FormatWind(apiResponse.Current.WindSpeed, apiResponse.Current.WindDirection),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userResponse)
}

func getBooksHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.MarshalIndent(books, "", "  ")
	if err != nil {
		RespondWithError(w, http.StatusInternalServerError, "Something went wrong")
		return
	}
	RespondWithJSON(w, http.StatusOK, resp)
}

func postBooksHandler(w http.ResponseWriter, r *http.Request) {
	var book models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid JSON format")
		return
	}

	if err := book.Validate(); err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	book.ID = bookID
	bookID++
	books = append(books, &book)

	RespondWithMessage(w, http.StatusOK, "Book was successfully added")
}

func updateBookHandler(w http.ResponseWriter, r *http.Request) {
	respId := r.PathValue("id")

	id, err := strconv.Atoi(respId)
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	for i, book := range books {
		if book.ID == id {
			var respBook = &models.Book{}

			if err := json.NewDecoder(r.Body).Decode(respBook); err != nil {
				RespondWithError(w, http.StatusBadRequest, "Invalid JSON format")
				return
			}

			if err := respBook.Validate(); err != nil {
				RespondWithError(w, http.StatusBadRequest, err.Error())
				return
			}

			respBook.ID = id
			books[i] = respBook

			RespondWithMessage(w, http.StatusOK, "Book successfully updated")
			return
		}
	}

	RespondWithError(w, http.StatusNotFound, "No book found with given ID")
}

func deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractIDFromURL(r.URL.Path, "/book/")
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...)
			RespondWithMessage(w, http.StatusOK, "Book successfully deleted")
			return
		}
	}

	RespondWithError(w, http.StatusNotFound, "No book found with given ID")
}

func extractIDFromURL(path, prefix string) (int, error) {
	if !strings.HasPrefix(path, prefix) {
		return 0, errors.New("invalid URL format")
	}

	idStr := strings.TrimPrefix(path, prefix)
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return 0, errors.New("invalid ID")
	}
	return id, nil
}
