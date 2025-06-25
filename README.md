# Qiscus Agent Assignment

A Service to automatically assign available agents to customer chats in Qiscus Omnichannel. This tool helps streamline customer service routing by integrating with Qiscus APIs and Redis for queue management.

## Getting Started

### Prerequisites

- Redis
- Qiscus Account

### Installation

1.  Copy and complate your environtment

        cp .env.example .env

2.  Update package

        go get

### Run Service

1. Run webhook to listen customer chat

   `go run main.go webhook --port[optional] 8080`

2. Run dequeue to consume queue customer chat

   `go run main.go dequeue`

3. Run resolve to listen admin resolve

   `go run main.go resolve`

### Technical Documentation

### Flowchart

![Flowchart](https://raw.githubusercontent.com/ahmad8taufiq/qiscus-omnichannel/refs/heads/main/flowchart.png)

```
flowchart TD
    A@{ shape: circle, label: "Customer" }--> |Chat| B[Qiscus]
    B --> |Forward chat| C[Webhook]
    C --> |Submit chat to queue| D[(new_session_queue)]
    D[(new_session_queue)] -->|Consume chat| E[Dequeue]
    E --> |Check| F{Is Agent Available ?}
    F --> |No| D
    F --> |Yes| H{Is Agent Customer Count Less Than Two ?}
    H --> |No| D
    H --> |Yes| I[Assign Customer to Agent]
    I --> |Increment| J[(Agent Customer Count)]

    AA@{ shape: circle, label: "Qiscus" }--> |Notify Resolved Chat| BB[Webhook Resolve]
    BB --> |Decrement| CC[(Agent Customer Count)]
```

#### Sequence Diagram

![Sequence Diagram](https://raw.githubusercontent.com/ahmad8taufiq/qiscus-omnichannel/refs/heads/main/sequence_diagram.png)

#### Database Design

![Database Design](https://raw.githubusercontent.com/ahmad8taufiq/qiscus-omnichannel/refs/heads/main/database_design.png)
