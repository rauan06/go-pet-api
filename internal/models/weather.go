package models

type (
	WeatherResponse struct {
		Location struct {
			Name    string `json:"name"`
			Country string `json:"country"`
		} `json:"location"`
		Current struct {
			Temperature         int      `json:"temperature"`
			WeatherDescriptions []string `json:"weather_descriptions"`
			FeelsLike           int      `json:"feelslike"`
			Humidity            int      `json:"humidity"`
			WindSpeed           int      `json:"wind_speed"`
			WindDirection       string   `json:"wind_dir"`
			WeatherIcons        []string `json:"weather_icons"`
		} `json:"current"`
	}

	UserFriendlyWeather struct {
		City        string `json:"city"`
		Temperature string `json:"temperature"`
		Condition   string `json:"condition"`
		FeelsLike   string `json:"feels_like"`
		Humidity    string `json:"humidity"`
		Wind        string `json:"wind"`
	}
)
