CREATE TABLE item (
    id          BIGSERIAL PRIMARY KEY,
    name        TEXT NOT NULL
    --
);

CREATE TABLE item_located (
    item        BIGSERIAL PRIMARY KEY,
    located     BIGSERIAL NOT NULL
    --
);

CREATE TABLE locate_abstract (
    id          BIGSERIAL PRIMARY KEY,
    locate_kind BIGSERIAL NOT NULL,
    locate_real BIGSERIAL NOT NULL
    --
);
