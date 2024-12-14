package model

type RAM struct {
	ID             int32  `json:"id_ram"`
	Name           string `json:"name"`
	ModuleCapacity int16  `json:"module_capacity"`
	ModulesNumber  int16  `json:"modules_number"`
	TotalCapacity  int16  `json:"total_capacity"`
	Frequency      int16  `json:"frequency"`
	ImageID        int32  `json:"image_id"`
	ManufacturerID int32  `json:"manufacturer_id"`
	TypeID         int32  `json:"type_id"`
}

type RAMTiny struct {
	ID   int32  `json:"id_ram"`
	Name string `json:"name"`
}
