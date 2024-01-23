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

	goodsResponse := domain.GoodsResponse{
		GoodsID:     goods.ID,
		GoodsName:   goods.GoodsName,
		GoodsCode:   goods.GoodsCode,
		GoodsImage:  goods.GoodsImage,
		GoodsAmount: goods.GoodsAmount,
		GoodsCost:   goods.GoodsCost,
		GoodsPrice:  goods.GoodsPrice,
		Tags:        tags,
	}

	return goodsResponse
}
