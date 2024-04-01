package model

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema --target=./ent --feature sql/upsert  --feature sql/modifier --feature sql/execquery
