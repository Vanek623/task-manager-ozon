package storage

import (
	"context"
	"encoding/json"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/counters"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/google/uuid"
	"gitlab.ozon.dev/Vanek623/task-manager-system/external/task/models"
)

const (
	expiration int32 = 60

	listKey = "list"

	listCachedCount = 10
)

type memcached struct {
	storage iTaskStorage
	client  *memcache.Client
	cs      *counters.Counters
}

func (m *memcached) Add(ctx context.Context, t *models.Task) error {
	err := m.storage.Add(ctx, t)
	if err != nil {
		return err
	}

	encoded, err := t.MarshalBinary()
	err = m.client.Add(&memcache.Item{
		Key:        t.ID.String(),
		Value:      encoded,
		Flags:      0,
		Expiration: expiration,
	})

	if err != nil {
		log.Error(err)
	}

	return nil
}

func (m *memcached) Delete(ctx context.Context, ID *uuid.UUID) error {
	err := m.storage.Delete(ctx, ID)
	if err != nil {
		return err
	}

	if err = m.client.Delete(ID.String()); err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		log.Error(err)
	}

	return nil
}

func (m *memcached) Update(ctx context.Context, t *models.Task) error {
	err := m.storage.Update(ctx, t)
	if err != nil {
		return err
	}

	encoded, err := t.MarshalBinary()
	err = m.client.Set(&memcache.Item{
		Key:        t.ID.String(),
		Value:      encoded,
		Expiration: expiration,
	})

	if err != nil {
		log.Error(err)
	}

	return nil
}

func newMemcached(host string, storage iTaskStorage) (*memcached, error) {
	client := memcache.New(host)
	if err := client.Ping(); err != nil {
		return nil, err
	}

	return &memcached{
		storage: storage,
		client:  client,
	}, nil
}

func (m *memcached) List(ctx context.Context, limit, offset uint64) ([]*models.Task, error) {
	tasks := make([]*models.Task, 0, limit)

	item, err := m.client.Get(listKey)
	if err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		return nil, err
	}

	canBeCached := listCachedCount-(offset+limit) >= 0

	if err == nil && canBeCached {
		m.cs.Inc(counters.CacheHit)
		var tmp []*models.Task
		if err = json.Unmarshal(item.Value, &tmp); err != nil {
			return nil, err
		}

		for i := offset; i < offset+limit; i++ {
			tasks = append(tasks, tmp[i])
		}
	} else if err != nil && canBeCached {
		m.cs.Inc(counters.CacheMiss)

		tmp, err := m.storage.List(ctx, 0, listCachedCount)
		if err != nil {
			return nil, err
		}

		for i := offset; i < offset+limit; i++ {
			tasks = append(tasks, tmp[i])
		}
	} else {
		tasks, err = m.storage.List(ctx, limit, offset)
		if err != nil {
			return nil, err
		}
	}

	return tasks, nil
}

func (m *memcached) Get(ctx context.Context, ID *uuid.UUID) (*models.Task, error) {
	var t *models.Task
	item, err := m.client.Get(ID.String())

	if err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		return nil, err
	}

	if err == nil {
		m.cs.Inc(counters.CacheHit)
		t = &models.Task{}
		if err = t.UnmarshalBinary(item.Value); err != nil {
			return nil, err
		}
	} else {
		m.cs.Inc(counters.CacheMiss)
		t, err = m.storage.Get(ctx, ID)
		if err != nil {
			return nil, err
		}

		if err = m.addToCache(t); err != nil {
			log.Error(err)
		}
	}

	return t, nil
}

func (m *memcached) addToCache(t *models.Task) error {
	encoded, err := t.MarshalBinary()
	if err != nil {
		return err
	}

	err = m.client.Set(&memcache.Item{
		Key:        t.ID.String(),
		Value:      encoded,
		Expiration: expiration,
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *memcached) removeFromCache(ID *uuid.UUID) error {
	if err := m.client.Delete(ID.String()); err != nil && !errors.Is(err, memcache.ErrCacheMiss) {
		return err
	}

	return nil
}
