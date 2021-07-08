package web

//CurrentPriceResponse is a pre-defined json struct
type CurrentPriceResponse struct {
	Time  string `json:"time"`
	Price string `json:"price"`
}
