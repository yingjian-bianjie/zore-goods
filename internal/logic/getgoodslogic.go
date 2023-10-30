package logic

import (
	"context"

	"github.com/yingjian-bianjie/zore-goods/internal/svc"
	"github.com/yingjian-bianjie/zore-goods/types/goods"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetGoodsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetGoodsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetGoodsLogic {
	return &GetGoodsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// rpc方法
func (l *GetGoodsLogic) GetGoods(in *goods.GoodsRequest) (res *goods.GoodsResponse, err error) {
	// todo: add your logic here and delete this line
	goodsId := in.GoodsId
	res = new(goods.GoodsResponse)
	res.GoodsId = goodsId
	res.Name = "茅台"
	return
}
