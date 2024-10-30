package blockchain

import (
	"encoding/hex"
	"fmt"
	"os"
	"runtime"

	"github.com/dgraph-io/badger"
)

const (
	dbPath      = "./tmp/blocks"
	dbFile      = "./tmp/blocks/MANIFEST"
	genesisData = "First Transaction From Genesis"
)

type BlockChain struct {
	LastHash []byte
	DataBase *badger.DB
}

type BlockChainInterator struct {
	CurrentHash []byte
	DataBase    *badger.DB
}

func DBexists() bool {
	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		return false
	}

	return true
}

func (chain *BlockChain) AddBlock(transaction []*Transaction) {
	var lasthash []byte

	err := chain.DataBase.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			lasthash = append([]byte{}, val...)
			return nil
		})
		return err
	})

	Handle(err)

	newBlock := CreateBlock(transaction, lasthash)

	err = chain.DataBase.Update(func(txn *badger.Txn) error {
		err := txn.Set(newBlock.Hash, newBlock.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), newBlock.Hash)

		chain.LastHash = newBlock.Hash

		return err
	})
	Handle(err)
}

func InitBlockChain(address string) *BlockChain {
	var lasthash []byte

	if DBexists() {
		fmt.Println("Blockchain already exists")
		runtime.Goexit()
	}

	opts := badger.DefaultOptions("")
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		cbtx := CoinbaseTx(address, genesisData)
		genesis := Genesis(cbtx)
		fmt.Println("Genesis proved")
		err = txn.Set(genesis.Hash, genesis.Serialize())
		Handle(err)
		err = txn.Set([]byte("lh"), genesis.Hash)

		lasthash = genesis.Hash

		return err
	})

	Handle(err)
	blockchain := BlockChain{lasthash, db}
	return &blockchain
}

func ContinueBlockChain(address string) *BlockChain {
	if DBexists() == false {
		fmt.Println("No existing blockchain found, create one!")
		runtime.Goexit()
	}

	var lastHash []byte

	opts := badger.DefaultOptions("")
	opts.Dir = dbPath
	opts.ValueDir = dbPath

	db, err := badger.Open(opts)
	Handle(err)

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		Handle(err)
		err = item.Value(func(val []byte) error {
			lastHash = val
			return nil
		})
		return err
	})
	Handle(err)

	chain := BlockChain{lastHash, db}

	return &chain
}

func (chain *BlockChain) Interator() *BlockChainInterator {
	iter := &BlockChainInterator{chain.LastHash, chain.DataBase}

	return iter
}

func (iter *BlockChainInterator) Next() *Block {
	var block *Block

	err := iter.DataBase.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		if err != nil {
			return err
		}

		var encodedBlock []byte
		encodedBlock, err = item.ValueCopy(nil)
		if err != nil {
			return err
		}

		block = Deserialize(encodedBlock)
		return nil
	})

	Handle(err)

	iter.CurrentHash = block.PrevHash

	return block
}

func (chain *BlockChain) FindUnspentTransaction(address string) []Transaction {
	var (
		unspentTxs []Transaction
	)

	spentTXOs := make(map[string][]int)

	iter := chain.Interator()

	for {
		block := iter.Next()

		for _, tx := range block.Transaction {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for outIdx, out := range tx.Output {
				if spentTXOs[txID] != nil {
					for _, spentOut := range spentTXOs[txID] {
						if spentOut == outIdx {
							continue Outputs
						}
					}
				}
				if out.CanBeUnlocked(address) {
					unspentTxs = append(unspentTxs, *tx)
				}
			}
			if tx.IsCoinbase() == false {
				for _, in := range tx.Input {
					if in.CanUnlock(address) {
						inTxID := hex.EncodeToString(in.ID)
						spentTXOs[inTxID] = append(spentTXOs[inTxID], in.Out)
					}
				}
			}
		}

		if len(block.PrevHash) == 0 {
			break
		}
	}

	return unspentTxs
}

func (chain *BlockChain) FindUTXO(address string) []TxOutput {
	var UTXOs []TxOutput
	upsentTransactions := chain.FindUnspentTransaction(address)

	for _, tx := range upsentTransactions {
		for _, out := range tx.Output {
			if out.CanBeUnlocked(address) {
				UTXOs = append(UTXOs, out)
			}
		}
	}

	return UTXOs
}

func (chain *BlockChain) FindSpendableOutputs(address string, amount int) (int, map[string][]int) {
	unspentOuts := make(map[string][]int)
	upsentTxs := chain.FindUnspentTransaction(address)
	accamulated := 0

Work:
	for _, tx := range upsentTxs {
		txID := hex.EncodeToString(tx.ID)

		for outIdx, out := range tx.Output {
			if out.CanBeUnlocked(address) && accamulated < amount {
				accamulated += out.Value
				unspentOuts[txID] = append(unspentOuts[txID], outIdx)

				if accamulated >= amount {
					break Work
				}
			}
		}
	}

	return accamulated, unspentOuts
}
