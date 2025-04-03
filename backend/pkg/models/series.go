package models

import (
	"gorm.io/gorm"
)

type Series struct {
	gorm.Model     `swaggerignore:"true"`
	Title          string `json:"title"`
	Description    string `json:"description"`
	Seasons        int    `json:"seasons"`
	Episodes       int    `json:"episodes"`
	Genre          string `json:"genre"`
	Status         string `json:"status" gorm:"default:'To Watch'"`
	CurrentEpisode int    `json:"current_episode" gorm:"default:0"`
	Score          int    `json:"score" gorm:"default:0"`
}
