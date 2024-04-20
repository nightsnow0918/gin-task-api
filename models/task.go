package models

type Task struct {
	Id     int    `json:"id" gorm:"primaryKey"`
	Name   string `json:"name" gorm:"not null" example:"Test Task"`
	Status int    `json:"status" gorm:"type:int;enum:0,1;default:0" example:"0"`
}
