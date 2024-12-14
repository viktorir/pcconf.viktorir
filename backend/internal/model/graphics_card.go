package model

type GraphicsCard struct {
	ID             int32  `json:"id_card"`
	Name           string `json:"name"`
	BaseFrequency  int16  `json:"base_frequency"`
	MaxFrequency   int16  `json:"max_frequency"`
	MemoryCapacity int16  `json:"memory_capacity"`
	ImageID        int32  `json:"image_id"`
	ManufacturerID int32  `json:"manufacturer_id"`
	TypeID         int32  `json:"type_id"`
}

type GraphicsCardTiny struct {
	ID   int32  `json:"id_card"`
	Name string `json:"name"`
}
