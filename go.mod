module github.com/razin99/cq-source-dnsimple

go 1.21.1

toolchain go1.21.4

require (
	github.com/apache/arrow/go/v14 v14.0.0-20231031200323-c49e24273160
	github.com/cloudquery/plugin-pb-go v1.14.4
	github.com/cloudquery/plugin-sdk/v4 v4.20.0
	github.com/rs/zerolog v1.29.1
)

replace github.com/apache/arrow/go/v14 => github.com/cloudquery/arrow/go/v14 v14.0.0-20230904001200-cd3d4114faa0

require (
	github.com/dnsimple/dnsimple-go v1.5.1 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/shopspring/decimal v1.3.1 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/oauth2 v0.14.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20231127180814-3a041ad873d4 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
)
