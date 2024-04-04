// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"wego/apps/service/user/model/ent/predicate"
	"wego/apps/service/user/model/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetNickName sets the "nickName" field.
func (uu *UserUpdate) SetNickName(s string) *UserUpdate {
	uu.mutation.SetNickName(s)
	return uu
}

// SetNillableNickName sets the "nickName" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNickName(s *string) *UserUpdate {
	if s != nil {
		uu.SetNickName(*s)
	}
	return uu
}

// SetRealName sets the "realName" field.
func (uu *UserUpdate) SetRealName(s string) *UserUpdate {
	uu.mutation.SetRealName(s)
	return uu
}

// SetNillableRealName sets the "realName" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRealName(s *string) *UserUpdate {
	if s != nil {
		uu.SetRealName(*s)
	}
	return uu
}

// SetEmail sets the "email" field.
func (uu *UserUpdate) SetEmail(s string) *UserUpdate {
	uu.mutation.SetEmail(s)
	return uu
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uu *UserUpdate) SetNillableEmail(s *string) *UserUpdate {
	if s != nil {
		uu.SetEmail(*s)
	}
	return uu
}

// SetTel sets the "tel" field.
func (uu *UserUpdate) SetTel(s string) *UserUpdate {
	uu.mutation.SetTel(s)
	return uu
}

// SetNillableTel sets the "tel" field if the given value is not nil.
func (uu *UserUpdate) SetNillableTel(s *string) *UserUpdate {
	if s != nil {
		uu.SetTel(*s)
	}
	return uu
}

// SetPassword sets the "password" field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.mutation.SetPassword(s)
	return uu
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// SetLastLoginTime sets the "lastLoginTime" field.
func (uu *UserUpdate) SetLastLoginTime(i int64) *UserUpdate {
	uu.mutation.ResetLastLoginTime()
	uu.mutation.SetLastLoginTime(i)
	return uu
}

// SetNillableLastLoginTime sets the "lastLoginTime" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastLoginTime(i *int64) *UserUpdate {
	if i != nil {
		uu.SetLastLoginTime(*i)
	}
	return uu
}

// AddLastLoginTime adds i to the "lastLoginTime" field.
func (uu *UserUpdate) AddLastLoginTime(i int64) *UserUpdate {
	uu.mutation.AddLastLoginTime(i)
	return uu
}

// SetLastIPAddr sets the "lastIPAddr" field.
func (uu *UserUpdate) SetLastIPAddr(s string) *UserUpdate {
	uu.mutation.SetLastIPAddr(s)
	return uu
}

// SetNillableLastIPAddr sets the "lastIPAddr" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastIPAddr(s *string) *UserUpdate {
	if s != nil {
		uu.SetLastIPAddr(*s)
	}
	return uu
}

// SetCreateTime sets the "createTime" field.
func (uu *UserUpdate) SetCreateTime(i int64) *UserUpdate {
	uu.mutation.ResetCreateTime()
	uu.mutation.SetCreateTime(i)
	return uu
}

// SetNillableCreateTime sets the "createTime" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCreateTime(i *int64) *UserUpdate {
	if i != nil {
		uu.SetCreateTime(*i)
	}
	return uu
}

// AddCreateTime adds i to the "createTime" field.
func (uu *UserUpdate) AddCreateTime(i int64) *UserUpdate {
	uu.mutation.AddCreateTime(i)
	return uu
}

// SetUpdateTime sets the "updateTime" field.
func (uu *UserUpdate) SetUpdateTime(i int64) *UserUpdate {
	uu.mutation.ResetUpdateTime()
	uu.mutation.SetUpdateTime(i)
	return uu
}

// SetNillableUpdateTime sets the "updateTime" field if the given value is not nil.
func (uu *UserUpdate) SetNillableUpdateTime(i *int64) *UserUpdate {
	if i != nil {
		uu.SetUpdateTime(*i)
	}
	return uu
}

