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

func AddArtista(w http.ResponseWriter, r *http.Request) {
	var artista models.Artista
	err := json.NewDecoder(r.Body).Decode(&artista)
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

	artistaRepository := repositories.NewArtistaRepository(db)
	lastId, err := artistaRepository.GetLastId()
	artista.ArtistaId = lastId + 1

	id, err := artistaRepository.Insert(&artista)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Write([]byte(fmt.Sprintf("Inserito artista con id %d", id)))
}

func GetAllArtisti(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	artistaRepository := repositories.NewArtistaRepository(db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	artisti, err := artistaRepository.FindAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	artistiJson, err := json.Marshal(&artisti)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(artistiJson)
}

func GetArtistaById(w http.ResponseWriter, r *http.Request) {
	artistaId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
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

	artistaRepository := repositories.NewArtistaRepository(db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	artista, err := artistaRepository.FindById(uint(artistaId))
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	artistaJson, err := json.Marshal(artista)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(artistaJson)
}

func UpdateArtista(w http.ResponseWriter, r *http.Request) {
	var artista models.Artista
	err := json.NewDecoder(r.Body).Decode(&artista)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	artistaId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	artista.ArtistaId = uint(artistaId)
	fmt.Printf("%#v", artista)

	db, err := utils.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	artistaRepository := repositories.NewArtistaRepository(db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	err = artistaRepository.Update(&artista)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Write([]byte(fmt.Sprintf("Aggiornato artista con id %d", artista.ArtistaId)))
}

func DeleteArtista(w http.ResponseWriter, r *http.Request) {
	var artista models.Artista
	artistaId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	artista.ArtistaId = uint(artistaId)

	db, err := utils.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	artistaRepository := repositories.NewArtistaRepository(db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	err = artistaRepository.Delete(&artista)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Write([]byte(fmt.Sprintf("Eliminato artista con id %d", artista.ArtistaId)))

}
