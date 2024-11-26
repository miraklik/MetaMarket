// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";

contract FastPaymentSystem is Ownable, ReentrancyGuard {
    IERC20 public stablecoin;

    // Структура для управления транзакциями через эскроу
    struct Escrow {
        address payer;
        address payee;
        uint256 amount;
        bool isCompleted;
    }

    // Хранение информации о транзакциях эскроу
    mapping(uint256 => Escrow) public escrows;
    uint256 public escrowCounter;

    event PaymentInitiated(address indexed payer, address indexed payee, uint256 amount, uint256 escrowId);
    event PaymentCompleted(uint256 indexed escrowId);
    event PaymentCancelled(uint256 indexed escrowId);

    constructor(address _stablecoin) Ownable(msg.sender) {
        require(_stablecoin != address(0), "Stablecoin address cannot be zero");
        stablecoin = IERC20(_stablecoin);
    }

    // Функция для начала платежа через эскроу
    function initiatePayment(address _payee, uint256 _amount) external nonReentrant returns (uint256) {
        require(_payee != address(0), "Payee address cannot be zero");
        require(_amount > 0, "Amount must be greater than zero");
        require(stablecoin.transferFrom(msg.sender, address(this), _amount), "Transfer failed");

        escrowCounter++;
        escrows[escrowCounter] = Escrow({
            payer: msg.sender,
            payee: _payee,
            amount: _amount,
            isCompleted: false
        });

        emit PaymentInitiated(msg.sender, _payee, _amount, escrowCounter);
        return escrowCounter;
    }

    // Функция для завершения платежа через эскроу
    function completePayment(uint256 _escrowId) external nonReentrant {
        Escrow storage escrow = escrows[_escrowId];
        require(!escrow.isCompleted, "Payment already completed");
        require(escrow.payer == msg.sender || msg.sender == owner(), "Only payer or owner can complete payment");

        escrow.isCompleted = true;
        require(stablecoin.transfer(escrow.payee, escrow.amount), "Transfer to payee failed");

        emit PaymentCompleted(_escrowId);
    }

    // Функция для отмены платежа и возврата средств плательщику
    function cancelPayment(uint256 _escrowId) external nonReentrant {
        Escrow storage escrow = escrows[_escrowId];
        require(!escrow.isCompleted, "Payment already completed");
        require(escrow.payer == msg.sender || msg.sender == owner(), "Only payer or owner can cancel payment");

        escrow.isCompleted = true;
        require(stablecoin.transfer(escrow.payer, escrow.amount), "Refund failed");

        emit PaymentCancelled(_escrowId);
    }

    // Функция для изменения адреса стейблкоина (только для владельца)
    function updateStablecoin(address _newStablecoin) external onlyOwner {
        require(_newStablecoin != address(0), "Stablecoin address cannot be zero");
        stablecoin = IERC20(_newStablecoin);
    }
}
