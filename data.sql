TRUNCATE table enemies cascade;
INSERT INTO "enemies" 
("name", "level", "hp", "min_attack", "max_attack", "defense", "critical_chance_percent", "evasion_chance_percent", "gold_reward_min", "gold_reward_max", "xp_reward_min", "xp_reward_max", "boss_type") 
VALUES

-- 🌲 Лес (1–2)
('goblin',         1, 35, 6, 8, 1, 5, 3, 4, 8, 10, 20, null),
('wild_dog',       1, 30, 6, 9, 0, 4, 5, 5, 9, 10, 20, null),
('forest_snake',   1, 28, 5, 7, 0, 6, 7, 4, 8, 10, 20, null),
('bandit',         2, 40, 7, 10, 2, 5, 4, 5, 10, 15, 25, null),
('slow_trent',     2, 70, 4, 6, 8, 0, 0, 2, 8, 15, 25, null),
('wolf',           2, 38, 8, 10, 1, 7, 6, 6, 12, 15, 25, null),


-- ⚰️ Кладбище (3–4)
('skeleton',       3, 45, 7, 10, 2, 6, 7, 6, 14, 30, 50, null),
('orc',            3, 60, 9, 11, 3, 6, 2, 5, 16, 30, 50, null),
('haunted_soul',   3, 40, 7, 10, 1, 9, 9, 6, 15, 30, 50, null),
('zombie',         4, 65, 9, 12, 4, 3, 2, 3, 16, 40, 60, null),
('vampire',        4, 50, 8, 11, 2, 8, 5, 7, 18, 40, 60, null),
('rotting_corpse', 4, 55, 10, 12, 4, 2, 1, 4, 14, 40, 60, null),


-- 🏰 Руины (5–6)
('orc',            5, 80, 11, 13, 5, 3, 2, 5, 20, 50, 80, null),
('undead_knight',  5, 75, 10, 12, 6, 3, 2, 4, 18, 50, 80, null),
('giant_bat',      5, 60, 9, 12, 2, 7, 6, 8, 20, 50, 80, null),
('bandit_leader',  6, 70, 10, 13, 4, 5, 4, 6, 25, 60, 90, null),
('dark_sorcerer',  6, 55, 9, 13, 3, 9, 5, 9, 26, 60, 90, null),
('cursed_monk',    6, 65, 11, 13, 5, 5, 4, 5, 24, 60, 90, null),


-- 🌋 Огненные пещеры (7–8)
('red_ogre',        7, 110, 13, 17, 7, 6, 2, 15, 38, 80, 120, null),
('satan_servant',  7, 95, 12, 15, 4, 15, 10, 12, 36, 80, 120, null),
('flame_imp',      7, 85, 14, 16, 3, 10, 8, 14, 35, 80, 120, null),
('fire_knight',    8, 105, 15, 18, 6, 7, 4, 18, 42, 100, 140, null),
('fire_elemental', 8, 90, 16, 19, 5, 12, 6, 20, 45, 100, 140, null),
('volcanic_golem', 8, 130, 12, 15, 10, 3, 2, 22, 50, 100, 140, null),


-- 😈 Преисподняя (9–10)
('demon',           9, 140, 17, 21, 9, 4, 3, 22, 58, 120, 180, null),
('ancient_warlord', 9, 135, 15, 20, 8, 3, 2, 18, 55, 120, 180, null),
('infernal_hound',  9, 120, 16, 19, 5, 6, 9, 20, 52, 120, 180, null),
('dragonspawn',     10, 140, 18, 22, 10, 6, 7, 25, 65, 150, 200, null),
('chaos_champion',  10, 150, 20, 24, 10, 5, 5, 28, 75, 150, 200, null),
('hellfire_witch',  10, 110, 17, 21, 7, 10, 8, 26, 72, 150, 200, null),


-- 🧊 Ледяные пещеры (11–12)
('ice_golem',         11, 160, 16, 20, 12, 2, 1, 22, 78, 170, 220, null),
('frost_wraith',      11, 125, 15, 18, 8, 9, 10, 20, 70, 170, 220, null),
('snow_wolf',         11, 110, 14, 18, 7, 8, 7, 18, 68, 170, 220, null),
('glacier_elemental', 12, 140, 16, 21, 10, 4, 3, 25, 82, 200, 250, null),
('frozen_cultist',    12, 130, 17, 20, 9, 11, 8, 22, 80, 200, 250, null),
('yeti',              12, 160, 18, 23, 12, 2, 2, 24, 85, 200, 250, null),


-- 🏜️ Пустыня (13-14)
('sand_wyrm', 13, 180, 16, 24, 12, 22, 8, 60, 90, 250, 300, null),
('scorpion_king', 13, 170, 15, 22, 14, 24, 10, 58, 88, 250, 300, null),
('desert_raider', 13, 150, 14, 21, 10, 28, 12, 55, 85, 250, 300, null),
('mummy_guardian', 14, 190, 17, 25, 15, 18, 6, 65, 95, 270, 320, null),
('ancient_serpent', 14, 180, 16, 23, 13, 26, 9, 68, 98, 270, 320, null),
('sand_elemental', 14, 200, 18, 26, 16, 20, 6, 70, 100, 270, 320, null),


