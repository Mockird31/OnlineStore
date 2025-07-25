CREATE TABLE IF NOT EXISTS category (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title TEXT NOT NULL UNIQUE,
    CONSTRAINT title_length CHECK (LENGTH(title) > 0 AND LENGTH(title) <= 50),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS item (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title TEXT NOT NULL,
    CONSTRAINT title_length CHECK (LENGTH(title) > 0 AND LENGTH(title) <= 255),
    description TEXT,
    CONSTRAINT description_length CHECK (LENGTH(description) <= 1000),
    price DECIMAL(10, 2) NOT NULL CHECK (price >= 0),
    image_url TEXT NOT NULL DEFAULT '/default_item_image.png',
    quantity BIGINT NOT NULL,
    is_active BOOLEAN NOT NULL DEFAULT TRUE, 
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS item_category (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    item_id BIGINT NOT NULL,
    category_id BIGINT NOT NULL,
    FOREIGN KEY (item_id) 
        REFERENCES item (id) 
        ON DELETE CASCADE 
        ON UPDATE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    CONSTRAINT unique_item_category_check UNIQUE (item_id, category_id)
);

---- create above / drop below ----

DROP TABLE IF EXISTS item;
DROP TABLE IF EXISTS category;