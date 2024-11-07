import React, { useState } from "react";
import ProductList from "../components/ProductList";
import ProductForm from "../components/ProductForm";

export default function ProductsPage() {
    const [refresh, setRefresh] = useState(false);

    const handleProductCreated = () => {
        setRefresh(!refresh);
    };

    return (
        <div>
            <h1>Manage Products</h1>
            <ProductForm onProductCreated={handleProductCreated} />
            <ProductList key={refresh} />
        </div>
    );
}