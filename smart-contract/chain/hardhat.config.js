require("@nomicfoundation/hardhat-toolbox");
require("@typechain/hardhat");

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: {
    compilers: [
      { version: "0.8.20" },
      { version: "0.8.27"},
      { version: "0.8.3" },
    ],
  },
  typechain: {
    outDir: "typechain",
    target: "ethers-v5",
  },
};

