TRUNCATE table enemies cascade;
INSERT INTO "enemies" 
("name", "level", "hp", "min_attack", "max_attack", "defense", "critical_chance_percent", "evasion_chance_percent", "boss_type") 
VALUES

-- 🌲 Лес (1–2)
('goblin',         1, 40, 3, 5, 3, 5, 3, null),
('wolf',           1, 43, 4, 6, 3, 4, 5, null),
('forest_spider',  1, 41, 3, 5, 2, 6, 7, null),
('bandit',         2, 45, 4, 6, 3, 5, 4, null),

-- ⚰️ Кладбище (3–4)
('skeleton',       3, 60, 5, 7, 5, 6, 7, null),
('orc',            3, 63, 6, 8, 5, 6, 2, null),
('ghost',          3, 61, 5, 7, 4, 9, 9, null),
('zombie',         4, 65, 6, 8, 5, 2, 1, null),

-- 🏰 Замок (5–6)
('knight',         5, 65, 5, 7, 5, 7, 5, null),
('castle_guard',   5, 68, 6, 8, 5, 5, 6, null),
('vampire',        6, 70, 6, 8, 5, 8, 8, null),
('troll',          6, 67, 5, 7, 6, 3, 10, null);


TRUNCATE table items cascade;
INSERT INTO "items" 
("name", "category", "equips_on", "attack_bonus", "defense_bonus", "critical_chance_percent_bonus", "evasion_percent_bonus", "base_price") 
VALUES

('rusty_sword', 'weapon', 'hand', 3, NULL, 1, NULL, 50),
('leather_armor', 'armor', 'body', NULL, 3, NULL, 2, 80);

-- Potions --
-- INSERT INTO "items" 
-- ("name", "category", "potion_effect_type", "potion_effect_value", "potion_effect_duration", "base_price") 
-- VALUES
-- ('health_potion', 'potion', 'heal', 30, NULL, 50),
-- ('greater_health_potion', 'potion', 'heal', 60, NULL, 100);
-- ('strength_potion', 'potion', 'attack_buff', 5, 3, 80),
-- ('defense_potion', 'potion', 'defense_buff', 5, 3, 80),
-- ('evasion_potion', 'potion', 'evasion_buff', 10, 2, 70),
-- ('invisibility_potion', 'potion', 'evasion_buff', 50, 1, 200);
