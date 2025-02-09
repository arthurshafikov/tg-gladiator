INSERT INTO "enemies" 
("name", "hp", "min_attack", "max_attack", "defense", "critical_chance_percent", "evasion_chance_percent", "gold_reward_min", "gold_reward_max") 
VALUES

('Goblin', 50, 3, 6, 2, 5, 5, 5, 10),
('Orc', 80, 5, 9, 4, 10, 3, 10, 20),
('Skeleton', 60, 4, 7, 3, 7, 10, 7, 15),
('Bandit', 70, 6, 10, 3, 15, 5, 12, 18),
('Troll', 100, 8, 12, 6, 12, 3, 20, 30),
;

TRUNCATE table items cascade;
INSERT INTO "items" 
("name", "category", "equips_on", "attack_bonus", "defense_bonus", "critical_chance_percent_bonus", "evasion_percent_bonus", "base_price") 
VALUES

('rusty_sword', 'weapon', 'hand', 3, NULL, 1, NULL, 50),
('iron_sword', 'weapon', 'hand', 5, NULL, 2, NULL, 100),
('steel_sword', 'weapon', 'hand', 7, NULL, 3, NULL, 200),
('knights_blade', 'weapon', 'hand', 10, NULL, 5, NULL, 350),
('shadow_dagger', 'weapon', 'hand', 4, NULL, 3, 2, 120),
('berserker_axe', 'weapon', 'two_hands', 12, NULL, 7, -2, 400),
('greatsword', 'weapon', 'two_hands', 15, NULL, 8, -3, 500),
('warhammer', 'weapon', 'two_hands', 18, NULL, 6, -4, 550),
('twin_daggers', 'weapon', 'hand', 6, NULL, 5, 3, 250),
('firebrand', 'weapon', 'hand', 9, NULL, 4, NULL, 300),
('shield_of_valor', 'weapon', 'hand', NULL, 6, NULL, -1, 250),
('wardens_shield', 'weapon', 'hand', NULL, 8, NULL, -2, 350),

('leather_armor', 'armor', 'body', NULL, 3, NULL, 2, 80),
('chainmail_armor', 'armor', 'body', NULL, 5, NULL, 1, 150),
('iron_plate', 'armor', 'body', NULL, 7, NULL, -1, 250),
('paladins_armor', 'armor', 'body', NULL, 10, NULL, -3, 400),
('shadow_cloak', 'armor', 'body', NULL, 4, NULL, 5, 200),
('dragon_scale_armor', 'armor', 'body', NULL, 12, NULL, -5, 600),
('reinforced_vest', 'armor', 'body', NULL, 6, NULL, 1, 180),
('iron_helmet', 'armor', 'head', NULL, 3, NULL, -1, 150),
('knights_helmet', 'armor', 'head', NULL, 5, NULL, -2, 300),
('shadow_hood', 'armor', 'head', NULL, 2, NULL, 3, 180),
('crown_of_thorns', 'armor', 'head', NULL, 4, 3, -1, 250),
('battle_mask', 'armor', 'head', NULL, 3, NULL, 2, 200),
('hunters_gloves', 'armor', 'hands', NULL, NULL, 3, 3, 160),
('titan_gauntlets', 'accessory', 'hands', 3, 3, NULL, -2, 180),
('agile_boots', 'armor', 'feet', NULL, NULL, NULL, 5, 120),
('boots_of_speed', 'armor', 'feet', NULL, NULL, NULL, 6, 200),
('elven_boots', 'armor', 'feet', NULL, NULL, NULL, 7, 280),
('dark_cloak', 'armor', 'body', NULL, 5, NULL, 5, 320),
('phantom_gloves', 'armor', 'hands', NULL, NULL, 5, 5, 270),

('lucky_amulet', 'accessory', 'neck', NULL, NULL, 5, NULL, 100),
('ring_of_protection', 'accessory', 'finger', NULL, 2, NULL, NULL, 90),
('ring_of_agility', 'accessory', 'finger', NULL, NULL, NULL, 4, 110),
('ring_of_strength', 'accessory', 'finger', 2, NULL, NULL, NULL, 130),
('amulet_of_power', 'accessory', 'neck', 3, NULL, 4, NULL, 220),
('ancient_talisman', 'accessory', 'neck', NULL, NULL, 7, 2, 300),
('blood_ring', 'accessory', 'finger', 4, NULL, NULL, -2, 150)
;
