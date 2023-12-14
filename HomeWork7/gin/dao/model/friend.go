package model

import "gorm.io/gorm"

// FriendApplication 该表用于好友申请，接受申请则删除字段，并在Friend表中添加字段，拒绝则相反
type FriendApplication struct {
	gorm.Model
	SrcId string
	DstId string
}
type Friend struct {
	gorm.Model
	SrcId string
	DesId string
}

func (f *Friend) TableName() string {
	return "friend"
}
func (f *FriendApplication) TableName() string {
	return "friend_application"
}
