// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package controller

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

// ComptrollerABI is the input ABI used to generate the binding from.
const ComptrollerABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"action\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"ActionPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSpeed\",\"type\":\"uint256\"}],\"name\":\"CompSpeedUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"creditLimit\",\"type\":\"uint256\"}],\"name\":\"CreditLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedBorrowerComp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"supplier\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compDelta\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"compSupplyIndex\",\"type\":\"uint256\"}],\"name\":\"DistributedSupplierComp\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"MarketExited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"MarketListed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBorrowCap\",\"type\":\"uint256\"}],\"name\":\"NewBorrowCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldBorrowCapGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"NewBorrowCapGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCloseFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCloseFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCollateralFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewCollateralFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldLiquidationIncentiveMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"NewLiquidationIncentive\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPauseGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"NewPauseGuardian\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"oldPriceOracle\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractPriceOracle\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"NewPriceOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newSupplyCap\",\"type\":\"uint256\"}],\"name\":\"NewSupplyCap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldSupplyCapGuardian\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newSupplyCapGuardian\",\"type\":\"address\"}],\"name\":\"NewSupplyCapGuardian\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractUnitroller\",\"name\":\"unitroller\",\"type\":\"address\"}],\"name\":\"_become\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"_mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newBorrowCapGuardian\",\"type\":\"address\"}],\"name\":\"_setBorrowCapGuardian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setBorrowPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newCloseFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCloseFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setCollateralFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"speeds\",\"type\":\"uint256[]\"}],\"name\":\"_setCompSpeeds\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"protocol\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"creditLimit\",\"type\":\"uint256\"}],\"name\":\"_setCreditLimit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newLiquidationIncentiveMantissa\",\"type\":\"uint256\"}],\"name\":\"_setLiquidationIncentive\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newBorrowCaps\",\"type\":\"uint256[]\"}],\"name\":\"_setMarketBorrowCaps\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"newSupplyCaps\",\"type\":\"uint256[]\"}],\"name\":\"_setMarketSupplyCaps\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setMintPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPauseGuardian\",\"type\":\"address\"}],\"name\":\"_setPauseGuardian\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"newOracle\",\"type\":\"address\"}],\"name\":\"_setPriceOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setSeizePaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newSupplyCapGuardian\",\"type\":\"address\"}],\"name\":\"_setSupplyCapGuardian\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"state\",\"type\":\"bool\"}],\"name\":\"_setTransferPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"_supportMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"contractCToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"contractCToken\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowCapGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"borrowGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractCToken\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"checkMembership\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"},{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"claimComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"holders\",\"type\":\"address[]\"},{\"internalType\":\"contractCToken[]\",\"name\":\"cTokens\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"borrowers\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"suppliers\",\"type\":\"bool\"}],\"name\":\"claimComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"holder\",\"type\":\"address\"}],\"name\":\"claimComp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"closeFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compBorrowState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compBorrowerIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compClaimThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"compInitialIndex\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"\",\"type\":\"uint224\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSpeeds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSupplierIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"compSupplyState\",\"outputs\":[{\"internalType\":\"uint224\",\"name\":\"index\",\"type\":\"uint224\"},{\"internalType\":\"uint32\",\"name\":\"block\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"creditLimits\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenAddress\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"contractCToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCompAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenModify\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getHypotheticalAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"liquidationIncentiveMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isListed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralFactorMantissa\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mintGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualMintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractPriceOracle\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pauseGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingComptrollerImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"seizeGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyCapGuardian\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"supplyCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transferGuardianPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Comptroller is an auto generated Go binding around an Ethereum contract.
type Comptroller struct {
	ComptrollerCaller     // Read-only binding to the contract
	ComptrollerTransactor // Write-only binding to the contract
	ComptrollerFilterer   // Log filterer for contract events
}

// ComptrollerCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComptrollerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComptrollerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComptrollerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComptrollerSession struct {
	Contract     *Comptroller      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ComptrollerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComptrollerCallerSession struct {
	Contract *ComptrollerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ComptrollerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComptrollerTransactorSession struct {
	Contract     *ComptrollerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ComptrollerRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComptrollerRaw struct {
	Contract *Comptroller // Generic contract binding to access the raw methods on
}

// ComptrollerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComptrollerCallerRaw struct {
	Contract *ComptrollerCaller // Generic read-only contract binding to access the raw methods on
}

// ComptrollerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComptrollerTransactorRaw struct {
	Contract *ComptrollerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComptroller creates a new instance of Comptroller, bound to a specific deployed contract.
func NewComptroller(address common.Address, backend bind.ContractBackend) (*Comptroller, error) {
	contract, err := bindComptroller(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Comptroller{ComptrollerCaller: ComptrollerCaller{contract: contract}, ComptrollerTransactor: ComptrollerTransactor{contract: contract}, ComptrollerFilterer: ComptrollerFilterer{contract: contract}}, nil
}

// NewComptrollerCaller creates a new read-only instance of Comptroller, bound to a specific deployed contract.
func NewComptrollerCaller(address common.Address, caller bind.ContractCaller) (*ComptrollerCaller, error) {
	contract, err := bindComptroller(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerCaller{contract: contract}, nil
}

// NewComptrollerTransactor creates a new write-only instance of Comptroller, bound to a specific deployed contract.
func NewComptrollerTransactor(address common.Address, transactor bind.ContractTransactor) (*ComptrollerTransactor, error) {
	contract, err := bindComptroller(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerTransactor{contract: contract}, nil
}

// NewComptrollerFilterer creates a new log filterer instance of Comptroller, bound to a specific deployed contract.
func NewComptrollerFilterer(address common.Address, filterer bind.ContractFilterer) (*ComptrollerFilterer, error) {
	contract, err := bindComptroller(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComptrollerFilterer{contract: contract}, nil
}

// bindComptroller binds a generic wrapper to an already deployed contract.
func bindComptroller(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComptrollerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comptroller *ComptrollerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Comptroller.Contract.ComptrollerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comptroller *ComptrollerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comptroller.Contract.ComptrollerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comptroller *ComptrollerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comptroller.Contract.ComptrollerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Comptroller *ComptrollerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Comptroller.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Comptroller *ComptrollerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Comptroller.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Comptroller *ComptrollerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Comptroller.Contract.contract.Transact(opts, method, params...)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCaller) BorrowGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "_borrowGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerSession) BorrowGuardianPaused() (bool, error) {
	return _Comptroller.Contract.BorrowGuardianPaused(&_Comptroller.CallOpts)
}

// BorrowGuardianPaused is a free data retrieval call binding the contract method 0xe6653f3d.
//
// Solidity: function _borrowGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCallerSession) BorrowGuardianPaused() (bool, error) {
	return _Comptroller.Contract.BorrowGuardianPaused(&_Comptroller.CallOpts)
}

// MintGuardianPausedAlias is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCaller) MintGuardianPausedAlias(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "_mintGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPausedAlias is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerSession) MintGuardianPausedAlias() (bool, error) {
	return _Comptroller.Contract.MintGuardianPausedAlias(&_Comptroller.CallOpts)
}

// MintGuardianPausedAlias is a free data retrieval call binding the contract method 0x3c94786f.
//
// Solidity: function _mintGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCallerSession) MintGuardianPausedAlias() (bool, error) {
	return _Comptroller.Contract.MintGuardianPausedAlias(&_Comptroller.CallOpts)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Comptroller *ComptrollerCaller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "accountAssets", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Comptroller *ComptrollerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Comptroller.Contract.AccountAssets(&_Comptroller.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Comptroller *ComptrollerCallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Comptroller.Contract.AccountAssets(&_Comptroller.CallOpts, arg0, arg1)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Comptroller *ComptrollerCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Comptroller *ComptrollerSession) Admin() (common.Address, error) {
	return _Comptroller.Contract.Admin(&_Comptroller.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Comptroller *ComptrollerCallerSession) Admin() (common.Address, error) {
	return _Comptroller.Contract.Admin(&_Comptroller.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Comptroller *ComptrollerCaller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "allMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Comptroller *ComptrollerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Comptroller.Contract.AllMarkets(&_Comptroller.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Comptroller *ComptrollerCallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Comptroller.Contract.AllMarkets(&_Comptroller.CallOpts, arg0)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Comptroller *ComptrollerCaller) BorrowCapGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "borrowCapGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Comptroller *ComptrollerSession) BorrowCapGuardian() (common.Address, error) {
	return _Comptroller.Contract.BorrowCapGuardian(&_Comptroller.CallOpts)
}

// BorrowCapGuardian is a free data retrieval call binding the contract method 0x21af4569.
//
// Solidity: function borrowCapGuardian() view returns(address)
func (_Comptroller *ComptrollerCallerSession) BorrowCapGuardian() (common.Address, error) {
	return _Comptroller.Contract.BorrowCapGuardian(&_Comptroller.CallOpts)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) BorrowCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "borrowCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.BorrowCaps(&_Comptroller.CallOpts, arg0)
}

// BorrowCaps is a free data retrieval call binding the contract method 0x4a584432.
//
// Solidity: function borrowCaps(address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) BorrowCaps(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.BorrowCaps(&_Comptroller.CallOpts, arg0)
}

// BorrowGuardianPausedAlias is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Comptroller *ComptrollerCaller) BorrowGuardianPausedAlias(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "borrowGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BorrowGuardianPausedAlias is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Comptroller *ComptrollerSession) BorrowGuardianPausedAlias(arg0 common.Address) (bool, error) {
	return _Comptroller.Contract.BorrowGuardianPausedAlias(&_Comptroller.CallOpts, arg0)
}

// BorrowGuardianPausedAlias is a free data retrieval call binding the contract method 0x6d154ea5.
//
// Solidity: function borrowGuardianPaused(address ) view returns(bool)
func (_Comptroller *ComptrollerCallerSession) BorrowGuardianPausedAlias(arg0 common.Address) (bool, error) {
	return _Comptroller.Contract.BorrowGuardianPausedAlias(&_Comptroller.CallOpts, arg0)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_Comptroller *ComptrollerCaller) CheckMembership(opts *bind.CallOpts, account common.Address, cToken common.Address) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "checkMembership", account, cToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_Comptroller *ComptrollerSession) CheckMembership(account common.Address, cToken common.Address) (bool, error) {
	return _Comptroller.Contract.CheckMembership(&_Comptroller.CallOpts, account, cToken)
}

// CheckMembership is a free data retrieval call binding the contract method 0x929fe9a1.
//
// Solidity: function checkMembership(address account, address cToken) view returns(bool)
func (_Comptroller *ComptrollerCallerSession) CheckMembership(account common.Address, cToken common.Address) (bool, error) {
	return _Comptroller.Contract.CheckMembership(&_Comptroller.CallOpts, account, cToken)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Comptroller *ComptrollerCaller) CloseFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "closeFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Comptroller *ComptrollerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Comptroller.Contract.CloseFactorMantissa(&_Comptroller.CallOpts)
}

