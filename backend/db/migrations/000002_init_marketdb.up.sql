CREATE TABLE listings (
    id SERIAL PRIMARY KEY,
    nft_id BIGINT REFERENCES nfts(token_id),
    seller_address TEXT NOT NULL,
    price NUMERIC(10, 2) NOT NULL,
    status TEXT CHECK (status IN ('active', 'sold', 'cancelled')) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);