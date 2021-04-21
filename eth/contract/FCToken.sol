// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

contract FCToken {

  string public name;
  string public symbol;
  uint8 public decimals;
  uint256 public totalSupply;
  address public owner;

  mapping (address => uint256) public balanceOf;
  mapping (address => mapping(address => uint256)) public allowance;
  mapping(address => uint8) public signer;
  mapping(address => uint8) public allowers;
  address[] public accounts;
  string[] public nodes;
  mapping(address=>bool) public accountsMap;

  event Transfer(address indexed  from, address indexed  to, uint256 value);
  event Approval(address indexed  from, address indexed  delegatee, uint256 value);
//   event ApproveWithArray(address indexed  from, address[] indexed  delegatee, uint256 value);
  struct Transcation {
    address from;
    address to;
    uint256 amount;
    uint8 signatureCounts;
    mapping(address => uint8) signatures;
  }

  uint256 constant MIN_SIGNATURE = 2;
  uint8[] private pendingTranscation;
  uint8 public transcationId;
  mapping(uint256 => Transcation) public transcations;

  event CreateTranscation(
    address from,
    address to,
    uint256 amount,
    uint8 transcationId
  );

  constructor(string memory _name, string memory _symbol, uint8 _decimals, uint256 _initiSuppy){
    name = _name;
    symbol = _symbol;
    decimals = _decimals;
    totalSupply = _initiSuppy * 10 ** uint256(decimals);
    balanceOf[msg.sender] = totalSupply;
    owner = msg.sender;

  }

  function _transfer(address _from, address _to, uint256 _value) internal {
    require(msg.sender!=owner);
    require(balanceOf[_from] >= _value);
    require(_to != address(0x0));
    require(balanceOf[_to] + _value > balanceOf[_to]);

    uint256 previousBanlance = balanceOf[_from] + balanceOf[_to];
    balanceOf[_from] -= _value;
    balanceOf[_to] += _value;
    emit Transfer(_from, _to, _value);
    assert (balanceOf[_from] + balanceOf[_to] == previousBanlance);
  }

  function transfer(address _to, uint256 _value) public {

    _transfer(msg.sender, _to, _value);

  }


  function transferFrom(address _from, address _to, uint256 _value) public {
    require(_to != address(0x0));
    require(allowance[owner][msg.sender] >= _value); //授权Token必须大于等于value
    allowance[_from][msg.sender] -= _value;
    balanceOf[_from] -= _value;
    balanceOf[_to] += _value;
  }

  function createMiner(address miner,string memory nodeId)public {
      require(miner!=address(0x0));
      bool isEx = accountsMap[miner];
      if(isEx != true){
          nodes.push(nodeId);
          accounts.push(miner);
          accountsMap[miner] = true;
      }

  }


  function approve(address _delegatee, uint256 _value) public {
    require(allowers[msg.sender] == 1);
    require(_delegatee != address(0x0));
    require(balanceOf[owner] >= _value);
    allowance[owner][_delegatee] += _value;
    emit Approval(msg.sender, _delegatee, _value);
  }

  function approveWithArray(uint8[] memory b) public {
    require(b.length>0);
    require(allowers[msg.sender] == 1);

    uint len = b.length / 8;
  if (b.length % 8 > 0) {
      len++;
  }
  uint[] memory a = new uint[](len);
  uint c;
  for(uint i = b.length -1;i>=0;i--){
      uint8 v=b[i];
      for ( uint8 j = 0; j < 8; j++) {
         if (v&(1<<j) > 0) {
            a[c] = (b.length-i-1)*8+j;
            c++;
         }
      }
  }

    int size = 10;
    if(a.length<10){
        size = int(a.length);
    }

    uint256[] memory arrWard = new uint256[](uint(size));

    for(int i=0;i<size;i++){
         uint256 random = uint256(keccak256(abi.encodePacked(block.difficulty, block.timestamp)));
         uint256 num = random%a.length;
         uint index = a[num];
         bool isHave = false;
         for(uint j=0;j<arrWard.length;j++){
             uint256 n = arrWard[j];
             if(n ==index+1&&n!=0){
                 isHave = true;
             }
         }
         if(isHave == false){
             arrWard[uint256(i)] = index+1;
             address who = accounts[index-1];
             allowance[owner][who] += 1;
         }


    }



  }


    modifier onlyOwner() {
     require(owner == msg.sender);
     _;
    }

    modifier onlySigner() {
     require(owner == msg.sender || signer[msg.sender] == 1);
     _;
    }

    function setAllower(address s) public onlyOwner {

     allowers[s] = 1;
    }

  function removeAllower(address s) public onlyOwner {

      allowers[s] = 0;
      delete allowers[s];
  }

  function setSigner(address s) public onlyOwner {

     signer[s] = 1;
  }

  function removeSigner(address s) public onlyOwner {

      signer[s] = 0;
      delete signer[s];
  }


  function startTransfer(address _to, uint256 _amount) public {

    require(_to != address(0x0));
    require(balanceOf[msg.sender] >= _amount);
    _transfer(msg.sender, address(this), _amount);
    transcationId = transcationId+1;
    Transcation storage transcation = transcations[transcationId];
    transcation.from = msg.sender;
    transcation.to = _to;
    transcation.amount = _amount;
    transcation.signatureCounts = 0;
    pendingTranscation.push(transcationId);
    emit CreateTranscation(msg.sender, _to, _amount, transcationId);
  }

  function getPendingTranscation() public view returns (uint) {
    return pendingTranscation.length;
  }


  function getCounts(uint8 id) public view returns (uint256) {
      return transcations[id].signatureCounts;
  }

  function getBalance() public view returns (uint256) {
      return balanceOf[msg.sender];
  }

  function getAccountCount() public view returns (uint) {
      return accounts.length;
  }

  function multiSignature(uint8 id) public onlySigner {
    Transcation storage transcation = transcations[id];
    require(transcation.from != address(0x0));
    require(msg.sender != transcation.from);
    require(transcation.signatures[msg.sender] != 1);

    transcation.signatures[msg.sender] = 1;
    transcation.signatureCounts++;

    if (transcation.signatureCounts >= MIN_SIGNATURE) {
      require(balanceOf[address(this)] >= transcation.amount);

      _transfer(address(this), transcation.to, transcation.amount);

      delete pendingTranscation;
      delete transcations[id];

    }
  }

}



