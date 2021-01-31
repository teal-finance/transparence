# Transparence
![](images/logo.png)
*Transparence, French word for transparency.*

This started during the [ETHOnline Hackathon](https://ethonline.org/) to improve transparency in cryptocurrencies.

We improve on this work for the [Encode Club Hackathon](https://hack.encode.club/).

# Description

`Transparence`, is a collection of tools, to improve transparency in cryptocurrencies. We want to make it easier for users interacting with different protocols to extract data from the chain that gives them guarantees that everything is fine.

Specially with the trends from DeFi and Yield farming, users are taking more risks without properly assessing the assumptions and trade-off. For example, using protocols like Compound, users will add leverage to their positions without considering if they could unwind their positions if we had to get through black Thursday again. Projects with liquidation tend to show only information related to the individuals and not network health metrics, such as a price at which half of the positions would need to be liquidated.

Another example is cross chain projects like Keep Network, RenVM, WBTC, who want to bring Bitcoin on Ethereum or on Binance. It is important for users to verify that they are backed 1:1. The only way to have this insurance is to verify on both chains that the data are consistent.


The project ethos is to enable users to verify data on chain using their full node and not an API, when possible. We started with a Bitcoin full node, but for now, can only use Binance and Ethereum APIs. 
From there, we developed the functions to retrieve the data of interest and process it.

Our goals are to:
* verify pegged assets cross chain (Bitcoin on Ethereum/Binance/..., ERC20 on other chains) and raise alarm when there is a breach of trust
* identify pain prices from protocols with a liquidation, CDP, loans, insurance, options,  to avoid snow ball effect
* provide those health/risk metrics on chain for smart contracts

## Technical Summary 
- [Pegged tokens cross-chain solvability audit](https://github.com/cryptohazard/transparence/tree/refactoring#pegged-tokens-crosschain-solvability-audit)
- [Proof of Audit on Tellor's Oracle platform](https://github.com/cryptohazard/transparence/tree/refactoring#proof-of-audit-on-tellors-orace-platform)
- [Compound and Cream Finance protocol auditability features](https://github.com/cryptohazard/transparence/tree/refactoring#compound-and-creamfinance-protocol-auditability-features)
- [Technical details](https://github.com/cryptohazard/transparence/tree/refactoring#technical-details)
- [Further works](https://github.com/cryptohazard/transparence/tree/refactoring#further-works)
- [References](https://github.com/cryptohazard/transparence/tree/refactoring#references)

### Pegged tokens crosschain solvability audit
The first feature enable to verify the solvency of pegged-tokens accross respective blockchains such as BTC pegged tokens (tBTC, wBTC) or Binance Pegged tokens. The tool query both blockchains and verify the solvency of tokens by comparing the token supply on a chain and the reserve address balance on the other. Transparence is compatible with several blockchains : Bitcoin, Ethereum, BinanceChain, BinanceSmartChain. Some protocols are fairly straightforward, like wBTC, while other use more complex cryptography/trust (tBTC).

#### Proof of Audit on Tellor's Orace platform 
By publishing audit results to Tellor's smart contracts, we are making available to everyone a Proof of Audit enabling to assess the solvency of a token without having to rely on a third-party (such as explorers or interfaces provided by token creators [1](https://wbtc.network/dashboard/audit) [2](https://www.binance.org/en/assets-proof)) or even having to execute audit by themself. In addition, once uploaded to Tellor, the audit results might also be used as conditions to smart contracts execution to protect users funds etc. However, in the case of a "centralized" pegged-token (ie issuance not governed by a smart contract protocol) this "Proof of Audit" cannot at all represent a full audit process because Transparence only checks the balance of addresses and is not able to verify the possession of an address by a claiming entity.


### Compound and CreamFinance protocol auditability features
The second feature is more DeFi-oriented and checks the metrics of Compound-fork protocols.
- Access to different metrics (TotalBorrow etc)

Many other interesting metrics, for both tokens and platforms, can be deduced from this data, such as the percentage of tokens locked across different defi platforms/blockchains and so on.


### Technical details
*In this part we will describe briefly the different sources of data and processes used to communicate with several blockchains.*

- Bitcoin:
    -  Data source : Local Node
    -  Set labels on Bitcoin node for addresses of interest
    -  Retrieve addresses involved in a tx
    -  Pain point: retrieve the actual balance, which is hard with utxos (compare to balance model)
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

### Further works
Transparence will continue to add different blockchains and protocols with a focus on the transparency and "health check" of cryptocurrencies. We will also continue to extract some metrics from our data in order to better represent the risks/state of DeFi.

### References:
- [Compound Protocol Docs](https://compound.finance/docs)
- [CreamFinance Docs](https://docs.cream.finance/)
- [Binance Proof of Asset](https://www.binance.org/en/assets-proof)
- [Tellor Docs](https://docs.tellor.io/tellor/)
- [Go-ethereum](https://github.com/ethereum/go-ethereum)
- [Defi-Pulse-Adapters](https://github.com/ConcourseOpen/DeFi-Pulse-Adapters)
