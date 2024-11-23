// Import ethers.js from the Hardhat runtime environment
const { ethers } = require("hardhat");
require("dotenv").config(); // Load environment variables from .env

async function main() {
    // Deployer's wallet address
    const [deployer] = await ethers.getSigners();

    console.log("Deploying contract with the account:", deployer.address);

    // Log the deployer's balance
    const balance = await deployer.getBalance();
    console.log("Deployer's balance:", ethers.utils.formatEther(balance), "ETH");

    // ERC721 contract address and marketplace commission percentage
    const nftContractAddress = "0xd9145CCE52D386f254917e481eB44e9943F39138"; // Replace with the deployed ERC721 contract address
    const commissionPercent = 2; // Example: 2% commission

    // Get the Marketplace contract factory
    const Marketplace = await ethers.getContractFactory("Marketplace");

    // Deploy the Marketplace contract
    const marketplace = await Marketplace.deploy(nftContractAddress, commissionPercent);

    console.log("Waiting for deployment...");
    await marketplace.deployed();

    console.log("Marketplace deployed to:", marketplace.address);
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });
