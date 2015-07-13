package main

import (
	"fmt"
	"log"

	"github.com/kkdai/goFbAlbum"
)

func main() {
	fbToken := "" //Get your token here. fbAlbim.GetMyAlbums()
	fbAlbum := goFbAlbum.NewFbAlbum(fbToken)

	var targetAlbumId string
	var albumPhotoCount int
	var err error

	//Get and display all my albums name
	myAlbums, err := fbAlbum.GetMyAlbums()

	if err != nil {
		log.Fatalln("error:", err)
	}

	for _, album := range myAlbums.Data {
		fmt.Println(album.Name)
		targetAlbumId = album.ID
		albumPhotoCount = album.Count
		break
	}

	fmt.Println("album ID:", targetAlbumId)

	//Get and display all photo in first album
	myPhotos, err := fbAlbum.GetPhotoByAlbum(targetAlbumId, albumPhotoCount)

	if err != nil {
		log.Fatalln("error:", err)
	}

	fmt.Println("Myphoto:", len(myPhotos.Data))
	for _, photo := range myPhotos.Data {
		fmt.Println(photo.ID)
	}
}
