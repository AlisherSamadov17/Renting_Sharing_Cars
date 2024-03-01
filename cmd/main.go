package main

import (
	"log"
	"main/config"
	"main/controller"
	"main/store/postgres"
	_ "github.com/lib/pq"
)

func main()  {
	cfg := config.Load()
	store, err := postgres.New(cfg)
	if err != nil {
		log.Fatalln("error while connecting to db",err.Error())
	return
	}
	defer store.Db.Close()

	con := controller.New(store)
	// con.CreateCar()
	// con.DeleteFromCarsOne()
	// con.UpdateCarOne()
	con.GetBYIdCar()
}








