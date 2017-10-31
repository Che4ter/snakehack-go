package main

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"github.com/stair-ch/snakehack-go/api"
	"github.com/stair-ch/snakehack-go/configuration"
	"github.com/stair-ch/snakehack-go/handler"
	"github.com/stair-ch/snakehack-go/helper"
	"github.com/stair-ch/snakehack-go/logic"
	"net/http"
	"os"
)

var config configuration.Configuration

func main() {
	fmt.Printf("\n\n")
	fmt.Println("{OO}")
	fmt.Println("\\__/              snakehack 2017                  ")
	fmt.Println(" |^|              go starter kit                 /\\")
	fmt.Println(" | |____________________________________________/ / ")
	fmt.Println(" \\_______________________________________________/ ")
	fmt.Println("                                       STAIR - v1.0")
	fmt.Printf("\n\n")

	config, _ = configuration.ParseConfiguration()
	gameStartBridge := make(chan api.StartRequest)
	gameRequestBridge := make(chan api.MoveRequest)
	gameResponseBridge := make(chan api.MoveResponse)

	// Instantiate a new router
	r := httprouter.New()

	r.POST("/start", handler.Start(gameStartBridge, config))
	r.POST("/move", handler.Move(gameRequestBridge, gameResponseBridge))

	r.ServeFiles("/static/*filepath", http.Dir("static"))
	r.GET("/", handler.Index)

	go logic.RunGameLogic(config, gameRequestBridge, gameResponseBridge, gameStartBridge)

	port := os.Getenv("PORT")

	if port == "" {
		port = config.Port
	}

	fmt.Printf("snake running on http://%v:%v\n", helper.GetOutboundIP(), port)
	// Fire up the server
	http.ListenAndServe(":"+port, r)
}
