
## 边缘节点挖矿教程

### 编译代码
```
git clone https://github.com/filepp/go-ipfs
cd go-ipfs
make build
cp cmd/ipfs /usr/local/bin
```

### 运行节点

运行
```
ipfs init
ipfs daemon
```

获取节点id
```
ipfs id
```


### 创建矿工
1 使用chrome浏览器，安装MetaMask钱包
2 MetaMask钱包配置好BSC地址。  
```
主网
https://bsc-dataseed1.binance.org/
链id: 56
```
```
测试网
https://data-seed-prebsc-1-s1.binance.org:8545/
链id： 97 
```
3 打开创建矿工页面，连接钱包。填入钱包地址和ipfs节点的id，点击创建矿工按钮。（确保钱包中有足够的BNB）

