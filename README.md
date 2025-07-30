# EvmVanityToken
Create A Vanity Address Token In EVM

# Follow the steps below to operate  
1. use https://remix.ethereum.org/ to build a ERC20(It can also be of other types of contracts) contract.
2. copy the contract file binary from remix IDE (The binary file compiled in this step can also be configured for optimization).
3. deploy factory.sol on EVM by remix(or other tools.).
4. edit vanity_generator/main.go
5. fill creationCodeHex by step 2. content
6. fill factoryAddr by step 3. factory.sol address.
7. fill targetSuffix The suffix you desire for the contract.
8. make && ./vanity_generator
9. you can get some output like ```seed[ 0x000000000000000000000000000000000000000000000000000000020000015a ] addr[ 0xA6386c6bEddBEfc83b6f848761D1CAcF65814999 ]``` after step 8. addr is preview of token address, choose you like addr and copy seed.
10. call the factory.sol::deployWithSeed, and fill seed(step 9. copy) and tokenBytecode(step 2. copy) to params to deploy the token.
11. call factory.sol::transferToken to move token to your wallet.