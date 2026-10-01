-- 001: items table (with soft delete)

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE items (
    id         UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    sku        TEXT        NOT NULL,
    name       TEXT        NOT NULL,
    category   TEXT        NOT NULL,
    unit       TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    deleted_at TIMESTAMPTZ,
    CONSTRAINT items_sku_unique UNIQUE (sku)
);

CREATE INDEX idx_items_sku ON items (sku) WHERE deleted_at IS NULL;
CREATE INDEX idx_items_category ON items (category) WHERE deleted_at IS NULL;
