package routes

import (
	"fmt"
	"github.com/mikelpsv/ttu-service/app"
	"github.com/mikelpsv/ttu-service/models"
	"net/http"
)

/*
Создать трек: POST /track
Создать трек и присвоить свой номер: POST /track/свой номер
Получить трек: GET /users/1
Удалить трек: DELETE /users/1
Получить все треки: GET /users
*/

func RegisterTrackHandlers(routeItems *Routes) *Routes {
	*routeItems = append(*routeItems, Route{
		Name:          "CreateTrack",
		Method:        "POST",
		Pattern:       fmt.Sprintf("/%s/{namespace}/track", app.ApiVersion),
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   handlerCreateTrack,
	})
	*routeItems = append(*routeItems, Route{
		Name:          "CreateTrackWithId",
		Method:        "POST",
		Pattern:       fmt.Sprintf("/%s/{namespace}/track/{trackId}", app.ApiVersion),
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   handlerCreateTrackWithId,
	})
	*routeItems = append(*routeItems, Route{
		Name:          "GetTrack",
		Method:        "GET",
		Pattern:       fmt.Sprintf("/%s/{namespace}/track/{trackId}", app.ApiVersion),
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   handlerGetTrack,
	})
	*routeItems = append(*routeItems, Route{
		Name:          "DeleteTrack",
		Method:        "GET",
		Pattern:       fmt.Sprintf("/%s/{namespace}/track/{trackId}", app.ApiVersion),
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   handlerDeleteTrack,
	})
	*routeItems = append(*routeItems, Route{
		Name:          "UpdateTrack",
		Method:        "PUT",
		Pattern:       fmt.Sprintf("/%s/{namespace}/track/{trackId}", app.ApiVersion),
		SetHeaderJSON: true,
		ValidateToken: false,
		HandlerFunc:   handlerUpdateTrack,
	})
	return routeItems

}
func handlerCreateTrack(w http.ResponseWriter, r *http.Request) {
	var namespace, trackUid string
	trServ := models.TrackService{app.Db}
	track, err := trServ.CreateTrack(namespace, trackUid)
	if err != nil{
		//
	}
	app.ResponseJSON(w, http.StatusOK, track)
}

func handlerCreateTrackWithId(w http.ResponseWriter, r *http.Request) {
	var namespace, trackUid string
	trServ := models.TrackService{app.Db}
	track, err := trServ.CreateTrack(namespace, trackUid)
	if err != nil{
		//
	}
	app.ResponseJSON(w, http.StatusOK, track)
}

func handlerGetTrack(w http.ResponseWriter, r *http.Request) {
	app.ResponseJSON(w, http.StatusOK, "self id")
}
func handlerDeleteTrack(w http.ResponseWriter, r *http.Request) {
	app.ResponseJSON(w, http.StatusOK, "self id")
}

func handlerUpdateTrack(w http.ResponseWriter, r *http.Request) {
	app.ResponseJSON(w, http.StatusOK, "self id")
}
