// +build go1.9

// Copyright 2019 Microsoft Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This code was auto-generated by:
// github.com/Azure/azure-sdk-for-go/tools/profileBuilder

package documentdb

import (
	"context"

	original "github.com/Azure/azure-sdk-for-go/services/cosmos-db/mgmt/2015-04-08/documentdb"
)

const (
	DefaultBaseURI = original.DefaultBaseURI
)

type ConflictResolutionMode = original.ConflictResolutionMode

const (
	Custom         ConflictResolutionMode = original.Custom
	LastWriterWins ConflictResolutionMode = original.LastWriterWins
)

type DataType = original.DataType

const (
	LineString   DataType = original.LineString
	MultiPolygon DataType = original.MultiPolygon
	Number       DataType = original.Number
	Point        DataType = original.Point
	Polygon      DataType = original.Polygon
	String       DataType = original.String
)

type DatabaseAccountKind = original.DatabaseAccountKind

const (
	GlobalDocumentDB DatabaseAccountKind = original.GlobalDocumentDB
	MongoDB          DatabaseAccountKind = original.MongoDB
	Parse            DatabaseAccountKind = original.Parse
)

type DatabaseAccountOfferType = original.DatabaseAccountOfferType

const (
	Standard DatabaseAccountOfferType = original.Standard
)

type DefaultConsistencyLevel = original.DefaultConsistencyLevel

const (
	BoundedStaleness DefaultConsistencyLevel = original.BoundedStaleness
	ConsistentPrefix DefaultConsistencyLevel = original.ConsistentPrefix
	Eventual         DefaultConsistencyLevel = original.Eventual
	Session          DefaultConsistencyLevel = original.Session
	Strong           DefaultConsistencyLevel = original.Strong
)

type IndexKind = original.IndexKind

const (
	Hash    IndexKind = original.Hash
	Range   IndexKind = original.Range
	Spatial IndexKind = original.Spatial
)

type IndexingMode = original.IndexingMode

const (
	Consistent IndexingMode = original.Consistent
	Lazy       IndexingMode = original.Lazy
	None       IndexingMode = original.None
)

type KeyKind = original.KeyKind

const (
	Primary           KeyKind = original.Primary
	PrimaryReadonly   KeyKind = original.PrimaryReadonly
	Secondary         KeyKind = original.Secondary
	SecondaryReadonly KeyKind = original.SecondaryReadonly
)

type PartitionKind = original.PartitionKind

const (
	PartitionKindHash  PartitionKind = original.PartitionKindHash
	PartitionKindRange PartitionKind = original.PartitionKindRange
)

type PrimaryAggregationType = original.PrimaryAggregationType

const (
	PrimaryAggregationTypeAverage   PrimaryAggregationType = original.PrimaryAggregationTypeAverage
	PrimaryAggregationTypeLast      PrimaryAggregationType = original.PrimaryAggregationTypeLast
	PrimaryAggregationTypeMaximum   PrimaryAggregationType = original.PrimaryAggregationTypeMaximum
	PrimaryAggregationTypeMinimimum PrimaryAggregationType = original.PrimaryAggregationTypeMinimimum
	PrimaryAggregationTypeNone      PrimaryAggregationType = original.PrimaryAggregationTypeNone
	PrimaryAggregationTypeTotal     PrimaryAggregationType = original.PrimaryAggregationTypeTotal
)

type UnitType = original.UnitType

const (
	Bytes          UnitType = original.Bytes
	BytesPerSecond UnitType = original.BytesPerSecond
	Count          UnitType = original.Count
	CountPerSecond UnitType = original.CountPerSecond
	Milliseconds   UnitType = original.Milliseconds
	Percent        UnitType = original.Percent
	Seconds        UnitType = original.Seconds
)

