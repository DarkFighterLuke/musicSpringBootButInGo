package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"musicSpringBootButInGo/models"
	"musicSpringBootButInGo/repositories"
	"musicSpringBootButInGo/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func AddAlbum(w http.ResponseWriter, r *http.Request) {
	var album models.Album
	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	db, err := utils.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	albumRepository := repositories.NewAlbumRepository(db)
	lastId, err := albumRepository.GetLastId()
	album.AlbumId = lastId + 1

	id, err := albumRepository.Insert(&album)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Write([]byte(fmt.Sprintf("Inserito album con id %d", id)))
}

func GetAllAlbums(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	albumRepository := repositories.NewAlbumRepository(db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	album, err := albumRepository.FindAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	albumJson, err := json.Marshal(&album)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(albumJson)
}

func GetAlbumById(w http.ResponseWriter, r *http.Request) {
	albumId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	db, err := utils.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	albumRepository := repositories.NewAlbumRepository(db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	album, err := albumRepository.FindById(uint(albumId))
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	albumJson, err := json.Marshal(album)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(albumJson)
}

func UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	var album models.Album
	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	albumId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	album.AlbumId = uint(albumId)

	db, err := utils.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	albumRepository := repositories.NewAlbumRepository(db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	err = albumRepository.Update(&album)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Write([]byte(fmt.Sprintf("Aggiornato album con id %d", album.AlbumId)))
}

func DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	var album models.Album
	albumId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	album.AlbumId = uint(albumId)

	db, err := utils.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	albumRepository := repositories.NewAlbumRepository(db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	err = albumRepository.Delete(&album)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Write([]byte(fmt.Sprintf("Eliminato album con id %d", album.AlbumId)))

}
