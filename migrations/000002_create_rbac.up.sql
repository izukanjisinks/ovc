-- Roles
CREATE TABLE roles (
    id         SERIAL PRIMARY KEY,
    name       VARCHAR(50) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- Permissions: natural composite PK (resource + action)
CREATE TABLE permissions (
    resource   VARCHAR(50) NOT NULL,
    action     VARCHAR(50) NOT NULL,
    PRIMARY KEY (resource, action)
);

-- Role → Permission assignments
CREATE TABLE role_permissions (
    role_id  INTEGER NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
    resource VARCHAR(50) NOT NULL,
    action   VARCHAR(50) NOT NULL,
    PRIMARY KEY (role_id, resource, action),
    FOREIGN KEY (resource, action) REFERENCES permissions(resource, action) ON DELETE CASCADE
);

-- Wire FK on users now that roles exists
ALTER TABLE users
    ADD CONSTRAINT fk_users_role FOREIGN KEY (role_id) REFERENCES roles(id);

-- -------------------------------------------------------
-- Seed: Roles
-- -------------------------------------------------------
INSERT INTO roles (id, name) VALUES
    (1, 'admin'),
    (2, 'user');

-- -------------------------------------------------------
-- Seed: Permissions
-- -------------------------------------------------------
INSERT INTO permissions (resource, action) VALUES
    -- user management
    ('user-management', 'create'),
    ('user-management', 'read'),
    ('user-management', 'update'),
    ('user-management', 'delete'),
    -- children
    ('children', 'create'),
    ('children', 'read'),
    ('children', 'update'),
    ('children', 'delete'),
    -- reports
    ('reports', 'create'),
    ('reports', 'read'),
    ('reports', 'update'),
    ('reports', 'delete'),
    ('reports', 'export'),
    -- dashboard
    ('dashboard', 'read'),
    -- highlights
    ('highlights', 'create'),
    ('highlights', 'read'),
    ('highlights', 'delete'),
    -- lookups (categories, requisites, sponsors)
    ('lookups', 'read');

-- -------------------------------------------------------
-- Seed: Role permissions
-- admin gets everything
-- -------------------------------------------------------
INSERT INTO role_permissions (role_id, resource, action)
SELECT 1, resource, action FROM permissions;

-- user gets everything except user-management and highlights create/delete
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

-- -------------------------------------------------------
-- Seed: Default admin user  (password: Admin@1234)
-- -------------------------------------------------------
INSERT INTO users (email, password_hash, full_name, role_id)
VALUES (
    'admin@ovc.local',
    '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',
    'System Administrator',
    1
) ON CONFLICT (email) DO NOTHING;
