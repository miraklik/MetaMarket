const registrationForm = document.getElementById('registrationForm');
        const errorMessage = document.getElementById('error-message');

        registrationForm.addEventListener('submit', async (e) => {
            e.preventDefault();

            const login = document.getElementById('login').value.trim();
            const email = document.getElementById('email').value.trim();
            const password = document.getElementById('password').value.trim();
            const confirmPassword = document.getElementById('confirm-password').value.trim();

            if (password !== confirmPassword) {
                alert('Passwords do not match!');
                return;
            }

            if (login === "" || email === "" || password === "" || confirmPassword === "") {
                errorMessage.style.display = "block";
            } else {
                errorMessage.style.display = "none";

                // Отправляем данные на сервер
                const response = await fetch('/register', {
                    method: 'POST',
                    headers: { 'Content-Type': 'application/json' },
                    body: JSON.stringify({ nickname, email, password })
                });

                const result = await response.text();
                alert(result); // Сообщение об успехе или ошибке
            }
        });