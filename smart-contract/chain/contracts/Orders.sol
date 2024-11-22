// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.18;

import "@openzeppelin/contracts/token/ERC721/IERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/IERC721Enumerable.sol";

/**
 * @title NFT Marketplace
 * @dev This contract facilitates the listing and sale of ERC721 tokens with added enumerable support.
 */
contract Marketplace {
    /// @notice Represents a listing on the marketplace.
    struct Listing {
        uint256 id;
        address seller;
        uint256 tokenId;
        uint256 price;
        bool sold;
    }

    /// @notice Owner of the marketplace.
    address public owner;

    /// @notice Total number of listings created.
    uint256 public listingCount;

    /// @notice Commission percentage charged by the marketplace.
    uint256 public commissionPercent;

    /// @dev Stores all listings by their unique ID.
    mapping(uint256 => Listing) public listings;

    /// @notice Interface for the ERC721 token contract.
    IERC721 public nftContract;

    /// @notice Interface for enumerable ERC721 functionality.
    IERC721Enumerable public nftEnumerable;

    /// @notice Event emitted when a new listing is created.
    event ListingCreated(
        uint256 indexed id,
        address indexed seller,
        uint256 tokenId,
        uint256 price
    );

    /// @notice Event emitted when an NFT is purchased.
    event PurchaseCompleted(
        uint256 indexed id,
        address indexed buyer,
        uint256 tokenId,
        uint256 price
    );

    /// @notice Event emitted when a listing is cancelled.
    event ListingCancelled(uint256 indexed id, address indexed seller);

    /// @notice Event emitted when the marketplace commission is updated.
    event CommissionUpdated(uint256 newPercent);

    /// @notice Event emitted when funds are withdrawn from the contract.
    event FundsWithdrawn(uint256 amount, address indexed owner);

    /// @dev Modifier to restrict certain actions to the contract owner.
    modifier onlyOwner() {
        require(msg.sender == owner, "Only the owner can perform this action");
        _;
    }

    /// @dev Reentrancy guard modifier.
    uint256 private unlocked = 1;
    modifier nonReentrant() {
        require(unlocked == 1, "Reentrant call");
        unlocked = 0;
        _;
        unlocked = 1;
    }

    /**
     * @dev Initializes the marketplace with an ERC721 contract and commission percentage.
     * @param _nftContractAddress Address of the ERC721 token contract.
     * @param _commissionPercent Marketplace commission percentage.
     */
    constructor(address _nftContractAddress, uint256 _commissionPercent) {
        require(_commissionPercent <= 100, "Commission cannot exceed 100%");
        owner = msg.sender;
        nftContract = IERC721(_nftContractAddress);
        nftEnumerable = IERC721Enumerable(_nftContractAddress);
        commissionPercent = _commissionPercent;
    }

    /**
     * @notice Creates a listing to sell an NFT.
     * @param _tokenId ID of the token to sell.
     * @param _price Sale price in wei.
     */
    function createListing(uint256 _tokenId, uint256 _price) external {
        require(_price > 0, "Price must be greater than 0");
        require(
            nftContract.ownerOf(_tokenId) == msg.sender,
            "You are not the owner of this token"
        );
        require(
            nftContract.getApproved(_tokenId) == address(this),
            "Marketplace is not approved to transfer this token"
        );

        listingCount++;
        listings[listingCount] = Listing(
            listingCount,
            msg.sender,
            _tokenId,
            _price,
            false
        );

        emit ListingCreated(listingCount, msg.sender, _tokenId, _price);
    }

    /**
     * @notice Purchases an NFT from an active listing.
     * @param _listingId ID of the listing to purchase.
     */
    function purchaseListing(uint256 _listingId) external payable nonReentrant {
        Listing storage listing = listings[_listingId];
        require(_listingId > 0 && _listingId <= listingCount, "Invalid listing ID");
        require(!listing.sold, "Item already sold");
        require(msg.value == listing.price, "Incorrect payment amount");

        listing.sold = true;

        uint256 commissionAmount = (listing.price * commissionPercent) / 100;
        uint256 sellerAmount = listing.price - commissionAmount;

        // Transfer commission to the marketplace owner
        (bool commissionSent, ) = owner.call{value: commissionAmount}("");
        require(commissionSent, "Failed to send commission");

        // Transfer payment to the seller
        (bool sellerPaid, ) = listing.seller.call{value: sellerAmount}("");
        require(sellerPaid, "Failed to send payment to seller");

        // Transfer NFT to the buyer
        nftContract.safeTransferFrom(listing.seller, msg.sender, listing.tokenId);

        emit PurchaseCompleted(_listingId, msg.sender, listing.tokenId, listing.price);
    }

    /**
     * @notice Cancels an active listing.
     * @param _listingId ID of the listing to cancel.
     */
    function cancelListing(uint256 _listingId) external {
        Listing storage listing = listings[_listingId];
        require(_listingId > 0 && _listingId <= listingCount, "Invalid listing ID");
        require(listing.seller == msg.sender, "You are not the seller");
        require(!listing.sold, "Item already sold");

        delete listings[_listingId];

        emit ListingCancelled(_listingId, msg.sender);
    }

    /**
     * @notice Updates the marketplace commission percentage.
     * @param _newPercent New commission percentage.
     */
    function setCommissionPercent(uint256 _newPercent) external onlyOwner {
        require(_newPercent <= 100, "Commission cannot exceed 100%");
        commissionPercent = _newPercent;

        emit CommissionUpdated(_newPercent);
    }

    /**
     * @notice Withdraws all funds from the contract.
     */
    function withdrawFunds() external onlyOwner nonReentrant {
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");

        (bool success, ) = owner.call{value: balance}("");
        require(success, "Withdraw failed");

        emit FundsWithdrawn(balance, owner);
    }

    /**
    * @notice Returns the token ID owned by `_owner` at a specific `index`.
    * @param _owner Address of the token owner.
    * @param index Index in the list of the `_owner`'s tokens.
    * @return tokenId The ID of the token owned by the `_owner` at the given index.
    */
    function tokenOfOwnerByIndex(address _owner, uint256 index) public view returns (uint256 tokenId) {
        return nftEnumerable.tokenOfOwnerByIndex(_owner, index);
    }
}
