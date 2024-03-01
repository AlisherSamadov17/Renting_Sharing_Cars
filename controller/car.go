package controller

import (
	"fmt"
	"main/models"

	"github.com/google/uuid"
)

func (c Controller) CreateCar() {
car := getCarInfo_1()

id,err := c.Store.CarStorage.CreateCar(car)
if err != nil {
	fmt.Println("error while creating data",err.Error())
}
fmt.Println("id is :",id)
}

func (c Controller) DeleteFromCarsOne() {
	car := getCarInfo_2().Id
   if err := c.Store.CarStorage.DeleteCar(car);err != nil{
	panic(err)
   }
   fmt.Println("car deleted !")
}

func (c Controller) UpdateCarOne() {
	car := getCarInfo_3()

	err := c.Store.CarStorage.UpdateCar(car)
	if err != nil{
		fmt.Println("error while updating car",err)
	return
	}
	fmt.Println("Successfully updated !")
	return
}

func (c Controller) GetBYIdCar() {
	idStr :=""
	fmt.Print("enter the id to find the car :")
	fmt.Scan(&idStr)

    id,err := uuid.Parse(idStr)
	if err != nil {
		fmt.Println("id is not uuid :",err.Error())
		return
	}

	car,err := c.Store.CarStorage.GetByIdCar(id)
	if err != nil {
		fmt.Println("error while getting car by id :",err.Error())
		return
	}
  fmt.Println("searched car is :",car)
}



func getCarInfo_1() models.Car {
	var (
		name string
		brand string
		model string
		hourse_power int
		colour string
        engine_cap float32
	)

	fmt.Print("enter the name: ")
	fmt.Scan(&name)
	fmt.Print("enter the brand: ")
	fmt.Scan(&brand)
	fmt.Print("enter the model: ")
	fmt.Scan(&model)
	fmt.Print("enter the hourse_power: ")
	fmt.Scan(&hourse_power)
	fmt.Print("enter the color: ")
	fmt.Scan(&colour)
	fmt.Print("enter the engine capacity: ")
	fmt.Scan(&engine_cap)

	return models.Car{
		Name: name,
		Brand: brand,
		Model: model,
		HoursePower: hourse_power,
		Colour: colour,
		EngineCap: engine_cap,
	}
}

func getCarInfo_2() models.Car {
	var (
		id string
	)
	fmt.Print("insert 'ID' to delete :")
	fmt.Scan(&id)
	return models.Car{
	  Id: id,
	}
}

func getCarInfo_3() models.Car {
	var (
		id string
		name string
		brand string
		model string
		hourse_power int
		colour string
        engine_cap float32
	)
	fmt.Print("enter the id: ")
	fmt.Scan(&id)
	fmt.Print("enter the name: ")
	fmt.Scan(&name)
	fmt.Print("enter the brand: ")
	fmt.Scan(&brand)
	fmt.Print("enter the model: ")
	fmt.Scan(&model)
	fmt.Print("enter the hourse_power: ")
	fmt.Scan(&hourse_power)
	fmt.Print("enter the color: ")
	fmt.Scan(&colour)
	fmt.Print("enter the engine capacity: ")
	fmt.Scan(&engine_cap)

	return models.Car{
		Id: id,
		Name: name,
		Brand: brand,
		Model: model,
		HoursePower: hourse_power,
		Colour: colour,
		EngineCap: engine_cap,
	}
}