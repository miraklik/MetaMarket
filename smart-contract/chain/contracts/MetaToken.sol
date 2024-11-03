// SPDX-License-Identifier: MIT

pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/utils/Base64.sol";

contract BadgeToken is ERC721 {
    uint265 private _currentTokenId = 0;

    constructor(string _name, string _symbol) ERC721(_name, _symbol) {
    }

    function mintTo(address to) public returns (uint256) {
        uint256 newTokenId = _getNextTokenId;
        _mint(to, newTokenId);
        return newTokenId;
    }

    function _getNextTokenId() private view returns (uint256) {
        return _currentTokenId+1;
    }

    function _incrementTokenId() private {
        _currentTokenId++;
    }

     function tokenURI(uint256 tokenId) override public pure returns (string memory) {
        string[3] memory parts;

        parts[0] = "<svg xmlns='http://www.w3.org/2000/svg' preserveAspectRatio='xMinYMin meet' viewBox='0 0 350 350'><style>.base { fill: white; font-family: serif; font-size: 300px; }</style><rect width='100%' height='100%' fill='brown' /><text x='100' y='260' class='base'>";

        parts[1] = Strings.toString(tokenId);

        parts[2] = "</text></svg>";

        string memory json = Base64.encode(bytes(string(abi.encodePacked(
            "{\"name\":\"Badge #", 
            Strings.toString(tokenId), 
            "\",\"description\":\"Badge NFT with on-chain SVG image.\",",
            "\"image\": \"data:image/svg+xml;base64,", 
            // Base64.encode(bytes(output)), 
            Base64.encode(bytes(abi.encodePacked(parts[0], parts[1], parts[2]))),     
            "\"}"
            ))));

        return string(abi.encodePacked("data:application/json;base64,", json));
    }    

}