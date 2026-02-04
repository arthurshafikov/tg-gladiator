SELECT
    date_trunc('hour', timestamp) AS hour,
    COUNT(*) AS chat_events_count
FROM analytic_events
WHERE type = 'chat_update_processed'
  AND timestamp >= $__timeFrom(timestamp)
  AND timestamp <= $__timeTo(timestamp)
GROUP BY hour
ORDER BY hour;

SELECT
    i.name AS item_name,
    i.base_price,
    COUNT(*) AS purchase_count
FROM analytic_events ae
JOIN items i ON (ae.payload->>'item_id')::bigint = i.id
WHERE ae.type = 'item_purchased'
  AND ae.timestamp >= $__timeFrom(timestamp)
  AND ae.timestamp <= $__timeTo(timestamp)
GROUP BY i.id, i.name, i.base_price
ORDER BY purchase_count DESC, i.base_price DESC;

SELECT
    e.name AS enemy_name,
    COUNT(*) AS fight_count
FROM analytic_events ae
JOIN fights f ON (ae.payload->>'fight_id')::bigint = f.id
JOIN enemies e ON f.opponent_id = e.id
WHERE ae.type = 'fight_started'
  AND f.opponent_type = 'mob'
  AND ae.timestamp >= $__timeFrom(timestamp)
  AND ae.timestamp <= $__timeTo(timestamp)
GROUP BY e.id, e.name
ORDER BY fight_count DESC;

SELECT
    e.level AS enemy_level,
    COUNT(*) AS fight_count
FROM analytic_events ae
JOIN fights f ON (ae.payload->>'fight_id')::bigint = f.id
JOIN enemies e ON f.opponent_id = e.id
WHERE ae.type = 'fight_started'
  AND f.opponent_type = 'mob'
  AND ae.timestamp >= $__timeFrom(timestamp)
  AND ae.timestamp <= $__timeTo(timestamp)
GROUP BY e.level
ORDER BY fight_count DESC;
