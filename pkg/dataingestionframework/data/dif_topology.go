package data

import (
	"github.com/alecthomas/jsonschema"
	"time"
)

//Data injection framework topology
type Topology struct {
	Version    string       `json:"version"`
	Updatetime int64        `json:"updateTime" jsonschema:"description=Epoch timestamp for when files are generated"`
	Scope      string       `json:"scope"`
	Source     string       `json:"source,omitempty"`
	Entities   []*DIFEntity `json:"topology" jsonschema:"description=List of topology entities"`
}

func GenerateJSONSchema() *jsonschema.Schema {
	jsonschema.Version = "http://json-schema.org/draft-07/schema#"
	return jsonschema.Reflect(&Topology{})
}

func NewTopology() *Topology {
	return &Topology{
		Version:    "v1",
		Updatetime: 0,
		Entities:   []*DIFEntity{},
		Scope:      "",
	}
}

func (r *Topology) SetUpdateTime() *Topology {
	t := time.Now()
	r.Updatetime = t.Unix()
	return r
}

func (r *Topology) AddEntity(entity *DIFEntity) {
	r.Entities = append(r.Entities, entity)
}

func (r *Topology) AddEntities(entities []*DIFEntity) {
	r.Entities = append(r.Entities, entities...)
}
