let walletConnected = false;  // Переменная для отслеживания подключения кошелька
let connectedWalletAddress = '';  // Хранит адрес подключенного кошелька

// Настройка провайдера WalletConnect
const walletConnectProvider = new WalletConnectProvider.default({
    rpc: {
        1: "https://mainnet.infura.io/v3/YOUR_INFURA_PROJECT_ID", // Ethereum mainnet
        56: "https://bsc-dataseed.binance.org/", // Binance Smart Chain mainnet
        137: "https://polygon-rpc.com/", // Polygon mainnet
    },
});

// Функция для обновления интерфейса после подключения кошелька
function updateUIAfterConnect(address) {
    walletConnected = true;
    connectedWalletAddress = address;
    const walletActionsDiv = document.getElementById('wallet-actions');
    walletActionsDiv.innerHTML = `
                <span>Кошелек: ${connectedWalletAddress.slice(0, 6)}...${connectedWalletAddress.slice(-4)}</span>
                <button class="disconnect-btn" id="disconnectButton">Отключить кошелек</button>
            `;
    document.getElementById('disconnectButton').addEventListener('click', disconnectWallet);
}

// Функция для отключения кошелька
function disconnectWallet() {
    walletConnected = false;
    connectedWalletAddress = '';
    const walletActionsDiv = document.getElementById('wallet-actions');
    walletActionsDiv.innerHTML = `<button class="wallet-btn" id="connectButton">Подключить кошелек</button>`;
    document.getElementById('connectButton').addEventListener('click', async () => {
        if (!walletConnected) {
            await connectWallet();
        }
    });
    // Если Trust Wallet, отключаем через WalletConnect
    if (walletConnectProvider.connected) {
        walletConnectProvider.disconnect();
    }
}

// Подключение MetaMask
async function connectMetaMask() {
    if (typeof window.ethereum !== 'undefined') {
        console.log('MetaMask обнаружен!');
        try {
            const accounts = await ethereum.request({ method: 'eth_requestAccounts' });
            console.log('Подключенный аккаунт (MetaMask):', accounts[0]);
            updateUIAfterConnect(accounts[0]);
        } catch (error) {
            console.error('Ошибка подключения MetaMask:', error);
        }
    } else {
        alert('Пожалуйста, установите MetaMask!');
    }
}

// Подключение Trust Wallet через WalletConnect
async function connectTrustWallet() {
    try {
        await walletConnectProvider.enable();
        const web3 = new Web3(walletConnectProvider);
        const accounts = await web3.eth.getAccounts();
        console.log('Подключенный аккаунт (Trust Wallet):', accounts[0]);
        updateUIAfterConnect(accounts[0]);

        // Подписка на изменения аккаунтов
        walletConnectProvider.on("accountsChanged", (accounts) => {
            console.log("Аккаунты изменены:", accounts);
            updateUIAfterConnect(accounts[0]);
        });

        // Подписка на изменения сети
        walletConnectProvider.on("chainChanged", (chainId) => {
            console.log("Сеть изменена:", chainId);
        });

        // Отключение провайдера
        walletConnectProvider.on("disconnect", (code, reason) => {
            console.log("WalletConnect отключен:", code, reason);
            disconnectWallet();
        });
    } catch (error) {
        console.error('Ошибка подключения Trust Wallet через WalletConnect:', error);
    }
}

// Подключение Phantom для Solana
async function connectPhantom() {
    if (window.solana && window.solana.isPhantom) {
        console.log('Phantom найден!');
        try {
            const response = await window.solana.connect();
            console.log('Подключенный аккаунт (Phantom):', response.publicKey.toString());
            updateUIAfterConnect(response.publicKey.toString());
        } catch (err) {
            console.error('Ошибка подключения Phantom:', err);
        }
    } else {
        alert('Пожалуйста, установите Phantom Wallet!');
    }
}

// Универсальная функция для подключения кошелька
async function connectWallet() {
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
}

// Инициализируем подключение по клику на кнопку
document.getElementById('connectButton').addEventListener('click', async () => {
    if (!walletConnected) {
        await connectWallet();
    }
});

