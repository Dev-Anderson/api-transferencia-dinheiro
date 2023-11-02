package schemas

type Transfer struct {
	IDOrigin  int     `json:"idorigin"`
	IDDestiny int     `json:"iddestiny"`
	Value     float64 `json:"value"`
}
