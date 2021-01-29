package main

import (
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/mikelpsv/ttu-service/app"
	"github.com/mikelpsv/ttu-service/routes"
	"log"
	"net/http"
)

func main() {
	app.AppVersion = "0.1"
	app.ApiVersion = "v1.0"

	err := app.Cfg.ReadEnv()
	if err != nil {
		log.Printf("error reading env, %s", err)
	}

	err = app.InitDb()
	if err != nil {
		log.Printf("error init database, %s", err)
	}

	routeItems := new(routes.Routes)
	routeItems = routes.RegisterHandlers(routeItems)
	router := NewRouter(routeItems)

	err = http.ListenAndServe(app.Cfg.AppAddr+":"+app.Cfg.AppPort, router)
	if err != nil {
		log.Println(err)
	}

}

func NewRouter(routeItems *routes.Routes) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range *routeItems {
		handlerFunc := route.HandlerFunc
		//if route.ValidateToken {
		//	handlerFunc = SetMiddlewareAuth(handlerFunc)
		//}

		if route.SetHeaderJSON {
			handlerFunc = app.SetMiddlewareJSON(handlerFunc)
		}

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			HandlerFunc(handlerFunc)
	}

	return router
}
