import React, { useState } from "react";
import { createProduct } from "../api";

export default function ProductForm({ onProductCreated }) {
    const [name, setName] = useState("");
    const [description, setDescription] = useState("");
    const [price, setPrice] = useState("");

    const handleSubmit = async (e) => {
        e.preventDefault();
        const newProduct = { name, description, price: parseFloat(price) };
        await createProduct(newProduct);
        onProductCreated();
        setName("");
        setDescription("");
        setPrice("");
    };

    return (
        <form onSubmit={handleSubmit}>
            <h3>Add New Product</h3>
            <input
                type="text"
                placeholder="Name"
                value={name}
                onChange={(e) => setName(e.target.value)}
            />
            <input
                type="text"
                placeholder="Description"
                value={description}
                onChange={(e) => setDescription(e.target.value)}
            />
            <input
                type="number"
                placeholder="Price"
                value={price}
                onChange={(e) => setPrice(e.target.value)}
            />
            <button type="submit">Add Product</button>
        </form>
    );
}