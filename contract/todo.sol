// SPDX-License-Identifier: UNLICENSED
pragma solidity >=0.7.0 <0.9.0;

contract Todo {

    Task[] task;

    address public owner;

    struct Task {
        bool m_status;
        string m_data;
    }

    constructor(){
        owner = msg.sender;
    }

    modifier isOwner() {
        require(owner == msg.sender);
        _;
    }

    function add(string memory data) public isOwner {
        task.push(Task(false, data));
    }

    function get(uint id) public isOwner view returns (Task memory) {
        return task[id];
    }

    function list() public isOwner view returns (Task[] memory) {
        return task;
    }

    function update(uint id, string memory data) public isOwner {
        task[id].m_data = data;
    }

    function remove(uint id) public isOwner {

        for (uint i = id; i < task.length - 1; i++) {
            task[i] = task[i + 1];
        }
    
        task.pop();
    }
}