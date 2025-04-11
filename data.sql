TRUNCATE table enemies cascade;
INSERT INTO "enemies" 
("name", "level", "hp", "min_attack", "max_attack", "defense", "critical_chance_percent", "evasion_chance_percent", "gold_reward_min", "gold_reward_max", "xp_reward_min", "xp_reward_max", "boss_type") 
VALUES

-- 🌲 Лес (1-2)
('goblin', 1, 50, 3, 6, 2, 5, 5, 5, 10, 10, 20, null),
('wild_dog', 1, 45, 3, 6, 1, 3, 7, 4, 9, 10, 20, null),
('forest_snake', 1, 40, 2, 5, 1, 4, 10, 4, 8, 10, 20, null),
('bandit', 2, 55, 4, 7, 3, 6, 4, 6, 12, 15, 25, null),
('slow_trent', 2, 100, 2, 4, 10, 0, 0, 5, 10, 15, 25, null),
('wolf', 2, 60, 5, 7, 2, 8, 7, 6, 12, 15, 25, null),

-- ⚰️ Кладбище (3-4)
('skeleton', 3, 60, 4, 7, 3, 7, 10, 7, 15, 30, 50, null),
('orc', 3, 80, 5, 9, 4, 10, 3, 10, 20, 30, 50, null),
('haunted_soul', 3, 55, 5, 8, 2, 12, 12, 8, 18, 30, 50, null),
('zombie', 4, 70, 6, 10, 3, 15, 5, 12, 18, 40, 60, null),
('vampire', 4, 65, 5, 8, 4, 12, 6, 9, 16, 40, 60, null),
('rotting_corpse', 4, 75, 6, 9, 5, 6, 3, 10, 15, 40, 60, null),

-- 🏰 Руины (5-6)
('orc', 5, 100, 8, 12, 6, 12, 3, 20, 30, 50, 80, null),
('undead_knight', 5, 90, 7, 11, 5, 14, 4, 18, 25, 50, 80, null),
('giant_bat', 5, 85, 6, 10, 3, 16, 10, 14, 22, 50, 80, null),
('bandit_leader', 6, 85, 7, 11, 5, 17, 7, 15, 28, 60, 90, null),
('dark_sorcerer', 6, 75, 6, 12, 3, 20, 7, 22, 30, 60, 90, null),
('cursed_monk', 6, 80, 8, 11, 6, 15, 6, 18, 26, 60, 90, null),

-- 🌋 Огненные пещеры (7-8)
('red_ogre', 7, 120, 10, 15, 7, 15, 2, 25, 40, 80, 120, null),
('satan_servant', 7, 90, 9, 14, 4, 25, 15, 20, 35, 80, 120, null),
('flame_imp', 7, 100, 10, 13, 5, 18, 8, 18, 32, 80, 120, null),
('fire_knight', 8, 110, 11, 16, 8, 18, 5, 30, 45, 100, 140, null),
('fire_elemental', 8, 100, 12, 17, 6, 22, 8, 28, 50, 100, 140, null),
('volcanic_golem', 8, 130, 10, 15, 10, 10, 4, 35, 50, 100, 140, null),

-- 😈 Преисподняя (9-10)
('demon', 9, 150, 14, 20, 9, 25, 5, 35, 60, 120, 180, null),
('ancient_warlord', 9, 140, 13, 18, 10, 20, 7, 30, 55, 120, 180, null),
('infernal_hound', 9, 130, 12, 17, 8, 15, 10, 28, 50, 120, 180, null),
('dragonspawn', 10, 160, 15, 22, 12, 28, 10, 40, 70, 150, 200, null),
('chaos_champion', 10, 170, 16, 24, 11, 30, 12, 45, 80, 150, 200, null),
('hellfire_witch', 10, 150, 13, 21, 9, 27, 9, 42, 75, 150, 200, null),

-- 🧊 Ледяные пещеры (11-12)
('ice_golem', 11, 180, 14, 20, 14, 10, 5, 50, 80, 170, 220, null),
('frost_wraith', 11, 140, 13, 18, 9, 22, 15, 45, 75, 170, 220, null),
('snow_wolf', 11, 130, 12, 17, 8, 18, 12, 40, 70, 170, 220, null),
('glacier_elemental', 12, 160, 15, 22, 12, 20, 6, 55, 85, 200, 250, null),
('frozen_cultist', 12, 150, 14, 20, 10, 25, 10, 48, 80, 200, 250, null),
('yeti', 12, 190, 16, 23, 13, 12, 4, 60, 90, 200, 250, null),

