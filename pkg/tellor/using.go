// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tellor

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

// Lensvalue is an auto generated low-level Go binding around an user-defined struct.
type Lensvalue struct {
	Timestamp *big.Int
	Value     *big.Int
}

// UsingTellorABI is the input ABI used to generate the binding from.
const UsingTellorABI = "[{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_master\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"_deity\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_tBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentReward\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"currentTotalTips\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"difficulty\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"disputeCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"disputeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getCurrentValue\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"ifRetrieve\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestampRetrieved\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getDataBefore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"_ifRetrieve\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestampRetrieved\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"getIndexForDataBefore\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"found\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"getLastNewValues\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"internalType\":\"structLens.value[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"}],\"name\":\"getNewValueCountbyRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getTimestampbyRequestIDandIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"isInDispute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pending_owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"proxy\",\"outputs\":[{\"internalType\":\"contractTellorMaster\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_requestId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_timestamp\",\"type\":\"uint256\"}],\"name\":\"retrieveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"slotProgress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"stakerCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tellorContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeOfLastNewValue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"timeTarget\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestID\",\"type\":\"uint256\"}],\"name\":\"totalTip\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UsingTellor is an auto generated Go binding around an Ethereum contract.
type UsingTellor struct {
	UsingTellorCaller     // Read-only binding to the contract
	UsingTellorTransactor // Write-only binding to the contract
	UsingTellorFilterer   // Log filterer for contract events
}

// UsingTellorCaller is an auto generated read-only Go binding around an Ethereum contract.
type UsingTellorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingTellorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UsingTellorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingTellorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UsingTellorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UsingTellorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UsingTellorSession struct {
	Contract     *UsingTellor      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UsingTellorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UsingTellorCallerSession struct {
	Contract *UsingTellorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// UsingTellorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UsingTellorTransactorSession struct {
	Contract     *UsingTellorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// UsingTellorRaw is an auto generated low-level Go binding around an Ethereum contract.
type UsingTellorRaw struct {
	Contract *UsingTellor // Generic contract binding to access the raw methods on
}

// UsingTellorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UsingTellorCallerRaw struct {
	Contract *UsingTellorCaller // Generic read-only contract binding to access the raw methods on
}

// UsingTellorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UsingTellorTransactorRaw struct {
	Contract *UsingTellorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUsingTellor creates a new instance of UsingTellor, bound to a specific deployed contract.
func NewUsingTellor(address common.Address, backend bind.ContractBackend) (*UsingTellor, error) {
	contract, err := bindUsingTellor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UsingTellor{UsingTellorCaller: UsingTellorCaller{contract: contract}, UsingTellorTransactor: UsingTellorTransactor{contract: contract}, UsingTellorFilterer: UsingTellorFilterer{contract: contract}}, nil
}

// NewUsingTellorCaller creates a new read-only instance of UsingTellor, bound to a specific deployed contract.
func NewUsingTellorCaller(address common.Address, caller bind.ContractCaller) (*UsingTellorCaller, error) {
	contract, err := bindUsingTellor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UsingTellorCaller{contract: contract}, nil
}

// NewUsingTellorTransactor creates a new write-only instance of UsingTellor, bound to a specific deployed contract.
func NewUsingTellorTransactor(address common.Address, transactor bind.ContractTransactor) (*UsingTellorTransactor, error) {
	contract, err := bindUsingTellor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UsingTellorTransactor{contract: contract}, nil
}

// NewUsingTellorFilterer creates a new log filterer instance of UsingTellor, bound to a specific deployed contract.
func NewUsingTellorFilterer(address common.Address, filterer bind.ContractFilterer) (*UsingTellorFilterer, error) {
	contract, err := bindUsingTellor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UsingTellorFilterer{contract: contract}, nil
}

// bindUsingTellor binds a generic wrapper to an already deployed contract.
func bindUsingTellor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UsingTellorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingTellor *UsingTellorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UsingTellor.Contract.UsingTellorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingTellor *UsingTellorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingTellor.Contract.UsingTellorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingTellor *UsingTellorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingTellor.Contract.UsingTellorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UsingTellor *UsingTellorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UsingTellor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UsingTellor *UsingTellorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UsingTellor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UsingTellor *UsingTellorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UsingTellor.Contract.contract.Transact(opts, method, params...)
}

