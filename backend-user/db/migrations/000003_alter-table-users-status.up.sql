-- Criar tipo enum para status
CREATE TYPE user_status AS ENUM ('active', 'inactive');

-- Adicionar coluna status com valor padrão 'active'
ALTER TABLE users 
ADD COLUMN status user_status DEFAULT 'active' NOT NULL;
