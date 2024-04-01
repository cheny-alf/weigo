package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"wego/apps/service/user/api/internal/svc"
	"wego/apps/service/user/api/internal/types"
	"wego/apps/service/user/model/dao"
	ent2 "wego/apps/service/user/model/ent"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	client *ent2.Client
	dao    *dao.DBDAO
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		client: svcCtx.Client,
		dao:    dao.NewDBDAO(svcCtx.Client),
	}
}

func (l *RegisterLogic) Register(ctx context.Context, req *types.RegisterReq) (resp *types.RegisterResp, err error) {
	user, err := l.dao.User.CreateUser(ctx, &ent2.User{
		NickName: req.NickName,
		RealName: req.RealName,
		Email:    req.Email,
		Tel:      req.Tel,
		Password: req.Password,
	})
	resp = &types.RegisterResp{
		NickName:      user.NickName,
		RealName:      user.RealName,
		Email:         user.Email,
		Tel:           user.Tel,
		LastLoginTime: user.LastLoginTime,
		LastIPAddr:    user.LastIPAddr}
	return resp, nil
}
