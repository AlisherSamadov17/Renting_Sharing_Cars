package postgres

import (
	"database/sql"
	"fmt"
	"main/config"
)



type Store struct{
	Db *sql.DB
	CarStorage carRepo
}

func New(cfg config.Config) (Store, error) {
	url := fmt.Sprintf(`host=%s port=%s user=%s password=%s database=%s sslmode=disable`, cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB)

	db, err := sql.Open("postgres", url)
	if err != nil {
		return Store{}, err
	}

	carRepo :=NewCarRepo(db)
          
	return Store{
		Db: db,
		CarStorage: carRepo,
	}, nil
}

// new user repo enter this dont ignore