package db

import (
	"log"
)

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
		return Nfts{}, err
	}

	if err := db.Where("name LIKE ?", "%"+name+"%").Find(&nfts).Error; err != nil {
		return Nfts{}, err
	}

	return nfts, nil
}

// GetNFTsByID retrieves an NFT by its ID from the database.
//
// The function takes a single parameter `id` which is the ID of the NFT to retrieve.
// It returns the NFT with the specified ID. If there is an error during the database query,
// it returns an error.
//
// Returns:
// - A `db.Nfts` containing the NFT with the specified ID.
// - An `error` if the database query fails.
func GetNFTsByID(id uint) (Nfts, error) {
	var nfts Nfts

	db, err := ConnectDB()
	if err != nil {
		return Nfts{}, err
	}

	if err := db.Where("id= ?", id).Find(&nfts).Error; err != nil {
		return Nfts{}, err
	}

	return nfts, nil
}

// DeleteNFT deletes an NFT by its ID from the database.
//
// The function takes a single parameter `id` which is the ID of the NFT to delete.
// It logs the deletion operation and returns an error if the deletion fails.
//
// Returns:
// - An `error` if the deletion fails.
func DeleteNFT(id string) error {
	var nfts Nfts

	db, err := ConnectDB()
	if err != nil {
		log.Printf("Failed to connect to database: %v", err)
		return err
	}

	if err := db.Where("id = ?", id).Delete(&nfts).Error; err != nil {
		log.Fatalf("Failed to delete NFT: %v", err)
	}

	return nil
}

// GetAllNFTs retrieves all NFTs from the database.
//
// The function returns a list of NFTs and an error. If there is an error during the database query,
// it returns an error.
//
// Returns:
// - A `[]db.Nfts` containing the list of NFTs.
// - An `error` if the database query fails.
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