-- 🏯 Восточная крепость (15-16)
('samurai_ghost', 15, 170, 17, 25, 14, 32, 14, 65, 95, 290, 340, null),
('ninja_specter', 15, 150, 16, 23, 12, 35, 20, 63, 93, 290, 340, null),
('ronin_traitor', 15, 160, 15, 24, 13, 30, 16, 60, 90, 290, 340, null),
('shogun_spirit', 16, 190, 18, 26, 16, 28, 12, 70, 100, 310, 360, null),
('dark_monk', 16, 180, 17, 25, 15, 34, 10, 73, 103, 310, 360, null),
('crimson_archer', 16, 170, 16, 24, 14, 36, 18, 68, 98, 310, 360, null),


-- 🕸️ Паутиные катакомбы (17-18)
('giant_spider', 17, 200, 18, 26, 15, 22, 14, 75, 105, 340, 390, null),
('web_caster', 17, 180, 17, 24, 13, 26, 22, 70, 100, 340, 390, null),
('poison_widow', 17, 170, 16, 25, 12, 30, 18, 73, 103, 340, 390, null),
('spider_matriarch', 18, 210, 19, 28, 17, 24, 12, 80, 110, 370, 420, null),
('venom_seeker', 18, 190, 18, 26, 15, 32, 16, 78, 108, 370, 420, null),
('dark_arachna', 18, 180, 17, 27, 14, 36, 20, 75, 105, 370, 420, null),


-- ⚡ Паровой город (19-20)
('steam_soldier', 19, 220, 19, 28, 16, 18, 10, 85, 115, 390, 440, null),
('gadget_assassin', 19, 180, 18, 25, 13, 34, 16, 83, 113, 390, 440, null),
('clockwork_beast', 19, 230, 20, 29, 18, 14, 8, 90, 120, 390, 440, null),
('tesla_conductor', 20, 210, 19, 30, 15, 28, 12, 95, 125, 410, 460, null),
('iron_engineer', 20, 200, 18, 27, 14, 24, 14, 93, 123, 410, 460, null),
('lightning_golem', 20, 240, 21, 32, 17, 22, 6, 100, 130, 410, 460, null),


-- 🧠 Академия магии (21-22)
('magic_sentinel', 21, 220, 20, 28, 16, 26, 12, 100, 135, 430, 480, null),
('arcane_scholar', 21, 200, 19, 27, 15, 30, 16, 98, 133, 430, 480, null),
('mana_wisp', 21, 170, 18, 25, 13, 38, 20, 95, 130, 430, 480, null),
('spell_reaver', 22, 230, 21, 30, 17, 32, 12, 105, 140, 460, 510, null),
('chrono_mage', 22, 210, 20, 29, 15, 36, 14, 103, 138, 460, 510, null),
('elder_arcanist', 22, 200, 19, 28, 16, 38, 18, 100, 135, 460, 510, null),


-- 🗡️ Финальные земли (23–25)
('dark_paladin', 23, 240, 22, 30, 19, 30, 10, 115, 145, 510, 560, null),
('void_walker', 23, 220, 21, 29, 17, 34, 16, 110, 140, 510, 560, null),
('nightmare_hound', 23, 230, 22, 28, 18, 28, 14, 113, 143, 510, 560, null),
('doom_bringer', 24, 260, 23, 31, 20, 32, 12, 120, 150, 540, 590, null),
('abyssal_knight', 24, 250, 22, 30, 19, 34, 14, 118, 148, 540, 590, null),
('avatar_of_chaos', 25, 280, 25, 34, 22, 40, 18, 125, 155, 610, 660, null),


-- Начальные боссы (уровни 5-10)
('goblin_warlord', 6, 200, 12, 18, 6, 10, 5, 50, 80, 180, 270, 'usual'),
('undead_captain', 7, 220, 14, 20, 8, 12, 4, 55, 90, 240, 360, 'usual'),
('troll_berserker', 8, 270, 18, 26, 9, 15, 3, 60, 100, 300, 420, 'usual'),
('shadow_assassin', 9, 180, 18, 24, 5, 25, 15, 70, 120, 360, 540, 'usual'),
('dark_sorcerer_lord', 10, 230, 15, 25, 7, 20, 10, 80, 130, 450, 600, 'usual'),

-- Средние боссы (уровни 11-15)
('demonic_knight', 11, 300, 20, 30, 12, 18, 7, 100, 150, 510, 660, 'usual'),
('cursed_warlord', 12, 320, 22, 32, 14, 20, 6, 110, 160, 600, 750, 'usual'),
('fire_golem', 13, 350, 25, 35, 16, 15, 5, 120, 170, 690, 840, 'usual'),
('frost_titan', 14, 370, 26, 38, 18, 17, 4, 130, 180, 780, 900, 'usual'),
('necromancer_king', 15, 400, 28, 40, 20, 22, 3, 140, 200, 840, 990, 'usual'),

