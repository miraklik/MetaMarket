import React, { useState, useEffect } from 'react';
import Web3 from 'web3';

const ProductSearch = () => {
    const [products, setProducts] = useState([]);
    const [search, setSearch] = useState('');

    useEffect(() => {
        const web3 = new Web3(window.ethereum);
        const contract = new web3.eth.Contract(abi, address);


        contract.methods.getProducts().call().then((products) => {
            setProducts(products);
        });
    }, []);

    const handleSearchChange = (event) => {
        setSearch(event.target.value);
    };

    const searchedProducts = products.filter((product) => {
        if (product.name.toLowerCase().includes(search.toLowerCase())) return true;
        if (product.description.toLowerCase().includes(search.toLowerCase())) return true;
        return false;
    });

    return (
        <div>
            <h1>Поиск товаров</h1>
            <form>
                <input type="text" value={search} onChange={handleSearchChange} placeholder="Введите имя или описание товара" />
            </form>
            <ul>
                {searchedProducts.map((product, index) => (
                    <li key={index}>
                        <p>Название товара: {product.name}</p>
                        <p>Описание: {product.description}</p>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default ProductSearch;