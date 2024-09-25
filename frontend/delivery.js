import React, { useState, useEffect } from 'react';
import Web3 from 'web3';

const DigitalDelivery = () => {
    const [digitalProducts, setDigitalProducts] = useState([]);
    const [downloadUrl, setDownloadUrl] = useState('');


    useEffect(() => {
    
        const web3 = new Web3(window.ethereum);
        const contract = new web3.eth.Contract(abi, address);

 
        contract.methods.getDigitalProducts().call().then((products) => {
            setDigitalProducts(products);
        });
    }, []);

    const handleDownload = (productId) => {
        const web3 = new Web3(window.ethereum);
        const contract = new web3.eth.Contract(abi, address);
        contract.methods.getDownloadUrl(productId).call().then((url) => {
            setDownloadUrl(url);
        });
    };

    return (
        <div>
            <h1>Электронная доставка</h1>
            <ul>
                {digitalProducts.map((product, index) => (
                    <li key={index}>
                        <p>Название товара: {product.name}</p>
                        <p>Описание: {product.description}</p>
                        <button onClick={() => handleDownload(product.id)}>Скачать</button>
                    </li>
                ))}
            </ul>
            {downloadUrl && (
                <a href={downloadUrl} download>
                    Скачать цифровой товар
                </a>
            )}
        </div>
    )};