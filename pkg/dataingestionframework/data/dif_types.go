package data

import (
	"fmt"
	set "github.com/deckarep/golang-set"
)

//Data ingestion framework topology entity
type DIFEntity struct {
	UID                 string                            `json:"uniqueId" jsonschema:"description=Entity ID"`
	Type                DIFEntityType                     `json:"type" jsonschema:"description=Entity type,enum=businessApplication,enum=businessTransaction,enum=service,enum=application,enum=databaseServer,enum=virtualMachine,enum=container"`
	Name                string                            `json:"name" jsonschema:"description=Entity display name"`
	HostedOn            *DIFHostedOn                      `json:"hostedOn,omitempty" jsonschema:"description=Attributes used to find the entity that hosts this entity"`
	MatchingIdentifiers *DIFMatchingIdentifiers           `json:"matchIdentifiers,omitempty" jsonschema:"description=Attributes used to find the entity that matches this entity"`
	PartOf              []*DIFPartOf                      `json:"partOf,omitempty" jsonschema:"description=Attributes used to find all the entities that this entity is part of"`
	Metrics             map[DIFMetricType][]*DIFMetricVal `json:"metrics,omitempty" jsonschema:"description=Metrics and values for the entity"`
	namespace           string
	partOfSet           set.Set
	hostTypeSet         set.Set
}

type DIFMatchingIdentifiers struct {
	IPAddress string `json:"ipAddress" jsonschema:"description=IP Address of the entity used to find the matching entity"`
}

type DIFHostedOn struct {
	HostType  []DIFEntityType `json:"hostType" jsonschema:"description=Entity type of the provider of this entity,enum=container,enum=virtualMachine"`
	IPAddress string          `json:"ipAddress" jsonschema:"description=IP Address of the host entity"`
	HostUuid  string          `json:"hostUuid" jsonschema:"description=Unique identifier for the host entity"`
}

type DIFPartOf struct {
	ParentEntity DIFEntityType `json:"entity" jsonschema:"description=Entity type of the parent entity,enum=businessApplication,enum=businessTransaction,enum=service,enum=application,enum=databaseServer"`
	UniqueId     string        `json:"uniqueId"`
	Label        string        `json:"label,omitempty"`
}

func NewDIFEntity(uid string, eType DIFEntityType) *DIFEntity {
	return &DIFEntity{
		UID:         uid,
		Type:        eType,
		Name:        uid,
		partOfSet:   set.NewSet(),
		hostTypeSet: set.NewSet(),
		Metrics:     make(map[DIFMetricType][]*DIFMetricVal),
	}
}

func (e *DIFEntity) WithName(name string) *DIFEntity {
	e.Name = name
	return e
}

func (e *DIFEntity) WithNamespace(namespace string) *DIFEntity {
	e.namespace = namespace
	return e
}

func (e *DIFEntity) GetNamespace() string {
	return e.namespace
}

func (e *DIFEntity) PartOfEntity(entity DIFEntityType, id, label string) *DIFEntity {
	if e.partOfSet.Contains(id) {
		return e
	}
	e.partOfSet.Add(id)
	e.PartOf = append(e.PartOf, &DIFPartOf{entity, id, label})
	return e
}

func (e *DIFEntity) HostedOnType(hostType DIFEntityType) *DIFEntity {
	if e.hostTypeSet.Contains(hostType) {
		return e
	}
	if e.HostedOn == nil {
		e.HostedOn = &DIFHostedOn{}
	}
	e.HostedOn.HostType = append(e.HostedOn.HostType, hostType)
	e.hostTypeSet.Add(hostType)
	return e
}

func (e *DIFEntity) GetHostedOnType() []DIFEntityType {
	var hostTypes []DIFEntityType
	for _, hostType := range e.hostTypeSet.ToSlice() {
		hostTypes = append(hostTypes, hostType.(DIFEntityType))
	}
	return hostTypes
}

func (e *DIFEntity) HostedOnIP(ip string) *DIFEntity {
	if e.HostedOn == nil {
		e.HostedOn = &DIFHostedOn{}
	}
	e.HostedOn.IPAddress = ip
	return e
}

func (e *DIFEntity) HostedOnUID(uid string) *DIFEntity {
	if e.HostedOn == nil {
		e.HostedOn = &DIFHostedOn{}
	}
	e.HostedOn.HostUuid = uid
	return e
}

func (e *DIFEntity) Matching(id string) *DIFEntity {
	if e.MatchingIdentifiers == nil {
		e.MatchingIdentifiers = &DIFMatchingIdentifiers{id}
		return e
	}
	// Overwrite
	e.MatchingIdentifiers.IPAddress = id
	return e
}

func (e *DIFEntity) AddMetric(metricType DIFMetricType, kind DIFMetricValKind, value float64, key string) {
	meList, found := e.Metrics[metricType]
	if !found {
		meList = append(meList, &DIFMetricVal{})
		e.Metrics[metricType] = meList
	}
	if len(meList) < 1 {
		return
	}
	if kind == AVERAGE {
		meList[0].Average = &value
	} else if kind == CAPACITY {
		meList[0].Capacity = &value
	}
	if key != "" {
		meList[0].Key = &key
	}
}

func (e *DIFEntity) AddMetrics(metricType DIFMetricType, metricVals []*DIFMetricVal) {
	e.Metrics[metricType] = append(e.Metrics[metricType], metricVals...)
}

func (e *DIFEntity) String() string {
	s := fmt.Sprintf("%s[%s:%s]", e.Type, e.UID, e.Name)
	if e.MatchingIdentifiers != nil {
		s += fmt.Sprintf(" IP[%s]", e.MatchingIdentifiers.IPAddress)
	}
	if e.PartOf != nil {
		s += fmt.Sprintf(" PartOf")
		for _, partOf := range e.PartOf {
			s += fmt.Sprintf("[%s:%s]", partOf.ParentEntity, partOf.UniqueId)
		}
	}
	if e.HostedOn != nil {
		s += fmt.Sprintf(" HostedOn")
		s += fmt.Sprintf("[%s:%s]",
			e.HostedOn.HostUuid, e.HostedOn.IPAddress)
	}
	for metricName, metricList := range e.Metrics {
		for _, metric := range metricList {
			s += fmt.Sprintf(" Metric %s:[%v]", metricName, metric)
		}
	}
	return s
}
