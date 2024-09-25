import React, { useState, useEffect } from 'react';
import Web3 from 'web3';

const ProductFilter = () => {
    const [products, setProducts] = useState([]);
    const [filter, setFilter] = useState({
        price: '',
        rating: '',
        category: ''
    });

    useEffect(() => {
        const web3 = new Web3(window.ethereum);
        const contract = new web3.eth.Contract(abi, address);

        contract.methods.getProducts().call().then((products) => {
            setProducts(products);
        });
    }, []);

    const handleFilterChange = (event) => {
        setFilter({ ...filter, [event.target.name]: event.target.value });
    };

    const filteredProducts = products.filter((product) => {
        if (filter.price && product.price !== filter.price) return false;
        if (filter.rating && product.rating !== filter.rating) return false;
        if (filter.category && product.category !== filter.category) return false;
        return true;
    });

    return (
        <div>
            <h1>Фильтр товаров</h1>
            <form>
                <label>
                    Цена:
                    <input type="number" value={filter.price} onChange={handleFilterChange} name="price" />
                </label>
                <label>
                    Рейтинг:
                    <input type="number" value={filter.rating} onChange={handleFilterChange} name="rating" />
                </label>
                <label>
                    Категория:
                    <select value={filter.category} onChange={handleFilterChange} name="category">
                        <option value="">Все категории</option>
                        <option value="electronics">Электроника</option>
                        <option value="clothing">Одежда</option>
                        <option value="home">Дом и сад</option>
                    </select>
                </label>
            </form>
            <ul>
                {filteredProducts.map((product, index) => (
                    <li key={index}>
                        <p>Название товара: {product.name}</p>
                        <p>Цена: {product.price}</p>
                        <p>Рейтинг: {product.rating}</p>
                        <p>Категория: {product.category}</p>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default ProductFilter;