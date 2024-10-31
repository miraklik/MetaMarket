async function main() {
    const network = hre.network.name;
  
    let usdtAddress;
    if (network === "bsc") {
      usdtAddress = process.env.USDT_ADDRESS_BSC;
    } else if (network === "polygon") {
      usdtAddress = process.env.USDT_ADDRESS_POLYGON;
    }else if (network == "ethereum"){
        usdtAddress = process.env.USDT_ADDRESS_ETHEREUM
    }
  
    const Marketplace = await ethers.getContractFactory("Marketplace");
    const commissionPercent = 5; // Укажите процент комиссии
  
    // Разворачиваем контракт с указанием адреса USDT и процента комиссии
    const marketplace = await Marketplace.deploy(usdtAddress, commissionPercent);
  
    await marketplace.deployed();
  
    console.log(`Marketplace deployed to ${network} at address:`, marketplace.address);
  }
  
  main()
    .then(() => process.exit(0))
    .catch((error) => {
      console.error(error);
      process.exit(1);
    });
  