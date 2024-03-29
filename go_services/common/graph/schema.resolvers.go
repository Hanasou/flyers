package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/Hanasou/flyers/go_services/common/databases"
	"github.com/Hanasou/flyers/go_services/common/graph/model"
)

// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented: CreateTodo - createTodo"))
}

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	byteContents, err := r.Driver.GetAll("todo")
	if err != nil {
		log.Println("CreateTodo unable to get results")
		return nil, err
	}

	results := &[]*model.Todo{}
	err = json.Unmarshal(byteContents, results)
	if err != nil {
		log.Println("CreateTodo failed to unmarshal byte data into results")
		return nil, err
	}

	return *results, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver {
	driver, err := databases.NewDriver("json")
	if err != nil {
		log.Println("QueryResolver failed to initialize database driver")
		return nil
	}
	return &queryResolver{
		r,
		driver,
	}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct {
	*Resolver
	databases.Driver
}
