document.addEventListener('DOMContentLoaded', () => {

    const connectWalletButton = document.querySelector('.connect-wallet');
    const walletAddressElement = document.getElementById('wallet-address');
    const tokenBalanceElement = document.getElementById('token-balance');
    const connectedWalletsElement = document.getElementById('connected-wallets');

    // Mock: Connect Wallet (Simulate Web3 Wallet Connect)
    connectWalletButton.addEventListener('click', async () => {
        // Mock wallet address (in a real app, integrate Web3.js or ethers.js)
        const walletAddress = '0x123...abc';
        walletAddressElement.textContent = walletAddress;
        connectWalletButton.textContent = 'Connected';
        connectWalletButton.disabled = true;

        // Mock token balance
        const tokenBalance = 125.50; // Assume 125.50 USDT balance
        tokenBalanceElement.textContent = `${tokenBalance} USDT`;

        // Update connected wallets section
        connectedWalletsElement.textContent = 'Metamask, WalletConnect';
    });

    // Handle logout (disconnect wallet)
    const logoutButton = document.querySelector('.logout');
    logoutButton.addEventListener('click', () => {
        alert('Wallet disconnected!');
        walletAddressElement.textContent = 'Not Connected';
        tokenBalanceElement.textContent = '0 USDT';
        connectedWalletsElement.textContent = 'No wallets connected';
        connectWalletButton.textContent = 'Connect Wallet';
        connectWalletButton.disabled = false;
    });
});
