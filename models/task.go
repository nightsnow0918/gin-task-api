package models

type TaskStatus int

const (
	INCOMPLETED TaskStatus = 0
	COMPLETED   TaskStatus = 1
)

type Task struct {
	Id     int        `json:"id" gorm:"primaryKey"`
	Name   string     `json:"name" gorm:"not null" example:"Test Task"`
	Status TaskStatus `json:"status" gorm:"type:int;enum:0,1;default:0" example:"0"`
}
