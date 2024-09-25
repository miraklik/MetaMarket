import React, { useState, useEffect } from 'react';
import Web3 from 'web3';

const Web3Marketplace = () => {
    const [orders, setOrders] = useState([]);
    const [newOrder, setNewOrder] = useState({
        seller: '',
        price: 0,
        itemName: ''
    });

    useEffect(() => {
        // Подключаемся к смарт-контракту
        const web3 = new Web3(window.ethereum);
        const contract = new web3.eth.Contract(abi, address);

        // Получаем список заказов
        contract.methods.getOrders().call().then((orders) => {
            setOrders(orders);
        });
    }, []);

    const handleCreateOrder = () => {
        // Создаем новый заказ
        const web3 = new Web3(window.ethereum);
        const contract = new web3.eth.Contract(abi, address);
        contract.methods.createOrder(newOrder.seller, newOrder.price, newOrder.itemName).send({
            from: window.ethereum.selectedAddress
        }).then((receipt) => {
            console.log(receipt);
        });
    };

    return (
        <div>
            <h1>Web3 Marketplace</h1>
            <ul>
                {orders.map((order, index) => (
                    <li key={index}>
                        <p>Продавец: {order.seller}</p>
                        <p>Цена: {order.price}</p>
                        <p>Название товара: {order.itemName}</p>
                        <button onClick={() => handleConfirmOrder(order.id)}>Подтвердить</button>
                        <button onClick={() => handleCancelOrder(order.id)}>Отменить</button>
                    </li>
                ))}
            </ul>
            <form>
                <input type="text" value={newOrder.seller} onChange={(e) => setNewOrder({ ...newOrder, seller: e.target.value })} placeholder="Продавец" />
                <input type="number" value={newOrder.price} onChange={(e) => setNewOrder({ ...newOrder, price: e.target.value })} placeholder="Цена" />
                <input type="text" value={newOrder.itemName} onChange={(e) => setNewOrder({ ...newOrder, itemName: e.target.value })} placeholder="Название товара" />
                <button onClick={handleCreateOrder}>Создать заказ</button>
            </form>
        </div>
    );
};

export default Web3Marketplace;