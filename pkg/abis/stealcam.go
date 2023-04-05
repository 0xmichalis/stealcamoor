// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abis

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
)

// StealcamMetaData contains all meta data concerning the Stealcam contract.
var StealcamMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Stolen\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseUri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"previousStealPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseUri\",\"type\":\"string\"}],\"name\":\"setBaseUri\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"steal\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StealcamABI is the input ABI used to generate the binding from.
// Deprecated: Use StealcamMetaData.ABI instead.
var StealcamABI = StealcamMetaData.ABI

// Stealcam is an auto generated Go binding around an Ethereum contract.
type Stealcam struct {
	StealcamCaller     // Read-only binding to the contract
	StealcamTransactor // Write-only binding to the contract
	StealcamFilterer   // Log filterer for contract events
}

// StealcamCaller is an auto generated read-only Go binding around an Ethereum contract.
type StealcamCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StealcamTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StealcamTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StealcamFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StealcamFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StealcamSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StealcamSession struct {
	Contract     *Stealcam         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StealcamCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StealcamCallerSession struct {
	Contract *StealcamCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// StealcamTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StealcamTransactorSession struct {
	Contract     *StealcamTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// StealcamRaw is an auto generated low-level Go binding around an Ethereum contract.
type StealcamRaw struct {
	Contract *Stealcam // Generic contract binding to access the raw methods on
}

// StealcamCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StealcamCallerRaw struct {
	Contract *StealcamCaller // Generic read-only contract binding to access the raw methods on
}

// StealcamTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StealcamTransactorRaw struct {
	Contract *StealcamTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStealcam creates a new instance of Stealcam, bound to a specific deployed contract.
func NewStealcam(address common.Address, backend bind.ContractBackend) (*Stealcam, error) {
	contract, err := bindStealcam(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stealcam{StealcamCaller: StealcamCaller{contract: contract}, StealcamTransactor: StealcamTransactor{contract: contract}, StealcamFilterer: StealcamFilterer{contract: contract}}, nil
}

// NewStealcamCaller creates a new read-only instance of Stealcam, bound to a specific deployed contract.
func NewStealcamCaller(address common.Address, caller bind.ContractCaller) (*StealcamCaller, error) {
	contract, err := bindStealcam(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StealcamCaller{contract: contract}, nil
}

// NewStealcamTransactor creates a new write-only instance of Stealcam, bound to a specific deployed contract.
func NewStealcamTransactor(address common.Address, transactor bind.ContractTransactor) (*StealcamTransactor, error) {
	contract, err := bindStealcam(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StealcamTransactor{contract: contract}, nil
}

// NewStealcamFilterer creates a new log filterer instance of Stealcam, bound to a specific deployed contract.
func NewStealcamFilterer(address common.Address, filterer bind.ContractFilterer) (*StealcamFilterer, error) {
	contract, err := bindStealcam(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StealcamFilterer{contract: contract}, nil
}

// bindStealcam binds a generic wrapper to an already deployed contract.
func bindStealcam(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StealcamABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stealcam *StealcamRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stealcam.Contract.StealcamCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stealcam *StealcamRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stealcam.Contract.StealcamTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stealcam *StealcamRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stealcam.Contract.StealcamTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stealcam *StealcamCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stealcam.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stealcam *StealcamTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stealcam.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stealcam *StealcamTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stealcam.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Stealcam *StealcamCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Stealcam.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Stealcam *StealcamSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Stealcam.Contract.BalanceOf(&_Stealcam.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Stealcam *StealcamCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Stealcam.Contract.BalanceOf(&_Stealcam.CallOpts, owner)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_Stealcam *StealcamCaller) BaseUri(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stealcam.contract.Call(opts, &out, "baseUri")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_Stealcam *StealcamSession) BaseUri() (string, error) {
	return _Stealcam.Contract.BaseUri(&_Stealcam.CallOpts)
}

// BaseUri is a free data retrieval call binding the contract method 0x9abc8320.
//
// Solidity: function baseUri() view returns(string)
func (_Stealcam *StealcamCallerSession) BaseUri() (string, error) {
	return _Stealcam.Contract.BaseUri(&_Stealcam.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 ) view returns(address)
func (_Stealcam *StealcamCaller) Creator(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stealcam.contract.Call(opts, &out, "creator", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 ) view returns(address)
func (_Stealcam *StealcamSession) Creator(arg0 *big.Int) (common.Address, error) {
	return _Stealcam.Contract.Creator(&_Stealcam.CallOpts, arg0)
}

// Creator is a free data retrieval call binding the contract method 0x510b5158.
//
// Solidity: function creator(uint256 ) view returns(address)
func (_Stealcam *StealcamCallerSession) Creator(arg0 *big.Int) (common.Address, error) {
	return _Stealcam.Contract.Creator(&_Stealcam.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Stealcam *StealcamCaller) GetApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stealcam.contract.Call(opts, &out, "getApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Stealcam *StealcamSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Stealcam.Contract.GetApproved(&_Stealcam.CallOpts, arg0)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 ) view returns(address)
func (_Stealcam *StealcamCallerSession) GetApproved(arg0 *big.Int) (common.Address, error) {
	return _Stealcam.Contract.GetApproved(&_Stealcam.CallOpts, arg0)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Stealcam *StealcamCaller) IsApprovedForAll(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (bool, error) {
	var out []interface{}
	err := _Stealcam.contract.Call(opts, &out, "isApprovedForAll", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Stealcam *StealcamSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Stealcam.Contract.IsApprovedForAll(&_Stealcam.CallOpts, arg0, arg1)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address , address ) view returns(bool)
func (_Stealcam *StealcamCallerSession) IsApprovedForAll(arg0 common.Address, arg1 common.Address) (bool, error) {
	return _Stealcam.Contract.IsApprovedForAll(&_Stealcam.CallOpts, arg0, arg1)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Stealcam *StealcamCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stealcam.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Stealcam *StealcamSession) Name() (string, error) {
	return _Stealcam.Contract.Name(&_Stealcam.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Stealcam *StealcamCallerSession) Name() (string, error) {
	return _Stealcam.Contract.Name(&_Stealcam.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stealcam *StealcamCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stealcam.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stealcam *StealcamSession) Owner() (common.Address, error) {
	return _Stealcam.Contract.Owner(&_Stealcam.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Stealcam *StealcamCallerSession) Owner() (common.Address, error) {
	return _Stealcam.Contract.Owner(&_Stealcam.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Stealcam *StealcamCaller) OwnerOf(opts *bind.CallOpts, id *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stealcam.contract.Call(opts, &out, "ownerOf", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Stealcam *StealcamSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Stealcam.Contract.OwnerOf(&_Stealcam.CallOpts, id)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 id) view returns(address owner)
func (_Stealcam *StealcamCallerSession) OwnerOf(id *big.Int) (common.Address, error) {
	return _Stealcam.Contract.OwnerOf(&_Stealcam.CallOpts, id)
}

// PreviousStealPrice is a free data retrieval call binding the contract method 0xaac5e508.
//
// Solidity: function previousStealPrice(uint256 ) view returns(uint256)
func (_Stealcam *StealcamCaller) PreviousStealPrice(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stealcam.contract.Call(opts, &out, "previousStealPrice", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PreviousStealPrice is a free data retrieval call binding the contract method 0xaac5e508.
//
// Solidity: function previousStealPrice(uint256 ) view returns(uint256)
func (_Stealcam *StealcamSession) PreviousStealPrice(arg0 *big.Int) (*big.Int, error) {
	return _Stealcam.Contract.PreviousStealPrice(&_Stealcam.CallOpts, arg0)
}

// PreviousStealPrice is a free data retrieval call binding the contract method 0xaac5e508.
//
// Solidity: function previousStealPrice(uint256 ) view returns(uint256)
func (_Stealcam *StealcamCallerSession) PreviousStealPrice(arg0 *big.Int) (*big.Int, error) {
	return _Stealcam.Contract.PreviousStealPrice(&_Stealcam.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stealcam *StealcamCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Stealcam.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stealcam *StealcamSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stealcam.Contract.SupportsInterface(&_Stealcam.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stealcam *StealcamCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stealcam.Contract.SupportsInterface(&_Stealcam.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Stealcam *StealcamCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Stealcam.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Stealcam *StealcamSession) Symbol() (string, error) {
	return _Stealcam.Contract.Symbol(&_Stealcam.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Stealcam *StealcamCallerSession) Symbol() (string, error) {
	return _Stealcam.Contract.Symbol(&_Stealcam.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Stealcam *StealcamCaller) TokenURI(opts *bind.CallOpts, id *big.Int) (string, error) {
	var out []interface{}
	err := _Stealcam.contract.Call(opts, &out, "tokenURI", id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Stealcam *StealcamSession) TokenURI(id *big.Int) (string, error) {
	return _Stealcam.Contract.TokenURI(&_Stealcam.CallOpts, id)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 id) view returns(string)
func (_Stealcam *StealcamCallerSession) TokenURI(id *big.Int) (string, error) {
	return _Stealcam.Contract.TokenURI(&_Stealcam.CallOpts, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Stealcam *StealcamTransactor) Approve(opts *bind.TransactOpts, spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Stealcam.contract.Transact(opts, "approve", spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Stealcam *StealcamSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Stealcam.Contract.Approve(&_Stealcam.TransactOpts, spender, id)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 id) returns()
func (_Stealcam *StealcamTransactorSession) Approve(spender common.Address, id *big.Int) (*types.Transaction, error) {
	return _Stealcam.Contract.Approve(&_Stealcam.TransactOpts, spender, id)
}

// Mint is a paid mutator transaction binding the contract method 0x73c02519.
//
// Solidity: function mint(uint256 id, address _creator, bytes signature) returns()
func (_Stealcam *StealcamTransactor) Mint(opts *bind.TransactOpts, id *big.Int, _creator common.Address, signature []byte) (*types.Transaction, error) {
	return _Stealcam.contract.Transact(opts, "mint", id, _creator, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x73c02519.
//
// Solidity: function mint(uint256 id, address _creator, bytes signature) returns()
func (_Stealcam *StealcamSession) Mint(id *big.Int, _creator common.Address, signature []byte) (*types.Transaction, error) {
	return _Stealcam.Contract.Mint(&_Stealcam.TransactOpts, id, _creator, signature)
}

// Mint is a paid mutator transaction binding the contract method 0x73c02519.
//
// Solidity: function mint(uint256 id, address _creator, bytes signature) returns()
func (_Stealcam *StealcamTransactorSession) Mint(id *big.Int, _creator common.Address, signature []byte) (*types.Transaction, error) {
	return _Stealcam.Contract.Mint(&_Stealcam.TransactOpts, id, _creator, signature)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 id, address _creator) returns()
func (_Stealcam *StealcamTransactor) Mint0(opts *bind.TransactOpts, id *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _Stealcam.contract.Transact(opts, "mint0", id, _creator)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 id, address _creator) returns()
func (_Stealcam *StealcamSession) Mint0(id *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _Stealcam.Contract.Mint0(&_Stealcam.TransactOpts, id, _creator)
}

// Mint0 is a paid mutator transaction binding the contract method 0x94bf804d.
//
// Solidity: function mint(uint256 id, address _creator) returns()
func (_Stealcam *StealcamTransactorSession) Mint0(id *big.Int, _creator common.Address) (*types.Transaction, error) {
	return _Stealcam.Contract.Mint0(&_Stealcam.TransactOpts, id, _creator)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stealcam *StealcamTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stealcam.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stealcam *StealcamSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stealcam.Contract.RenounceOwnership(&_Stealcam.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Stealcam *StealcamTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Stealcam.Contract.RenounceOwnership(&_Stealcam.TransactOpts)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Stealcam *StealcamTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Stealcam.contract.Transact(opts, "safeTransferFrom", from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Stealcam *StealcamSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Stealcam.Contract.SafeTransferFrom(&_Stealcam.TransactOpts, from, to, id)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id) returns()
func (_Stealcam *StealcamTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Stealcam.Contract.SafeTransferFrom(&_Stealcam.TransactOpts, from, to, id)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Stealcam *StealcamTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Stealcam.contract.Transact(opts, "safeTransferFrom0", from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Stealcam *StealcamSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Stealcam.Contract.SafeTransferFrom0(&_Stealcam.TransactOpts, from, to, id, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, bytes data) returns()
func (_Stealcam *StealcamTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, id *big.Int, data []byte) (*types.Transaction, error) {
	return _Stealcam.Contract.SafeTransferFrom0(&_Stealcam.TransactOpts, from, to, id, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Stealcam *StealcamTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Stealcam.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Stealcam *StealcamSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Stealcam.Contract.SetApprovalForAll(&_Stealcam.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Stealcam *StealcamTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Stealcam.Contract.SetApprovalForAll(&_Stealcam.TransactOpts, operator, approved)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string _baseUri) returns()
func (_Stealcam *StealcamTransactor) SetBaseUri(opts *bind.TransactOpts, _baseUri string) (*types.Transaction, error) {
	return _Stealcam.contract.Transact(opts, "setBaseUri", _baseUri)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string _baseUri) returns()
func (_Stealcam *StealcamSession) SetBaseUri(_baseUri string) (*types.Transaction, error) {
	return _Stealcam.Contract.SetBaseUri(&_Stealcam.TransactOpts, _baseUri)
}

// SetBaseUri is a paid mutator transaction binding the contract method 0xa0bcfc7f.
//
// Solidity: function setBaseUri(string _baseUri) returns()
func (_Stealcam *StealcamTransactorSession) SetBaseUri(_baseUri string) (*types.Transaction, error) {
	return _Stealcam.Contract.SetBaseUri(&_Stealcam.TransactOpts, _baseUri)
}

// Steal is a paid mutator transaction binding the contract method 0x9bcf5f0c.
//
// Solidity: function steal(uint256 id) payable returns()
func (_Stealcam *StealcamTransactor) Steal(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Stealcam.contract.Transact(opts, "steal", id)
}

// Steal is a paid mutator transaction binding the contract method 0x9bcf5f0c.
//
// Solidity: function steal(uint256 id) payable returns()
func (_Stealcam *StealcamSession) Steal(id *big.Int) (*types.Transaction, error) {
	return _Stealcam.Contract.Steal(&_Stealcam.TransactOpts, id)
}

// Steal is a paid mutator transaction binding the contract method 0x9bcf5f0c.
//
// Solidity: function steal(uint256 id) payable returns()
func (_Stealcam *StealcamTransactorSession) Steal(id *big.Int) (*types.Transaction, error) {
	return _Stealcam.Contract.Steal(&_Stealcam.TransactOpts, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Stealcam *StealcamTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Stealcam.contract.Transact(opts, "transferFrom", from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Stealcam *StealcamSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Stealcam.Contract.TransferFrom(&_Stealcam.TransactOpts, from, to, id)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 id) returns()
func (_Stealcam *StealcamTransactorSession) TransferFrom(from common.Address, to common.Address, id *big.Int) (*types.Transaction, error) {
	return _Stealcam.Contract.TransferFrom(&_Stealcam.TransactOpts, from, to, id)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stealcam *StealcamTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Stealcam.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stealcam *StealcamSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stealcam.Contract.TransferOwnership(&_Stealcam.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Stealcam *StealcamTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Stealcam.Contract.TransferOwnership(&_Stealcam.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Stealcam *StealcamTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stealcam.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Stealcam *StealcamSession) Withdraw() (*types.Transaction, error) {
	return _Stealcam.Contract.Withdraw(&_Stealcam.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Stealcam *StealcamTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Stealcam.Contract.Withdraw(&_Stealcam.TransactOpts)
}

// StealcamApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Stealcam contract.
type StealcamApprovalIterator struct {
	Event *StealcamApproval // Event containing the contract specifics and raw log

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
func (it *StealcamApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StealcamApproval)
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
		it.Event = new(StealcamApproval)
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
func (it *StealcamApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StealcamApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StealcamApproval represents a Approval event raised by the Stealcam contract.
type StealcamApproval struct {
	Owner   common.Address
	Spender common.Address
	Id      *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Stealcam *StealcamFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address, id []*big.Int) (*StealcamApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Stealcam.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &StealcamApprovalIterator{contract: _Stealcam.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Stealcam *StealcamFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StealcamApproval, owner []common.Address, spender []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Stealcam.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StealcamApproval)
				if err := _Stealcam.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id)
func (_Stealcam *StealcamFilterer) ParseApproval(log types.Log) (*StealcamApproval, error) {
	event := new(StealcamApproval)
	if err := _Stealcam.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StealcamApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Stealcam contract.
type StealcamApprovalForAllIterator struct {
	Event *StealcamApprovalForAll // Event containing the contract specifics and raw log

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
func (it *StealcamApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StealcamApprovalForAll)
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
		it.Event = new(StealcamApprovalForAll)
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
func (it *StealcamApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StealcamApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StealcamApprovalForAll represents a ApprovalForAll event raised by the Stealcam contract.
type StealcamApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Stealcam *StealcamFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*StealcamApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Stealcam.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StealcamApprovalForAllIterator{contract: _Stealcam.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Stealcam *StealcamFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StealcamApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Stealcam.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StealcamApprovalForAll)
				if err := _Stealcam.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Stealcam *StealcamFilterer) ParseApprovalForAll(log types.Log) (*StealcamApprovalForAll, error) {
	event := new(StealcamApprovalForAll)
	if err := _Stealcam.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StealcamOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Stealcam contract.
type StealcamOwnershipTransferredIterator struct {
	Event *StealcamOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StealcamOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StealcamOwnershipTransferred)
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
		it.Event = new(StealcamOwnershipTransferred)
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
func (it *StealcamOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StealcamOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StealcamOwnershipTransferred represents a OwnershipTransferred event raised by the Stealcam contract.
type StealcamOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stealcam *StealcamFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StealcamOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stealcam.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StealcamOwnershipTransferredIterator{contract: _Stealcam.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stealcam *StealcamFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StealcamOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Stealcam.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StealcamOwnershipTransferred)
				if err := _Stealcam.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Stealcam *StealcamFilterer) ParseOwnershipTransferred(log types.Log) (*StealcamOwnershipTransferred, error) {
	event := new(StealcamOwnershipTransferred)
	if err := _Stealcam.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StealcamStolenIterator is returned from FilterStolen and is used to iterate over the raw logs and unpacked data for Stolen events raised by the Stealcam contract.
type StealcamStolenIterator struct {
	Event *StealcamStolen // Event containing the contract specifics and raw log

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
func (it *StealcamStolenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StealcamStolen)
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
		it.Event = new(StealcamStolen)
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
func (it *StealcamStolenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StealcamStolenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StealcamStolen represents a Stolen event raised by the Stealcam contract.
type StealcamStolen struct {
	From  common.Address
	To    common.Address
	Id    *big.Int
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterStolen is a free log retrieval operation binding the contract event 0xbd1d1f579700e7d5d89a06ef937990d5f920f734ad1b9b945b354c9643dfd322.
//
// Solidity: event Stolen(address from, address to, uint256 id, uint256 value)
func (_Stealcam *StealcamFilterer) FilterStolen(opts *bind.FilterOpts) (*StealcamStolenIterator, error) {

	logs, sub, err := _Stealcam.contract.FilterLogs(opts, "Stolen")
	if err != nil {
		return nil, err
	}
	return &StealcamStolenIterator{contract: _Stealcam.contract, event: "Stolen", logs: logs, sub: sub}, nil
}

// WatchStolen is a free log subscription operation binding the contract event 0xbd1d1f579700e7d5d89a06ef937990d5f920f734ad1b9b945b354c9643dfd322.
//
// Solidity: event Stolen(address from, address to, uint256 id, uint256 value)
func (_Stealcam *StealcamFilterer) WatchStolen(opts *bind.WatchOpts, sink chan<- *StealcamStolen) (event.Subscription, error) {

	logs, sub, err := _Stealcam.contract.WatchLogs(opts, "Stolen")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StealcamStolen)
				if err := _Stealcam.contract.UnpackLog(event, "Stolen", log); err != nil {
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

// ParseStolen is a log parse operation binding the contract event 0xbd1d1f579700e7d5d89a06ef937990d5f920f734ad1b9b945b354c9643dfd322.
//
// Solidity: event Stolen(address from, address to, uint256 id, uint256 value)
func (_Stealcam *StealcamFilterer) ParseStolen(log types.Log) (*StealcamStolen, error) {
	event := new(StealcamStolen)
	if err := _Stealcam.contract.UnpackLog(event, "Stolen", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StealcamTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Stealcam contract.
type StealcamTransferIterator struct {
	Event *StealcamTransfer // Event containing the contract specifics and raw log

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
func (it *StealcamTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StealcamTransfer)
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
		it.Event = new(StealcamTransfer)
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
func (it *StealcamTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StealcamTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StealcamTransfer represents a Transfer event raised by the Stealcam contract.
type StealcamTransfer struct {
	From common.Address
	To   common.Address
	Id   *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Stealcam *StealcamFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*StealcamTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Stealcam.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &StealcamTransferIterator{contract: _Stealcam.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Stealcam *StealcamFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StealcamTransfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Stealcam.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StealcamTransfer)
				if err := _Stealcam.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed id)
func (_Stealcam *StealcamFilterer) ParseTransfer(log types.Log) (*StealcamTransfer, error) {
	event := new(StealcamTransfer)
	if err := _Stealcam.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
