// Code generated by goctl. DO NOT EDIT.
// Source: goods.proto

package server

import (
	"context"

	"github.com/yingjian/zore-goods/internal/logic"
	"github.com/yingjian/zore-goods/internal/svc"
	"github.com/yingjian/zore-goods/types/goods"
)

type GoodsServer struct {
	svcCtx *svc.ServiceContext
	goods.UnimplementedGoodsServer
}

func NewGoodsServer(svcCtx *svc.ServiceContext) *GoodsServer {
	return &GoodsServer{
		svcCtx: svcCtx,
	}
}

// rpc方法
func (s *GoodsServer) GetGoods(ctx context.Context, in *goods.GoodsRequest) (*goods.GoodsResponse, error) {
	l := logic.NewGetGoodsLogic(ctx, s.svcCtx)
	return l.GetGoods(in)
}
