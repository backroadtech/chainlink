pragma solidity >=0.5.0;

interface AggregatorV2Interface {
  function decimals() external view returns (uint8);
  function description() external view returns (string memory);
  function getRoundData(uint256 _roundId)
    external
    view
    returns (
      uint256 roundId,
      int256 answer,
      uint256 startedAt,
      uint256 updatedAt,
      uint256 answeredInRound
    );
  function latestRoundData()
    external
    view
    returns (
      uint256 roundId,
      int256 answer,
      uint256 startedAt,
      uint256 updatedAt,
      uint256 answeredInRound
    );
  function version() external view returns (uint256);
}
