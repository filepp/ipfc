# ipfc

IPFC（Inter-Planetary File Cache） 星际文件缓存

![](docs/image/topology.jpg)

### 编译代码
IPFS
```
clone https://github.com/filepp/go-ipfs
make build
```
IPFC
```
clone https://github.com/filepp/go-ipfs
make build
```

### 部署中心节点
运行ipfs，需要指定矿工钱包地址
```
ipfs init
ipfs daemon --enable-pubsub-experiment=true --enable-mining=true --wallet-address=0x3a9e6cf4e3157670a3b991c25d6f4fcbd9419c03 --miner-role=0
```
运行lotus（略）

运行ipfc, 需修改配置 `resource/ipfc.yaml`
运行ipfc，(ipfc和lotus必须运行同一台主机上，或者将IPFC的文件目录挂载到lotus主机上)
```
./ipfc
```

### 边沿节点
运行边沿节点，需要指定矿工钱包地址
```
ipfs init
ipfs daemon --enable-pubsub-experiment=true --enable-mining=true --wallet-address=0x3a9e6cf4e3157670a3b991c25d6f4fcbd9419c03 --miner-role=1
```

### 巡检节点
运行ipfs
```
ipfs init
ipfs daemon --enable-pubsub-experiment=true
```
运行inspector
```
./inspector
```