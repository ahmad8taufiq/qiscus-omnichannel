package cmd

import (
	"fmt"
	"net/http"

	"qiscus-omnichannel/app"
	"qiscus-omnichannel/repository"
	"qiscus-omnichannel/service"
	"qiscus-omnichannel/tools/console"
	"qiscus-omnichannel/tools/logger"

	"github.com/spf13/cobra"
)

var serverPort int

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start server",
	Run:   runServer,
}

func init() {
	serverCmd.Flags().IntVarP(&serverPort, "port", "p", 8081, "Port for constant server")
	rootCmd.AddCommand(serverCmd)
}

func runServer(_ *cobra.Command, _ []string) {
	log := logger.Logger
	redisRepo := repository.NewRedisRepository()
	redisService := service.RedisService(redisRepo)
	
	chatRepo := repository.NewChatRepository()
	chatService := service.ChatService(chatRepo)

	commentRepo := repository.NewCommentRepository()
	commentService := service.CommentService(commentRepo)

	authRepo := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepo)

	roomRepo := repository.NewRoomRepository()
	roomService := service.RoomService(roomRepo)

	http.HandleFunc("/config", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			app.GetMaxCustomerPerAgentHandler(redisService)(w, r)
		case http.MethodPut:
			app.SetMaxCustomerPerAgentHandler(redisService)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			app.ChatWithDelayHandler(chatService, commentService, authService, roomService, redisService)(w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	addr := fmt.Sprintf(":%d", serverPort)
	console.ConsoleGreet("Web Server", "1.0.0", "", serverPort)
	
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("❌ Failed to start constant server: %v", err)
	}
}
