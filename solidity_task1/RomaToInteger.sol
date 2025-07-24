// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RomanToInteger {
    // 罗马数字到整数的映射
    mapping(bytes1 => uint256) private romanValues;
    
    constructor() {
        // 初始化罗马数字映射
        romanValues['I'] = 1;
        romanValues['V'] = 5;
        romanValues['X'] = 10;
        romanValues['L'] = 50;
        romanValues['C'] = 100;
        romanValues['D'] = 500;
        romanValues['M'] = 1000;
    }
    
    // 罗马数字转整数的主函数
    function romanToInt(string memory s) public view returns (uint256) {
        bytes memory roman = bytes(s);
        uint256 length = roman.length;
        uint256 total = 0;
        
        for (uint256 i = 0; i < length; i++) {
            uint256 current = romanValues[roman[i]];
            
            // 如果当前字符比下一个字符小，则减去当前值
            if (i < length - 1 && current < romanValues[roman[i + 1]]) {
                total -= current;
            } else {
                total += current;
            }
        }
        
        return total;
    }
}