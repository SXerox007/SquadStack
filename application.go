package main

import (
	"SquadStack/base/environment"
	"SquadStack/base/router"
	"SquadStack/base/router/server"
	"log"
)

// init
func Init() {
	env := environment.GetEnv()
	port := environment.GetPort()
	router.InitRouter()
	setupRouter(env, port)
}

func main() {
	Init()
}

func setupRouter(env, port string) {
	squadstack := router.SubRouter("/squadstack")
	squadstack.HandleFunc("/{version}/file/parking-lot", UploadParkingLotFile()).Methods("POST")

	log.Println("Server serve at", env+":"+port)
	server.StartServer(port)
}