-- Высшие боссы (уровни 16-20)
('doom_bringer', 16, 450, 30, 45, 22, 25, 2, 160, 220, 900, 1050, 'usual'),
('chaos_behemoth', 17, 480, 32, 48, 24, 27, 1, 170, 240, 990, 1140, 'usual'),
('ancient_dragon', 18, 500, 35, 50, 26, 30, 1, 200, 250, 1080, 1200, 'usual'),
('archdemon', 19, 550, 42, 60, 28, 35, 0, 220, 280, 1140, 1290, 'usual'),
('god_of_destruction', 20, 650, 45, 65, 30, 38, 0, 250, 300, 1200, 1350, 'usual'),

-- Легендарные боссы (уровни 21–25)
('avatar_of_the_void', 21, 700, 50, 70, 32, 30, 5, 280, 340, 1260, 1410, 'usual'),
('twilight_reaper', 22, 750, 52, 75, 34, 32, 6, 300, 360, 1350, 1500, 'usual'),
('elder_lich', 23, 800, 55, 80, 36, 34, 4, 320, 380, 1500, 1650, 'usual'),
('celestial_knight', 24, 850, 58, 85, 38, 36, 3, 340, 400, 1590, 1740, 'usual'),
('emperor_of_shadows', 25, 900, 60, 90, 40, 40, 2, 360, 420, 1800, 2000, 'usual'),

-- Апокалиптические боссы (уровни 26–30)
('fury_of_the_elements', 26, 1000, 65, 95, 42, 32, 4, 400, 460, 1900, 2100, 'usual'),
('prime_abyssal', 27, 1100, 68, 100, 45, 35, 3, 420, 480, 2050, 2250, 'usual'),
('draconic_overlord', 28, 1200, 70, 105, 48, 37, 3, 440, 500, 2300, 2500, 'usual'),
('void_tyrant', 29, 1300, 75, 110, 50, 38, 2, 470, 530, 2400, 2580, 'usual'),
('titan_of_oblivion', 30, 1500, 80, 120, 55, 40, 1, 500, 550, 2500, 3000, 'usual');


TRUNCATE table items cascade;
INSERT INTO "items" 
("name", "category", "equips_on", "attack_bonus", "defense_bonus", "critical_chance_percent_bonus", "evasion_percent_bonus", "base_price") 
VALUES

('rusty_sword', 'weapon', 'hand', 3, NULL, 1, NULL, 50),
('iron_sword', 'weapon', 'hand', 5, NULL, 2, NULL, 100),
('steel_sword', 'weapon', 'hand', 7, NULL, 3, NULL, 200),
('knights_blade', 'weapon', 'hand', 10, NULL, 5, NULL, 350),
('shadow_dagger', 'weapon', 'hand', 4, NULL, 2, 1, 120),
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
('bronze_armor', 'armor', 'body', NULL, 4, NULL, 0, 100),
('iron_helmet', 'armor', 'head', NULL, 3, NULL, -1, 150),
('knights_helmet', 'armor', 'head', NULL, 5, NULL, -2, 300),
('shadow_hood', 'armor', 'head', NULL, 2, NULL, 3, 180),
('crown_of_thorns', 'armor', 'head', NULL, 4, 3, -1, 250),
('battle_mask', 'armor', 'head', NULL, 3, NULL, 2, 200),
('hunters_gloves', 'armor', 'hands', NULL, NULL, 3, 3, 160),
('phantom_gloves', 'armor', 'hands', NULL, NULL, 4, 4, 270),
('titan_gauntlets', 'accessory', 'hands', 3, 3, NULL, -2, 180),
('agile_boots', 'armor', 'feet', NULL, NULL, NULL, 5, 120),
('boots_of_speed', 'armor', 'feet', NULL, NULL, NULL, 6, 200),
('elven_boots', 'armor', 'feet', NULL, NULL, NULL, 7, 280),
('dark_cloak', 'armor', 'body', NULL, 5, NULL, 5, 320),

('lucky_amulet', 'accessory', 'neck', NULL, NULL, 5, NULL, 100),
('ring_of_protection', 'accessory', 'finger', NULL, 3, NULL, NULL, 90),
('ring_of_agility', 'accessory', 'finger', NULL, NULL, NULL, 4, 110),
('ring_of_strength', 'accessory', 'finger', 2, NULL, NULL, NULL, 130),
('amulet_of_power', 'accessory', 'neck', 3, NULL, 4, NULL, 180),
('ancient_talisman', 'accessory', 'neck', NULL, NULL, 7, 2, 300),
('blood_ring', 'accessory', 'finger', 4, NULL, NULL, -2, 150);

-- Potions --
INSERT INTO "items" 
("name", "category", "potion_effect_type", "potion_effect_value", "potion_effect_duration", "base_price") 
VALUES
('health_potion', 'potion', 'heal', 30, NULL, 50),
('greater_health_potion', 'potion', 'heal', 60, NULL, 100);
-- ('strength_potion', 'potion', 'attack_buff', 5, 3, 80),
-- ('defense_potion', 'potion', 'defense_buff', 5, 3, 80),
-- ('evasion_potion', 'potion', 'evasion_buff', 10, 2, 70),
-- ('invisibility_potion', 'potion', 'evasion_buff', 50, 1, 200);
