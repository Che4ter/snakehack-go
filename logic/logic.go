package logic

import (
	"fmt"
	"github.com/stair-ch/snakehack-go/api"
	"github.com/stair-ch/snakehack-go/configuration"
)

type GameData struct {
	config    configuration.Configuration
	gameID    string
	width     int
	height    int
	fieldData []api.MoveRequest
}

var gameData GameData
var count int

func RunGameLogic(config configuration.Configuration, gameRequestBridge chan api.MoveRequest, gameResponseBridge chan api.MoveResponse, gameStartBridge chan api.StartRequest) {
	gameData.config = config

	for {
		select {
		case start := <-gameStartBridge:
			gameData.gameID = start.GameId
			gameData.width = start.Width
			gameData.height = start.Height
			gameData.fieldData = []api.MoveRequest{}
			count = 0

			fmt.Printf("\n\n")
			fmt.Println("  {0O}")
			fmt.Println("  \\__/   new game started")
			fmt.Printf("  /^/    height %d, width %d\n", gameData.height, gameData.width)
			fmt.Println(" ( ( 		")
			fmt.Println("  \\_\\____")
			fmt.Println(" (________)")
			fmt.Println("(_________()Oo")
			fmt.Printf("\n\n")

		case newField := <-gameRequestBridge:
			gameData.fieldData = append(gameData.fieldData, newField)

			var direction configuration.MoveDirection

			if count < 3 {
				direction = configuration.MOVE_UP
			} else if count >= 3 && count < 6 {
				direction = configuration.MOVE_RIGHT
			} else if count >= 6 && count < 9 {
				direction = configuration.MOVE_DOWN
			} else if count >= 9 && count < 11 {
				direction = configuration.MOVE_LEFT
			} else if count == 11 {
				direction = configuration.MOVE_LEFT
				count = -1
			}
			count++

			m := api.MoveResponse{
				Move:  direction.String(),
				Taunt: "Lalala",
			}
			gameResponseBridge <- m
		}
	}
}
