package models

type AnalysisInfoTheosAutoM1 struct {
	Id 		int 	`json:"id"`
	TmName 	string 	`json:"tm_name"`
	Freq 	string 	`json:"freq"`
	AnomalyResultTable string `json:"anomaly_result_table"`
}


type TmDetail struct {
	TmName 		string 			`json:"tm"`
	FreqDates 	[]FreqDate		`json:"freq_dates"`
}

type FreqDate map[string]string

type TmFreq map[string]interface{}

// type SubsystemFilter map[string]interface{}

type SubsystemCount struct {
	Subsystemname string 
	Subsystemcount int 
}

type Tth1Tmprogmodel struct {
	Id int `json:"id"`
	Tmname string `json:"tmname"`
	Subsystemname string `json:"subsystemname"`
}

type JoinAnalysisAnomalyTB struct {
	Id    int     `json:"id"`
	UTC   string  `json:"utc"`
	Avg   float32 `json:"avg"`
	Std   float32 `json:"std"`
	Count float32 `json:"count"`
	Min   float32 `json:"min"`
	Max   float32 `json:"max"`
	Q1    float32 `json:"q1"`
	Q2    float32 `json:"q2"`
	Q3    float32 `json:"q3"`
	// Skew 			float32		`json:"skew"`
	LostState          bool    `json:"lost_state"`
	EpochTen           uint32  `json:"epoch_ten"`
	AnomalyStateAutoM1 float32 `json:"anomaly_state_auto_m1"`
}

type StatisticsTHEOSSlice struct {
    Utc_tm   []string       `json:"tm_utc"`
    Avg_tm   []float32      `json:"tm_avg"`
    Std_tm   []float32      `json:"tm_std"`
	Count_tm []float32 		`json:"tm_count"`
    Min_tm   []float32      `json:"tm_min"`
    Max_tm   []float32      `json:"tm_max"`
    Q1_tm    []float32      `json:"tm_q1"`
    Q2_tm    []float32      `json:"tm_q2"`
    Q3_tm    []float32      `json:"tm_q3"`
    Utc_ano1 []string       `json:"tm_utc_ano1"`
    Ano1     []float32      `json:"tm_ano1"`
    Utc_ano2 []string       `json:"tm_utc_ano2"`
    Ano2     []float32      `json:"tm_ano2"`
    Utc_ano3 []string       `json:"tm_utc_ano3"`
    Ano3     []float32      `json:"tm_ano3"`
    Ano_bar  []VerticalLine `json:"bar_ano"`
}

type AnaAnoCSVdownload struct {
	UTC                string `csv:"utc"`
	EpochTen           string `csv:"epoch_ten"`
	Avg                string `csv:"avg"`
	Std                string `csv:"std"`
	Count              string `csv:"count"`
	Min                string `csv:"min"`
	Max                string `csv:"max"`
	Q1                 string `csv:"q1"`
	Q2                 string `csv:"q2"`
	Q3                 string `csv:"q3"`
	AnomalyStateAutoM1 string `csv:"anomaly_state_auto_m1"`
}