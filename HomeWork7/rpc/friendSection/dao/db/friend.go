package db

import (
	"HomeWork7/rpc/userSection/dao/db"
	"gorm.io/gorm"
)

type Relationship struct {
	gorm.Model
	SrcId     string
	DstId     string
	GroupName string
}

func (r *Relationship) TableName() string {
	return "relationship"
}
func (r *Relationship) Add() {
	//先查看你之前是否已经follow过他了，
	db.DB.Model(&Relationship{}).Where("dst").Find(&Relationship{})
}
func (r *Relationship) Delete() {

}
func (r *Relationship) Group() {

}
func (r *Relationship) ShowAll() {

}
func (r *Relationship) ShowGroup() {

}
func (r *Relationship) Search() {

}
