CREATE TABLE IF NOT EXISTS dates
(
    date_uuid               UUID PRIMARY KEY DEFAULT gen_random_uuid()      NOT NULL,
    dairy_date              DATE UNIQUE                                     NOT NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_date ON dates (dairy_date);

CREATE TABLE IF NOT EXISTS tasks
(
    task_uuid               UUID PRIMARY KEY DEFAULT gen_random_uuid()      NOT NULL,
    dairy_time              TIMESTAMP                                       NOT NULL,
    dairy_task              TEXT                                            NOT NULL,
    date_uuid               UUID REFERENCES dates(date_uuid)                NOT NULL
    
);
CREATE UNIQUE INDEX IF NOT EXISTS idx_time ON tasks (dairy_time);
CREATE INDEX IF NOT EXISTS idx_task ON tasks (dairy_task);