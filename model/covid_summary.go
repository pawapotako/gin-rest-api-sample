package model

type CovidSummaryRequest struct {
	ConfirmDate    *string `json:"confirmDate"`
	No             *int    `json:"no"`
	Age            *int    `json:"age"`
	Gender         *string `json:"gender"`
	GenderEn       *string `json:"genderEn"`
	Nation         *string `json:"nation"`
	NationEn       *string `json:"nationEn"`
	Province       *string `json:"province"`
	ProvinceID     *int    `json:"provinceId"`
	District       *string `json:"district"`
	ProvinceEn     *string `json:"provinceEn"`
	StatQuarantine *int    `json:"statQuarantine"`
}

type CovidSummaryResponse struct {
	Province map[string]int `json:"province"`
	AgeGroup map[string]int `json:"ageGroup"`
}
