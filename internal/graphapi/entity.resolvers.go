package graphapi

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.34

import (
	"context"

	"go.infratographer.com/ipam-api/internal/ent/generated"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipaddress"
	"go.infratographer.com/x/gidx"
)

// FindIPAddressByID is the resolver for the findIPAddressByID field.
func (r *entityResolver) FindIPAddressByID(ctx context.Context, id gidx.PrefixedID) (*generated.IPAddress, error) {
	return r.client.IPAddress.Get(ctx, id)
}

// FindIPAddressableByID is the resolver for the findIPAddressableByID field.
func (r *entityResolver) FindIPAddressableByID(ctx context.Context, id gidx.PrefixedID) (*IPAddressable, error) {
	addrs, err := r.client.IPAddress.Query().Where(ipaddress.NodeID(id)).All(ctx)
	if err != nil {
		return nil, err
	}
	return &IPAddressable{ID: id, IPAddresses: addrs}, nil
}

// FindIPBlockByID is the resolver for the findIPBlockByID field.
func (r *entityResolver) FindIPBlockByID(ctx context.Context, id gidx.PrefixedID) (*generated.IPBlock, error) {
	return r.client.IPBlock.Get(ctx, id)
}

// FindIPBlockTypeByID is the resolver for the findIPBlockTypeByID field.
func (r *entityResolver) FindIPBlockTypeByID(ctx context.Context, id gidx.PrefixedID) (*generated.IPBlockType, error) {
	return r.client.IPBlockType.Get(ctx, id)
}

// FindResourceOwnerByID is the resolver for the findResourceOwnerByID field.
func (r *entityResolver) FindResourceOwnerByID(ctx context.Context, id gidx.PrefixedID) (*ResourceOwner, error) {
	return &ResourceOwner{ID: id}, nil
}

// Entity returns EntityResolver implementation.
func (r *Resolver) Entity() EntityResolver { return &entityResolver{r} }

type entityResolver struct{ *Resolver }
