package controllers

import (
	"vermouth-backend/models"
	"log"
	"net/http"
	"fmt"
	"os"
	"strings"
	"encoding/csv"
	"reflect"

	"github.com/gin-gonic/gin"
)

func GET_THEOSchartfilters(c *gin.Context) {
	var th1progmodel []models.Th1Tmprogmodel
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

	var analysis_infotheosm1 []models.AnalysisInfoTheosAutoM1
	err_info := dbAnalysis.Model((*models.AnalysisInfoTheosAutoM1)(nil)).
		Column("id", "tm_name", "freq", "anomaly_result_table").
		Where("id>?", 15).
		Select(&analysis_infotheosm1)

	if err != nil {
		log.Panicf("Error getting TM anomaly data, Reason: %v\n", err_info)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	var minmax_date models.MaxMinDate
	var tmfreqDetail []models.TmDetail
	subsystemfilters := make(map[string][]models.TmDetail)
	counttm := 0
	for _, thpro := range th1progmodel {
		// fmt.Println("----------------------------")
		fd := models.FreqDate {}
		freqdate := []models.FreqDate {}
		// tf := models.TmFreq {}

		td := models.TmDetail {}
		for _, anainfo := range analysis_infotheosm1 {
			if thpro.Tmname == anainfo.TmName {
			
				counttm = counttm+ 1

				sql_date := "select date(max(utc)) max_date, date(min(utc)) min_date from " + anainfo.AnomalyResultTable +";"
				// fmt.Println(sql_date)
				_, err_date := dbAnalysis.Query(&minmax_date, sql_date)
				if err_date != nil {
					log.Panicf("Error getting TM anomaly data, Reason: %v\n", err_info)
					c.JSON(http.StatusInternalServerError, gin.H{
						"status":  http.StatusInternalServerError,
						"massege": "Something went wrong",
					})
					return
				}
				fd = models.FreqDate {
					"freq": anainfo.Freq,
					"minDate": minmax_date.MinDate,
					"maxDate": minmax_date.MaxDate,
				}

				freqdate = append(freqdate, fd)

				td.TmName = thpro.Tmname
				td.FreqDates = freqdate
				
			}
		}

		tmfreqDetail = append(tmfreqDetail, td)
		
		// fmt.Println("==> ", freqdate)
		// fmt.Println("=====> ", thpro.Subsystemname, tf)
		// fmt.Println("=====----> ", subsystemfilters)
		// fmt.Println("--------  subs --------")
		subsystemfilters[thpro.Subsystemname] = append(subsystemfilters[thpro.Subsystemname], td)
		// fmt.Println("=====----> ", subsystemfilters)

	}




	c.JSON(http.StatusOK, gin.H{
		// "tmDetails":tmfreqDetail ,
		"THEOSchartfilter": subsystemfilters,
	})
	return
}

func GET_THEOS_chartanomaly(c *gin.Context) {
	// satname := c.Param("satname")
	// fmt.Println(satname)
	tmname := c.Param("tmname")
	// fmt.Println("TM name = ",strings.ToUpper(tmname))
	freq := c.Param("freq")

	// fmt.Println(satname, tmname, freq)
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
	// fmt.Println("Detail = ",*&joinpersub)
	// fmt.Println(detials_collect)

	var joinresult []models.JoinAnalysisAnomalyTB
	quejoin := "select ana.id, ana.utc, ana.avg, ana.std, ana.count, ana.min, ana.max, ana.q1, ana.q2, ana.q3, ana.lost_state, ana.epoch_ten, ano.anomaly_state_auto_m1 from " +
		anaanotb[0].FeatureTable + " as ana left join " +
		anaanotb[0].AnomalyResultTable + " as ano on ano.id = ana.id order by epoch_ten;"
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

	var ds models.StatisticsTHEOSSlice
	var ano_state []float32

	for _, s := range joinresult {
		// fmt.Println(s)
		ds.Utc_tm = append(ds.Utc_tm, s.UTC)
		ds.Avg_tm = append(ds.Avg_tm, s.Avg)
		ds.Std_tm = append(ds.Std_tm, s.Std)
		ds.Count_tm = append(ds.Count_tm, s.Count)
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
	AnoCount_Condition := 3
	// Variation
	NoCount := 0
	Ano_1 := 0
	start := ""
	end := ""
	s_collect := []string{}
	e_collect := []string{}

	// co := 0
	// for k := 0; k < AnoCount_Condition; k++ {
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

		if t == 0 && NoCount > NoCount_Condition && Ano_1 >= AnoCount_Condition {
			end = ds.Utc_tm[i-AnoCount_Condition]
			s_collect = append(s_collect, start)
			e_collect = append(e_collect, end)
			NoCount = 0
			Ano_1 = 0
			start = ""
			end = ""
		} else if t == 0 && NoCount > NoCount_Condition && Ano_1 < AnoCount_Condition {
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

func GET_THEOSDownloadcsv(c *gin.Context) {
	// var params 
	// var params models.CSVdownload
	// c.BindJSON(&params)
	// sat_name := params.SatName
	// tm_name := params.TmName
	// freq := params.Freq
	// ana_tb := params.AnalysisTb
	// ano_tb := params.AnomalyTb
	// start_utc := params.StartUtc
	// end_utc := params.EndUtc

	sat_name := c.Param("satname")
	tm_name := c.Param("tmname")
	freq := c.Param("freq")
	ana_tb := c.Param("analysis_table")
	ano_tb := c.Param("anomaly_table")
	start_utc := c.Param("start_utc")
	end_utc := c.Param("end_utc")


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

	// fmt.Println("Check root path = ", rootpath+"/download")

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
	// fmt.Println("CSV file = ",csvfile_name_path)
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

func POST_downloadCSV(c *gin.Context) {
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

func GET_THEOSDownload_staticscsv(c *gin.Context) {
	sat_name := c.Param("satname")
	tm_name := c.Param("tmname")
	freq := c.Param("freq")
	ana_tb := c.Param("analysis_table")
	ano_tb := c.Param("anomaly_table")
	start_utc := c.Param("start_utc")
	end_utc := c.Param("end_utc")


	fmt.Println("params => ", sat_name, tm_name, freq, ana_tb, ano_tb, start_utc, end_utc)

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
	fmt.Println(joincsvresult)
	fmt.Println(reflect.TypeOf(joincsvresult))

	// fmt.Println("== >", joincsvresult[0].Avg)

	// fmt.Println("Check root path = ", rootpath+"/download")

	// Convert struct data to CSV format
	var csvData [][]string

	header := []string{"satellite_name", "telemetry_name", "freq", "utc", "epoch", "average",
		"count", "maximum", "minimum", "standard_deviation",
		"quartile1", "quartile2", "quartile3", "anomaly_state"}

	csvData = append(csvData, header)

	for _, rows := range joincsvresult {
		row := []string{}
		row = append(row, sat_name, tm_name, freq, rows.UTC, rows.EpochTen, rows.Avg,
			rows.Count, rows.Max, rows.Min, rows.Std, rows.Q1, rows.Q2, rows.Q3, rows.AnomalyStateAutoM1)
		csvData = append(csvData, row)
	}

	// fmt.Println("data sat = > ", data)

	// Set response headers for CSV download
	c.Writer.Header().Set("Content-Type", "text/csv")
	c.Writer.Header().Set("Content-Disposition", "attachment; filename=data.csv")

	// Create a csv writer
	writer := csv.NewWriter(c.Writer)
	defer writer.Flush()

	// Write thw csv data
	for _, dataRow := range csvData {
		err := writer.Write(dataRow)
		if err != nil {
			c.String(http.StatusInternalServerError, err.Error())
			return
		}
	}
	c.Status(http.StatusOK)
}