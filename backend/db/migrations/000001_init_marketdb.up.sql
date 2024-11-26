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