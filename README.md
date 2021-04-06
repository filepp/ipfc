# ipfc

IPFC（Inter-Planetary File Cache） 星际文件缓存

### 安装依赖

### IPFS
[部署IPFS](docs/IPFS.md)

### Lotus
安装（略）  
获取地址
```
lotus auth api-info --perm admin
```

### 部署IPFC
编译代码
```
make
```
修改配置 `conf/ipfc.yaml`

运行IPFC，(IPFC和lotus必须运行同一台主机上，或者将IPFC的文件目录挂载到lotus主机上)
```
./ipfc
```
