# ipfc

IPFC（Inter-Planetary File Cache） 星际文件缓存

![](docs/image/topology.jpg)

### 部署依赖
- mysql
- ipfs
- lotus

### 编译代码
IPFS
```
git clone https://github.com/filepp/go-ipfs
cd go-ipfs
make build
cp cmd/ipfs /usr/local/bin
```
IPFC
```
git clone  https://git.filep.vip/ruitai/ipfc
cd ipfc
make
```

### 部署中心节点
运行ipfs，需要指定矿工钱包地址和矿工角色,(0代表中心节点)
```
ipfs init
ipfs daemon --enable-pubsub-experiment=true --enable-mining=true --miner-role=0
```
运行lotus（略）

运行ipfc, 需修改配置 `conf/ipfc.yaml`
运行ipfc，(ipfc和lotus必须运行同一台主机上，或者将IPFC的文件目录挂载到lotus主机上)
```
./ipfc
```

### 边缘节点
运行边缘节点，需要指定矿工钱包地址和矿工角色，(1代表边缘节点)
```
ipfs init
ipfs daemon --enable-pubsub-experiment=true --enable-mining=true --miner-role=1
```
查看 peerID, 该id在web页面创建矿工时作为参数使用。
```
ipfs id
```

### 巡检节点
运行ipfs
```
ipfs init
ipfs daemon --enable-pubsub-experiment=true
```

运行inspector， 需修改配置 `conf/inspector.yaml`
```
./inspector
```
