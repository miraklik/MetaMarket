// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
     
    //Merkle tree

contract Test {
    // TX1 TX2 TX3 TX4
    bytes32[] public hashes;
    string[4] transaction = [
        "TX1: 0x2945 -> 0xdk24124",
        "TX2: 0xdmc2 -> 0xlaoe4",
        "TX3: 0x2031 -> 0xlorj24",
        "TX4: 0xlsi4 -> 0x1dcef"
    ];

    constructor() {
        for(uint i = 0; i < transaction.length; i++) {
            hashes.push(makeHash(transaction[i]));
        }

        uint count = transaction.length;
        uint offset = 0;

        while(count > 0) {
            for (uint i = 0; i < count - 1; i += 2) {
                hashes.push(keccak256(
                    abi.encodePacked(
                        hashes[offset + i], hashes[offset + i + 1]
                    )
                ));
            }
            offset += count;
            count /= 2;
        }
    }

    function verify(string memory transactions, uint index, bytes32 root, bytes32[] memory proof) public pure returns(bool) {
        bytes32 hash = makeHash(transactions);
        for(uint i = 0; i < proof.length; i++) {
            bytes32 element = proof[i];
            if(index % 2 == 0){
                hash = keccak256(abi.encodePacked(hash, element));
            } else {
                hash = keccak256(abi.encodePacked(element, hash));
            }

            index /= 2;
        }

        return hash == root;
    }

    function encode(string memory input) public pure returns(bytes memory) {
        bytes memory encoded = abi.encodePacked(input);
        return encoded;
    }

    function makeHash(string memory input) public pure returns(bytes32) {
        bytes32 hash = keccak256(encode(input));
        return hash;
    }

}