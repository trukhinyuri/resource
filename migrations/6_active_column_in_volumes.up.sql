BEGIN TRANSACTION;
ALTER TABLE volumes
    ADD COLUMN IF NOT EXISTS is_active BOOLEAN DEFAULT FALSE;
COMMIT TRANSACTION;