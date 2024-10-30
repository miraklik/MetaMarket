// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;

import "../.deps/npm/@openzeppelin/contracts/token/ERC20/IERC20.sol";

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

    struct Order {
        bool isConfirmed;
    }

    address public owner;
    uint public listingCount;
    uint256 public totalOrders;
    uint public commissionPercent; // проценты комиссии 
    mapping(uint => Listing) public listings;
    mapping(uint => Order) public orders;

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

    event PurchaseCancelled(
        uint indexed id,
        address indexed buyer,
        uint price
    );

    event EscrowReleased(uint indexed listingId, address indexed seller, uint amount);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can perform this action");
        _;
    }

    constructor(address _usdtTokenAddress, uint8 _commissionPercent) {
        owner = msg.sender; // Устанавливаем владельца контракта
        usdtToken = IERC20(_usdtTokenAddress); // Указать адрес USDT контракта
        commissionPercent = _commissionPercent;
    }

    // Создание объявления
    function createListing(
        string memory _title,
        string memory _description,
        string memory _imageIPFSHash,
        uint _price
    ) public {
        require(_price > 0, "Price must be greater than 0");
        require(bytes(_title).length > 0, "Title cannot be empty");
        require(bytes(_description).length > 0, "Description cannot be empty");

        // Дополнительные проверки
        require(_price <= 1e18, "Price is too high"); // Максимальная цена — 1e18 (примерно 1 USDT в wei)

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

    // Покупка товара с использованием USDT и с комиссией
    function purchaseListing(uint _listingId) public {
        Listing storage listing = listings[_listingId];
        require(_listingId > 0 && _listingId <= listingCount, "Invalid listing ID");
        require(!listing.sold, "Item already sold");
        require(
            usdtToken.allowance(msg.sender, address(this)) >= listing.price,
            "Not enough allowance"
        );

        // Рассчитываем комиссию
        uint commissionAmount = (listing.price * commissionPercent) / 100;
        uint sellerAmount = listing.price - commissionAmount;

        // Переводим полную сумму с покупателя на контракт
        bool success = usdtToken.transferFrom(msg.sender, address(this), listing.price);
        require(success, "Transfer failed");

        // Переводим комиссию владельцу контракта
        success = usdtToken.transfer(owner, commissionAmount);
        require(success, "Commission transfer failed");

        // Переводим оставшиеся средства продавцу
        success = usdtToken.transfer(listing.seller, sellerAmount);
        require(success, "Seller payment failed");

        listing.sold = true;

        emit PurchaseCompleted(_listingId, msg.sender, listing.price);
    }

    // Функция изменения процента комиссии (только для владельца контракта)
    function setCommissionPercent(uint _newPercent) public {
        require(msg.sender == owner, "Only owner can change commission");
        require(_newPercent <= 100, "Commission cannot exceed 100%");
        commissionPercent = _newPercent;
    }

    // Подтверждение получения товара покупателем и выплата продавцу
    function confirmPurchase(uint _listingId) public {
        require(
            escrowBuyer[_listingId] == msg.sender,
            "Only buyer can confirm purchase"
        );
        require(escrowAmount[_listingId] > 0, "No funds in escrow");
        require(listings[_listingId].sold, "Listing is not sold yet");

        Listing storage listing = listings[_listingId];

        // Выплата продавцу
        bool success = usdtToken.transfer(listing.seller, escrowAmount[_listingId]);
        require(success, "USDT transfer to seller failed");

        // Очищаем эскроу
        escrowAmount[_listingId] = 0;
        escrowBuyer[_listingId] = address(0);

        emit EscrowReleased(_listingId, listing.seller, escrowAmount[_listingId]);
    }

    // Отмена покупки, возврат средств покупателю
    function cancelPurchase(uint _listingId) public {
        require(escrowBuyer[_listingId] == msg.sender, "Only the buyer can cancel");
        require(escrowAmount[_listingId] > 0, "No funds in escrow");
        require(!listings[_listingId].sold, "Cannot cancel, item already sold");

        uint amount = escrowAmount[_listingId];

        // Возврат средств покупателю
        bool success = usdtToken.transfer(escrowBuyer[_listingId], amount);
        require(success, "USDT transfer to buyer failed");

        // Очищаем эскроу
        escrowAmount[_listingId] = 0;
        escrowBuyer[_listingId] = address(0);

        emit PurchaseCancelled(_listingId, msg.sender, amount);
    }
    
    // Функция для владельца контракта вывести все токены
    function withdrawToken() public onlyOwner {
        uint balance = usdtToken.balanceOf(address(this));
        require(balance > 0, "No tokens to withdraw");

        bool success = usdtToken.transfer(owner, balance);
        require(success, "Withdraw failed");
    }
}
