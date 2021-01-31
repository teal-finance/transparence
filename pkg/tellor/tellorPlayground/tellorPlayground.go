// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tellorPlayground

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// TellorPlaygroundABI is the input ABI used to generate the binding from.
const TellorPlaygroundABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_time\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"NewValue\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tip\",\"type\":\"uint256\"}],\"name\":\"TipAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"addTip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"balances\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"disputeValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"faucet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"isDisputed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"submitValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"timestamps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TellorPlayground is an auto generated Go binding around an Ethereum contract.
type TellorPlayground struct {
	TellorPlaygroundCaller     // Read-only binding to the contract
	TellorPlaygroundTransactor // Write-only binding to the contract
	TellorPlaygroundFilterer   // Log filterer for contract events
}

// TellorPlaygroundCaller is an auto generated read-only Go binding around an Ethereum contract.
type TellorPlaygroundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorPlaygroundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TellorPlaygroundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorPlaygroundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TellorPlaygroundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TellorPlaygroundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TellorPlaygroundSession struct {
	Contract     *TellorPlayground // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TellorPlaygroundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TellorPlaygroundCallerSession struct {
	Contract *TellorPlaygroundCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// TellorPlaygroundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TellorPlaygroundTransactorSession struct {
	Contract     *TellorPlaygroundTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// TellorPlaygroundRaw is an auto generated low-level Go binding around an Ethereum contract.
type TellorPlaygroundRaw struct {
	Contract *TellorPlayground // Generic contract binding to access the raw methods on
}

// TellorPlaygroundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TellorPlaygroundCallerRaw struct {
	Contract *TellorPlaygroundCaller // Generic read-only contract binding to access the raw methods on
}

// TellorPlaygroundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TellorPlaygroundTransactorRaw struct {
	Contract *TellorPlaygroundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTellorPlayground creates a new instance of TellorPlayground, bound to a specific deployed contract.
