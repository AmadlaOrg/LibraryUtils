-- All the tables used for testing

CREATE TABLE IF NOT EXISTS content (
    txt TEXT UNIQUE, -- Just the name of the plugin (unique)
    num INTEGER,
    bool BOOLEAN,
    real REAL,
    date_time DATETIME DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS connect (
    content_rowid INTEGER,
    FOREIGN KEY (content_rowid) REFERENCES content(rowid)
);
