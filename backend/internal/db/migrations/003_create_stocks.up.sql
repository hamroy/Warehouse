-- 003: stocks table + updated_at trigger

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

CREATE INDEX idx_stocks_item ON stocks (item_id);
CREATE INDEX idx_stocks_location ON stocks (location_id);

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
