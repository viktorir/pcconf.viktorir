package model

type CPU struct {
	ID             int32  `json:"id_cpu"`
	Name           string `json:"name"`
	Cores          int16  `json:"number_cores"`
	Threads        int16  `json:"number_threads"`
	BaseFrequency  int16  `json:"base_frequency"`
	MaxFrequency   int16  `json:"max_frequency"`
	ImageID        int32  `json:"image_id"`
	ManufacturerID int32  `json:"manufacturer_id"`
	SocketID       int32  `json:"socket_id"`
}

type CPUTiny struct {
	ID   int32  `json:"id_cpu"`
	Name string `json:"name"`
}
