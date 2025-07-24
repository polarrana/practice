// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract BinarySearch {
    /**
     * @dev 在已排序的数组中执行二分查找
     * @param arr 已排序的数组（升序）
     * @param target 要查找的目标值
     * @return 找到则返回索引，未找到则返回type(uint256).max
     */
    function search(uint256[] memory arr, uint256 target) public pure returns (uint256) {
        // 检查数组是否为空
        if (arr.length == 0) {
            return type(uint256).max;
        }

        uint256 left = 0;
        uint256 right = arr.length - 1;

        while (left <= right) {
            // 防止大数组相加溢出
            uint256 mid = left + (right - left) / 2;
            
            if (arr[mid] == target) {
                return mid;
            } else if (arr[mid] < target) {
                left = mid + 1;
            } else {
                // 防止下溢
                if (mid == 0) {
                    break;
                }
                right = mid - 1;
            }
        }

        return type(uint256).max; // 表示未找到
    }
}