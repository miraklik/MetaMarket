const { ethers } = require("hardhat");
require("dotenv").config();
import { Web3Storage } from 'web3.storage'

async function main() {
    const [deployer] = await ethers.getSigners();
    console.log("Deploying contract with the account:", deployer.address);

    const balance = await deployer.getBalance();
    console.log("Deployer's balance:", ethers.utils.formatEther(balance), "ETH");

    if (balance.lt(ethers.utils.parseEther("0.01"))) {
        throw new Error("Insufficient balance for deployment. Ensure the deployer wallet has enough ETH.");
    }

    const MakeStorageClient = () => {
        return new Web3Storage({token: process.env.WEB3STORAGE_TOKEN})
    }

    const saveToIpfs = async (files) => {
        const client = MakeStorageClient()
        const cid = await client.put([files])
        console.log(`File stored with CID: ${cid}`)
        return cid
      }

    const retriveData = async (cid) => {
        const client = MakeStorageClient()
        const res = await client.get(cid)
        console.log(`Got a response! [${res.status}] ${res.statusText}`)
        if (!res.ok) {
          throw new Error(`failed to get ${cid} - ${res.statusText}`)
        }

        const files = await res.files()
        for (const file of files) {
          console.log(`${file.cid} -- ${file.path} -- ${file.size}`)
        }
    }

    const nftContractAddress = "0x9823dda4Bac5331a6dFe7A2883075A7f3D72Bb64";
    const commissionPercent = 1;

    try {
        const Marketplace = await ethers.getContractFactory("Marketplace");
        const marketplace = await Marketplace.deploy(nftContractAddress, commissionPercent);

        console.log("Deployment transaction hash:", marketplace.deployTransaction.hash);
        console.log("Waiting for deployment confirmation...");

        await marketplace.deployed();

        console.log("Marketplace deployed to:", marketplace.address);
    } catch (error) {
        console.error("Error during deployment:", error);
        throw error;
    }
}

main()
    .then(() => process.exit(0))
    .catch((error) => {
        console.error("Deployment failed:", error);
        process.exit(1);
    });
