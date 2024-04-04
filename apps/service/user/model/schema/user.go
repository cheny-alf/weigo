package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickName").Comment("nickname").Annotations(
			entsql.Annotation{Size: 64}),
		field.String("realName").Comment("realName").Annotations(
			entsql.Annotation{Size: 64}).Unique(),
		field.String("email").Comment("email").Annotations(
			entsql.Annotation{Size: 64}),
		field.String("tel").Comment("tel").Annotations(
			entsql.Annotation{Size: 12}).Unique(),
		field.String("password").Comment("password").Annotations(
			entsql.Annotation{Size: 20}),
		field.Int64("lastLoginTime").Comment("lastLoginTime").Default(0),
		field.String("lastIPAddr").Comment("lastIPAddr").Default(""),
		field.Int64("createTime").Comment("createTime"),
		field.Int64("updateTime").Comment("updateTime"),
		field.Int64("isDeleted").Comment("isDeleted 0:否，1:是").Default(0),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("following", User.Type).
			From("followers"),
	}
}
