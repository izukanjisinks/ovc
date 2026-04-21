-- +goose Up
ALTER TABLE users DROP CONSTRAINT fk_users_role;
ALTER TABLE role_permissions DROP CONSTRAINT role_permissions_role_id_fkey;
ALTER TABLE role_permissions DROP CONSTRAINT role_permissions_pkey;
ALTER TABLE role_permissions DROP CONSTRAINT role_permissions_resource_action_fkey;

-- roles: replace SERIAL PK with UUID
ALTER TABLE roles ADD COLUMN uuid UUID DEFAULT gen_random_uuid();
UPDATE roles SET uuid = gen_random_uuid();

ALTER TABLE users ADD COLUMN role_uuid UUID;
ALTER TABLE role_permissions ADD COLUMN role_uuid UUID;

UPDATE users u SET role_uuid = r.uuid FROM roles r WHERE r.id = u.role_id;
UPDATE role_permissions rp SET role_uuid = r.uuid FROM roles r WHERE r.id = rp.role_id;

ALTER TABLE roles DROP CONSTRAINT roles_pkey;
ALTER TABLE roles DROP COLUMN id;
ALTER TABLE roles RENAME COLUMN uuid TO id;
ALTER TABLE roles ADD PRIMARY KEY (id);
ALTER TABLE roles ALTER COLUMN id SET DEFAULT gen_random_uuid();

ALTER TABLE users DROP COLUMN role_id;
ALTER TABLE users RENAME COLUMN role_uuid TO role_id;
ALTER TABLE users ALTER COLUMN role_id SET NOT NULL;
ALTER TABLE users ADD CONSTRAINT fk_users_role FOREIGN KEY (role_id) REFERENCES roles(id);

ALTER TABLE role_permissions DROP COLUMN role_id;
ALTER TABLE role_permissions RENAME COLUMN role_uuid TO role_id;
ALTER TABLE role_permissions ALTER COLUMN role_id SET NOT NULL;
ALTER TABLE role_permissions ADD PRIMARY KEY (role_id, resource, action);
ALTER TABLE role_permissions ADD CONSTRAINT role_permissions_role_id_fkey
    FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE CASCADE;

-- permissions: add UUID id column as PK, keep (resource, action) unique
ALTER TABLE permissions ADD COLUMN id UUID DEFAULT gen_random_uuid();
ALTER TABLE permissions DROP CONSTRAINT permissions_pkey;
ALTER TABLE permissions ADD PRIMARY KEY (id);
ALTER TABLE permissions ADD CONSTRAINT permissions_resource_action_key UNIQUE (resource, action);

ALTER TABLE role_permissions ADD CONSTRAINT role_permissions_resource_action_fkey
    FOREIGN KEY (resource, action) REFERENCES permissions(resource, action) ON DELETE CASCADE;

-- +goose Down
-- This migration is not safely reversible. Restore from backup if needed.
