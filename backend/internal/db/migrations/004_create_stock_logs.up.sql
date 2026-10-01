-- 004: stock transaction logs

CREATE TABLE stock_logs (
    id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    type        TEXT        NOT NULL,
    item_id     UUID        NOT NULL REFERENCES items(id),
    location_id UUID        NOT NULL REFERENCES locations(id),
    qty         INTEGER     NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE INDEX idx_stock_logs_item ON stock_logs (item_id);
CREATE INDEX idx_stock_logs_location ON stock_logs (location_id);
