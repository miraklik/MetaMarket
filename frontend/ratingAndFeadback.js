import React, { useState, useEffect } from 'react';
import Web3 from 'web3';

const RatingAndReviews = () => {
    const [sellerRating, setSellerRating] = useState(0);
    const [reviews, setReviews] = useState([]);
    const [newReview, setNewReview] = useState({
        rating: 0,
        comment: ''
    });

    useEffect(() => {
        const web3 = new Web3(window.ethereum);
        const contract = new web3.eth.Contract(abi, address);

        // Получаем рейтинг продавца
        contract.methods.getSellerRating().call().then((rating) => {
            setSellerRating(rating);
        });

        contract.methods.getReviews().call().then((reviews) => {
            setReviews(reviews);
        });
    }, []);

    const handleNewReviewChange = (event) => {
        setNewReview({ ...newReview, [event.target.name]: event.target.value });
    };

    const handleAddReview = () => {
        // Создаем новый отзыв
        const web3 = new Web3(window.ethereum);
        const contract = new web3.eth.Contract(abi, address);
        contract.methods.addReview(newReview.rating, newReview.comment).send({
            from: window.ethereum.selectedAddress
        }).then((receipt) => {
            console.log(receipt);
        });
    };

    return (
        <div>
            <h1>Рейтинг и отзывы</h1>
            <p>Рейтинг продавца: {sellerRating}</p>
            <ul>
                {reviews.map((review, index) => (
                    <li key={index}>
                        <p>Рейтинг: {review.rating}</p>
                        <p>Комментарий: {review.comment}</p>
                    </li>
                ))}
            </ul>
            <form>
                <label>
                    Рейтинг:
                    <input type="number" value={newReview.rating} onChange={handleNewReviewChange} name="rating" />
                </label>
                <label>
                    Комментарий:
                    <textarea value={newReview.comment} onChange={handleNewReviewChange} name="comment" />
                </label>
                <button onClick={handleAddReview}>Добавить отзыв</button>
            </form>
        </div>
    );
};

export default RatingAndReviews;