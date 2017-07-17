package game

import (
	"edroity.com/server/common/log"
	"gopkg.in/mgo.v2/bson"
	"edroity.com/server/common/db/mongo"
)

func QueryUser()  {
	db := mongo.NewDBHandler()
	err := db.Find("nba", "user", bson.M{"name": "admin"})
	if err != nil {
		log.Error("[QueryUser] 查询用户基本信息异常:", err)
	}
}