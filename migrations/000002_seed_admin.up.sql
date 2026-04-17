-- Default admin: admin@ovc.local / Admin@1234
-- Password hash generated with bcrypt cost 10
INSERT INTO users (email, password_hash, full_name, role)
VALUES (
    'admin@ovc.local',
    '$2a$10$N9qo8uLOickgx2ZMRZoMyeIjZAgcfl7p92ldGxad68LJZdL17lhWy',
    'System Administrator',
    'admin'
) ON CONFLICT (email) DO NOTHING;
