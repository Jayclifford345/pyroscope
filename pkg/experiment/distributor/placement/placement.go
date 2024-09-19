package placement

import (
	"github.com/grafana/dskit/ring"

	"github.com/grafana/pyroscope/pkg/iter"
)

// Key represents the distribution key.
type Key struct {
	TenantID    string
	DatasetName string

	Tenant      uint64
	Dataset     uint64
	Fingerprint uint64
}

type Strategy interface {
	// NumTenantShards returns the number of shards
	// for a tenant from n total.
	NumTenantShards(k Key, n int) (size int)
	// NumDatasetShards returns the number of shards
	// for a dataset from n total.
	NumDatasetShards(k Key, n int) (size int)
	// PickShard returns the shard index
	// for a given key from n total.
	PickShard(k Key, n int) (shard int)
}

var DefaultPlacement = defaultPlacement{}

type defaultPlacement struct{}

func (defaultPlacement) NumTenantShards(Key, int) int { return 0 }

func (defaultPlacement) NumDatasetShards(Key, int) int { return 2 }

func (defaultPlacement) PickShard(k Key, n int) int { return int(k.Fingerprint % uint64(n)) }

// Placement represents the placement for the given distribution key.
type Placement struct {
	// Note that the instances reference shared objects, and must not be modified.
	Instances iter.Iterator[*ring.InstanceDesc]
	Shard     uint32
}

// Next returns the next available location address.
func (p *Placement) Next() (*ring.InstanceDesc, bool) {
	for p.Instances.Next() {
		if instance := p.Instances.At(); instance.State == ring.ACTIVE {
			return instance, true
		}
	}
	return nil, false
}
