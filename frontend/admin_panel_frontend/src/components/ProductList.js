import React, { useEffect, useState } from "react";
import { fetchProducts } from "../api";

export default function ProductList() {
    const [products, setProducts] = useState([]);

    useEffect(() => {
        async function loadProducts() {
            const products = await fetchProducts();
            setProducts(products);
        }
        loadProducts();
    }, []);

    return (
        <div>
            <h2>Products</h2>
            <ul>
                {products.map((product) => (
                    <li key={product.id}>
                        {product.name} - ${product.price}
                    </li>
                ))}
            </ul>
        </div>
    );
}