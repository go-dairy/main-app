CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
NEW.updated_at = now();
RETURN NEW;
END;
$$ language 'plpgsql';

CREATE TABLE IF NOT EXISTS dates
(
    uuid               UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    dairy_date              DATE UNIQUE NOT NULL,
    is_deleted				BOOLEAN NOT NULL DEFAULT FALSE,
    created_at				TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at				TIMESTAMP NOT NULL DEFAULT NOW()
);
CREATE TRIGGER update_modified_time BEFORE UPDATE ON dates FOR EACH ROW EXECUTE PROCEDURE update_modified_column();

CREATE TABLE IF NOT EXISTS tasks
(
    uuid               UUID PRIMARY key NOT NULL DEFAULT gen_random_uuid(),
    dairy_time              TIMESTAMP NOT NULL UNIQUE,
    dairy_task              TEXT  NOT NULL,
    date_uuid               UUID NOT NULL REFERENCES dates(uuid),
    is_deleted				BOOLEAN NOT NULL DEFAULT FALSE,
    created_at				TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at				TIMESTAMP NOT NULL DEFAULT NOW()
    
);
CREATE TRIGGER update_modified_time BEFORE UPDATE ON tasks FOR EACH ROW EXECUTE PROCEDURE update_modified_column();
CREATE INDEX IF NOT EXISTS idx_time ON tasks (dairy_time);
CREATE INDEX IF NOT EXISTS idx_task ON tasks (dairy_task);