// CloseFactorMantissa is a free data retrieval call binding the contract method 0xe8755446.
//
// Solidity: function closeFactorMantissa() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) CloseFactorMantissa() (*big.Int, error) {
	return _Comptroller.Contract.CloseFactorMantissa(&_Comptroller.CallOpts)
}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) CompAccrued(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "compAccrued", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) CompAccrued(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.CompAccrued(&_Comptroller.CallOpts, arg0)
}

// CompAccrued is a free data retrieval call binding the contract method 0xcc7ebdc4.
//
// Solidity: function compAccrued(address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) CompAccrued(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.CompAccrued(&_Comptroller.CallOpts, arg0)
}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Comptroller *ComptrollerCaller) CompBorrowState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "compBorrowState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})

	outstruct.Index = out[0].(*big.Int)
	outstruct.Block = out[1].(uint32)

	return *outstruct, err

}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Comptroller *ComptrollerSession) CompBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Comptroller.Contract.CompBorrowState(&_Comptroller.CallOpts, arg0)
}

// CompBorrowState is a free data retrieval call binding the contract method 0x8c57804e.
//
// Solidity: function compBorrowState(address ) view returns(uint224 index, uint32 block)
func (_Comptroller *ComptrollerCallerSession) CompBorrowState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Comptroller.Contract.CompBorrowState(&_Comptroller.CallOpts, arg0)
}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) CompBorrowerIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "compBorrowerIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) CompBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.CompBorrowerIndex(&_Comptroller.CallOpts, arg0, arg1)
}

// CompBorrowerIndex is a free data retrieval call binding the contract method 0xca0af043.
//
// Solidity: function compBorrowerIndex(address , address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) CompBorrowerIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.CompBorrowerIndex(&_Comptroller.CallOpts, arg0, arg1)
}

// CompClaimThreshold is a free data retrieval call binding the contract method 0x747026c9.
//
// Solidity: function compClaimThreshold() view returns(uint256)
func (_Comptroller *ComptrollerCaller) CompClaimThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "compClaimThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompClaimThreshold is a free data retrieval call binding the contract method 0x747026c9.
//
// Solidity: function compClaimThreshold() view returns(uint256)
func (_Comptroller *ComptrollerSession) CompClaimThreshold() (*big.Int, error) {
	return _Comptroller.Contract.CompClaimThreshold(&_Comptroller.CallOpts)
}

// CompClaimThreshold is a free data retrieval call binding the contract method 0x747026c9.
//
// Solidity: function compClaimThreshold() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) CompClaimThreshold() (*big.Int, error) {
	return _Comptroller.Contract.CompClaimThreshold(&_Comptroller.CallOpts)
}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_Comptroller *ComptrollerCaller) CompInitialIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "compInitialIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_Comptroller *ComptrollerSession) CompInitialIndex() (*big.Int, error) {
	return _Comptroller.Contract.CompInitialIndex(&_Comptroller.CallOpts)
}

// CompInitialIndex is a free data retrieval call binding the contract method 0xa7f0e231.
//
// Solidity: function compInitialIndex() view returns(uint224)
func (_Comptroller *ComptrollerCallerSession) CompInitialIndex() (*big.Int, error) {
	return _Comptroller.Contract.CompInitialIndex(&_Comptroller.CallOpts)
}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) CompSpeeds(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "compSpeeds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) CompSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.CompSpeeds(&_Comptroller.CallOpts, arg0)
}

// CompSpeeds is a free data retrieval call binding the contract method 0x1d7b33d7.
//
// Solidity: function compSpeeds(address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) CompSpeeds(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.CompSpeeds(&_Comptroller.CallOpts, arg0)
}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) CompSupplierIndex(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "compSupplierIndex", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) CompSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.CompSupplierIndex(&_Comptroller.CallOpts, arg0, arg1)
}

// CompSupplierIndex is a free data retrieval call binding the contract method 0xb21be7fd.
//
// Solidity: function compSupplierIndex(address , address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) CompSupplierIndex(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.CompSupplierIndex(&_Comptroller.CallOpts, arg0, arg1)
}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Comptroller *ComptrollerCaller) CompSupplyState(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "compSupplyState", arg0)

	outstruct := new(struct {
		Index *big.Int
		Block uint32
	})

	outstruct.Index = out[0].(*big.Int)
	outstruct.Block = out[1].(uint32)

	return *outstruct, err

}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Comptroller *ComptrollerSession) CompSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Comptroller.Contract.CompSupplyState(&_Comptroller.CallOpts, arg0)
}

// CompSupplyState is a free data retrieval call binding the contract method 0x6b79c38d.
//
// Solidity: function compSupplyState(address ) view returns(uint224 index, uint32 block)
func (_Comptroller *ComptrollerCallerSession) CompSupplyState(arg0 common.Address) (struct {
	Index *big.Int
	Block uint32
}, error) {
	return _Comptroller.Contract.CompSupplyState(&_Comptroller.CallOpts, arg0)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Comptroller *ComptrollerCaller) ComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "comptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Comptroller *ComptrollerSession) ComptrollerImplementation() (common.Address, error) {
	return _Comptroller.Contract.ComptrollerImplementation(&_Comptroller.CallOpts)
}

// ComptrollerImplementation is a free data retrieval call binding the contract method 0xbb82aa5e.
//
// Solidity: function comptrollerImplementation() view returns(address)
func (_Comptroller *ComptrollerCallerSession) ComptrollerImplementation() (common.Address, error) {
	return _Comptroller.Contract.ComptrollerImplementation(&_Comptroller.CallOpts)
}

// CreditLimits is a free data retrieval call binding the contract method 0x99bc1873.
//
// Solidity: function creditLimits(address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) CreditLimits(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "creditLimits", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CreditLimits is a free data retrieval call binding the contract method 0x99bc1873.
//
// Solidity: function creditLimits(address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) CreditLimits(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.CreditLimits(&_Comptroller.CallOpts, arg0)
}

// CreditLimits is a free data retrieval call binding the contract method 0x99bc1873.
//
// Solidity: function creditLimits(address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) CreditLimits(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.CreditLimits(&_Comptroller.CallOpts, arg0)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerCaller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "getAccountLiquidity", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Comptroller.Contract.GetAccountLiquidity(&_Comptroller.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerCallerSession) GetAccountLiquidity(account common.Address) (*big.Int, *big.Int, *big.Int, error) {
	return _Comptroller.Contract.GetAccountLiquidity(&_Comptroller.CallOpts, account)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Comptroller *ComptrollerCaller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "getAllMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Comptroller *ComptrollerSession) GetAllMarkets() ([]common.Address, error) {
	return _Comptroller.Contract.GetAllMarkets(&_Comptroller.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Comptroller *ComptrollerCallerSession) GetAllMarkets() ([]common.Address, error) {
	return _Comptroller.Contract.GetAllMarkets(&_Comptroller.CallOpts)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Comptroller *ComptrollerCaller) GetAssetsIn(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "getAssetsIn", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Comptroller *ComptrollerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Comptroller.Contract.GetAssetsIn(&_Comptroller.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Comptroller *ComptrollerCallerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Comptroller.Contract.GetAssetsIn(&_Comptroller.CallOpts, account)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Comptroller *ComptrollerCaller) GetBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "getBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Comptroller *ComptrollerSession) GetBlockNumber() (*big.Int, error) {
	return _Comptroller.Contract.GetBlockNumber(&_Comptroller.CallOpts)
}

// GetBlockNumber is a free data retrieval call binding the contract method 0x42cbb15c.
//
// Solidity: function getBlockNumber() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) GetBlockNumber() (*big.Int, error) {
	return _Comptroller.Contract.GetBlockNumber(&_Comptroller.CallOpts)
}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_Comptroller *ComptrollerCaller) GetCompAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "getCompAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_Comptroller *ComptrollerSession) GetCompAddress() (common.Address, error) {
	return _Comptroller.Contract.GetCompAddress(&_Comptroller.CallOpts)
}

// GetCompAddress is a free data retrieval call binding the contract method 0x9d1b5a0a.
//
// Solidity: function getCompAddress() view returns(address)
func (_Comptroller *ComptrollerCallerSession) GetCompAddress() (common.Address, error) {
	return _Comptroller.Contract.GetCompAddress(&_Comptroller.CallOpts)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerCaller) GetHypotheticalAccountLiquidity(opts *bind.CallOpts, account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "getHypotheticalAccountLiquidity", account, cTokenModify, redeemTokens, borrowAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerSession) GetHypotheticalAccountLiquidity(account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Comptroller.Contract.GetHypotheticalAccountLiquidity(&_Comptroller.CallOpts, account, cTokenModify, redeemTokens, borrowAmount)
}

// GetHypotheticalAccountLiquidity is a free data retrieval call binding the contract method 0x4e79238f.
//
// Solidity: function getHypotheticalAccountLiquidity(address account, address cTokenModify, uint256 redeemTokens, uint256 borrowAmount) view returns(uint256, uint256, uint256)
func (_Comptroller *ComptrollerCallerSession) GetHypotheticalAccountLiquidity(account common.Address, cTokenModify common.Address, redeemTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, *big.Int, error) {
	return _Comptroller.Contract.GetHypotheticalAccountLiquidity(&_Comptroller.CallOpts, account, cTokenModify, redeemTokens, borrowAmount)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Comptroller *ComptrollerCaller) IsComptroller(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "isComptroller")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Comptroller *ComptrollerSession) IsComptroller() (bool, error) {
	return _Comptroller.Contract.IsComptroller(&_Comptroller.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_Comptroller *ComptrollerCallerSession) IsComptroller() (bool, error) {
	return _Comptroller.Contract.IsComptroller(&_Comptroller.CallOpts)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comptroller *ComptrollerCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "liquidateCalculateSeizeTokens", cTokenBorrowed, cTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comptroller *ComptrollerSession) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Comptroller.Contract.LiquidateCalculateSeizeTokens(&_Comptroller.CallOpts, cTokenBorrowed, cTokenCollateral, actualRepayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 actualRepayAmount) view returns(uint256, uint256)
func (_Comptroller *ComptrollerCallerSession) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Comptroller.Contract.LiquidateCalculateSeizeTokens(&_Comptroller.CallOpts, cTokenBorrowed, cTokenCollateral, actualRepayAmount)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Comptroller *ComptrollerCaller) LiquidationIncentiveMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "liquidationIncentiveMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Comptroller *ComptrollerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Comptroller.Contract.LiquidationIncentiveMantissa(&_Comptroller.CallOpts)
}