// Deity is a free data retrieval call binding the contract method 0xffbe5e9e.
//
// Solidity: function _deity() view returns(address)
func (_UsingTellor *UsingTellorCaller) Deity(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "_deity")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Deity is a free data retrieval call binding the contract method 0xffbe5e9e.
//
// Solidity: function _deity() view returns(address)
func (_UsingTellor *UsingTellorSession) Deity() (common.Address, error) {
	return _UsingTellor.Contract.Deity(&_UsingTellor.CallOpts)
}

// Deity is a free data retrieval call binding the contract method 0xffbe5e9e.
//
// Solidity: function _deity() view returns(address)
func (_UsingTellor *UsingTellorCallerSession) Deity() (common.Address, error) {
	return _UsingTellor.Contract.Deity(&_UsingTellor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_UsingTellor *UsingTellorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_UsingTellor *UsingTellorSession) Owner() (common.Address, error) {
	return _UsingTellor.Contract.Owner(&_UsingTellor.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0xb2bdfa7b.
//
// Solidity: function _owner() view returns(address)
func (_UsingTellor *UsingTellorCallerSession) Owner() (common.Address, error) {
	return _UsingTellor.Contract.Owner(&_UsingTellor.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(uint256)
func (_UsingTellor *UsingTellorCaller) TBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "_tBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(uint256)
func (_UsingTellor *UsingTellorSession) TBlock() (*big.Int, error) {
	return _UsingTellor.Contract.TBlock(&_UsingTellor.CallOpts)
}

// TBlock is a free data retrieval call binding the contract method 0x6e3cf885.
//
// Solidity: function _tBlock() view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) TBlock() (*big.Int, error) {
	return _UsingTellor.Contract.TBlock(&_UsingTellor.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(uint256)
func (_UsingTellor *UsingTellorCaller) CurrentReward(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "currentReward")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(uint256)
func (_UsingTellor *UsingTellorSession) CurrentReward() (*big.Int, error) {
	return _UsingTellor.Contract.CurrentReward(&_UsingTellor.CallOpts)
}

// CurrentReward is a free data retrieval call binding the contract method 0x07621eca.
//
// Solidity: function currentReward() view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) CurrentReward() (*big.Int, error) {
	return _UsingTellor.Contract.CurrentReward(&_UsingTellor.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(uint256)
func (_UsingTellor *UsingTellorCaller) CurrentTotalTips(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "currentTotalTips")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(uint256)
func (_UsingTellor *UsingTellorSession) CurrentTotalTips() (*big.Int, error) {
	return _UsingTellor.Contract.CurrentTotalTips(&_UsingTellor.CallOpts)
}

// CurrentTotalTips is a free data retrieval call binding the contract method 0x75ad1a2a.
//
// Solidity: function currentTotalTips() view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) CurrentTotalTips() (*big.Int, error) {
	return _UsingTellor.Contract.CurrentTotalTips(&_UsingTellor.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_UsingTellor *UsingTellorCaller) Difficulty(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "difficulty")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_UsingTellor *UsingTellorSession) Difficulty() (*big.Int, error) {
	return _UsingTellor.Contract.Difficulty(&_UsingTellor.CallOpts)
}

// Difficulty is a free data retrieval call binding the contract method 0x19cae462.
//
// Solidity: function difficulty() view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) Difficulty() (*big.Int, error) {
	return _UsingTellor.Contract.Difficulty(&_UsingTellor.CallOpts)
}

// DisputeCount is a free data retrieval call binding the contract method 0xa28889e1.
//
// Solidity: function disputeCount() view returns(uint256)
func (_UsingTellor *UsingTellorCaller) DisputeCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "disputeCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeCount is a free data retrieval call binding the contract method 0xa28889e1.
//
// Solidity: function disputeCount() view returns(uint256)
func (_UsingTellor *UsingTellorSession) DisputeCount() (*big.Int, error) {
	return _UsingTellor.Contract.DisputeCount(&_UsingTellor.CallOpts)
}

// DisputeCount is a free data retrieval call binding the contract method 0xa28889e1.
//
// Solidity: function disputeCount() view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) DisputeCount() (*big.Int, error) {
	return _UsingTellor.Contract.DisputeCount(&_UsingTellor.CallOpts)
}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(uint256)
func (_UsingTellor *UsingTellorCaller) DisputeFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "disputeFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(uint256)
func (_UsingTellor *UsingTellorSession) DisputeFee() (*big.Int, error) {
	return _UsingTellor.Contract.DisputeFee(&_UsingTellor.CallOpts)
}

// DisputeFee is a free data retrieval call binding the contract method 0xb9ce896b.
//
// Solidity: function disputeFee() view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) DisputeFee() (*big.Int, error) {
	return _UsingTellor.Contract.DisputeFee(&_UsingTellor.CallOpts)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool ifRetrieve, uint256 value, uint256 _timestampRetrieved)
func (_UsingTellor *UsingTellorCaller) GetCurrentValue(opts *bind.CallOpts, _requestId *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "getCurrentValue", _requestId)

	outstruct := new(struct {
		IfRetrieve         bool
		Value              *big.Int
		TimestampRetrieved *big.Int
	})

	outstruct.IfRetrieve = out[0].(bool)
	outstruct.Value = out[1].(*big.Int)
	outstruct.TimestampRetrieved = out[2].(*big.Int)

	return *outstruct, err

}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool ifRetrieve, uint256 value, uint256 _timestampRetrieved)
func (_UsingTellor *UsingTellorSession) GetCurrentValue(_requestId *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	return _UsingTellor.Contract.GetCurrentValue(&_UsingTellor.CallOpts, _requestId)
}

// GetCurrentValue is a free data retrieval call binding the contract method 0x3fcad964.
//
// Solidity: function getCurrentValue(uint256 _requestId) view returns(bool ifRetrieve, uint256 value, uint256 _timestampRetrieved)
func (_UsingTellor *UsingTellorCallerSession) GetCurrentValue(_requestId *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	return _UsingTellor.Contract.GetCurrentValue(&_UsingTellor.CallOpts, _requestId)
}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool _ifRetrieve, uint256 _value, uint256 _timestampRetrieved)
func (_UsingTellor *UsingTellorCaller) GetDataBefore(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "getDataBefore", _requestId, _timestamp)

	outstruct := new(struct {
		IfRetrieve         bool
		Value              *big.Int
		TimestampRetrieved *big.Int
	})

	outstruct.IfRetrieve = out[0].(bool)
	outstruct.Value = out[1].(*big.Int)
	outstruct.TimestampRetrieved = out[2].(*big.Int)

	return *outstruct, err

}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool _ifRetrieve, uint256 _value, uint256 _timestampRetrieved)
func (_UsingTellor *UsingTellorSession) GetDataBefore(_requestId *big.Int, _timestamp *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	return _UsingTellor.Contract.GetDataBefore(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// GetDataBefore is a free data retrieval call binding the contract method 0x66b44611.
//
// Solidity: function getDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool _ifRetrieve, uint256 _value, uint256 _timestampRetrieved)
func (_UsingTellor *UsingTellorCallerSession) GetDataBefore(_requestId *big.Int, _timestamp *big.Int) (struct {
	IfRetrieve         bool
	Value              *big.Int
	TimestampRetrieved *big.Int
}, error) {
	return _UsingTellor.Contract.GetDataBefore(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// GetIndexForDataBefore is a free data retrieval call binding the contract method 0xb73e4979.
//
// Solidity: function getIndexForDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool found, uint256 index)
func (_UsingTellor *UsingTellorCaller) GetIndexForDataBefore(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (struct {
	Found bool
	Index *big.Int
}, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "getIndexForDataBefore", _requestId, _timestamp)

	outstruct := new(struct {
		Found bool
		Index *big.Int
	})

	outstruct.Found = out[0].(bool)
	outstruct.Index = out[1].(*big.Int)

	return *outstruct, err

}

// GetIndexForDataBefore is a free data retrieval call binding the contract method 0xb73e4979.
//
// Solidity: function getIndexForDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool found, uint256 index)
func (_UsingTellor *UsingTellorSession) GetIndexForDataBefore(_requestId *big.Int, _timestamp *big.Int) (struct {
	Found bool
	Index *big.Int
}, error) {
	return _UsingTellor.Contract.GetIndexForDataBefore(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// GetIndexForDataBefore is a free data retrieval call binding the contract method 0xb73e4979.
//
// Solidity: function getIndexForDataBefore(uint256 _requestId, uint256 _timestamp) view returns(bool found, uint256 index)
func (_UsingTellor *UsingTellorCallerSession) GetIndexForDataBefore(_requestId *big.Int, _timestamp *big.Int) (struct {
	Found bool
	Index *big.Int
}, error) {
	return _UsingTellor.Contract.GetIndexForDataBefore(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// GetLastNewValues is a free data retrieval call binding the contract method 0x5b026f61.
//
// Solidity: function getLastNewValues(uint256 requestID, uint256 count) view returns((uint256,uint256)[])
func (_UsingTellor *UsingTellorCaller) GetLastNewValues(opts *bind.CallOpts, requestID *big.Int, count *big.Int) ([]Lensvalue, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "getLastNewValues", requestID, count)

	if err != nil {
		return *new([]Lensvalue), err
	}

	out0 := *abi.ConvertType(out[0], new([]Lensvalue)).(*[]Lensvalue)

	return out0, err

}

// GetLastNewValues is a free data retrieval call binding the contract method 0x5b026f61.
//
// Solidity: function getLastNewValues(uint256 requestID, uint256 count) view returns((uint256,uint256)[])
func (_UsingTellor *UsingTellorSession) GetLastNewValues(requestID *big.Int, count *big.Int) ([]Lensvalue, error) {
	return _UsingTellor.Contract.GetLastNewValues(&_UsingTellor.CallOpts, requestID, count)
}

// GetLastNewValues is a free data retrieval call binding the contract method 0x5b026f61.
//
// Solidity: function getLastNewValues(uint256 requestID, uint256 count) view returns((uint256,uint256)[])
func (_UsingTellor *UsingTellorCallerSession) GetLastNewValues(requestID *big.Int, count *big.Int) ([]Lensvalue, error) {
	return _UsingTellor.Contract.GetLastNewValues(&_UsingTellor.CallOpts, requestID, count)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_UsingTellor *UsingTellorCaller) GetNewValueCountbyRequestId(opts *bind.CallOpts, _requestId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "getNewValueCountbyRequestId", _requestId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_UsingTellor *UsingTellorSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.GetNewValueCountbyRequestId(&_UsingTellor.CallOpts, _requestId)
}

// GetNewValueCountbyRequestId is a free data retrieval call binding the contract method 0x46eee1c4.
//
// Solidity: function getNewValueCountbyRequestId(uint256 _requestId) view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) GetNewValueCountbyRequestId(_requestId *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.GetNewValueCountbyRequestId(&_UsingTellor.CallOpts, _requestId)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_UsingTellor *UsingTellorCaller) GetTimestampbyRequestIDandIndex(opts *bind.CallOpts, _requestId *big.Int, _index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "getTimestampbyRequestIDandIndex", _requestId, _index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_UsingTellor *UsingTellorSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, _index *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.GetTimestampbyRequestIDandIndex(&_UsingTellor.CallOpts, _requestId, _index)
}

// GetTimestampbyRequestIDandIndex is a free data retrieval call binding the contract method 0x77fbb663.
//
// Solidity: function getTimestampbyRequestIDandIndex(uint256 _requestId, uint256 _index) view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) GetTimestampbyRequestIDandIndex(_requestId *big.Int, _index *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.GetTimestampbyRequestIDandIndex(&_UsingTellor.CallOpts, _requestId, _index)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_UsingTellor *UsingTellorCaller) IsInDispute(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (bool, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "isInDispute", _requestId, _timestamp)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_UsingTellor *UsingTellorSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _UsingTellor.Contract.IsInDispute(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// IsInDispute is a free data retrieval call binding the contract method 0x3df0777b.
//
// Solidity: function isInDispute(uint256 _requestId, uint256 _timestamp) view returns(bool)
func (_UsingTellor *UsingTellorCallerSession) IsInDispute(_requestId *big.Int, _timestamp *big.Int) (bool, error) {
	return _UsingTellor.Contract.IsInDispute(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(address)
func (_UsingTellor *UsingTellorCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "pending_owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(address)
func (_UsingTellor *UsingTellorSession) PendingOwner() (common.Address, error) {
	return _UsingTellor.Contract.PendingOwner(&_UsingTellor.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0x7f4ec4c3.
//
// Solidity: function pending_owner() view returns(address)
func (_UsingTellor *UsingTellorCallerSession) PendingOwner() (common.Address, error) {
	return _UsingTellor.Contract.PendingOwner(&_UsingTellor.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_UsingTellor *UsingTellorCaller) Proxy(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "proxy")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_UsingTellor *UsingTellorSession) Proxy() (common.Address, error) {
	return _UsingTellor.Contract.Proxy(&_UsingTellor.CallOpts)
}

// Proxy is a free data retrieval call binding the contract method 0xec556889.
//
// Solidity: function proxy() view returns(address)
func (_UsingTellor *UsingTellorCallerSession) Proxy() (common.Address, error) {
	return _UsingTellor.Contract.Proxy(&_UsingTellor.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_UsingTellor *UsingTellorCaller) RequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "requestCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_UsingTellor *UsingTellorSession) RequestCount() (*big.Int, error) {
	return _UsingTellor.Contract.RequestCount(&_UsingTellor.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) RequestCount() (*big.Int, error) {
	return _UsingTellor.Contract.RequestCount(&_UsingTellor.CallOpts)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_UsingTellor *UsingTellorCaller) RetrieveData(opts *bind.CallOpts, _requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "retrieveData", _requestId, _timestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_UsingTellor *UsingTellorSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.RetrieveData(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// RetrieveData is a free data retrieval call binding the contract method 0x93fa4915.
//
// Solidity: function retrieveData(uint256 _requestId, uint256 _timestamp) view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) RetrieveData(_requestId *big.Int, _timestamp *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.RetrieveData(&_UsingTellor.CallOpts, _requestId, _timestamp)
}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(uint256)
func (_UsingTellor *UsingTellorCaller) SlotProgress(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "slotProgress")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(uint256)
func (_UsingTellor *UsingTellorSession) SlotProgress() (*big.Int, error) {
	return _UsingTellor.Contract.SlotProgress(&_UsingTellor.CallOpts)
}

// SlotProgress is a free data retrieval call binding the contract method 0x03b3160f.
//
// Solidity: function slotProgress() view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) SlotProgress() (*big.Int, error) {
	return _UsingTellor.Contract.SlotProgress(&_UsingTellor.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_UsingTellor *UsingTellorCaller) StakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "stakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_UsingTellor *UsingTellorSession) StakeAmount() (*big.Int, error) {
	return _UsingTellor.Contract.StakeAmount(&_UsingTellor.CallOpts)
}

// StakeAmount is a free data retrieval call binding the contract method 0x60c7dc47.
//
// Solidity: function stakeAmount() view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) StakeAmount() (*big.Int, error) {
	return _UsingTellor.Contract.StakeAmount(&_UsingTellor.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_UsingTellor *UsingTellorCaller) StakerCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "stakerCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_UsingTellor *UsingTellorSession) StakerCount() (*big.Int, error) {
	return _UsingTellor.Contract.StakerCount(&_UsingTellor.CallOpts)
}

// StakerCount is a free data retrieval call binding the contract method 0xdff69787.
//
// Solidity: function stakerCount() view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) StakerCount() (*big.Int, error) {
	return _UsingTellor.Contract.StakerCount(&_UsingTellor.CallOpts)
}

// TellorContract is a free data retrieval call binding the contract method 0xa339ac74.
//
// Solidity: function tellorContract() view returns(address)
func (_UsingTellor *UsingTellorCaller) TellorContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "tellorContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TellorContract is a free data retrieval call binding the contract method 0xa339ac74.
//
// Solidity: function tellorContract() view returns(address)
func (_UsingTellor *UsingTellorSession) TellorContract() (common.Address, error) {
	return _UsingTellor.Contract.TellorContract(&_UsingTellor.CallOpts)
}

// TellorContract is a free data retrieval call binding the contract method 0xa339ac74.
//
// Solidity: function tellorContract() view returns(address)
func (_UsingTellor *UsingTellorCallerSession) TellorContract() (common.Address, error) {
	return _UsingTellor.Contract.TellorContract(&_UsingTellor.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(uint256)
func (_UsingTellor *UsingTellorCaller) TimeOfLastNewValue(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "timeOfLastNewValue")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(uint256)
func (_UsingTellor *UsingTellorSession) TimeOfLastNewValue() (*big.Int, error) {
	return _UsingTellor.Contract.TimeOfLastNewValue(&_UsingTellor.CallOpts)
}

// TimeOfLastNewValue is a free data retrieval call binding the contract method 0x6fd4f229.
//
// Solidity: function timeOfLastNewValue() view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) TimeOfLastNewValue() (*big.Int, error) {
	return _UsingTellor.Contract.TimeOfLastNewValue(&_UsingTellor.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(uint256)
func (_UsingTellor *UsingTellorCaller) TimeTarget(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "timeTarget")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(uint256)
func (_UsingTellor *UsingTellorSession) TimeTarget() (*big.Int, error) {
	return _UsingTellor.Contract.TimeTarget(&_UsingTellor.CallOpts)
}

// TimeTarget is a free data retrieval call binding the contract method 0x6fc37811.
//
// Solidity: function timeTarget() view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) TimeTarget() (*big.Int, error) {
	return _UsingTellor.Contract.TimeTarget(&_UsingTellor.CallOpts)
}

// TotalTip is a free data retrieval call binding the contract method 0x44b12aea.
//
// Solidity: function totalTip(uint256 requestID) view returns(uint256)
func (_UsingTellor *UsingTellorCaller) TotalTip(opts *bind.CallOpts, requestID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _UsingTellor.contract.Call(opts, &out, "totalTip", requestID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTip is a free data retrieval call binding the contract method 0x44b12aea.
//
// Solidity: function totalTip(uint256 requestID) view returns(uint256)
func (_UsingTellor *UsingTellorSession) TotalTip(requestID *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.TotalTip(&_UsingTellor.CallOpts, requestID)
}

// TotalTip is a free data retrieval call binding the contract method 0x44b12aea.
//
// Solidity: function totalTip(uint256 requestID) view returns(uint256)
func (_UsingTellor *UsingTellorCallerSession) TotalTip(requestID *big.Int) (*big.Int, error) {
	return _UsingTellor.Contract.TotalTip(&_UsingTellor.CallOpts, requestID)
}
