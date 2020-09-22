package data

import (
	set "github.com/deckarep/golang-set"
)

type DIFEntityType string

const (
	VM              DIFEntityType = "virtualMachine"
	CONTAINER       DIFEntityType = "container"
	APPLICATION     DIFEntityType = "application"
	SERVICE         DIFEntityType = "service"
	DATABASE_SERVER DIFEntityType = "databaseServer"
	BUSINESS_APP    DIFEntityType = "businessApplication"
	BUSINESS_TRANS  DIFEntityType = "businessTransaction"
)

var validDIFEntities = []interface{}{
	APPLICATION,
	BUSINESS_APP,
	BUSINESS_TRANS,
	DATABASE_SERVER,
	SERVICE,
	VM,
	CONTAINER,
}

type DIFMetricType string

const (
	RESPONSE_TIME    DIFMetricType = "responseTime"
	TRANSACTION      DIFMetricType = "transaction"
	CONNECTION       DIFMetricType = "connection"
	CPU              DIFMetricType = "cpu"
	MEMORY           DIFMetricType = "memory"
	THREADS          DIFMetricType = "threads"
	HEAP             DIFMetricType = "heap"
	COLLECTION_TIME  DIFMetricType = "collectionTime"
	DBMEM            DIFMetricType = "dbMem"
	DBCACHEHITRATE   DIFMetricType = "dbCacheHitRate"
	KPI              DIFMetricType = "kpi"
	CLUSTER          DIFMetricType = "cluster"
	IO               DIFMetricType = "io"
	NET_THROUPUT     DIFMetricType = "netThroughput"
	APPLICATION_COMM DIFMetricType = "application"
	TRANSACTION_LOG  DIFMetricType = "transactionLog"
)

var validDIFMetrics = []interface{}{
	COLLECTION_TIME,
	CONNECTION,
	CPU,
	DBCACHEHITRATE,
	DBMEM,
	HEAP,
	KPI,
	MEMORY,
	THREADS,
	RESPONSE_TIME,
	TRANSACTION,
}

var DIFEntities = set.NewSetFromSlice(validDIFEntities)

func IsValidDIFEntity(entity DIFEntityType) bool {
	return DIFEntities.Contains(entity)
}

var DIFMetrics = set.NewSetFromSlice(validDIFMetrics)

func IsValidDIFMetric(metric DIFMetricType) bool {
	return DIFMetrics.Contains(metric)
}
