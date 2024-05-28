package controllers

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"vermouth-backend/models"

	"github.com/gin-gonic/gin"
)

func GET_dynamic_tmsystem(c *gin.Context) {
	// var data []models.Th1Tmname
	// fmt.Println("-- data 1 == ", data)
	// err := dbRecord.Model((*models.Th1Tmname)(nil)).
	// 	Column("id", "tmname", "property", "description", "tmsubsystem_id", "tmoperation_id").
	// Where("tmsubsystem_id=?", 3).
	// Select(&data)
	var data2 []models.TmStruct

	err := dbRecord.Model().
		TableExpr("th1_tmname AS tmname").
		Column("tmname.id", "tmname.tmname", "tmname.property", "tmname.description", "tmname.tmsubsystem_id", "tmname.tmoperation_id").
		Select(&data2)
	// fmt.Println("-- data 2 == ", data2)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data2,
	})
	return
}

func GET_tmsystem(c *gin.Context) {

	// freqslice := []string{"o","p"}

	type Th1Subsystem struct {
		Id     uint16   `json:"id"`
		Tmname string   `json:"tmname"`
		Freqs  []string `json:"freqs"`
	}
	// freq := []freqslice{}

	ThSatSlice := make(map[string]map[string][]Th1Subsystem)
	var th1progmodel []models.Th1Tmprogmodel
	var subprogress []Th1Subsystem

	err := dbRecord.Model((*models.Th1Tmprogmodel)(nil)).
		Column("id", "tmname", "subsystemname").
		Select(&th1progmodel)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	ThSatSlice["THEOS"] = make(map[string][]Th1Subsystem)
	ThSatSlice["THEOS2M"] = make(map[string][]Th1Subsystem)
	ThSatSlice["THEOS2S"] = make(map[string][]Th1Subsystem)
	for m := range th1progmodel {
		ThSatSlice["THEOS"][th1progmodel[m].Subsystemname] = subprogress
	}

	// fr := []string{"1D","1H"}

	// var info []models.AnalysisInfoTheos
	// err = dbAnalysis.Model().
	// 		TableExpr("analysis_info_theos_auto_m1 AS infotheos").
	// 		Column("infotheos.id",
	// 		"infotheos.tm_name",
	// 		"infotheos.freq").
	// "infotheos.feauture_table",
	// "infotheos.transform_method",
	// "infotheos.algorithm_name",
	// "infotheos.model_address",
	// "infotheos.model_name",
	// "infotheos.create_date",
	// "infotheos.update_date",
	// "infotheos.anomaly_result_table",
	// "infotheos.start_trainpoint",
	// "infotheos.end_trainpoint",
	// "infotheos.start_traindate",
	// "infotheos.end_traindate").
	// 		Where("tm_name = ?","TMONOFFSET").
	// 		Select(&info)
	// fmt.Println("info = ", info)
	var freqs []models.FrequencyInfo
	// var freq []string
	for n := range th1progmodel {
		var freq []string

		// fmt.Println("tm name = ", th1progmodel[n].Tmname)
		err = dbAnalysis.Model().
			TableExpr("analysis_info_theos_auto_m1 AS infotb").
			Column("infotb.freq").
			Where("tm_name = ?", th1progmodel[n].Tmname).
			Select(&freqs)

		if err != nil {
			log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"massege": "Something went wrong",
			})
			return
		}
		// fmt.Println("----- ", n, " -----> ", th1progmodel[n].Id, th1progmodel[n].Tmname, freqs)

		for f := range freqs {
			freq = append(freq, freqs[f].Freq)

		}
		// fmt.Println("=> ", freq)
		tm := Th1Subsystem{th1progmodel[n].Id, th1progmodel[n].Tmname, freq}
		ThSatSlice["THEOS"][th1progmodel[n].Subsystemname] = append(ThSatSlice["THEOS"][th1progmodel[n].Subsystemname], tm)
	}

	c.JSON(http.StatusOK, gin.H{
		"Data": ThSatSlice,
	})
	return
}

