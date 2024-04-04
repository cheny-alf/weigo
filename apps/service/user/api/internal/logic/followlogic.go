package logic

import (
	"context"
	"wego/apps/service/user/model/dao"

	"wego/apps/service/user/api/internal/svc"
	"wego/apps/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FollowLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	dao *dao.DBDAO
}

func NewFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FollowLogic {
	return &FollowLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		dao:    dao.NewDBDAO(svcCtx.Client),
	}
}

func (l *FollowLogic) Follow(req *types.FollowReq) (resp *types.FollowResp, err error) {
	err = l.dao.User.FollowUser(l.ctx, req.UserID, req.FollowingUserID)
	if err != nil {
		return nil, err
	}
	return
}
