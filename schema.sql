CREATE TABLE IF NOT EXISTS stock
(
    id      INTEGER PRIMARY KEY,
    symbol  TEXT NOT NULL UNIQUE,
    company TEXT NOT NULL
) STRICT;

CREATE TABLE IF NOT EXISTS stock_price
(
    id             INTEGER PRIMARY KEY,
    stock_id       INTEGER NOT NULL,
    date           TEXT    NOT NULL,
    open           TEXT    NOT NULL,
    high           TEXT    NOT NULL,
    low            TEXT    NOT NULL,
    close          TEXT    NOT NULL,
    adjusted_close TEXT    NOT NULL,
    volume         TEXT    NOT NULL,

    FOREIGN KEY (stock_id) REFERENCES stock (id),
    UNIQUE (stock_id, date)
) STRICT;