-- 🏜️ Пустыня (13-14)
('sand_wyrm', 13, 170, 15, 22, 10, 18, 6, 55, 85, 230, 280, null),
('scorpion_king', 13, 160, 14, 21, 11, 20, 8, 50, 80, 230, 280, null),
('desert_raider', 13, 140, 13, 19, 9, 24, 10, 45, 75, 230, 280, null),
('mummy_guardian', 14, 180, 16, 24, 13, 14, 5, 58, 88, 260, 300, null),
('ancient_serpent', 14, 170, 15, 23, 12, 22, 7, 60, 90, 260, 300, null),
('sand_elemental', 14, 190, 17, 25, 14, 16, 5, 63, 93, 260, 300, null),

-- 🏯 Восточная крепость (15-16)
('samurai_ghost', 15, 160, 16, 23, 12, 28, 12, 60, 90, 280, 330, null),
('ninja_specter', 15, 140, 15, 21, 10, 30, 18, 58, 88, 280, 330, null),
('ronin_traitor', 15, 150, 14, 22, 11, 26, 14, 56, 86, 280, 330, null),
('shogun_spirit', 16, 180, 17, 25, 14, 25, 10, 65, 95, 300, 350, null),
('dark_monk', 16, 170, 16, 24, 13, 30, 8, 68, 98, 300, 350, null),
('crimson_archer', 16, 160, 15, 22, 12, 32, 16, 62, 92, 300, 350, null),

-- 🕸️ Паутиные катакомбы (17-18)
('giant_spider', 17, 190, 17, 24, 13, 18, 12, 70, 100, 330, 380, null),
('web_caster', 17, 170, 16, 22, 11, 22, 20, 65, 95, 330, 380, null),
('poison_widow', 17, 160, 15, 23, 10, 26, 16, 68, 98, 330, 380, null),
('spider_matriarch', 18, 200, 18, 26, 15, 20, 10, 75, 105, 360, 400, null),
('venom_seeker', 18, 180, 17, 24, 13, 28, 14, 73, 103, 360, 400, null),
('dark_arachna', 18, 170, 16, 25, 12, 32, 18, 70, 100, 360, 400, null),

-- ⚡ Паровой город (19-20)
('steam_soldier', 19, 200, 18, 25, 14, 15, 8, 80, 110, 380, 430, null),
('gadget_assassin', 19, 170, 17, 23, 11, 30, 14, 78, 108, 380, 430, null),
('clockwork_beast', 19, 210, 19, 26, 16, 12, 6, 85, 115, 380, 430, null),
('tesla_conductor', 20, 190, 18, 27, 13, 25, 10, 90, 120, 400, 450, null),
('iron_engineer', 20, 180, 17, 24, 12, 20, 12, 88, 118, 400, 450, null),
('lightning_golem', 20, 220, 20, 28, 15, 18, 5, 95, 125, 400, 450, null),

-- 🧠 Академия магии (21-22)
('magic_sentinel', 21, 200, 19, 26, 14, 22, 10, 95, 130, 420, 470, null),
('arcane_scholar', 21, 180, 18, 25, 13, 26, 14, 92, 127, 420, 470, null),
('mana_wisp', 21, 160, 17, 23, 11, 35, 18, 88, 125, 420, 470, null),
('spell_reaver', 22, 210, 20, 28, 15, 28, 10, 100, 135, 450, 500, null),
('chrono_mage', 22, 190, 19, 27, 13, 32, 12, 98, 132, 450, 500, null),
('elder_arcanist', 22, 180, 18, 26, 14, 34, 15, 95, 130, 450, 500, null),

-- 🗡️ Финальные земли (23–25)
('dark_paladin', 23, 220, 21, 29, 17, 26, 8, 110, 140, 500, 550, null),
('void_walker', 23, 200, 20, 28, 15, 30, 14, 105, 135, 500, 550, null),
('nightmare_hound', 23, 210, 21, 27, 16, 24, 12, 108, 138, 500, 550, null),
('doom_bringer', 24, 240, 22, 30, 18, 28, 10, 115, 145, 530, 580, null),
('abyssal_knight', 24, 230, 21, 29, 17, 30, 12, 112, 142, 530, 580, null),
('avatar_of_chaos', 25, 260, 24, 32, 20, 35, 15, 120, 150, 600, 650, null),

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
