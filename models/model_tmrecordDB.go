package models

type Th1TmnameOld struct {
	TmStructOld
}
type Th1Tmname struct {
	TmStruct
}

type TmStruct struct {
	Id             string `json:"id"`
	Tmname         string `json:"tmname"`
	Property       string `json:"property"`
	Description    string `json:"description"`
	Tmsubsystem_id string `json:"tmsubsystem_id"`
	Tmoperation_id string `json:"tmoperation_id"`
	// Remark         string `json:"remark"`
}
type TmStructOld struct {
	Id             string `json:"id"`
	Tmname         string `json:"tmname"`
	Property       string `json:"property"`
	Description    string `json:"description"`
	Tmsubsystem_id string `json:"tmsubsystem_id"`
	Tmoperation_id string `json:"tmoperation_id"`
}

// type TmIdName struct {
// 	Id   string `json:"TmId"`
// 	TmName string `json:"tmName"`
// }

type Th1Tmprogmodel struct {
	Id				uint16 `json:"id"`
	Tmname         	string `json:"tmname"`
	Subsystemname 	string `json:"subsystemname"`
}

type JoinOperationSubsystem struct {
	Property		string	`json:"property"`
	Description		string	`json:"description"`
	Operationname	string 	`json:"operationname"`
	Subsystemname	string 	`json:"subsystemname"`
}