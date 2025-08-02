-- Inserir usuário administrador inicial
-- Senha: 123 (hash bcrypt)
INSERT INTO users (name, email, password, status, created_at, updated_at) 
VALUES (
    'Administrador',
    'admin@email.com',
    '$2a$10$lVFRSLalmhmskzPVppv3m.5PMTdwW36g3GrvTlYhYK5v3/98SJuIa', -- senha: 123
    'active',
    CURRENT_TIMESTAMP,
    CURRENT_TIMESTAMP
) ON CONFLICT (email) DO NOTHING;
