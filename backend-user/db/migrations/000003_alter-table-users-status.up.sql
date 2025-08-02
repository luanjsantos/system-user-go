-- Criar tipo enum para status (se não existir)
DO $$ BEGIN
    CREATE TYPE user_status AS ENUM ('active', 'inactive');
EXCEPTION
    WHEN duplicate_object THEN null;
END $$;

-- Adicionar coluna status com valor padrão 'active'
ALTER TABLE users 
ADD COLUMN status user_status DEFAULT 'active' NOT NULL;
