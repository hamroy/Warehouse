-- 004 rollback: stock transaction logs

DROP INDEX IF EXISTS idx_stock_logs_location;
DROP INDEX IF EXISTS idx_stock_logs_item;
DROP TABLE IF EXISTS stock_logs CASCADE;
