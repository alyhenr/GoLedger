DROP TRIGGER IF EXISTS check_balance_trigger ON entries;
DROP FUNCTION IF EXISTS validate_transaction_balance;
DROP TABLE IF EXISTS entries;
DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS accounts;
