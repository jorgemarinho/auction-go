package bid_usecase

import (
	"context"

	"github.com/jorgemarinho/auction-go/internal/internal_error"
)

func (bu *BidUseCase) FindByIdAuctionId(
	ctx context.Context,
	auctionId string) ([]BidOutputDTO, *internal_error.InternalError) {

	bidList, err := bu.BidRepository.FindByIdAuctionId(ctx, auctionId)

	if err != nil {
		return nil, err
	}

	var bidOutputList []BidOutputDTO

	for _, bid := range bidList {
		bidOutputList = append(bidOutputList, BidOutputDTO{
			Id:        bid.Id,
			UserId:    bid.UserId,
			AuctionId: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp,
		})
	}

	return bidOutputList, nil
}

func (bu *BidUseCase) FindWinningBiAuctionId(
	ctx context.Context,
	auctionId string) (*BidOutputDTO, *internal_error.InternalError) {

	bidEntity, err := bu.BidRepository.FindWinningBiAuctionId(ctx, auctionId)

	if err != nil {
		return nil, err
	}

	bidOutput := &BidOutputDTO{
		Id:        bidEntity.Id,
		UserId:    bidEntity.UserId,
		AuctionId: bidEntity.AuctionId,
		Amount:    bidEntity.Amount,
		Timestamp: bidEntity.Timestamp,
	}

	return bidOutput, nil
}
