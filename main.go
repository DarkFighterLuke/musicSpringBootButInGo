package main

import (
	"log"
	"musicSpringBootButInGo/routers"
	"net/http"
)

func main() {
	r := routers.GetInitializedRouter()
	err := http.ListenAndServe("0.0.0.0:8090", r)
	if err != nil {
		log.Fatalln(err)
	}
}
