-- Adicionar colunas sem NOT NULL primeiro
ALTER TABLE users 
ADD COLUMN password VARCHAR(255) DEFAULT '',
ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;

-- Atualizar registros existentes com senha padrão
UPDATE users SET password = '$2a$10$default.hash.for.existing.users' WHERE password = '';

-- Adicionar NOT NULL constraint
ALTER TABLE users ALTER COLUMN password SET NOT NULL;
