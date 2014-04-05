package models

import (
	"database/sql"
	_ "github.com/Go-SQL-Driver/MySQL"
	"github.com/coopernurse/gorp"
	"log"
)

var (
	dbmap *gorp.DbMap
)

func init() {
	// initialize the DbMap

	// connect to db using standard Go database/sql API
	// use whatever database/sql driver you wish
	db, err := sql.Open("mysql", "root:Chase0the0Wolf@/jobgeniusdb")
	if err != nil {
		log.Printf("%s", err)
	}

	// construct a gorp DbMap
	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}
	defer dbmap.Db.Close()
}
