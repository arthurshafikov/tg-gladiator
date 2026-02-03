TRUNCATE table enemies cascade;
INSERT INTO "enemies" 
("name", "level", "hp", "min_attack", "max_attack", "defense", "critical_chance_percent", "evasion_chance_percent", "boss_type") 
VALUES

-- 🌲 Лес (1–2)
('goblin',         1, 40, 3, 5, 3, 5, 3, null),
('wolf',           1, 43, 4, 6, 3, 4, 5, null),
('forest_spider',  1, 41, 3, 5, 2, 6, 7, null),
('forest_giant',   2, 48, 4, 6, 3, 5, 4, null),
('woodland_sprite',2, 46, 4, 6, 3, 5, 5, null),
('forest_goblin',  2, 47, 4, 6, 3, 4, 4, null),

-- ⚰️ Кладбище (3–4)
('skeleton',       3, 60, 5, 7, 5, 6, 7, null),
('orc',            3, 63, 6, 8, 5, 6, 2, null),
('ghost',          3, 61, 5, 7, 4, 9, 9, null),
('skeleton_warrior',4, 68, 6, 8, 5, 5, 4, null),
('ghost_knight',   4, 70, 6, 8, 5, 6, 3, null),
('undead_knight',  4, 69, 6, 8, 6, 4, 3, null),

-- 🏰 Замок (5–6)
('knight',         5, 65, 5, 7, 5, 7, 5, null),
('castle_guard',   5, 68, 6, 8, 5, 5, 6, null),
('castle_wizard',  5, 72, 7, 9, 5, 6, 4, null),
('dark_knight',    6, 74, 6, 8, 5, 6, 5, null),
('castle_sorcerer',6, 78, 7, 9, 6, 7, 4, null),
('royal_guard',    6, 75, 7, 9, 6, 5, 5, null),

-- 🔥 Драконьи горы (7–8)
('young_dragon',   7, 75, 7, 9, 6, 8, 3, null),
('fire_elemental', 7, 78, 8, 10, 6, 7, 4, null),
('mountain_ogre',  7, 76, 7, 9, 7, 5, 2, null),
('fire_golem',     8, 85, 9, 11, 7, 7, 3, null),
('mountain_dragon',8, 88, 9, 11, 7, 8, 4, null),
('lava_titan',     8, 84, 8, 10, 7, 6, 1, null),

-- 🌑 Тёмный лес (9–10)
('dark_elf',       9, 85, 9, 11, 7, 10, 10, null),
('werewolf',       9, 88, 10, 12, 7, 9, 8, null),
('shadow_beast',   9, 86, 9, 11, 6, 12, 12, null),
('dark_knight',    10, 94, 11, 13, 8, 10, 6, null),
('shadow_assassin',10, 96, 11, 13, 8, 9, 5, null),
('night_stalker',  10, 93, 10, 12, 7, 8, 6, null),

-- ❄️ Ледяные пещеры (11–12)
('ice_giant',      11, 95, 11, 13, 8, 7, 2, null),
('frost_troll',    11, 98, 12, 14, 8, 6, 3, null),
('snow_yeti',      11, 96, 11, 13, 9, 5, 1, null),
('blizzard_wolf',  12, 105, 13, 15, 9, 8, 4, null),
('ice_wraith',     12, 102, 12, 14, 8, 9, 5, null),
('frozen_guardian',12, 103, 12, 14, 8, 7, 3, null),

-- 🌋 Вулканические равнины (13–14)
('ash_dragon',     13, 105, 13, 15, 9, 9, 3, null),
('magma_golem',    13, 108, 14, 16, 9, 7, 2, null),
('fire_phoenix',   13, 106, 13, 15, 8, 11, 11, null),
('volcano_titan',  14, 115, 15, 17, 10, 8, 1, null),
('magma_dragon',   14, 118, 15, 17, 10, 9, 2, null),
('lava_beast',     14, 112, 14, 16, 9, 8, 3, null),

-- 🏜️ Пустынные дюны (15–16)
('sand_elemental', 15, 115, 15, 17, 10, 10, 5, null),
('desert_scorpion',15, 118, 16, 18, 10, 9, 7, null),
('mirage_warrior', 15, 116, 15, 17, 9, 12, 9, null),
('dune_behemoth',  16, 125, 17, 19, 11, 7, 2, null),
('sand_golem',     16, 128, 17, 19, 11, 8, 3, null),
('desert_serpent', 16, 122, 16, 18, 10, 10, 4, null),

-- 🌊 Океанские глубины (17–18)
('sea_serpent',    17, 125, 17, 19, 11, 11, 6, null),
('kraken_tentacle',17, 128, 18, 20, 11, 8, 4, null),
('abyssal_horror', 17, 126, 17, 19, 10, 13, 10, null),
('leviathan',      18, 135, 19, 21, 12, 9, 3, null),
('ocean_guardian', 18, 138, 19, 21, 12, 10, 4, null),
('deep_dweller',   18, 132, 18, 20, 11, 11, 5, null),

