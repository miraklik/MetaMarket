// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract myContract {
    
    struct Listing {
        uint amount;
        uint timestamp;
        address from;
        string airdropToken;
    }

    struct Balance{
        uint totalPayments;
        mapping(uint => Listing) listings;
    }

    mapping(address => Balance) public balances;

    function getPayment(address _addr, uint _index) public view returns(Listing memory) {
        return balances[_addr].listings[_index];
    }

    function pay(string memory airdropToken) public payable {
        uint paymentNum = balances[msg.sender].totalPayments++;

        balances[msg.sender].totalPayments++;

        Listing memory newListing = Listing(
            msg.value,
            block.timestamp,
            msg.sender,
            airdropToken
        );

        balances[msg.sender].listings[paymentNum] = newListing;
    }
}