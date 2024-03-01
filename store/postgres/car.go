package postgres

import (
	"database/sql"
	"fmt"
	"main/models"

	"github.com/google/uuid"
)

type carRepo struct{
	DB * sql.DB
}

func NewCarRepo(db *sql.DB) carRepo {
	return carRepo{
		DB: db,
	}
}

func (c carRepo) CreateCar(car models.Car) (string, error) {
	fmt.Println("car :",car)
	id := uuid.New()
	if _,err := c.DB.Exec(`insert into cars (id,name,brand,model,hourse_power,colour,engine_cap) values($1,$2,$3,$4,$5,$6,$7)`,id,car.Name,car.Brand,car.Model,car.HoursePower,car.Colour,car.EngineCap);err != nil{
		return "",err
	}
	return id.String(),nil
}

func (c carRepo) DeleteCar(id string) error {
	if _,err := c.DB.Exec(`delete from cars where id = $1`,id);err!=nil{
		return err
	}
	return nil
}

func (c carRepo) UpdateCar(car models.Car) error {
	if _,err := c.DB.Exec(`update cars set name=$2,brand=$3,model=$4,hourse_power=$5,colour=$6,engine_cap=$7 where id=$1`,car.Id,car.Name,car.Brand,car.Model,car.HoursePower,car.Colour,car.EngineCap); err != nil {
		return err
	}
	return nil
}

func (c carRepo) GetByIdCar(id uuid.UUID) (models.Car,error) {
	car :=models.Car{}

	if err := c.DB.QueryRow(`select id,name,brand,model,hourse_power,colour,engine_cap,created_at where id = $1`,id).Scan(
		&car.Id,
		&car.Name,
		&car.Brand,
		&car.Model,
		&car.HoursePower,
		&car.Colour,
		&car.EngineCap,
		&car.CreatedAt,
	);err != nil{
		return models.Car{},err
	}
	return car,nil
}


func (c carRepo) GetAllCars() ([]models.Car,error) {
	cars:=[]models.Car{}
	rows,err := c.DB.Query(`select * from users`)
	if err != nil {
		return nil,err
	}
	for rows.Next(){
		car := models.Car{}

		if err := rows.Scan(&car.Id,&car.Name,&car.Brand,&car.Model,&car.HoursePower,&car.Colour,&car.EngineCap,&car.CreatedAt);err != nil {
			return nil,err
		}
		cars = append(cars, car)
	}
	return cars,nil
}