// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract IntegerToRoman {
    // 定义数值到罗马数字的映射
    struct RomanNumeral {
        uint256 value;
        string symbol;
    }
    
    RomanNumeral[] private romanNumerals;
    
    constructor() {
        // 初始化映射表（按从大到小顺序排列）
        romanNumerals.push(RomanNumeral(1000, "M"));
        romanNumerals.push(RomanNumeral(900, "CM"));
        romanNumerals.push(RomanNumeral(500, "D"));
        romanNumerals.push(RomanNumeral(400, "CD"));
        romanNumerals.push(RomanNumeral(100, "C"));
        romanNumerals.push(RomanNumeral(90, "XC"));
        romanNumerals.push(RomanNumeral(50, "L"));
        romanNumerals.push(RomanNumeral(40, "XL"));
        romanNumerals.push(RomanNumeral(10, "X"));
        romanNumerals.push(RomanNumeral(9, "IX"));
        romanNumerals.push(RomanNumeral(5, "V"));
        romanNumerals.push(RomanNumeral(4, "IV"));
        romanNumerals.push(RomanNumeral(1, "I"));
    }
    
    // 整数转罗马数字的主函数
    function intToRoman(uint256 num) public view returns (string memory) {
        require(num > 0 && num < 4000, "Number must be between 1 and 3999");
        
        bytes memory roman;
        
        for (uint256 i = 0; i < romanNumerals.length; i++) {
            while (num >= romanNumerals[i].value) {
                roman = abi.encodePacked(roman, romanNumerals[i].symbol);
                num -= romanNumerals[i].value;
            }
        }
        
        return string(roman);
    }
}