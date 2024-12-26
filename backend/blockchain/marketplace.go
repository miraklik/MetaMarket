// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package marketplace

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// MarketplaceListing is an auto generated low-level Go binding around an user-defined struct.
type MarketplaceListing struct {
	Seller   common.Address
	TokenId  *big.Int
	Price    *big.Int
	IsActive bool
}

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nftContractAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_commissionPercent\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPercent\",\"type\":\"uint256\"}],\"name\":\"CommissionUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"FundsWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"}],\"name\":\"ListingCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ListingCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"PurchaseCompleted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MAX_COMMISSION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"}],\"name\":\"cancelListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commissionPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"_tokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"_price\",\"type\":\"uint128\"}],\"name\":\"createListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getListingId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"getListingsBySeller\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"tokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structMarketplace.Listing[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"isTokenListed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"uint128\",\"name\":\"tokenId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"price\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftContract\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftEnumerable\",\"outputs\":[{\"internalType\":\"contractIERC721Enumerable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingWithdrawals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"}],\"name\":\"purchaseListing\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPercent\",\"type\":\"uint256\"}],\"name\":\"setCommissionPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// MarketplaceABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketplaceMetaData.ABI instead.
var MarketplaceABI = MarketplaceMetaData.ABI

// Marketplace is an auto generated Go binding around an Ethereum contract.
type Marketplace struct {
	MarketplaceCaller     // Read-only binding to the contract
	MarketplaceTransactor // Write-only binding to the contract
	MarketplaceFilterer   // Log filterer for contract events
}

// MarketplaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketplaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketplaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketplaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketplaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketplaceSession struct {
	Contract     *Marketplace      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketplaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketplaceCallerSession struct {
	Contract *MarketplaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// MarketplaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketplaceTransactorSession struct {
	Contract     *MarketplaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// MarketplaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketplaceRaw struct {
	Contract *Marketplace // Generic contract binding to access the raw methods on
}

// MarketplaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketplaceCallerRaw struct {
	Contract *MarketplaceCaller // Generic read-only contract binding to access the raw methods on
}

// MarketplaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketplaceTransactorRaw struct {
	Contract *MarketplaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarketplace creates a new instance of Marketplace, bound to a specific deployed contract.
func NewMarketplace(address common.Address, backend bind.ContractBackend) (*Marketplace, error) {
	contract, err := bindMarketplace(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Marketplace{MarketplaceCaller: MarketplaceCaller{contract: contract}, MarketplaceTransactor: MarketplaceTransactor{contract: contract}, MarketplaceFilterer: MarketplaceFilterer{contract: contract}}, nil
}

// NewMarketplaceCaller creates a new read-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceCaller(address common.Address, caller bind.ContractCaller) (*MarketplaceCaller, error) {
	contract, err := bindMarketplace(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceCaller{contract: contract}, nil
}

// NewMarketplaceTransactor creates a new write-only instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketplaceTransactor, error) {
	contract, err := bindMarketplace(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketplaceTransactor{contract: contract}, nil
}

// NewMarketplaceFilterer creates a new log filterer instance of Marketplace, bound to a specific deployed contract.
func NewMarketplaceFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketplaceFilterer, error) {
	contract, err := bindMarketplace(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFilterer{contract: contract}, nil
}

// bindMarketplace binds a generic wrapper to an already deployed contract.
func bindMarketplace(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := MarketplaceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.MarketplaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.MarketplaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Marketplace *MarketplaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Marketplace.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Marketplace *MarketplaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Marketplace *MarketplaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Marketplace.Contract.contract.Transact(opts, method, params...)
}

// MAXCOMMISSION is a free data retrieval call binding the contract method 0xae97dde8.
//
// Solidity: function MAX_COMMISSION() view returns(uint256)
func (_Marketplace *MarketplaceCaller) MAXCOMMISSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "MAX_COMMISSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCOMMISSION is a free data retrieval call binding the contract method 0xae97dde8.
//
// Solidity: function MAX_COMMISSION() view returns(uint256)
func (_Marketplace *MarketplaceSession) MAXCOMMISSION() (*big.Int, error) {
	return _Marketplace.Contract.MAXCOMMISSION(&_Marketplace.CallOpts)
}

// MAXCOMMISSION is a free data retrieval call binding the contract method 0xae97dde8.
//
// Solidity: function MAX_COMMISSION() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) MAXCOMMISSION() (*big.Int, error) {
	return _Marketplace.Contract.MAXCOMMISSION(&_Marketplace.CallOpts)
}

