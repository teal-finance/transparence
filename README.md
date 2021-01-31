# Transparence
![](images/logo.png)
*Transparence, French word for transparency.*

This is a collection of tools, build during the [ETHOnline Hackathon](https://ethonline.org/) to improve transparency in cryptocurrencies.

# Pegged tokens crosschain solvability audit
The first feature enable to verify the solvency of pegged-tokens accross respective blockchains such as BTC pegged tokens (tBTC, wBTC) or Binance Pegged tokens. The tool query both blockchains and verify the solvency of tokens by comparing the token supply on a chain and the reserve address balance on the other. Transparence is compatible with several blockchains : Bitcoin, Ethereum, BinanceChain, BinanceSmartChain.

#### Proof of Audit on Tellor's Orace platform 
By publishing audit results to Tellor's smart contracts, we are making available to everyone a Proof of Audit enabling to assess the solvency of a token without having to rely on a third-party such as explorers or having to execute audit by themself. In addition, once uploaded to Tellor, the audit results might also be used as conditions to smart contracts execution etc.


# Compound and CreamFinance protocol auditability features
The second feature is more defi-oriented and checks the metrics of Compound-fork protocols.
- Access to different metrics (TotalBorrow etc)
- Checking the solvency of underlying cAssets

Many other interesting metrics, for both tokens and platforms, can be deduced from this data, such as the percentage of tokens locked across different defi platforms and so on.


# Technical details
*In this part we will describe briefly the different technology used and how we are communicating with several blockchains.*

- Bitcoin:
    -  Data source : Local Node
- Ethereum:
    - Data source : Infura Mainnet,Rinkeby
    - ERC20 smart contract ABI
    - go-ethereum
- Binance Smart Chain:
    - Data source : Blockchain Node, Binance Chain API
- Tellor:
    - Data source : 
        - Infura Rinkeby : to update/get new data feed from TellorPlayground
        - Infura Mainnet : to get the price of btc or other assets from UsingTellor
    - TellorPlayground & UsingTellor smart contracts ABIs
- Compound & CreamFinance:
    - Data source : Infura Mainnet / Binance Blockchain Node
    - cToken/crToken & Comptroller smart contract ABIs
- Coingecko :
    - API to get informations about cryptocurrencies (not used in audits)

# References:
- [Compound Protocol Dos](https://compound.finance/docs)
- [CreamFinance Docs](https://docs.cream.finance/)
- [Binance Proof of Asset](https://www.binance.org/en/assets-proof)
- [Tellor Docs](https://docs.tellor.io/tellor/)
- [Go-ethereum](https://github.com/ethereum/go-ethereum)
- [Defi-Pulse-Adapters](https://github.com/ConcourseOpen/DeFi-Pulse-Adapters)