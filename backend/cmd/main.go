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

// MarketplaceMetaData contains all meta data concerning the Marketplace contract.
var MarketplaceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_commissionPercent\",\"type\":\"uint8\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"listingId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EscrowReleased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"ListingCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PurchaseCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"PurchaseCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"}],\"name\":\"cancelPurchase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"commissionPercent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"}],\"name\":\"confirmPurchase\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_imageIPFSHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"createListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"escrowAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"escrowBuyer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"listingCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"listings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"seller\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"imageIPFSHash\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isConfirmed\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_listingId\",\"type\":\"uint256\"}],\"name\":\"purchaseListing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newPercent\",\"type\":\"uint256\"}],\"name\":\"setCommissionPercent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdtToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// EscrowAmount is a free data retrieval call binding the contract method 0x5d19e56c.
//
// Solidity: function escrowAmount(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCaller) EscrowAmount(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "escrowAmount", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EscrowAmount is a free data retrieval call binding the contract method 0x5d19e56c.
//
// Solidity: function escrowAmount(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceSession) EscrowAmount(arg0 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.EscrowAmount(&_Marketplace.CallOpts, arg0)
}

// EscrowAmount is a free data retrieval call binding the contract method 0x5d19e56c.
//
// Solidity: function escrowAmount(uint256 ) view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) EscrowAmount(arg0 *big.Int) (*big.Int, error) {
	return _Marketplace.Contract.EscrowAmount(&_Marketplace.CallOpts, arg0)
}

// EscrowBuyer is a free data retrieval call binding the contract method 0x6c5d4bec.
//
// Solidity: function escrowBuyer(uint256 ) view returns(address)
func (_Marketplace *MarketplaceCaller) EscrowBuyer(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "escrowBuyer", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EscrowBuyer is a free data retrieval call binding the contract method 0x6c5d4bec.
//
// Solidity: function escrowBuyer(uint256 ) view returns(address)
func (_Marketplace *MarketplaceSession) EscrowBuyer(arg0 *big.Int) (common.Address, error) {
	return _Marketplace.Contract.EscrowBuyer(&_Marketplace.CallOpts, arg0)
}

// EscrowBuyer is a free data retrieval call binding the contract method 0x6c5d4bec.
//
// Solidity: function escrowBuyer(uint256 ) view returns(address)
func (_Marketplace *MarketplaceCallerSession) EscrowBuyer(arg0 *big.Int) (common.Address, error) {
	return _Marketplace.Contract.EscrowBuyer(&_Marketplace.CallOpts, arg0)
}

// ListingCount is a free data retrieval call binding the contract method 0xa9b07c26.
//
// Solidity: function listingCount() view returns(uint256)
func (_Marketplace *MarketplaceCaller) ListingCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "listingCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ListingCount is a free data retrieval call binding the contract method 0xa9b07c26.
//
// Solidity: function listingCount() view returns(uint256)
func (_Marketplace *MarketplaceSession) ListingCount() (*big.Int, error) {
	return _Marketplace.Contract.ListingCount(&_Marketplace.CallOpts)
}

// ListingCount is a free data retrieval call binding the contract method 0xa9b07c26.
//
// Solidity: function listingCount() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) ListingCount() (*big.Int, error) {
	return _Marketplace.Contract.ListingCount(&_Marketplace.CallOpts)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 id, address seller, string title, string description, string imageIPFSHash, uint256 price, bool sold)
