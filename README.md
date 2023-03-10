# XELIS-INDEX

The indexer helps make the decentralized data across multiple nodes more accessible for Xelis users.

When a stable block or transaction appear on the blockchain, they are recorded into
a centralized Postgres database.

The indexed data can be easily search for and
retrieve based on very specific parameters.

Developers can quickly build and test their application,
without having to manually query the entire blockchain.

## Use service

`Popular trusted index`

Mainnet

| Endpoint                 | Maintainer          |
| ------------------------ | ------------------- |
| <https://index.xelis.io> | `g45t345rt` `Slixe` |

Testnet

| Endpoint                         | Maintainer          |
| -------------------------------- | ------------------- |
| <https://testnet-index.xelis.io> | `g45t345rt` `Slixe` |

Dev

| Endpoint                     | Maintainer          |
| ---------------------------- | ------------------- |
| <https://dev-index.xelis.io> | `g45t345rt` `Slixe` |

## Run your own

If you don't want to have limited ressource you should run your own indexer and modify code based on your needs.

## Build

Pull xgo docker first

`docker pull crazymax/xgo:latest`

and quick build for linux

`./xgo_build.sh`

or target build directly

`xgo -v -targets linux/amd64 -dest build .`
