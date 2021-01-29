package routes

/**
  Данный файл содержит изменяемую часть сервера Api
  Список методов и функции-обработчики
*/
import (
	"github.com/mikelpsv/ttu-service/app"
	"net/http"
)

type Route struct {
	Name          string
	Method        string
	Pattern       string
	SetHeaderJSON bool
	ValidateToken bool
	HandlerFunc   http.HandlerFunc
}

type Routes []Route

type ResponsePing struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

type ResponseVersion struct {
	AppVersion string `json:"app_version"`
	ApiVersion string `json:"api_version"`
}

type ResponseHealth struct {
	Db bool `json:"db"`
}

func RegisterHandlers(routeItems *Routes) *Routes {
	routeItems = RegisterBaseHandlers(routeItems) // ping etc
	routeItems = RegisterTTuHandlers(routeItems)
	routeItems = RegisterTrackHandlers(routeItems)

	return routeItems
}

func RegisterBaseHandlers(routeItems *Routes) *Routes {
	*routeItems = append(*routeItems, Route{
		Name:          "Ping",
		Method:        "GET",
		Pattern:       "/ping",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   handlerPing,
	})
	*routeItems = append(*routeItems, Route{
		Name:          "Version",
		Method:        "GET",
		Pattern:       "/version",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   handlerVersion,
	})
	*routeItems = append(*routeItems, Route{
		Name:          "Health",
		Method:        "GET",
		Pattern:       "/health",
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   handlerHealth,
	})
	return routeItems
}

func handlerPing(w http.ResponseWriter, r *http.Request) {
	app.ResponseJSON(w, http.StatusOK, ResponsePing{
		Code:        http.StatusOK,
		Description: "",
	})
}

func handlerVersion(w http.ResponseWriter, r *http.Request) {
	app.ResponseJSON(w, http.StatusOK, ResponseVersion{
		AppVersion: app.AppVersion,
		ApiVersion: app.ApiVersion,
	})
}

func handlerHealth(w http.ResponseWriter, r *http.Request) {
	err := app.Db.Ping()
	dbOk := err == nil
	app.ResponseJSON(w, http.StatusOK, ResponseHealth{
		Db: dbOk,
	})
}
