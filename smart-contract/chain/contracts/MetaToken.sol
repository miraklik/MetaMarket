// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.20;

contract MonToken {
    string public constant name = "MetaToken";

    string public constant symbol = "META";

    uint256 public constant decimals = 18;

    uint256 public totalSupply;

    address public owner;

    mapping(address => uint256) public balanceOf;

    mapping(address => mapping(address => uint256)) public allowance;

    event Transfer(address indexed _from, address indexed _to, uint256 _value);

    event Approval(
        address indexed _owner,
        address indexed _spender,
        uint256 _value
    );

    event Ownership(address indexed owner, address indexed ownerNew);

    constructor(uint256 _totalSupply) {
        owner = msg.sender;
        totalSupply = _totalSupply;
        balanceOf[msg.sender] = totalSupply;
    }

    modifier onlyOwner() {
        require(msg.sender == owner, "pas autorise");
        _;
    }

    function transfer(address _to, uint256 _value)
        public
        returns (bool success)
    {
        require(_to != address(0), "sets a normal address");

        require(balanceOf[msg.sender] >= _value, "you don't have enough funds");

        balanceOf[msg.sender] -= _value;

        balanceOf[_to] += _value;

        emit Transfer(msg.sender, _to, _value);

        return true;
    }

    function approve(address _spender, uint256 _value)
        public
        returns (bool succes)
    {
        allowance[msg.sender][_spender] = _value;

        emit Approval(msg.sender, _spender, _value);

        return true;
    }

    function transferFrom(
        address _from,
        address _to,
        uint256 _value
    ) public returns (bool success) {
        require(
            balanceOf[_from] >= _value,
            "we don't have enough tokens to serve you"
        );

        require(allowance[_from][msg.sender] >= _value, "not authorized");

        allowance[_from][msg.sender] -= _value;

        balanceOf[_from] -= _value;

        balanceOf[_to] += _value;

        emit Transfer(msg.sender, _to, _value);

        return true;
    }

    function mint(address _to, uint256 _value)
        public
        onlyOwner
        returns (bool success)
    {
        require(_to != address(0));

        totalSupply += _value;

        balanceOf[_to] += _value;

        emit Transfer(msg.sender, _to, _value);

        return true;
    }

    function burn(uint256 _value) public onlyOwner returns (bool success) {
        totalSupply -= _value;

        balanceOf[msg.sender] -= _value;

        emit Transfer(msg.sender, address(0), _value);

        return true;
    }

    function transferOwnerShip(address _newOwner) public onlyOwner {
        owner = _newOwner;

        emit Ownership(msg.sender, _newOwner);
    }
}