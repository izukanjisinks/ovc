-- +goose Up
CREATE TABLE roles (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE permissions (
    resource   VARCHAR(50) NOT NULL,
    action     VARCHAR(50) NOT NULL,
    PRIMARY KEY (resource, action)
);

CREATE TABLE role_permissions (
    role_id  INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    resource VARCHAR(50) NOT NULL,
    action   VARCHAR(50) NOT NULL,
    PRIMARY KEY (role_id, resource, action),
    FOREIGN KEY (resource, action) REFERENCES permissions(resource, action) ON DELETE CASCADE
);

ALTER TABLE users
    ADD CONSTRAINT fk_users_role FOREIGN KEY (role_id) REFERENCES roles(id);

INSERT INTO roles (id, name) VALUES
    (1, 'admin'),
    (2, 'user');

INSERT INTO permissions (resource, action) VALUES
    ('user-management', 'create'),
    ('user-management', 'read'),
    ('user-management', 'update'),
    ('user-management', 'delete'),
    ('children', 'create'),
    ('children', 'read'),
    ('children', 'update'),
    ('children', 'delete'),
    ('reports', 'create'),
    ('reports', 'read'),
    ('reports', 'update'),
    ('reports', 'delete'),
    ('reports', 'export'),
    ('dashboard', 'read'),
    ('highlights', 'create'),
    ('highlights', 'read'),
    ('highlights', 'delete'),
    ('lookups', 'read');

INSERT INTO role_permissions (role_id, resource, action)
SELECT 1, resource, action FROM permissions;

INSERT INTO role_permissions (role_id, resource, action)
SELECT 2, resource, action FROM permissions
WHERE (resource, action) NOT IN (
    ('user-management', 'create'),
    ('user-management', 'read'),
    ('user-management', 'update'),
    ('user-management', 'delete'),
    ('highlights', 'create'),
    ('highlights', 'delete')
);

-- password: Admin@1234
INSERT INTO users (email, password_hash, full_name, role_id)
VALUES (
    'admin@ovc.local',
    '$2a$10$G5OW.dVQ3uUpDyFyzMo6w.udtgvAQz3ekne56MDVCxQeAeO01JcFO',
    'System Administrator',
    1
) ON CONFLICT (email) DO NOTHING;

-- +goose Down
DELETE FROM users WHERE email = 'admin@ovc.local';
ALTER TABLE users DROP CONSTRAINT IF EXISTS fk_users_role;
DROP TABLE IF EXISTS role_permissions;
DROP TABLE IF EXISTS permissions;
DROP TABLE IF EXISTS roles;
