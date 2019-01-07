package Config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Db struct{
	Type string
	Host string
	Port string
	database *gorm.DB
}


func (db *Db)Connect() error{
	conn, err := gorm.Open(db.Type, db.Host+":"+db.Port)
	if err != nil {
		fmt.Println("db err: ", err)
		return err;
	}
	conn.DB().SetMaxIdleConns(10)
	conn.LogMode(true)
	db.database = conn
	return nil
}

