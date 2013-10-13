package models

//import

type ViewUser struct {
	TotalRows uint64 `json:"total_rows"`
	Offset    uint64
	Rows      []ViewUserRow	
}

type ViewUserRow struct {
	Id    string
	Key   string 
	Value []string 
}

type User struct {
	Id          string `json:"_id"`
	Rev         string `json:"_rev"`
	Username    string 
	DisplayName string 
	Password    string 
}
