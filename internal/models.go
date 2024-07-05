package internal

type Response struct {
	ClientIP string `json:"client_ip"`
	Location string `json:"location"`
	Greeting string `json:"greeting"`
}

type IPinfoResponse struct {
	IP       string `json:"ip"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Location string `json:"loc"`
}

type LocationData struct {
	City     string `json:"city"`
	Country  string `json:"country"`
	IP       string `json:"ip"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Readme   string `json:"readme"`
	Region   string `json:"region"`
	Timezone string `json:"timezone"`
}

type WeatherInfo struct {
	Current  CurrentWeather  `json:"current"`
	Location LocationDetails `json:"location"`
}

type CurrentWeather struct {
	Cloud            int       `json:"cloud"`
	Condition        Condition `json:"condition"`
	DewpointC        float64   `json:"dewpoint_c"`
	DewpointF        float64   `json:"dewpoint_f"`
	FeelslikeC       float64   `json:"feelslike_c"`
	FeelslikeF       float64   `json:"feelslike_f"`
	GustKph          float64   `json:"gust_kph"`
	GustMph          float64   `json:"gust_mph"`
	HeatindexC       float64   `json:"heatindex_c"`
	HeatindexF       float64   `json:"heatindex_f"`
	Humidity         int       `json:"humidity"`
	IsDay            int       `json:"is_day"`
	LastUpdated      string    `json:"last_updated"`
	LastUpdatedEpoch int64     `json:"last_updated_epoch"`
	PrecipIn         float64   `json:"precip_in"`
	PrecipMm         float64   `json:"precip_mm"`
	PressureIn       float64   `json:"pressure_in"`
	PressureMb       float64   `json:"pressure_mb"`
	TempC            float64   `json:"temp_c"`
	TempF            float64   `json:"temp_f"`
	Uv               float64   `json:"uv"`
	VisKm            float64   `json:"vis_km"`
	VisMiles         float64   `json:"vis_miles"`
	WindDegree       int       `json:"wind_degree"`
	WindDir          string    `json:"wind_dir"`
	WindKph          float64   `json:"wind_kph"`
	WindMph          float64   `json:"wind_mph"`
	WindchillC       float64   `json:"windchill_c"`
	WindchillF       float64   `json:"windchill_f"`
}

type Condition struct {
	Code int    `json:"code"`
	Icon string `json:"icon"`
	Text string `json:"text"`
}

type LocationDetails struct {
	Country        string  `json:"country"`
	Lat            float64 `json:"lat"`
	Localtime      string  `json:"localtime"`
	LocaltimeEpoch int64   `json:"localtime_epoch"`
	Lon            float64 `json:"lon"`
	Name           string  `json:"name"`
	Region         string  `json:"region"`
	TzId           string  `json:"tz_id"`
}
