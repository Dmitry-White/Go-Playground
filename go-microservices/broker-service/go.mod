module go-microservices/broker-service

go 1.20

require github.com/go-chi/chi/v5 v5.0.10

require github.com/go-chi/cors v1.2.1

require (
	github.com/rabbitmq/amqp091-go v1.8.1
	google.golang.org/grpc v1.57.0
	google.golang.org/protobuf v1.30.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	golang.org/x/net v0.9.0 // indirect
	golang.org/x/sys v0.7.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230525234030-28d5490b6b19 // indirect
)
