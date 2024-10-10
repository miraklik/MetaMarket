// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "../.deps/npm/@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/secu";

contract Marketplace is ReentrancyGuard {
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

    address owner;
    uint public listingCount;
    uint256 public totalOrders;
    mapping(uint => Listing) public listings;
    mapping(uint => Order) public orders;

    // Адрес контракта USDT
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

    event EscrowReleased(uint indexed listingId, address indexed seller, uint amount);
    event EscrowCancelled(uint indexed listingId, address indexed buyer, uint amount);

    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can perform this action");
        _;
    }

    constructor(address _usdtTokenAddress) {
        owner = msg.sender; // Устанавливаем владельца контракта
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
        require(bytes(_title).length > 0, "Title cannot be empty");
        require(bytes(_description).length > 0, "Description cannot be empty");

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
    function purchaseListing(uint _listingId) public nonReentrant {
        Listing storage listing = listings[_listingId];
        require(_listingId > 0 && _listingId <= listingCount, "Invalid listing ID");
        require(!listing.sold, "Item already sold");
        require(
            usdtToken.allowance(msg.sender, address(this)) >= listing.price,
            "Not enough allowance"
        );

        // Блокируем средства в эскроу
        bool success = usdtToken.transferFrom(msg.sender, address(this), listing.price);
        require(success, "USDT transfer failed");

        escrowBuyer[_listingId] = msg.sender;
        escrowAmount[_listingId] = listing.price;

        listing.sold = true;

        emit PurchaseCompleted(_listingId, msg.sender, listing.price);
    }

    // Подтверждение получения товара покупателем и выплата продавцу
    function confirmPurchase(uint _listingId) public nonReentrant {
        require(
            escrowBuyer[_listingId] == msg.sender,
            "Only buyer can confirm purchase"
        );
        require(escrowAmount[_listingId] > 0, "No funds in escrow");

        Listing storage listing = listings[_listingId];

        // Выплата продавцу
        bool success = usdtToken.transfer(listing.seller, escrowAmount[_listingId]);
        require(success, "USDT transfer to seller failed");

        // Очищаем эскроу
        escrowAmount[_listingId] = 0;
        escrowBuyer[_listingId] = address(0);

        emit EscrowReleased(_listingId, listing.seller, escrowAmount[_listingId]);
    }

    // Отмена заказа, возвращение средств покупателю
    function cancelOrder(uint _listingId) public nonReentrant {
        require(escrowBuyer[_listingId] == msg.sender, "Only the buyer can cancel the order");
        require(escrowAmount[_listingId] > 0, "No funds in escrow");
        require(!listings[_listingId].sold, "Order already completed, cannot cancel");

        uint amount = escrowAmount[_listingId];

        // Возвращаем средства покупателю
        bool success = usdtToken.transfer(escrowBuyer[_listingId], amount);
        require(success, "USDT transfer to buyer failed");

        // Очищаем эскроу
        escrowAmount[_listingId] = 0;
        escrowBuyer[_listingId] = address(0);

        emit EscrowCancelled(_listingId, msg.sender, amount);
    }

    // Функция для владельца контракта вывести все токены
    function withdrawToken() public onlyOwner {
        uint balance = usdtToken.balanceOf(address(this));
        require(balance > 0, "No tokens to withdraw");

        bool success = usdtToken.transfer(owner, balance);
        require(success, "Withdraw failed");
    }
}
