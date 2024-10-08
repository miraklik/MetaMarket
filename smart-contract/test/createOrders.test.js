const { expect } = require("chai")
const { expect } = require("hardhat")

describe("Orders", function () {
    let acc1
    let acc2
    let oredersdapl

    beforeEach(async function () {
        [acc1, acc2] = await ethers.getSigners()
        const Orders = await ethers.getContractFactory("Orders", acc1)
        oredersdapl = await Orders.deploy()
        await oredersdapl.deployed()
        console.log(oredersdapl.address)
    })

    it("should be deployed", async function () {
        console.log("success!")
    })
})