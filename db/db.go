package db

import (
	"database/sql"
	"log"
)

func DatabaseInit(db *sql.DB) error{
	//TODO INIT FUNC
	DDLs:= []string{UserDDL}
	for _, ddl :=  range DDLs{
		_, err := db.Exec(ddl)
		if err != nil {
			log.Fatal("can't init $s err is %e", ddl, err)
		}
	}
	return nil
}
