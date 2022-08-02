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

func AddTraccia(w http.ResponseWriter, r *http.Request) {
	var traccia models.Traccia
	err := json.NewDecoder(r.Body).Decode(&traccia)
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

	tracciaRepository := repositories.NewTracciaRepository(db)
	lastId, err := tracciaRepository.GetLastId()
	traccia.TracceId = lastId + 1

	id, err := tracciaRepository.Insert(&traccia)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Write([]byte(fmt.Sprintf("Inserita traccia con id %d", id)))
}

func GetAllTracce(w http.ResponseWriter, r *http.Request) {
	db, err := utils.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	tracciaRepository := repositories.NewTracciaRepository(db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	tracce, err := tracciaRepository.FindAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	tracceJson, err := json.Marshal(&tracce)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(tracceJson)
}

func GetTracciaById(w http.ResponseWriter, r *http.Request) {
	tracciaId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
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

	tracciaRepository := repositories.NewTracciaRepository(db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	traccia, err := tracciaRepository.FindById(uint(tracciaId))
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	tracciaJson, err := json.Marshal(traccia)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(tracciaJson)
}

func UpdateTraccia(w http.ResponseWriter, r *http.Request) {
	var traccia models.Traccia
	err := json.NewDecoder(r.Body).Decode(&traccia)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	tracciaId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	traccia.TracceId = uint(tracciaId)
	fmt.Printf("%#v", traccia)

	db, err := utils.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	tracciaRepository := repositories.NewTracciaRepository(db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	err = tracciaRepository.Update(&traccia)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Write([]byte(fmt.Sprintf("Aggiornata traccia con id %d", traccia.TracceId)))
}

func DeleteTraccia(w http.ResponseWriter, r *http.Request) {
	var traccia models.Traccia
	tracciaId, err := strconv.ParseUint(mux.Vars(r)["id"], 10, 64)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}
	traccia.TracceId = uint(tracciaId)

	db, err := utils.GetDatabaseConnection()
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	tracciaRepository := repositories.NewTracciaRepository(db)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	err = tracciaRepository.Delete(&traccia)
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Write([]byte(fmt.Sprintf("Eliminata traccia con id %d", traccia.TracceId)))

}
