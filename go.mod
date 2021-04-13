module ipfc

go 1.15

require (
	github.com/filecoin-project/go-address v0.0.5
	github.com/filecoin-project/go-fil-markets v1.1.9
	github.com/filecoin-project/go-jsonrpc v0.1.4-0.20210217175800-45ea43ac2bec
	github.com/filecoin-project/lotus v1.5.3
	github.com/filecoin-project/specs-actors/v3 v3.1.0 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/go-sql-driver/mysql v1.5.0
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/ipfs/go-cid v0.0.7
	github.com/ipfs/go-ipfs v0.8.0
	github.com/ipfs/go-ipfs-files v0.0.8
	github.com/ipfs/go-ipfs-http-client v0.1.0
	github.com/ipfs/go-log/v2 v2.1.2
	github.com/ipfs/interface-go-ipfs-core v0.4.0
	github.com/jinzhu/configor v1.2.1
	github.com/jinzhu/gorm v1.9.16
	github.com/libp2p/go-libp2p-core v0.8.5
	github.com/liyue201/golib v0.0.0-20210225015707-e527cc337867
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/prometheus/common v0.18.0
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1
)

replace (
	github.com/ipfs/go-ipfs v0.8.0 => github.com/filepp/go-ipfs v0.8.1-0.20210413041357-db8e769044a8
	github.com/ipfs/go-log/v2 v2.1.2 => github.com/liyue201/go-log/v2 v2.1.2-0.20210312022747-08ac38b3a792
)
