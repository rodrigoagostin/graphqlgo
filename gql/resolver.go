package gql

import (
	"graphgo/postgres"

	"github.com/graphql-go/graphql"
)

type Resolver struct {
	db *postgres.Db
}

func (r *Resolver) UserResolver(p graphql.ResolveParams) (interface{}, error) {
	name, ok := p.Args["name"].(string)

	if ok {
		users := r.db.GetUserNameByName(name)
		return users, nil
	}

	return nil, nil
}
