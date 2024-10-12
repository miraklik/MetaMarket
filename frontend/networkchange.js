const Web3 = require('web3');

// В зависимости от выбранной сети устанавливаем адрес контракта USDT
let usdtAddress;
let networkId = await web3.eth.net.getId(); // Получаем идентификатор сети

if (networkId === 1) {  // Ethereum Mainnet
    usdtAddress = "0xdAC17F958D2ee523a2206206994597C13D831ec7";
} else if (networkId === 56) {  // Binance Smart Chain
    usdtAddress = "0x55d398326f99059ff775485246999027b3197955";
} else if (networkId === 137) {  // Polygon
    usdtAddress = "0x3813e0826c49b429c1a28dbeed4a150d23840e0d";
}

// Инициализируем контракт с правильным адресом USDT для текущей сети
const usdtContract = new web3.eth.Contract(usdtAbi, usdtAddress);

// Далее работаем с контрактом, например, одобряем средства для отправки на маркетплейс
await usdtContract.methods.approve(marketplaceAddress, amount).send({ from: userAddress });
