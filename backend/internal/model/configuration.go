package model

import (
	"pcconf.viktorir/internal/database"
)

type PCConfiguration struct {
	ID            int    `json:"id_config"`
	Name          string `json:"name"`
	Description   string `json:"description"`
	Price         string `json:"price"`
	CPUID         int    `json:"cpu_id"`
	MotherboardID int    `json:"motherboard_id"`
	RAMID         int    `json:"ram_id"`
	CoolingID     int    `json:"cooling_id"`
	CardID        int    `json:"card_id"`
	PSUID         int    `json:"psu_id"`
	StorageID     int    `json:"storage_id"`
	CaseID        int    `json:"case_id"`
}

type PCConfigurationCard struct {
	ID          int              `json:"id_config"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Price       string           `json:"price"`
	CPU         CPUTiny          `json:"cpu"`
	Motherboard MotherboardTiny  `json:"motherboard"`
	RAM         RAMTiny          `json:"ram"`
	Cooling     CoolingTiny      `json:"cooling"`
	Card        GraphicsCardTiny `json:"card"`
	PSU         PSUTiny          `json:"psu"`
	Storage     StorageTiny      `json:"storage"`
	Case        CaseTiny         `json:"case"`
}

func GetConfigs(page, limit int) (configs []PCConfigurationCard, err error) {
	query, err := database.ReadSQLFile("internal/model/queries/pc_configuration/get_all.sql")
	if err != nil {
		return nil, err
	}
	rows, err := database.Sql.Query(query, limit, (page-1)*limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var config PCConfigurationCard
		err := rows.Scan(
			&config.ID,
			&config.Name,
			&config.Description,
			&config.Price,
			&config.CPU.ID,
			&config.CPU.Name,
			&config.Motherboard.ID,
			&config.Motherboard.Name,
			&config.RAM.ID,
			&config.RAM.Name,
			&config.Cooling.ID,
			&config.Cooling.Name,
			&config.Card.ID,
			&config.Card.Name,
			&config.PSU.ID,
			&config.PSU.Name,
			&config.Storage.ID,
			&config.Storage.Name,
			&config.Case.ID,
			&config.Case.Name,
		)
		if err != nil {
			return nil, err
		}
		configs = append(configs, config)
	}

	return configs, nil
}

func GetConfig(id int) (config PCConfigurationCard, err error) {
	query, err := database.ReadSQLFile("internal/model/queries/pc_configuration/get.sql")
	if err != nil {
		return PCConfigurationCard{}, err
	}

	err = database.Sql.QueryRow(query, id).Scan(
		&config.ID,
		&config.Name,
		&config.Description,
		&config.Price,
		&config.CPU.ID,
		&config.CPU.Name,
		&config.Motherboard.ID,
		&config.Motherboard.Name,
		&config.RAM.ID,
		&config.RAM.Name,
		&config.Cooling.ID,
		&config.Cooling.Name,
		&config.Card.ID,
		&config.Card.Name,
		&config.PSU.ID,
		&config.PSU.Name,
		&config.Storage.ID,
		&config.Storage.Name,
		&config.Case.ID,
		&config.Case.Name,
	)
	if err != nil {
		return PCConfigurationCard{}, err
	}

	return config, nil
}

func CreateConfig(config PCConfiguration) (int, error) {
	query, err := database.ReadSQLFile("internal/model/queries/pc_configuration/create.sql")
	if err != nil {
		return 0, err
	}
	var newID int
	err = database.Sql.QueryRow(
		query,
		config.Name,
		config.Description,
		config.Price,
		config.CPUID,
		config.MotherboardID,
		config.RAMID,
		config.CoolingID,
		config.CardID,
		config.PSUID,
		config.StorageID,
		config.CaseID,
	).Scan(&newID)
	if err != nil {
		return 0, err
	}
	return newID, nil
}

func UpdateConfig(config PCConfiguration) error {
	query, err := database.ReadSQLFile("internal/model/queries/pc_configuration/update.sql")
	if err != nil {
		return err
	}
	_, err = database.Sql.Exec(
		query,
		config.Name,
		config.Description,
		config.Price,
		config.CPUID,
		config.MotherboardID,
		config.RAMID,
		config.CoolingID,
		config.CardID,
		config.PSUID,
		config.StorageID,
		config.CaseID,
		config.ID,
	)
	return err
}

func DeleteConfig(id int) error {
	query, err := database.ReadSQLFile("internal/model/queries/pc_configuration/delete.sql")
	if err != nil {
		return err
	}
	_, err = database.Sql.Exec(query, id)
	return err
}
