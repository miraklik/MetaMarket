// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

contract Payments {

    struct Payment {
        uint amount;
        uint timestamp;
        address from;
        string message;
    }

    struct Balance {
        uint totalPayments;
        mapping(uint => Payment) payments;
    }

    mapping(address => Balance) public balances;

    function currentBalance() public view returns(uint) {
        return address(this).balance;
    }

    function getPayment(address _addr, uint _index) public view returns(Payment memory) {
        require(_index < balances[_addr].totalPayments, "Payment does not exist");
        return balances[_addr].payments[_index];
    }

    function pay(string memory message) public payable {
        require(msg.value < 0, "there are no funds in the wallet");

        uint paymentNum = balances[msg.sender].totalPayments;
        balances[msg.sender].totalPayments++;
        Payment memory newPayment = Payment(
            msg.value,
            block.timestamp,
            msg.sender,
            message
        );
        balances[msg.sender].payments[paymentNum] = newPayment;
    }
}