// LiquidationIncentiveMantissa is a free data retrieval call binding the contract method 0x4ada90af.
//
// Solidity: function liquidationIncentiveMantissa() view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) LiquidationIncentiveMantissa() (*big.Int, error) {
	return _Comptroller.Contract.LiquidationIncentiveMantissa(&_Comptroller.CallOpts)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa)
func (_Comptroller *ComptrollerCaller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
}, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		IsListed                 bool
		CollateralFactorMantissa *big.Int
	})

	outstruct.IsListed = out[0].(bool)
	outstruct.CollateralFactorMantissa = out[1].(*big.Int)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa)
func (_Comptroller *ComptrollerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
}, error) {
	return _Comptroller.Contract.Markets(&_Comptroller.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(bool isListed, uint256 collateralFactorMantissa)
func (_Comptroller *ComptrollerCallerSession) Markets(arg0 common.Address) (struct {
	IsListed                 bool
	CollateralFactorMantissa *big.Int
}, error) {
	return _Comptroller.Contract.Markets(&_Comptroller.CallOpts, arg0)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Comptroller *ComptrollerCaller) MintGuardianPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "mintGuardianPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Comptroller *ComptrollerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Comptroller.Contract.MintGuardianPaused(&_Comptroller.CallOpts, arg0)
}

// MintGuardianPaused is a free data retrieval call binding the contract method 0x731f0c2b.
//
// Solidity: function mintGuardianPaused(address ) view returns(bool)
func (_Comptroller *ComptrollerCallerSession) MintGuardianPaused(arg0 common.Address) (bool, error) {
	return _Comptroller.Contract.MintGuardianPaused(&_Comptroller.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Comptroller *ComptrollerCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Comptroller *ComptrollerSession) Oracle() (common.Address, error) {
	return _Comptroller.Contract.Oracle(&_Comptroller.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Comptroller *ComptrollerCallerSession) Oracle() (common.Address, error) {
	return _Comptroller.Contract.Oracle(&_Comptroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Comptroller *ComptrollerCaller) PauseGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "pauseGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Comptroller *ComptrollerSession) PauseGuardian() (common.Address, error) {
	return _Comptroller.Contract.PauseGuardian(&_Comptroller.CallOpts)
}

// PauseGuardian is a free data retrieval call binding the contract method 0x24a3d622.
//
// Solidity: function pauseGuardian() view returns(address)
func (_Comptroller *ComptrollerCallerSession) PauseGuardian() (common.Address, error) {
	return _Comptroller.Contract.PauseGuardian(&_Comptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Comptroller *ComptrollerCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Comptroller *ComptrollerSession) PendingAdmin() (common.Address, error) {
	return _Comptroller.Contract.PendingAdmin(&_Comptroller.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_Comptroller *ComptrollerCallerSession) PendingAdmin() (common.Address, error) {
	return _Comptroller.Contract.PendingAdmin(&_Comptroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Comptroller *ComptrollerCaller) PendingComptrollerImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "pendingComptrollerImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Comptroller *ComptrollerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Comptroller.Contract.PendingComptrollerImplementation(&_Comptroller.CallOpts)
}

// PendingComptrollerImplementation is a free data retrieval call binding the contract method 0xdcfbc0c7.
//
// Solidity: function pendingComptrollerImplementation() view returns(address)
func (_Comptroller *ComptrollerCallerSession) PendingComptrollerImplementation() (common.Address, error) {
	return _Comptroller.Contract.PendingComptrollerImplementation(&_Comptroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCaller) SeizeGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "seizeGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerSession) SeizeGuardianPaused() (bool, error) {
	return _Comptroller.Contract.SeizeGuardianPaused(&_Comptroller.CallOpts)
}

// SeizeGuardianPaused is a free data retrieval call binding the contract method 0xac0b0bb7.
//
// Solidity: function seizeGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCallerSession) SeizeGuardianPaused() (bool, error) {
	return _Comptroller.Contract.SeizeGuardianPaused(&_Comptroller.CallOpts)
}

// SupplyCapGuardian is a free data retrieval call binding the contract method 0xa979f0c5.
//
// Solidity: function supplyCapGuardian() view returns(address)
func (_Comptroller *ComptrollerCaller) SupplyCapGuardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "supplyCapGuardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupplyCapGuardian is a free data retrieval call binding the contract method 0xa979f0c5.
//
// Solidity: function supplyCapGuardian() view returns(address)
func (_Comptroller *ComptrollerSession) SupplyCapGuardian() (common.Address, error) {
	return _Comptroller.Contract.SupplyCapGuardian(&_Comptroller.CallOpts)
}

// SupplyCapGuardian is a free data retrieval call binding the contract method 0xa979f0c5.
//
// Solidity: function supplyCapGuardian() view returns(address)
func (_Comptroller *ComptrollerCallerSession) SupplyCapGuardian() (common.Address, error) {
	return _Comptroller.Contract.SupplyCapGuardian(&_Comptroller.CallOpts)
}

// SupplyCaps is a free data retrieval call binding the contract method 0x02c3bcbb.
//
// Solidity: function supplyCaps(address ) view returns(uint256)
func (_Comptroller *ComptrollerCaller) SupplyCaps(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "supplyCaps", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyCaps is a free data retrieval call binding the contract method 0x02c3bcbb.
//
// Solidity: function supplyCaps(address ) view returns(uint256)
func (_Comptroller *ComptrollerSession) SupplyCaps(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.SupplyCaps(&_Comptroller.CallOpts, arg0)
}

// SupplyCaps is a free data retrieval call binding the contract method 0x02c3bcbb.
//
// Solidity: function supplyCaps(address ) view returns(uint256)
func (_Comptroller *ComptrollerCallerSession) SupplyCaps(arg0 common.Address) (*big.Int, error) {
	return _Comptroller.Contract.SupplyCaps(&_Comptroller.CallOpts, arg0)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCaller) TransferGuardianPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Comptroller.contract.Call(opts, &out, "transferGuardianPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerSession) TransferGuardianPaused() (bool, error) {
	return _Comptroller.Contract.TransferGuardianPaused(&_Comptroller.CallOpts)
}

// TransferGuardianPaused is a free data retrieval call binding the contract method 0x87f76303.
//
// Solidity: function transferGuardianPaused() view returns(bool)
func (_Comptroller *ComptrollerCallerSession) TransferGuardianPaused() (bool, error) {
	return _Comptroller.Contract.TransferGuardianPaused(&_Comptroller.CallOpts)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Comptroller *ComptrollerTransactor) Become(opts *bind.TransactOpts, unitroller common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_become", unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Comptroller *ComptrollerSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.Become(&_Comptroller.TransactOpts, unitroller)
}

// Become is a paid mutator transaction binding the contract method 0x1d504dc6.
//
// Solidity: function _become(address unitroller) returns()
func (_Comptroller *ComptrollerTransactorSession) Become(unitroller common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.Become(&_Comptroller.TransactOpts, unitroller)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Comptroller *ComptrollerTransactor) SetBorrowCapGuardian(opts *bind.TransactOpts, newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setBorrowCapGuardian", newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Comptroller *ComptrollerSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetBorrowCapGuardian(&_Comptroller.TransactOpts, newBorrowCapGuardian)
}

// SetBorrowCapGuardian is a paid mutator transaction binding the contract method 0x391957d7.
//
// Solidity: function _setBorrowCapGuardian(address newBorrowCapGuardian) returns()
func (_Comptroller *ComptrollerTransactorSession) SetBorrowCapGuardian(newBorrowCapGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetBorrowCapGuardian(&_Comptroller.TransactOpts, newBorrowCapGuardian)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_Comptroller *ComptrollerTransactor) SetBorrowPaused(opts *bind.TransactOpts, cToken common.Address, state bool) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setBorrowPaused", cToken, state)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_Comptroller *ComptrollerSession) SetBorrowPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetBorrowPaused(&_Comptroller.TransactOpts, cToken, state)
}

// SetBorrowPaused is a paid mutator transaction binding the contract method 0x18c882a5.
//
// Solidity: function _setBorrowPaused(address cToken, bool state) returns(bool)
func (_Comptroller *ComptrollerTransactorSession) SetBorrowPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetBorrowPaused(&_Comptroller.TransactOpts, cToken, state)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetCloseFactor(opts *bind.TransactOpts, newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setCloseFactor", newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCloseFactor(&_Comptroller.TransactOpts, newCloseFactorMantissa)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x317b0b77.
//
// Solidity: function _setCloseFactor(uint256 newCloseFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetCloseFactor(newCloseFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCloseFactor(&_Comptroller.TransactOpts, newCloseFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetCollateralFactor(opts *bind.TransactOpts, cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setCollateralFactor", cToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerSession) SetCollateralFactor(cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCollateralFactor(&_Comptroller.TransactOpts, cToken, newCollateralFactorMantissa)
}

// SetCollateralFactor is a paid mutator transaction binding the contract method 0xe4028eee.
//
// Solidity: function _setCollateralFactor(address cToken, uint256 newCollateralFactorMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetCollateralFactor(cToken common.Address, newCollateralFactorMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCollateralFactor(&_Comptroller.TransactOpts, cToken, newCollateralFactorMantissa)
}

// SetCompSpeeds is a paid mutator transaction binding the contract method 0x7e60c7a0.
//
// Solidity: function _setCompSpeeds(address[] cTokens, uint256[] speeds) returns()
func (_Comptroller *ComptrollerTransactor) SetCompSpeeds(opts *bind.TransactOpts, cTokens []common.Address, speeds []*big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setCompSpeeds", cTokens, speeds)
}

// SetCompSpeeds is a paid mutator transaction binding the contract method 0x7e60c7a0.
//
// Solidity: function _setCompSpeeds(address[] cTokens, uint256[] speeds) returns()
func (_Comptroller *ComptrollerSession) SetCompSpeeds(cTokens []common.Address, speeds []*big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCompSpeeds(&_Comptroller.TransactOpts, cTokens, speeds)
}

// SetCompSpeeds is a paid mutator transaction binding the contract method 0x7e60c7a0.
//
// Solidity: function _setCompSpeeds(address[] cTokens, uint256[] speeds) returns()
func (_Comptroller *ComptrollerTransactorSession) SetCompSpeeds(cTokens []common.Address, speeds []*big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCompSpeeds(&_Comptroller.TransactOpts, cTokens, speeds)
}

// SetCreditLimit is a paid mutator transaction binding the contract method 0xbc93082b.
//
// Solidity: function _setCreditLimit(address protocol, uint256 creditLimit) returns()
func (_Comptroller *ComptrollerTransactor) SetCreditLimit(opts *bind.TransactOpts, protocol common.Address, creditLimit *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setCreditLimit", protocol, creditLimit)
}

// SetCreditLimit is a paid mutator transaction binding the contract method 0xbc93082b.
//
// Solidity: function _setCreditLimit(address protocol, uint256 creditLimit) returns()
func (_Comptroller *ComptrollerSession) SetCreditLimit(protocol common.Address, creditLimit *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCreditLimit(&_Comptroller.TransactOpts, protocol, creditLimit)
}

// SetCreditLimit is a paid mutator transaction binding the contract method 0xbc93082b.
//
// Solidity: function _setCreditLimit(address protocol, uint256 creditLimit) returns()
func (_Comptroller *ComptrollerTransactorSession) SetCreditLimit(protocol common.Address, creditLimit *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetCreditLimit(&_Comptroller.TransactOpts, protocol, creditLimit)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetLiquidationIncentive(opts *bind.TransactOpts, newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setLiquidationIncentive", newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Comptroller *ComptrollerSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetLiquidationIncentive(&_Comptroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x4fd42e17.
//
// Solidity: function _setLiquidationIncentive(uint256 newLiquidationIncentiveMantissa) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetLiquidationIncentive(newLiquidationIncentiveMantissa *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetLiquidationIncentive(&_Comptroller.TransactOpts, newLiquidationIncentiveMantissa)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] cTokens, uint256[] newBorrowCaps) returns()
func (_Comptroller *ComptrollerTransactor) SetMarketBorrowCaps(opts *bind.TransactOpts, cTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setMarketBorrowCaps", cTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] cTokens, uint256[] newBorrowCaps) returns()
func (_Comptroller *ComptrollerSession) SetMarketBorrowCaps(cTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMarketBorrowCaps(&_Comptroller.TransactOpts, cTokens, newBorrowCaps)
}

// SetMarketBorrowCaps is a paid mutator transaction binding the contract method 0x607ef6c1.
//
// Solidity: function _setMarketBorrowCaps(address[] cTokens, uint256[] newBorrowCaps) returns()
func (_Comptroller *ComptrollerTransactorSession) SetMarketBorrowCaps(cTokens []common.Address, newBorrowCaps []*big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMarketBorrowCaps(&_Comptroller.TransactOpts, cTokens, newBorrowCaps)
}

// SetMarketSupplyCaps is a paid mutator transaction binding the contract method 0x51a485e4.
//
// Solidity: function _setMarketSupplyCaps(address[] cTokens, uint256[] newSupplyCaps) returns()
func (_Comptroller *ComptrollerTransactor) SetMarketSupplyCaps(opts *bind.TransactOpts, cTokens []common.Address, newSupplyCaps []*big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setMarketSupplyCaps", cTokens, newSupplyCaps)
}

// SetMarketSupplyCaps is a paid mutator transaction binding the contract method 0x51a485e4.
//
// Solidity: function _setMarketSupplyCaps(address[] cTokens, uint256[] newSupplyCaps) returns()
func (_Comptroller *ComptrollerSession) SetMarketSupplyCaps(cTokens []common.Address, newSupplyCaps []*big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMarketSupplyCaps(&_Comptroller.TransactOpts, cTokens, newSupplyCaps)
}

// SetMarketSupplyCaps is a paid mutator transaction binding the contract method 0x51a485e4.
//
// Solidity: function _setMarketSupplyCaps(address[] cTokens, uint256[] newSupplyCaps) returns()
func (_Comptroller *ComptrollerTransactorSession) SetMarketSupplyCaps(cTokens []common.Address, newSupplyCaps []*big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMarketSupplyCaps(&_Comptroller.TransactOpts, cTokens, newSupplyCaps)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_Comptroller *ComptrollerTransactor) SetMintPaused(opts *bind.TransactOpts, cToken common.Address, state bool) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setMintPaused", cToken, state)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_Comptroller *ComptrollerSession) SetMintPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMintPaused(&_Comptroller.TransactOpts, cToken, state)
}

// SetMintPaused is a paid mutator transaction binding the contract method 0x3bcf7ec1.
//
// Solidity: function _setMintPaused(address cToken, bool state) returns(bool)
func (_Comptroller *ComptrollerTransactorSession) SetMintPaused(cToken common.Address, state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetMintPaused(&_Comptroller.TransactOpts, cToken, state)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetPauseGuardian(opts *bind.TransactOpts, newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setPauseGuardian", newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Comptroller *ComptrollerSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetPauseGuardian(&_Comptroller.TransactOpts, newPauseGuardian)
}

// SetPauseGuardian is a paid mutator transaction binding the contract method 0x5f5af1aa.
//
// Solidity: function _setPauseGuardian(address newPauseGuardian) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetPauseGuardian(newPauseGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetPauseGuardian(&_Comptroller.TransactOpts, newPauseGuardian)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SetPriceOracle(opts *bind.TransactOpts, newOracle common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setPriceOracle", newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Comptroller *ComptrollerSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetPriceOracle(&_Comptroller.TransactOpts, newOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x55ee1fe1.
//
// Solidity: function _setPriceOracle(address newOracle) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SetPriceOracle(newOracle common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetPriceOracle(&_Comptroller.TransactOpts, newOracle)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactor) SetSeizePaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setSeizePaused", state)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Comptroller *ComptrollerSession) SetSeizePaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetSeizePaused(&_Comptroller.TransactOpts, state)
}

// SetSeizePaused is a paid mutator transaction binding the contract method 0x2d70db78.
//
// Solidity: function _setSeizePaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactorSession) SetSeizePaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetSeizePaused(&_Comptroller.TransactOpts, state)
}

// SetSupplyCapGuardian is a paid mutator transaction binding the contract method 0x38b8f4c3.
//
// Solidity: function _setSupplyCapGuardian(address newSupplyCapGuardian) returns()
func (_Comptroller *ComptrollerTransactor) SetSupplyCapGuardian(opts *bind.TransactOpts, newSupplyCapGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setSupplyCapGuardian", newSupplyCapGuardian)
}

// SetSupplyCapGuardian is a paid mutator transaction binding the contract method 0x38b8f4c3.
//
// Solidity: function _setSupplyCapGuardian(address newSupplyCapGuardian) returns()
func (_Comptroller *ComptrollerSession) SetSupplyCapGuardian(newSupplyCapGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetSupplyCapGuardian(&_Comptroller.TransactOpts, newSupplyCapGuardian)
}

// SetSupplyCapGuardian is a paid mutator transaction binding the contract method 0x38b8f4c3.
//
// Solidity: function _setSupplyCapGuardian(address newSupplyCapGuardian) returns()
func (_Comptroller *ComptrollerTransactorSession) SetSupplyCapGuardian(newSupplyCapGuardian common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SetSupplyCapGuardian(&_Comptroller.TransactOpts, newSupplyCapGuardian)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactor) SetTransferPaused(opts *bind.TransactOpts, state bool) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_setTransferPaused", state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerSession) SetTransferPaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetTransferPaused(&_Comptroller.TransactOpts, state)
}

// SetTransferPaused is a paid mutator transaction binding the contract method 0x8ebf6364.
//
// Solidity: function _setTransferPaused(bool state) returns(bool)
func (_Comptroller *ComptrollerTransactorSession) SetTransferPaused(state bool) (*types.Transaction, error) {
	return _Comptroller.Contract.SetTransferPaused(&_Comptroller.TransactOpts, state)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SupportMarket(opts *bind.TransactOpts, cToken common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "_supportMarket", cToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Comptroller *ComptrollerSession) SupportMarket(cToken common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SupportMarket(&_Comptroller.TransactOpts, cToken)
}

// SupportMarket is a paid mutator transaction binding the contract method 0xa76b3fda.
//
// Solidity: function _supportMarket(address cToken) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SupportMarket(cToken common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.SupportMarket(&_Comptroller.TransactOpts, cToken)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactor) BorrowAllowed(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "borrowAllowed", cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Comptroller *ComptrollerSession) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.BorrowAllowed(&_Comptroller.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.BorrowAllowed(&_Comptroller.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Comptroller *ComptrollerTransactor) BorrowVerify(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "borrowVerify", cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Comptroller *ComptrollerSession) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.BorrowVerify(&_Comptroller.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_Comptroller *ComptrollerTransactorSession) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.BorrowVerify(&_Comptroller.TransactOpts, cToken, borrower, borrowAmount)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x1c3db2e0.
//
// Solidity: function claimComp(address holder, address[] cTokens) returns()
func (_Comptroller *ComptrollerTransactor) ClaimComp(opts *bind.TransactOpts, holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "claimComp", holder, cTokens)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x1c3db2e0.
//
// Solidity: function claimComp(address holder, address[] cTokens) returns()
func (_Comptroller *ComptrollerSession) ClaimComp(holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ClaimComp(&_Comptroller.TransactOpts, holder, cTokens)
}

// ClaimComp is a paid mutator transaction binding the contract method 0x1c3db2e0.
//
// Solidity: function claimComp(address holder, address[] cTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) ClaimComp(holder common.Address, cTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ClaimComp(&_Comptroller.TransactOpts, holder, cTokens)
}

// ClaimComp0 is a paid mutator transaction binding the contract method 0x6810dfa6.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_Comptroller *ComptrollerTransactor) ClaimComp0(opts *bind.TransactOpts, holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "claimComp0", holders, cTokens, borrowers, suppliers)
}

// ClaimComp0 is a paid mutator transaction binding the contract method 0x6810dfa6.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_Comptroller *ComptrollerSession) ClaimComp0(holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Comptroller.Contract.ClaimComp0(&_Comptroller.TransactOpts, holders, cTokens, borrowers, suppliers)
}

// ClaimComp0 is a paid mutator transaction binding the contract method 0x6810dfa6.
//
// Solidity: function claimComp(address[] holders, address[] cTokens, bool borrowers, bool suppliers) returns()
func (_Comptroller *ComptrollerTransactorSession) ClaimComp0(holders []common.Address, cTokens []common.Address, borrowers bool, suppliers bool) (*types.Transaction, error) {
	return _Comptroller.Contract.ClaimComp0(&_Comptroller.TransactOpts, holders, cTokens, borrowers, suppliers)
}

// ClaimComp1 is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_Comptroller *ComptrollerTransactor) ClaimComp1(opts *bind.TransactOpts, holder common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "claimComp1", holder)
}

