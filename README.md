# Transparence
![](images/logo.png)
*Transparence, French word for transparency.*

This is a collection of tools, build during the [ETHOnline Hackathon](https://ethonline.org/) to improve transparency in cryptocurrencies.


## Summary 
- [Pegged tokens crosschain solvability audit](https://github.com/cryptohazard/transparence/tree/refactoring#pegged-tokens-crosschain-solvability-audit)
- [Proof of Audit on Tellor's Orace platform](https://github.com/cryptohazard/transparence/tree/refactoring#proof-of-audit-on-tellors-orace-platform)
- [Compound and CreamFinance protocol auditability features](https://github.com/cryptohazard/transparence/tree/refactoring#compound-and-creamfinance-protocol-auditability-features)
- [Technical details](https://github.com/cryptohazard/transparence/tree/refactoring#technical-details)
- [Further works](https://github.com/cryptohazard/transparence/tree/refactoring#further-works)
- [References](https://github.com/cryptohazard/transparence/tree/refactoring#references)

### Pegged tokens crosschain solvability audit
The first feature enable to verify the solvency of pegged-tokens accross respective blockchains such as BTC pegged tokens (tBTC, wBTC) or Binance Pegged tokens. The tool query both blockchains and verify the solvency of tokens by comparing the token supply on a chain and the reserve address balance on the other. Transparence is compatible with several blockchains : Bitcoin, Ethereum, BinanceChain, BinanceSmartChain.

#### Proof of Audit on Tellor's Orace platform 
By publishing audit results to Tellor's smart contracts, we are making available to everyone a Proof of Audit enabling to assess the solvency of a token without having to rely on a third-party (such as explorers or interfaces provided by token creators [1](https://wbtc.network/dashboard/audit) [2](https://www.binance.org/en/assets-proof)) or even having to execute audit by themself. In addition, once uploaded to Tellor, the audit results might also be used as conditions to smart contracts execution to protect users funds etc. However, in the case of a "centralized" pegged-token (ie issuance not governed by a smart contract protocol) this "Proof of Audit" cannot at all represent a full audit process because Transparence only checks the balance of addresses and is not able to verify the possession of an address by a claiming entity.


### Compound and CreamFinance protocol auditability features
The second feature is more defi-oriented and checks the metrics of Compound-fork protocols.
- Access to different metrics (TotalBorrow etc)

Many other interesting metrics, for both tokens and platforms, can be deduced from this data, such as the percentage of tokens locked across different defi platforms/blockchains and so on.


### Technical details
*In this part we will describe briefly the different datasources and processes used to communicate with several blockchains.*

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

### Command Line interface
There is 4 arguments that can be used for Transparence's CLI, "blockchain", "protocol", "symbol" ,"tellor".
TODO: add args description
Here are some examples:
- go run ./cmd/cli/main.go --blockchain binance --protocol pegged-tokens 
- go run ./cmd/cli/main.go --blockchain binance --protocol pegged-tokens --symbol UNI
- go run ./cmd/cli/main.go --blockchain binance --protocol pegged-tokens --symbol UNI --tellor
- go run ./cmd/cli/main.go --blockchain ethereum --protocol compound
- go run ./cmd/cli/main.go --blockchain binance --protocol cream
- go run ./cmd/cli/main.go --blockchain bitcoin --protocol pegged-tokens

### Further works
Transparence will continue to add different blockchains and protocols with a focus on the transparency and "health check" of cryptocurrencies. We will also continue to extract some metrics from our data in order to better represent the risks/state of defi.

### References:
- [Compound Protocol Docs](https://compound.finance/docs)
- [CreamFinance Docs](https://docs.cream.finance/)
- [Binance Proof of Asset](https://www.binance.org/en/assets-proof)
- [Tellor Docs](https://docs.tellor.io/tellor/)
- [Go-ethereum](https://github.com/ethereum/go-ethereum)
- [Defi-Pulse-Adapters](https://github.com/ConcourseOpen/DeFi-Pulse-Adapters)
