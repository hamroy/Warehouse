-- 003 rollback: stocks + trigger

DROP TRIGGER IF EXISTS stocks_set_updated_at ON stocks;
DROP TRIGGER IF EXISTS locations_set_updated_at ON locations;
DROP TRIGGER IF EXISTS items_set_updated_at ON items;
DROP FUNCTION IF EXISTS set_updated_at();
DROP TABLE IF EXISTS stocks CASCADE;
