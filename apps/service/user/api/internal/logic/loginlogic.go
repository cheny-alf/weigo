package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"

	"wego/apps/common/constant"
	"wego/apps/service/user/api/internal/svc"
	"wego/apps/service/user/api/internal/types"
	"wego/apps/service/user/model/dao"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext

	dao *dao.DBDAO
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		dao:    dao.NewDBDAO(svcCtx.Client),
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	user, err := l.dao.User.GetUserByTel(l.ctx, req.Tel)
	if err != nil {
		return nil, err
	}
	if user.Password != req.Password {
		resp.Code = constant.ErrorUserNotFound
		resp.Msg = constant.GetErrorMsg(constant.ErrorUserNotFound)
		return resp, nil
	}
	// todo 更新ip相关
	resp = &types.LoginResp{
		NickName:      user.NickName,
		RealName:      user.RealName,
		Email:         user.Email,
		Tel:           user.Tel,
		LastLoginTime: int(user.LastLoginTime),
		LastIPAddr:    user.LastIPAddr,
	}
	return resp, nil
}