func NewTellorPlayground(address common.Address, backend bind.ContractBackend) (*TellorPlayground, error) {
	contract, err := bindTellorPlayground(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TellorPlayground{TellorPlaygroundCaller: TellorPlaygroundCaller{contract: contract}, TellorPlaygroundTransactor: TellorPlaygroundTransactor{contract: contract}, TellorPlaygroundFilterer: TellorPlaygroundFilterer{contract: contract}}, nil
}

// NewTellorPlaygroundCaller creates a new read-only instance of TellorPlayground, bound to a specific deployed contract.
func NewTellorPlaygroundCaller(address common.Address, caller bind.ContractCaller) (*TellorPlaygroundCaller, error) {
	contract, err := bindTellorPlayground(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TellorPlaygroundCaller{contract: contract}, nil
}

// NewTellorPlaygroundTransactor creates a new write-only instance of TellorPlayground, bound to a specific deployed contract.
func NewTellorPlaygroundTransactor(address common.Address, transactor bind.ContractTransactor) (*TellorPlaygroundTransactor, error) {
	contract, err := bindTellorPlayground(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TellorPlaygroundTransactor{contract: contract}, nil
}

// NewTellorPlaygroundFilterer creates a new log filterer instance of TellorPlayground, bound to a specific deployed contract.
func NewTellorPlaygroundFilterer(address common.Address, filterer bind.ContractFilterer) (*TellorPlaygroundFilterer, error) {
	contract, err := bindTellorPlayground(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TellorPlaygroundFilterer{contract: contract}, nil
}

// bindTellorPlayground binds a generic wrapper to an already deployed contract.
func bindTellorPlayground(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TellorPlaygroundABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorPlayground *TellorPlaygroundRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorPlayground.Contract.TellorPlaygroundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorPlayground *TellorPlaygroundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorPlayground.Contract.TellorPlaygroundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorPlayground *TellorPlaygroundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorPlayground.Contract.TellorPlaygroundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TellorPlayground *TellorPlaygroundCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TellorPlayground.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TellorPlayground *TellorPlaygroundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TellorPlayground.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TellorPlayground *TellorPlaygroundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TellorPlayground.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TellorPlayground.Contract.Allowance(&_TellorPlayground.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _TellorPlayground.Contract.Allowance(&_TellorPlayground.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TellorPlayground.Contract.BalanceOf(&_TellorPlayground.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _TellorPlayground.Contract.BalanceOf(&_TellorPlayground.CallOpts, account)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCaller) Balances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "balances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _TellorPlayground.Contract.Balances(&_TellorPlayground.CallOpts, arg0)
}

// Balances is a free data retrieval call binding the contract method 0x27e235e3.
//
// Solidity: function balances(address ) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCallerSession) Balances(arg0 common.Address) (*big.Int, error) {
	return _TellorPlayground.Contract.Balances(&_TellorPlayground.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TellorPlayground *TellorPlaygroundCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TellorPlayground *TellorPlaygroundSession) Decimals() (uint8, error) {
	return _TellorPlayground.Contract.Decimals(&_TellorPlayground.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_TellorPlayground *TellorPlaygroundCallerSession) Decimals() (uint8, error) {
	return _TellorPlayground.Contract.Decimals(&_TellorPlayground.CallOpts)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorPlayground.Contract.GetNewValueCountbyRequestId(&_TellorPlayground.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _TellorPlayground.Contract.GetNewValueCountbyRequestId(&_TellorPlayground.CallOpts, _requestId)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 index) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestId *big.Int, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestId, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 index) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, index *big.Int) (*big.Int, error) {
	return _TellorPlayground.Contract.GetTimestampbyRequestIDandIndex(&_TellorPlayground.CallOpts, _requestId, index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 index) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCallerSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, index *big.Int) (*big.Int, error) {
	return _TellorPlayground.Contract.GetTimestampbyRequestIDandIndex(&_TellorPlayground.CallOpts, _requestId, index)
}

// IsDisputed is a free data retrieval call binding the contract method 0xb041d696.
//
// Solidity: function isDisputed(uint256 , uint256 ) view returns(bool)
func (_TellorPlayground *TellorPlaygroundCaller) IsDisputed(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "isDisputed", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDisputed is a free data retrieval call binding the contract method 0xb041d696.
//
// Solidity: function isDisputed(uint256 , uint256 ) view returns(bool)
func (_TellorPlayground *TellorPlaygroundSession) IsDisputed(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _TellorPlayground.Contract.IsDisputed(&_TellorPlayground.CallOpts, arg0, arg1)
}

// IsDisputed is a free data retrieval call binding the contract method 0xb041d696.
//
// Solidity: function isDisputed(uint256 , uint256 ) view returns(bool)
func (_TellorPlayground *TellorPlaygroundCallerSession) IsDisputed(arg0 *big.Int, arg1 *big.Int) (bool, error) {
	return _TellorPlayground.Contract.IsDisputed(&_TellorPlayground.CallOpts, arg0, arg1)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_TellorPlayground *TellorPlaygroundCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_TellorPlayground *TellorPlaygroundSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _TellorPlayground.Contract.IsInDispute(&_TellorPlayground.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_TellorPlayground *TellorPlaygroundCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _TellorPlayground.Contract.IsInDispute(&_TellorPlayground.CallOpts, _requestId, _timestamp)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TellorPlayground *TellorPlaygroundCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TellorPlayground *TellorPlaygroundSession) Name() (string, error) {
	return _TellorPlayground.Contract.Name(&_TellorPlayground.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_TellorPlayground *TellorPlaygroundCallerSession) Name() (string, error) {
	return _TellorPlayground.Contract.Name(&_TellorPlayground.CallOpts)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorPlayground.Contract.RetrieveData(&_TellorPlayground.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _TellorPlayground.Contract.RetrieveData(&_TellorPlayground.CallOpts, _requestId, _timestamp)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TellorPlayground *TellorPlaygroundCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TellorPlayground *TellorPlaygroundSession) Symbol() (string, error) {
	return _TellorPlayground.Contract.Symbol(&_TellorPlayground.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_TellorPlayground *TellorPlaygroundCallerSession) Symbol() (string, error) {
	return _TellorPlayground.Contract.Symbol(&_TellorPlayground.CallOpts)
}

// Timestamps is a free data retrieval call binding the contract method 0xfb0ceb04.
//
// Solidity: function timestamps(uint256 , uint256 ) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCaller) Timestamps(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "timestamps", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Timestamps is a free data retrieval call binding the contract method 0xfb0ceb04.
//
// Solidity: function timestamps(uint256 , uint256 ) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundSession) Timestamps(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorPlayground.Contract.Timestamps(&_TellorPlayground.CallOpts, arg0, arg1)
}

// Timestamps is a free data retrieval call binding the contract method 0xfb0ceb04.
//
// Solidity: function timestamps(uint256 , uint256 ) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCallerSession) Timestamps(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorPlayground.Contract.Timestamps(&_TellorPlayground.CallOpts, arg0, arg1)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TellorPlayground *TellorPlaygroundSession) TotalSupply() (*big.Int, error) {
	return _TellorPlayground.Contract.TotalSupply(&_TellorPlayground.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCallerSession) TotalSupply() (*big.Int, error) {
	return _TellorPlayground.Contract.TotalSupply(&_TellorPlayground.CallOpts)
}

// Values is a free data retrieval call binding the contract method 0xa3183701.
//
// Solidity: function values(uint256 , uint256 ) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCaller) Values(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TellorPlayground.contract.Call(opts, &out, "values", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Values is a free data retrieval call binding the contract method 0xa3183701.
//
// Solidity: function values(uint256 , uint256 ) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundSession) Values(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorPlayground.Contract.Values(&_TellorPlayground.CallOpts, arg0, arg1)
}

// Values is a free data retrieval call binding the contract method 0xa3183701.
//
// Solidity: function values(uint256 , uint256 ) view returns(uint256)
func (_TellorPlayground *TellorPlaygroundCallerSession) Values(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _TellorPlayground.Contract.Values(&_TellorPlayground.CallOpts, arg0, arg1)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _amount) returns()
func (_TellorPlayground *TellorPlaygroundTransactor) AddTip(opts *bind.TransactOpts, _requestId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.contract.Transact(opts, "addTip", _requestId, _amount)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _amount) returns()
func (_TellorPlayground *TellorPlaygroundSession) AddTip(_requestId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.AddTip(&_TellorPlayground.TransactOpts, _requestId, _amount)
}

// AddTip is a paid mutator transaction binding the contract method 0x752d49a1.
//
// Solidity: function addTip(uint256 _requestId, uint256 _amount) returns()
func (_TellorPlayground *TellorPlaygroundTransactorSession) AddTip(_requestId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.AddTip(&_TellorPlayground.TransactOpts, _requestId, _amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TellorPlayground *TellorPlaygroundTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TellorPlayground *TellorPlaygroundSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.Approve(&_TellorPlayground.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_TellorPlayground *TellorPlaygroundTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.Approve(&_TellorPlayground.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TellorPlayground *TellorPlaygroundTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TellorPlayground *TellorPlaygroundSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.DecreaseAllowance(&_TellorPlayground.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_TellorPlayground *TellorPlaygroundTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.DecreaseAllowance(&_TellorPlayground.TransactOpts, spender, subtractedValue)
}

// DisputeValue is a paid mutator transaction binding the contract method 0xacebfc54.
//
// Solidity: function disputeValue(uint256 _requestId, uint256 _timestamp) returns()
func (_TellorPlayground *TellorPlaygroundTransactor) DisputeValue(opts *bind.TransactOpts, _requestId *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.contract.Transact(opts, "disputeValue", _requestId, _timestamp)
}

// DisputeValue is a paid mutator transaction binding the contract method 0xacebfc54.
//
// Solidity: function disputeValue(uint256 _requestId, uint256 _timestamp) returns()
func (_TellorPlayground *TellorPlaygroundSession) DisputeValue(_requestId *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.DisputeValue(&_TellorPlayground.TransactOpts, _requestId, _timestamp)
}

// DisputeValue is a paid mutator transaction binding the contract method 0xacebfc54.
//
// Solidity: function disputeValue(uint256 _requestId, uint256 _timestamp) returns()
func (_TellorPlayground *TellorPlaygroundTransactorSession) DisputeValue(_requestId *big.Int, _timestamp *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.DisputeValue(&_TellorPlayground.TransactOpts, _requestId, _timestamp)
}

// Faucet is a paid mutator transaction binding the contract method 0xb86d1d63.
//
// Solidity: function faucet(address user) returns()
func (_TellorPlayground *TellorPlaygroundTransactor) Faucet(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _TellorPlayground.contract.Transact(opts, "faucet", user)
}

// Faucet is a paid mutator transaction binding the contract method 0xb86d1d63.
//
// Solidity: function faucet(address user) returns()
func (_TellorPlayground *TellorPlaygroundSession) Faucet(user common.Address) (*types.Transaction, error) {
	return _TellorPlayground.Contract.Faucet(&_TellorPlayground.TransactOpts, user)
}

// Faucet is a paid mutator transaction binding the contract method 0xb86d1d63.
//
// Solidity: function faucet(address user) returns()
func (_TellorPlayground *TellorPlaygroundTransactorSession) Faucet(user common.Address) (*types.Transaction, error) {
	return _TellorPlayground.Contract.Faucet(&_TellorPlayground.TransactOpts, user)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TellorPlayground *TellorPlaygroundTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TellorPlayground *TellorPlaygroundSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.IncreaseAllowance(&_TellorPlayground.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_TellorPlayground *TellorPlaygroundTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.IncreaseAllowance(&_TellorPlayground.TransactOpts, spender, addedValue)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x62f55112.
//
// Solidity: function submitValue(uint256 _requestId, uint256 _value) returns()
func (_TellorPlayground *TellorPlaygroundTransactor) SubmitValue(opts *bind.TransactOpts, _requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.contract.Transact(opts, "submitValue", _requestId, _value)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x62f55112.
//
// Solidity: function submitValue(uint256 _requestId, uint256 _value) returns()
func (_TellorPlayground *TellorPlaygroundSession) SubmitValue(_requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.SubmitValue(&_TellorPlayground.TransactOpts, _requestId, _value)
}

// SubmitValue is a paid mutator transaction binding the contract method 0x62f55112.
//
// Solidity: function submitValue(uint256 _requestId, uint256 _value) returns()
func (_TellorPlayground *TellorPlaygroundTransactorSession) SubmitValue(_requestId *big.Int, _value *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.SubmitValue(&_TellorPlayground.TransactOpts, _requestId, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TellorPlayground *TellorPlaygroundTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TellorPlayground *TellorPlaygroundSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.Transfer(&_TellorPlayground.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_TellorPlayground *TellorPlaygroundTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.Transfer(&_TellorPlayground.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TellorPlayground *TellorPlaygroundTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TellorPlayground *TellorPlaygroundSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.TransferFrom(&_TellorPlayground.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_TellorPlayground *TellorPlaygroundTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _TellorPlayground.Contract.TransferFrom(&_TellorPlayground.TransactOpts, sender, recipient, amount)
}

// TellorPlaygroundApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the TellorPlayground contract.
type TellorPlaygroundApprovalIterator struct {
	Event *TellorPlaygroundApproval // Event containing the contract specifics and raw log

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
func (it *TellorPlaygroundApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorPlaygroundApproval)
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
		it.Event = new(TellorPlaygroundApproval)
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
func (it *TellorPlaygroundApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorPlaygroundApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorPlaygroundApproval represents a Approval event raised by the TellorPlayground contract.
type TellorPlaygroundApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TellorPlayground *TellorPlaygroundFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*TellorPlaygroundApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TellorPlayground.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &TellorPlaygroundApprovalIterator{contract: _TellorPlayground.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TellorPlayground *TellorPlaygroundFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *TellorPlaygroundApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _TellorPlayground.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorPlaygroundApproval)
				if err := _TellorPlayground.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_TellorPlayground *TellorPlaygroundFilterer) ParseApproval(log types.Log) (*TellorPlaygroundApproval, error) {
	event := new(TellorPlaygroundApproval)
	if err := _TellorPlayground.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorPlaygroundNewValueIterator is returned from FilterNewValue and is used to iterate over the raw logs and unpacked data for NewValue events raised by the TellorPlayground contract.
type TellorPlaygroundNewValueIterator struct {
	Event *TellorPlaygroundNewValue // Event containing the contract specifics and raw log

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
func (it *TellorPlaygroundNewValueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorPlaygroundNewValue)
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
		it.Event = new(TellorPlaygroundNewValue)
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
func (it *TellorPlaygroundNewValueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorPlaygroundNewValueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorPlaygroundNewValue represents a NewValue event raised by the TellorPlayground contract.
type TellorPlaygroundNewValue struct {
	RequestId *big.Int
	Time      *big.Int
	Value     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterNewValue is a free log retrieval operation binding the contract event 0xba11e319aee26e7bbac889432515ba301ec8f6d27bf6b94829c21a65c5f6ff25.
//
// Solidity: event NewValue(uint256 _requestId, uint256 _time, uint256 _value)
func (_TellorPlayground *TellorPlaygroundFilterer) FilterNewValue(opts *bind.FilterOpts) (*TellorPlaygroundNewValueIterator, error) {

	logs, sub, err := _TellorPlayground.contract.FilterLogs(opts, "NewValue")
	if err != nil {
		return nil, err
	}
	return &TellorPlaygroundNewValueIterator{contract: _TellorPlayground.contract, event: "NewValue", logs: logs, sub: sub}, nil
}

// WatchNewValue is a free log subscription operation binding the contract event 0xba11e319aee26e7bbac889432515ba301ec8f6d27bf6b94829c21a65c5f6ff25.
//
// Solidity: event NewValue(uint256 _requestId, uint256 _time, uint256 _value)
func (_TellorPlayground *TellorPlaygroundFilterer) WatchNewValue(opts *bind.WatchOpts, sink chan<- *TellorPlaygroundNewValue) (event.Subscription, error) {

	logs, sub, err := _TellorPlayground.contract.WatchLogs(opts, "NewValue")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorPlaygroundNewValue)
				if err := _TellorPlayground.contract.UnpackLog(event, "NewValue", log); err != nil {
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

// ParseNewValue is a log parse operation binding the contract event 0xba11e319aee26e7bbac889432515ba301ec8f6d27bf6b94829c21a65c5f6ff25.
//
// Solidity: event NewValue(uint256 _requestId, uint256 _time, uint256 _value)
func (_TellorPlayground *TellorPlaygroundFilterer) ParseNewValue(log types.Log) (*TellorPlaygroundNewValue, error) {
	event := new(TellorPlaygroundNewValue)
	if err := _TellorPlayground.contract.UnpackLog(event, "NewValue", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorPlaygroundTipAddedIterator is returned from FilterTipAdded and is used to iterate over the raw logs and unpacked data for TipAdded events raised by the TellorPlayground contract.
type TellorPlaygroundTipAddedIterator struct {
	Event *TellorPlaygroundTipAdded // Event containing the contract specifics and raw log

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
func (it *TellorPlaygroundTipAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorPlaygroundTipAdded)
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
		it.Event = new(TellorPlaygroundTipAdded)
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
func (it *TellorPlaygroundTipAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorPlaygroundTipAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorPlaygroundTipAdded represents a TipAdded event raised by the TellorPlayground contract.
type TellorPlaygroundTipAdded struct {
	Sender    common.Address
	RequestId *big.Int
	Tip       *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTipAdded is a free log retrieval operation binding the contract event 0x9e771e1220a6c2e407f3601f70a769ca9fff75a110d1687e0b582824673a1f5c.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip)
func (_TellorPlayground *TellorPlaygroundFilterer) FilterTipAdded(opts *bind.FilterOpts, _sender []common.Address, _requestId []*big.Int) (*TellorPlaygroundTipAddedIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorPlayground.contract.FilterLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return &TellorPlaygroundTipAddedIterator{contract: _TellorPlayground.contract, event: "TipAdded", logs: logs, sub: sub}, nil
}

// WatchTipAdded is a free log subscription operation binding the contract event 0x9e771e1220a6c2e407f3601f70a769ca9fff75a110d1687e0b582824673a1f5c.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip)
func (_TellorPlayground *TellorPlaygroundFilterer) WatchTipAdded(opts *bind.WatchOpts, sink chan<- *TellorPlaygroundTipAdded, _sender []common.Address, _requestId []*big.Int) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _requestIdRule []interface{}
	for _, _requestIdItem := range _requestId {
		_requestIdRule = append(_requestIdRule, _requestIdItem)
	}

	logs, sub, err := _TellorPlayground.contract.WatchLogs(opts, "TipAdded", _senderRule, _requestIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorPlaygroundTipAdded)
				if err := _TellorPlayground.contract.UnpackLog(event, "TipAdded", log); err != nil {
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

// ParseTipAdded is a log parse operation binding the contract event 0x9e771e1220a6c2e407f3601f70a769ca9fff75a110d1687e0b582824673a1f5c.
//
// Solidity: event TipAdded(address indexed _sender, uint256 indexed _requestId, uint256 _tip)
func (_TellorPlayground *TellorPlaygroundFilterer) ParseTipAdded(log types.Log) (*TellorPlaygroundTipAdded, error) {
	event := new(TellorPlaygroundTipAdded)
	if err := _TellorPlayground.contract.UnpackLog(event, "TipAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TellorPlaygroundTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the TellorPlayground contract.
type TellorPlaygroundTransferIterator struct {
	Event *TellorPlaygroundTransfer // Event containing the contract specifics and raw log

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
func (it *TellorPlaygroundTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TellorPlaygroundTransfer)
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
		it.Event = new(TellorPlaygroundTransfer)
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
func (it *TellorPlaygroundTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TellorPlaygroundTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TellorPlaygroundTransfer represents a Transfer event raised by the TellorPlayground contract.
type TellorPlaygroundTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TellorPlayground *TellorPlaygroundFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*TellorPlaygroundTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TellorPlayground.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &TellorPlaygroundTransferIterator{contract: _TellorPlayground.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TellorPlayground *TellorPlaygroundFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *TellorPlaygroundTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _TellorPlayground.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TellorPlaygroundTransfer)
				if err := _TellorPlayground.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_TellorPlayground *TellorPlaygroundFilterer) ParseTransfer(log types.Log) (*TellorPlaygroundTransfer, error) {
	event := new(TellorPlaygroundTransfer)
	if err := _TellorPlayground.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