// CommissionPercent is a free data retrieval call binding the contract method 0x77d3550b.
//
// Solidity: function commissionPercent() view returns(uint256)
func (_Marketplace *MarketplaceCaller) CommissionPercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "commissionPercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommissionPercent is a free data retrieval call binding the contract method 0x77d3550b.
//
// Solidity: function commissionPercent() view returns(uint256)
func (_Marketplace *MarketplaceSession) CommissionPercent() (*big.Int, error) {
	return _Marketplace.Contract.CommissionPercent(&_Marketplace.CallOpts)
}

// CommissionPercent is a free data retrieval call binding the contract method 0x77d3550b.
//
// Solidity: function commissionPercent() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) CommissionPercent() (*big.Int, error) {
	return _Marketplace.Contract.CommissionPercent(&_Marketplace.CallOpts)
}

// GetListingId is a free data retrieval call binding the contract method 0xcc85a55b.
//
// Solidity: function getListingId(uint256 _tokenId) view returns(uint256)
func (_Marketplace *MarketplaceCaller) GetListingId(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getListingId", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetListingId is a free data retrieval call binding the contract method 0xcc85a55b.
//
// Solidity: function getListingId(uint256 _tokenId) view returns(uint256)
func (_Marketplace *MarketplaceSession) GetListingId(_tokenId *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.GetListingId(&_Marketplace.CallOpts, _tokenId)
}

// GetListingId is a free data retrieval call binding the contract method 0xcc85a55b.
//
// Solidity: function getListingId(uint256 _tokenId) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) GetListingId(_tokenId *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.GetListingId(&_Marketplace.CallOpts, _tokenId)
}

// GetListingsBySeller is a free data retrieval call binding the contract method 0xd8cba251.
//
// Solidity: function getListingsBySeller(address _seller) view returns((address,uint128,uint128,bool)[])
func (_Marketplace *MarketplaceCaller) GetListingsBySeller(opts *bind.CallOpts, _seller common.Address) ([]MarketplaceListing, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "getListingsBySeller", _seller)

	if err != nil {
		return *new([]MarketplaceListing), err
	}

	out0 := *abi.ConvertType(out[0], new([]MarketplaceListing)).(*[]MarketplaceListing)

	return out0, err

}

// GetListingsBySeller is a free data retrieval call binding the contract method 0xd8cba251.
//
// Solidity: function getListingsBySeller(address _seller) view returns((address,uint128,uint128,bool)[])
func (_Marketplace *MarketplaceSession) GetListingsBySeller(_seller common.Address) ([]MarketplaceListing, error) {
	return _Marketplace.Contract.GetListingsBySeller(&_Marketplace.CallOpts, _seller)
}

// GetListingsBySeller is a free data retrieval call binding the contract method 0xd8cba251.
//
// Solidity: function getListingsBySeller(address _seller) view returns((address,uint128,uint128,bool)[])
func (_Marketplace *MarketplaceCallerSession) GetListingsBySeller(_seller common.Address) ([]MarketplaceListing, error) {
	return _Marketplace.Contract.GetListingsBySeller(&_Marketplace.CallOpts, _seller)
}

