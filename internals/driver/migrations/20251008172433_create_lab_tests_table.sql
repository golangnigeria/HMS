-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE lab_tests (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    patient_id UUID NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
    doctor_id UUID REFERENCES doctors(id) ON DELETE SET NULL,
    lab_id UUID REFERENCES labs(id) ON DELETE SET NULL,
    test_name VARCHAR(255) NOT NULL,
    test_code VARCHAR(100),
    sample_collected_date TIMESTAMP,
    result_date TIMESTAMP,
    result TEXT,
    status VARCHAR(50) DEFAULT 'pending', -- pending, processing, completed, cancelled
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS lab_tests;
DROP EXTENSION IF EXISTS pgcrypto;
-- +goose StatementEnd
