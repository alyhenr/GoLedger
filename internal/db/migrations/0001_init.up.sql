CREATE TABLE accounts (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE transactions (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  reference TEXT,
  created_at TIMESTAMPTZ DEFAULT now()
);

CREATE TABLE entries (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  transaction_id UUID NOT NULL REFERENCES transactions(id),
  account_id UUID NOT NULL REFERENCES accounts(id),
  amount NUMERIC NOT NULL,
  description TEXT,
  created_at TIMESTAMPTZ DEFAULT now()
);

-- Ensure double-entry (transaction balance = 0)
CREATE FUNCTION validate_transaction_balance() RETURNS TRIGGER AS $$
BEGIN
  IF (
    SELECT SUM(amount)
    FROM entries
    WHERE transaction_id = NEW.transaction_id
  ) != 0 THEN
    RAISE EXCEPTION 'Transaction not balanced!';
  END IF;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER check_balance_trigger
AFTER INSERT ON entries
FOR EACH ROW
EXECUTE FUNCTION validate_transaction_balance();
