package after

import (
	"context"
	"sync"
)

type (
	Entity struct {
		ID   int64
		name string
	}

	EntityRepo interface {
		FindByID(ctx context.Context, ID int64) (*Entity, error)
		FindAllByIDs(ctx context.Context, ids []int64) (entities []*Entity)
	}

	entityRepo struct {
		// Put your Database and Cache here
	}
)

// NewEntityRepository create new repository
func NewEntityRepository() EntityRepo {
	return &entityRepo{}
}

func (r *entityRepo) FindAllByIDs(ctx context.Context, ids []int64) (entities []*Entity) {
	var wg sync.WaitGroup
	c := make(chan *Entity, len(ids))
	for _, id := range ids {
		wg.Add(1)
		go func(id int64) {
			defer wg.Done()

			entity, err := r.FindByID(ctx, id)
			if err != nil { // Ignore error
				return
			}
			c <- entity
		}(id)
	}
	wg.Wait()
	close(c)

	rs := map[int64]*Entity{}
	for t := range c {
		if t != nil {
			rs[t.ID] = t
		}
	}

	for _, id := range ids {
		if t, ok := rs[id]; ok {
			entities = append(entities, t)
		}
	}

	return
}

func (r *entityRepo) FindByID(ctx context.Context, id int64) (*Entity, error) {

	// Find Entity by Cache or Database
	var entity Entity
	return &entity, nil
}
