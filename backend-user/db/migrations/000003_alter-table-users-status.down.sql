-- Remover coluna status
ALTER TABLE users DROP COLUMN status;

-- Remover tipo enum
DROP TYPE user_status;
