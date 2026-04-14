CREATE TABLE IF NOT EXISTS registrations (
    id SERIAL PRIMARY KEY,
    form_id UUID NOT NULL,
    member_type_id UUID NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(20) NOT NULL,
    birth_date DATE NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    UNIQUE (form_id, name, email)
);

CREATE INDEX IF NOT EXISTS idx_registrations_form_id ON registrations (form_id);