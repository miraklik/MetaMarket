const { expect } = require("chai");
const { ethers } = require("hardhat");
const { solidity } = require("chai");

chai.use(solidity);

describe("Payments", function () {
  let acc1;
  let acc2;
  let payments;

  beforeEach(async function () {
    [acc1, acc2] = await ethers.getSigners();
    const Payments = await ethers.getContractFactory("Payments");
    payments = await Payments.deploy();
    await payments.deployed();  // убедитесь, что контракт развёрнут
  });

  it("should be deployed", async function () {
    expect(payments.address).to.be.properAddress;
  });

  it("should have 0 ether by default", async function () {
    const balance = await ethers.provider.getBalance(payments.address);  // используйте стандартный метод
    expect(balance).to.eq(0);
  });

  it("should be possible to send funds", async function () {
    const sum = ethers.utils.parseEther("0.1"); // Отправляем 0.1 Ether
    const msg = "hello from hardhat";
    const tx = await payments.connect(acc2).pay(msg, { value: sum });

    await expect(() => tx).to.changeEtherBalances([acc2, payments], [-sum, sum]);  // проверяем изменения балансов
    await tx.wait();

    const newPayment = await payments.getPayment(acc2.address, 0);  // проверяем детали платежа
    expect(newPayment.message).to.eq(msg);
    expect(newPayment.amount).to.eq(sum);
    expect(newPayment.from).to.eq(acc2.address);
  });
});
