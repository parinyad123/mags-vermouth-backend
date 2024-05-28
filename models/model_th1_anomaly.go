package models

// import (
	// "time"
	// "github.com/golang/protobuf/ptypes/timestamp"
// )


type DataSlice struct {
	Utc_tm   []string       `json:"tm_utc"`
	// Utc_tm   []time.Time         `json:"tm_utc"`
	Avg_tm   []float32      `json:"tm_avg"`
	Std_tm   []float32      `json:"tm_std"`
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

type VerticalLine struct {
	Tyte      string     `json:"type"`
	X0        string     `json:"x0"`
	Y0        uint8      `json:"y0"`
	Xref      string     `json:"xref"`
	Yref      string     `json:"yref"`
	X1        string     `json:"x1"`
	Y1        uint8      `json:"y1"`
	Fillcolor string     `json:"fillcolor"`
	Opacity   float32    `json:"opacity"`
	Layer     string     `json:"layer"`
	Line      LineStruct `json:"line"`
}

type LineStruct struct {
	// Color string `json:"color"`
	Width float32 `json:"width"`
}

type CSVdownload struct {
	SatName    string `json:"sat_name"`
	TmName     string `json:"tm_name"`
	Freq       string `json:"freq"`
	AnalysisTb string `json:"analysis_table"`
	AnomalyTb  string `json:"anomaly_table"`
	StartUtc   string `json:"start_utc"`
	EndUtc     string `json:"end_utc"`
}