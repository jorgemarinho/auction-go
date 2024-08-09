package auction

import (
	"context"
	"fmt"
	"time"

	"github.com/jorgemarinho/auction-go/configuration/logger"
	"github.com/jorgemarinho/auction-go/internal/entity/auction_entity"
	"github.com/jorgemarinho/auction-go/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (ar *AuctionRepository) FindAuctionById(ctx context.Context, id string) (*auction_entity.Auction, *internal_error.InternalError) {
	filter := bson.M{"_id": id}
	var auctionEntityMongo AuctionEntityMongo

	if err := ar.Collection.FindOne(ctx, filter).Decode(&auctionEntityMongo); err != nil {
		logger.Error(fmt.Sprintf("Error trying to find auction by id = %s", id), err)
		return nil, internal_error.NewInternalServerError("Error trying to find auction")
	}
	return &auction_entity.Auction{
		Id:          auctionEntityMongo.Id,
		ProductName: auctionEntityMongo.ProductName,
		Category:    auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Condition:   auctionEntityMongo.Condition,
		Status:      auctionEntityMongo.Status,
		Timestamp:   time.Unix(auctionEntityMongo.Timestamp, 0),
	}, nil
}

func (ar *AuctionRepository) FindAuctions(
	ctx context.Context,
	status auction_entity.AuctionStatus,
	category string,
	productName string) ([]auction_entity.Auction, *internal_error.InternalError) {

	filter := bson.M{}

	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = category
	}

	if productName != "" {
		filter["product_name"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}

	cursor, err := ar.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error("Error trying to find auctions", err)
		return nil, internal_error.NewInternalServerError("Error trying to find auctions")
	}

	defer cursor.Close(ctx)

	var auctionEntityMongo []AuctionEntityMongo
	if err := cursor.All(ctx, &auctionEntityMongo); err != nil {
		logger.Error("Error trying to decode auctions", err)
		return nil, internal_error.NewInternalServerError("Error trying to find auctions")
	}

	var auctionEntity []auction_entity.Auction
	for _, a := range auctionEntityMongo {
		auctionEntity = append(auctionEntity, auction_entity.Auction{
			Id:          a.Id,
			ProductName: a.ProductName,
			Category:    a.Category,
			Description: a.Description,
			Condition:   a.Condition,
			Status:      a.Status,
			Timestamp:   time.Unix(a.Timestamp, 0),
		})
	}

	return auctionEntity, nil
}
