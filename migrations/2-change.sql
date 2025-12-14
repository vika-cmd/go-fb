UPDATE note
SET readydat = '0001-01-01'
WHERE id = 3

-- err
ALTER TABLE note
ALTER COLUMN readydat SET DEFAULT '0001-01-01'

--

ALTER TABLE alarm
RENAME COLUMN readydat TO createdat

-- TIME  DATE
ALTER TABLE alarm
ADD COLUMN readydat TIMESTAMP DEFAULT CURRENT_TIMESTAMP;


ALTER TABLE tasks
ADD COLUMN priority INT DEFAULT 4;

UPDATE note
SET readydat = '0001-01-01'
WHERE readydat = null;




UPDATE note
SET readydat = '0001-01-01'
WHERE id = null;



DELETE FROM "tasks" WHERE id=11;

ALTER TABLE tasks
DROP COLUMN readydat;

DROP TABLE alarm;

-- tasks;

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    text TEXT NOT NULL,
    completed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE NOT NULL DEFAULT NOW()
);
