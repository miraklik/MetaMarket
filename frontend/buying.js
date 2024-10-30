const Web3 = require('web3');
const web3 = new Web3(window.ethereum);

// Адреса контрактов и ABI
const marketplaceAddress = "0x5042a3463C65581A19944EA239f245f05326f1E1"; // Адрес вашего маркетплейса
const usdtAddress = "0x9823dda4Bac5331a6dFe7A2883075A7f3D72Bb64";       // Адрес контракта USDT
const marketplaceABI = "./smart-contract/artifacts/contracts/orders.sol/Marketplace.json"   // ABI контракта маркетплейса
const usdtABI = "./smart-contract/artifacts/@openzeppelin/contracts/token/ERC20/IERC20.sol/IERC20.json"              // ABI контракта USDT

// Инстансы контрактов
const marketplace = new web3.eth.Contract(marketplaceABI, marketplaceAddress);
const usdtToken = new web3.eth.Contract(usdtABI, usdtAddress);

// Пример взаимодействия для покупки товара
async function buyItem(listingId, amount) {
    const accounts = await window.ethereum.request({ method: 'eth_requestAccounts' });
    const buyer = accounts[0];

    // Одобряем контракт маркетплейса для списания средств в USDT
    await usdtToken.methods.approve(marketplaceAddress, amount).send({ from: buyer });

    // Выполняем покупку товара
    await marketplace.methods.purchaseListing(listingId).send({ from: buyer });
}
