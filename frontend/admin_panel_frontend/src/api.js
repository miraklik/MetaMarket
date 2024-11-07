const API_BASE = "http://localhost:8080/admin";


export async function fetchProducts() {
    const response = await fetch(`${API_BASE}/products`);
    return response.json();
}


export async function createProduct(product) {
    const response = await fetch(`${API_BASE}/products`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json"
        },
        body: JSON.stringify(product)
    });
    return response.json();
}