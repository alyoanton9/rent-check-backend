CREATE TABLE flats (
    id INTEGER PRIMARY KEY,
    address VARCHAR NOT NULL UNIQUE,
    notes VARCHAR
);

CREATE TABLE tags (
    id INTEGER PRIMARY KEY,
    title VARCHAR NOT NULL UNIQUE
);

CREATE TABLE items (
    id INTEGER PRIMARY KEY,
    title VARCHAR NOT NULL UNIQUE,
    description VARCHAR,
    tag_id INTEGER REFERENCES tags (id)
);

CREATE TABLE flat_items (
    flat_id INTEGER REFERENCES flats (id),
    item_id INTEGER REFERENCES items (id),
    state BOOLEAN NOT NULL,
    PRIMARY KEY (flat_id, item_id)
);
