# FL-Chain

![Go Version][go-image]
![CosmosSDK Version][sdk-image]
![Node_Version][node-image]

### Software Installation

1.  [Install Git](https://git-scm.com/downloads)
2.  [Install Rust](https://www.rust-lang.org/tools/install)
3.  [Install Go](https://go.dev/)
4.  [Install Node.js](https://nodejs.org/en)
5.  [Install Docker](https://www.docker.com/)
6.  [Install VSCode](https://code.visualstudio.com/download)

### Description

We have built an AppChain for the reliability of federated learning data and learning using Cosmos-SDK. Additionally, federated learning is conducted off-chain through the flwr framework.

A. AppChain (built with CosmosSDK)
Features (Modules) include Record, Token (Mint, Transfer, Lock, UnLock), Event, and ZK Prove.

- Record : Registers Data Info (Hash of Data, MetaData(features, count, memo)) and records learning results.
- Token : An AppChain Token used to pay for learning. When applying for learning, a certain amount of Token is Locked. After learning is completed, participants - prove their learning and UnLock (Transfer) their tokens.
- Event : Emits an Event for the start of federated learning.
- ZK Prove : A Prove Module that verifies honest participation in learning with the recorded Data. It uses Raw Data as a Private Input and Data Hash, w (pre-learning weight), and a (learning rate) as Public Input. The Output is w' (post-learning weight).

Entities include DataProvider and DataConsumer (learning applicants).

The sequence is as follows:

1. The DataProvider records its data with a Hash value and associated metadata (information describing the data).
2. The DataConsumer queries the Chain for data information and finds the required DataProvider.
3. The DataConsumer applies for learning, sending a designated DataProvider List, Model URI, and their verification key (for the ZK Verifier).
4. Chain Nodes that receive this Tx (Token) from step 3 validate the Token, Lock it, and then emit an Event.


B. FlowerFramework (Off-chain)

5. DataProviders that were listening for Events will set up configurations for federated learning once they listen to the event from step 4.
6. The DataConsumer and DataProviders proceed with learning. The reasons for building our AppChain for data and trust resolution are revealed:
-> The authenticity of the RawData used for learning is secured by comparing it with the Hash value registered on the Chain.
-> During learning, DataProviders create a ZK Proof of the learning process using their data and send it to the Chain.
7. The ZK Verifier on the Chain leaves a record of whether the Proof is Valid/Invalid.
8. After the learning is completed, the participating DataProviders apply for Unlock, distribute the Token to the Valid DataProviders, and return the remaining Tokens to the DataConsumer.


This structure encourages both DataProviders and DataConsumers to actively participate without any disruptive behavior in the learning process.




<!-- Markdown link & img dfn's -->
[go-image]: https://img.shields.io/badge/Go-1.21.1-blue
[sdk-image]: https://img.shields.io/badge/CosmosSDK-v0.45.4-purple
[Express-image]: https://img.shields.io/badge/Express-4.18.2-orange
[Node-image]: https://img.shields.io/badge/Node-18.13.0-yellow
