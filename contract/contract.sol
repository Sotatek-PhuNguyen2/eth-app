// SPDX-License-Identifier: MIT
pragma solidity ^0.8.17;

contract ReceiveEther {
    /*
    Which function is called, fallback() or receive()?

           send Ether
               |
         msg.data is empty?
              / \
            yes  no
            /     \
receive() exists?  fallback()
         /   \
        yes   no
        /      \
    receive()   fallback()
    */

    // Function to receive Ether. msg.data must be empty
    receive() external payable {}

    // Fallback function is called when msg.data is not empty
    fallback() external payable {}

    function getBalance() public view returns (uint) {
        return address(this).balance;
    }
}

contract Store {

    address admin;
    mapping (address => uint)  ownerRegisterCount;

    constructor (){
        admin = msg.sender;
        ownerRegisterCount[admin] = 2 ;
    }

    event Unsubscribe(address direction);
    struct User{
        string username;
        address direction;
    }

    // function sendMeMoneyContract() public payable{

    // }

    // function sendAdminMoneyContract () public  {

    // }

    // function refundMoneyUser() private{

    // }


    User[] users;

    function register (string memory _username) public payable{
        users.push(User(_username, msg.sender));
        require (ownerRegisterCount[msg.sender]==0,"Register failed");
        ownerRegisterCount[msg.sender]++;
        require(msg.value >= 0.00001 ether, "Failed to send Ether");
    }

    function list() public view returns (User[] memory){
        //require (ownerRegisterCount[msg.sender] == 2);
        return users;
    }

    function unsubscribe() public{
        uint id;
        require (ownerRegisterCount[msg.sender] == 1,"Unsubscribe failed");
        for (uint i = 0; i < users.length; i++){
            if (users[i].direction == msg.sender){
               id = i;
            }
        }
        for (uint i = id; i < users.length - 1; i++){
            users[i] = users[i+1];
        }
        users.pop();
        ownerRegisterCount[msg.sender] = 3;
        emit Unsubscribe(msg.sender);
    }

    function refundMoneyUser()public payable {
        require(ownerRegisterCount[msg.sender] == 3,"Failed to refund");
        (bool sent,) = msg.sender.call{value: msg.value}("");
        require(sent, "Failed to send Ether");
        ownerRegisterCount[msg.sender] == 2;
    }
}