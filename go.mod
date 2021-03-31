module ipfc

go 1.15

require (
	github.com/filecoin-project/go-jsonrpc v0.1.4-0.20210217175800-45ea43ac2bec
	github.com/filecoin-project/lotus v1.5.3
	github.com/filecoin-project/specs-actors/v3 v3.1.0 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/gofrs/uuid v4.0.0+incompatible
	github.com/ipfs/go-ipfs-files v0.0.8
	github.com/ipfs/go-ipfs-http-client v0.0.5
	github.com/ipfs/go-log/v2 v2.1.2
	github.com/ipfs/interface-go-ipfs-core v0.2.3
	github.com/jinzhu/configor v1.2.1
	github.com/liyue201/golib v0.0.0-20210225015707-e527cc337867
	github.com/multiformats/go-multiaddr v0.3.1
	github.com/prometheus/common v0.10.0
)

replace github.com/ipfs/go-log/v2 v2.1.2 => github.com/liyue201/go-log/v2 v2.1.2-0.20210312022747-08ac38b3a792
