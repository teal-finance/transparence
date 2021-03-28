// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package creamfinance

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

// CrTokenABI is the input ABI used to generate the binding from.
const CrTokenABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying_\",\"type\":\"address\"},{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals_\",\"type\":\"uint8\"},{\"internalType\":\"addresspayable\",\"name\":\"admin_\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"becomeImplementationData\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"cashPrior\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"NewImplementation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"benefactor\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"}],\"name\":\"_addReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"implementation_\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"allowResign\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"becomeImplementationData\",\"type\":\"bytes\"}],\"name\":\"_setImplementation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"delegateToViewImplementation\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"implementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isCToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractCTokenInterface\",\"name\":\"cTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CrToken is an auto generated Go binding around an Ethereum contract.
type CrToken struct {
	CrTokenCaller     // Read-only binding to the contract
	CrTokenTransactor // Write-only binding to the contract
	CrTokenFilterer   // Log filterer for contract events
}

// CrTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CrTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CrTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CrTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CrTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CrTokenSession struct {
	Contract     *CrToken          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CrTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CrTokenCallerSession struct {
	Contract *CrTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CrTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CrTokenTransactorSession struct {
	Contract     *CrTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CrTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CrTokenRaw struct {
	Contract *CrToken // Generic contract binding to access the raw methods on
}

// CrTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CrTokenCallerRaw struct {
	Contract *CrTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CrTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CrTokenTransactorRaw struct {
	Contract *CrTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCrToken creates a new instance of CrToken, bound to a specific deployed contract.
func NewCrToken(address common.Address, backend bind.ContractBackend) (*CrToken, error) {
	contract, err := bindCrToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CrToken{CrTokenCaller: CrTokenCaller{contract: contract}, CrTokenTransactor: CrTokenTransactor{contract: contract}, CrTokenFilterer: CrTokenFilterer{contract: contract}}, nil
}

// NewCrTokenCaller creates a new read-only instance of CrToken, bound to a specific deployed contract.
func NewCrTokenCaller(address common.Address, caller bind.ContractCaller) (*CrTokenCaller, error) {
	contract, err := bindCrToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CrTokenCaller{contract: contract}, nil
}

// NewCrTokenTransactor creates a new write-only instance of CrToken, bound to a specific deployed contract.
func NewCrTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CrTokenTransactor, error) {
	contract, err := bindCrToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CrTokenTransactor{contract: contract}, nil
}

// NewCrTokenFilterer creates a new log filterer instance of CrToken, bound to a specific deployed contract.
func NewCrTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CrTokenFilterer, error) {
	contract, err := bindCrToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CrTokenFilterer{contract: contract}, nil
}

// bindCrToken binds a generic wrapper to an already deployed contract.
func bindCrToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CrTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrToken *CrTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrToken.Contract.CrTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrToken *CrTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrToken.Contract.CrTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrToken *CrTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrToken.Contract.CrTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CrToken *CrTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CrToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CrToken *CrTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CrToken *CrTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CrToken.Contract.contract.Transact(opts, method, params...)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_CrToken *CrTokenCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_CrToken *CrTokenSession) AccrualBlockNumber() (*big.Int, error) {
	return _CrToken.Contract.AccrualBlockNumber(&_CrToken.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_CrToken *CrTokenCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _CrToken.Contract.AccrualBlockNumber(&_CrToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrToken *CrTokenCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrToken *CrTokenSession) Admin() (common.Address, error) {
	return _CrToken.Contract.Admin(&_CrToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CrToken *CrTokenCallerSession) Admin() (common.Address, error) {
	return _CrToken.Contract.Admin(&_CrToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CrToken *CrTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CrToken *CrTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CrToken.Contract.Allowance(&_CrToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CrToken *CrTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CrToken.Contract.Allowance(&_CrToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CrToken *CrTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CrToken *CrTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CrToken.Contract.BalanceOf(&_CrToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CrToken *CrTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CrToken.Contract.BalanceOf(&_CrToken.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_CrToken *CrTokenCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_CrToken *CrTokenSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _CrToken.Contract.BorrowBalanceStored(&_CrToken.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_CrToken *CrTokenCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _CrToken.Contract.BorrowBalanceStored(&_CrToken.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_CrToken *CrTokenCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_CrToken *CrTokenSession) BorrowIndex() (*big.Int, error) {
	return _CrToken.Contract.BorrowIndex(&_CrToken.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_CrToken *CrTokenCallerSession) BorrowIndex() (*big.Int, error) {
	return _CrToken.Contract.BorrowIndex(&_CrToken.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_CrToken *CrTokenCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_CrToken *CrTokenSession) BorrowRatePerBlock() (*big.Int, error) {
	return _CrToken.Contract.BorrowRatePerBlock(&_CrToken.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_CrToken *CrTokenCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _CrToken.Contract.BorrowRatePerBlock(&_CrToken.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_CrToken *CrTokenCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_CrToken *CrTokenSession) Comptroller() (common.Address, error) {
	return _CrToken.Contract.Comptroller(&_CrToken.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_CrToken *CrTokenCallerSession) Comptroller() (common.Address, error) {
	return _CrToken.Contract.Comptroller(&_CrToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CrToken *CrTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CrToken *CrTokenSession) Decimals() (uint8, error) {
	return _CrToken.Contract.Decimals(&_CrToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CrToken *CrTokenCallerSession) Decimals() (uint8, error) {
	return _CrToken.Contract.Decimals(&_CrToken.CallOpts)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_CrToken *CrTokenCaller) DelegateToViewImplementation(opts *bind.CallOpts, data []byte) ([]byte, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "delegateToViewImplementation", data)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_CrToken *CrTokenSession) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _CrToken.Contract.DelegateToViewImplementation(&_CrToken.CallOpts, data)
}

// DelegateToViewImplementation is a free data retrieval call binding the contract method 0x4487152f.
//
// Solidity: function delegateToViewImplementation(bytes data) view returns(bytes)
func (_CrToken *CrTokenCallerSession) DelegateToViewImplementation(data []byte) ([]byte, error) {
	return _CrToken.Contract.DelegateToViewImplementation(&_CrToken.CallOpts, data)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_CrToken *CrTokenCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_CrToken *CrTokenSession) ExchangeRateStored() (*big.Int, error) {
	return _CrToken.Contract.ExchangeRateStored(&_CrToken.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_CrToken *CrTokenCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _CrToken.Contract.ExchangeRateStored(&_CrToken.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_CrToken *CrTokenCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "getAccountSnapshot", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_CrToken *CrTokenSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _CrToken.Contract.GetAccountSnapshot(&_CrToken.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_CrToken *CrTokenCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _CrToken.Contract.GetAccountSnapshot(&_CrToken.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_CrToken *CrTokenCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_CrToken *CrTokenSession) GetCash() (*big.Int, error) {
	return _CrToken.Contract.GetCash(&_CrToken.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_CrToken *CrTokenCallerSession) GetCash() (*big.Int, error) {
	return _CrToken.Contract.GetCash(&_CrToken.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_CrToken *CrTokenCaller) Implementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "implementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_CrToken *CrTokenSession) Implementation() (common.Address, error) {
	return _CrToken.Contract.Implementation(&_CrToken.CallOpts)
}

// Implementation is a free data retrieval call binding the contract method 0x5c60da1b.
//
// Solidity: function implementation() view returns(address)
func (_CrToken *CrTokenCallerSession) Implementation() (common.Address, error) {
	return _CrToken.Contract.Implementation(&_CrToken.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_CrToken *CrTokenCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_CrToken *CrTokenSession) InterestRateModel() (common.Address, error) {
	return _CrToken.Contract.InterestRateModel(&_CrToken.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_CrToken *CrTokenCallerSession) InterestRateModel() (common.Address, error) {
	return _CrToken.Contract.InterestRateModel(&_CrToken.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_CrToken *CrTokenCaller) IsCToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "isCToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_CrToken *CrTokenSession) IsCToken() (bool, error) {
	return _CrToken.Contract.IsCToken(&_CrToken.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_CrToken *CrTokenCallerSession) IsCToken() (bool, error) {
	return _CrToken.Contract.IsCToken(&_CrToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CrToken *CrTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CrToken *CrTokenSession) Name() (string, error) {
	return _CrToken.Contract.Name(&_CrToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CrToken *CrTokenCallerSession) Name() (string, error) {
	return _CrToken.Contract.Name(&_CrToken.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_CrToken *CrTokenCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_CrToken *CrTokenSession) PendingAdmin() (common.Address, error) {
	return _CrToken.Contract.PendingAdmin(&_CrToken.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_CrToken *CrTokenCallerSession) PendingAdmin() (common.Address, error) {
	return _CrToken.Contract.PendingAdmin(&_CrToken.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_CrToken *CrTokenCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_CrToken *CrTokenSession) ReserveFactorMantissa() (*big.Int, error) {
	return _CrToken.Contract.ReserveFactorMantissa(&_CrToken.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_CrToken *CrTokenCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _CrToken.Contract.ReserveFactorMantissa(&_CrToken.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CrToken *CrTokenCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CrToken *CrTokenSession) SupplyRatePerBlock() (*big.Int, error) {
	return _CrToken.Contract.SupplyRatePerBlock(&_CrToken.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CrToken *CrTokenCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _CrToken.Contract.SupplyRatePerBlock(&_CrToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CrToken *CrTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CrToken *CrTokenSession) Symbol() (string, error) {
	return _CrToken.Contract.Symbol(&_CrToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CrToken *CrTokenCallerSession) Symbol() (string, error) {
	return _CrToken.Contract.Symbol(&_CrToken.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_CrToken *CrTokenCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_CrToken *CrTokenSession) TotalBorrows() (*big.Int, error) {
	return _CrToken.Contract.TotalBorrows(&_CrToken.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_CrToken *CrTokenCallerSession) TotalBorrows() (*big.Int, error) {
	return _CrToken.Contract.TotalBorrows(&_CrToken.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_CrToken *CrTokenCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_CrToken *CrTokenSession) TotalReserves() (*big.Int, error) {
	return _CrToken.Contract.TotalReserves(&_CrToken.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_CrToken *CrTokenCallerSession) TotalReserves() (*big.Int, error) {
	return _CrToken.Contract.TotalReserves(&_CrToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CrToken *CrTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CrToken *CrTokenSession) TotalSupply() (*big.Int, error) {
	return _CrToken.Contract.TotalSupply(&_CrToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CrToken *CrTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CrToken.Contract.TotalSupply(&_CrToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CrToken *CrTokenCaller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CrToken.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CrToken *CrTokenSession) Underlying() (common.Address, error) {
	return _CrToken.Contract.Underlying(&_CrToken.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CrToken *CrTokenCallerSession) Underlying() (common.Address, error) {
	return _CrToken.Contract.Underlying(&_CrToken.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_CrToken *CrTokenTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_CrToken *CrTokenSession) AcceptAdmin() (*types.Transaction, error) {
	return _CrToken.Contract.AcceptAdmin(&_CrToken.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_CrToken *CrTokenTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _CrToken.Contract.AcceptAdmin(&_CrToken.TransactOpts)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_CrToken *CrTokenTransactor) AddReserves(opts *bind.TransactOpts, addAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "_addReserves", addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_CrToken *CrTokenSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.AddReserves(&_CrToken.TransactOpts, addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0x3e941010.
//
// Solidity: function _addReserves(uint256 addAmount) returns(uint256)
func (_CrToken *CrTokenTransactorSession) AddReserves(addAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.AddReserves(&_CrToken.TransactOpts, addAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_CrToken *CrTokenTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_CrToken *CrTokenSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.ReduceReserves(&_CrToken.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_CrToken *CrTokenTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.ReduceReserves(&_CrToken.TransactOpts, reduceAmount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_CrToken *CrTokenTransactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_CrToken *CrTokenSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _CrToken.Contract.SetComptroller(&_CrToken.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_CrToken *CrTokenTransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _CrToken.Contract.SetComptroller(&_CrToken.TransactOpts, newComptroller)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_CrToken *CrTokenTransactor) SetImplementation(opts *bind.TransactOpts, implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "_setImplementation", implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_CrToken *CrTokenSession) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _CrToken.Contract.SetImplementation(&_CrToken.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetImplementation is a paid mutator transaction binding the contract method 0x555bcc40.
//
// Solidity: function _setImplementation(address implementation_, bool allowResign, bytes becomeImplementationData) returns()
func (_CrToken *CrTokenTransactorSession) SetImplementation(implementation_ common.Address, allowResign bool, becomeImplementationData []byte) (*types.Transaction, error) {
	return _CrToken.Contract.SetImplementation(&_CrToken.TransactOpts, implementation_, allowResign, becomeImplementationData)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_CrToken *CrTokenTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_CrToken *CrTokenSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _CrToken.Contract.SetInterestRateModel(&_CrToken.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_CrToken *CrTokenTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _CrToken.Contract.SetInterestRateModel(&_CrToken.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_CrToken *CrTokenTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_CrToken *CrTokenSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _CrToken.Contract.SetPendingAdmin(&_CrToken.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_CrToken *CrTokenTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _CrToken.Contract.SetPendingAdmin(&_CrToken.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_CrToken *CrTokenTransactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_CrToken *CrTokenSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.SetReserveFactor(&_CrToken.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_CrToken *CrTokenTransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.SetReserveFactor(&_CrToken.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_CrToken *CrTokenTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_CrToken *CrTokenSession) AccrueInterest() (*types.Transaction, error) {
	return _CrToken.Contract.AccrueInterest(&_CrToken.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_CrToken *CrTokenTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _CrToken.Contract.AccrueInterest(&_CrToken.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CrToken *CrTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CrToken *CrTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.Approve(&_CrToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CrToken *CrTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.Approve(&_CrToken.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_CrToken *CrTokenTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_CrToken *CrTokenSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _CrToken.Contract.BalanceOfUnderlying(&_CrToken.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_CrToken *CrTokenTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _CrToken.Contract.BalanceOfUnderlying(&_CrToken.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_CrToken *CrTokenTransactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_CrToken *CrTokenSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.Borrow(&_CrToken.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_CrToken *CrTokenTransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.Borrow(&_CrToken.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_CrToken *CrTokenTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_CrToken *CrTokenSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _CrToken.Contract.BorrowBalanceCurrent(&_CrToken.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_CrToken *CrTokenTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _CrToken.Contract.BorrowBalanceCurrent(&_CrToken.TransactOpts, account)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_CrToken *CrTokenTransactor) DelegateToImplementation(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "delegateToImplementation", data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_CrToken *CrTokenSession) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _CrToken.Contract.DelegateToImplementation(&_CrToken.TransactOpts, data)
}

// DelegateToImplementation is a paid mutator transaction binding the contract method 0x0933c1ed.
//
// Solidity: function delegateToImplementation(bytes data) returns(bytes)
func (_CrToken *CrTokenTransactorSession) DelegateToImplementation(data []byte) (*types.Transaction, error) {
	return _CrToken.Contract.DelegateToImplementation(&_CrToken.TransactOpts, data)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_CrToken *CrTokenTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_CrToken *CrTokenSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _CrToken.Contract.ExchangeRateCurrent(&_CrToken.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_CrToken *CrTokenTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _CrToken.Contract.ExchangeRateCurrent(&_CrToken.TransactOpts)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_CrToken *CrTokenTransactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "liquidateBorrow", borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_CrToken *CrTokenSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CrToken.Contract.LiquidateBorrow(&_CrToken.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_CrToken *CrTokenTransactorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CrToken.Contract.LiquidateBorrow(&_CrToken.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_CrToken *CrTokenTransactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "mint", mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_CrToken *CrTokenSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.Mint(&_CrToken.TransactOpts, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_CrToken *CrTokenTransactorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.Mint(&_CrToken.TransactOpts, mintAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_CrToken *CrTokenTransactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_CrToken *CrTokenSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.Redeem(&_CrToken.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_CrToken *CrTokenTransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.Redeem(&_CrToken.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_CrToken *CrTokenTransactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_CrToken *CrTokenSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.RedeemUnderlying(&_CrToken.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_CrToken *CrTokenTransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.RedeemUnderlying(&_CrToken.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_CrToken *CrTokenTransactor) RepayBorrow(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "repayBorrow", repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_CrToken *CrTokenSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.RepayBorrow(&_CrToken.TransactOpts, repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_CrToken *CrTokenTransactorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.RepayBorrow(&_CrToken.TransactOpts, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_CrToken *CrTokenTransactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "repayBorrowBehalf", borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_CrToken *CrTokenSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.RepayBorrowBehalf(&_CrToken.TransactOpts, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_CrToken *CrTokenTransactorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.RepayBorrowBehalf(&_CrToken.TransactOpts, borrower, repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_CrToken *CrTokenTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_CrToken *CrTokenSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.Seize(&_CrToken.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_CrToken *CrTokenTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.Seize(&_CrToken.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_CrToken *CrTokenTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_CrToken *CrTokenSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _CrToken.Contract.TotalBorrowsCurrent(&_CrToken.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_CrToken *CrTokenTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _CrToken.Contract.TotalBorrowsCurrent(&_CrToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_CrToken *CrTokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_CrToken *CrTokenSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.Transfer(&_CrToken.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_CrToken *CrTokenTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.Transfer(&_CrToken.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_CrToken *CrTokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrToken.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_CrToken *CrTokenSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.TransferFrom(&_CrToken.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_CrToken *CrTokenTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CrToken.Contract.TransferFrom(&_CrToken.TransactOpts, src, dst, amount)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CrToken *CrTokenTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _CrToken.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CrToken *CrTokenSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CrToken.Contract.Fallback(&_CrToken.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_CrToken *CrTokenTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _CrToken.Contract.Fallback(&_CrToken.TransactOpts, calldata)
}

// CrTokenAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the CrToken contract.
type CrTokenAccrueInterestIterator struct {
	Event *CrTokenAccrueInterest // Event containing the contract specifics and raw log

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
func (it *CrTokenAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenAccrueInterest)
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
		it.Event = new(CrTokenAccrueInterest)
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
func (it *CrTokenAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenAccrueInterest represents a AccrueInterest event raised by the CrToken contract.
type CrTokenAccrueInterest struct {
	CashPrior           *big.Int
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_CrToken *CrTokenFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*CrTokenAccrueInterestIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &CrTokenAccrueInterestIterator{contract: _CrToken.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_CrToken *CrTokenFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *CrTokenAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenAccrueInterest)
				if err := _CrToken.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0x4dec04e750ca11537cabcd8a9eab06494de08da3735bc8871cd41250e190bc04.
//
// Solidity: event AccrueInterest(uint256 cashPrior, uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_CrToken *CrTokenFilterer) ParseAccrueInterest(log types.Log) (*CrTokenAccrueInterest, error) {
	event := new(CrTokenAccrueInterest)
	if err := _CrToken.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CrToken contract.
type CrTokenApprovalIterator struct {
	Event *CrTokenApproval // Event containing the contract specifics and raw log

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
func (it *CrTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenApproval)
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
		it.Event = new(CrTokenApproval)
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
func (it *CrTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenApproval represents a Approval event raised by the CrToken contract.
type CrTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_CrToken *CrTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CrTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CrTokenApprovalIterator{contract: _CrToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_CrToken *CrTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CrTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenApproval)
				if err := _CrToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_CrToken *CrTokenFilterer) ParseApproval(log types.Log) (*CrTokenApproval, error) {
	event := new(CrTokenApproval)
	if err := _CrToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the CrToken contract.
type CrTokenBorrowIterator struct {
	Event *CrTokenBorrow // Event containing the contract specifics and raw log

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
func (it *CrTokenBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenBorrow)
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
		it.Event = new(CrTokenBorrow)
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
func (it *CrTokenBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenBorrow represents a Borrow event raised by the CrToken contract.
type CrTokenBorrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CrToken *CrTokenFilterer) FilterBorrow(opts *bind.FilterOpts) (*CrTokenBorrowIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &CrTokenBorrowIterator{contract: _CrToken.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CrToken *CrTokenFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *CrTokenBorrow) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenBorrow)
				if err := _CrToken.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CrToken *CrTokenFilterer) ParseBorrow(log types.Log) (*CrTokenBorrow, error) {
	event := new(CrTokenBorrow)
	if err := _CrToken.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the CrToken contract.
type CrTokenFailureIterator struct {
	Event *CrTokenFailure // Event containing the contract specifics and raw log

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
func (it *CrTokenFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenFailure)
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
		it.Event = new(CrTokenFailure)
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
func (it *CrTokenFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenFailure represents a Failure event raised by the CrToken contract.
type CrTokenFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_CrToken *CrTokenFilterer) FilterFailure(opts *bind.FilterOpts) (*CrTokenFailureIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &CrTokenFailureIterator{contract: _CrToken.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_CrToken *CrTokenFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *CrTokenFailure) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenFailure)
				if err := _CrToken.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_CrToken *CrTokenFilterer) ParseFailure(log types.Log) (*CrTokenFailure, error) {
	event := new(CrTokenFailure)
	if err := _CrToken.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenLiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the CrToken contract.
type CrTokenLiquidateBorrowIterator struct {
	Event *CrTokenLiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *CrTokenLiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenLiquidateBorrow)
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
		it.Event = new(CrTokenLiquidateBorrow)
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
func (it *CrTokenLiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenLiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenLiquidateBorrow represents a LiquidateBorrow event raised by the CrToken contract.
type CrTokenLiquidateBorrow struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	CTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrow is a free log retrieval operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_CrToken *CrTokenFilterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*CrTokenLiquidateBorrowIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &CrTokenLiquidateBorrowIterator{contract: _CrToken.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_CrToken *CrTokenFilterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *CrTokenLiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenLiquidateBorrow)
				if err := _CrToken.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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

// ParseLiquidateBorrow is a log parse operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_CrToken *CrTokenFilterer) ParseLiquidateBorrow(log types.Log) (*CrTokenLiquidateBorrow, error) {
	event := new(CrTokenLiquidateBorrow)
	if err := _CrToken.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the CrToken contract.
type CrTokenMintIterator struct {
	Event *CrTokenMint // Event containing the contract specifics and raw log

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
func (it *CrTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenMint)
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
		it.Event = new(CrTokenMint)
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
func (it *CrTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenMint represents a Mint event raised by the CrToken contract.
type CrTokenMint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_CrToken *CrTokenFilterer) FilterMint(opts *bind.FilterOpts) (*CrTokenMintIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &CrTokenMintIterator{contract: _CrToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_CrToken *CrTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *CrTokenMint) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenMint)
				if err := _CrToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_CrToken *CrTokenFilterer) ParseMint(log types.Log) (*CrTokenMint, error) {
	event := new(CrTokenMint)
	if err := _CrToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the CrToken contract.
type CrTokenNewAdminIterator struct {
	Event *CrTokenNewAdmin // Event containing the contract specifics and raw log

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
func (it *CrTokenNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenNewAdmin)
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
		it.Event = new(CrTokenNewAdmin)
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
func (it *CrTokenNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenNewAdmin represents a NewAdmin event raised by the CrToken contract.
type CrTokenNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_CrToken *CrTokenFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*CrTokenNewAdminIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &CrTokenNewAdminIterator{contract: _CrToken.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_CrToken *CrTokenFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *CrTokenNewAdmin) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenNewAdmin)
				if err := _CrToken.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_CrToken *CrTokenFilterer) ParseNewAdmin(log types.Log) (*CrTokenNewAdmin, error) {
	event := new(CrTokenNewAdmin)
	if err := _CrToken.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenNewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the CrToken contract.
type CrTokenNewComptrollerIterator struct {
	Event *CrTokenNewComptroller // Event containing the contract specifics and raw log

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
func (it *CrTokenNewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenNewComptroller)
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
		it.Event = new(CrTokenNewComptroller)
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
func (it *CrTokenNewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenNewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenNewComptroller represents a NewComptroller event raised by the CrToken contract.
type CrTokenNewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_CrToken *CrTokenFilterer) FilterNewComptroller(opts *bind.FilterOpts) (*CrTokenNewComptrollerIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &CrTokenNewComptrollerIterator{contract: _CrToken.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_CrToken *CrTokenFilterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *CrTokenNewComptroller) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenNewComptroller)
				if err := _CrToken.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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

// ParseNewComptroller is a log parse operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_CrToken *CrTokenFilterer) ParseNewComptroller(log types.Log) (*CrTokenNewComptroller, error) {
	event := new(CrTokenNewComptroller)
	if err := _CrToken.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenNewImplementationIterator is returned from FilterNewImplementation and is used to iterate over the raw logs and unpacked data for NewImplementation events raised by the CrToken contract.
type CrTokenNewImplementationIterator struct {
	Event *CrTokenNewImplementation // Event containing the contract specifics and raw log

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
func (it *CrTokenNewImplementationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenNewImplementation)
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
		it.Event = new(CrTokenNewImplementation)
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
func (it *CrTokenNewImplementationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenNewImplementationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenNewImplementation represents a NewImplementation event raised by the CrToken contract.
type CrTokenNewImplementation struct {
	OldImplementation common.Address
	NewImplementation common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewImplementation is a free log retrieval operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_CrToken *CrTokenFilterer) FilterNewImplementation(opts *bind.FilterOpts) (*CrTokenNewImplementationIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return &CrTokenNewImplementationIterator{contract: _CrToken.contract, event: "NewImplementation", logs: logs, sub: sub}, nil
}

// WatchNewImplementation is a free log subscription operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_CrToken *CrTokenFilterer) WatchNewImplementation(opts *bind.WatchOpts, sink chan<- *CrTokenNewImplementation) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "NewImplementation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenNewImplementation)
				if err := _CrToken.contract.UnpackLog(event, "NewImplementation", log); err != nil {
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

// ParseNewImplementation is a log parse operation binding the contract event 0xd604de94d45953f9138079ec1b82d533cb2160c906d1076d1f7ed54befbca97a.
//
// Solidity: event NewImplementation(address oldImplementation, address newImplementation)
func (_CrToken *CrTokenFilterer) ParseNewImplementation(log types.Log) (*CrTokenNewImplementation, error) {
	event := new(CrTokenNewImplementation)
	if err := _CrToken.contract.UnpackLog(event, "NewImplementation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenNewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the CrToken contract.
type CrTokenNewMarketInterestRateModelIterator struct {
	Event *CrTokenNewMarketInterestRateModel // Event containing the contract specifics and raw log

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
func (it *CrTokenNewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenNewMarketInterestRateModel)
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
		it.Event = new(CrTokenNewMarketInterestRateModel)
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
func (it *CrTokenNewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenNewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenNewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the CrToken contract.
type CrTokenNewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_CrToken *CrTokenFilterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*CrTokenNewMarketInterestRateModelIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &CrTokenNewMarketInterestRateModelIterator{contract: _CrToken.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_CrToken *CrTokenFilterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *CrTokenNewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenNewMarketInterestRateModel)
				if err := _CrToken.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
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

// ParseNewMarketInterestRateModel is a log parse operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_CrToken *CrTokenFilterer) ParseNewMarketInterestRateModel(log types.Log) (*CrTokenNewMarketInterestRateModel, error) {
	event := new(CrTokenNewMarketInterestRateModel)
	if err := _CrToken.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the CrToken contract.
type CrTokenNewPendingAdminIterator struct {
	Event *CrTokenNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *CrTokenNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenNewPendingAdmin)
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
		it.Event = new(CrTokenNewPendingAdmin)
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
func (it *CrTokenNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenNewPendingAdmin represents a NewPendingAdmin event raised by the CrToken contract.
type CrTokenNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_CrToken *CrTokenFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*CrTokenNewPendingAdminIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &CrTokenNewPendingAdminIterator{contract: _CrToken.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_CrToken *CrTokenFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *CrTokenNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenNewPendingAdmin)
				if err := _CrToken.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_CrToken *CrTokenFilterer) ParseNewPendingAdmin(log types.Log) (*CrTokenNewPendingAdmin, error) {
	event := new(CrTokenNewPendingAdmin)
	if err := _CrToken.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenNewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the CrToken contract.
type CrTokenNewReserveFactorIterator struct {
	Event *CrTokenNewReserveFactor // Event containing the contract specifics and raw log

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
func (it *CrTokenNewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenNewReserveFactor)
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
		it.Event = new(CrTokenNewReserveFactor)
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
func (it *CrTokenNewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenNewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenNewReserveFactor represents a NewReserveFactor event raised by the CrToken contract.
type CrTokenNewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_CrToken *CrTokenFilterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*CrTokenNewReserveFactorIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &CrTokenNewReserveFactorIterator{contract: _CrToken.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_CrToken *CrTokenFilterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *CrTokenNewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenNewReserveFactor)
				if err := _CrToken.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
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

// ParseNewReserveFactor is a log parse operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_CrToken *CrTokenFilterer) ParseNewReserveFactor(log types.Log) (*CrTokenNewReserveFactor, error) {
	event := new(CrTokenNewReserveFactor)
	if err := _CrToken.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the CrToken contract.
type CrTokenRedeemIterator struct {
	Event *CrTokenRedeem // Event containing the contract specifics and raw log

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
func (it *CrTokenRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenRedeem)
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
		it.Event = new(CrTokenRedeem)
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
func (it *CrTokenRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenRedeem represents a Redeem event raised by the CrToken contract.
type CrTokenRedeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_CrToken *CrTokenFilterer) FilterRedeem(opts *bind.FilterOpts) (*CrTokenRedeemIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &CrTokenRedeemIterator{contract: _CrToken.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_CrToken *CrTokenFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *CrTokenRedeem) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenRedeem)
				if err := _CrToken.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_CrToken *CrTokenFilterer) ParseRedeem(log types.Log) (*CrTokenRedeem, error) {
	event := new(CrTokenRedeem)
	if err := _CrToken.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the CrToken contract.
type CrTokenRepayBorrowIterator struct {
	Event *CrTokenRepayBorrow // Event containing the contract specifics and raw log

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
func (it *CrTokenRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenRepayBorrow)
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
		it.Event = new(CrTokenRepayBorrow)
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
func (it *CrTokenRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenRepayBorrow represents a RepayBorrow event raised by the CrToken contract.
type CrTokenRepayBorrow struct {
	Payer          common.Address
	Borrower       common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CrToken *CrTokenFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*CrTokenRepayBorrowIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &CrTokenRepayBorrowIterator{contract: _CrToken.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CrToken *CrTokenFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *CrTokenRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenRepayBorrow)
				if err := _CrToken.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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

// ParseRepayBorrow is a log parse operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CrToken *CrTokenFilterer) ParseRepayBorrow(log types.Log) (*CrTokenRepayBorrow, error) {
	event := new(CrTokenRepayBorrow)
	if err := _CrToken.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenReservesAddedIterator is returned from FilterReservesAdded and is used to iterate over the raw logs and unpacked data for ReservesAdded events raised by the CrToken contract.
type CrTokenReservesAddedIterator struct {
	Event *CrTokenReservesAdded // Event containing the contract specifics and raw log

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
func (it *CrTokenReservesAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenReservesAdded)
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
		it.Event = new(CrTokenReservesAdded)
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
func (it *CrTokenReservesAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenReservesAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenReservesAdded represents a ReservesAdded event raised by the CrToken contract.
type CrTokenReservesAdded struct {
	Benefactor       common.Address
	AddAmount        *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesAdded is a free log retrieval operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_CrToken *CrTokenFilterer) FilterReservesAdded(opts *bind.FilterOpts) (*CrTokenReservesAddedIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return &CrTokenReservesAddedIterator{contract: _CrToken.contract, event: "ReservesAdded", logs: logs, sub: sub}, nil
}

// WatchReservesAdded is a free log subscription operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_CrToken *CrTokenFilterer) WatchReservesAdded(opts *bind.WatchOpts, sink chan<- *CrTokenReservesAdded) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "ReservesAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenReservesAdded)
				if err := _CrToken.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
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

// ParseReservesAdded is a log parse operation binding the contract event 0xa91e67c5ea634cd43a12c5a482724b03de01e85ca68702a53d0c2f45cb7c1dc5.
//
// Solidity: event ReservesAdded(address benefactor, uint256 addAmount, uint256 newTotalReserves)
func (_CrToken *CrTokenFilterer) ParseReservesAdded(log types.Log) (*CrTokenReservesAdded, error) {
	event := new(CrTokenReservesAdded)
	if err := _CrToken.contract.UnpackLog(event, "ReservesAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the CrToken contract.
type CrTokenReservesReducedIterator struct {
	Event *CrTokenReservesReduced // Event containing the contract specifics and raw log

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
func (it *CrTokenReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenReservesReduced)
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
		it.Event = new(CrTokenReservesReduced)
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
func (it *CrTokenReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenReservesReduced represents a ReservesReduced event raised by the CrToken contract.
type CrTokenReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_CrToken *CrTokenFilterer) FilterReservesReduced(opts *bind.FilterOpts) (*CrTokenReservesReducedIterator, error) {

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &CrTokenReservesReducedIterator{contract: _CrToken.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_CrToken *CrTokenFilterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *CrTokenReservesReduced) (event.Subscription, error) {

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenReservesReduced)
				if err := _CrToken.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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

// ParseReservesReduced is a log parse operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_CrToken *CrTokenFilterer) ParseReservesReduced(log types.Log) (*CrTokenReservesReduced, error) {
	event := new(CrTokenReservesReduced)
	if err := _CrToken.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CrTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CrToken contract.
type CrTokenTransferIterator struct {
	Event *CrTokenTransfer // Event containing the contract specifics and raw log

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
func (it *CrTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CrTokenTransfer)
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
		it.Event = new(CrTokenTransfer)
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
func (it *CrTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CrTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CrTokenTransfer represents a Transfer event raised by the CrToken contract.
type CrTokenTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CrToken *CrTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CrTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CrTokenTransferIterator{contract: _CrToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CrToken *CrTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CrTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CrToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CrTokenTransfer)
				if err := _CrToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CrToken *CrTokenFilterer) ParseTransfer(log types.Log) (*CrTokenTransfer, error) {
	event := new(CrTokenTransfer)
	if err := _CrToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
