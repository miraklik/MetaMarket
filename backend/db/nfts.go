package db

type Nfts struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Name        string `gorm:"size:255; not null; unique" json:"name"`
	Symbol      string `gorm:"size:255; not null; unique" json:"symbol"`
	Description string `gorm:"size:255; not null;" json:"description"`
	Price       string `gorm:"size:255; not null;" json:"price"`
}

// GetNFTsByName retrieves a list of NFTs with the given name from the database.
//
// The function takes a single parameter `name` which is the name of the NFT to search for.
// It returns a list of NFTs that match the given name. If there is an error during the database query, it returns an error.
//
// Returns:
// - A `db.Nfts` containing the list of NFTs with the specified name.
// - An `error` if the database query fails.
func GetNFTsByName(name string) (Nfts, error) {
	var nfts Nfts

	db, err := ConnectDB()
	if err != nil {
		return nfts, err
	}

	if err := db.Where("name LIKE ?", "%"+name+"%").Find(&nfts).Error; err != nil {
		return nfts, err
	}

	return nfts, nil
}

func GetAllNFTs() ([]Nfts, error) {
	var nfts []Nfts

	db, err := ConnectDB()
	if err != nil {
		return nfts, err
	}

	if err := db.Find(&nfts).Error; err != nil {
		return nfts, err
	}

	return nfts, nil
}
