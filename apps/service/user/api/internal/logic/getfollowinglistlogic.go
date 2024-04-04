package logic

import (
	"context"
	"wego/apps/service/user/model/dao"

	"wego/apps/service/user/api/internal/svc"
	"wego/apps/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowingListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	dao *dao.DBDAO
}

func NewGetFollowingListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowingListLogic {
	return &GetFollowingListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		dao:    dao.NewDBDAO(svcCtx.Client),
	}
}

func (l *GetFollowingListLogic) GetFollowingList(req *types.GetFollowingListReq) (resp *types.GetFollowingListResp, err error) {
	user, err := l.dao.User.GetUserByID(l.ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	followers, err := l.dao.User.GetFollowing(l.ctx, user)
	if err != nil {
		return nil, err
	}
	fans := make([]types.Fan, len(followers))
	for i, follower := range followers {
		fans[i] = types.Fan{
			ID:       i,
			UserID:   follower.ID,
			NickName: follower.NickName,
		}
	}
	resp = &types.GetFollowingListResp{
		CommonResp: types.CommonResp{},
		Data:       fans,
	}
	return
}
