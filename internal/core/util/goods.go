package util

import "github.com/lMikadal/POS_Backend_golang.git/internal/core/domain"

func ConverGoodsToGoodsResponse(goods *domain.Goods) domain.GoodsResponse {
	var tags []domain.TagResponseWithOutGoods
	for _, tag := range goods.Tags {
		tag := domain.TagResponseWithOutGoods{
			TagID:   tag.ID,
			TagName: tag.TagName,
		}
		tags = append(tags, tag)
	}
	var prices []domain.PriceResponse
	for _, price := range goods.Prices {
		price := domain.PriceResponse{
			PriceID:     price.ID,
			PriceAmount: price.PriceAmount,
			PricePrice:  price.PricePrice,
		}
		prices = append(prices, price)
	}

	goodsResponse := domain.GoodsResponse{
		GoodsID:     goods.ID,
		GoodsName:   goods.GoodsName,
		GoodsCode:   goods.GoodsCode,
		GoodsAmount: goods.GoodsAmount,
		GoodsCost:   goods.GoodsCost,
		Tags:        tags,
		Prices:      prices,
	}

	return goodsResponse
}
