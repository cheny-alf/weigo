package logic

import (
	"context"
	"fmt"
	"time"
	"wego/apps/service/user/model/dao"

	"wego/apps/service/user/api/internal/svc"
	"wego/apps/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	dao *dao.DBDAO
}

func NewGetFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFollowerListLogic {
	return &GetFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		dao:    dao.NewDBDAO(svcCtx.Client),
	}
}

func (l *GetFollowerListLogic) GetFollowerList(req *types.GetFollowerListReq) (resp *types.GetFollowerListResp, err error) {
	fmt.Println(time.Now())
	deadline, _ := l.ctx.Deadline()
	fmt.Println(deadline)
	fmt.Println(deadline.Format("2006:03:04 15:16:17"))
	user, err := l.dao.User.GetUserByID(l.ctx, req.UserID)
	if err != nil {
		return nil, err
	}

	followers, err := l.dao.User.GetFollowers(l.ctx, user)
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
	resp = &types.GetFollowerListResp{
		CommonResp: types.CommonResp{},
		Data:       fans,
	}
	return
}
