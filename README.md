# eth-p2p-challenger

![GitHub Tag](https://img.shields.io/github/v/tag/kamikazechaser/eth-p2p-challenger)

> Ethereum execution client devp2p probe and challenger

## Description

A proof-of-concept P2P challenger that connects to Ethereum nodes and submits various random protocol challenges over RLPx to test node behavior.

The main goal is to probe node capabilities to verify execution client implementation authenticity.


```bash
time=2025-06-02T11:37:21.509+03:00 level=INFO msg="starting eth-p2p-challenger" build=dev
time=2025-06-02T11:37:32.027+03:00 level=INFO msg="sending random block headers request" start_block=35055942 reverse=true
time=2025-06-02T11:37:32.281+03:00 level=INFO msg="received block header" number=35055942 hash=0x7b76a67e39d2e93ccbae023efd9b8b47fd3c551db17298cbaf6515725bc5b4f5
time=2025-06-02T11:37:32.281+03:00 level=INFO msg="received block header" number=35055941 hash=0xb24e898c1ec473dd75693d1e85523934ed801bf137294bbbe45aafcb84bbb3ad
time=2025-06-02T11:37:32.281+03:00 level=INFO msg="received block header" number=35055940 hash=0x032473c23b2961be8c5b7e2c82e26feb6da44098899590fa1ee682b7a5346990
time=2025-06-02T11:37:32.281+03:00 level=INFO msg="received block header" number=35055939 hash=0x9465e109c235e9b8a0bc342e4d2fbaa71d435efb575b55db9eb89820a800d86a
time=2025-06-02T11:37:32.281+03:00 level=INFO msg="received block header" number=35055938 hash=0xd5fcc2b40ce474d7984151cc6d3ef6094ff120710ad6e54d1f0a2a36f9282f8f
time=2025-06-02T11:37:42.028+03:00 level=INFO msg="sending random block headers request" start_block=35013227 reverse=true
time=2025-06-02T11:37:42.281+03:00 level=INFO msg="received block header" number=35013227 hash=0x241fcdba4685d04ccc904055eaf2c166500c5b1543b19f70035e4bf5372228ba
time=2025-06-02T11:37:42.281+03:00 level=INFO msg="received block header" number=35013226 hash=0x3e0a14a3f183737fc7e4648ccde9e682081ba7b8f95e249fc33f0f5ce0f4c74e
time=2025-06-02T11:37:42.281+03:00 level=INFO msg="received block header" number=35013225 hash=0x372f7d5454f2aaf9fb0374b8e0329411e32615eadb8f218051dce55105fb81bb
time=2025-06-02T11:37:42.281+03:00 level=INFO msg="received block header" number=35013224 hash=0x2e219b5b63beb680e09c9d4b4b1613d493a5c7b0ffc9d44117aa16b74ff5214f
time=2025-06-02T11:37:42.281+03:00 level=INFO msg="received block header" number=35013223 hash=0x66c2fb73c94b1b91398a865d85cd4adbf9f0e6c3ddcb17b1350b14c863a41f78
time=2025-06-02T11:37:52.024+03:00 level=INFO msg="sending random block headers request" start_block=35575873 reverse=false
time=2025-06-02T11:37:52.277+03:00 level=INFO msg="received block header" number=35575873 hash=0xb8b3a9d5a1035ec9699775dc013187120b3225c60f0c2b130558082c6ea91bce
time=2025-06-02T11:37:52.277+03:00 level=INFO msg="received block header" number=35575874 hash=0x7d68ee2bb82b38b5f247519b30b968094e2c5478a14890516d50d0314c996bec
time=2025-06-02T11:37:52.277+03:00 level=INFO msg="received block header" number=35575875 hash=0x7f4ead0403f3dccb3b72e9e6af97c154b4e4d680f1522aafb719f154adc78e10
time=2025-06-02T11:37:52.277+03:00 level=INFO msg="received block header" number=35575876 hash=0x2b49c3af88502b3fa045a9c8c7e1fe89ffb71ed7d83b60e8053bf2937accdfd9
time=2025-06-02T11:37:52.277+03:00 level=INFO msg="received block header" number=35575877 hash=0x0ea565c4e2d766448f6c085b0800e9ff2da4f81cfb3113860a8ee70d3f1ca249
time=2025-06-02T11:38:02.032+03:00 level=INFO msg="sending random block headers request" start_block=35751080 reverse=true
time=2025-06-02T11:38:02.286+03:00 level=INFO msg="received block header" number=35751080 hash=0xdf45a9d10dac16eb8451d2d89eb83dd41c9d2bd64ef2cd1a073b472c50b98509
time=2025-06-02T11:38:02.286+03:00 level=INFO msg="received block header" number=35751079 hash=0xdb9664c60839208a1b10677767ddfa15f97288d850a27746d63c0efb618ec4bb
time=2025-06-02T11:38:02.286+03:00 level=INFO msg="received block header" number=35751078 hash=0xcda624859948319f91a8b72542cd37e4b517e3ddaf9f78d05df7d4579391e862
time=2025-06-02T11:38:02.286+03:00 level=INFO msg="received block header" number=35751077 hash=0xdcbd321f2cfadd2bed3ffc77835062eff01f40aa92dca965cbf8ac120031d25f
time=2025-06-02T11:38:02.286+03:00 level=INFO msg="received block header" number=35751076 hash=0xd6e17efea611f99adb9c64dc8ac90dde05d3a588a3c15f3ef1eb039c0f4a3f37
```

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
