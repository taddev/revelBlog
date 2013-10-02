package models

//import

type View struct {
	TotalRows uint64 `json:"total_rows"`
	Offset    uint64
	Rows      []Row
}

type Row struct {
	Id    string //`json:"id"`
	Key   []int  //`json:"key"`
	Value string //`json:"value"`
}
