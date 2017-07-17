/*
  @@copyright: edroity.com
  @author zhaojianwei at 2017.07.14
  @reference http://godoc.org/gopkg.in/mgo.v2#Bulk.Update
 */
package mongo

import (
	"edroity.com/server/common/log"
	"gopkg.in/mgo.v2"
	"os"
)

var session *mgo.Session

type DBHandler struct {
	Session *mgo.Session
}

func init() {
	//TODO 解析配置文件
	var err error
	if session, err = mgo.Dial(""); err != nil {
		log.Error("[mgo] 数据连接失败:", err)
		os.Exit(1)
	}
}

func NewDBHandler() *DBHandler {
	return &DBHandler{Session: session}
}

func (dh *DBHandler) Find(db, c string, sql interface{}) *mgo.Query {
	return dh.Session.DB(db).C(c).Find(sql)
}

func (dh *DBHandler) Insert(db, c string, sql ...interface{}) error {
	return dh.Session.DB(db).C(c).Insert(sql)
}

func (dh *DBHandler) Remove(db, c string, sql interface{}) error {
	return dh.Session.DB(db).C(c).Remove(sql)
}

func (dh *DBHandler) RemoveAll(db, c string, sql ...interface{}) (*mgo.ChangeInfo, error) {
	return dh.Session.DB(db).C(c).RemoveAll(sql)
}

func (dh *DBHandler) Update(db, c string, sql ...interface{}) error {
	return dh.Session.DB(db).C(c).Update(sql[0], sql[1])
}

func Shutdown() {
	if session != nil {
		session.Close()
	}
}
