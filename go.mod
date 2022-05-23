module github.com/liamsi/convert-addrs

go 1.18

require github.com/cosmos/cosmos-sdk v0.45.4

require github.com/cosmos/btcutil v1.0.4 // indirect

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
)