func (_Marketplace *MarketplaceCaller) Listings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id            *big.Int
	Seller        common.Address
	Title         string
	Description   string
	ImageIPFSHash string
	Price         *big.Int
	Sold          bool
}, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "listings", arg0)

	outstruct := new(struct {
		Id            *big.Int
		Seller        common.Address
		Title         string
		Description   string
		ImageIPFSHash string
		Price         *big.Int
		Sold          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Title = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Description = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.ImageIPFSHash = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Price = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Sold = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 id, address seller, string title, string description, string imageIPFSHash, uint256 price, bool sold)
func (_Marketplace *MarketplaceSession) Listings(arg0 *big.Int) (struct {
	Id            *big.Int
	Seller        common.Address
	Title         string
	Description   string
	ImageIPFSHash string
	Price         *big.Int
	Sold          bool
}, error) {
	return _Marketplace.Contract.Listings(&_Marketplace.CallOpts, arg0)
}

// Listings is a free data retrieval call binding the contract method 0xde74e57b.
//
// Solidity: function listings(uint256 ) view returns(uint256 id, address seller, string title, string description, string imageIPFSHash, uint256 price, bool sold)
func (_Marketplace *MarketplaceCallerSession) Listings(arg0 *big.Int) (struct {
	Id            *big.Int
	Seller        common.Address
	Title         string
	Description   string
	ImageIPFSHash string
	Price         *big.Int
	Sold          bool
}, error) {
	return _Marketplace.Contract.Listings(&_Marketplace.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(bool isConfirmed)
func (_Marketplace *MarketplaceCaller) Orders(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "orders", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(bool isConfirmed)
func (_Marketplace *MarketplaceSession) Orders(arg0 *big.Int) (bool, error) {
	return _Marketplace.Contract.Orders(&_Marketplace.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0xa85c38ef.
//
// Solidity: function orders(uint256 ) view returns(bool isConfirmed)
func (_Marketplace *MarketplaceCallerSession) Orders(arg0 *big.Int) (bool, error) {
	return _Marketplace.Contract.Orders(&_Marketplace.CallOpts, arg0)
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

// TotalOrders is a free data retrieval call binding the contract method 0x1d834409.
//
// Solidity: function totalOrders() view returns(uint256)
func (_Marketplace *MarketplaceCaller) TotalOrders(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "totalOrders")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalOrders is a free data retrieval call binding the contract method 0x1d834409.
//
// Solidity: function totalOrders() view returns(uint256)
func (_Marketplace *MarketplaceSession) TotalOrders() (*big.Int, error) {
	return _Marketplace.Contract.TotalOrders(&_Marketplace.CallOpts)
}

// TotalOrders is a free data retrieval call binding the contract method 0x1d834409.
//
// Solidity: function totalOrders() view returns(uint256)
func (_Marketplace *MarketplaceCallerSession) TotalOrders() (*big.Int, error) {
	return _Marketplace.Contract.TotalOrders(&_Marketplace.CallOpts)
}

// UsdtToken is a free data retrieval call binding the contract method 0xa98ad46c.
//
// Solidity: function usdtToken() view returns(address)
func (_Marketplace *MarketplaceCaller) UsdtToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Marketplace.contract.Call(opts, &out, "usdtToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdtToken is a free data retrieval call binding the contract method 0xa98ad46c.
//
// Solidity: function usdtToken() view returns(address)
func (_Marketplace *MarketplaceSession) UsdtToken() (common.Address, error) {
	return _Marketplace.Contract.UsdtToken(&_Marketplace.CallOpts)
}

// UsdtToken is a free data retrieval call binding the contract method 0xa98ad46c.
//
// Solidity: function usdtToken() view returns(address)
func (_Marketplace *MarketplaceCallerSession) UsdtToken() (common.Address, error) {
	return _Marketplace.Contract.UsdtToken(&_Marketplace.CallOpts)
}

// CancelPurchase is a paid mutator transaction binding the contract method 0xc3634ddc.
//
// Solidity: function cancelPurchase(uint256 _listingId) returns()
func (_Marketplace *MarketplaceTransactor) CancelPurchase(opts *bind.TransactOpts, _listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "cancelPurchase", _listingId)
}

// CancelPurchase is a paid mutator transaction binding the contract method 0xc3634ddc.
//
// Solidity: function cancelPurchase(uint256 _listingId) returns()
func (_Marketplace *MarketplaceSession) CancelPurchase(_listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelPurchase(&_Marketplace.TransactOpts, _listingId)
}

// CancelPurchase is a paid mutator transaction binding the contract method 0xc3634ddc.
//
// Solidity: function cancelPurchase(uint256 _listingId) returns()
func (_Marketplace *MarketplaceTransactorSession) CancelPurchase(_listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CancelPurchase(&_Marketplace.TransactOpts, _listingId)
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0x4d24e902.
//
// Solidity: function confirmPurchase(uint256 _listingId) returns()
func (_Marketplace *MarketplaceTransactor) ConfirmPurchase(opts *bind.TransactOpts, _listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "confirmPurchase", _listingId)
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0x4d24e902.
//
// Solidity: function confirmPurchase(uint256 _listingId) returns()
func (_Marketplace *MarketplaceSession) ConfirmPurchase(_listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.ConfirmPurchase(&_Marketplace.TransactOpts, _listingId)
}

// ConfirmPurchase is a paid mutator transaction binding the contract method 0x4d24e902.
//
// Solidity: function confirmPurchase(uint256 _listingId) returns()
func (_Marketplace *MarketplaceTransactorSession) ConfirmPurchase(_listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.ConfirmPurchase(&_Marketplace.TransactOpts, _listingId)
}

// CreateListing is a paid mutator transaction binding the contract method 0x451d4107.
//
// Solidity: function createListing(string _title, string _description, string _imageIPFSHash, uint256 _price) returns()
func (_Marketplace *MarketplaceTransactor) CreateListing(opts *bind.TransactOpts, _title string, _description string, _imageIPFSHash string, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "createListing", _title, _description, _imageIPFSHash, _price)
}

// CreateListing is a paid mutator transaction binding the contract method 0x451d4107.
//
// Solidity: function createListing(string _title, string _description, string _imageIPFSHash, uint256 _price) returns()
func (_Marketplace *MarketplaceSession) CreateListing(_title string, _description string, _imageIPFSHash string, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateListing(&_Marketplace.TransactOpts, _title, _description, _imageIPFSHash, _price)
}

// CreateListing is a paid mutator transaction binding the contract method 0x451d4107.
//
// Solidity: function createListing(string _title, string _description, string _imageIPFSHash, uint256 _price) returns()
func (_Marketplace *MarketplaceTransactorSession) CreateListing(_title string, _description string, _imageIPFSHash string, _price *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.CreateListing(&_Marketplace.TransactOpts, _title, _description, _imageIPFSHash, _price)
}

// PurchaseListing is a paid mutator transaction binding the contract method 0x169d5a7d.
//
// Solidity: function purchaseListing(uint256 _listingId) returns()
func (_Marketplace *MarketplaceTransactor) PurchaseListing(opts *bind.TransactOpts, _listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "purchaseListing", _listingId)
}

// PurchaseListing is a paid mutator transaction binding the contract method 0x169d5a7d.
//
// Solidity: function purchaseListing(uint256 _listingId) returns()
func (_Marketplace *MarketplaceSession) PurchaseListing(_listingId *big.Int) (*types.Transaction, error) {
	return _Marketplace.Contract.PurchaseListing(&_Marketplace.TransactOpts, _listingId)
}

// PurchaseListing is a paid mutator transaction binding the contract method 0x169d5a7d.
//
// Solidity: function purchaseListing(uint256 _listingId) returns()
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

// WithdrawToken is a paid mutator transaction binding the contract method 0xca628c78.
//
// Solidity: function withdrawToken() returns()
func (_Marketplace *MarketplaceTransactor) WithdrawToken(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Marketplace.contract.Transact(opts, "withdrawToken")
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xca628c78.
//
// Solidity: function withdrawToken() returns()
func (_Marketplace *MarketplaceSession) WithdrawToken() (*types.Transaction, error) {
	return _Marketplace.Contract.WithdrawToken(&_Marketplace.TransactOpts)
}

// WithdrawToken is a paid mutator transaction binding the contract method 0xca628c78.
//
// Solidity: function withdrawToken() returns()
func (_Marketplace *MarketplaceTransactorSession) WithdrawToken() (*types.Transaction, error) {
	return _Marketplace.Contract.WithdrawToken(&_Marketplace.TransactOpts)
}

// MarketplaceEscrowReleasedIterator is returned from FilterEscrowReleased and is used to iterate over the raw logs and unpacked data for EscrowReleased events raised by the Marketplace contract.
type MarketplaceEscrowReleasedIterator struct {
	Event *MarketplaceEscrowReleased // Event containing the contract specifics and raw log

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
func (it *MarketplaceEscrowReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplaceEscrowReleased)
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
		it.Event = new(MarketplaceEscrowReleased)
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
func (it *MarketplaceEscrowReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplaceEscrowReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplaceEscrowReleased represents a EscrowReleased event raised by the Marketplace contract.
type MarketplaceEscrowReleased struct {
	ListingId *big.Int
	Seller    common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEscrowReleased is a free log retrieval operation binding the contract event 0x6244ed823ca6be0f11bc890c3fafcf3c29cb23420c14243642e930b5e07e6d0a.
//
// Solidity: event EscrowReleased(uint256 indexed listingId, address indexed seller, uint256 amount)
func (_Marketplace *MarketplaceFilterer) FilterEscrowReleased(opts *bind.FilterOpts, listingId []*big.Int, seller []common.Address) (*MarketplaceEscrowReleasedIterator, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "EscrowReleased", listingIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplaceEscrowReleasedIterator{contract: _Marketplace.contract, event: "EscrowReleased", logs: logs, sub: sub}, nil
}

// WatchEscrowReleased is a free log subscription operation binding the contract event 0x6244ed823ca6be0f11bc890c3fafcf3c29cb23420c14243642e930b5e07e6d0a.
//
// Solidity: event EscrowReleased(uint256 indexed listingId, address indexed seller, uint256 amount)
func (_Marketplace *MarketplaceFilterer) WatchEscrowReleased(opts *bind.WatchOpts, sink chan<- *MarketplaceEscrowReleased, listingId []*big.Int, seller []common.Address) (event.Subscription, error) {

	var listingIdRule []interface{}
	for _, listingIdItem := range listingId {
		listingIdRule = append(listingIdRule, listingIdItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "EscrowReleased", listingIdRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplaceEscrowReleased)
				if err := _Marketplace.contract.UnpackLog(event, "EscrowReleased", log); err != nil {
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

// ParseEscrowReleased is a log parse operation binding the contract event 0x6244ed823ca6be0f11bc890c3fafcf3c29cb23420c14243642e930b5e07e6d0a.
//
// Solidity: event EscrowReleased(uint256 indexed listingId, address indexed seller, uint256 amount)
func (_Marketplace *MarketplaceFilterer) ParseEscrowReleased(log types.Log) (*MarketplaceEscrowReleased, error) {
	event := new(MarketplaceEscrowReleased)
	if err := _Marketplace.contract.UnpackLog(event, "EscrowReleased", log); err != nil {
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
	Id     *big.Int
	Seller common.Address
	Title  string
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterListingCreated is a free log retrieval operation binding the contract event 0x6b7a3194eaddce2be5200d10785a15a5cd14a78b827cf902f4d23c35e6bd73f8.
//
// Solidity: event ListingCreated(uint256 indexed id, address indexed seller, string title, uint256 price)
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

// WatchListingCreated is a free log subscription operation binding the contract event 0x6b7a3194eaddce2be5200d10785a15a5cd14a78b827cf902f4d23c35e6bd73f8.
//
// Solidity: event ListingCreated(uint256 indexed id, address indexed seller, string title, uint256 price)
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

// ParseListingCreated is a log parse operation binding the contract event 0x6b7a3194eaddce2be5200d10785a15a5cd14a78b827cf902f4d23c35e6bd73f8.
//
// Solidity: event ListingCreated(uint256 indexed id, address indexed seller, string title, uint256 price)
func (_Marketplace *MarketplaceFilterer) ParseListingCreated(log types.Log) (*MarketplaceListingCreated, error) {
	event := new(MarketplaceListingCreated)
	if err := _Marketplace.contract.UnpackLog(event, "ListingCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketplacePurchaseCancelledIterator is returned from FilterPurchaseCancelled and is used to iterate over the raw logs and unpacked data for PurchaseCancelled events raised by the Marketplace contract.
type MarketplacePurchaseCancelledIterator struct {
	Event *MarketplacePurchaseCancelled // Event containing the contract specifics and raw log

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
func (it *MarketplacePurchaseCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketplacePurchaseCancelled)
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
		it.Event = new(MarketplacePurchaseCancelled)
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
func (it *MarketplacePurchaseCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketplacePurchaseCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketplacePurchaseCancelled represents a PurchaseCancelled event raised by the Marketplace contract.
type MarketplacePurchaseCancelled struct {
	Id    *big.Int
	Buyer common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPurchaseCancelled is a free log retrieval operation binding the contract event 0x83559618a86943fdd7e8104a56312f188df6d0e149a9868b514871d05ac785fb.
//
// Solidity: event PurchaseCancelled(uint256 indexed id, address indexed buyer, uint256 price)
func (_Marketplace *MarketplaceFilterer) FilterPurchaseCancelled(opts *bind.FilterOpts, id []*big.Int, buyer []common.Address) (*MarketplacePurchaseCancelledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Marketplace.contract.FilterLogs(opts, "PurchaseCancelled", idRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &MarketplacePurchaseCancelledIterator{contract: _Marketplace.contract, event: "PurchaseCancelled", logs: logs, sub: sub}, nil
}

// WatchPurchaseCancelled is a free log subscription operation binding the contract event 0x83559618a86943fdd7e8104a56312f188df6d0e149a9868b514871d05ac785fb.
//
// Solidity: event PurchaseCancelled(uint256 indexed id, address indexed buyer, uint256 price)
func (_Marketplace *MarketplaceFilterer) WatchPurchaseCancelled(opts *bind.WatchOpts, sink chan<- *MarketplacePurchaseCancelled, id []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Marketplace.contract.WatchLogs(opts, "PurchaseCancelled", idRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketplacePurchaseCancelled)
				if err := _Marketplace.contract.UnpackLog(event, "PurchaseCancelled", log); err != nil {
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

// ParsePurchaseCancelled is a log parse operation binding the contract event 0x83559618a86943fdd7e8104a56312f188df6d0e149a9868b514871d05ac785fb.
//
// Solidity: event PurchaseCancelled(uint256 indexed id, address indexed buyer, uint256 price)
func (_Marketplace *MarketplaceFilterer) ParsePurchaseCancelled(log types.Log) (*MarketplacePurchaseCancelled, error) {
	event := new(MarketplacePurchaseCancelled)
	if err := _Marketplace.contract.UnpackLog(event, "PurchaseCancelled", log); err != nil {
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
	Id    *big.Int
	Buyer common.Address
	Price *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterPurchaseCompleted is a free log retrieval operation binding the contract event 0x3915b4ab4f4d3ad54079c20a1bb2159e875d3d5b673d4dc93986f956d2581189.
//
// Solidity: event PurchaseCompleted(uint256 indexed id, address indexed buyer, uint256 price)
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

// WatchPurchaseCompleted is a free log subscription operation binding the contract event 0x3915b4ab4f4d3ad54079c20a1bb2159e875d3d5b673d4dc93986f956d2581189.
//
// Solidity: event PurchaseCompleted(uint256 indexed id, address indexed buyer, uint256 price)
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

// ParsePurchaseCompleted is a log parse operation binding the contract event 0x3915b4ab4f4d3ad54079c20a1bb2159e875d3d5b673d4dc93986f956d2581189.
//
// Solidity: event PurchaseCompleted(uint256 indexed id, address indexed buyer, uint256 price)
func (_Marketplace *MarketplaceFilterer) ParsePurchaseCompleted(log types.Log) (*MarketplacePurchaseCompleted, error) {
	event := new(MarketplacePurchaseCompleted)
	if err := _Marketplace.contract.UnpackLog(event, "PurchaseCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
