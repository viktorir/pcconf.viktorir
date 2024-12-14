package model

type Storage struct {
	ID             int32  `json:"id_storage"`
	Name           string `json:"name"`
	ImageID        int32  `json:"image_id"`
	ManufacturerID int32  `json:"manufacturer_id"`
	TypeID         int32  `json:"type_id"`
}

type StorageTiny struct {
	ID   int32  `json:"id_storage"`
	Name string `json:"name"`
}
