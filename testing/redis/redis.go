package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"qiscus-omnichannel/repository"
)

type Agent struct {
	ID    int    `json:"id"`
	CurrentCustomerCount  int `json:"current_customer_count"`
}

func main() {
	agent1 := Agent{
		ID:    12345,
		CurrentCustomerCount:  1,
	}
	agent2 := Agent{
		ID:    23456,
		CurrentCustomerCount:  0,
	}

	users := []Agent{agent1, agent2}

	jsonBytes, err := json.Marshal(users)
	if err != nil {
		log.Fatalf("Failed to marshal: %v", err)
	}

	jsonString := string(jsonBytes)

	redisRepo := repository.NewRedisRepository()

	err = redisRepo.SetCache("agents", jsonString, 10*time.Minute)
	if err != nil {
		log.Fatalf("SetJSON failed: %v", err)
	}

	fmt.Println("✅ JSON stored")

	data, _ := redisRepo.GetCache("agents")

	var agents []Agent
	err = json.Unmarshal([]byte(data), &agents)
	if err != nil {
		log.Fatalf("Failed to unmarshal: %v", err)
	}

	if len(agents) == 0 {
		log.Println("⚠️ No agents available")
	} else {
		minAgent := agents[0]
		for _, agent := range agents[1:] {
			if agent.CurrentCustomerCount < minAgent.CurrentCustomerCount {
				// minAgent = agent
			}
		}

		fmt.Printf("✅ Agent with lowest customer count: %+v\n", minAgent)
	}	
}
