package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

type bestSong struct {
	ID 	string 	`json:"id"`
	Title string `json:"title"`
	Album string `json:"album"`
	Price float64 `json:"price"`
}

type album struct {
	ID string `json:"id"`
	Name string `json:"name"`
	Year int `json:"year"`
	Price float64 `json:"price"`
}

var bestSongs = []bestSong{  
	{ID: "1", Title: "Zankokuna Tenshino These", Album: "Best Anime Songs in 90s", Price: 100},
	{ID: "2", Title: "Sobakasu", Album: "Best Anime Songs in 2000s", Price: 150},
	{ID: "3", Title: "Mezase Pokemon Master", Album: "Best Anime Songs in 200s", Price: 160},
	{ID: "4", Title: "Brave Heart", Album: "Best Digimon Songs", Price: 160},
	{ID: "5", Title: "Brave Heart", Album: "Best Digimon Songs", Price: 160},
	{ID: "6", Title: "Change The World", Album: "Best Inuyasha Songs", Price: 160},
}

var albums = []album{
	{ID: "1", Name: "Best Anime Songs in 90s", Year: 2001, Price: 1000},
	{ID: "2", Name: "Best Anime Songs in 2000s", Year: 2011, Price: 1300},
	{ID: "3", Name: "Best Digimon Songs", Year: 2011, Price: 1300},
	{ID: "4", Name: "Best Inuyasha Songs", Year: 2011, Price: 1200},
}

func getBestSongs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, bestSongs)
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func main() {
	router := gin.Default()
	router.GET("/best-songs", getBestSongs)
	router.GET("/best-songs/:id", getBestSongById)
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/new-album", postAlbum)
	router.POST("/best-song", postBestSong)
	router.Run("localhost:9000")
}

func postAlbum(c *gin.Context) {
	var newAlbum album

	if err:= c.BindJSON(&newAlbum); err != nil {
		return
	}

	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated, newAlbum)
}

func postBestSong(c *gin.Context) {
	var newBestSong bestSong

	if err:=c.BindJSON(&newBestSong); err != nil {
		return
	}

	bestSongs = append(bestSongs, newBestSong)
	c.IndentedJSON(http.StatusCreated, newBestSong)
}

func getBestSongById(c *gin.Context) {
	id := c.Param("id")

	for _, song := range bestSongs {
		if song.ID == id {
			 c.IndentedJSON(http.StatusOK, song)
			 return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Song not found"})
}

func getAlbumById(c *gin.Context) {
	id := c.Param("id")

	for _, albumReq := range albums {
		if albumReq.ID == id {
			c.IndentedJSON(http.StatusOK, albumReq)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "album not found, try with different id"})
}