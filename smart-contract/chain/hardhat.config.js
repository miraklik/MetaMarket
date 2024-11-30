require("@nomiclabs/hardhat-waffle");
require("@nomiclabs/hardhat-ethers");
require("dotenv").config();

const { APIKEY_EHTMAINNET, APIKEY_EHTSEPOLIA, APIKEY_BSCTESTNET, APIKEY_POLMAINNET, APIKEY_ARBSEPOLIA, APIKEY_STRKSEPOLIA, APIKEY_ARBMAINNET,APIKEY_OPMAINNET, APIKEY_OPSEPOLIA, APIKEY_STRKMAINNET, APIKEY_POLAMOY, APIKEY_BSCMAINNET, APIKEY_OPBNBMAINNET, APIKEY_MANTELMAINNET, PRIVATE_KEY } = process.env; 
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
      url: APIKEY_BSCMAINNET,
      accounts: [`0x${PRIVATE_KEY}`],
    },
    polygon: {
      url: APIKEY_POLMAINNET,
      accounts: [`0x${PRIVATE_KEY}`],
    },
    bscTestnet: {
      url: APIKEY_BSCTESTNET,
      accounts: [`0x${PRIVATE_KEY}`],
    },
    polygonTestnet: {
      url: APIKEY_POLAMOY,
      accounts: [`0x${PRIVATE_KEY}`],
    },
    ethereum: {
      url: APIKEY_EHTMAINNET,
      accounts: [`0x${PRIVATE_KEY}`],
    },
    ethereumSepolia: {
      url: APIKEY_EHTSEPOLIA,
      accounts: [`0x${PRIVATE_KEY}`],
    },
    starknet: {
      url: APIKEY_STRKMAINNET,
      accounts: [`0x${PRIVATE_KEY}`]
    },
    arbitrum: {
      url: APIKEY_ARBMAINNET,
      accounts: [`0x${PRIVATE_KEY}`]
    },
    optimism: {
      url: APIKEY_OPMAINNET,
      accounts: [`0x${PRIVATE_KEY}`]
    },
    opBNB: {
      url: APIKEY_OPBNBMAINNET,
      accounts: [`0x${PRIVATE_KEY}`]
    },
    arbitrumSepolia: {
      url: APIKEY_ARBSEPOLIA,
      accounts: [`0x${PRIVATE_KEY}`]
    },
    optimismSepolia: {
      url: APIKEY_OPSEPOLIA,
      accounts: [`0x${PRIVATE_KEY}`]
    },
    starknetSepolia: {
      url: APIKEY_STRKSEPOLIA,
      accounts: [`0x${PRIVATE_KEY}`]
    },
    mantle: {
      url: APIKEY_MANTELMAINNET,
      accounts: [`0x${PRIVATE_KEY}`]
    }
  },
};
