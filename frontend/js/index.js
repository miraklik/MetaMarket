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
                <button class="v56_34" id="disconnectButton">Disconnect</button>
            `;
            document.getElementById('disconnectButton').addEventListener('click', disconnectWallet);
        }

        // Функция для отключения кошелька
        function disconnectWallet() {
            walletConnected = false;
            connectedWalletAddress = '';
            const walletActionsDiv = document.getElementById('wallet-actions');
            walletActionsDiv.innerHTML = `<button class="v56_34" id="connectButton">Connect</button>`;
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
        
        // Получение списка товаров (если их нужно загрузить из базы данных или локального хранилища)
        const products = JSON.parse(localStorage.getItem('products')) || [];

        // Контейнер для отображения товаров
        const productsContainer = document.getElementById('productsContainer');

        // Функция для отображения товаров на главной странице
        function displayProducts() {
            productsContainer.innerHTML = ''; // Очистка контейнера
            products.forEach(product => {
                const productCard = document.createElement('div');
                productCard.classList.add('product');

                const productImage = document.createElement('img');
                productImage.src = product.image;

                const productName = document.createElement('h3');
                productName.textContent = product.name;

                const productPrice = document.createElement('p');
                productPrice.classList.add('price');
                productPrice.textContent = `${product.price} $`;

                const productDescription = document.createElement('p');
                productDescription.textContent = product.description;

                const deleteButton = document.createElement('button');
                deleteButton.classList.add('delete-btn');
                deleteButton.textContent = 'Удалить';
                deleteButton.addEventListener('click', () => deleteProduct(index));

                productCard.appendChild(productImage);
                productCard.appendChild(productName);
                productCard.appendChild(productPrice);
                productCard.appendChild(productDescription);

                productsContainer.appendChild(productCard);
            }); 
        }

        // Функция для удаления товара
        function deleteProduct(index) {
            // Удаляем товар из массива
            products.splice(index, 1);

            // Сохраняем обновленный список товаров в локальное хранилище
            localStorage.setItem('products', JSON.stringify(products));

            // Обновляем отображение товаров
            displayProducts();
        }

        // Вызов функции для отображения товаров при загрузке страницы
        displayProducts();