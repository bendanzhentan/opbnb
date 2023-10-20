pragma solidity 0.8.15;

import { L2StandardBridge } from "./L2StandardBridge.sol";

contract L2DelegateBridge {
    address public constant L2_STANDARD_BRIDGE_ADDRESS = 0x4200000000000000000000000000000000000010;
    L2StandardBridge public L2_STANDARD_BRIDGE = L2StandardBridge(payable(L2_STANDARD_BRIDGE_ADDRESS));

    function withdrawTo(
        address _l2Token,
        address _to,
        uint256 _amount,
        uint32 _minGasLimit,
        bytes calldata _extraData
    ) external payable virtual {
        L2_STANDARD_BRIDGE.withdrawTo{
            value: msg.value
        }(_l2Token, _to, _amount, _minGasLimit, _extraData);
    }
}
