# 🛒 Shopping Cart Application

This is a simple e-commerce shopping cart application built as part of an assignment.

## 🚀 Tech Stack

### Backend
- Go (Gin Framework)
- In-memory storage
- JWT-based authentication (single device login)

### Frontend
- React (Vite)
- Axios
- CSS (custom responsive UI)

---

## 🔑 Features Implemented

- User Signup & Login
- Single-device login using token
- Protected routes
- Items listing
- Add items to cart
- View cart
- Place order (checkout)
- View order history
- Logout functionality
- Responsive UI (3 items per row)

---

## 🔁 Application Flow

1. User signs up or logs in
2. Token is generated and stored
3. User views items
4. Items are added to cart
5. Cart is converted to order
6. Orders can be viewed
7. User can logout

---

## 🔌 API Endpoints

### User
- POST /users → Create user
- GET /users → List users
- POST /users/login → Login user

### Items
- POST /items → Create item
- GET /items → List items

### Cart
- POST /carts → Add item to cart
- GET /carts → View cart

### Orders
- POST /orders → Checkout
- GET /orders → View orders

---

## ▶️ How to Run

### Backend
```bash
cd backend
go run main.go
