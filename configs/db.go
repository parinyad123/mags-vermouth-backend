package configs

import (
	"vermouth-backend/controllers"
	// "vermouth-backend/models"
	"log"
	"os"

	"github.com/go-pg/pg/v10"
)

const (
	addr = "local:5432"
	user = ""	
	password = ""
	db_analysis = ""
	db_record = ""
)

func Connect_tmrecord() *pg.DB {
	opts_record := &pg.Options{
		User: user,
		Password: password,
		Addr: addr,
		Database: db_record,
	}

	var db_record_pg *pg.DB = pg.Connect(opts_record)

	if db_record_pg == nil {
		log.Printf("Failed to connect tm_record DB ...")
		os.Exit(100)
	}

	log.Printf("Connect to DB success tm_record DB ...")
	controllers.InitiateRecordDB(db_record_pg)
	return db_record_pg

}
 
func Connect_tmanalysis() *pg.DB {
	opts_analysis := &pg.Options{
		User: user,
		Password: password,
		Addr: addr,
		Database: db_analysis,
	}

	var db_analysis_pg *pg.DB = pg.Connect(opts_analysis)

	if db_analysis_pg == nil {
		log.Printf("Failed to connect tm_analysis DB ...")
		os.Exit(100)
	}

	log.Printf("Connect to DB success tm_analysis DB ...")
	controllers.InitiateAnalysisDB(db_analysis_pg)
	return db_analysis_pg

}