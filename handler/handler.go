package handler

import (
	"encoding/json"
	"fmt"
	"github.com/stair-ch/snakehack-go/api"
	"github.com/stair-ch/snakehack-go/configuration"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n ")
}

func Start(gameStartBridge chan api.StartRequest, config configuration.Configuration) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		startData := api.StartRequest{}

		json.NewDecoder(r.Body).Decode(&startData)

		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		gameStartBridge <- startData

		s := api.StartResponse{
			Color:          config.Color,
			Name:           config.Name,
			Taunt:          config.Taunt,
			HeadUrl:        fmt.Sprintf("%v://%v/static/head.png", scheme, r.Host),
			HeadType:       config.HeadType.String(),
			TailType:       config.TailType.String(),
			SecondaryColor: config.SecondaryColor,
		}

		sj, _ := json.Marshal(s)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprintf(w, "%s", sj)
	}
}

func Move(gameRequestBridge chan api.MoveRequest, gameResponseBridge chan api.MoveResponse) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		gameData := api.MoveRequest{}

		err := json.NewDecoder(r.Body).Decode(&gameData)
		if err == nil {
			gameRequestBridge <- gameData
			// Marshal provided interface into JSON structure
			mj, _ := json.Marshal(<-gameResponseBridge)

			// Write content-type, statuscode, payload
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(200)
			fmt.Fprintf(w, "%s", mj)

		} else {
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(500)
		}
	}
}
