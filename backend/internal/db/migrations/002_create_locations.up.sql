-- 002: locations table

CREATE TABLE locations (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code       TEXT        NOT NULL UNIQUE,
    zone       TEXT        NOT NULL,
    type       TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
