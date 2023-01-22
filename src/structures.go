package src

/*
** Structures
*/
type TimeCommanderData struct {
	Records []Record `json:"records"`
}

type Record struct {
	Date     string    `json:"date"`
	Sessions []Session `json:"sessions"`
}

type Session struct {
	Start int64 `json:"start"`
	End   int64 `json:"end"`
}
