package models

type CountAnomalyweekly struct {
	Date string `json:"weekly"`
	Lastweekday string `json:"lastweekday"`
	Count string `json:"count"`
}

type CountWeekly struct {
	Week string `json:"week"`
	SatTH string `json:"theos"`
	SatTHMain string `json:"theosmain"`
	SatTHSmall string `json:"theossmall"`
} 

type MaxMinDate struct {
	MaxDate string `json:"max_date"`
	MinDate string `json:"min_date"`
}

type Th1Dailyfilter struct {
	Tmname string `json:"tmname"`
	Subsystemname string `json:"subsystemname"`
}

type CountInfoTheosAutoM1 struct {
	Id int `json:"id"`
	TmName string `json:"tm_name"`
	Freq string `json:"freq"`
	AnomalyTable string `json:"anomaly_table"`
	FeatureTable string `json:"feature_table"`
	RecordTable string `json:"record_table"`
	AnalysisParamsId int `json:"analysis_params_id"`
	AnalysisInfoId int `json:"analysis_info_id"`
}

type LastDatecount struct {
	Lastdate string `json:"string"`
}

type Th1Grouptmsubsystem struct {
	Tmname string `json:"tmname"`
	Subsystemname string `json:"subsystemname"`
	Count int `json:"count"`
}

type CountDaily struct {
	// Id int `json:"id"`
	// CountInfoId int `json:"count_info_id"`
	// AnomalyTable string `json:"anomaly_table"`
	CountDate string `json:"count_date"`
	CountamountPerDay int `json:"countamount_per_day"`
	TmName string `json:"tm_name"` 
	Freq string `json:"freq"`
 }

 type ReportdDailyTable struct {
	Date string `json:"date"`
	Subsystem string `json:"subsystem"`
	Name string `json:"name"`
	Freq string `json:"freq"`
	AnomalyPoints int `json:"anomaly_points"`
 }

 type PostReportDaily struct {
	Satname string `json:"satname"`
	Tmnames []string `json:"tmnames"`
	Optiondate string `json:"optiondate"`
	Dates []string `json:"dates"`

 }