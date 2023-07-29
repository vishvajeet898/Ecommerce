package models

type Addresses struct {
	Address_ID     string `json:"address_ID"`
	User_ID        string `json:"user_ID"`
	Address_line_1 string `json:"address_Line_1"`
	City           string `json:"city"`
	State          string `json:"state"`
	Country        string `json:"country"`
	Is_default     bool   `json:"is_default"`
}
