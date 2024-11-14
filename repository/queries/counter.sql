
-- name: GetCounter :one
SELECT count FROM counter WHERE id = 1;

-- name: IncrementCounter :exec
UPDATE counter SET count = count + 1 WHERE id = 1;

-- name: DecrementCounter :exec
UPDATE counter SET count = count - 1 WHERE id = 1;

-- name: InitializeCounter :exec
INSERT INTO counter (id, count) VALUES (1, 0) ON CONFLICT (id) DO NOTHING;

