require("@nomiclabs/hardhat-waffle");
require("dotenv").config();

module.exports = {
  solidity: {
    version: "0.8.20", // Везде используйте версию 0.8.20
    settings: {
      optimizer: {
        enabled: true,
        runs: 200,
      },
    },
  },
  networks: {
    bsc: {
      url: "https://bscrpc.com/",
      accounts: [process.env.PRIVATE_KEY],
    },
    polygon: {
      url: "https://polygon-rpc.com/",
      accounts: [process.env.PRIVATE_KEY],
    },
    bscTestnet: {
      url: "https://data-seed-prebsc-1-s1.binance.org:8545/",
      accounts: [process.env.PRIVATE_KEY],
    },
    polygonTestnet: {
      url: "https://rpc-mumbai.maticvigil.com",
      accounts: [process.env.PRIVATE_KEY],
    },
    ethereum: {
      url: "https://eth.public-rpc.com/",
      accounts: [process.env.PRIVATE_KEY],
    },
    ethereumTestnet: {
      url: "https://mainnet.infura.io/v3/dfa0335a8d2b4364bd669159aa3dc734",
      accounts: [process.env.PRIVATE_KEY],
    },
  },
};