func GET_tmgraphTH1(c *gin.Context) {
	satname := c.Param("satname")
	fmt.Println(satname)
	tmname := c.Param("tmname")
	// fmt.Println("TM name = ",strings.ToUpper(tmname))
	freq := c.Param("freq")
	// fmt.Println(freq)

	// var analysis_tbname string
	// if satname == "THEOS" {

	var anaanotb []models.AnaAnoTableInfo
	err := dbAnalysis.Model().
		TableExpr("analysis_info_theos_auto_m1 AS infotb").
		Column("infotb.feature_table", "infotb.anomaly_result_table").
		Where("tm_name = ? and freq = ?", tmname, freq).Select(&anaanotb)
	// fmt.Println("Info --- ", anaanotb)
	// fmt.Println("Info --- ", anaanotb[0].FeatureTable)
	// fmt.Println("Info --- ", anaanotb[0].AnomalyResultTable)
	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Qurey error, Found feature_table and anomaly_result_table inside analysis_info_theos_auto_m1",
		})
		return
	}

	var joinpersub []models.JoinOperationSubsystem
	// queopersub := "select optb.property, optb.description, optb.operationname, sub.subsystemname from (select * from th1_tmname as tm left join th1_tmoperation as oper on tm.tmoperation_id=oper.id where tm.tmname = '" + tmname +"') as optb left join th1_tmsubsystem as sub on optb.tmsubsystem_id = sub.id;"
	queopersub := "select oper.property, oper.description, oper.operationname, sub.subsystemname from (select tm.property, tm.description, tm.tmsubsystem_id, oper.operationname from th1_tmname as tm left join th1_tmoperation as oper on tm.tmoperation_id=oper.id where tm.tmname = '" + tmname + "') as oper left join th1_tmsubsystem as sub on oper.tmsubsystem_id = sub.id"

	// fmt.Println(queopersub)
	_, ps_err := dbRecord.Query(&joinpersub, queopersub)
	if ps_err != nil {
		panic(ps_err)
	}

	// tm_details := *&joinpersub
	// var detials_collect []models.JoinOperationSubsystem
	// for _, detial := range *&joinpersub {
	// 	// fmt.Println(detial)
	// 	detials := new(models.JoinOperationSubsystem)
	// 	detials.Property = detial.Property
	// 	detials.Description = detial.Description
	// 	detials.Operationname = detial.Operationname
	// 	detials.Subsystemname = detial.Subsystemname
	// 	detials_collect = append(detials_collect, *detials)

	// }
	// fmt.Println(*&joinpersub)
	// fmt.Println(detials_collect)

	var joinresult []models.JoinAnalyAnomalyTB
	quejoin := "select ana.id, ana.utc, ana.avg, ana.std, ana.count, ana.min, ana.max, ana.q1, ana.q2, ana.q3, ana.lost_state, ana.epoch_ten, ano.anomaly_state_auto_m1 from " +
		anaanotb[0].FeatureTable + " as ana left join " +
		anaanotb[0].AnomalyResultTable + " as ano on ano.id = ana.id;"
	// fmt.Println("Que = ", quejoin)
	_, aa_err := dbAnalysis.Query(&joinresult, quejoin)
	if aa_err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", aa_err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	// fmt.Println(*&jrows)
	// fmt.Println("===================")
	// fmt.Println(*&joinresult[0])
	// fmt.Println(*&joinresult[0].UTC, *&joinresult[0].Avg)

	var ds models.DataSlice
	var ano_state []float32

	for _, s := range joinresult {
		// fmt.Println(s)
		ds.Utc_tm = append(ds.Utc_tm, s.UTC)
		ds.Avg_tm = append(ds.Avg_tm, s.Avg)
		ds.Std_tm = append(ds.Std_tm, s.Std)
		ds.Min_tm = append(ds.Min_tm, s.Min)
		ds.Max_tm = append(ds.Max_tm, s.Max)
		ds.Q1_tm = append(ds.Q1_tm, s.Q1)
		ds.Q2_tm = append(ds.Q2_tm, s.Q2)
		ds.Q3_tm = append(ds.Q3_tm, s.Q3)

		if s.LostState == false {
			ano_state = append(ano_state, s.AnomalyStateAutoM1)
			if s.AnomalyStateAutoM1 == 1 {
				ds.Utc_ano1 = append(ds.Utc_ano1, s.UTC)
				ds.Ano1 = append(ds.Ano1, s.Avg)
			} else if s.AnomalyStateAutoM1 == 2 {
				ds.Utc_ano2 = append(ds.Utc_ano2, s.UTC)
				ds.Ano2 = append(ds.Ano2, s.Avg)
			} else if s.AnomalyStateAutoM1 == 3 {
				ds.Utc_ano3 = append(ds.Utc_ano3, s.UTC)
				ds.Ano3 = append(ds.Ano3, s.Avg)
			}
		} else if s.LostState == true {
			ano_state = append(ano_state, 0)
		}

	}
	// fmt.Println("=== before ====")
	// fmt.Println(ano_state)

	// Create Anomaly bar
	// Constant
	NoCount_Condition := 2
	AnoConut_Condition := 3
	// Variation
	NoCount := 0
	Ano_1 := 0
	start := ""
	end := ""
	s_collect := []string{}
	e_collect := []string{}

	// co := 0
	// for k := 0; k < AnoConut_Condition; k++ {
	// 	co += 1
	// 	fmt.Println("count loop = ",co,"  ---- k = ",k)
	// 	ano_state = append(ano_state, 0)
	// }
	for i, t := range ano_state {
		// Find start date begin
		if t != 0 && NoCount == 0 && Ano_1 == 0 && start == "" && end == "" {
			Ano_1 = 1
			start = ds.Utc_tm[i]
		} else if t != 0 && start != "" {
			Ano_1 += 1
			NoCount = 0
		} else if t == 0 && start != "" {
			NoCount += 1
		}

		if t == 0 && NoCount > NoCount_Condition && Ano_1 >= AnoConut_Condition {
			end = ds.Utc_tm[i-AnoConut_Condition]
			s_collect = append(s_collect, start)
			e_collect = append(e_collect, end)
			NoCount = 0
			Ano_1 = 0
			start = ""
			end = ""
		} else if t == 0 && NoCount > NoCount_Condition && Ano_1 < AnoConut_Condition {
			NoCount = 0
			Ano_1 = 0
			start = ""
			end = ""
		}

	}

	var collect []models.VerticalLine

	if len(s_collect) == len(e_collect) {
		for i := 0; i < len(s_collect); i++ {
			// fmt.Println(i)
			col := new(models.VerticalLine)
			col.Tyte = "rect"
			col.X0 = s_collect[i]
			col.Y0 = 0
			col.X1 = e_collect[i]
			col.Y1 = 1
			col.Xref = "x"
			col.Yref = "paper"
			col.Opacity = 0.7
			col.Fillcolor = "rgb(255, 128, 255)"
			col.Layer = "below"
			col.Line.Width = 0
			collect = append(collect, *col)
		}
	} else {
		fmt.Println("s_collect's length is not equal e_collect's length")
	}

	ds.Ano_bar = collect

	c.JSON(http.StatusOK, gin.H{
		"analysis_table": anaanotb[0].FeatureTable,
		"anomaly_table":  anaanotb[0].AnomalyResultTable,
		"data_detail":    *&joinpersub,
		"data_tm":        ds,
	})
	return
}

