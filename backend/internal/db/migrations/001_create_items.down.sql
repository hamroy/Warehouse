-- 001 rollback: items

DROP INDEX IF EXISTS idx_items_category;
DROP INDEX IF EXISTS idx_items_sku;
DROP TABLE IF EXISTS items CASCADE;
