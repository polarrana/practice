// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
// 合约地址0xBbc4632fcE69956fDDe870Ffb976E3DBDf6316F3
contract BeggingContract {
    // 合约所有者
    address payable public owner;
    
    // 记录每个地址的捐赠金额
    mapping(address => uint256) public donations;

    // 捐赠者信息结构体
    struct DonorInfo {
        address donor;
        uint256 amount;
    }

    // 记录所有捐赠者
    DonorInfo[] private allDonors;
    
    // 捐赠事件
    event Donation(address indexed donor, uint256 amount);
    
    // 构造函数，设置合约所有者
    constructor() {
        owner = payable(msg.sender);
    }
    
    // 修饰符：仅所有者可调用
    modifier onlyOwner() {
        require(msg.sender == owner, "Only owner can call this function");
        _;
    }

    // 修改 donate 函数，添加捐赠者到列表
    function donate() external payable {
        require(msg.value > 0, "Donation amount must be greater than 0");
        
        donations[msg.sender] += msg.value;
        emit Donation(msg.sender, msg.value);
        
        // 添加到捐赠者列表或更新现有记录
        bool found = false;
        for (uint i = 0; i < allDonors.length; i++) {
            if (allDonors[i].donor == msg.sender) {
                allDonors[i].amount = donations[msg.sender];
                found = true;
                break;
            }
        }
        
        if (!found) {
            allDonors.push(DonorInfo(msg.sender, msg.value));
        }
    }

    // 改进的排行榜函数
    function getTopDonors() external view returns (DonorInfo[] memory) {
        // 复制数组进行排序（注意：链上排序可能消耗大量gas）
        DonorInfo[] memory sortedDonors = new DonorInfo[](allDonors.length);
        for (uint i = 0; i < allDonors.length; i++) {
            sortedDonors[i] = allDonors[i];
        }
        
        // 简单的冒泡排序（实际项目中应考虑更高效的算法或离链处理）
        for (uint i = 0; i < sortedDonors.length; i++) {
            for (uint j = i+1; j < sortedDonors.length; j++) {
                if (sortedDonors[i].amount < sortedDonors[j].amount) {
                    DonorInfo memory temp = sortedDonors[i];
                    sortedDonors[i] = sortedDonors[j];
                    sortedDonors[j] = temp;
                }
            }
        }
        
        // 返回前3名
        DonorInfo[] memory topDonors = new DonorInfo[](3);
        for (uint i = 0; i < 3; i++) {
            topDonors[i] = sortedDonors[i];
        }
        
        return topDonors;
    }
    
    // 提取合约资金（仅所有者）
    function withdraw() external onlyOwner {
        uint256 balance = address(this).balance;
        require(balance > 0, "No funds to withdraw");
        
        // 转移资金给所有者
        owner.transfer(balance);
    }
    
    // 查询指定地址的捐赠金额
    function getDonation(address _donor) external view returns (uint256) {
        return donations[_donor];
    }
    
    // 获取合约余额
    function getContractBalance() external view returns (uint256) {
        return address(this).balance;
    }
}