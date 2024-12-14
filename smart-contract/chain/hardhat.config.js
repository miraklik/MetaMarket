require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-ethers");
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
      url: process.env.APIKEY_BSCMAINNET,
      accounts: [process.env.PRIVATE_KEY],
    },
    polygon: { 
      url: process.env.APIKEY_POLMAINNET,
      accounts: [process.env.PRIVATE_KEY],
    },
    bscTestnet: {
      url: process.env.APIKEY_BSCTESTNET,
      accounts: [process.env.PRIVATE_KEY],
    },
    polygonTestnet: {
      url: process.env.APIKEY_POLAMOY,
      accounts: [process.env.PRIVATE_KEY],
    },
    ethereum: {
      url: process.env.APIKEY_EHTMAINNET,
      accounts: [process.env.PRIVATE_KEY],
    },
    sepolia: {
      url: process.env.APIKEY_EHTSEPOLIA,
      accounts: [process.env.PRIVATE_KEY],
    },
    starknet: {
      url: process.env.APIKEY_STRKMAINNET,
      accounts: [process.env.PRIVATE_KEY]
    },
    arbitrum: {
      url: process.env.APIKEY_ARBMAINNET,
      accounts: [process.env.PRIVATE_KEY]
    },
    optimism: {
      url: process.env.APIKEY_OPMAINNET,
      accounts: [process.env.PRIVATE_KEY]
    },
    opBNB: {
      url: process.env.APIKEY_OPBNBMAINNET,
      accounts: [process.env.PRIVATE_KEY]
    },
    arbitrumSepolia: {
      url: process.env.APIKEY_ARBSEPOLIA,
      accounts: [process.env.PRIVATE_KEY]
    },
    optimismSepolia: {
      url: process.env.APIKEY_OPSEPOLIA,
      accounts: [process.env.PRIVATE_KEY]
    },
    starknetSepolia: {
      url: process.env.APIKEY_STRKSEPOLIA,
      accounts: [process.env.PRIVATE_KEY]
    },
    mantle: {
      url: process.env.APIKEY_MANTELMAINNET,
      accounts: [process.env.PRIVATE_KEY]
    }
  },
};
