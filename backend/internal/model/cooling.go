package model

type Cooling struct {
	ID               int32  `json:"id_cooling"`
	Name             string `json:"name"`
	PowerDissipation int16  `json:"power_dissipation"`
	ImageID          int32  `json:"image_id"`
	ManufacturerID   int32  `json:"manufacturer_id"`
	TypeID           int32  `json:"type_id"`
}

type CoolingTiny struct {
	ID   int32  `json:"id_ram"`
	Name string `json:"name"`
}
