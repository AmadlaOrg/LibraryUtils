-- All the tables used for plugin system

-- This table stores information about plugin types
CREATE TABLE IF NOT EXISTS plugin_types (
    name TEXT UNIQUE -- Just the name of the plugin (unique)
);

-- CREATE INDEX IF NOT EXISTS idx_plugin_types_name ON plugin_types(Name);

-- This table stores information about plugins
CREATE TABLE IF NOT EXISTS plugins (
    plugin_types_rowid INTEGER,
    name TEXT, -- Just the name of the plugin (unique)
    uri TEXT UNIQUE,  -- Contains the entities URI
    repo_url TEXT,  -- The full URL to the repository containing the entity
    origin TEXT,   -- Contains the partial path of the entity
    version TEXT,  -- The entity version
    is_latest_version BOOLEAN,  -- Whether the entity version is the latest
    is_pseudo_version BOOLEAN,
    abs_path TEXT UNIQUE, -- The full system path to the entity files (unique)
    have BOOLEAN,  -- Indicates if the entity is on the local machine
    hash TEXT UNIQUE,     -- The hash of the entity content for validation
    exist BOOLEAN, -- Indicates if the repository was found
    insert_date_time DATETIME DEFAULT CURRENT_TIMESTAMP,  -- The date when the row was added
    update_date_time DATETIME DEFAULT CURRENT_TIMESTAMP,   -- The latest date when the row was updated
    FOREIGN KEY (plugin_types_rowid) REFERENCES plugin_types(rowid)
);

CREATE INDEX IF NOT EXISTS idx_plugins_name ON plugins(Name);
CREATE INDEX IF NOT EXISTS idx_plugins_repo_url ON plugins(repo_url);
CREATE INDEX IF NOT EXISTS idx_plugins_version ON plugins(version);
CREATE INDEX IF NOT EXISTS idx_plugins_exist ON plugins(exist);
CREATE INDEX IF NOT EXISTS idx_plugins_is_latest_version ON plugins(is_latest_version);
CREATE INDEX IF NOT EXISTS idx_plugins_is_pseudo_version ON plugins(is_pseudo_version);
CREATE INDEX IF NOT EXISTS idx_plugins_have ON plugins(have);
