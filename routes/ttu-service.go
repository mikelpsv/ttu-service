package routes

import (
	"fmt"
	"github.com/mikelpsv/ttu-service/app"
	"net/http"
)

func RegisterTTuHandlers(routeItems *Routes) *Routes {
	*routeItems = append(*routeItems, Route{
		Name:          "GetUnitId",
		Method:        "GET",
		Pattern:       fmt.Sprintf("/%s/unit/id", app.ApiVersion),
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   handlerGetUnitId,
	})
	return routeItems
}
func handlerGetUnitId(w http.ResponseWriter, r *http.Request) {
	app.ResponseJSON(w, http.StatusOK, "test")
}