// AddUpdateTime adds i to the "updateTime" field.
func (uu *UserUpdate) AddUpdateTime(i int64) *UserUpdate {
	uu.mutation.AddUpdateTime(i)
	return uu
}

// SetIsDeleted sets the "isDeleted" field.
func (uu *UserUpdate) SetIsDeleted(i int64) *UserUpdate {
	uu.mutation.ResetIsDeleted()
	uu.mutation.SetIsDeleted(i)
	return uu
}

// SetNillableIsDeleted sets the "isDeleted" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsDeleted(i *int64) *UserUpdate {
	if i != nil {
		uu.SetIsDeleted(*i)
	}
	return uu
}

// AddIsDeleted adds i to the "isDeleted" field.
func (uu *UserUpdate) AddIsDeleted(i int64) *UserUpdate {
	uu.mutation.AddIsDeleted(i)
	return uu
}

// AddFollowerIDs adds the "followers" edge to the User entity by IDs.
func (uu *UserUpdate) AddFollowerIDs(ids ...int) *UserUpdate {
	uu.mutation.AddFollowerIDs(ids...)
	return uu
}

// AddFollowers adds the "followers" edges to the User entity.
func (uu *UserUpdate) AddFollowers(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the "following" edge to the User entity by IDs.
func (uu *UserUpdate) AddFollowingIDs(ids ...int) *UserUpdate {
	uu.mutation.AddFollowingIDs(ids...)
	return uu
}

// AddFollowing adds the "following" edges to the User entity.
func (uu *UserUpdate) AddFollowing(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowingIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearFollowers clears all "followers" edges to the User entity.
func (uu *UserUpdate) ClearFollowers() *UserUpdate {
	uu.mutation.ClearFollowers()
	return uu
}

// RemoveFollowerIDs removes the "followers" edge to User entities by IDs.
func (uu *UserUpdate) RemoveFollowerIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveFollowerIDs(ids...)
	return uu
}

// RemoveFollowers removes "followers" edges to User entities.
func (uu *UserUpdate) RemoveFollowers(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowerIDs(ids...)
}

// ClearFollowing clears all "following" edges to the User entity.
func (uu *UserUpdate) ClearFollowing() *UserUpdate {
	uu.mutation.ClearFollowing()
	return uu
}

// RemoveFollowingIDs removes the "following" edge to User entities by IDs.
func (uu *UserUpdate) RemoveFollowingIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveFollowingIDs(ids...)
	return uu
}