type BaseClient = original.BaseClient
type Capability = original.Capability
type CassandraKeyspace = original.CassandraKeyspace
type CassandraKeyspaceCreateUpdateParameters = original.CassandraKeyspaceCreateUpdateParameters
type CassandraKeyspaceCreateUpdateProperties = original.CassandraKeyspaceCreateUpdateProperties
type CassandraKeyspaceListResult = original.CassandraKeyspaceListResult
type CassandraKeyspaceProperties = original.CassandraKeyspaceProperties
type CassandraKeyspaceResource = original.CassandraKeyspaceResource
type CassandraPartitionKey = original.CassandraPartitionKey
type CassandraSchema = original.CassandraSchema
type CassandraTable = original.CassandraTable
type CassandraTableCreateUpdateParameters = original.CassandraTableCreateUpdateParameters
type CassandraTableCreateUpdateProperties = original.CassandraTableCreateUpdateProperties
type CassandraTableListResult = original.CassandraTableListResult
type CassandraTableProperties = original.CassandraTableProperties
type CassandraTableResource = original.CassandraTableResource
type ClusterKey = original.ClusterKey
type CollectionClient = original.CollectionClient
type CollectionPartitionClient = original.CollectionPartitionClient
type CollectionPartitionRegionClient = original.CollectionPartitionRegionClient
type CollectionRegionClient = original.CollectionRegionClient
type Column = original.Column
type ConflictResolutionPolicy = original.ConflictResolutionPolicy
type ConsistencyPolicy = original.ConsistencyPolicy
type Container = original.Container
type ContainerCreateUpdateParameters = original.ContainerCreateUpdateParameters
type ContainerCreateUpdateProperties = original.ContainerCreateUpdateProperties
type ContainerListResult = original.ContainerListResult
type ContainerPartitionKey = original.ContainerPartitionKey
type ContainerProperties = original.ContainerProperties
type ContainerResource = original.ContainerResource
type DatabaseAccount = original.DatabaseAccount
type DatabaseAccountConnectionString = original.DatabaseAccountConnectionString
type DatabaseAccountCreateUpdateParameters = original.DatabaseAccountCreateUpdateParameters
type DatabaseAccountCreateUpdateProperties = original.DatabaseAccountCreateUpdateProperties
type DatabaseAccountListConnectionStringsResult = original.DatabaseAccountListConnectionStringsResult
type DatabaseAccountListKeysResult = original.DatabaseAccountListKeysResult
type DatabaseAccountListReadOnlyKeysResult = original.DatabaseAccountListReadOnlyKeysResult
type DatabaseAccountPatchParameters = original.DatabaseAccountPatchParameters
type DatabaseAccountPatchProperties = original.DatabaseAccountPatchProperties
type DatabaseAccountProperties = original.DatabaseAccountProperties
type DatabaseAccountRegenerateKeyParameters = original.DatabaseAccountRegenerateKeyParameters
type DatabaseAccountRegionClient = original.DatabaseAccountRegionClient
type DatabaseAccountsClient = original.DatabaseAccountsClient
type DatabaseAccountsCreateOrUpdateFuture = original.DatabaseAccountsCreateOrUpdateFuture
type DatabaseAccountsCreateUpdateCassandraKeyspaceFuture = original.DatabaseAccountsCreateUpdateCassandraKeyspaceFuture
type DatabaseAccountsCreateUpdateCassandraTableFuture = original.DatabaseAccountsCreateUpdateCassandraTableFuture
type DatabaseAccountsCreateUpdateGremlinContainerFuture = original.DatabaseAccountsCreateUpdateGremlinContainerFuture
type DatabaseAccountsCreateUpdateGremlinDatabaseFuture = original.DatabaseAccountsCreateUpdateGremlinDatabaseFuture
type DatabaseAccountsCreateUpdateMongoCollectionFuture = original.DatabaseAccountsCreateUpdateMongoCollectionFuture
type DatabaseAccountsCreateUpdateMongoDatabaseFuture = original.DatabaseAccountsCreateUpdateMongoDatabaseFuture
type DatabaseAccountsCreateUpdateSQLContainerFuture = original.DatabaseAccountsCreateUpdateSQLContainerFuture
type DatabaseAccountsCreateUpdateSQLDatabaseFuture = original.DatabaseAccountsCreateUpdateSQLDatabaseFuture
type DatabaseAccountsCreateUpdateTableFuture = original.DatabaseAccountsCreateUpdateTableFuture
type DatabaseAccountsDeleteCassandraKeyspaceFuture = original.DatabaseAccountsDeleteCassandraKeyspaceFuture
type DatabaseAccountsDeleteCassandraTableFuture = original.DatabaseAccountsDeleteCassandraTableFuture
type DatabaseAccountsDeleteFuture = original.DatabaseAccountsDeleteFuture
type DatabaseAccountsDeleteGremlinContainerFuture = original.DatabaseAccountsDeleteGremlinContainerFuture
type DatabaseAccountsDeleteGremlinDatabaseFuture = original.DatabaseAccountsDeleteGremlinDatabaseFuture
type DatabaseAccountsDeleteMongoCollectionFuture = original.DatabaseAccountsDeleteMongoCollectionFuture
type DatabaseAccountsDeleteMongoDatabaseFuture = original.DatabaseAccountsDeleteMongoDatabaseFuture
type DatabaseAccountsDeleteSQLContainerFuture = original.DatabaseAccountsDeleteSQLContainerFuture
type DatabaseAccountsDeleteSQLDatabaseFuture = original.DatabaseAccountsDeleteSQLDatabaseFuture
type DatabaseAccountsDeleteTableFuture = original.DatabaseAccountsDeleteTableFuture
type DatabaseAccountsFailoverPriorityChangeFuture = original.DatabaseAccountsFailoverPriorityChangeFuture
type DatabaseAccountsListResult = original.DatabaseAccountsListResult
type DatabaseAccountsOfflineRegionFuture = original.DatabaseAccountsOfflineRegionFuture
type DatabaseAccountsOnlineRegionFuture = original.DatabaseAccountsOnlineRegionFuture
type DatabaseAccountsPatchFuture = original.DatabaseAccountsPatchFuture
type DatabaseAccountsRegenerateKeyFuture = original.DatabaseAccountsRegenerateKeyFuture
type DatabaseClient = original.DatabaseClient
type ErrorResponse = original.ErrorResponse
type ExcludedPath = original.ExcludedPath
type ExtendedResourceProperties = original.ExtendedResourceProperties
type FailoverPolicies = original.FailoverPolicies
type FailoverPolicy = original.FailoverPolicy
type GremlinDatabase = original.GremlinDatabase
type GremlinDatabaseCreateUpdateParameters = original.GremlinDatabaseCreateUpdateParameters
type GremlinDatabaseCreateUpdateProperties = original.GremlinDatabaseCreateUpdateProperties
type GremlinDatabaseListResult = original.GremlinDatabaseListResult
type GremlinDatabaseProperties = original.GremlinDatabaseProperties
type GremlinDatabaseResource = original.GremlinDatabaseResource
type IncludedPath = original.IncludedPath
type Indexes = original.Indexes
type IndexingPolicy = original.IndexingPolicy
type Location = original.Location
type Metric = original.Metric
type MetricAvailability = original.MetricAvailability
type MetricDefinition = original.MetricDefinition
type MetricDefinitionsListResult = original.MetricDefinitionsListResult
type MetricListResult = original.MetricListResult
type MetricName = original.MetricName
type MetricValue = original.MetricValue
type MongoCollection = original.MongoCollection
type MongoCollectionCreateUpdateParameters = original.MongoCollectionCreateUpdateParameters
type MongoCollectionCreateUpdateProperties = original.MongoCollectionCreateUpdateProperties
type MongoCollectionListResult = original.MongoCollectionListResult
type MongoCollectionProperties = original.MongoCollectionProperties
type MongoCollectionResource = original.MongoCollectionResource
type MongoDatabase = original.MongoDatabase
type MongoDatabaseCreateUpdateParameters = original.MongoDatabaseCreateUpdateParameters
type MongoDatabaseCreateUpdateProperties = original.MongoDatabaseCreateUpdateProperties
type MongoDatabaseListResult = original.MongoDatabaseListResult
type MongoDatabaseProperties = original.MongoDatabaseProperties
type MongoDatabaseResource = original.MongoDatabaseResource
type MongoIndex = original.MongoIndex
type MongoIndexKeys = original.MongoIndexKeys
type MongoIndexOptions = original.MongoIndexOptions
type Operation = original.Operation
type OperationDisplay = original.OperationDisplay
type OperationListResult = original.OperationListResult
type OperationListResultIterator = original.OperationListResultIterator
type OperationListResultPage = original.OperationListResultPage
type OperationsClient = original.OperationsClient
type PartitionKeyRangeIDClient = original.PartitionKeyRangeIDClient
type PartitionKeyRangeIDRegionClient = original.PartitionKeyRangeIDRegionClient
type PartitionMetric = original.PartitionMetric
type PartitionMetricListResult = original.PartitionMetricListResult
type PartitionUsage = original.PartitionUsage
type PartitionUsagesResult = original.PartitionUsagesResult
type PercentileClient = original.PercentileClient
type PercentileMetric = original.PercentileMetric
type PercentileMetricListResult = original.PercentileMetricListResult
type PercentileMetricValue = original.PercentileMetricValue
type PercentileSourceTargetClient = original.PercentileSourceTargetClient
type PercentileTargetClient = original.PercentileTargetClient
type RegionForOnlineOffline = original.RegionForOnlineOffline
type Resource = original.Resource
type SQLDatabase = original.SQLDatabase
type SQLDatabaseCreateUpdateParameters = original.SQLDatabaseCreateUpdateParameters
type SQLDatabaseCreateUpdateProperties = original.SQLDatabaseCreateUpdateProperties
type SQLDatabaseListResult = original.SQLDatabaseListResult
type SQLDatabaseProperties = original.SQLDatabaseProperties
type SQLDatabaseResource = original.SQLDatabaseResource
type Table = original.Table
type TableCreateUpdateParameters = original.TableCreateUpdateParameters
type TableCreateUpdateProperties = original.TableCreateUpdateProperties
type TableListResult = original.TableListResult
type TableProperties = original.TableProperties
type TableResource = original.TableResource
type UniqueKey = original.UniqueKey
type UniqueKeyPolicy = original.UniqueKeyPolicy
type Usage = original.Usage
type UsagesResult = original.UsagesResult
type VirtualNetworkRule = original.VirtualNetworkRule