-- ⚡ Штормовые вершины (19–20)
('storm_giant',    19, 135, 19, 21, 12, 12, 5, null),
('lightning_drake',19, 138, 20, 22, 12, 10, 7, null),
('tempest_spirit', 19, 136, 19, 21, 11, 14, 11, null),
('thunder_beast',  20, 145, 21, 23, 13, 11, 4, null),
('storm_elemental',20, 148, 21, 23, 13, 12, 6, null),
('celestial_titan',20, 142, 20, 22, 12, 10, 5, null);


TRUNCATE table items cascade;
INSERT INTO "items" 
("name", "category", "equips_on", "attack_bonus", "defense_bonus", "critical_chance_percent_bonus", "evasion_percent_bonus", "base_price") 
VALUES

('rusty_sword', 'weapon', 'hand', 1, NULL, 1, NULL, 50),
('leather_armor', 'armor', 'body', NULL, 2, NULL, 1, 80),
('iron_sword', 'weapon', 'hand', 2, NULL, 2, NULL, 100),
('chain_armor', 'armor', 'body', NULL, 3, NULL, 1, 120),
('steel_sword', 'weapon', 'hand', 3, NULL, 3, NULL, 150),
('mithril_armor', 'armor', 'body', NULL, 4, NULL, 2, 180),
('adamantine_sword', 'weapon', 'hand', 4, NULL, 4, NULL, 300),
('dragon_armor', 'armor', 'body', NULL, 5, NULL, 3, 350),
('legendary_sword', 'weapon', 'hand', 5, NULL, 3, NULL, 500),
('divine_armor', 'armor', 'body', NULL, 6, NULL, 4, 600),
('wooden_helmet', 'armor', 'head', NULL, 2, NULL, 1, 40),
('iron_helmet', 'armor', 'head', NULL, 3, NULL, 1, 90),
('steel_helmet', 'armor', 'head', NULL, 4, NULL, 2, 140),
('mithril_helmet', 'armor', 'head', NULL, 5, NULL, 2, 250),
('dragon_helmet', 'armor', 'head', NULL, 6, NULL, 3, 400),
('leather_gloves', 'armor', 'hands', NULL, 2, 1, NULL, 60),
('iron_gloves', 'armor', 'hands', NULL, 3, 1, NULL, 110),
('steel_gloves', 'armor', 'hands', NULL, 4, 1, NULL, 160),
('mithril_gloves', 'armor', 'hands', NULL, 5, 1, NULL, 300),
('dragon_gloves', 'armor', 'hands', NULL, 6, 2, NULL, 450),
('leather_boots', 'armor', 'feet', NULL, 2, NULL, 1, 70),
('iron_boots', 'armor', 'feet', NULL, 3, NULL, 1, 120),
('steel_boots', 'armor', 'feet', NULL, 4, NULL, 2, 170),
('mithril_boots', 'armor', 'feet', NULL, 5, NULL, 2, 320),
('dragon_boots', 'armor', 'feet', NULL, 6, NULL, 3, 500),

-- 🗡️ Two-handed weapons
('wooden_staff', 'weapon', 'two_hands', 2, NULL, NULL, NULL, 60),
('iron_axe', 'weapon', 'two_hands', 4, NULL, NULL, NULL, 150),
('steel_greatsword', 'weapon', 'two_hands', 6, NULL, 1, NULL, 250),
('mithril_halberd', 'weapon', 'two_hands', 8, NULL, 2, NULL, 400),
('dragon_cleaver', 'weapon', 'two_hands', 10, NULL, 3, NULL, 700),

-- 📿 Neck accessories
('leather_amulet', 'accessory', 'neck', NULL, 1, NULL, NULL, 70),
('silver_pendant', 'accessory', 'neck', NULL, 2, 1, NULL, 120),
('gold_necklace', 'accessory', 'neck', 1, 3, 2, NULL, 200),
('mithril_amulet', 'accessory', 'neck', 2, 4, 2, 1, 350),
('dragon_totem', 'accessory', 'neck', 3, 5, 3, 2, 600),

-- 💍 Rings
('copper_ring', 'accessory', 'finger', NULL, 1, NULL, 1, 40),
('silver_ring', 'accessory', 'finger', 1, 1, 1, 2, 100),
('gold_ring', 'accessory', 'finger', 2, 2, 2, 3, 180),
('sapphire_ring', 'accessory', 'finger', 3, 3, 3, 4, 300),
('dragon_ring', 'accessory', 'finger', 4, 4, 4, 5, 550);

-- Potions --
-- INSERT INTO "items" 
-- ("name", "category", "potion_effect_type", "potion_effect_value", "potion_effect_duration", "base_price") 
-- VALUES
-- ('health_potion', 'potion', 'heal', 30, NULL, 50),
-- ('strength_potion', 'potion', 'attack_buff', 3, 3, 80),
-- ('greater_health_potion', 'potion', 'heal', 60, NULL, 100),
-- ('defense_potion', 'potion', 'defense_buff', 3, 3, 120),
-- ('evasion_potion', 'potion', 'evasion_buff', 5, 3, 120),
-- ('crit_potion', 'potion', 'crit_buff', 3, 3, 120);
