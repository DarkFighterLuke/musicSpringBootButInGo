package routers

import (
	"musicSpringBootButInGo/controllers"

	"github.com/gorilla/mux"
)

func GetInitializedRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/music-spring-boot/album/add", controllers.AddAlbum).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/music-spring-boot/album/list", controllers.GetAllAlbums).Methods("GET")
	r.HandleFunc("/music-spring-boot/album/get/{id}", controllers.GetAlbumById).Methods("GET")
	r.HandleFunc("/music-spring-boot/album/set/{id}", controllers.UpdateAlbum).Methods("PUT").Headers("Content-Type", "application/json")
	r.HandleFunc("/music-spring-boot/album/remove/{id}", controllers.DeleteAlbum).Methods("DELETE")

	r.HandleFunc("/music-spring-boot/artista/add", controllers.AddArtista).Methods("POST").Headers("Content-Type", "application/json")
	r.HandleFunc("/music-spring-boot/artista/list", controllers.GetAllArtisti).Methods("GET")
	r.HandleFunc("/music-spring-boot/artista/get/{id}", controllers.GetArtistaById).Methods("GET")
	r.HandleFunc("/music-spring-boot/artista/set/{id}", controllers.UpdateArtista).Methods("PUT").Headers("Content-Type", "application/json")
	r.HandleFunc("/music-spring-boot/artista/remove/{id}", controllers.DeleteArtista).Methods("DELETE")

	return r
}
