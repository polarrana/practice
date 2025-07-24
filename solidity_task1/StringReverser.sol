// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract StringReverser {
    /**
     * @dev 反转字符串
     * @param _str 要反转的字符串
     * @return 反转后的字符串
     */
    function reverse(string memory _str) public pure returns (string memory) {
        // 将字符串转换为bytes类型以便处理
        bytes memory strBytes = bytes(_str);
        uint256 length = strBytes.length;
        
        // 如果字符串为空或长度为1，直接返回
        if (length == 0 || length == 1) {
            return _str;
        }
        
        // 创建新的bytes数组存储反转结果
        bytes memory reversed = new bytes(length);
        
        // 反转操作
        for (uint256 i = 0; i < length; i++) {
            reversed[i] = strBytes[length - 1 - i];
        }
        
        // 将bytes转换回string类型返回
        return string(reversed);
    }
}