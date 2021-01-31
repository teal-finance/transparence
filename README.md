# transparence
![](images/logo.png)
*Transparence, French word for transparency.*

This is a collection of tools, build during the [ETHOnline Hackathon](https://ethonline.org/) to improve transparency in cryptocurrencies.

# Pegged tokens crosschain solvability audit
The first feature enable to verify the solvency of Ethereum ERC20 backed tokens present in other blockchains such as BTC pegged tokens (tBTC, wBTC) or Binance Pegged tokens. The tool uses the smart contracts ABIs of ERC20/BEP20 to query the blockchain nodes and verify the solvency of tokens by comparing the token address and the reserve address amounts. Transparence is compatible with several blockchains : Bitcoin, Ethereum, BinanceChain, BinanceSmartChain.

#### Tellor's Orace platform
By publishing audit results to Tellor's smart contracts, we are making available to everyone a Proof of Audit enabling to assess the solvency of a token without having to rely on a third-party or having to execute audit by themself. In addition, once uploaded to Tellor the audit results might also be used as conditions to smart contracts execution etc.


# Compound and CreamFinance protocol auditability features
The second feature is more defi-oriented and checks the metrics of Compound-fork protocols.
- Access to different metrics (TotalBorrow etc)
- Checking the solvency of underlying cAssets


# Technical details
*In this part we will describe briefly the different technology used and how we are communicating with several blockchains.*
Bitcoin:
- Local Node

Ethereum:
- Infura Mainnet,Rinkeby
- ERC20 smart contract ABI
- go-ethereum

Binance Smart Chain:
- Blockchain Node
- Binance Chain API

Tellor:
- TellorPlayground (Rinkeby) to update/get new data feed
- UsingTellor (Mainnet) to get the price of btc or other assets

Compound:
- Comptroller smart contract ABI
- cToken smart contract ABI

Coingecko :
- API to get informations about cryptocurrencies (not used in audits)

# References:
- Compound
- Cream
- https://www.binance.org/en/assets-proof
- Tellor
- go eth
- defipulse github




