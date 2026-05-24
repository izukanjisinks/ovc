-- +goose Up

-- OVC Categories
CREATE TABLE ovc_categories (
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(150) UNIQUE NOT NULL
);

-- Requisites
CREATE TABLE requisites (
    id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name          VARCHAR(100) UNIQUE NOT NULL,
    default_price NUMERIC(10,2) NOT NULL DEFAULT 0.00
);

-- Sponsors
CREATE TABLE sponsors (
    id   UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(100) UNIQUE NOT NULL
);

-- Junction: children <-> categories
CREATE TABLE child_categories (
    child_id    UUID NOT NULL REFERENCES children(id) ON DELETE CASCADE,
    category_id UUID NOT NULL REFERENCES ovc_categories(id) ON DELETE CASCADE,
    PRIMARY KEY (child_id, category_id)
);

-- Junction: children <-> requisites
CREATE TABLE child_requisites (
    child_id       UUID NOT NULL REFERENCES children(id) ON DELETE CASCADE,
    requisite_id   UUID NOT NULL REFERENCES requisites(id) ON DELETE CASCADE,
    quantity       INTEGER NOT NULL DEFAULT 1,
    checked        BOOLEAN NOT NULL DEFAULT false,
    price_per_item NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    PRIMARY KEY (child_id, requisite_id)
);

-- Junction: children <-> sponsors
CREATE TABLE child_sponsors (
    child_id   UUID NOT NULL REFERENCES children(id) ON DELETE CASCADE,
    sponsor_id UUID NOT NULL REFERENCES sponsors(id) ON DELETE CASCADE,
    PRIMARY KEY (child_id, sponsor_id)
);

-- Seed: OVC Categories
INSERT INTO ovc_categories (name) VALUES
    ('SINGLE ORPHAN'),
    ('DOUBLE ORPHAN'),
    ('EXTENDED FAMILY UNABLE TO MEET SCHOOL COST'),
    ('CHILD HEADED HOUSEHOLD'),
    ('ELDERLY PERSON HEADED HOUSEHOLDS AGED 65'),
    ('PARENT WITH NO RELIABLE SOURCE OF INCOME'),
    ('ON SCHOOL RE-ENTRY PROGRAM WITHOUT FAMILY SUPPORT'),
    ('POOR AS DEFINED BY CWAC');

-- Seed: Requisites
INSERT INTO requisites (name) VALUES
    ('SCHOOL UNIFORM'),
    ('SCHOOL BAGS'),
    ('SANITARY TOWELS'),
    ('NOTE BOOKS'),
    ('PENS AND PENCILS'),
    ('MATHEMATICAL SETS'),
    ('CALCULATORS'),
    ('TEXT BOOKS');

-- Seed: Sponsors
INSERT INTO sponsors (name) VALUES
    ('GOVERNMENT REPUBLIC OF ZAMBIA (GRZ)'),
    ('NON GOVERNMENTAL ORGANIZATION (NGO)'),
    ('OTHER SPONSORS');

-- +goose Down
DROP TABLE IF EXISTS child_sponsors;
DROP TABLE IF EXISTS child_requisites;
DROP TABLE IF EXISTS child_categories;
DROP TABLE IF EXISTS sponsors;
DROP TABLE IF EXISTS requisites;
DROP TABLE IF EXISTS ovc_categories;
