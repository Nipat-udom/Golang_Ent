package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// // Fields of the User.
// func (User) Fields() []ent.Field {
// 	return []ent.Field{
// 		field.UUID("id", uuid.UUID{}).
// 			Default(uuid.New).StorageKey("userid"),
// 		field.String("firstname").Default("Unknown"),
// 		field.String("lastname").Optional(),
// 		field.Int("age"),
// 		field.Time("updated_at").
// 			Default(time.Now).
// 			UpdateDefault(time.Now),
// 	}
// }

// // Edges of the User.
// func (User) Edges() []ent.Edge {
// 	return []ent.Edge{
// 		edge.To("card", Card.Type).
// 			Unique(),
// 		edge.To("pets", Pet.Type),
// 		edge.From("groups", Group.Type).
// 			Ref("users"),
// 	}
// }

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New).StorageKey("userid"),
		field.String("name").Default("Unknown"),
		field.Int("age"),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type),
		edge.To("card", Card.Type).Unique(),
		edge.From("groups", Group.Type).Ref("users"),
	}
}
