#### Setup

- go mod tidy

#### Run Webhook

- go run main.go webhook --port[optional] 8080

#### Set Max Customer per Agent

1. Start server : go run main.go server --port[optional] 8081
2. Setter : curl -X PUT https://40f2-103-102-12-15.ngrok-free.app/config -H "Content-Type: application/json" -d '{"max_customer_per_agent": 2}'
3. Getter : curl -X GET https://40f2-103-102-12-15.ngrok-free.app/config

#### Run Dequeue Listener

- go run main.go webhook

#### Run Resolver Listener

- go run main.go resolve
