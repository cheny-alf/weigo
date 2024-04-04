package logic

import (
	"context"
	"wego/apps/service/user/model/dao"

	"wego/apps/service/user/api/internal/svc"
	"wego/apps/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnfollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	dao *dao.DBDAO
}

func NewUnfollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnfollowLogic {
	return &UnfollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		dao:    dao.NewDBDAO(svcCtx.Client),
	}
}

func (l *UnfollowLogic) Unfollow(req *types.FollowReq) (resp *types.FollowResp, err error) {
	err = l.dao.User.UnFollowUser(l.ctx, req.UserID, req.FollowingUserID)
	if err != nil {
		return nil, err
	}
	return
}
