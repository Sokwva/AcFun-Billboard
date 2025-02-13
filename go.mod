module sokwva/acfun/billboard

go 1.23

toolchain go1.23.3

require (
	github.com/BurntSushi/toml v1.4.0
	github.com/PuerkitoBio/goquery v1.10.1
	google.golang.org/grpc v1.70.0
	sokwva/acfun/dougaInfo v0.0.0-00010101000000-000000000000
)

require (
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/influxdata/line-protocol v0.0.0-20210922203350-b1ad95c89adf // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/montanaflynn/stats v0.7.1 // indirect
	github.com/oapi-codegen/runtime v1.1.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.2 // indirect
	github.com/xdg-go/stringprep v1.0.4 // indirect
	github.com/youmark/pkcs8 v0.0.0-20240726163527-a2c0da244d78 // indirect
	golang.org/x/crypto v0.33.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250207221924-e9438ea467c6 // indirect
	google.golang.org/protobuf v1.36.5 // indirect
)

require (
	github.com/andybalholm/cascadia v1.3.3 // indirect
	github.com/influxdata/influxdb-client-go/v2 v2.14.0
	go.mongodb.org/mongo-driver v1.17.2
	golang.org/x/net v0.35.0 // indirect
)

replace sokwva/acfun/dougaInfo => ../dougaInfo
