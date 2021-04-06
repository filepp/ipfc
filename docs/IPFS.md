
### 安装IPFS
参考： https://docs.ipfs.io/recent-releases/go-ipfs-0-7/install/#linux    
ipfs下载地址 https://dist.ipfs.io/go-ipfs/v0.7.0  
```
wget https://dist.ipfs.io/go-ipfs/v0.7.0/go-ipfs_v0.7.0_linux-amd64.tar.gz
tar -xvzf go-ipfs_v0.7.0_linux-amd64.tar.gz
cd go-ipfs
./install.sh
ipfs --version
```
初始化仓库目录，会自动生成目录` ~/.ipfs`
```
ipfs init
```

启动节点
```
ipfs bootstrap rm --all　
nohup ipfs daemon &> ipfs.log &
```


### 部署webui， 

安装webui，指定端口3001
```
https://github.com/ipfs-shipyard/ipfs-webui
cd ipfs-webui
npm install -g cnpm --registry=https://registry.npm.taobao.org	
cnpm install
PORT=3001 npm start
```
或者使用docker
```
docker run -d --name webui -p 3001:3000  garychen/ipfs-webui:latest
```

打开web, 假设ip为192.168.3.101

http://192.168.3.101:3001

报错，根据提示修改配置 `~/.ipfs/config`,以下4处
```
{
  "API": {
    "HTTPHeaders": {
      "Access-Control-Allow-Methods": [  //1.这里
        "PUT",
        "POST"
      ],
      "Access-Control-Allow-Origin": [  
        "http://192.168.3.101:3001"   // 2.这里
      ]
    }
  },
  "Addresses": {
    "API": "/ip4/0.0.0.0/tcp/5001",    // 3.这里
    "Announce": [],
    "Gateway": "/ip4/192.168.2.101/tcp/8080",  // 4.这里
    "NoAnnounce": [], 
    "Swarm": [
      "/ip4/0.0.0.0/tcp/4001",
      "/ip6/::/tcp/4001",
      "/ip4/0.0.0.0/udp/4001/quic",
      "/ip6/::/udp/4001/quic"
    ]
  },
   
  // 省略
   
}

```

重启ipfs
```
nohup ipfs daemon &> ipfs.log &
```
在web页面配置节点地址 `/ip4/192.168.3.101/tcp/5001`




