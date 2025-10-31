package main

import (
	"context"
	"fmt"

	"gorm.io/gorm"
	"ventra.com/backend/dbconfig"
	"ventra.com/backend/models"
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

	ctx := context.Background()

	tag, err := gorm.G[models.LeadsTags](db).Where("id = ?", 1).First(ctx)
	if err != nil {
		panic("Query failed!")
	}

	fmt.Println(tag.TagsTitle)
}