# eth-p2p-challenger

![GitHub Tag](https://img.shields.io/github/v/tag/kamikazechaser/eth-p2p-challenger)

> Ethereum execution client devp2p probe and challenger

## Description

A proof-of-concept P2P challenger that connects to Ethereum nodes and submits various random protocol challenges over RLPx to test node behavior.

The main goal is to probe node capabilities to verify execution client implementation authenticity.

### Challenges Implemented

- [x] **Random block header requests** - Random range and specific block queries

### Other potential challenges

- **Random trie nodes over Snap protocol**
- **Random block bodies**

## Refs

- https://github.com/ethereum/go-ethereum/
- https://github.com/ethereum/devp2p
- https://github.com/ethereumjs/ethereumjs-monorepo/tree/master/packages/devp2p
- https://github.com/ethpandaops/xatu

## License

[AGPL-3.0](LICENSE)
