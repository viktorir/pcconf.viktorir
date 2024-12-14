package model

type Case struct {
	ID             int32  `json:"id_case"`
	Name           string `json:"name"`
	ImageID        int32  `json:"image_id"`
	ManufacturerID int32  `json:"manufacturer_id"`
}

type CaseTiny struct {
	ID   int32  `json:"id_case"`
	Name string `json:"name"`
}
