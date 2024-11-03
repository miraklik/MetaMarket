const registrationForm = document.getElementById('register-form');
const errorMessage = document.getElementById('error-message');

registrationForm.addEventListener('submit', async (e) => {
    e.preventDefault();

    const login = document.getElementById('login').value.trim();
    const email = document.getElementById('email').value.trim();
    const password = document.getElementById('password').value.trim();
    const confirmPassword = document.getElementById('confirm-password').value.trim();

    // * Проверка на длину пароля
    if (password.length < 5) {
        errorMessage.innerText = "Пароль должен быть не менее 5 символов.";
        errorMessage.style.display = "block";
        return;
    }

    // * проверка что пароли подходят
    if (password !== confirmPassword) {
        errorMessage.innerText = "Пароли не совпадают.";
        errorMessage.style.display = "block";
        return;
    }

    // * Проверка на пустые поля
    if (login === "" || email === "" || password === "" || confirmPassword === "") {
        errorMessage.innerText = "Все поля обязательны для заполнения.";
        errorMessage.style.display = "block";
        return;
    }

    errorMessage.style.display = "none";

    // * Отправляем данные на сервер
    try {
        const response = await fetch('/register', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ nickname: login, email, password })
        });

        const result = await response.text();
        alert(result); // Сообщение об успехе или ошибке
        window.location.href = 'login.html';
    } catch (error) {
        errorMessage.innerText = "Произошла ошибка при отправке данных. Пожалуйста, попробуйте позже.";
        errorMessage.style.display = "block";
    }

    fetch('http://localhost:44044/api/data')
    .then(response => response.json())
    .then(data => {
        console.log(data.message);
    })
    .catch(error => console.error('Error:', error));

});