package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Album struct {
	ID     uint    `json:"id" gorm:"primarykey;autoIncrement"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Year   uint    `json:"year"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "The Number of the Beast", Artist: "Iron Maiden", Year: 1982, Price: 25.19},
	{ID: "2", Title: "Youthanasia", Artist: "Medadeth", Year: 1994, Price: 13.65},
	{ID: "3", Title: "Master of Puppets", Artist: "Metallica", Year: 1986, Price: 20.97},
}

func getAlbums(ctx *gin.Context) {
	var albums []album
	db.Find
	ctx.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(ctx *gin.Context) {
	var newAlbum Album

	if err := ctx.BindJSON(&newAlbum); err != nil {
		return
	}
	db.Create(&newAlbum)
	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	var album Album
	result := db.First(&album, "id= ?", id)
	if result.Error != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
	}
	ctx.IndentedJSON(http.StatusOK, album)
}

func putAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var modifyAlbum Album //datos ingresados por el usuario en el body del PUT
	var album Album       // datos obtenidos de la base de datos

	if err := ctx.BindJSON(&modifyAlbum); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"messge": "Datos incorrectos"})
		return
	}

	album.Title = modifyAlbum.Title
	album.Artist = modifyAlbum.Artist
	album.Year = modifyAlbum.Year
	album.Price = modifyAlbum.Price

	db.Save(&album)
	ctx.IndentedJSON(http.StatusOK, album)
	/*
		for index, a := range albums {
			if a.ID == id {
				albums[index] = modifyAlbum
				albums[index].ID = id
				ctx.IndentedJSON(http.StatusOK, albums[index])
				return
			}
		} */
	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "album no encontrado"})
}

func deleteAlbumByID(ctx *gin.Context) {
	id := ctx.Param("id")

	result := db.Delete(&Album{}, "id =?", id)
	if result.RowsAffected == 0 {
		ctx.IndentedJSON(http.StatusNotFound)
	}
	ctx.IndentedJSON(http.StatusNotFound)
}

var db *gorm.DB //dejar disponible de manera global

func initDB() (*gorm.DB, err) {
	//usuario:password@tcp8ruta:puerto)/baseDeDatos
	dsn := "root:romanovictoria_96@tcp(localhost:3306)/db_vinilos"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	log.Println("Conexion a la BD exitosa")
	return db, nil
}

func main() {
	var err error
	db, err = initDB()

	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.GET("/albums/:id", getAlbumByID)
	router.PUT("/albums/:id", putAlbumByID)
	router.DELETE("/albums/:id", deleteAlbumByID)
	router.Run("localhost:8080")
}
