// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

// 创建一个名为Voting的合约，包含以下功能：
// 一个mapping来存储候选人的得票数
// 一个vote函数，允许用户投票给某个候选人
// 一个getVotes函数，返回某个候选人的得票数
// 一个resetVotes函数，重置所有候选人的得票数
contract Voting {
    // 存储候选人得票数
    mapping(string => uint256) private votesReceived;
    
    // 存储所有候选人列表
    string[] private candidates;
    
    // 标记候选人是否存在
    mapping(string => bool) private candidateExists;
    
    // 合约所有者地址
    address public owner;
    
    // 事件定义
    event Voted(address indexed voter, string candidate);
    event CandidateAdded(string candidate);
    event VotesReset(address indexed resetBy);
    
    // 构造函数，设置合约部署者为所有者
    constructor() {
        owner = msg.sender;
    }
    
    // 修饰器：限制只有所有者能执行
    modifier onlyOwner() {
        require(msg.sender == owner, "Only contract owner can perform this action");
        _;
    }
    
    /**
     * @dev 添加新候选人
     * @param candidate 候选人名称
     */
    function addCandidate(string memory candidate) public onlyOwner {
        require(bytes(candidate).length > 0, "Candidate name cannot be empty");
        require(!candidateExists[candidate], "Candidate already exists");
        
        candidates.push(candidate);
        votesReceived[candidate] = 0; // 显式初始化为0
        candidateExists[candidate] = true; // 标记为已存在
        emit CandidateAdded(candidate);
    }
    
    /**
     * @dev 给指定候选人投票
     * @param candidate 候选人名称
     */
    function vote(string memory candidate) public {
        require(bytes(candidate).length > 0, "Candidate name cannot be empty");
        require(candidateExists[candidate], "Candidate does not exist");
        
        votesReceived[candidate] += 1;
        emit Voted(msg.sender, candidate);
    }
    
    /**
     * @dev 获取指定候选人的得票数
     * @param candidate 候选人名称
     * @return 该候选人的得票数
     */
    function getVotes(string memory candidate) public view returns (uint256) {
        return votesReceived[candidate];
    }
    
    /**
     * @dev 获取所有候选人列表
     * @return 包含所有候选人的数组
     */
    function getAllCandidates() public view returns (string[] memory) {
        return candidates;
    }
    
    /**
     * @dev 重置所有候选人的得票数（仅所有者可调用）
     */
    function resetVotes() public onlyOwner {
        for(uint i = 0; i < candidates.length; i++) {
            votesReceived[candidates[i]] = 0;
        }
        emit VotesReset(msg.sender);
    }
    
    /**
     * @dev 获取候选人数量
     * @return 候选人总数
     */
    function getCandidateCount() public view returns (uint256) {
        return candidates.length;
    }
    
    /**
     * @dev 检查候选人是否存在
     * @param candidate 候选人名称
     * @return 是否存在
     */
    function isCandidateExist(string memory candidate) public view returns (bool) {
        return candidateExists[candidate];
    }
}