module github.com/abrahamSN/shippy/shippy-cli-consignment

go 1.18

replace github.com/<YourUsername>/shippy/shippy-service-consignment => ../shippy-service-consignment

require (
	github.com/abrahamSN/shippy/shippy-service-consignment v0.0.0-20220412180755-9a81af223081
	google.golang.org/grpc v1.45.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5 // indirect
	golang.org/x/sys v0.0.0-20220412071739-889880a91fd5 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20220407144326-9054f6ed7bac // indirect
	google.golang.org/protobuf v1.28.0 // indirect
)
