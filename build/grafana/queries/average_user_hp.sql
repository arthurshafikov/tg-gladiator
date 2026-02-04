-- Average (median) user HP at the end of the fight
SELECT
    percentile_cont(0.5) WITHIN GROUP (ORDER BY hero_hp) AS median_hero_hp,
    avg(hero_hp) AS avg_hero_hp,
    count(*) AS fights_count
FROM fights
WHERE status = 'ended'
  AND created_at >= $__timeFrom(created_at)
  AND created_at <= $__timeTo(created_at);
