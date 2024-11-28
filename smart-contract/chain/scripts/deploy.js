const { ethers } = require("hardhat");
require("dotenv").config(); 

async function main() {
    const [deployer] = await ethers.getSigners();

    console.log("Deploying contract with the account:", deployer.address);

    const balance = await deployer.getBalance();
    console.log("Deployer's balance:", ethers.utils.formatEther(balance), "ETH");

    const nftContractAddress = "0xd9145CCE52D386f254917e481eB44e9943F39138"; // Замените на адрес вашего ERC721 контракта
    const commissionPercent = 2;

    const Marketplace = await ethers.getContractFactory("Marketplace");

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
