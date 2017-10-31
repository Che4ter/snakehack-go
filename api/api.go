package api

import (
	"encoding/json"
	"errors"
)

type StartRequest struct {
	GameId string `json:"game_id"`
	Height int    `json:"height"`
	Width  int    `json:"width"`
}

type StartResponse struct {
	Color          string `json:"color"`
	Name           string `json:"name"`
	HeadUrl        string `json:"head_url,omitempty"`
	Taunt          string `json:"taunt,omitempty"`
	HeadType       string `json:"head_type,omitempty"`
	TailType       string `json:"tail_type,omitempty"`
	SecondaryColor string `json:"secondary_color,omitempty"`
}

type MoveRequest struct {
	Food       []Point `json:"food"`
	GameId     string  `json:"game_id"`
	Height     int     `json:"height"`
	Snakes     []Snake `json:"snakes"`
	DeadSnakes []Snake `json:"dead_snake"`
	Turn       int     `json:"turn"`
	Width      int     `json:"width"`
	You        string  `json:"you"`
}

type MoveResponse struct {
	Move  string `json:"move"`
	Taunt string `json:"taunt,omitempty"`
}

type Point struct {
	X int `json:"x"`
	Y int `json:"y"`
}

type Snake struct {
	Coords       []Point `json:"coords"`
	HealthPoints int     `json:"health_points"`
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Taunt        string  `json:"taunt"`
}

// Decode [number, number] JSON array into a Point
func (point *Point) UnmarshalJSON(data []byte) error {
	var coords []int
	json.Unmarshal(data, &coords)
	if len(coords) != 2 {
		return errors.New("Bad set of coordinates: " + string(data))
	}
	*point = Point{X: coords[0], Y: coords[1]}
	return nil
}
