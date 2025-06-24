#### Setup

- go mod tidy

#### Run Webhook

- go run main.go webhook --port[optional] 8080

#### Run Dequeue Listener to FIFO customers to agents

- go run main.go dequeue

#### Run Resolve Listener to assign next customer to agent

- go run main.go resolve

#### Technical Documentation

#### Flowchart

![Flowchart](https://raw.githubusercontent.com/ahmad8taufiq/qiscus-omnichannel/refs/heads/main/flowchart.png)

#### Sequence Diagram

![Sequence Diagram](https://raw.githubusercontent.com/ahmad8taufiq/qiscus-omnichannel/refs/heads/main/sequence_diagram.png)

#### Database Design

![Database Design](https://raw.githubusercontent.com/ahmad8taufiq/qiscus-omnichannel/refs/heads/main/database_design.png)
