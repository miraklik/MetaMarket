// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "../.deps/npm/@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "./chancheprice.sol";

contract Marketplace {
    // Структура объявления
    struct Listing {
        uint id;
        address seller;
        string title;
        string description;
        string imageIPFSHash;
        uint price;
        bool sold;
    }

    address owner;
    uint public listingCount;
    uint256 public totalOrders;
    mapping(uint => Listing) public listings;
    mapping(uint => address) public orders;
    

    // Адрес контракта USDT (для работы с любой сетью, контракт адрес необходимо указать отдельно)
    IERC20 public usdtToken;

    // Эскроу для блокировки средств
    mapping(uint => address) public escrowBuyer;
    mapping(uint => uint) public escrowAmount;

    event ListingCreated(
        uint indexed id,
        address indexed seller,
        string title,
        uint price
    );

    event PurchaseCompleted(
        uint indexed id,
        address indexed buyer,
        uint price
    );

    constructor(address _usdtTokenAddress) {
        usdtToken = IERC20(_usdtTokenAddress); // Указать адрес USDT контракта
    }

    // Создание объявления
    function createListing(
        string memory _title,
        string memory _description,
        string memory _imageIPFSHash,
        uint _price
    ) public {
        require(_price > 0, "Price must be greater than 0");

        listingCount++;
        listings[listingCount] = Listing(
            listingCount,
            msg.sender,
            _title,
            _description,
            _imageIPFSHash,
            _price,
            false
        );

        emit ListingCreated(listingCount, msg.sender, _title, _price);
    }

    // Покупка товара с использованием USDT
    function purchaseListing(uint _listingId) public {
        Listing storage listing = listings[_listingId];
        require(!listing.sold, "Item already sold");
        require(
            usdtToken.allowance(msg.sender, address(this)) >= listing.price,
            "Not enough allowance"
        );

        // Блокируем средства в эскроу
        usdtToken.transferFrom(msg.sender, address(this), listing.price);
        escrowBuyer[_listingId] = msg.sender;
        escrowAmount[_listingId] = listing.price;

        listing.sold = true;

        emit PurchaseCompleted(_listingId, msg.sender, listing.price);
    }

    // Подтверждение получения товара покупателем и выплата продавцу
    function confirmPurchase(uint _listingId) public {
        require(
            escrowBuyer[_listingId] == msg.sender,
            "Only buyer can confirm purchase"
        );
        require(escrowAmount[_listingId] > 0, "No funds in escrow");

        Listing storage listing = listings[_listingId];

        // Выплата продавцу
        usdtToken.transfer(listing.seller, escrowAmount[_listingId]);

        // Очищаем эскроу
        escrowAmount[_listingId] = 0;
        escrowBuyer[_listingId] = address(0);
    }

     /*function cancelOrder(uint256 _orderId) public {
        // Проверяем, что заказ существует
        require(_orderId < totalOrders);

        // Проверяем, что заказ не подтвержден ранее
        require(!orders[_orderId].isConfirmed);

        // Отменяем заказ
        orders[_orderId].isConfirmed = false;
    }*/

    function withdrawToken() public {
        address _to = payable(owner);
        address _thisContract = address(this);
        usdtToken.transfer(_to, usdtToken.balanceOf(_thisContract));
    } 
}