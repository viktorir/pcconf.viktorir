INSERT INTO public.pc_configuration (
    name,
    description,
    price,
    cpu_id,
    motherboard_id,
    ram_id,
    cooling_id,
    card_id,
    psu_id,
    storage_id,
    case_id
)
VALUES
($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING id_config;