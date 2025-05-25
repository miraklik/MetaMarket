# 🛍️ MetaMarket — Web3 NFT Platform

**MetaMarket** is a decentralized platform that allows users to create, manage, and trade NFTs. The project includes smart contracts written in Solidity, a Golang backend, a React-based frontend, and monitoring via Prometheus.

---

## 🚀 Features

- 📦 Mint and manage NFTs
- 🌐 Blockchain integration (Ethereum, BSC, etc.)
- 💼 User-friendly interface
- 🧠 Smart contract interaction via Web3
- 📊 Monitoring with Prometheus

---

## 🛠️ Tech Stack

- **Frontend:** React.js, Ethers.js
- **Backend:** Golang, Geth, Gin, GORM, PostgreSQL
- **Blockchain:** Solidity, Hardhat, Openzeppelin
- **Web3 Tools:** Ethers.js, Web3.js
- **Monitoring:** Prometheus
- **Other:** JavaScript, REST API, Docker (optional)

---

## 📁 Project Structure
MetaMarket/
├── backend/ # Golang server
├── frontend/ # React frontend
├── smart-contract/ # Solidity smart contracts
├── prometheus/ # Prometheus monitoring config
├── .vscode/ # Editor settings
├── .gitignore
├── README.md
└── package.json

## 📦 Installation & Setup

### 1. Clone the repository

```bash
git clone https://github.com/miraklik/MetaMarket.git
cd MetaMarket
```

### 2. Install dependencies

#### 2.1. Backend

```bash
cd backend
go mod tidy
```

#### 2.2. Frontend

```bash
cd frontend
npm install
```

## 🔨 Deploy Smart Contracts

```bash
cd smart-contract/chain
npm install
npx hardhat compile
npx hardhat run scripts/deploy.js --network mainnet
```


## 🖥️ Run the App

### 1. Backend

```bash
cd backend
go run cmd/worker/main.go
```

### 2. Frontend

```bash
cd ../frontend
npm start
```

## 🔍 Monitoring with Prometheus

```bash
cd prometheus
docker-compose up
```
