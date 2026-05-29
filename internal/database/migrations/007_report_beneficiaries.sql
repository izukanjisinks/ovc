-- +goose Up
CREATE TABLE report_beneficiaries (
    report_id UUID NOT NULL REFERENCES reports(id) ON DELETE CASCADE,
    child_id  UUID NOT NULL REFERENCES children(id) ON DELETE CASCADE,
    PRIMARY KEY (report_id, child_id)
);

-- +goose Down
DROP TABLE IF EXISTS report_beneficiaries;
