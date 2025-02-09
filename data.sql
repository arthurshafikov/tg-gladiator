INSERT INTO "enemies" 
("name", "hp", "min_attack", "max_attack", "defense", "critical_chance_percent", "evasion_chance_percent", "gold_reward_min", "gold_reward_max") 
VALUES

('Goblin', 50, 3, 6, 2, 5, 5, 5, 10),
('Orc', 80, 5, 9, 4, 10, 3, 10, 20),
('Skeleton', 60, 4, 7, 3, 7, 10, 7, 15),
('Bandit', 70, 6, 10, 3, 15, 5, 12, 18),
('Troll', 100, 8, 12, 6, 12, 3, 20, 30),
;

INSERT INTO "items" 
("name", "category", "equips_on", "attack_bonus", "defense_bonus", "critical_chance_percent_bonus", "evasion_percent_bonus", "base_price") 
VALUES

('Rusty Sword', 'weapon', 'hand', 3, NULL, 1, NULL, 50),
('Iron Sword', 'weapon', 'hand', 5, NULL, 2, NULL, 100),
('Steel Sword', 'weapon', 'hand', 7, NULL, 3, NULL, 200),
('Knight''s Blade', 'weapon', 'hand', 10, NULL, 5, NULL, 350),
('Shadow Dagger', 'weapon', 'hand', 4, NULL, 3, 2, 120),
('Berserker Axe', 'weapon', 'two_hands', 12, NULL, 7, -2, 400),
('Greatsword', 'weapon', 'two_hands', 15, NULL, 8, -3, 500),
('Warhammer', 'weapon', 'two_hands', 18, NULL, 6, -4, 550),
('Twin Daggers', 'weapon', 'hand', 6, NULL, 5, 3, 250),
('Firebrand', 'weapon', 'hand', 9, NULL, 4, NULL, 300),
('Shield of Valor', 'weapon', 'hand', NULL, 6, NULL, -1, 250),
('Warden''s Shield', 'weapon', 'hand', NULL, 8, NULL, -2, 350),

('Leather Armor', 'armor', 'body', NULL, 3, NULL, 2, 80),
('Chainmail Armor', 'armor', 'body', NULL, 5, NULL, 1, 150),
('Iron Plate', 'armor', 'body', NULL, 7, NULL, -1, 250),
('Paladin''s Armor', 'armor', 'body', NULL, 10, NULL, -3, 400),
('Shadow Cloak', 'armor', 'body', NULL, 4, NULL, 5, 200),
('Dragon Scale Armor', 'armor', 'body', NULL, 12, NULL, -5, 600),
('Reinforced Vest', 'armor', 'body', NULL, 6, NULL, 1, 180),
('Iron Helmet', 'armor', 'head', NULL, 3, NULL, -1, 150),
('Knight''s Helmet', 'armor', 'head', NULL, 5, NULL, -2, 300),
('Shadow Hood', 'armor', 'head', NULL, 2, NULL, 3, 180),
('Crown of Thorns', 'armor', 'head', NULL, 4, 3, -1, 250),
('Battle Mask', 'armor', 'head', NULL, 3, NULL, 2, 200),
('Hunter''s Gloves', 'armor', 'hands', NULL, NULL, 3, 3, 160),
('Titan Gauntlets', 'accessory', 'hands', 3, 3, NULL, -2, 180),
('Agile Boots', 'armor', 'feet', NULL, NULL, NULL, 5, 120),
('Boots of Speed', 'armor', 'feet', NULL, NULL, NULL, 6, 200),
('Elven Boots', 'armor', 'feet', NULL, NULL, NULL, 7, 280),
('Dark Cloak', 'armor', 'body', NULL, 5, NULL, 5, 320),
('Phantom Gloves', 'armor', 'hands', NULL, NULL, 5, 5, 270),

('Lucky Amulet', 'accessory', 'neck', NULL, NULL, 5, NULL, 100),
('Ring of Protection', 'accessory', 'finger', NULL, 2, NULL, NULL, 90),
('Ring of Agility', 'accessory', 'finger', NULL, NULL, NULL, 4, 110),
('Ring of Strength', 'accessory', 'finger', 2, NULL, NULL, NULL, 130),
('Amulet of Power', 'accessory', 'neck', 3, NULL, 4, NULL, 220),
('Ancient Talisman', 'accessory', 'neck', NULL, NULL, 7, 2, 300),
('Blood Ring', 'accessory', 'finger', 4, NULL, NULL, -2, 150)
;
