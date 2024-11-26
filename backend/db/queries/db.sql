CREATE TABLE nfts (
    id SERIAL PRIMARY KEY,
    token_id BIGINT NOT NULL,
    owner_address TEXT NOT NULL,
    nft_name TEXT NOT NULL,
    metadata_url TEXT NOT NULL,
    image_url TEXT NOT NULL,
    title TEXT,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE listings (
    id SERIAL PRIMARY KEY,
    nft_id BIGINT REFERENCES nfts(token_id),
    seller_address TEXT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    status TEXT CHECK (status IN ('active', 'sold', 'cancelled')) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    wallet_address TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