func New(subscriptionID string) BaseClient {
	return original.New(subscriptionID)
}
func NewCollectionClient(subscriptionID string) CollectionClient {
	return original.NewCollectionClient(subscriptionID)
}
func NewCollectionClientWithBaseURI(baseURI string, subscriptionID string) CollectionClient {
	return original.NewCollectionClientWithBaseURI(baseURI, subscriptionID)
}
func NewCollectionPartitionClient(subscriptionID string) CollectionPartitionClient {
	return original.NewCollectionPartitionClient(subscriptionID)
}
func NewCollectionPartitionClientWithBaseURI(baseURI string, subscriptionID string) CollectionPartitionClient {
	return original.NewCollectionPartitionClientWithBaseURI(baseURI, subscriptionID)
}
func NewCollectionPartitionRegionClient(subscriptionID string) CollectionPartitionRegionClient {
	return original.NewCollectionPartitionRegionClient(subscriptionID)
}
func NewCollectionPartitionRegionClientWithBaseURI(baseURI string, subscriptionID string) CollectionPartitionRegionClient {
	return original.NewCollectionPartitionRegionClientWithBaseURI(baseURI, subscriptionID)
}
func NewCollectionRegionClient(subscriptionID string) CollectionRegionClient {
	return original.NewCollectionRegionClient(subscriptionID)
}
func NewCollectionRegionClientWithBaseURI(baseURI string, subscriptionID string) CollectionRegionClient {
	return original.NewCollectionRegionClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabaseAccountRegionClient(subscriptionID string) DatabaseAccountRegionClient {
	return original.NewDatabaseAccountRegionClient(subscriptionID)
}
func NewDatabaseAccountRegionClientWithBaseURI(baseURI string, subscriptionID string) DatabaseAccountRegionClient {
	return original.NewDatabaseAccountRegionClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabaseAccountsClient(subscriptionID string) DatabaseAccountsClient {
	return original.NewDatabaseAccountsClient(subscriptionID)
}
func NewDatabaseAccountsClientWithBaseURI(baseURI string, subscriptionID string) DatabaseAccountsClient {
	return original.NewDatabaseAccountsClientWithBaseURI(baseURI, subscriptionID)
}
func NewDatabaseClient(subscriptionID string) DatabaseClient {
	return original.NewDatabaseClient(subscriptionID)
}
func NewDatabaseClientWithBaseURI(baseURI string, subscriptionID string) DatabaseClient {
	return original.NewDatabaseClientWithBaseURI(baseURI, subscriptionID)
}
func NewOperationListResultIterator(page OperationListResultPage) OperationListResultIterator {
	return original.NewOperationListResultIterator(page)
}
func NewOperationListResultPage(getNextPage func(context.Context, OperationListResult) (OperationListResult, error)) OperationListResultPage {
	return original.NewOperationListResultPage(getNextPage)
}
func NewOperationsClient(subscriptionID string) OperationsClient {
	return original.NewOperationsClient(subscriptionID)
}
func NewOperationsClientWithBaseURI(baseURI string, subscriptionID string) OperationsClient {
	return original.NewOperationsClientWithBaseURI(baseURI, subscriptionID)
}
func NewPartitionKeyRangeIDClient(subscriptionID string) PartitionKeyRangeIDClient {
	return original.NewPartitionKeyRangeIDClient(subscriptionID)
}
func NewPartitionKeyRangeIDClientWithBaseURI(baseURI string, subscriptionID string) PartitionKeyRangeIDClient {
	return original.NewPartitionKeyRangeIDClientWithBaseURI(baseURI, subscriptionID)
}
func NewPartitionKeyRangeIDRegionClient(subscriptionID string) PartitionKeyRangeIDRegionClient {
	return original.NewPartitionKeyRangeIDRegionClient(subscriptionID)
}
func NewPartitionKeyRangeIDRegionClientWithBaseURI(baseURI string, subscriptionID string) PartitionKeyRangeIDRegionClient {
	return original.NewPartitionKeyRangeIDRegionClientWithBaseURI(baseURI, subscriptionID)
}
func NewPercentileClient(subscriptionID string) PercentileClient {
	return original.NewPercentileClient(subscriptionID)
}
func NewPercentileClientWithBaseURI(baseURI string, subscriptionID string) PercentileClient {
	return original.NewPercentileClientWithBaseURI(baseURI, subscriptionID)
}
func NewPercentileSourceTargetClient(subscriptionID string) PercentileSourceTargetClient {
	return original.NewPercentileSourceTargetClient(subscriptionID)
}
func NewPercentileSourceTargetClientWithBaseURI(baseURI string, subscriptionID string) PercentileSourceTargetClient {
	return original.NewPercentileSourceTargetClientWithBaseURI(baseURI, subscriptionID)
}
func NewPercentileTargetClient(subscriptionID string) PercentileTargetClient {
	return original.NewPercentileTargetClient(subscriptionID)
}
func NewPercentileTargetClientWithBaseURI(baseURI string, subscriptionID string) PercentileTargetClient {
	return original.NewPercentileTargetClientWithBaseURI(baseURI, subscriptionID)
}
func NewWithBaseURI(baseURI string, subscriptionID string) BaseClient {
	return original.NewWithBaseURI(baseURI, subscriptionID)
}
func PossibleConflictResolutionModeValues() []ConflictResolutionMode {
	return original.PossibleConflictResolutionModeValues()
}
func PossibleDataTypeValues() []DataType {
	return original.PossibleDataTypeValues()
}
func PossibleDatabaseAccountKindValues() []DatabaseAccountKind {
	return original.PossibleDatabaseAccountKindValues()
}
func PossibleDatabaseAccountOfferTypeValues() []DatabaseAccountOfferType {
	return original.PossibleDatabaseAccountOfferTypeValues()
}
func PossibleDefaultConsistencyLevelValues() []DefaultConsistencyLevel {
	return original.PossibleDefaultConsistencyLevelValues()
}
func PossibleIndexKindValues() []IndexKind {
	return original.PossibleIndexKindValues()
}
func PossibleIndexingModeValues() []IndexingMode {
	return original.PossibleIndexingModeValues()
}
func PossibleKeyKindValues() []KeyKind {
	return original.PossibleKeyKindValues()
}
func PossiblePartitionKindValues() []PartitionKind {
	return original.PossiblePartitionKindValues()
}
func PossiblePrimaryAggregationTypeValues() []PrimaryAggregationType {
	return original.PossiblePrimaryAggregationTypeValues()
}
func PossibleUnitTypeValues() []UnitType {
	return original.PossibleUnitTypeValues()
}
func UserAgent() string {
	return original.UserAgent() + " profiles/latest"
}
func Version() string {
	return original.Version()
}