// RemoveFollowing removes "following" edges to User entities.
func (uu *UserUpdate) RemoveFollowing(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowingIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uu *UserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdate {
	uu.modifiers = append(uu.modifiers, modifiers...)
	return uu
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.NickName(); ok {
		_spec.SetField(user.FieldNickName, field.TypeString, value)
	}
	if value, ok := uu.mutation.RealName(); ok {
		_spec.SetField(user.FieldRealName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uu.mutation.Tel(); ok {
		_spec.SetField(user.FieldTel, field.TypeString, value)
	}
	if value, ok := uu.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastLoginTime(); ok {
		_spec.SetField(user.FieldLastLoginTime, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedLastLoginTime(); ok {
		_spec.AddField(user.FieldLastLoginTime, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.LastIPAddr(); ok {
		_spec.SetField(user.FieldLastIPAddr, field.TypeString, value)
	}
	if value, ok := uu.mutation.CreateTime(); ok {
		_spec.SetField(user.FieldCreateTime, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedCreateTime(); ok {
		_spec.AddField(user.FieldCreateTime, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedUpdateTime(); ok {
		_spec.AddField(user.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.IsDeleted(); ok {
		_spec.SetField(user.FieldIsDeleted, field.TypeInt64, value)
	}
	if value, ok := uu.mutation.AddedIsDeleted(); ok {
		_spec.AddField(user.FieldIsDeleted, field.TypeInt64, value)
	}
	if uu.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFollowersIDs(); len(nodes) > 0 && !uu.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FollowersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedFollowingIDs(); len(nodes) > 0 && !uu.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.FollowingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(uu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetNickName sets the "nickName" field.
func (uuo *UserUpdateOne) SetNickName(s string) *UserUpdateOne {
	uuo.mutation.SetNickName(s)
	return uuo
}

// SetNillableNickName sets the "nickName" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNickName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNickName(*s)
	}
	return uuo
}

// SetRealName sets the "realName" field.
func (uuo *UserUpdateOne) SetRealName(s string) *UserUpdateOne {
	uuo.mutation.SetRealName(s)
	return uuo
}

// SetNillableRealName sets the "realName" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRealName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetRealName(*s)
	}
	return uuo
}

// SetEmail sets the "email" field.
func (uuo *UserUpdateOne) SetEmail(s string) *UserUpdateOne {
	uuo.mutation.SetEmail(s)
	return uuo
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableEmail(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetEmail(*s)
	}
	return uuo
}

// SetTel sets the "tel" field.
func (uuo *UserUpdateOne) SetTel(s string) *UserUpdateOne {
	uuo.mutation.SetTel(s)
	return uuo
}

// SetNillableTel sets the "tel" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTel(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTel(*s)
	}
	return uuo
}

// SetPassword sets the "password" field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.mutation.SetPassword(s)
	return uuo
}

// SetNillablePassword sets the "password" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// SetLastLoginTime sets the "lastLoginTime" field.
func (uuo *UserUpdateOne) SetLastLoginTime(i int64) *UserUpdateOne {
	uuo.mutation.ResetLastLoginTime()
	uuo.mutation.SetLastLoginTime(i)
	return uuo
}

// SetNillableLastLoginTime sets the "lastLoginTime" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastLoginTime(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetLastLoginTime(*i)
	}
	return uuo
}

// AddLastLoginTime adds i to the "lastLoginTime" field.
func (uuo *UserUpdateOne) AddLastLoginTime(i int64) *UserUpdateOne {
	uuo.mutation.AddLastLoginTime(i)
	return uuo
}

// SetLastIPAddr sets the "lastIPAddr" field.
func (uuo *UserUpdateOne) SetLastIPAddr(s string) *UserUpdateOne {
	uuo.mutation.SetLastIPAddr(s)
	return uuo
}

// SetNillableLastIPAddr sets the "lastIPAddr" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastIPAddr(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLastIPAddr(*s)
	}
	return uuo
}

// SetCreateTime sets the "createTime" field.
func (uuo *UserUpdateOne) SetCreateTime(i int64) *UserUpdateOne {
	uuo.mutation.ResetCreateTime()
	uuo.mutation.SetCreateTime(i)
	return uuo
}

// SetNillableCreateTime sets the "createTime" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCreateTime(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetCreateTime(*i)
	}
	return uuo
}

// AddCreateTime adds i to the "createTime" field.
func (uuo *UserUpdateOne) AddCreateTime(i int64) *UserUpdateOne {
	uuo.mutation.AddCreateTime(i)
	return uuo
}

// SetUpdateTime sets the "updateTime" field.
func (uuo *UserUpdateOne) SetUpdateTime(i int64) *UserUpdateOne {
	uuo.mutation.ResetUpdateTime()
	uuo.mutation.SetUpdateTime(i)
	return uuo
}

// SetNillableUpdateTime sets the "updateTime" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableUpdateTime(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetUpdateTime(*i)
	}
	return uuo
}

// AddUpdateTime adds i to the "updateTime" field.
func (uuo *UserUpdateOne) AddUpdateTime(i int64) *UserUpdateOne {
	uuo.mutation.AddUpdateTime(i)
	return uuo
}

// SetIsDeleted sets the "isDeleted" field.
func (uuo *UserUpdateOne) SetIsDeleted(i int64) *UserUpdateOne {
	uuo.mutation.ResetIsDeleted()
	uuo.mutation.SetIsDeleted(i)
	return uuo
}

