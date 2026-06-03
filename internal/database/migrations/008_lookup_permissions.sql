-- +goose Up

INSERT INTO permissions (resource, action) VALUES
    ('lookups', 'create'),
    ('lookups', 'update'),
    ('lookups', 'delete');

-- Grant to admin role only
INSERT INTO role_permissions (role_id, resource, action)
SELECT id, p.resource, p.action
FROM roles, (VALUES ('lookups', 'create'), ('lookups', 'update'), ('lookups', 'delete')) AS p(resource, action)
WHERE roles.name = 'admin';

-- +goose Down

DELETE FROM role_permissions WHERE resource = 'lookups' AND action IN ('create', 'update', 'delete');
DELETE FROM permissions WHERE resource = 'lookups' AND action IN ('create', 'update', 'delete');
