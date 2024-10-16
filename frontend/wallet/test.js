const nodemailer = require('nodemailer');
const express = require('express');
const bodyParser = require('body-parser');
const app = express();

app.use(bodyParser.json());

// Настройка транспортера для отправки email
const transporter = nodemailer.createTransport({
    service: 'Gmail',
    auth: {
        user: 'karmau473@gmail.com', // Замените на свою почту
        pass: 'mirak1991'   // Введите пароль от почты или используйте app-specific password
    }
});

// Маршрут для регистрации пользователя и отправки кода на email
app.post('/register', (req, res) => {
    const { email, nickname, password } = req.body;

    // Генерируем случайный код подтверждения (например, 6-значный код)
    const confirmationCode = Math.floor(100000 + Math.random() * 900000);

    // Отправляем письмо с кодом
    const mailOptions = {
        from: 'karmau473@gmail.com',
        to: email,
        subject: 'Подтверждение регистрации',
        text: `Ваш код подтверждения: ${confirmationCode}`
    };

    transporter.sendMail(mailOptions, (error, info) => {
        if (error) {
            return res.status(500).send('Ошибка при отправке email: ' + error.message);
        } else {
            console.log('Email отправлен: ' + info.response);
            // Можно сохранить код в базе данных или временном хранилище
            res.status(200).send('Регистрация прошла успешно! Проверьте свою почту.');
        }
    });
});

// Запуск сервера
app.listen(3000, () => {
    console.log('Сервер запущен на порту 3000');
});
