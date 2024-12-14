SELECT
    pc.id_config,
    pc.name,
    pc.description,
    pc.price,
    pc.cpu_id,
    cpu.name AS cpu_name,
    pc.motherboard_id,
    motherboard.name AS motherboard_name,
    pc.ram_id,
    ram.name AS ram_name,
    pc.cooling_id,
    cooling.name AS cooling_name,
    pc.card_id,
    card.name AS card_name,
    pc.psu_id,
    psu.name AS psu_name,
    pc.storage_id,
    storage.name AS storage_name,
    pc.case_id,
    cases.name AS case_name
FROM public.pc_configuration pc
         LEFT JOIN public.cpus AS cpu ON pc.cpu_id = cpu.id_cpu
         LEFT JOIN public.motherboards AS motherboard ON pc.motherboard_id = motherboard.id_motherboard
         LEFT JOIN public.rams AS ram ON pc.ram_id = ram.id_ram
         LEFT JOIN public.coolings AS cooling ON pc.cooling_id = cooling.id_cooling
         LEFT JOIN public.graphics_cards AS card ON pc.card_id = card.id_card
         LEFT JOIN public.psus AS psu ON pc.psu_id = psu.id_psu
         LEFT JOIN public.storages AS storage ON pc.storage_id = storage.id_storage
         LEFT JOIN public.cases AS cases ON pc.case_id = cases.id_case
ORDER BY id_config DESC;