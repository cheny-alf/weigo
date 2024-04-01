package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"wego/apps/service/user/api/internal/svc"
	"wego/apps/service/user/api/internal/types"
	"wego/apps/service/user/model/dao"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	dao *dao.DBDAO
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		dao:    dao.NewDBDAO(svcCtx.Client),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(ctx context.Context, req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	user, err := l.dao.User.GetUserByID(ctx, req.UserID)
	if err != nil {
		return nil, err
	}
	resp = &types.GetUserInfoResp{
		NickName:      user.NickName,
		RealName:      user.RealName,
		Email:         user.Email,
		Tel:           user.Tel,
		LastLoginTime: int(user.LastLoginTime),
		LastIPAddr:    user.LastIPAddr,
	}
	return resp, nil
}