// SetNillableIsDeleted sets the "isDeleted" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsDeleted(i *int64) *UserUpdateOne {
	if i != nil {
		uuo.SetIsDeleted(*i)
	}
	return uuo
}

// AddIsDeleted adds i to the "isDeleted" field.
func (uuo *UserUpdateOne) AddIsDeleted(i int64) *UserUpdateOne {
	uuo.mutation.AddIsDeleted(i)
	return uuo
}

// AddFollowerIDs adds the "followers" edge to the User entity by IDs.
func (uuo *UserUpdateOne) AddFollowerIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddFollowerIDs(ids...)
	return uuo
}

// AddFollowers adds the "followers" edges to the User entity.
func (uuo *UserUpdateOne) AddFollowers(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the "following" edge to the User entity by IDs.
func (uuo *UserUpdateOne) AddFollowingIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddFollowingIDs(ids...)
	return uuo
}

// AddFollowing adds the "following" edges to the User entity.
func (uuo *UserUpdateOne) AddFollowing(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowingIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearFollowers clears all "followers" edges to the User entity.
func (uuo *UserUpdateOne) ClearFollowers() *UserUpdateOne {
	uuo.mutation.ClearFollowers()
	return uuo
}

// RemoveFollowerIDs removes the "followers" edge to User entities by IDs.
func (uuo *UserUpdateOne) RemoveFollowerIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveFollowerIDs(ids...)
	return uuo
}

// RemoveFollowers removes "followers" edges to User entities.
func (uuo *UserUpdateOne) RemoveFollowers(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowerIDs(ids...)
}

// ClearFollowing clears all "following" edges to the User entity.
func (uuo *UserUpdateOne) ClearFollowing() *UserUpdateOne {
	uuo.mutation.ClearFollowing()
	return uuo
}

// RemoveFollowingIDs removes the "following" edge to User entities by IDs.
func (uuo *UserUpdateOne) RemoveFollowingIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveFollowingIDs(ids...)
	return uuo
}

// RemoveFollowing removes "following" edges to User entities.
func (uuo *UserUpdateOne) RemoveFollowing(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowingIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uuo *UserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdateOne {
	uuo.modifiers = append(uuo.modifiers, modifiers...)
	return uuo
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.NickName(); ok {
		_spec.SetField(user.FieldNickName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.RealName(); ok {
		_spec.SetField(user.FieldRealName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Email(); ok {
		_spec.SetField(user.FieldEmail, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Tel(); ok {
		_spec.SetField(user.FieldTel, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Password(); ok {
		_spec.SetField(user.FieldPassword, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastLoginTime(); ok {
		_spec.SetField(user.FieldLastLoginTime, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedLastLoginTime(); ok {
		_spec.AddField(user.FieldLastLoginTime, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.LastIPAddr(); ok {
		_spec.SetField(user.FieldLastIPAddr, field.TypeString, value)
	}
	if value, ok := uuo.mutation.CreateTime(); ok {
		_spec.SetField(user.FieldCreateTime, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedCreateTime(); ok {
		_spec.AddField(user.FieldCreateTime, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.UpdateTime(); ok {
		_spec.SetField(user.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedUpdateTime(); ok {
		_spec.AddField(user.FieldUpdateTime, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.IsDeleted(); ok {
		_spec.SetField(user.FieldIsDeleted, field.TypeInt64, value)
	}
	if value, ok := uuo.mutation.AddedIsDeleted(); ok {
		_spec.AddField(user.FieldIsDeleted, field.TypeInt64, value)
	}
	if uuo.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFollowersIDs(); len(nodes) > 0 && !uuo.mutation.FollowersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FollowersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedFollowingIDs(); len(nodes) > 0 && !uuo.mutation.FollowingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.FollowingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(uuo.modifiers...)
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
