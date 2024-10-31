const loginForm = document.getElementById('loginForm');

        loginForm.addEventListener('submit', (e) => {
            e.preventDefault();
            const login = document.getElementById('login').value;
            const password = document.getElementById('password').value;

            // Логика входа, например, проверка данных пользователя
            console.log("login:", login);
            console.log("Пароль:", password);

            // Перенаправление на index.html после успешного входа
            window.location.href = 'index.html';
        });