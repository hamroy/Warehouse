-- TASK-101: Database Design
--
-- Entities: items, locations, stocks
-- Conventions:
--   - All primary keys are UUID v4 (gen_random_uuid())
--   - All tables carry created_at / updated_at timestamps
--   - items supports soft delete via deleted_at
--   - stocks is a join of items x locations with a unique pair

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

CREATE TABLE locations (
    id        UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    code      TEXT        NOT NULL UNIQUE,
    zone      TEXT        NOT NULL,
    type      TEXT        NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE TABLE stocks (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    item_id     UUID        NOT NULL REFERENCES items(id) ON DELETE CASCADE,
    location_id UUID        NOT NULL REFERENCES locations(id) ON DELETE CASCADE,
    qty         INTEGER     NOT NULL DEFAULT 0,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at  TIMESTAMPTZ NOT NULL DEFAULT now(),
    CONSTRAINT stocks_item_location_unique UNIQUE (item_id, location_id),
    CONSTRAINT stocks_qty_non_negative CHECK (qty >= 0)
);

-- Keep updated_at fresh on every UPDATE without app-side bookkeeping.
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER items_set_updated_at
    BEFORE UPDATE ON items
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER locations_set_updated_at
    BEFORE UPDATE ON locations
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

CREATE TRIGGER stocks_set_updated_at
    BEFORE UPDATE ON stocks
    FOR EACH ROW EXECUTE FUNCTION set_updated_at();

-- Soft-deleted items must be excluded from unique SKU enforcement is NOT
-- required by the PRD (unique stays global); however queries for "active"
-- items filter deleted_at IS NULL at the repository layer.

CREATE INDEX idx_items_sku ON items (sku) WHERE deleted_at IS NULL;
CREATE INDEX idx_items_category ON items (category) WHERE deleted_at IS NULL;
CREATE INDEX idx_stocks_item ON stocks (item_id);
CREATE INDEX idx_stocks_location ON stocks (location_id);
