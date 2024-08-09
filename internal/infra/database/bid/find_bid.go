package bid

import (
	"context"
	"fmt"
	"time"

	"github.com/jorgemarinho/auction-go/configuration/logger"
	"github.com/jorgemarinho/auction-go/internal/entity/bid_entity"
	"github.com/jorgemarinho/auction-go/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (bd *BidRepository) FindByIdAuctionId(
	ctx context.Context,
	auctionId string) ([]bid_entity.Bid, *internal_error.InternalError) {

	filter := bson.M{"auction_id": auctionId}

	cursor, err := bd.Collection.Find(ctx, filter)
	if err != nil {
		logger.Error(fmt.Sprintf("Error trying to find bids by auction id = %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error trying to find bids by auction id = %s", auctionId))
	}

	var bidEntityMongo []BidEntityMongo
	if err := cursor.All(ctx, &bidEntityMongo); err != nil {
		logger.Error(fmt.Sprintf("Error trying to find bids by auction id = %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error trying to find bids by auction id = %s", auctionId))
	}

	bidEntities := make([]bid_entity.Bid, 0)

	for _, bidEntityMongo := range bidEntityMongo {
		bidEntities = append(bidEntities, bid_entity.Bid{
			Id:        bidEntityMongo.Id,
			UserId:    bidEntityMongo.UserId,
			AuctionId: bidEntityMongo.AuctionId,
			Amount:    bidEntityMongo.Amount,
			Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
		})
	}

	return bidEntities, nil
}

func (bd *BidRepository) FindWinningBiAuctionId(
	ctx context.Context,
	auctionId string) (*bid_entity.Bid, *internal_error.InternalError) {

	filter := bson.M{"auction_id": auctionId}

	var bidEntityMongo BidEntityMongo
	opts := options.FindOne().SetSort(bson.D{{Key: "amount", Value: -1}})

	if err := bd.Collection.FindOne(ctx, filter, opts).Decode(&bidEntityMongo); err != nil {
		logger.Error(fmt.Sprintf("Error trying to find winning bid by auction id = %s", auctionId), err)
		return nil, internal_error.NewInternalServerError(fmt.Sprintf("Error trying to find winning bid by auction id = %s", auctionId))
	}

	return &bid_entity.Bid{
		Id:        bidEntityMongo.Id,
		UserId:    bidEntityMongo.UserId,
		AuctionId: bidEntityMongo.AuctionId,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil

}
