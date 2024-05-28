package models


type AnalysisInfoTheos struct {
	Id     int    `json:"id"`
	TmName string `json:"tm_name"`
	Freq   string `json:"freq"`
	
}

type FrequencyInfo struct {
	Freq string `json:"freq"`
}

type AnaAnoTableInfo struct {
	FeatureTable       string `json:"feature_table"`
	AnomalyResultTable string `json:"anomaly_result_table"`
}

type JoinAnalyAnomalyTB struct {
	Id    int     `json:"id"`
	UTC   string     `json:"utc"`
	// UTC   time.Time     `json:"utc"`
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

type AnomalyTb struct {
	Id            int         `json:"id"`
	Name          string      `json:"name"`
	UTC           string      `json:"utc"`
	EpochTen      uint32      `json:"epoch_ten"`
	Avg           float32     `json:"avg"`
	AnomalyStatus float32     `json:"anomaly_state_auto_m1"`
	AnalysisTb    *AnalysisTb `pg:"rel:has-one"`
}

type AnalysisTb struct {
	Id int `json:"id"`
	// AnomalyState	float32		`json:"anomaly_state_auto_m1"`
	UTC        string      `json:"utc"`
	Avg        float32     `json:"avg"`
	Std        float32     `json:"std"`
	Count      float32     `json:"count"`
	Min        float32     `json:"min"`
	Max        float32     `json:"max"`
	Q1         float32     `json:"q1"`
	Q2         float32     `json:"q2"`
	Q3         float32     `json:"q3"`
	Skew       float32     `json:"skew"`
	LostState  float32     `json:"lost_state"`
	EpochTen   uint32      `json:"epoch_ten"`
	Name       string      `json:"name"`
	AnomalyTbs []AnomalyTb `pg:"rel:has-many"`
}

type AnomalythTb struct {
	TableName     struct{}        `sql:"anomalythtbs"`
	Id            int             `json:"id"`
	Name          string          `json:"name"`
	UTC           string          `json:"utc"`
	EpochTen      uint32          `json:"epoch_ten"`
	Avg           float32         `json:"avg"`
	AnomalyStatus float32         `json:"anomaly_state_auto_m1"`
	AnalysisthTb  []*AnalysisthTb `pg:",many2many:anomalythtbs_analysisthtbs"`
}

type AnalysisthTb struct {
	TableName    struct{}       `sql:"analysisthtbs"`
	Id           int            `json:"id"`
	UTC          string         `json:"utc"`
	Avg          float32        `json:"avg"`
	Std          float32        `json:"std"`
	Count        float32        `json:"count"`
	Min          float32        `json:"min"`
	Max          float32        `json:"max"`
	Q1           float32        `json:"q1"`
	Q2           float32        `json:"q2"`
	Q3           float32        `json:"q3"`
	Skew         float32        `json:"skew"`
	LostState    float32        `json:"lost_state"`
	EpochTen     uint32         `json:"epoch_ten"`
	Name         string         `json:"name"`
	AnomalythTbs []*AnomalythTb `pg:",many2many:anomalythtbs_analysisthtbs"`
}