// ClaimComp1 is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_Comptroller *ComptrollerSession) ClaimComp1(holder common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ClaimComp1(&_Comptroller.TransactOpts, holder)
}

// ClaimComp1 is a paid mutator transaction binding the contract method 0xe9af0292.
//
// Solidity: function claimComp(address holder) returns()
func (_Comptroller *ComptrollerTransactorSession) ClaimComp1(holder common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ClaimComp1(&_Comptroller.TransactOpts, holder)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Comptroller *ComptrollerTransactor) EnterMarkets(opts *bind.TransactOpts, cTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "enterMarkets", cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Comptroller *ComptrollerSession) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.EnterMarkets(&_Comptroller.TransactOpts, cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_Comptroller *ComptrollerTransactorSession) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.EnterMarkets(&_Comptroller.TransactOpts, cTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Comptroller *ComptrollerTransactor) ExitMarket(opts *bind.TransactOpts, cTokenAddress common.Address) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "exitMarket", cTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Comptroller *ComptrollerSession) ExitMarket(cTokenAddress common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ExitMarket(&_Comptroller.TransactOpts, cTokenAddress)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cTokenAddress) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) ExitMarket(cTokenAddress common.Address) (*types.Transaction, error) {
	return _Comptroller.Contract.ExitMarket(&_Comptroller.TransactOpts, cTokenAddress)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "liquidateBorrowAllowed", cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerSession) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.LiquidateBorrowAllowed(&_Comptroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.LiquidateBorrowAllowed(&_Comptroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerTransactor) LiquidateBorrowVerify(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "liquidateBorrowVerify", cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerSession) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.LiquidateBorrowVerify(&_Comptroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 actualRepayAmount, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, actualRepayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.LiquidateBorrowVerify(&_Comptroller.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, actualRepayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactor) MintAllowed(opts *bind.TransactOpts, cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "mintAllowed", cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Comptroller *ComptrollerSession) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.MintAllowed(&_Comptroller.TransactOpts, cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.MintAllowed(&_Comptroller.TransactOpts, cToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Comptroller *ComptrollerTransactor) MintVerify(opts *bind.TransactOpts, cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "mintVerify", cToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Comptroller *ComptrollerSession) MintVerify(cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.MintVerify(&_Comptroller.TransactOpts, cToken, minter, actualMintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 actualMintAmount, uint256 mintTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) MintVerify(cToken common.Address, minter common.Address, actualMintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.MintVerify(&_Comptroller.TransactOpts, cToken, minter, actualMintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactor) RedeemAllowed(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "redeemAllowed", cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Comptroller *ComptrollerSession) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RedeemAllowed(&_Comptroller.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RedeemAllowed(&_Comptroller.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Comptroller *ComptrollerTransactor) RedeemVerify(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "redeemVerify", cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Comptroller *ComptrollerSession) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RedeemVerify(&_Comptroller.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RedeemVerify(&_Comptroller.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactor) RepayBorrowAllowed(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "repayBorrowAllowed", cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerSession) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RepayBorrowAllowed(&_Comptroller.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RepayBorrowAllowed(&_Comptroller.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Comptroller *ComptrollerTransactor) RepayBorrowVerify(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "repayBorrowVerify", cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Comptroller *ComptrollerSession) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RepayBorrowVerify(&_Comptroller.TransactOpts, cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 actualRepayAmount, uint256 borrowerIndex) returns()
func (_Comptroller *ComptrollerTransactorSession) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, actualRepayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.RepayBorrowVerify(&_Comptroller.TransactOpts, cToken, payer, borrower, actualRepayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactor) SeizeAllowed(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "seizeAllowed", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Comptroller *ComptrollerSession) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SeizeAllowed(&_Comptroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SeizeAllowed(&_Comptroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerTransactor) SeizeVerify(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "seizeVerify", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerSession) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SeizeVerify(&_Comptroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.SeizeVerify(&_Comptroller.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactor) TransferAllowed(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "transferAllowed", cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Comptroller *ComptrollerSession) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.TransferAllowed(&_Comptroller.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_Comptroller *ComptrollerTransactorSession) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.TransferAllowed(&_Comptroller.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Comptroller *ComptrollerTransactor) TransferVerify(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.contract.Transact(opts, "transferVerify", cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Comptroller *ComptrollerSession) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.TransferVerify(&_Comptroller.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_Comptroller *ComptrollerTransactorSession) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Comptroller.Contract.TransferVerify(&_Comptroller.TransactOpts, cToken, src, dst, transferTokens)
}

// ComptrollerActionPausedIterator is returned from FilterActionPaused and is used to iterate over the raw logs and unpacked data for ActionPaused events raised by the Comptroller contract.
type ComptrollerActionPausedIterator struct {
	Event *ComptrollerActionPaused // Event containing the contract specifics and raw log

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
func (it *ComptrollerActionPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerActionPaused)
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
		it.Event = new(ComptrollerActionPaused)
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
func (it *ComptrollerActionPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerActionPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerActionPaused represents a ActionPaused event raised by the Comptroller contract.
type ComptrollerActionPaused struct {
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused is a free log retrieval operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) FilterActionPaused(opts *bind.FilterOpts) (*ComptrollerActionPausedIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return &ComptrollerActionPausedIterator{contract: _Comptroller.contract, event: "ActionPaused", logs: logs, sub: sub}, nil
}

// WatchActionPaused is a free log subscription operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) WatchActionPaused(opts *bind.WatchOpts, sink chan<- *ComptrollerActionPaused) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "ActionPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerActionPaused)
				if err := _Comptroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
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

// ParseActionPaused is a log parse operation binding the contract event 0xef159d9a32b2472e32b098f954f3ce62d232939f1c207070b584df1814de2de0.
//
// Solidity: event ActionPaused(string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) ParseActionPaused(log types.Log) (*ComptrollerActionPaused, error) {
	event := new(ComptrollerActionPaused)
	if err := _Comptroller.contract.UnpackLog(event, "ActionPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerActionPaused0Iterator is returned from FilterActionPaused0 and is used to iterate over the raw logs and unpacked data for ActionPaused0 events raised by the Comptroller contract.
type ComptrollerActionPaused0Iterator struct {
	Event *ComptrollerActionPaused0 // Event containing the contract specifics and raw log

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
func (it *ComptrollerActionPaused0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerActionPaused0)
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
		it.Event = new(ComptrollerActionPaused0)
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
func (it *ComptrollerActionPaused0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerActionPaused0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerActionPaused0 represents a ActionPaused0 event raised by the Comptroller contract.
type ComptrollerActionPaused0 struct {
	CToken     common.Address
	Action     string
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterActionPaused0 is a free log retrieval operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address cToken, string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) FilterActionPaused0(opts *bind.FilterOpts) (*ComptrollerActionPaused0Iterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return &ComptrollerActionPaused0Iterator{contract: _Comptroller.contract, event: "ActionPaused0", logs: logs, sub: sub}, nil
}

// WatchActionPaused0 is a free log subscription operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address cToken, string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) WatchActionPaused0(opts *bind.WatchOpts, sink chan<- *ComptrollerActionPaused0) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "ActionPaused0")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerActionPaused0)
				if err := _Comptroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
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

// ParseActionPaused0 is a log parse operation binding the contract event 0x71aec636243f9709bb0007ae15e9afb8150ab01716d75fd7573be5cc096e03b0.
//
// Solidity: event ActionPaused(address cToken, string action, bool pauseState)
func (_Comptroller *ComptrollerFilterer) ParseActionPaused0(log types.Log) (*ComptrollerActionPaused0, error) {
	event := new(ComptrollerActionPaused0)
	if err := _Comptroller.contract.UnpackLog(event, "ActionPaused0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerCompSpeedUpdatedIterator is returned from FilterCompSpeedUpdated and is used to iterate over the raw logs and unpacked data for CompSpeedUpdated events raised by the Comptroller contract.
type ComptrollerCompSpeedUpdatedIterator struct {
	Event *ComptrollerCompSpeedUpdated // Event containing the contract specifics and raw log

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
func (it *ComptrollerCompSpeedUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerCompSpeedUpdated)
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
		it.Event = new(ComptrollerCompSpeedUpdated)
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
func (it *ComptrollerCompSpeedUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerCompSpeedUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerCompSpeedUpdated represents a CompSpeedUpdated event raised by the Comptroller contract.
type ComptrollerCompSpeedUpdated struct {
	CToken   common.Address
	NewSpeed *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCompSpeedUpdated is a free log retrieval operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_Comptroller *ComptrollerFilterer) FilterCompSpeedUpdated(opts *bind.FilterOpts, cToken []common.Address) (*ComptrollerCompSpeedUpdatedIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "CompSpeedUpdated", cTokenRule)
	if err != nil {
		return nil, err
	}
	return &ComptrollerCompSpeedUpdatedIterator{contract: _Comptroller.contract, event: "CompSpeedUpdated", logs: logs, sub: sub}, nil
}

// WatchCompSpeedUpdated is a free log subscription operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_Comptroller *ComptrollerFilterer) WatchCompSpeedUpdated(opts *bind.WatchOpts, sink chan<- *ComptrollerCompSpeedUpdated, cToken []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "CompSpeedUpdated", cTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerCompSpeedUpdated)
				if err := _Comptroller.contract.UnpackLog(event, "CompSpeedUpdated", log); err != nil {
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

// ParseCompSpeedUpdated is a log parse operation binding the contract event 0x2ab93f65628379309f36cb125e90d7c902454a545c4f8b8cb0794af75c24b807.
//
// Solidity: event CompSpeedUpdated(address indexed cToken, uint256 newSpeed)
func (_Comptroller *ComptrollerFilterer) ParseCompSpeedUpdated(log types.Log) (*ComptrollerCompSpeedUpdated, error) {
	event := new(ComptrollerCompSpeedUpdated)
	if err := _Comptroller.contract.UnpackLog(event, "CompSpeedUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerCreditLimitChangedIterator is returned from FilterCreditLimitChanged and is used to iterate over the raw logs and unpacked data for CreditLimitChanged events raised by the Comptroller contract.
type ComptrollerCreditLimitChangedIterator struct {
	Event *ComptrollerCreditLimitChanged // Event containing the contract specifics and raw log

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
func (it *ComptrollerCreditLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerCreditLimitChanged)
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
		it.Event = new(ComptrollerCreditLimitChanged)
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
func (it *ComptrollerCreditLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerCreditLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerCreditLimitChanged represents a CreditLimitChanged event raised by the Comptroller contract.
type ComptrollerCreditLimitChanged struct {
	Protocol    common.Address
	CreditLimit *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreditLimitChanged is a free log retrieval operation binding the contract event 0xacba5197b02f201109f99752a9adb58c2598809f6acb3c34b7f1445f0cbeee87.
//
// Solidity: event CreditLimitChanged(address protocol, uint256 creditLimit)
func (_Comptroller *ComptrollerFilterer) FilterCreditLimitChanged(opts *bind.FilterOpts) (*ComptrollerCreditLimitChangedIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "CreditLimitChanged")
	if err != nil {
		return nil, err
	}
	return &ComptrollerCreditLimitChangedIterator{contract: _Comptroller.contract, event: "CreditLimitChanged", logs: logs, sub: sub}, nil
}

// WatchCreditLimitChanged is a free log subscription operation binding the contract event 0xacba5197b02f201109f99752a9adb58c2598809f6acb3c34b7f1445f0cbeee87.
//
// Solidity: event CreditLimitChanged(address protocol, uint256 creditLimit)
func (_Comptroller *ComptrollerFilterer) WatchCreditLimitChanged(opts *bind.WatchOpts, sink chan<- *ComptrollerCreditLimitChanged) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "CreditLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerCreditLimitChanged)
				if err := _Comptroller.contract.UnpackLog(event, "CreditLimitChanged", log); err != nil {
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

// ParseCreditLimitChanged is a log parse operation binding the contract event 0xacba5197b02f201109f99752a9adb58c2598809f6acb3c34b7f1445f0cbeee87.
//
// Solidity: event CreditLimitChanged(address protocol, uint256 creditLimit)
func (_Comptroller *ComptrollerFilterer) ParseCreditLimitChanged(log types.Log) (*ComptrollerCreditLimitChanged, error) {
	event := new(ComptrollerCreditLimitChanged)
	if err := _Comptroller.contract.UnpackLog(event, "CreditLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerDistributedBorrowerCompIterator is returned from FilterDistributedBorrowerComp and is used to iterate over the raw logs and unpacked data for DistributedBorrowerComp events raised by the Comptroller contract.
type ComptrollerDistributedBorrowerCompIterator struct {
	Event *ComptrollerDistributedBorrowerComp // Event containing the contract specifics and raw log

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
func (it *ComptrollerDistributedBorrowerCompIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerDistributedBorrowerComp)
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
		it.Event = new(ComptrollerDistributedBorrowerComp)
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
func (it *ComptrollerDistributedBorrowerCompIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerDistributedBorrowerCompIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerDistributedBorrowerComp represents a DistributedBorrowerComp event raised by the Comptroller contract.
type ComptrollerDistributedBorrowerComp struct {
	CToken          common.Address
	Borrower        common.Address
	CompDelta       *big.Int
	CompBorrowIndex *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributedBorrowerComp is a free log retrieval operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_Comptroller *ComptrollerFilterer) FilterDistributedBorrowerComp(opts *bind.FilterOpts, cToken []common.Address, borrower []common.Address) (*ComptrollerDistributedBorrowerCompIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "DistributedBorrowerComp", cTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &ComptrollerDistributedBorrowerCompIterator{contract: _Comptroller.contract, event: "DistributedBorrowerComp", logs: logs, sub: sub}, nil
}

// WatchDistributedBorrowerComp is a free log subscription operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_Comptroller *ComptrollerFilterer) WatchDistributedBorrowerComp(opts *bind.WatchOpts, sink chan<- *ComptrollerDistributedBorrowerComp, cToken []common.Address, borrower []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "DistributedBorrowerComp", cTokenRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerDistributedBorrowerComp)
				if err := _Comptroller.contract.UnpackLog(event, "DistributedBorrowerComp", log); err != nil {
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

// ParseDistributedBorrowerComp is a log parse operation binding the contract event 0x1fc3ecc087d8d2d15e23d0032af5a47059c3892d003d8e139fdcb6bb327c99a6.
//
// Solidity: event DistributedBorrowerComp(address indexed cToken, address indexed borrower, uint256 compDelta, uint256 compBorrowIndex)
func (_Comptroller *ComptrollerFilterer) ParseDistributedBorrowerComp(log types.Log) (*ComptrollerDistributedBorrowerComp, error) {
	event := new(ComptrollerDistributedBorrowerComp)
	if err := _Comptroller.contract.UnpackLog(event, "DistributedBorrowerComp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerDistributedSupplierCompIterator is returned from FilterDistributedSupplierComp and is used to iterate over the raw logs and unpacked data for DistributedSupplierComp events raised by the Comptroller contract.
type ComptrollerDistributedSupplierCompIterator struct {
	Event *ComptrollerDistributedSupplierComp // Event containing the contract specifics and raw log

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
func (it *ComptrollerDistributedSupplierCompIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerDistributedSupplierComp)
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
		it.Event = new(ComptrollerDistributedSupplierComp)
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
func (it *ComptrollerDistributedSupplierCompIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerDistributedSupplierCompIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerDistributedSupplierComp represents a DistributedSupplierComp event raised by the Comptroller contract.
type ComptrollerDistributedSupplierComp struct {
	CToken          common.Address
	Supplier        common.Address
	CompDelta       *big.Int
	CompSupplyIndex *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterDistributedSupplierComp is a free log retrieval operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_Comptroller *ComptrollerFilterer) FilterDistributedSupplierComp(opts *bind.FilterOpts, cToken []common.Address, supplier []common.Address) (*ComptrollerDistributedSupplierCompIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "DistributedSupplierComp", cTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return &ComptrollerDistributedSupplierCompIterator{contract: _Comptroller.contract, event: "DistributedSupplierComp", logs: logs, sub: sub}, nil
}

// WatchDistributedSupplierComp is a free log subscription operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_Comptroller *ComptrollerFilterer) WatchDistributedSupplierComp(opts *bind.WatchOpts, sink chan<- *ComptrollerDistributedSupplierComp, cToken []common.Address, supplier []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}
	var supplierRule []interface{}
	for _, supplierItem := range supplier {
		supplierRule = append(supplierRule, supplierItem)
	}

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "DistributedSupplierComp", cTokenRule, supplierRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerDistributedSupplierComp)
				if err := _Comptroller.contract.UnpackLog(event, "DistributedSupplierComp", log); err != nil {
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

// ParseDistributedSupplierComp is a log parse operation binding the contract event 0x2caecd17d02f56fa897705dcc740da2d237c373f70686f4e0d9bd3bf0400ea7a.
//
// Solidity: event DistributedSupplierComp(address indexed cToken, address indexed supplier, uint256 compDelta, uint256 compSupplyIndex)
func (_Comptroller *ComptrollerFilterer) ParseDistributedSupplierComp(log types.Log) (*ComptrollerDistributedSupplierComp, error) {
	event := new(ComptrollerDistributedSupplierComp)
	if err := _Comptroller.contract.UnpackLog(event, "DistributedSupplierComp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the Comptroller contract.
type ComptrollerFailureIterator struct {
	Event *ComptrollerFailure // Event containing the contract specifics and raw log

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
func (it *ComptrollerFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerFailure)
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
		it.Event = new(ComptrollerFailure)
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
func (it *ComptrollerFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerFailure represents a Failure event raised by the Comptroller contract.
type ComptrollerFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Comptroller *ComptrollerFilterer) FilterFailure(opts *bind.FilterOpts) (*ComptrollerFailureIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &ComptrollerFailureIterator{contract: _Comptroller.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_Comptroller *ComptrollerFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *ComptrollerFailure) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerFailure)
				if err := _Comptroller.contract.UnpackLog(event, "Failure", log); err != nil {
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
func (_Comptroller *ComptrollerFilterer) ParseFailure(log types.Log) (*ComptrollerFailure, error) {
	event := new(ComptrollerFailure)
	if err := _Comptroller.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerMarketEnteredIterator is returned from FilterMarketEntered and is used to iterate over the raw logs and unpacked data for MarketEntered events raised by the Comptroller contract.
type ComptrollerMarketEnteredIterator struct {
	Event *ComptrollerMarketEntered // Event containing the contract specifics and raw log

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
func (it *ComptrollerMarketEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerMarketEntered)
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
		it.Event = new(ComptrollerMarketEntered)
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
func (it *ComptrollerMarketEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerMarketEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerMarketEntered represents a MarketEntered event raised by the Comptroller contract.
type ComptrollerMarketEntered struct {
	CToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketEntered is a free log retrieval operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Comptroller *ComptrollerFilterer) FilterMarketEntered(opts *bind.FilterOpts) (*ComptrollerMarketEnteredIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return &ComptrollerMarketEnteredIterator{contract: _Comptroller.contract, event: "MarketEntered", logs: logs, sub: sub}, nil
}

// WatchMarketEntered is a free log subscription operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Comptroller *ComptrollerFilterer) WatchMarketEntered(opts *bind.WatchOpts, sink chan<- *ComptrollerMarketEntered) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "MarketEntered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerMarketEntered)
				if err := _Comptroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
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

// ParseMarketEntered is a log parse operation binding the contract event 0x3ab23ab0d51cccc0c3085aec51f99228625aa1a922b3a8ca89a26b0f2027a1a5.
//
// Solidity: event MarketEntered(address cToken, address account)
func (_Comptroller *ComptrollerFilterer) ParseMarketEntered(log types.Log) (*ComptrollerMarketEntered, error) {
	event := new(ComptrollerMarketEntered)
	if err := _Comptroller.contract.UnpackLog(event, "MarketEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerMarketExitedIterator is returned from FilterMarketExited and is used to iterate over the raw logs and unpacked data for MarketExited events raised by the Comptroller contract.
type ComptrollerMarketExitedIterator struct {
	Event *ComptrollerMarketExited // Event containing the contract specifics and raw log

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
func (it *ComptrollerMarketExitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerMarketExited)
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
		it.Event = new(ComptrollerMarketExited)
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
func (it *ComptrollerMarketExitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerMarketExitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerMarketExited represents a MarketExited event raised by the Comptroller contract.
type ComptrollerMarketExited struct {
	CToken  common.Address
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMarketExited is a free log retrieval operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Comptroller *ComptrollerFilterer) FilterMarketExited(opts *bind.FilterOpts) (*ComptrollerMarketExitedIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return &ComptrollerMarketExitedIterator{contract: _Comptroller.contract, event: "MarketExited", logs: logs, sub: sub}, nil
}

// WatchMarketExited is a free log subscription operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Comptroller *ComptrollerFilterer) WatchMarketExited(opts *bind.WatchOpts, sink chan<- *ComptrollerMarketExited) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "MarketExited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerMarketExited)
				if err := _Comptroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
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

// ParseMarketExited is a log parse operation binding the contract event 0xe699a64c18b07ac5b7301aa273f36a2287239eb9501d81950672794afba29a0d.
//
// Solidity: event MarketExited(address cToken, address account)
func (_Comptroller *ComptrollerFilterer) ParseMarketExited(log types.Log) (*ComptrollerMarketExited, error) {
	event := new(ComptrollerMarketExited)
	if err := _Comptroller.contract.UnpackLog(event, "MarketExited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerMarketListedIterator is returned from FilterMarketListed and is used to iterate over the raw logs and unpacked data for MarketListed events raised by the Comptroller contract.
type ComptrollerMarketListedIterator struct {
	Event *ComptrollerMarketListed // Event containing the contract specifics and raw log

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
func (it *ComptrollerMarketListedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerMarketListed)
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
		it.Event = new(ComptrollerMarketListed)
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
func (it *ComptrollerMarketListedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerMarketListedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerMarketListed represents a MarketListed event raised by the Comptroller contract.
type ComptrollerMarketListed struct {
	CToken common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMarketListed is a free log retrieval operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Comptroller *ComptrollerFilterer) FilterMarketListed(opts *bind.FilterOpts) (*ComptrollerMarketListedIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return &ComptrollerMarketListedIterator{contract: _Comptroller.contract, event: "MarketListed", logs: logs, sub: sub}, nil
}

// WatchMarketListed is a free log subscription operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Comptroller *ComptrollerFilterer) WatchMarketListed(opts *bind.WatchOpts, sink chan<- *ComptrollerMarketListed) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "MarketListed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerMarketListed)
				if err := _Comptroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
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

// ParseMarketListed is a log parse operation binding the contract event 0xcf583bb0c569eb967f806b11601c4cb93c10310485c67add5f8362c2f212321f.
//
// Solidity: event MarketListed(address cToken)
func (_Comptroller *ComptrollerFilterer) ParseMarketListed(log types.Log) (*ComptrollerMarketListed, error) {
	event := new(ComptrollerMarketListed)
	if err := _Comptroller.contract.UnpackLog(event, "MarketListed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewBorrowCapIterator is returned from FilterNewBorrowCap and is used to iterate over the raw logs and unpacked data for NewBorrowCap events raised by the Comptroller contract.
type ComptrollerNewBorrowCapIterator struct {
	Event *ComptrollerNewBorrowCap // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewBorrowCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewBorrowCap)
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
		it.Event = new(ComptrollerNewBorrowCap)
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
func (it *ComptrollerNewBorrowCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewBorrowCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewBorrowCap represents a NewBorrowCap event raised by the Comptroller contract.
type ComptrollerNewBorrowCap struct {
	CToken       common.Address
	NewBorrowCap *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCap is a free log retrieval operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed cToken, uint256 newBorrowCap)
func (_Comptroller *ComptrollerFilterer) FilterNewBorrowCap(opts *bind.FilterOpts, cToken []common.Address) (*ComptrollerNewBorrowCapIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewBorrowCap", cTokenRule)
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewBorrowCapIterator{contract: _Comptroller.contract, event: "NewBorrowCap", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCap is a free log subscription operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed cToken, uint256 newBorrowCap)
func (_Comptroller *ComptrollerFilterer) WatchNewBorrowCap(opts *bind.WatchOpts, sink chan<- *ComptrollerNewBorrowCap, cToken []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewBorrowCap", cTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewBorrowCap)
				if err := _Comptroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
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

// ParseNewBorrowCap is a log parse operation binding the contract event 0x6f1951b2aad10f3fc81b86d91105b413a5b3f847a34bbc5ce1904201b14438f6.
//
// Solidity: event NewBorrowCap(address indexed cToken, uint256 newBorrowCap)
func (_Comptroller *ComptrollerFilterer) ParseNewBorrowCap(log types.Log) (*ComptrollerNewBorrowCap, error) {
	event := new(ComptrollerNewBorrowCap)
	if err := _Comptroller.contract.UnpackLog(event, "NewBorrowCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewBorrowCapGuardianIterator is returned from FilterNewBorrowCapGuardian and is used to iterate over the raw logs and unpacked data for NewBorrowCapGuardian events raised by the Comptroller contract.
type ComptrollerNewBorrowCapGuardianIterator struct {
	Event *ComptrollerNewBorrowCapGuardian // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewBorrowCapGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewBorrowCapGuardian)
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
		it.Event = new(ComptrollerNewBorrowCapGuardian)
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
func (it *ComptrollerNewBorrowCapGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewBorrowCapGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewBorrowCapGuardian represents a NewBorrowCapGuardian event raised by the Comptroller contract.
type ComptrollerNewBorrowCapGuardian struct {
	OldBorrowCapGuardian common.Address
	NewBorrowCapGuardian common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewBorrowCapGuardian is a free log retrieval operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Comptroller *ComptrollerFilterer) FilterNewBorrowCapGuardian(opts *bind.FilterOpts) (*ComptrollerNewBorrowCapGuardianIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewBorrowCapGuardianIterator{contract: _Comptroller.contract, event: "NewBorrowCapGuardian", logs: logs, sub: sub}, nil
}

// WatchNewBorrowCapGuardian is a free log subscription operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Comptroller *ComptrollerFilterer) WatchNewBorrowCapGuardian(opts *bind.WatchOpts, sink chan<- *ComptrollerNewBorrowCapGuardian) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewBorrowCapGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewBorrowCapGuardian)
				if err := _Comptroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
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

// ParseNewBorrowCapGuardian is a log parse operation binding the contract event 0xeda98690e518e9a05f8ec6837663e188211b2da8f4906648b323f2c1d4434e29.
//
// Solidity: event NewBorrowCapGuardian(address oldBorrowCapGuardian, address newBorrowCapGuardian)
func (_Comptroller *ComptrollerFilterer) ParseNewBorrowCapGuardian(log types.Log) (*ComptrollerNewBorrowCapGuardian, error) {
	event := new(ComptrollerNewBorrowCapGuardian)
	if err := _Comptroller.contract.UnpackLog(event, "NewBorrowCapGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewCloseFactorIterator is returned from FilterNewCloseFactor and is used to iterate over the raw logs and unpacked data for NewCloseFactor events raised by the Comptroller contract.
type ComptrollerNewCloseFactorIterator struct {
	Event *ComptrollerNewCloseFactor // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewCloseFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewCloseFactor)
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
		it.Event = new(ComptrollerNewCloseFactor)
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
func (it *ComptrollerNewCloseFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewCloseFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewCloseFactor represents a NewCloseFactor event raised by the Comptroller contract.
type ComptrollerNewCloseFactor struct {
	OldCloseFactorMantissa *big.Int
	NewCloseFactorMantissa *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterNewCloseFactor is a free log retrieval operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Comptroller *ComptrollerFilterer) FilterNewCloseFactor(opts *bind.FilterOpts) (*ComptrollerNewCloseFactorIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewCloseFactorIterator{contract: _Comptroller.contract, event: "NewCloseFactor", logs: logs, sub: sub}, nil
}

// WatchNewCloseFactor is a free log subscription operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Comptroller *ComptrollerFilterer) WatchNewCloseFactor(opts *bind.WatchOpts, sink chan<- *ComptrollerNewCloseFactor) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewCloseFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewCloseFactor)
				if err := _Comptroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
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

// ParseNewCloseFactor is a log parse operation binding the contract event 0x3b9670cf975d26958e754b57098eaa2ac914d8d2a31b83257997b9f346110fd9.
//
// Solidity: event NewCloseFactor(uint256 oldCloseFactorMantissa, uint256 newCloseFactorMantissa)
func (_Comptroller *ComptrollerFilterer) ParseNewCloseFactor(log types.Log) (*ComptrollerNewCloseFactor, error) {
	event := new(ComptrollerNewCloseFactor)
	if err := _Comptroller.contract.UnpackLog(event, "NewCloseFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewCollateralFactorIterator is returned from FilterNewCollateralFactor and is used to iterate over the raw logs and unpacked data for NewCollateralFactor events raised by the Comptroller contract.
type ComptrollerNewCollateralFactorIterator struct {
	Event *ComptrollerNewCollateralFactor // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewCollateralFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewCollateralFactor)
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
		it.Event = new(ComptrollerNewCollateralFactor)
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
func (it *ComptrollerNewCollateralFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewCollateralFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewCollateralFactor represents a NewCollateralFactor event raised by the Comptroller contract.
type ComptrollerNewCollateralFactor struct {
	CToken                      common.Address
	OldCollateralFactorMantissa *big.Int
	NewCollateralFactorMantissa *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterNewCollateralFactor is a free log retrieval operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Comptroller *ComptrollerFilterer) FilterNewCollateralFactor(opts *bind.FilterOpts) (*ComptrollerNewCollateralFactorIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewCollateralFactorIterator{contract: _Comptroller.contract, event: "NewCollateralFactor", logs: logs, sub: sub}, nil
}

// WatchNewCollateralFactor is a free log subscription operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Comptroller *ComptrollerFilterer) WatchNewCollateralFactor(opts *bind.WatchOpts, sink chan<- *ComptrollerNewCollateralFactor) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewCollateralFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewCollateralFactor)
				if err := _Comptroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
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

// ParseNewCollateralFactor is a log parse operation binding the contract event 0x70483e6592cd5182d45ac970e05bc62cdcc90e9d8ef2c2dbe686cf383bcd7fc5.
//
// Solidity: event NewCollateralFactor(address cToken, uint256 oldCollateralFactorMantissa, uint256 newCollateralFactorMantissa)
func (_Comptroller *ComptrollerFilterer) ParseNewCollateralFactor(log types.Log) (*ComptrollerNewCollateralFactor, error) {
	event := new(ComptrollerNewCollateralFactor)
	if err := _Comptroller.contract.UnpackLog(event, "NewCollateralFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewLiquidationIncentiveIterator is returned from FilterNewLiquidationIncentive and is used to iterate over the raw logs and unpacked data for NewLiquidationIncentive events raised by the Comptroller contract.
type ComptrollerNewLiquidationIncentiveIterator struct {
	Event *ComptrollerNewLiquidationIncentive // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewLiquidationIncentiveIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewLiquidationIncentive)
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
		it.Event = new(ComptrollerNewLiquidationIncentive)
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
func (it *ComptrollerNewLiquidationIncentiveIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewLiquidationIncentiveIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewLiquidationIncentive represents a NewLiquidationIncentive event raised by the Comptroller contract.
type ComptrollerNewLiquidationIncentive struct {
	OldLiquidationIncentiveMantissa *big.Int
	NewLiquidationIncentiveMantissa *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterNewLiquidationIncentive is a free log retrieval operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Comptroller *ComptrollerFilterer) FilterNewLiquidationIncentive(opts *bind.FilterOpts) (*ComptrollerNewLiquidationIncentiveIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewLiquidationIncentiveIterator{contract: _Comptroller.contract, event: "NewLiquidationIncentive", logs: logs, sub: sub}, nil
}

// WatchNewLiquidationIncentive is a free log subscription operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Comptroller *ComptrollerFilterer) WatchNewLiquidationIncentive(opts *bind.WatchOpts, sink chan<- *ComptrollerNewLiquidationIncentive) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewLiquidationIncentive")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewLiquidationIncentive)
				if err := _Comptroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
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

// ParseNewLiquidationIncentive is a log parse operation binding the contract event 0xaeba5a6c40a8ac138134bff1aaa65debf25971188a58804bad717f82f0ec1316.
//
// Solidity: event NewLiquidationIncentive(uint256 oldLiquidationIncentiveMantissa, uint256 newLiquidationIncentiveMantissa)
func (_Comptroller *ComptrollerFilterer) ParseNewLiquidationIncentive(log types.Log) (*ComptrollerNewLiquidationIncentive, error) {
	event := new(ComptrollerNewLiquidationIncentive)
	if err := _Comptroller.contract.UnpackLog(event, "NewLiquidationIncentive", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewPauseGuardianIterator is returned from FilterNewPauseGuardian and is used to iterate over the raw logs and unpacked data for NewPauseGuardian events raised by the Comptroller contract.
type ComptrollerNewPauseGuardianIterator struct {
	Event *ComptrollerNewPauseGuardian // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewPauseGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewPauseGuardian)
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
		it.Event = new(ComptrollerNewPauseGuardian)
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
func (it *ComptrollerNewPauseGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewPauseGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewPauseGuardian represents a NewPauseGuardian event raised by the Comptroller contract.
type ComptrollerNewPauseGuardian struct {
	OldPauseGuardian common.Address
	NewPauseGuardian common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewPauseGuardian is a free log retrieval operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Comptroller *ComptrollerFilterer) FilterNewPauseGuardian(opts *bind.FilterOpts) (*ComptrollerNewPauseGuardianIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewPauseGuardianIterator{contract: _Comptroller.contract, event: "NewPauseGuardian", logs: logs, sub: sub}, nil
}

// WatchNewPauseGuardian is a free log subscription operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Comptroller *ComptrollerFilterer) WatchNewPauseGuardian(opts *bind.WatchOpts, sink chan<- *ComptrollerNewPauseGuardian) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewPauseGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewPauseGuardian)
				if err := _Comptroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
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

// ParseNewPauseGuardian is a log parse operation binding the contract event 0x0613b6ee6a04f0d09f390e4d9318894b9f6ac7fd83897cd8d18896ba579c401e.
//
// Solidity: event NewPauseGuardian(address oldPauseGuardian, address newPauseGuardian)
func (_Comptroller *ComptrollerFilterer) ParseNewPauseGuardian(log types.Log) (*ComptrollerNewPauseGuardian, error) {
	event := new(ComptrollerNewPauseGuardian)
	if err := _Comptroller.contract.UnpackLog(event, "NewPauseGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewPriceOracleIterator is returned from FilterNewPriceOracle and is used to iterate over the raw logs and unpacked data for NewPriceOracle events raised by the Comptroller contract.
type ComptrollerNewPriceOracleIterator struct {
	Event *ComptrollerNewPriceOracle // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewPriceOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewPriceOracle)
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
		it.Event = new(ComptrollerNewPriceOracle)
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
func (it *ComptrollerNewPriceOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewPriceOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewPriceOracle represents a NewPriceOracle event raised by the Comptroller contract.
type ComptrollerNewPriceOracle struct {
	OldPriceOracle common.Address
	NewPriceOracle common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewPriceOracle is a free log retrieval operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Comptroller *ComptrollerFilterer) FilterNewPriceOracle(opts *bind.FilterOpts) (*ComptrollerNewPriceOracleIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewPriceOracleIterator{contract: _Comptroller.contract, event: "NewPriceOracle", logs: logs, sub: sub}, nil
}

// WatchNewPriceOracle is a free log subscription operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Comptroller *ComptrollerFilterer) WatchNewPriceOracle(opts *bind.WatchOpts, sink chan<- *ComptrollerNewPriceOracle) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewPriceOracle")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewPriceOracle)
				if err := _Comptroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
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

// ParseNewPriceOracle is a log parse operation binding the contract event 0xd52b2b9b7e9ee655fcb95d2e5b9e0c9f69e7ef2b8e9d2d0ea78402d576d22e22.
//
// Solidity: event NewPriceOracle(address oldPriceOracle, address newPriceOracle)
func (_Comptroller *ComptrollerFilterer) ParseNewPriceOracle(log types.Log) (*ComptrollerNewPriceOracle, error) {
	event := new(ComptrollerNewPriceOracle)
	if err := _Comptroller.contract.UnpackLog(event, "NewPriceOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewSupplyCapIterator is returned from FilterNewSupplyCap and is used to iterate over the raw logs and unpacked data for NewSupplyCap events raised by the Comptroller contract.
type ComptrollerNewSupplyCapIterator struct {
	Event *ComptrollerNewSupplyCap // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewSupplyCapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewSupplyCap)
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
		it.Event = new(ComptrollerNewSupplyCap)
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
func (it *ComptrollerNewSupplyCapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewSupplyCapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewSupplyCap represents a NewSupplyCap event raised by the Comptroller contract.
type ComptrollerNewSupplyCap struct {
	CToken       common.Address
	NewSupplyCap *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewSupplyCap is a free log retrieval operation binding the contract event 0x9e0ad9cee10bdf36b7fbd38910c0bdff0f275ace679b45b922381c2723d676f8.
//
// Solidity: event NewSupplyCap(address indexed cToken, uint256 newSupplyCap)
func (_Comptroller *ComptrollerFilterer) FilterNewSupplyCap(opts *bind.FilterOpts, cToken []common.Address) (*ComptrollerNewSupplyCapIterator, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewSupplyCap", cTokenRule)
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewSupplyCapIterator{contract: _Comptroller.contract, event: "NewSupplyCap", logs: logs, sub: sub}, nil
}

// WatchNewSupplyCap is a free log subscription operation binding the contract event 0x9e0ad9cee10bdf36b7fbd38910c0bdff0f275ace679b45b922381c2723d676f8.
//
// Solidity: event NewSupplyCap(address indexed cToken, uint256 newSupplyCap)
func (_Comptroller *ComptrollerFilterer) WatchNewSupplyCap(opts *bind.WatchOpts, sink chan<- *ComptrollerNewSupplyCap, cToken []common.Address) (event.Subscription, error) {

	var cTokenRule []interface{}
	for _, cTokenItem := range cToken {
		cTokenRule = append(cTokenRule, cTokenItem)
	}

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewSupplyCap", cTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewSupplyCap)
				if err := _Comptroller.contract.UnpackLog(event, "NewSupplyCap", log); err != nil {
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

// ParseNewSupplyCap is a log parse operation binding the contract event 0x9e0ad9cee10bdf36b7fbd38910c0bdff0f275ace679b45b922381c2723d676f8.
//
// Solidity: event NewSupplyCap(address indexed cToken, uint256 newSupplyCap)
func (_Comptroller *ComptrollerFilterer) ParseNewSupplyCap(log types.Log) (*ComptrollerNewSupplyCap, error) {
	event := new(ComptrollerNewSupplyCap)
	if err := _Comptroller.contract.UnpackLog(event, "NewSupplyCap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerNewSupplyCapGuardianIterator is returned from FilterNewSupplyCapGuardian and is used to iterate over the raw logs and unpacked data for NewSupplyCapGuardian events raised by the Comptroller contract.
type ComptrollerNewSupplyCapGuardianIterator struct {
	Event *ComptrollerNewSupplyCapGuardian // Event containing the contract specifics and raw log

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
func (it *ComptrollerNewSupplyCapGuardianIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerNewSupplyCapGuardian)
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
		it.Event = new(ComptrollerNewSupplyCapGuardian)
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
func (it *ComptrollerNewSupplyCapGuardianIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerNewSupplyCapGuardianIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerNewSupplyCapGuardian represents a NewSupplyCapGuardian event raised by the Comptroller contract.
type ComptrollerNewSupplyCapGuardian struct {
	OldSupplyCapGuardian common.Address
	NewSupplyCapGuardian common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewSupplyCapGuardian is a free log retrieval operation binding the contract event 0xb0d3622c24ac9bd967d8f37a25808b3e668fe7ed4f3075bbe82842d3e287c044.
//
// Solidity: event NewSupplyCapGuardian(address oldSupplyCapGuardian, address newSupplyCapGuardian)
func (_Comptroller *ComptrollerFilterer) FilterNewSupplyCapGuardian(opts *bind.FilterOpts) (*ComptrollerNewSupplyCapGuardianIterator, error) {

	logs, sub, err := _Comptroller.contract.FilterLogs(opts, "NewSupplyCapGuardian")
	if err != nil {
		return nil, err
	}
	return &ComptrollerNewSupplyCapGuardianIterator{contract: _Comptroller.contract, event: "NewSupplyCapGuardian", logs: logs, sub: sub}, nil
}

// WatchNewSupplyCapGuardian is a free log subscription operation binding the contract event 0xb0d3622c24ac9bd967d8f37a25808b3e668fe7ed4f3075bbe82842d3e287c044.
//
// Solidity: event NewSupplyCapGuardian(address oldSupplyCapGuardian, address newSupplyCapGuardian)
func (_Comptroller *ComptrollerFilterer) WatchNewSupplyCapGuardian(opts *bind.WatchOpts, sink chan<- *ComptrollerNewSupplyCapGuardian) (event.Subscription, error) {

	logs, sub, err := _Comptroller.contract.WatchLogs(opts, "NewSupplyCapGuardian")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerNewSupplyCapGuardian)
				if err := _Comptroller.contract.UnpackLog(event, "NewSupplyCapGuardian", log); err != nil {
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

// ParseNewSupplyCapGuardian is a log parse operation binding the contract event 0xb0d3622c24ac9bd967d8f37a25808b3e668fe7ed4f3075bbe82842d3e287c044.
//
// Solidity: event NewSupplyCapGuardian(address oldSupplyCapGuardian, address newSupplyCapGuardian)
func (_Comptroller *ComptrollerFilterer) ParseNewSupplyCapGuardian(log types.Log) (*ComptrollerNewSupplyCapGuardian, error) {
	event := new(ComptrollerNewSupplyCapGuardian)
	if err := _Comptroller.contract.UnpackLog(event, "NewSupplyCapGuardian", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
