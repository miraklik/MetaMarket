// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

contract PriceOracle {
    address public owner;
    mapping (address => uint256) public prices;

    constructor() {
        owner = msg.sender;
    }

    function setPrice(address _token, uint256 _price) public {
        require(msg.sender == owner, "Only the owner can set the price");
        prices[_token] = _price;
    }

    function getPrice(address _token) public view returns (uint256) {
        return prices[_token];
    }
}

contract Product {
    address public owner;
    uint256 public price;
    address public priceOracle;

    constructor(address _priceOracle) {
        owner = msg.sender;
        priceOracle = _priceOracle;
    }

    function updatePrice() public {
        uint256 currentPrice = PriceOracle(priceOracle).getPrice(address(this));
        price = currentPrice;
    }

    function buyProduct() public payable {
        require(msg.value >= price, "Insufficient funds");
    }
}