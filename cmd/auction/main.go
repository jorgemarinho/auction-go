package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jorgemarinho/auction-go/configuration/database/mongodb"
	"github.com/jorgemarinho/auction-go/internal/infra/api/web/controller/auction_controller"
	"github.com/jorgemarinho/auction-go/internal/infra/api/web/controller/bid_controller"
	"github.com/jorgemarinho/auction-go/internal/infra/api/web/controller/user_controller"
	"github.com/jorgemarinho/auction-go/internal/infra/database/auction"
	"github.com/jorgemarinho/auction-go/internal/infra/database/bid"
	"github.com/jorgemarinho/auction-go/internal/infra/database/user"
	"github.com/jorgemarinho/auction-go/internal/usecase/auction_usecase"
	"github.com/jorgemarinho/auction-go/internal/usecase/bid_usecase"
	"github.com/jorgemarinho/auction-go/internal/usecase/user_usecase"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {

	ctx := context.Background()

	if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
		return
	}

	databaseConnection, err := mongodb.NewMongoDBConnection(ctx)

	if err != nil {
		log.Fatal(err.Error())
		return
	}

	router := gin.Default()

	userController, bidController, auctionController := initDependecies(databaseConnection)

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "welcome to auction api",
		})
	})

	router.GET("/auction", auctionController.FindAuctions)
	router.GET("/auction/:auctionId", auctionController.FindAuctionById)
	router.GET("/auction/winner/:auctionId", auctionController.FindWinningBidByAuctionId)
	router.POST("/auction", auctionController.CreateAuction)
	router.POST("/bid", bidController.CreateBid)
	router.GET("/bid/:auctionId", bidController.FindBidAuctionById)
	router.GET("/user/:userId", userController.FindUserById)

	router.Run(":8080")
}

func initDependecies(database *mongo.Database) (
	userController *user_controller.UserController,
	bidController *bid_controller.BidController,
	auctionController *auction_controller.AuctionController) {

	auctionRepository := auction.NewAuctionRepository(database)
	bidRepository := bid.NewBidRepository(database, auctionRepository)
	userRepository := user.NewUserRepository(database)

	userController = user_controller.NewUserController(user_usecase.NewUserUseCase(userRepository))

	auctionController = auction_controller.NewAuctionController(auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))

	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))

	return

}