// IsTokenListed is a free data retrieval call binding the contract method 0x9b83cddc.
//
// Solidity: function isTokenListed(uint256 _tokenId) view returns(bool)
func (_Marketplace *MarketplaceCaller) IsTokenListed(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "isTokenListed", _tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsTokenListed is a free data retrieval call binding the contract method 0x9b83cddc.
//
// Solidity: function isTokenListed(uint256 _tokenId) view returns(bool)
func (_Marketplace *MarketplaceSession) IsTokenListed(_tokenId *big.Int) (bool, error) {
	return _Marketplace.Contract.IsTokenListed(&_Marketplace.CallOpts, _tokenId)
}

// IsTokenListed is a free data retrieval call binding the contract method 0x9b83cddc.
//
// Solidity: function isTokenListed(uint256 _tokenId) view returns(bool)
func (_Marketplace *MarketplaceCallerSession) IsTokenListed(_tokenId *big.Int) (bool, error) {
	return _Marketplace.Contract.IsTokenListed(&_Marketplace.CallOpts, _tokenId)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, uint128 tokenId, uint128 price, bool isActive)
func (_Marketplace *MarketplaceCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller   common.Address
	TokenId  *big.Int
	Price    *big.Int
	IsActive bool
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		Seller   common.Address
		TokenId  *big.Int
		Price    *big.Int
		IsActive bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.IsActive = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, uint128 tokenId, uint128 price, bool isActive)
func (_Marketplace *MarketplaceSession) Listings(arg0 *big.Int) (struct {
	Seller   common.Address
	TokenId  *big.Int
	Price    *big.Int
	IsActive bool
}, error) {
	return _Marketplace.Contract.Listings(&_Marketplace.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(address seller, uint128 tokenId, uint128 price, bool isActive)
func (_Marketplace *MarketplaceCallerSession) Listings(arg0 *big.Int) (struct {
	Seller   common.Address
	TokenId  *big.Int
	Price    *big.Int
	IsActive bool
}, error) {
	return _Marketplace.Contract.Listings(&_Marketplace.CallOpts, arg0)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_Marketplace *MarketplaceCaller) NftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "nftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_Marketplace *MarketplaceSession) NftContract() (common.Address, error) {
	return _Marketplace.Contract.NftContract(&_Marketplace.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_Marketplace *MarketplaceCallerSession) NftContract() (common.Address, error) {
	return _Marketplace.Contract.NftContract(&_Marketplace.CallOpts)
}

// NftEnumerable is a free data retrieval call binding the contract method 0x51e61b25.
//
// Solidity: function nftEnumerable() view returns(address)
func (_Marketplace *MarketplaceCaller) NftEnumerable(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "nftEnumerable")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftEnumerable is a free data retrieval call binding the contract method 0x51e61b25.
//
// Solidity: function nftEnumerable() view returns(address)
func (_Marketplace *MarketplaceSession) NftEnumerable() (common.Address, error) {
	return _Marketplace.Contract.NftEnumerable(&_Marketplace.CallOpts)
}

// NftEnumerable is a free data retrieval call binding the contract method 0x51e61b25.
//
// Solidity: function nftEnumerable() view returns(address)
func (_Marketplace *MarketplaceCallerSession) NftEnumerable() (common.Address, error) {
	return _Marketplace.Contract.NftEnumerable(&_Marketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Marketplace *MarketplaceCallerSession) Owner() (common.Address, error) {
	return _Marketplace.Contract.Owner(&_Marketplace.CallOpts)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) view returns(uint256)
func (_Marketplace *MarketplaceCaller) PendingWithdrawals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "pendingWithdrawals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) view returns(uint256)
func (_Marketplace *MarketplaceSession) PendingWithdrawals(arg0 common.Address) (*big.Int, error) {
	return _Marketplace.Contract.PendingWithdrawals(&_Marketplace.CallOpts, arg0)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) PendingWithdrawals(arg0 common.Address) (*big.Int, error) {
	return _Marketplace.Contract.PendingWithdrawals(&_Marketplace.CallOpts, arg0)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 index) view returns(uint256)
func (_Marketplace *MarketplaceCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "tokenOfOwnerByIndex", _owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 index) view returns(uint256)
func (_Marketplace *MarketplaceSession) TokenOfOwnerByIndex(_owner common.Address, index *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.TokenOfOwnerByIndex(&_Marketplace.CallOpts, _owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address _owner, uint256 index) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) TokenOfOwnerByIndex(_owner common.Address, index *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.TokenOfOwnerByIndex(&_Marketplace.CallOpts, _owner, index)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 _listingId) returns()
func (_Marketplace *MarketplaceTransactor) CancelListing(opts *bind.TransactOpts, _listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelListing", _listingId)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 _listingId) returns()
func (_Marketplace *MarketplaceSession) CancelListing(_listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelListing(&_Marketplace.TransactOpts, _listingId)
}

// CancelListing is a paid mutator transaction binding the contract method 0x305a67a8.
//
// Solidity: function cancelListing(uint256 _listingId) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelListing(_listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelListing(&_Marketplace.TransactOpts, _listingId)
}

// CreateListing is a paid mutator transaction binding the contract method 0x9acaa6fb.
//
// Solidity: function createListing(uint128 _tokenId, uint128 _price) returns()
func (_Marketplace *MarketplaceTransactor) CreateListing(opts *bind.TransactOpts, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "createListing", _tokenId, _price)
}

// CreateListing is a paid mutator transaction binding the contract method 0x9acaa6fb.
//
// Solidity: function createListing(uint128 _tokenId, uint128 _price) returns()
func (_Marketplace *MarketplaceSession) CreateListing(_tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateListing(&_Marketplace.TransactOpts, _tokenId, _price)
}

// CreateListing is a paid mutator transaction binding the contract method 0x9acaa6fb.
//
// Solidity: function createListing(uint128 _tokenId, uint128 _price) returns()
func (_Marketplace *MarketplaceTransactorSession) CreateListing(_tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateListing(&_Marketplace.TransactOpts, _tokenId, _price)
}

// PurchaseListing is a paid mutator transaction binding the contract method 0x169d5a7d.
//
// Solidity: function purchaseListing(uint256 _listingId) payable returns()
func (_Marketplace *MarketplaceTransactor) PurchaseListing(opts *bind.TransactOpts, _listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "purchaseListing", _listingId)
}

// PurchaseListing is a paid mutator transaction binding the contract method 0x169d5a7d.
//
// Solidity: function purchaseListing(uint256 _listingId) payable returns()
func (_Marketplace *MarketplaceSession) PurchaseListing(_listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PurchaseListing(&_Marketplace.TransactOpts, _listingId)
}

// PurchaseListing is a paid mutator transaction binding the contract method 0x169d5a7d.
//
// Solidity: function purchaseListing(uint256 _listingId) payable returns()
func (_Marketplace *MarketplaceTransactorSession) PurchaseListing(_listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PurchaseListing(&_Marketplace.TransactOpts, _listingId)
}

// SetCommissionPercent is a paid mutator transaction binding the contract method 0x404a9ab8.
//
// Solidity: function setCommissionPercent(uint256 _newPercent) returns()
func (_Marketplace *MarketplaceTransactor) SetCommissionPercent(opts *bind.TransactOpts, _newPercent *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "setCommissionPercent", _newPercent)
}

// SetCommissionPercent is a paid mutator transaction binding the contract method 0x404a9ab8.
//
// Solidity: function setCommissionPercent(uint256 _newPercent) returns()
func (_Marketplace *MarketplaceSession) SetCommissionPercent(_newPercent *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetCommissionPercent(&_Marketplace.TransactOpts, _newPercent)
}

// SetCommissionPercent is a paid mutator transaction binding the contract method 0x404a9ab8.
//
// Solidity: function setCommissionPercent(uint256 _newPercent) returns()
func (_Marketplace *MarketplaceTransactorSession) SetCommissionPercent(_newPercent *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.SetCommissionPercent(&_Marketplace.TransactOpts, _newPercent)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_Marketplace *MarketplaceTransactor) WithdrawFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "withdrawFunds")
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_Marketplace *MarketplaceSession) WithdrawFunds() (*types.Transaction, error) {
	return _Marketplace.Contract.WithdrawFunds(&_Marketplace.TransactOpts)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_Marketplace *MarketplaceTransactorSession) WithdrawFunds() (*types.Transaction, error) {
	return _Marketplace.Contract.WithdrawFunds(&_Marketplace.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace *MarketplaceTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace *MarketplaceSession) Receive() (*types.Transaction, error) {
	return _Marketplace.Contract.Receive(&_Marketplace.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Marketplace *MarketplaceTransactorSession) Receive() (*types.Transaction, error) {
	return _Marketplace.Contract.Receive(&_Marketplace.TransactOpts)
}

// MarketplaceCommissionUpdatedIterator is returned from FilterCommissionUpdated and is used to iterate over the raw logs and unpacked data for CommissionUpdated events raised by the Marketplace contract.
type MarketplaceCommissionUpdatedIterator struct {
	Event *MarketplaceCommissionUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceCommissionUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceCommissionUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceCommissionUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceCommissionUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceCommissionUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceCommissionUpdated represents a CommissionUpdated event raised by the Marketplace contract.
type MarketplaceCommissionUpdated struct {
	NewPercent *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterCommissionUpdated is a free log retrieval operation binding the contract event 0x13f60dd2b588490608c3ff1932a3daeb4087f3d5af04b97e5c2a16b5b4c0302e.
//
// Solidity: event CommissionUpdated(uint256 newPercent)
func (_Marketplace *MarketplaceFilterer) FilterCommissionUpdated(opts *bind.FilterOpts) (*MarketplaceCommissionUpdatedIterator, error) {

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "CommissionUpdated")
	if err != nil {
		return nil, err
	}
	return &MarketplaceCommissionUpdatedIterator{contract: _Marketplace.contract, event: "CommissionUpdated", logs: logs, sub: sub}, nil
}

// WatchCommissionUpdated is a free log subscription operation binding the contract event 0x13f60dd2b588490608c3ff1932a3daeb4087f3d5af04b97e5c2a16b5b4c0302e.
//
// Solidity: event CommissionUpdated(uint256 newPercent)
func (_Marketplace *MarketplaceFilterer) WatchCommissionUpdated(opts *bind.WatchOpts, sink chan<- *MarketplaceCommissionUpdated) (event.Subscription, error) {

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "CommissionUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceCommissionUpdated)
				if err := _Marketplace.contract.UnpackLog(event, "CommissionUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCommissionUpdated is a log parse operation binding the contract event 0x13f60dd2b588490608c3ff1932a3daeb4087f3d5af04b97e5c2a16b5b4c0302e.
//
// Solidity: event CommissionUpdated(uint256 newPercent)
func (_Marketplace *MarketplaceFilterer) ParseCommissionUpdated(log types.Log) (*MarketplaceCommissionUpdated, error) {
	event := new(MarketplaceCommissionUpdated)
	if err := _Marketplace.contract.UnpackLog(event, "CommissionUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceFundsWithdrawnIterator is returned from FilterFundsWithdrawn and is used to iterate over the raw logs and unpacked data for FundsWithdrawn events raised by the Marketplace contract.
type MarketplaceFundsWithdrawnIterator struct {
	Event *MarketplaceFundsWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceFundsWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceFundsWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceFundsWithdrawn represents a FundsWithdrawn event raised by the Marketplace contract.
type MarketplaceFundsWithdrawn struct {
	Amount    *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFundsWithdrawn is a free log retrieval operation binding the contract event 0x6141b54b56b8a52a8c6f5cd2a857f6117b18ffbf4d46bd3106f300a839cbf5ea.
//
// Solidity: event FundsWithdrawn(uint256 amount, address indexed recipient)
func (_Marketplace *MarketplaceFilterer) FilterFundsWithdrawn(opts *bind.FilterOpts, recipient []common.Address) (*MarketplaceFundsWithdrawnIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "FundsWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceFundsWithdrawnIterator{contract: _Marketplace.contract, event: "FundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchFundsWithdrawn is a free log subscription operation binding the contract event 0x6141b54b56b8a52a8c6f5cd2a857f6117b18ffbf4d46bd3106f300a839cbf5ea.
//
// Solidity: event FundsWithdrawn(uint256 amount, address indexed recipient)
func (_Marketplace *MarketplaceFilterer) WatchFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *MarketplaceFundsWithdrawn, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "FundsWithdrawn", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceFundsWithdrawn)
				if err := _Marketplace.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFundsWithdrawn is a log parse operation binding the contract event 0x6141b54b56b8a52a8c6f5cd2a857f6117b18ffbf4d46bd3106f300a839cbf5ea.
//
// Solidity: event FundsWithdrawn(uint256 amount, address indexed recipient)
func (_Marketplace *MarketplaceFilterer) ParseFundsWithdrawn(log types.Log) (*MarketplaceFundsWithdrawn, error) {
	event := new(MarketplaceFundsWithdrawn)
	if err := _Marketplace.contract.UnpackLog(event, "FundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceListingCancelledIterator is returned from FilterListingCancelled and is used to iterate over the raw logs and unpacked data for ListingCancelled events raised by the Marketplace contract.
type MarketplaceListingCancelledIterator struct {
	Event *MarketplaceListingCancelled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceListingCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceListingCancelled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceListingCancelled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceListingCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceListingCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceListingCancelled represents a ListingCancelled event raised by the Marketplace contract.
type MarketplaceListingCancelled struct {
	Id     *big.Int
	Seller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterListingCancelled is a free log retrieval operation binding the contract event 0x8e25282255ab31897df2b0456bb993ac7f84d376861aefd84901d2d63a7428a2.
//
// Solidity: event ListingCancelled(uint256 indexed id, address indexed seller)
func (_Marketplace *MarketplaceFilterer) FilterListingCancelled(opts *bind.FilterOpts, id []*big.Int, seller []common.Address) (*MarketplaceListingCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "ListingCancelled", idRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceListingCancelledIterator{contract: _Marketplace.contract, event: "ListingCancelled", logs: logs, sub: sub}, nil
}

// WatchListingCancelled is a free log subscription operation binding the contract event 0x8e25282255ab31897df2b0456bb993ac7f84d376861aefd84901d2d63a7428a2.
//
// Solidity: event ListingCancelled(uint256 indexed id, address indexed seller)
func (_Marketplace *MarketplaceFilterer) WatchListingCancelled(opts *bind.WatchOpts, sink chan<- *MarketplaceListingCancelled, id []*big.Int, seller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "ListingCancelled", idRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceListingCancelled)
				if err := _Marketplace.contract.UnpackLog(event, "ListingCancelled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseListingCancelled is a log parse operation binding the contract event 0x8e25282255ab31897df2b0456bb993ac7f84d376861aefd84901d2d63a7428a2.
//
// Solidity: event ListingCancelled(uint256 indexed id, address indexed seller)
func (_Marketplace *MarketplaceFilterer) ParseListingCancelled(log types.Log) (*MarketplaceListingCancelled, error) {
	event := new(MarketplaceListingCancelled)
	if err := _Marketplace.contract.UnpackLog(event, "ListingCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplaceListingCreatedIterator is returned from FilterListingCreated and is used to iterate over the raw logs and unpacked data for ListingCreated events raised by the Marketplace contract.
type MarketplaceListingCreatedIterator struct {
	Event *MarketplaceListingCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplaceListingCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceListingCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplaceListingCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplaceListingCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceListingCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceListingCreated represents a ListingCreated event raised by the Marketplace contract.
type MarketplaceListingCreated struct {
	Id        *big.Int
	Seller    common.Address
	TokenId   *big.Int
	Price     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterListingCreated is a free log retrieval operation binding the contract event 0xa3941cfa07eecf2734875ab71cea0e47395be644da6e17b039a7e29efba74ef9.
//
// Solidity: event ListingCreated(uint256 indexed id, address indexed seller, uint256 tokenId, uint256 price, uint256 timestamp)
func (_Marketplace *MarketplaceFilterer) FilterListingCreated(opts *bind.FilterOpts, id []*big.Int, seller []common.Address) (*MarketplaceListingCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "ListingCreated", idRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceListingCreatedIterator{contract: _Marketplace.contract, event: "ListingCreated", logs: logs, sub: sub}, nil
}

// WatchListingCreated is a free log subscription operation binding the contract event 0xa3941cfa07eecf2734875ab71cea0e47395be644da6e17b039a7e29efba74ef9.
//
// Solidity: event ListingCreated(uint256 indexed id, address indexed seller, uint256 tokenId, uint256 price, uint256 timestamp)
func (_Marketplace *MarketplaceFilterer) WatchListingCreated(opts *bind.WatchOpts, sink chan<- *MarketplaceListingCreated, id []*big.Int, seller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "ListingCreated", idRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceListingCreated)
				if err := _Marketplace.contract.UnpackLog(event, "ListingCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseListingCreated is a log parse operation binding the contract event 0xa3941cfa07eecf2734875ab71cea0e47395be644da6e17b039a7e29efba74ef9.
//
// Solidity: event ListingCreated(uint256 indexed id, address indexed seller, uint256 tokenId, uint256 price, uint256 timestamp)
func (_Marketplace *MarketplaceFilterer) ParseListingCreated(log types.Log) (*MarketplaceListingCreated, error) {
	event := new(MarketplaceListingCreated)
	if err := _Marketplace.contract.UnpackLog(event, "ListingCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplacePurchaseCompletedIterator is returned from FilterPurchaseCompleted and is used to iterate over the raw logs and unpacked data for PurchaseCompleted events raised by the Marketplace contract.
type MarketplacePurchaseCompletedIterator struct {
	Event *MarketplacePurchaseCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MarketplacePurchaseCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacePurchaseCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MarketplacePurchaseCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MarketplacePurchaseCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacePurchaseCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacePurchaseCompleted represents a PurchaseCompleted event raised by the Marketplace contract.
type MarketplacePurchaseCompleted struct {
	Id        *big.Int
	Buyer     common.Address
	TokenId   *big.Int
	Price     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPurchaseCompleted is a free log retrieval operation binding the contract event 0xe29b93709a7ac5e91a6057482d1df985483725988417e08dec23377234003c80.
//
// Solidity: event PurchaseCompleted(uint256 indexed id, address indexed buyer, uint256 tokenId, uint256 price, uint256 timestamp)
func (_Marketplace *MarketplaceFilterer) FilterPurchaseCompleted(opts *bind.FilterOpts, id []*big.Int, buyer []common.Address) (*MarketplacePurchaseCompletedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "PurchaseCompleted", idRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplacePurchaseCompletedIterator{contract: _Marketplace.contract, event: "PurchaseCompleted", logs: logs, sub: sub}, nil
}

// WatchPurchaseCompleted is a free log subscription operation binding the contract event 0xe29b93709a7ac5e91a6057482d1df985483725988417e08dec23377234003c80.
//
// Solidity: event PurchaseCompleted(uint256 indexed id, address indexed buyer, uint256 tokenId, uint256 price, uint256 timestamp)
func (_Marketplace *MarketplaceFilterer) WatchPurchaseCompleted(opts *bind.WatchOpts, sink chan<- *MarketplacePurchaseCompleted, id []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "PurchaseCompleted", idRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacePurchaseCompleted)
				if err := _Marketplace.contract.UnpackLog(event, "PurchaseCompleted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePurchaseCompleted is a log parse operation binding the contract event 0xe29b93709a7ac5e91a6057482d1df985483725988417e08dec23377234003c80.
//
// Solidity: event PurchaseCompleted(uint256 indexed id, address indexed buyer, uint256 tokenId, uint256 price, uint256 timestamp)
func (_Marketplace *MarketplaceFilterer) ParsePurchaseCompleted(log types.Log) (*MarketplacePurchaseCompleted, error) {
	event := new(MarketplacePurchaseCompleted)
	if err := _Marketplace.contract.UnpackLog(event, "PurchaseCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
