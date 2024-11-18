module sokwva/acfun/billboard

go 1.23

toolchain go1.23.3

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/PuerkitoBio/goquery v1.10.0
	google.golang.org/grpc v1.68.0
	sokwva/acfun/dougaInfo v0.0.0-00010101000000-000000000000
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/deepmap/oapi-codegen v1.12.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/influxdata/line-protocol v0.0.0-20210922203350-b1ad95c89adf // indirect
	golang.org/x/sys v0.27.0 // indirect
	golang.org/x/text v0.20.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20241113202542-65e8d215514f // indirect
	google.golang.org/protobuf v1.35.1 // indirect
)

require (
	github.com/andybalholm/cascadia v1.3.2 // indirect
	github.com/influxdata/influxdb-client-go/v2 v2.12.4
	golang.org/x/net v0.31.0 // indirect
)

replace sokwva/acfun/dougaInfo => ../dougaInfo
