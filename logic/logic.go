package logic

import (
	"fmt"
	"github.com/Che4ter/snakehack-go/api"
	"github.com/Che4ter/snakehack-go/configuration"
	"math"
	"math/rand"
)

type GameData struct {
	config    configuration.Configuration
	gameID    string
	width     int
	height    int
	fieldData []api.MoveRequest
}

type CurrentData struct {
	headCoords api.Point
	shortestFoodCoords api.Point
	nexMove configuration.MoveDirection
	nextHeadCoords api.Point
}

var gameData GameData
var currentData CurrentData
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
			getCurrentGameData(newField)
			getShortestFood()

			fmt.Println(currentData.shortestFoodCoords)
			goToFood()

			for !checkMove(currentData.nexMove) {
				randNumber := rand.Intn(4)
				switch randNumber {
				case 0: currentData.nexMove = configuration.MOVE_DOWN
				case 1: currentData.nexMove = configuration.MOVE_UP
				case 2: currentData.nexMove = configuration.MOVE_RIGHT
				case 3: currentData.nexMove = configuration.MOVE_LEFT
				default:
					currentData.nexMove = configuration.MOVE_LEFT
				}

			}

			m := api.MoveResponse{
				Move: currentData.nexMove.String() ,
				Taunt: "go go go",
			}
			gameResponseBridge <- m
		}
	}
}

func getCurrentGameData(newField api.MoveRequest){
	gameData.fieldData = append(gameData.fieldData, newField)

	for _,snake := range newField.Snakes {
		if snake.Id == newField.You {
			currentData.headCoords = snake.Coords[0]
			break
		}
	}
}
func getShortestFood(){
	currentFoods := gameData.fieldData[len(gameData.fieldData)-1].Food

	var shortestDistance float64
	var shortestIndex int = 0
	for index, food := range currentFoods {
		a := math.Abs(float64(food.X - currentData.headCoords.X))
		b := math.Abs(float64(food.Y - currentData.headCoords.Y))

		c := math.Hypot(float64(a),float64(b))

		if shortestDistance == 0 {
			shortestDistance = c
			shortestIndex = index
		}else if c < shortestDistance {
			shortestDistance = c
			shortestIndex = index
		}
	}

	currentData.shortestFoodCoords = currentFoods[shortestIndex]
}
func goToFood(){
	h := math.Abs(float64(currentData.shortestFoodCoords.X - currentData.headCoords.X))
	v := math.Abs(float64(currentData.shortestFoodCoords.Y - currentData.headCoords.Y))
	fmt.Println("horizontal ", h)
	fmt.Println("vertical", v)
	if h < v{

		fmt.Println("Food ")
		//Move v
		if currentData.shortestFoodCoords.Y > currentData.headCoords.Y {
			currentData.nexMove = configuration.MOVE_DOWN
		} else {
			currentData.nexMove = configuration.MOVE_UP
		}
	}else if h > v{
		//Move h
		if currentData.shortestFoodCoords.X > currentData.headCoords.X {
			currentData.nexMove = configuration.MOVE_RIGHT
		} else {
			currentData.nexMove = configuration.MOVE_LEFT
		}

	}else{
		if currentData.shortestFoodCoords.X > currentData.headCoords.X {
			currentData.nexMove = configuration.MOVE_RIGHT
		} else {
			currentData.nexMove = configuration.MOVE_LEFT
		}
	}
}

func checkMove(nextMove configuration.MoveDirection) bool{
	var nextCoords api.Point

	switch nextMove {
		case configuration.MOVE_LEFT:
			nextCoords.X = currentData.headCoords.X - 1
			nextCoords.Y = currentData.headCoords.Y
		case configuration.MOVE_RIGHT:
			nextCoords.X = currentData.headCoords.X + 1
			nextCoords.Y = currentData.headCoords.Y
		case configuration.MOVE_UP:
			nextCoords.X = currentData.headCoords.X
			nextCoords.Y = currentData.headCoords.Y - 1
		case configuration.MOVE_DOWN:
			nextCoords.X = currentData.headCoords.X
			nextCoords.Y = currentData.headCoords.Y	+ 1
	}
	fmt.Println("Current Coords", currentData.headCoords)

	fmt.Println("Next Coords", nextCoords)

	if nextCoords.X < 0 || nextCoords.Y < 0{
		return false
	}

	if nextCoords.X > gameData.width || nextCoords.Y > gameData.height{
		return false
	}

	for _, snake := range gameData.fieldData[len(gameData.fieldData)-1].Snakes {
		for _, point := range snake.Coords {
			if point.X == nextCoords.X && point.Y == nextCoords.Y {
				fmt.Println("Used coords", point)
				return false

			}
		}
	}

	return true
}