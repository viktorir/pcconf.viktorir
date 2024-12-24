UPDATE public.pc_configuration
SET
    name = $1,
    description = $2,
    price = $3,
    cpu_id = $4,
    motherboard_id = $5,
    ram_id = $6,
    cooling_id = $7,
    card_id = $8,
    psu_id = $9,
    storage_id = $10,
    case_id = $11
WHERE id_config = $12;