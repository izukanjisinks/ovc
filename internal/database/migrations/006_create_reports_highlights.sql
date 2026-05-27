-- +goose Up
CREATE TABLE reports (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title      VARCHAR(255) NOT NULL,
    body       TEXT NOT NULL,
    term       VARCHAR(10) NOT NULL CHECK (term IN ('TERM_1', 'TERM_2', 'TERM_3')),
    year       INTEGER NOT NULL,
    created_by UUID REFERENCES users(id) ON DELETE SET NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX idx_reports_term_year ON reports(term, year);

CREATE TABLE highlights (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    image_url  VARCHAR(500) NOT NULL,
    caption    VARCHAR(255),
    term       VARCHAR(10) CHECK (term IN ('TERM_1', 'TERM_2', 'TERM_3')),
    year       INTEGER,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);

-- +goose Down
DROP TABLE IF EXISTS highlights;
DROP TABLE IF EXISTS reports;
