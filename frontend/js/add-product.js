const productForm = document.getElementById('productForm');

        productForm.addEventListener('submit', (e) => {
            e.preventDefault();

            const name = document.getElementById('productName').value;
            const price = document.getElementById('productPrice').value;
            const description = document.getElementById('productDescription').value;
            const image = document.getElementById('productImage').files[0];

            const reader = new FileReader();
            reader.onload = function(event) {
                const imageUrl = event.target.result;

                // Получаем текущий список товаров из локального хранилища
                const products = JSON.parse(localStorage.getItem('products')) || [];

                // Добавляем новый товар в список
                const newProduct = {
                    name: name,
                    price: price,
                    description: description,
                    image: imageUrl
                };
                products.push(newProduct);

                // Сохраняем обновленный список товаров в локальное хранилище
                localStorage.setItem('products', JSON.stringify(products));

                // Перенаправляем пользователя обратно на главную страницу
                window.location.href = 'index.html';
            };

            reader.readAsDataURL(image); // Чтение изображения как URL
        });