CREATE TABLE IF NOT EXISTS dates
(
    uuid               UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid(),
    dairy_date              DATE UNIQUE NOT NULL,
    is_deleted				BOOLEAN NOT NULL DEFAULT 0,
    created_at				DATETIME NOT NULL DEFAULT NOW(),
    updated_at				DATETIME NOT NULL ON UPDATE NOW()
);

CREATE TABLE IF NOT EXISTS tasks
(
    uuid               UUID PRIMARY key NOT NULL DEFAULT gen_random_uuid(),
    dairy_time              TIMESTAMP NOT NULL UNIQUE,
    dairy_task              TEXT  NOT NULL,
    date_uuid               UUID NOT NULL REFERENCES dates(uuid),
    is_deleted				BOOLEAN NOT NULL DEFAULT 0,
    created_at				DATETIME NOT NULL DEFAULT NOW(),
    updated_at				DATETIME NOT NULL ON UPDATE NOW()
    
);
CREATE INDEX IF NOT EXISTS idx_time ON tasks (dairy_time);
CREATE INDEX IF NOT EXISTS idx_task ON tasks (dairy_task);