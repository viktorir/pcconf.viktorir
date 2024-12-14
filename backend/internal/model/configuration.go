package model

import (
	"pcconf.viktorir/internal/database"
)

type PCConfigurationCard struct {
	ID          int32            `json:"id_config"`
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

func GetFullConfigs() (configs []PCConfigurationCard, err error) {
	query, err := database.ReadSQLFile("internal/model/queries/get_full_pc_configurations.sql")
	if err != nil {
		return nil, err
	}
	rows, err := database.Sql.Query(query)
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
	query, err := database.ReadSQLFile("internal/model/queries/get_pc_configuration.sql")
	if err != nil {
		return
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
		return
	}
	return
}
