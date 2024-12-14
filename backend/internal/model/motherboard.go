package model

type Motherboard struct {
	ID             int32  `json:"id_motherboard"`
	Name           string `json:"name"`
	RamSlots       int16  `json:"ram_slots"`
	M2Slots        int16  `json:"m2_connectors"`
	SataSlots      int16  `json:"sata_connectors"`
	ImageID        int32  `json:"image_id"`
	ManufacturerID int32  `json:"manufacturer_id"`
	SocketId       int32  `json:"socket_id"`
	FormfactorID   int32  `json:"formfactor_id"`
	ChipsetID      int32  `json:"chipset_id"`
	TypeID         int32  `json:"type_id"`
}

type MotherboardTiny struct {
	ID   int32  `json:"id_motherboard"`
	Name string `json:"name"`
}
