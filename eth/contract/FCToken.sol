// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

contract FCToken {

  string public name; 
  string public symbol; 
  uint8 public decimals; 
  uint256 public totalSupply;
  address public owner; 

  mapping (address => uint256) public balanceOf;
  mapping (address => mapping(address => uint256)) public approvalBanlance;
  mapping(address => uint8) public signer; 
  mapping (address => uint256) public freezeOf;
  
  event Transfer(address from, address to, uint256 value);
  event Approval(address from, address delegatee, uint256 value);
  event Destory(address destorys, uint256 value);	
  /* This notifies clients about the amount frozen */
  event Freeze(address indexed from, uint256 value);
  /* This notifies clients about the amount unfrozen */
  event Unfreeze(address indexed from, uint256 value);
    
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
    
  constructor(string memory _name, string memory _symbol, uint8 _decimals, uint256 _initiSupply){
    name = _name;
    symbol = _symbol;
    decimals = _decimals;
    totalSupply = _initiSupply * 10 ** uint256(decimals);
    balanceOf[msg.sender] = totalSupply;
    owner = msg.sender;
  }

  function _transfer(address _from, address _to, uint256 _value) internal {
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
    require(_to != address(_to));
    require(balanceOf[msg.sender] >= _value);
    require(_value > 0);
    if(owner==msg.sender){
        startTransfer(_to,_value);
    }else{
        _transfer(msg.sender, _to, _value);
    }
    
  }

  function transferFrom(address _from, address _to, uint256 _value) public {
    require(_from != address(0x0)&&_to != address(0x0));
    require(approvalBanlance[_from][msg.sender] >= _value); //授权Token必须大于等于value
    approvalBanlance[_from][msg.sender] -= _value;
    _transfer(_from, _to, _value);
  }
  

  function approval(address _delegatee, uint256 _value) public {
    
    require(signer[msg.sender] == 1||msg.sender == owner);
    require(_delegatee != address(0x0));
    require(balanceOf[msg.sender] >= _value); 
    approvalBanlance[msg.sender][_delegatee] = _value;
    emit Approval(msg.sender, _delegatee, _value);
  }
  

 function approveWthArray(address[] memory _spenders, uint256[] memory _values) public
        returns (bool success) {
        require(signer[msg.sender] == 1||msg.sender == owner);   
		require (_values.length > 0&&_spenders.length>0);
        for (uint256 i = 0;i < _spenders.length;i++){
            require(balanceOf[msg.sender] >= _values[i]);
            approvalBanlance[msg.sender][_spenders[i]] += _values[i];
        }
        return true;
    }

  function destory(uint _value) public {
    require(balanceOf[msg.sender] >= _value);
    balanceOf[msg.sender] -= _value;
    totalSupply -= _value;
    emit Destory(msg.sender, _value);
  }

  
  function destoryFrom(address _from, uint _value) public {
    require(approvalBanlance[_from][msg.sender] >= _value);
    require(balanceOf[_from] >= _value);
    balanceOf[_from] -= _value;
    approvalBanlance[_from][msg.sender] -= _value;
    totalSupply -= _value;
    emit Destory(msg.sender, _value);
  }
  
  function freeze(address who,uint256 _value) public returns (bool success) {
	    require (msg.sender==owner);
	    require (who != address(0x0));
	    require (_value > 0);
	    require (balanceOf[who] > _value);
        balanceOf[who] = balanceOf[who] - _value;                      // Subtract from the sender
        freezeOf[who] = freezeOf[who] + _value;                                // Updates totalSupply
        emit Freeze(who, _value);
        return true;
    }

	function unfreeze(address who,uint256 _value) public returns (bool success) {
	    require (msg.sender==owner);
	    require (who != address(0x0));
        require (freezeOf[who] >= _value);            // Check if the sender has enough
		require (_value > 0); 
        freezeOf[who] = freezeOf[who] - _value;                      // Subtract from the sender
		balanceOf[who] = balanceOf[who] + _value;
        emit Unfreeze(who, _value);
        return true;
    }
    
    modifier onlyOwner() {
     require(owner == msg.sender);
     _;
    }

  modifier onlySigner() {
    require(owner == msg.sender || signer[msg.sender] == 1);
    _;
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



