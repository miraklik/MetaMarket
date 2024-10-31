// Инициализация провайдера WalletConnect для Trust Wallet
const walletConnectProvider = new WalletConnectProvider.default({
    rpc: {
        1: "https://mainnet.infura.io/v3/dfa0335a8d2b4364bd669159aa3dc734", // Ethereum mainnet
        56: "https://bsc-dataseed.binance.org/", //"https://bsc-mainnet.infura.io/v3/dfa0335a8d2b4364bd669159aa3dc734", // Binance Smart Chain mainnet
        137: "https://polygon-rpc.com/", //"https://polygon-mainnet.infura.io/v3/dfa0335a8d2b4364bd669159aa3dc734", // Polygon mainnet
    },
});

// Проверка наличия MetaMask
async function connectMetaMask() {
    if (typeof window.ethereum !== 'undefined') {
        console.log('MetaMask обнаружен!');
        try {
            const accounts = await ethereum.request({ method: 'eth_requestAccounts' });
            console.log('Подключенный аккаунт (MetaMask):', accounts[0]);
            const provider = new ethers.providers.Web3Provider(window.ethereum);
            const signer = provider.getSigner();
            console.log('MetaMask провайдер и подписант инициализированы:', provider, signer);
        } catch (error) {
            console.error('Ошибка подключения MetaMask:', error);
        }
    } else {
        alert('Пожалуйста, установите MetaMask!');
    }
}

// Функция для подключения Trust Wallet через WalletConnect
async function connectTrustWallet() {
    try {
        await walletConnectProvider.enable();
        const web3 = new Web3(walletConnectProvider);
        const accounts = await web3.eth.getAccounts();
        console.log('Подключенный аккаунт (Trust Wallet):', accounts[0]);

        // Подписка на изменения аккаунта
        walletConnectProvider.on("accountsChanged", (accounts) => {
            console.log("Аккаунты изменены:", accounts);
        });

        // Подписка на изменения сети
        walletConnectProvider.on("chainChanged", (chainId) => {
            console.log("Сеть изменена:", chainId);
        });

        // Отключение провайдера
        walletConnectProvider.on("disconnect", (code, reason) => {
            console.log("WalletConnect отключен:", code, reason);
        });
    } catch (error) {
        console.error('Ошибка подключения Trust Wallet:', error);
    }
}

// Проверка наличия Phantom для Solana
async function connectPhantom() {
    if (window.solana && window.solana.isPhantom) {
        console.log('Phantom найден!');
        try {
            const response = await window.solana.connect();
            console.log('Подключенный аккаунт (Phantom):', response.publicKey.toString());
        } catch (err) {
            console.error('Ошибка подключения Phantom:', err);
        }
    } else {
        alert('Пожалуйста, установите Phantom Wallet!');
    }
}

// Единый обработчик для кнопки "Connect Wallet"
document.getElementById('connectButton').addEventListener('click', async () => {
    if (typeof window.ethereum !== 'undefined') {
        // Подключаем MetaMask
        await connectMetaMask();
    } else if (window.solana && window.solana.isPhantom) {
        // Подключаем Phantom (Solana)
        await connectPhantom();
    } else {
        // Подключаем Trust Wallet через WalletConnect
        await connectTrustWallet();
    }
});