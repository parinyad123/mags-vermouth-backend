package controllers

import (
	"vermouth-backend/models"
	"fmt"
	"log"
	"net/http"
	"strings"

	"time"

	"github.com/gin-gonic/gin"
)

func GET_anomalyweekly(c *gin.Context) {

	var countweeklies []models.CountAnomalyweekly
	err := dbAnalysis.Model((*models.CountAnomalyweekly)(nil)).
		Column("date", "lastweekday", "count").
		Select(&countweeklies)

	if err != nil {
		log.Panicf("Error getting Amount anomaly points weekly, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	var ck models.CountWeekly
	var count_weekly []models.CountWeekly

	const (
		layoutISO = "2006-01-02"
		layoutUS  = "Jan 2, 2006"
		layoutTH  = "2 Jan 2006"
	)

	for _, c := range countweeklies {
		t_date, _ := time.Parse(layoutISO, c.Date)
		t_lastweekday, _ := time.Parse(layoutISO, c.Lastweekday)
		ck.Week = t_date.Format(layoutTH) + " - " + t_lastweekday.Format(layoutTH)
		ck.SatTH = c.Count
		ck.SatTHMain = "-"
		ck.SatTHSmall = "-"
		count_weekly = append(count_weekly, ck)
	}

	// fmt.Println(count_weekly)
	c.JSON(http.StatusOK, gin.H{
		"anomalyweekly": count_weekly,
	})
	return
}

func GET_dailyfilter(c *gin.Context) {
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

	subtm := make(map[string][]string)
	for _, m := range th1progmodel {
		subtm[m.Subsystemname] = append(subtm[m.Subsystemname], m.Tmname)
	}

	var subsystems []string
	for s, _ := range subtm {
		subsystems = append(subsystems, s)
	}
	// 	fmt.Println(subsystems)
	// 	fmt.Println(subtm)
	var maxmin_date []models.MaxMinDate
	maxmindate_sql := "select date(max(count_date)) max_date , date(min(count_date)) min_date from countdaily_tmanomaly_theos_auto_m1;"
	_, mm_err := dbAnalysis.Query(&maxmin_date, maxmindate_sql)
	if mm_err != nil {
		log.Panicf("Error getting Max, Min date, Reason: %v\n", mm_err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	var startend_date []string
	// startend_date := [maxmin_date[0].MaxDate mamaxmin_date[0].MinDate]
	startend_date = append(startend_date, maxmin_date[0].MinDate)
	startend_date = append(startend_date, maxmin_date[0].MaxDate)

	// fmt.Println(maxmin_date)
	c.JSON(http.StatusOK, gin.H{
		"subsystems":     subsystems,
		"tm":             subtm,
		"startend_dates": startend_date,
	})
	return
}

func GET_reportdailyfilter(c *gin.Context) {
	var tmsub []models.Th1Dailyfilter
	err_ts := dbRecord.Model((*models.Th1Dailyfilter)(nil)).
		Column("tmname", "subsystemname").Select(&tmsub)

	if err_ts != nil {
		log.Panicf("Error fetch data from , Reason: %v\n", err_ts)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}
	// fmt.Println(tmsub)

	var countinfo []models.CountInfoTheosAutoM1
	err_coin := dbAnalysis.Model((*models.CountInfoTheosAutoM1)(nil)).
		Column("id", "tm_name", "freq", "anomaly_table", "feature_table", "record_table", "analysis_params_id", "analysis_info_id").Select(&countinfo)
	if err_coin != nil {
		log.Panicf("Error fetch data from , Reason: %v\n", err_coin)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}
	// fmt.Println(countinfo)
	subtm := make(map[string][]string)
	var subsystems []string
	for _, c := range countinfo {
		// fmt.Println(c.TmName)

		for _, ts := range tmsub {
			if c.TmName == ts.Tmname {
				subtm[ts.Subsystemname] = append(subtm[ts.Subsystemname], c.TmName)
				subsystems = append(subsystems, ts.Subsystemname)
			}
		}
	}

	// remove duplicated subsystem
	subkey := make(map[string]bool)
	subsystemsSlice := []string{}
	for _, s := range subsystems {
		if _, value := subkey[s]; !value {
			subkey[s] = true
			subsystemsSlice = append(subsystemsSlice, s)
		}
	}

	// fmt.Println(subsystems)
	// fmt.Println(subsystemsSlice)
	// fmt.Println(subtm)

	var maxmin_date []models.MaxMinDate
	maxmindate_sql := "select date(max(count_date)) max_date , date(min(count_date)) min_date from countdaily_tmanomaly_theos_auto_m1;"
	_, mm_err := dbAnalysis.Query(&maxmin_date, maxmindate_sql)
	if mm_err != nil {
		log.Panicf("Error getting Max, Min date, Reason: %v\n", mm_err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}
	var startend_date []string
	startend_date = append(startend_date, maxmin_date[0].MinDate)
	startend_date = append(startend_date, maxmin_date[0].MaxDate)

	// fmt.Println(maxmin_date)
	c.JSON(http.StatusOK, gin.H{
		"subsystems":     subsystemsSlice,
		"tm":             subtm,
		"startend_dates": startend_date,
	})
	return
}

func GET_reportalldaily_proviousmonth(c *gin.Context) {
	var lastdate []models.LastDatecount
	lastdate_sql := "select date(max(count_date)) lastdate from countdaily_tmanomaly_theos_auto_m1;"
	_, last_err := dbAnalysis.Query(&lastdate, lastdate_sql)
	if last_err != nil {
		log.Panicf("Error getting last date, Reason: %v\n", last_err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}
	// convert string to time
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "Jan 2, 2006"
		layoutTH  = "2 Jan 2006"
	)
	last_date, _ := time.Parse(layoutISO, lastdate[0].Lastdate)

	// provious 30 days
	proviousmonth_date := last_date.AddDate(0, 0, -30)

	proviousmonth_dateISO := proviousmonth_date.Format(layoutISO)

	part1 := "select proviousmonth.count_date, proviousmonth.countamount_per_day, countinfo.tm_name, countinfo.freq from (select * from countdaily_tmanomaly_theos_auto_m1 where count_date > '"
	part2 := "') as proviousmonth left join count_info_theos_auto_m1 as countinfo on countinfo.id = proviousmonth.count_info_id;"
	proviousmonth_sql := part1 + proviousmonth_dateISO + part2
	// fmt.Println(proviousmonth_sql)

	// var proviousmonth []models.ProviousMonth
	// _, prem_err := dbAnalysis.Query(&proviousmonth, proviousmonth_sql)
	// if prem_err != nil {
	// 	log.Panicf("Error getting proviou smonth date, Reason: %v\n", prem_err)
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"status":  http.StatusInternalServerError,
	// 		"massege": "Something went wrong",
	// 	})
	// 	return
	// }
	// fmt.Println("proviousmonth = ",proviousmonth)

	// var tmsubsystems []models.Th1Dailyfilter
	// err_ts := dbRecord.Model((*models.Th1Dailyfilter)(nil)).
	// Column("tmname", "subsystemname").Select(&tmsubsystems)
	// if err_ts != nil {
	// 	log.Panicf("Error fetch data from , Reason: %v\n", err_ts)
	// 	c.JSON(http.StatusInternalServerError, gin.H{
	// 		"status":  http.StatusInternalServerError,
	// 		"massege": "Something went wrong",
	// 	})
	// 	return
	// }

	// tmsubsystem := make(map[string][]string)

	// for _, tms := range(tmsubsystems) {
	// 	tmsubsystem[tms.Tmname] = append(tmsubsystem[tms.Tmname], tms.Subsystemname)
	// }

	// fmt.Println(tmsubsystem)

	// var reportddailytables []models.ReportdDailyTable
	// var report models.ReportdDailyTable
	// for _, pro := range(proviousmonth) {
	// 	report.Date = strings.Split(pro.CountDate, " ")[0]
	// 	if len(tmsubsystem[pro.TmName]) > 1 {
	// 		report.Subsystem = tmsubsystem[pro.TmName][0]
	// 		for i := 1; i < len(tmsubsystem[pro.TmName]); i++ {
	// 			report.Subsystem = report.Subsystem +  "," + tmsubsystem[pro.TmName][i]
	// 		}
	// 	} else {
	// 		report.Subsystem = tmsubsystem[pro.TmName][0]
	// 	}
	// 	report.Name = pro.TmName
	// 	report.Freq = pro.Freq
	// 	report.AnomalyPoints = pro.CountamountPerDay
	// 	reportddailytables = append(reportddailytables, report)
	// }

	// c.JSON(http.StatusOK, gin.H{
	// 	"reportdailytable": reportddailytables,
	// })
	// return
	REPOST_dailyanomalytable(proviousmonth_sql, c)
}

func POST_reportdaily(c *gin.Context) {
	var rep models.PostReportDaily
	c.BindJSON(&rep)
	satname := rep.Satname
	tmnames := rep.Tmnames
	optiondate := rep.Optiondate
	dates := rep.Dates

	fmt.Println("POST_reportdaily = ", satname, tmnames, optiondate, dates)

	tmcondition := " where tmjoin.tm_name = '" + tmnames[0] + "'"
	// var cond string
	cond := ""
	if len(tmnames) > 1 {
		for i := 1; i < len(tmnames); i++ {
			cond = cond+" or tmjoin.tm_name = '" + tmnames[i] + "'"
		}
		tmcondition = tmcondition+cond
	}

	part1 := "select * from (select tmcount.count_date, tmcount.countamount_per_day, count_info_theos_auto_m1.tm_name, count_info_theos_auto_m1.freq from (select * from countdaily_tmanomaly_theos_auto_m1 "
	part2 := ") as tmcount left join count_info_theos_auto_m1 on tmcount.count_info_id = count_info_theos_auto_m1.id) as tmjoin"
	
	var datecondition string
	if optiondate == "range" {
		datecondition = " where count_date between '" + dates[0] + "' and '" + dates[1] + "'"
	} else if optiondate == "multiple" {
		datecondition = " where count_date = '" + dates[0] + "'"
		// var condmul string
		condmul := ""
		if len(dates) > 1 {
			for m := 1; m<len(dates); m++ {
			condmul = condmul+" or count_date = '"+ dates[m] + "'"
			}
		datecondition = datecondition+condmul
		}		
	} else if optiondate == "all" {
		datecondition = ""

	} 

	sql := part1 + datecondition + part2 + tmcondition + ";"
	// fmt.Println(sql)
	REPOST_dailyanomalytable(sql, c)

}

func REPOST_dailyanomalytable(sql string, c *gin.Context) {
	var countdaily []models.CountDaily
	_, prem_err := dbAnalysis.Query(&countdaily, sql)
	if prem_err != nil {
		log.Panicf("Error getting proviou smonth date, Reason: %v\n", prem_err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	var tmsubsystems []models.Th1Dailyfilter
	err_ts := dbRecord.Model((*models.Th1Dailyfilter)(nil)).
		Column("tmname", "subsystemname").Select(&tmsubsystems)
	if err_ts != nil {
		log.Panicf("Error fetch data from , Reason: %v\n", err_ts)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"massege": "Something went wrong",
		})
		return
	}

	tmsubsystem := make(map[string][]string)

	for _, tms := range tmsubsystems {
		tmsubsystem[tms.Tmname] = append(tmsubsystem[tms.Tmname], tms.Subsystemname)
	}

	var reportddailytables []models.ReportdDailyTable
	var report models.ReportdDailyTable
	for _, pro := range countdaily {
		report.Date = strings.Split(pro.CountDate, " ")[0]
		if len(tmsubsystem[pro.TmName]) > 1 {
			report.Subsystem = tmsubsystem[pro.TmName][0]
			for i := 1; i < len(tmsubsystem[pro.TmName]); i++ {
				report.Subsystem = report.Subsystem + ", " + tmsubsystem[pro.TmName][i]
			}
		} else {
			report.Subsystem = tmsubsystem[pro.TmName][0]
		}
		report.Name = pro.TmName
		report.Freq = pro.Freq
		report.AnomalyPoints = pro.CountamountPerDay
		reportddailytables = append(reportddailytables, report)
	}

	c.JSON(http.StatusOK, gin.H{
		"reportdailytable": reportddailytables,
	})
	return
}
