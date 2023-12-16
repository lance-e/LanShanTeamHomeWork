package controller

import (
	"HomeWork7/rpc/friendSection/dao/db"
	"HomeWork7/rpc/friendSection/friend"
	"context"
	"errors"
)

type FriendServer struct {
	friend.UnimplementedFriendServer
}

func (r *FriendServer) FriendServe(ctx context.Context, relationship *friend.RelationShip) (*friend.Response, error) {
	mode := relationship.Mode
	relative := &db.Relationship{
		SrcId:     relationship.SrcId,
		DstId:     relationship.DstId,
		GroupName: relationship.DstInSrcGroup,
	}
	switch mode {
	case "0":
		//添加好友
		relative.Add()

	case "1":
		//删除好友
	case "2":
		//给好友分组
	case "3":
		//显示所有好友
	case "4":
		//显示分组好友
	case "5":
		//好友搜索
	default:
		break
	}

	return &friend.Response{
		Result: false,
	}, errors.New("undefined error")
}
