package internal

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func SayHello(ctx *gin.Context) {
	ip := ctx.GetHeader("X-FORWARDED-FOR")
	log.Printf("ip address is: %s", ip)
	ip = strings.Split(ip, ",")[0]
	log.Printf("remote address: %s", ctx.Request.RemoteAddr)

	visitorName := ctx.Query("visitor_name")

	if visitorName == "" {
		visitorName = "user"
	}

	token := os.Getenv("PINFO_ACCESS_TOKEN")
	weatherApiKey := os.Getenv("WEATHER_API_KEY")

	info, err := GetGeolocation(ip, token)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	weatherData, err := GetWeatherData(info.Loc, weatherApiKey)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	greeting := GenerateGreeting(visitorName, weatherData.Location.Name, float32(weatherData.Current.TempC))

	ctx.JSON(http.StatusOK, Response{
		ClientIP: info.IP,
		Location: weatherData.Location.Name,
		Greeting: greeting,
	})
}
