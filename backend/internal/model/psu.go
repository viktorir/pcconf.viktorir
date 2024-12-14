package model

type PSU struct {
	ID             int32  `json:"id_psu"`
	Name           string `json:"name"`
	RatedPower     int32  `json:"rated_power"`
	ImageID        int32  `json:"image_id"`
	ManufacturerID int32  `json:"manufacturer_id"`
	FormfactorID   int32  `json:"formfactor_id"`
	CertificateID  int32  `json:"certificate_id"`
}

type PSUTiny struct {
	ID   int32  `json:"id_psu"`
	Name string `json:"name"`
}
