package main

import (
	"ventra.com/backend/dbconfig"
	"ventra.com/backend/models"
	"ventra.com/backend/api/route"
)

func main() {
    db, err := dbconfig.ConnectDB()
	if err != nil {
		panic("Connection failed!")
	}

	db.AutoMigrate(
		&models.LeadsTags{},
		&models.LeadsData{},
		&models.StatsData{},
	)

	router := route.SetupRouter(db)
	router.Run()
}