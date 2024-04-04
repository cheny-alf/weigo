package dao

import (
	"context"
	"time"

	ent2 "wego/apps/service/user/model/ent"
	"wego/apps/service/user/model/ent/user"
)

type UserDao struct {
	db *ent2.Client
	tx *ent2.Tx
}

func NewUserDao(client *ent2.Client) *UserDao {
	return &UserDao{
		db: client,
	}
}

func NewUserDaoFromTx(tx *ent2.Tx) *UserDao {
	return &UserDao{
		tx: tx,
	}
}

func (d *UserDao) Client() *ent2.UserClient {
	if d.db != nil {
		return d.db.User
	}
	if d.tx != nil {
		return d.tx.User
	}
	return nil
}

func (d *UserDao) CreateUser(ctx context.Context, user *ent2.User) (*ent2.User, error) {
	user, err := d.Client().Create().
		SetNickName(user.NickName).
		SetRealName(user.RealName).
		SetEmail(user.Email).
		SetTel(user.Tel).
		SetPassword(user.Password).
		SetCreateTime(time.Now().Unix()).
		SetUpdateTime(time.Now().Unix()).Save(ctx)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (d *UserDao) GetUserByID(ctx context.Context, id int) (*ent2.User, error) {
	user, err := d.Client().Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UserDao) GetUserByTel(ctx context.Context, tel string) (*ent2.User, error) {
	user, err := d.Client().Query().
		Where(
			user.TelEQ(tel),
			user.IsDeletedEQ(0),
		).
		First(ctx)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (d *UserDao) UpdateUser(ctx context.Context, user *ent2.User) error {
	_, err := d.Client().UpdateOneID(user.ID).
		SetNickName(user.NickName).
		SetRealName(user.RealName).
		SetEmail(user.Email).
		SetTel(user.Tel).
		SetPassword(user.Password).
		SetUpdateTime(time.Now().Unix()).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *UserDao) GetFollowers(ctx context.Context, user *ent2.User) ([]*ent2.User, error) {
	users, err := user.QueryFollowers().All(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (d *UserDao) GetFollowing(ctx context.Context, user *ent2.User) ([]*ent2.User, error) {
	users, err := user.QueryFollowing().All(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (d *UserDao) FollowUser(ctx context.Context, userID, followingID int) error {
	_, err := d.Client().Update().Where(
		user.IDEQ(userID),
	).AddFollowingIDs([]int{followingID}...).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}

func (d *UserDao) UnFollowUser(ctx context.Context, userID, followingID int) error {
	_, err := d.Client().Update().Where(
		user.IDEQ(userID),
	).RemoveFollowingIDs([]int{followingID}...).Save(ctx)
	if err != nil {
		return err
	}
	return nil
}
