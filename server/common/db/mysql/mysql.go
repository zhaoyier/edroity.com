package mysql

import (
	"database/sql"
	"errors"

	"edroity.com/server/common/conf"
	"edroity.com/server/common/log"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type DBHandler struct {
	DB *sql.DB
}

func init() {
	var err error
	db, err = sql.Open(conf.ServerConf.MysqlConf.DriverName, conf.ServerConf.MysqlConf.DataSourceName)
	checkErr(err)
	// 设置最大打开的连接数，默认值为0表示不限制,避免并发太高导致连接mysql出现too many connections的错误
	db.SetMaxOpenConns(conf.ServerConf.MysqlConf.MaxOpenConns)
	// 设置闲置的连接数,当开启的一个连接使用完成后可以放在池里等候下一次使用
	db.SetMaxIdleConns(conf.ServerConf.MysqlConf.MaxIdleConns)
	db.Ping()
}

//get dbhandler
func NewDBHandler() *DBHandler {
	return &DBHandler{DB: db}
}

//insert
func (this *DBHandler) Insert(sql string, args ...interface{}) (int64, error) {
	stmt, err := this.DB.Prepare(sql)
	if err != nil {
		log.Error("prepare sql for insert faild, sql:", sql, " error:", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(args...)
	if err != nil {
		log.Error("execute sql for insert faild, sql:", sql, " error:", err)
		return 0, err
	}
	if res == nil { //just in case
		log.Error("execute sql for insert faild, sql:", sql, " returned null result")
		return 0, errors.New("execute sql for insert faild, sql:" + sql + " returned null result")
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Error("get LastInsertId failed, sql:", sql, " error:", err)
		return 0, nil
	}
	return id, nil
}

//insert,return the RowsAffected column
func (this *DBHandler) InsertWithRowsAffected(sql string, args ...interface{}) (int64, error) {
	stmt, err := this.DB.Prepare(sql)
	if err != nil {
		log.Error("prepare sql for insert faild, sql:", sql, " error:", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(args...)
	if err != nil {
		log.Error("execute sql for insert faild, sql:", sql, " error:", err)
		return 0, err
	}
	if res == nil { //just in case
		log.Error("execute sql for insert faild, sql:", sql, " returned null result")
		return 0, errors.New("execute sql for insert faild, sql:" + sql + " returned null result")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Error("get RowsAffected failed, sql:", sql, " error:", err)
		return 0, nil
	}
	return rowsAffected, nil
}

// Delete deletes date according to the sql and args passed in,
// and returns affected on success
func (this *DBHandler) Delete(sql string, args ...interface{}) (int64, error) {
	stmt, err := this.DB.Prepare(sql)
	if err != nil {
		log.Error("prepare sql for Delete faild, sql:", sql, " error:", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(args...)
	if err != nil {
		log.Error("execute sql for Delete faild, sql:", sql, " error:", err)
		return 0, err
	}
	if res == nil { //just in case
		log.Error("execute sql for Delete faild, sql:", sql, " returned null result")
		return 0, errors.New("execute sql for Delete faild, sql:" + sql + " returned null result")
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Error("get affected rows failed, sql:", sql, " error:", err)
		return 0, err
	}
	return affect, nil
}

//更新数据
func (this *DBHandler) Update(sql string, args ...interface{}) (int64, error) {
	stmt, err := this.DB.Prepare(sql)
	if err != nil {
		log.Error("prepare sql for Update faild, sql:", sql, " error:", err)
		return 0, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(args...)
	if err != nil {
		log.Error("execute sql for Update faild, sql:", sql, " error:", err)
		return 0, err
	}
	if res == nil { //just in case
		log.Error("execute sql for Update faild, sql:", sql, " returned null result")
		return 0, errors.New("execute sql for Update faild, sql:" + sql + " returned null result")
	}
	affect, err := res.RowsAffected()
	if err != nil {
		log.Error("get affected rows failed, sql:", sql, " error:", err)
		return 0, err
	}
	return affect, nil
}

//查询数据，带参数
func (this *DBHandler) Query(sql string, args ...interface{}) (records []map[string]string, err error) {
	rows, err := this.DB.Query(sql, args...)
	if err != nil {
		log.Error("Query faild, sql:", sql, " error:", err)
		return nil, err
	}
	defer rows.Close()
	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	if len(columns) == 0 {
		return records, nil
	}
	values := make([][]byte, len(columns))
	scanArgs := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Error("rows.Scan failed, sql:", sql, " error:", err)
			return nil, err
		}
		record := make(map[string]string)
		for i, col := range values {
			if col == nil {
				record[columns[i]] = ""
			} else {
				record[columns[i]] = string(col)
			}
		}
		records = append(records, record)
	}
	return records, nil
}

//查询数据
func QueryAll(sql string, db *sql.DB) (map[string]string, error) {
	rows, err := db.Query(sql)
	if err != nil {
		log.Error("Query faild, sql:", sql, " error:", err)
		return nil, err
	}
	defer rows.Close()
	//字典类型
	//构造scanArgs、values两个数组，scanArgs的每个值指向values相应值的地址
	columns, _ := rows.Columns()
	if len(columns) == 0 {
		return nil, nil
	}
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		if err != nil {
			log.Error("rows.Scan failed, sql:", sql, " error:", err)
			return nil, err
		}
		for i, col := range values {
			if col == nil {
				record[columns[i]] = ""
			} else {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}
	return record, nil
}

//查询单条数据
func QueryRow(sql string, db *sql.DB, values interface{}, args ...interface{}) {
	row := db.QueryRow(sql, args...)
	err := row.Scan(values)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Error(err)
	}
}
