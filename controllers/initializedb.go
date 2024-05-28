package controllers

import (
	// "path"
	"os"
	"fmt"

	"github.com/go-pg/pg/v10"
)

var dbRecord *pg.DB
var dbAnalysis *pg.DB
// var pathWindow string
var rootpath string

func InitiateRecordDB(dbrecord *pg.DB) {
	dbRecord = dbrecord
}

func InitiateAnalysisDB(dbanalysis *pg.DB) {
	dbAnalysis = dbanalysis
}

func InitiateConfiguration() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	rootpath = path
}


