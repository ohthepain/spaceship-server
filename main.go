package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"moul.io/banner"
)

// type State struct {
// 	Rooms map[string]Room
// }

// func NewState() *State {
// 	fmt.Printf("NewState\n")
// 	var s State
// 	s.Rooms = make(map[string]Room)
// 	s.Rooms["main"] = *NewRoom()
// 	return &s
// }

func CORS(c *gin.Context) {

	// TODO: CORS is wide open!
	// First, we add the headers with need to enable CORS
	// Make sure to adjust these headers to your needs
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	// Second, we handle the OPTIONS problem
	if c.Request.Method != "OPTIONS" {

		c.Next()

	} else {

		// Everytime we receive an OPTIONS request,
		// we just return an HTTP 200 Status Code
		// Like this, Angular can now do the real
		// request using any other method than OPTIONS
		c.AbortWithStatus(http.StatusOK)
	}
}

var ctx gin.Context

// var state State

func main() {
	fmt.Printf("%s\n\t\t\t... welcomes you!\n\n", banner.Inline("spaceship server"))

	// Rooms := Room{}
	// state := NewState()
	// fmt.Printf("room name %s\n", state.Rooms["main"].Name)
	// state.rooms["main"] = *NewRoom()

	router := gin.Default()
	router.Use(CORS)
	router.POST("/api/player/update", updatePlayer)
	router.GET("/api/rooms", getRooms)
	router.POST("/api/room/join", enterRoom)

	router.Run("localhost:8080")
}

func getRooms(c *gin.Context) {
	// name := c.DefaultQuery("param_1", "no param found")
	fmt.Printf("getRooms: %d\n", len(state.Rooms))
	// fmt.Printf("getRooms: room name %s\n", state.Rooms["main"].Name)

	c.JSON(http.StatusOK, state.Rooms)
}

type updatePlayerRequest struct {
	player   Player
	missiles []Missile
}

func enterRoom(c *gin.Context) {
	player := new(Player)

	if err := c.Bind(&player); err != nil {
		fmt.Printf("enterRoom: Could not bind player json")
		c.Error(err)
		c.Abort()
	} else {
		fmt.Printf("enterRoom: player name: %s score: %d\n", player.Name, player.Score)
		c.JSON(http.StatusOK, c.Request.Body)
	}
}

func updatePlayer(c *gin.Context) {
	player := new(Player)

	if err := c.Bind(&player); err != nil {
		fmt.Printf("updatePlayer: Could not bind player json")
		c.Error(err)
		c.Abort()
	} else {
		fmt.Printf("updatePlayer: player name: %s score: %d\n", player.Name, player.Score)
		fmt.Printf("updatePlayer: spaceship at %f,%f\n", player.Spaceship.XPos, player.Spaceship.YPos)
		c.JSON(http.StatusOK, c.Request.Body)
	}
}
