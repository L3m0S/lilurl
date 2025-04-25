CREATE TABLE
    IF NOT EXISTS urls (
        "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        "shortId" TEXT,
        "url" TEXT
    );