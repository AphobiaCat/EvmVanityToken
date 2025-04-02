// SPDX-License-Identifier: MIT
pragma solidity ^0.8.24;

interface IBEP20 {
    function transfer(address recipient, uint256 amount) external returns (bool);
    function balanceOf(address account) external view returns (uint256);
}

contract Factory {
    event TokenDeployed(address indexed tokenAddress, bytes32 indexed seed);
    event TokenTransferred(address indexed from, address indexed to, uint256 amount);

    mapping(address => address) private token_owner;

    function predictAddress(bytes32 seed, bytes calldata tokenBytecode) public view returns (address) {
        bytes32 hash = keccak256(
            abi.encodePacked(
                hex"ff",
                address(this),
                seed,
                keccak256(tokenBytecode)
            )
        );
        return address(uint160(uint256(hash)));
    }

    function deployWithSeed(bytes32 seed, bytes calldata tokenBytecode) external returns (address) {
        address tokenAddress;
        bytes memory bytecode = tokenBytecode;
        
        assembly {
            tokenAddress := create2(0, add(bytecode, 0x20), mload(bytecode), seed)
        }
        
        require(tokenAddress != address(0), "Create2 failed");

        token_owner[tokenAddress] = msg.sender;

        emit TokenDeployed(tokenAddress, seed);
        return tokenAddress;
    }

    function transferToken(address tokenAddress, uint256 amount) external {
        require(token_owner[tokenAddress] == msg.sender, "Not your token");
        
        IBEP20 token = IBEP20(tokenAddress);  // access token by IBEP20
        
        //uint256 amount = token.balanceOf(address(this));

        // call token transfer
        bool success = token.transfer(msg.sender, amount);
        require(success, "Token transfer failed");

        emit TokenTransferred(address(this), msg.sender, amount);
    }
}