func POST_csvdownload(c *gin.Context) {
	var params models.CSVdownload
	c.BindJSON(&params)
	sat_name := params.SatName
	tm_name := params.TmName
	freq := params.Freq
	ana_tb := params.AnalysisTb
	ano_tb := params.AnomalyTb
	start_utc := params.StartUtc
	end_utc := params.EndUtc

	// fmt.Println("params => ", sat_name, tm_name, freq, ana_tb, ano_tb, start_utc, end_utc)

	var joincsvresult []models.AnaAnoCSVdownload
	quejoin := "select ana.utc, ana.epoch_ten, ana.avg, ana.std, ana.count, ana.min, ana.max, ana.q1, ana.q2, ana.q3, ano.anomaly_state_auto_m1 from " +
		ana_tb + " as ana left join " +
		ano_tb + " as ano on ano.id = ana.id where ana.utc >= '" +
		start_utc + "' and ana.utc <= '" + end_utc + "' and ana.lost_state=false;"
	// fmt.Println("Que = ", quejoin)
	_, aa_err := dbAnalysis.Query(&joincsvresult, quejoin)
	if aa_err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", aa_err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}
	// fmt.Println(joincsvresult)

	// fmt.Println("== >", joincsvresult[0].Avg)

	// fmt.Println(rootpath+"/download")

	var data [][]string

	header := []string{"satellite_name", "telemetry_name", "freq", "utc", "epoch", "average",
		"count", "maximum", "minimum", "standard_deviation",
		"quartile1", "quartile2", "quartile3", "anomaly_state"}

	data = append(data, header)

	for _, rows := range joincsvresult {
		row := []string{}
		row = append(row, sat_name, tm_name, freq, rows.UTC, rows.EpochTen, rows.Avg,
			rows.Count, rows.Max, rows.Min, rows.Std, rows.Q1, rows.Q2, rows.Q3, rows.AnomalyStateAutoM1)
		data = append(data, row)
	}

	start_time := strings.Split(strings.Split(start_utc, " ")[0], "-")
	start_time_name := start_time[0] + start_time[1] + start_time[2]
	end_time := strings.Split(strings.Split(end_utc, " ")[0], "-")
	end_time_name := end_time[0] + end_time[1] + end_time[2]

	csvfile_name := sat_name + "_" + tm_name + "_" + start_time_name + "_" + end_time_name + ".csv"
	// fmt.Println(csvfile_name)
	InitiateConfiguration()
	csvfile_name_path := rootpath + "/download/" + csvfile_name
	csvFile, err := os.Create(csvfile_name_path)
	if err != nil {
		log.Fatalf("failed creating csv file : %s", err)
	}
	// fmt.Println(csvfile_name_path)

	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range data {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()

	c.Writer.Header().Add("Content-Disposition", fmt.Sprintf(csvfile_name))
	c.Writer.Header().Add("Content-Type", "application/octet-stream")
	// c.Writer.Header().Add("Content-Type", "text/csv")
	c.File(csvfile_name_path)

	// Romove .csv file after it was send
	err = os.Remove(csvfile_name_path)
	if err != nil {
		log.Fatal()
	}

}
