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
("name", "attack_bonus", "defense_bonus", "critical_chance_percent_bonus", "evasion_percent_bonus", "base_price") 
VALUES

('Rusty Sword', 3, NULL, 1, NULL, 50),
('Iron Sword', 5, NULL, 2, NULL, 100),
('Steel Sword', 7, NULL, 3, NULL, 200),
('Knight''s Blade', 10, NULL, 5, NULL, 350),
('Shadow Dagger', 4, NULL, 3, 2, 120),
('Berserker Axe', 12, NULL, 7, -2, 400),

('Leather Armor', NULL, 3, NULL, 2, 80),
('Chainmail Armor', NULL, 5, NULL, 1, 150),
('Iron Plate', NULL, 7, NULL, -1, 250),
('Paladin''s Armor', NULL, 10, NULL, -3, 400),
('Shadow Cloak', NULL, 4, NULL, 5, 200),

('Lucky Amulet', NULL, NULL, 5, NULL, 100),
('Agile Boots', NULL, NULL, NULL, 5, 120),
('Ring of Protection', NULL, 2, NULL, NULL, 90),
('Warrior''s Belt', 2, 2, NULL, NULL, 140),
('Hunter''s Gloves', NULL, NULL, 3, 3, 160)
