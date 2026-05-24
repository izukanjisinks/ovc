-- +goose Up
CREATE TABLE children (
    id                  UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    pupil_id            VARCHAR(50) UNIQUE NOT NULL,
    first_name          VARCHAR(100) NOT NULL,
    last_name           VARCHAR(100) NOT NULL,
    address             VARCHAR(255) NOT NULL,
    class_name          VARCHAR(20) NOT NULL,
    image_url           VARCHAR(500),
    guardian_first_name VARCHAR(100) NOT NULL,
    guardian_last_name  VARCHAR(100) NOT NULL,
    guardian_address    VARCHAR(255) NOT NULL,
    guardian_phone      VARCHAR(20) NOT NULL,
    created_by          UUID REFERENCES users(id) ON DELETE SET NULL,
    created_at          TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at          TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_children_pupil_id ON children(pupil_id);

-- +goose Down
DROP TABLE IF EXISTS children;
