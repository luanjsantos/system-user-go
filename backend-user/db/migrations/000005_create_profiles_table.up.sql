-- Criar tabela profiles
CREATE TABLE profiles (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    bio TEXT,
    avatar_url TEXT,
    phone VARCHAR(20),
    address TEXT,
    birth_date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(user_id)
);

-- Criar índices para melhor performance
CREATE INDEX idx_profiles_user_id ON profiles(user_id);

-- Inserir perfil para o usuário administrador
INSERT INTO profiles (user_id, bio, created_at, updated_at) 
SELECT id, 'Administrador do sistema', created_at, updated_at 
FROM users 
WHERE email = 'admin@email.com' 
ON CONFLICT (user_id) DO NOTHING;
