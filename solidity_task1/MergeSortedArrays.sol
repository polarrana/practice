// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MergeSortedArrays {
    // 合并两个升序排列的数组
    function merge(uint256[] memory nums1, uint256[] memory nums2) public pure returns (uint256[] memory) {
        uint256 m = nums1.length;
        uint256 n = nums2.length;
        uint256[] memory merged = new uint256[](m + n);
        
        uint256 i = 0; // nums1的指针
        uint256 j = 0; // nums2的指针
        uint256 k = 0; // merged数组的指针
        
        // 比较两个数组的元素并按顺序合并
        while (i < m && j < n) {
            if (nums1[i] <= nums2[j]) {
                merged[k] = nums1[i];
                i++;
            } else {
                merged[k] = nums2[j];
                j++;
            }
            k++;
        }
        
        // 将nums1剩余元素复制到merged
        while (i < m) {
            merged[k] = nums1[i];
            i++;
            k++;
        }
        
        // 将nums2剩余元素复制到merged
        while (j < n) {
            merged[k] = nums2[j];
            j++;
            k++;
        }
        
        return merged;
    }
}