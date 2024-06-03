
-- name: GetAssociatedRolePassives :many
SELECT pd.*
FROM role_passives_join rpj
JOIN passive_details pd ON rpj.passive_id = pd.id
WHERE rpj.role_id = $1;
