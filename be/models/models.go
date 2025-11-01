package models

import (
	"time"
	"gorm.io/gorm"
)

type LeadsTags struct {
	gorm.Model
	TagsTitle string `json:"tags_title" gorm:"unique;not null"`
}

type LeadsData struct {
	gorm.Model
	DateIn 			time.Time `json:"date_in"`
	WhatsappId		string `json:"whatsapp_id" gorm:"index;unique"`
	PhoneNumber		string `json:"phone_number" gorm:"index;unique"`
	Name			string `json:"name"`
	Company			string `json:"company"`
	Notes			string `json:"notes"`
	Tags			[]LeadsTags `gorm:"many2many:leads_data_tags;" json:"tags"`
}

type StatsData struct {
	gorm.Model
	LeadsAlltime	int	`json:"leads_alltime" gorm:"autoUpdateTime"`
	LeadsDaily		int `json:"leads_daily" gorm:"autoUpdateTime"`
	FollowUpToday	int `json:"follow_up_today" gorm:"autoUpdateTime"`
	DealsAlltime	int `json:"deals_alltime" gorm:"autoUpdateTime"`
}