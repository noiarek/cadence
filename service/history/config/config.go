// Copyright (c) 2020 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package config

import (
	"time"

	"github.com/uber/cadence/common"
	"github.com/uber/cadence/common/config"
	"github.com/uber/cadence/common/definition"
	"github.com/uber/cadence/common/dynamicconfig"
	"github.com/uber/cadence/common/task"
)

// Config represents configuration for cadence-history service
type Config struct {
	NumberOfShards                  int
	RPS                             dynamicconfig.IntPropertyFn
	MaxIDLengthWarnLimit            dynamicconfig.IntPropertyFn
	DomainNameMaxLength             dynamicconfig.IntPropertyFnWithDomainFilter
	IdentityMaxLength               dynamicconfig.IntPropertyFnWithDomainFilter
	WorkflowIDMaxLength             dynamicconfig.IntPropertyFnWithDomainFilter
	SignalNameMaxLength             dynamicconfig.IntPropertyFnWithDomainFilter
	WorkflowTypeMaxLength           dynamicconfig.IntPropertyFnWithDomainFilter
	RequestIDMaxLength              dynamicconfig.IntPropertyFnWithDomainFilter
	TaskListNameMaxLength           dynamicconfig.IntPropertyFnWithDomainFilter
	ActivityIDMaxLength             dynamicconfig.IntPropertyFnWithDomainFilter
	ActivityTypeMaxLength           dynamicconfig.IntPropertyFnWithDomainFilter
	MarkerNameMaxLength             dynamicconfig.IntPropertyFnWithDomainFilter
	TimerIDMaxLength                dynamicconfig.IntPropertyFnWithDomainFilter
	PersistenceMaxQPS               dynamicconfig.IntPropertyFn
	PersistenceGlobalMaxQPS         dynamicconfig.IntPropertyFn
	EnableVisibilitySampling        dynamicconfig.BoolPropertyFn
	EnableReadFromClosedExecutionV2 dynamicconfig.BoolPropertyFn
	VisibilityOpenMaxQPS            dynamicconfig.IntPropertyFnWithDomainFilter
	VisibilityClosedMaxQPS          dynamicconfig.IntPropertyFnWithDomainFilter
	AdvancedVisibilityWritingMode   dynamicconfig.StringPropertyFn
	EmitShardDiffLog                dynamicconfig.BoolPropertyFn
	MaxAutoResetPoints              dynamicconfig.IntPropertyFnWithDomainFilter
	ThrottledLogRPS                 dynamicconfig.IntPropertyFn
	EnableStickyQuery               dynamicconfig.BoolPropertyFnWithDomainFilter
	ShutdownDrainDuration           dynamicconfig.DurationPropertyFn

	// HistoryCache settings
	// Change of these configs require shard restart
	HistoryCacheInitialSize dynamicconfig.IntPropertyFn
	HistoryCacheMaxSize     dynamicconfig.IntPropertyFn
	HistoryCacheTTL         dynamicconfig.DurationPropertyFn

	// EventsCache settings
	// Change of these configs require shard restart
	EventsCacheInitialCount       dynamicconfig.IntPropertyFn
	EventsCacheMaxCount           dynamicconfig.IntPropertyFn
	EventsCacheMaxSize            dynamicconfig.IntPropertyFn
	EventsCacheTTL                dynamicconfig.DurationPropertyFn
	EventsCacheGlobalEnable       dynamicconfig.BoolPropertyFn
	EventsCacheGlobalInitialCount dynamicconfig.IntPropertyFn
	EventsCacheGlobalMaxCount     dynamicconfig.IntPropertyFn

	// ShardController settings
	RangeSizeBits           uint
	AcquireShardInterval    dynamicconfig.DurationPropertyFn
	AcquireShardConcurrency dynamicconfig.IntPropertyFn

	// the artificial delay added to standby cluster's view of active cluster's time
	StandbyClusterDelay                  dynamicconfig.DurationPropertyFn
	StandbyTaskMissingEventsResendDelay  dynamicconfig.DurationPropertyFn
	StandbyTaskMissingEventsDiscardDelay dynamicconfig.DurationPropertyFn

	// Task process settings
	TaskProcessRPS                          dynamicconfig.IntPropertyFnWithDomainFilter
	TaskSchedulerType                       dynamicconfig.IntPropertyFn
	TaskSchedulerWorkerCount                dynamicconfig.IntPropertyFn
	TaskSchedulerShardWorkerCount           dynamicconfig.IntPropertyFn
	TaskSchedulerQueueSize                  dynamicconfig.IntPropertyFn
	TaskSchedulerShardQueueSize             dynamicconfig.IntPropertyFn
	TaskSchedulerDispatcherCount            dynamicconfig.IntPropertyFn
	TaskSchedulerRoundRobinWeights          dynamicconfig.MapPropertyFn
	TaskCriticalRetryCount                  dynamicconfig.IntPropertyFn
	ActiveTaskRedispatchInterval            dynamicconfig.DurationPropertyFn
	StandbyTaskRedispatchInterval           dynamicconfig.DurationPropertyFn
	TaskRedispatchIntervalJitterCoefficient dynamicconfig.FloatPropertyFn
	StandbyTaskReReplicationContextTimeout  dynamicconfig.DurationPropertyFnWithDomainIDFilter
	EnableDropStuckTaskByDomainID           dynamicconfig.BoolPropertyFnWithDomainIDFilter
	ResurrectionCheckMinDelay               dynamicconfig.DurationPropertyFnWithDomainFilter

	// QueueProcessor settings
	QueueProcessorEnableSplit                          dynamicconfig.BoolPropertyFn
	QueueProcessorSplitMaxLevel                        dynamicconfig.IntPropertyFn
	QueueProcessorEnableRandomSplitByDomainID          dynamicconfig.BoolPropertyFnWithDomainIDFilter
	QueueProcessorRandomSplitProbability               dynamicconfig.FloatPropertyFn
	QueueProcessorEnablePendingTaskSplitByDomainID     dynamicconfig.BoolPropertyFnWithDomainIDFilter
	QueueProcessorPendingTaskSplitThreshold            dynamicconfig.MapPropertyFn
	QueueProcessorEnableStuckTaskSplitByDomainID       dynamicconfig.BoolPropertyFnWithDomainIDFilter
	QueueProcessorStuckTaskSplitThreshold              dynamicconfig.MapPropertyFn
	QueueProcessorSplitLookAheadDurationByDomainID     dynamicconfig.DurationPropertyFnWithDomainIDFilter
	QueueProcessorPollBackoffInterval                  dynamicconfig.DurationPropertyFn
	QueueProcessorPollBackoffIntervalJitterCoefficient dynamicconfig.FloatPropertyFn
	QueueProcessorEnablePersistQueueStates             dynamicconfig.BoolPropertyFn
	QueueProcessorEnableLoadQueueStates                dynamicconfig.BoolPropertyFn

	// TimerQueueProcessor settings
	TimerTaskBatchSize                                dynamicconfig.IntPropertyFn
	TimerTaskDeleteBatchSize                          dynamicconfig.IntPropertyFn
	TimerProcessorGetFailureRetryCount                dynamicconfig.IntPropertyFn
	TimerProcessorCompleteTimerFailureRetryCount      dynamicconfig.IntPropertyFn
	TimerProcessorUpdateAckInterval                   dynamicconfig.DurationPropertyFn
	TimerProcessorUpdateAckIntervalJitterCoefficient  dynamicconfig.FloatPropertyFn
	TimerProcessorCompleteTimerInterval               dynamicconfig.DurationPropertyFn
	TimerProcessorFailoverMaxPollRPS                  dynamicconfig.IntPropertyFn
	TimerProcessorMaxPollRPS                          dynamicconfig.IntPropertyFn
	TimerProcessorMaxPollInterval                     dynamicconfig.DurationPropertyFn
	TimerProcessorMaxPollIntervalJitterCoefficient    dynamicconfig.FloatPropertyFn
	TimerProcessorSplitQueueInterval                  dynamicconfig.DurationPropertyFn
	TimerProcessorSplitQueueIntervalJitterCoefficient dynamicconfig.FloatPropertyFn
	TimerProcessorMaxRedispatchQueueSize              dynamicconfig.IntPropertyFn
	TimerProcessorMaxTimeShift                        dynamicconfig.DurationPropertyFn
	TimerProcessorHistoryArchivalSizeLimit            dynamicconfig.IntPropertyFn
	TimerProcessorArchivalTimeLimit                   dynamicconfig.DurationPropertyFn

	// TransferQueueProcessor settings
	TransferTaskBatchSize                                dynamicconfig.IntPropertyFn
	TransferTaskDeleteBatchSize                          dynamicconfig.IntPropertyFn
	TransferProcessorCompleteTransferFailureRetryCount   dynamicconfig.IntPropertyFn
	TransferProcessorFailoverMaxPollRPS                  dynamicconfig.IntPropertyFn
	TransferProcessorMaxPollRPS                          dynamicconfig.IntPropertyFn
	TransferProcessorMaxPollInterval                     dynamicconfig.DurationPropertyFn
	TransferProcessorMaxPollIntervalJitterCoefficient    dynamicconfig.FloatPropertyFn
	TransferProcessorSplitQueueInterval                  dynamicconfig.DurationPropertyFn
	TransferProcessorSplitQueueIntervalJitterCoefficient dynamicconfig.FloatPropertyFn
	TransferProcessorUpdateAckInterval                   dynamicconfig.DurationPropertyFn
	TransferProcessorUpdateAckIntervalJitterCoefficient  dynamicconfig.FloatPropertyFn
	TransferProcessorCompleteTransferInterval            dynamicconfig.DurationPropertyFn
	TransferProcessorMaxRedispatchQueueSize              dynamicconfig.IntPropertyFn
	TransferProcessorEnableValidator                     dynamicconfig.BoolPropertyFn
	TransferProcessorValidationInterval                  dynamicconfig.DurationPropertyFn
	TransferProcessorVisibilityArchivalTimeLimit         dynamicconfig.DurationPropertyFn

	// CrossClusterQueueProcessor settings
	CrossClusterTaskBatchSize                                     dynamicconfig.IntPropertyFn
	CrossClusterTaskDeleteBatchSize                               dynamicconfig.IntPropertyFn
	CrossClusterTaskFetchBatchSize                                dynamicconfig.IntPropertyFnWithShardIDFilter
	CrossClusterSourceProcessorMaxPollRPS                         dynamicconfig.IntPropertyFn
	CrossClusterSourceProcessorMaxPollInterval                    dynamicconfig.DurationPropertyFn
	CrossClusterSourceProcessorMaxPollIntervalJitterCoefficient   dynamicconfig.FloatPropertyFn
	CrossClusterSourceProcessorUpdateAckInterval                  dynamicconfig.DurationPropertyFn
	CrossClusterSourceProcessorUpdateAckIntervalJitterCoefficient dynamicconfig.FloatPropertyFn
	CrossClusterSourceProcessorMaxRedispatchQueueSize             dynamicconfig.IntPropertyFn
	CrossClusterSourceProcessorMaxPendingTaskSize                 dynamicconfig.IntPropertyFn

	// CrossClusterTargetTaskProcessor settings
	CrossClusterTargetProcessorMaxPendingTasks            dynamicconfig.IntPropertyFn
	CrossClusterTargetProcessorMaxRetryCount              dynamicconfig.IntPropertyFn
	CrossClusterTargetProcessorTaskWaitInterval           dynamicconfig.DurationPropertyFn
	CrossClusterTargetProcessorServiceBusyBackoffInterval dynamicconfig.DurationPropertyFn
	CrossClusterTargetProcessorJitterCoefficient          dynamicconfig.FloatPropertyFn

	// CrossClusterTaskFetcher settings
	CrossClusterFetcherParallelism                dynamicconfig.IntPropertyFn
	CrossClusterFetcherAggregationInterval        dynamicconfig.DurationPropertyFn
	CrossClusterFetcherServiceBusyBackoffInterval dynamicconfig.DurationPropertyFn
	CrossClusterFetcherErrorBackoffInterval       dynamicconfig.DurationPropertyFn
	CrossClusterFetcherJitterCoefficient          dynamicconfig.FloatPropertyFn

	// ReplicatorQueueProcessor settings
	ReplicatorTaskBatchSize                               dynamicconfig.IntPropertyFn
	ReplicatorTaskDeleteBatchSize                         dynamicconfig.IntPropertyFn
	ReplicatorTaskWorkerCount                             dynamicconfig.IntPropertyFn
	ReplicatorReadTaskMaxRetryCount                       dynamicconfig.IntPropertyFn
	ReplicatorProcessorMaxPollRPS                         dynamicconfig.IntPropertyFn
	ReplicatorProcessorMaxPollInterval                    dynamicconfig.DurationPropertyFn
	ReplicatorProcessorMaxPollIntervalJitterCoefficient   dynamicconfig.FloatPropertyFn
	ReplicatorProcessorUpdateAckInterval                  dynamicconfig.DurationPropertyFn
	ReplicatorProcessorUpdateAckIntervalJitterCoefficient dynamicconfig.FloatPropertyFn
	ReplicatorProcessorMaxRedispatchQueueSize             dynamicconfig.IntPropertyFn
	ReplicatorProcessorEnablePriorityTaskProcessor        dynamicconfig.BoolPropertyFn
	ReplicatorProcessorFetchTasksBatchSize                dynamicconfig.IntPropertyFnWithShardIDFilter
	ReplicatorUpperLatency                                dynamicconfig.DurationPropertyFn

	// Persistence settings
	ExecutionMgrNumConns dynamicconfig.IntPropertyFn
	HistoryMgrNumConns   dynamicconfig.IntPropertyFn

	// System Limits
	MaximumBufferedEventsBatch dynamicconfig.IntPropertyFn
	MaximumSignalsPerExecution dynamicconfig.IntPropertyFnWithDomainFilter

	// ShardUpdateMinInterval the minimal time interval which the shard info can be updated
	ShardUpdateMinInterval dynamicconfig.DurationPropertyFn
	// ShardSyncMinInterval the minimal time interval which the shard info should be sync to remote
	ShardSyncMinInterval            dynamicconfig.DurationPropertyFn
	ShardSyncTimerJitterCoefficient dynamicconfig.FloatPropertyFn

	// Time to hold a poll request before returning an empty response
	// right now only used by GetMutableState
	LongPollExpirationInterval dynamicconfig.DurationPropertyFnWithDomainFilter

	// encoding the history events
	EventEncodingType dynamicconfig.StringPropertyFnWithDomainFilter
	// whether or not using ParentClosePolicy
	EnableParentClosePolicy dynamicconfig.BoolPropertyFnWithDomainFilter
	// whether or not enable system workers for processing parent close policy task
	EnableParentClosePolicyWorker dynamicconfig.BoolPropertyFn
	// whether to enable system workers for processing ElasticSearch Analyzer
	EnableESAnalyzer dynamicconfig.BoolPropertyFn
	// parent close policy will be processed by sys workers(if enabled) if
	// the number of children greater than or equal to this threshold
	ParentClosePolicyThreshold dynamicconfig.IntPropertyFnWithDomainFilter
	// total number of parentClosePolicy system workflows
	NumParentClosePolicySystemWorkflows dynamicconfig.IntPropertyFn

	// Archival settings
	NumArchiveSystemWorkflows       dynamicconfig.IntPropertyFn
	ArchiveRequestRPS               dynamicconfig.IntPropertyFn
	AllowArchivingIncompleteHistory dynamicconfig.BoolPropertyFn

	// Size limit related settings
	BlobSizeLimitError     dynamicconfig.IntPropertyFnWithDomainFilter
	BlobSizeLimitWarn      dynamicconfig.IntPropertyFnWithDomainFilter
	HistorySizeLimitError  dynamicconfig.IntPropertyFnWithDomainFilter
	HistorySizeLimitWarn   dynamicconfig.IntPropertyFnWithDomainFilter
	HistoryCountLimitError dynamicconfig.IntPropertyFnWithDomainFilter
	HistoryCountLimitWarn  dynamicconfig.IntPropertyFnWithDomainFilter

	// ValidSearchAttributes is legal indexed keys that can be used in list APIs
	ValidSearchAttributes             dynamicconfig.MapPropertyFn
	SearchAttributesNumberOfKeysLimit dynamicconfig.IntPropertyFnWithDomainFilter
	SearchAttributesSizeOfValueLimit  dynamicconfig.IntPropertyFnWithDomainFilter
	SearchAttributesTotalSizeLimit    dynamicconfig.IntPropertyFnWithDomainFilter

	// Decision settings
	// StickyTTL is to expire a sticky tasklist if no update more than this duration
	// TODO https://github.com/uber/cadence/issues/2357
	StickyTTL dynamicconfig.DurationPropertyFnWithDomainFilter
	// DecisionHeartbeatTimeout is to timeout behavior of: RespondDecisionTaskComplete with ForceCreateNewDecisionTask == true without any decisions
	// So that decision will be scheduled to another worker(by clear stickyness)
	DecisionHeartbeatTimeout dynamicconfig.DurationPropertyFnWithDomainFilter
	// MaxDecisionStartToCloseSeconds is the StartToCloseSeconds for decision
	MaxDecisionStartToCloseSeconds           dynamicconfig.IntPropertyFnWithDomainFilter
	DecisionRetryCriticalAttempts            dynamicconfig.IntPropertyFn
	DecisionRetryMaxAttempts                 dynamicconfig.IntPropertyFnWithDomainFilter
	NormalDecisionScheduleToStartMaxAttempts dynamicconfig.IntPropertyFnWithDomainFilter
	NormalDecisionScheduleToStartTimeout     dynamicconfig.DurationPropertyFnWithDomainFilter

	// The following is used by the new RPC replication stack
	ReplicationTaskFetcherParallelism                  dynamicconfig.IntPropertyFn
	ReplicationTaskFetcherAggregationInterval          dynamicconfig.DurationPropertyFn
	ReplicationTaskFetcherTimerJitterCoefficient       dynamicconfig.FloatPropertyFn
	ReplicationTaskFetcherErrorRetryWait               dynamicconfig.DurationPropertyFn
	ReplicationTaskFetcherServiceBusyWait              dynamicconfig.DurationPropertyFn
	ReplicationTaskProcessorErrorRetryWait             dynamicconfig.DurationPropertyFnWithShardIDFilter
	ReplicationTaskProcessorErrorRetryMaxAttempts      dynamicconfig.IntPropertyFnWithShardIDFilter
	ReplicationTaskProcessorErrorSecondRetryWait       dynamicconfig.DurationPropertyFnWithShardIDFilter
	ReplicationTaskProcessorErrorSecondRetryMaxWait    dynamicconfig.DurationPropertyFnWithShardIDFilter
	ReplicationTaskProcessorErrorSecondRetryExpiration dynamicconfig.DurationPropertyFnWithShardIDFilter
	ReplicationTaskProcessorNoTaskRetryWait            dynamicconfig.DurationPropertyFnWithShardIDFilter
	ReplicationTaskProcessorCleanupInterval            dynamicconfig.DurationPropertyFnWithShardIDFilter
	ReplicationTaskProcessorCleanupJitterCoefficient   dynamicconfig.FloatPropertyFnWithShardIDFilter
	ReplicationTaskProcessorReadHistoryBatchSize       dynamicconfig.IntPropertyFn
	ReplicationTaskProcessorStartWait                  dynamicconfig.DurationPropertyFnWithShardIDFilter
	ReplicationTaskProcessorStartWaitJitterCoefficient dynamicconfig.FloatPropertyFnWithShardIDFilter
	ReplicationTaskProcessorHostQPS                    dynamicconfig.FloatPropertyFn
	ReplicationTaskProcessorShardQPS                   dynamicconfig.FloatPropertyFn
	ReplicationTaskGenerationQPS                       dynamicconfig.FloatPropertyFn
	EnableReplicationTaskGeneration                    dynamicconfig.BoolPropertyFnWithDomainIDAndWorkflowIDFilter

	// The following are used by consistent query
	EnableConsistentQuery         dynamicconfig.BoolPropertyFn
	EnableConsistentQueryByDomain dynamicconfig.BoolPropertyFnWithDomainFilter
	MaxBufferedQueryCount         dynamicconfig.IntPropertyFn

	EnableCrossClusterOperations dynamicconfig.BoolPropertyFnWithDomainFilter

	// Data integrity check related config knobs
	MutableStateChecksumGenProbability    dynamicconfig.IntPropertyFnWithDomainFilter
	MutableStateChecksumVerifyProbability dynamicconfig.IntPropertyFnWithDomainFilter
	MutableStateChecksumInvalidateBefore  dynamicconfig.FloatPropertyFn

	// Failover marker heartbeat
	NotifyFailoverMarkerInterval               dynamicconfig.DurationPropertyFn
	NotifyFailoverMarkerTimerJitterCoefficient dynamicconfig.FloatPropertyFn
	EnableGracefulFailover                     dynamicconfig.BoolPropertyFn

	// Allows worker to dispatch activity tasks through local tunnel after decisions are made. This is an performance optimization to skip activity scheduling efforts.
	EnableActivityLocalDispatchByDomain dynamicconfig.BoolPropertyFnWithDomainFilter

	ActivityMaxScheduleToStartTimeoutForRetry dynamicconfig.DurationPropertyFnWithDomainFilter

	// Debugging configurations
	EnableDebugMode             bool // note that this value is initialized once on service start
	EnableTaskInfoLogByDomainID dynamicconfig.BoolPropertyFnWithDomainIDFilter
}

const (
	// DefaultHistoryMaxAutoResetPoints is the default maximum number for auto reset points
	DefaultHistoryMaxAutoResetPoints = 20
)

var (
	// DefaultTaskPriorityWeight is the default round robin weight used by task scheduler
	DefaultTaskPriorityWeight = map[int]int{
		task.GetTaskPriority(task.HighPriorityClass, task.DefaultPrioritySubclass):    500,
		task.GetTaskPriority(task.DefaultPriorityClass, task.DefaultPrioritySubclass): 20,
		task.GetTaskPriority(task.LowPriorityClass, task.DefaultPrioritySubclass):     5,
	}

	// DefaultPendingTaskSplitThreshold is the default pending task split threshold
	DefaultPendingTaskSplitThreshold = map[int]int{ // processing queue level -> threshold for # of pending tasks per domain
		0: 1000,
		1: 10000,
	}

	// DefaultStuckTaskSplitThreshold is the default stuck task split threshold
	DefaultStuckTaskSplitThreshold = map[int]int{ // processing queue level -> threshold for # of task attempts
		0: 100,
		1: 10000,
	}
)

// New returns new service config with default values
func New(dc *dynamicconfig.Collection, numberOfShards int, storeType string, isAdvancedVisConfigExist bool) *Config {
	cfg := &Config{
		NumberOfShards:                       numberOfShards,
		RPS:                                  dc.GetIntProperty(dynamicconfig.HistoryRPS, 3000),
		MaxIDLengthWarnLimit:                 dc.GetIntProperty(dynamicconfig.MaxIDLengthWarnLimit, common.DefaultIDLengthWarnLimit),
		DomainNameMaxLength:                  dc.GetIntPropertyFilteredByDomain(dynamicconfig.DomainNameMaxLength, common.DefaultIDLengthErrorLimit),
		IdentityMaxLength:                    dc.GetIntPropertyFilteredByDomain(dynamicconfig.IdentityMaxLength, common.DefaultIDLengthErrorLimit),
		WorkflowIDMaxLength:                  dc.GetIntPropertyFilteredByDomain(dynamicconfig.WorkflowIDMaxLength, common.DefaultIDLengthErrorLimit),
		SignalNameMaxLength:                  dc.GetIntPropertyFilteredByDomain(dynamicconfig.SignalNameMaxLength, common.DefaultIDLengthErrorLimit),
		WorkflowTypeMaxLength:                dc.GetIntPropertyFilteredByDomain(dynamicconfig.WorkflowTypeMaxLength, common.DefaultIDLengthErrorLimit),
		RequestIDMaxLength:                   dc.GetIntPropertyFilteredByDomain(dynamicconfig.RequestIDMaxLength, common.DefaultIDLengthErrorLimit),
		TaskListNameMaxLength:                dc.GetIntPropertyFilteredByDomain(dynamicconfig.TaskListNameMaxLength, common.DefaultIDLengthErrorLimit),
		ActivityIDMaxLength:                  dc.GetIntPropertyFilteredByDomain(dynamicconfig.ActivityIDMaxLength, common.DefaultIDLengthErrorLimit),
		ActivityTypeMaxLength:                dc.GetIntPropertyFilteredByDomain(dynamicconfig.ActivityTypeMaxLength, common.DefaultIDLengthErrorLimit),
		MarkerNameMaxLength:                  dc.GetIntPropertyFilteredByDomain(dynamicconfig.MarkerNameMaxLength, common.DefaultIDLengthErrorLimit),
		TimerIDMaxLength:                     dc.GetIntPropertyFilteredByDomain(dynamicconfig.TimerIDMaxLength, common.DefaultIDLengthErrorLimit),
		PersistenceMaxQPS:                    dc.GetIntProperty(dynamicconfig.HistoryPersistenceMaxQPS, 9000),
		PersistenceGlobalMaxQPS:              dc.GetIntProperty(dynamicconfig.HistoryPersistenceGlobalMaxQPS, 0),
		ShutdownDrainDuration:                dc.GetDurationProperty(dynamicconfig.HistoryShutdownDrainDuration, 0),
		EnableVisibilitySampling:             dc.GetBoolProperty(dynamicconfig.EnableVisibilitySampling, false),
		EnableReadFromClosedExecutionV2:      dc.GetBoolProperty(dynamicconfig.EnableReadFromClosedExecutionV2, false),
		VisibilityOpenMaxQPS:                 dc.GetIntPropertyFilteredByDomain(dynamicconfig.HistoryVisibilityOpenMaxQPS, 300),
		VisibilityClosedMaxQPS:               dc.GetIntPropertyFilteredByDomain(dynamicconfig.HistoryVisibilityClosedMaxQPS, 300),
		MaxAutoResetPoints:                   dc.GetIntPropertyFilteredByDomain(dynamicconfig.HistoryMaxAutoResetPoints, DefaultHistoryMaxAutoResetPoints),
		MaxDecisionStartToCloseSeconds:       dc.GetIntPropertyFilteredByDomain(dynamicconfig.MaxDecisionStartToCloseSeconds, 240),
		AdvancedVisibilityWritingMode:        dc.GetStringProperty(dynamicconfig.AdvancedVisibilityWritingMode, common.GetDefaultAdvancedVisibilityWritingMode(isAdvancedVisConfigExist)),
		EmitShardDiffLog:                     dc.GetBoolProperty(dynamicconfig.EmitShardDiffLog, false),
		HistoryCacheInitialSize:              dc.GetIntProperty(dynamicconfig.HistoryCacheInitialSize, 128),
		HistoryCacheMaxSize:                  dc.GetIntProperty(dynamicconfig.HistoryCacheMaxSize, 512),
		HistoryCacheTTL:                      dc.GetDurationProperty(dynamicconfig.HistoryCacheTTL, time.Hour),
		EventsCacheInitialCount:              dc.GetIntProperty(dynamicconfig.EventsCacheInitialCount, 128),
		EventsCacheMaxCount:                  dc.GetIntProperty(dynamicconfig.EventsCacheMaxCount, 512),
		EventsCacheMaxSize:                   dc.GetIntProperty(dynamicconfig.EventsCacheMaxSize, 0),
		EventsCacheTTL:                       dc.GetDurationProperty(dynamicconfig.EventsCacheTTL, time.Hour),
		EventsCacheGlobalEnable:              dc.GetBoolProperty(dynamicconfig.EventsCacheGlobalEnable, false),
		EventsCacheGlobalInitialCount:        dc.GetIntProperty(dynamicconfig.EventsCacheGlobalInitialCount, 4096),
		EventsCacheGlobalMaxCount:            dc.GetIntProperty(dynamicconfig.EventsCacheGlobalMaxCount, 131072),
		RangeSizeBits:                        20, // 20 bits for sequencer, 2^20 sequence number for any range
		AcquireShardInterval:                 dc.GetDurationProperty(dynamicconfig.AcquireShardInterval, time.Minute),
		AcquireShardConcurrency:              dc.GetIntProperty(dynamicconfig.AcquireShardConcurrency, 1),
		StandbyClusterDelay:                  dc.GetDurationProperty(dynamicconfig.StandbyClusterDelay, 5*time.Minute),
		StandbyTaskMissingEventsResendDelay:  dc.GetDurationProperty(dynamicconfig.StandbyTaskMissingEventsResendDelay, 15*time.Minute),
		StandbyTaskMissingEventsDiscardDelay: dc.GetDurationProperty(dynamicconfig.StandbyTaskMissingEventsDiscardDelay, 25*time.Minute),

		TaskProcessRPS:                          dc.GetIntPropertyFilteredByDomain(dynamicconfig.TaskProcessRPS, 1000),
		TaskSchedulerType:                       dc.GetIntProperty(dynamicconfig.TaskSchedulerType, int(task.SchedulerTypeWRR)),
		TaskSchedulerWorkerCount:                dc.GetIntProperty(dynamicconfig.TaskSchedulerWorkerCount, 200),
		TaskSchedulerShardWorkerCount:           dc.GetIntProperty(dynamicconfig.TaskSchedulerShardWorkerCount, 0),
		TaskSchedulerQueueSize:                  dc.GetIntProperty(dynamicconfig.TaskSchedulerQueueSize, 10000),
		TaskSchedulerShardQueueSize:             dc.GetIntProperty(dynamicconfig.TaskSchedulerShardQueueSize, 200),
		TaskSchedulerDispatcherCount:            dc.GetIntProperty(dynamicconfig.TaskSchedulerDispatcherCount, 1),
		TaskSchedulerRoundRobinWeights:          dc.GetMapProperty(dynamicconfig.TaskSchedulerRoundRobinWeights, common.ConvertIntMapToDynamicConfigMapProperty(DefaultTaskPriorityWeight)),
		TaskCriticalRetryCount:                  dc.GetIntProperty(dynamicconfig.TaskCriticalRetryCount, 50),
		ActiveTaskRedispatchInterval:            dc.GetDurationProperty(dynamicconfig.ActiveTaskRedispatchInterval, 5*time.Second),
		StandbyTaskRedispatchInterval:           dc.GetDurationProperty(dynamicconfig.StandbyTaskRedispatchInterval, 30*time.Second),
		TaskRedispatchIntervalJitterCoefficient: dc.GetFloat64Property(dynamicconfig.TaskRedispatchIntervalJitterCoefficient, 0.15),
		StandbyTaskReReplicationContextTimeout:  dc.GetDurationPropertyFilteredByDomainID(dynamicconfig.StandbyTaskReReplicationContextTimeout, 3*time.Minute),
		EnableDropStuckTaskByDomainID:           dc.GetBoolPropertyFilteredByDomainID(dynamicconfig.EnableDropStuckTaskByDomainID, false),
		ResurrectionCheckMinDelay:               dc.GetDurationPropertyFilteredByDomain(dynamicconfig.ResurrectionCheckMinDelay, 24*time.Hour),

		QueueProcessorEnableSplit:                          dc.GetBoolProperty(dynamicconfig.QueueProcessorEnableSplit, false),
		QueueProcessorSplitMaxLevel:                        dc.GetIntProperty(dynamicconfig.QueueProcessorSplitMaxLevel, 2), // 3 levels, start from 0
		QueueProcessorEnableRandomSplitByDomainID:          dc.GetBoolPropertyFilteredByDomainID(dynamicconfig.QueueProcessorEnableRandomSplitByDomainID, false),
		QueueProcessorRandomSplitProbability:               dc.GetFloat64Property(dynamicconfig.QueueProcessorRandomSplitProbability, 0.01),
		QueueProcessorEnablePendingTaskSplitByDomainID:     dc.GetBoolPropertyFilteredByDomainID(dynamicconfig.QueueProcessorEnablePendingTaskSplitByDomainID, false),
		QueueProcessorPendingTaskSplitThreshold:            dc.GetMapProperty(dynamicconfig.QueueProcessorPendingTaskSplitThreshold, common.ConvertIntMapToDynamicConfigMapProperty(DefaultPendingTaskSplitThreshold)),
		QueueProcessorEnableStuckTaskSplitByDomainID:       dc.GetBoolPropertyFilteredByDomainID(dynamicconfig.QueueProcessorEnableStuckTaskSplitByDomainID, false),
		QueueProcessorStuckTaskSplitThreshold:              dc.GetMapProperty(dynamicconfig.QueueProcessorStuckTaskSplitThreshold, common.ConvertIntMapToDynamicConfigMapProperty(DefaultStuckTaskSplitThreshold)),
		QueueProcessorSplitLookAheadDurationByDomainID:     dc.GetDurationPropertyFilteredByDomainID(dynamicconfig.QueueProcessorSplitLookAheadDurationByDomainID, 20*time.Minute),
		QueueProcessorPollBackoffInterval:                  dc.GetDurationProperty(dynamicconfig.QueueProcessorPollBackoffInterval, 5*time.Second),
		QueueProcessorPollBackoffIntervalJitterCoefficient: dc.GetFloat64Property(dynamicconfig.QueueProcessorPollBackoffIntervalJitterCoefficient, 0.15),
		QueueProcessorEnablePersistQueueStates:             dc.GetBoolProperty(dynamicconfig.QueueProcessorEnablePersistQueueStates, true),
		QueueProcessorEnableLoadQueueStates:                dc.GetBoolProperty(dynamicconfig.QueueProcessorEnableLoadQueueStates, true),

		TimerTaskBatchSize:                                dc.GetIntProperty(dynamicconfig.TimerTaskBatchSize, 100),
		TimerTaskDeleteBatchSize:                          dc.GetIntProperty(dynamicconfig.TimerTaskDeleteBatchSize, 4000),
		TimerProcessorGetFailureRetryCount:                dc.GetIntProperty(dynamicconfig.TimerProcessorGetFailureRetryCount, 5),
		TimerProcessorCompleteTimerFailureRetryCount:      dc.GetIntProperty(dynamicconfig.TimerProcessorCompleteTimerFailureRetryCount, 10),
		TimerProcessorUpdateAckInterval:                   dc.GetDurationProperty(dynamicconfig.TimerProcessorUpdateAckInterval, 30*time.Second),
		TimerProcessorUpdateAckIntervalJitterCoefficient:  dc.GetFloat64Property(dynamicconfig.TimerProcessorUpdateAckIntervalJitterCoefficient, 0.15),
		TimerProcessorCompleteTimerInterval:               dc.GetDurationProperty(dynamicconfig.TimerProcessorCompleteTimerInterval, 60*time.Second),
		TimerProcessorFailoverMaxPollRPS:                  dc.GetIntProperty(dynamicconfig.TimerProcessorFailoverMaxPollRPS, 1),
		TimerProcessorMaxPollRPS:                          dc.GetIntProperty(dynamicconfig.TimerProcessorMaxPollRPS, 20),
		TimerProcessorMaxPollInterval:                     dc.GetDurationProperty(dynamicconfig.TimerProcessorMaxPollInterval, 5*time.Minute),
		TimerProcessorMaxPollIntervalJitterCoefficient:    dc.GetFloat64Property(dynamicconfig.TimerProcessorMaxPollIntervalJitterCoefficient, 0.15),
		TimerProcessorSplitQueueInterval:                  dc.GetDurationProperty(dynamicconfig.TimerProcessorSplitQueueInterval, 1*time.Minute),
		TimerProcessorSplitQueueIntervalJitterCoefficient: dc.GetFloat64Property(dynamicconfig.TimerProcessorSplitQueueIntervalJitterCoefficient, 0.15),
		TimerProcessorMaxRedispatchQueueSize:              dc.GetIntProperty(dynamicconfig.TimerProcessorMaxRedispatchQueueSize, 10000),
		TimerProcessorMaxTimeShift:                        dc.GetDurationProperty(dynamicconfig.TimerProcessorMaxTimeShift, 1*time.Second),
		TimerProcessorHistoryArchivalSizeLimit:            dc.GetIntProperty(dynamicconfig.TimerProcessorHistoryArchivalSizeLimit, 500*1024),
		TimerProcessorArchivalTimeLimit:                   dc.GetDurationProperty(dynamicconfig.TimerProcessorArchivalTimeLimit, 1*time.Second),

		TransferTaskBatchSize:                                dc.GetIntProperty(dynamicconfig.TransferTaskBatchSize, 100),
		TransferTaskDeleteBatchSize:                          dc.GetIntProperty(dynamicconfig.TransferTaskDeleteBatchSize, 4000),
		TransferProcessorFailoverMaxPollRPS:                  dc.GetIntProperty(dynamicconfig.TransferProcessorFailoverMaxPollRPS, 1),
		TransferProcessorMaxPollRPS:                          dc.GetIntProperty(dynamicconfig.TransferProcessorMaxPollRPS, 20),
		TransferProcessorCompleteTransferFailureRetryCount:   dc.GetIntProperty(dynamicconfig.TransferProcessorCompleteTransferFailureRetryCount, 10),
		TransferProcessorMaxPollInterval:                     dc.GetDurationProperty(dynamicconfig.TransferProcessorMaxPollInterval, 1*time.Minute),
		TransferProcessorMaxPollIntervalJitterCoefficient:    dc.GetFloat64Property(dynamicconfig.TransferProcessorMaxPollIntervalJitterCoefficient, 0.15),
		TransferProcessorSplitQueueInterval:                  dc.GetDurationProperty(dynamicconfig.TransferProcessorSplitQueueInterval, 1*time.Minute),
		TransferProcessorSplitQueueIntervalJitterCoefficient: dc.GetFloat64Property(dynamicconfig.TransferProcessorSplitQueueIntervalJitterCoefficient, 0.15),
		TransferProcessorUpdateAckInterval:                   dc.GetDurationProperty(dynamicconfig.TransferProcessorUpdateAckInterval, 30*time.Second),
		TransferProcessorUpdateAckIntervalJitterCoefficient:  dc.GetFloat64Property(dynamicconfig.TransferProcessorUpdateAckIntervalJitterCoefficient, 0.15),
		TransferProcessorCompleteTransferInterval:            dc.GetDurationProperty(dynamicconfig.TransferProcessorCompleteTransferInterval, 60*time.Second),
		TransferProcessorMaxRedispatchQueueSize:              dc.GetIntProperty(dynamicconfig.TransferProcessorMaxRedispatchQueueSize, 10000),
		TransferProcessorEnableValidator:                     dc.GetBoolProperty(dynamicconfig.TransferProcessorEnableValidator, false),
		TransferProcessorValidationInterval:                  dc.GetDurationProperty(dynamicconfig.TransferProcessorValidationInterval, 30*time.Second),
		TransferProcessorVisibilityArchivalTimeLimit:         dc.GetDurationProperty(dynamicconfig.TransferProcessorVisibilityArchivalTimeLimit, 200*time.Millisecond),

		CrossClusterTaskBatchSize:                                     dc.GetIntProperty(dynamicconfig.CrossClusterTaskBatchSize, 100),
		CrossClusterTaskDeleteBatchSize:                               dc.GetIntProperty(dynamicconfig.CrossClusterTaskDeleteBatchSize, 4000),
		CrossClusterTaskFetchBatchSize:                                dc.GetIntPropertyFilteredByShardID(dynamicconfig.CrossClusterTaskFetchBatchSize, 100),
		CrossClusterSourceProcessorMaxPollRPS:                         dc.GetIntProperty(dynamicconfig.CrossClusterSourceProcessorMaxPollRPS, 20),
		CrossClusterSourceProcessorMaxPollInterval:                    dc.GetDurationProperty(dynamicconfig.CrossClusterSourceProcessorMaxPollInterval, 1*time.Minute),
		CrossClusterSourceProcessorMaxPollIntervalJitterCoefficient:   dc.GetFloat64Property(dynamicconfig.CrossClusterSourceProcessorMaxPollIntervalJitterCoefficient, 0.15),
		CrossClusterSourceProcessorUpdateAckInterval:                  dc.GetDurationProperty(dynamicconfig.CrossClusterSourceProcessorUpdateAckInterval, 30*time.Second),
		CrossClusterSourceProcessorUpdateAckIntervalJitterCoefficient: dc.GetFloat64Property(dynamicconfig.CrossClusterSourceProcessorUpdateAckIntervalJitterCoefficient, 0.15),
		CrossClusterSourceProcessorMaxRedispatchQueueSize:             dc.GetIntProperty(dynamicconfig.CrossClusterSourceProcessorMaxRedispatchQueueSize, 10000),
		CrossClusterSourceProcessorMaxPendingTaskSize:                 dc.GetIntProperty(dynamicconfig.CrossClusterSourceProcessorMaxPendingTaskSize, 500),

		CrossClusterTargetProcessorMaxPendingTasks:            dc.GetIntProperty(dynamicconfig.CrossClusterTargetProcessorMaxPendingTasks, 200),
		CrossClusterTargetProcessorMaxRetryCount:              dc.GetIntProperty(dynamicconfig.CrossClusterTargetProcessorMaxRetryCount, 20),
		CrossClusterTargetProcessorTaskWaitInterval:           dc.GetDurationProperty(dynamicconfig.CrossClusterTargetProcessorTaskWaitInterval, 3*time.Second),
		CrossClusterTargetProcessorServiceBusyBackoffInterval: dc.GetDurationProperty(dynamicconfig.CrossClusterTargetProcessorServiceBusyBackoffInterval, 5*time.Second),
		CrossClusterTargetProcessorJitterCoefficient:          dc.GetFloat64Property(dynamicconfig.CrossClusterTargetProcessorJitterCoefficient, 0.15),

		CrossClusterFetcherParallelism:                dc.GetIntProperty(dynamicconfig.CrossClusterFetcherParallelism, 1),
		CrossClusterFetcherAggregationInterval:        dc.GetDurationProperty(dynamicconfig.CrossClusterFetcherAggregationInterval, 2*time.Second),
		CrossClusterFetcherServiceBusyBackoffInterval: dc.GetDurationProperty(dynamicconfig.CrossClusterFetcherServiceBusyBackoffInterval, 5*time.Second),
		CrossClusterFetcherErrorBackoffInterval:       dc.GetDurationProperty(dynamicconfig.CrossClusterFetcherServiceBusyBackoffInterval, time.Second),
		CrossClusterFetcherJitterCoefficient:          dc.GetFloat64Property(dynamicconfig.CrossClusterFetcherJitterCoefficient, 0.15),

		ReplicatorTaskBatchSize:                               dc.GetIntProperty(dynamicconfig.ReplicatorTaskBatchSize, 100),
		ReplicatorTaskDeleteBatchSize:                         dc.GetIntProperty(dynamicconfig.ReplicatorTaskDeleteBatchSize, 4000),
		ReplicatorTaskWorkerCount:                             dc.GetIntProperty(dynamicconfig.ReplicatorTaskWorkerCount, 10),
		ReplicatorReadTaskMaxRetryCount:                       dc.GetIntProperty(dynamicconfig.ReplicatorReadTaskMaxRetryCount, 3),
		ReplicatorProcessorMaxPollRPS:                         dc.GetIntProperty(dynamicconfig.ReplicatorProcessorMaxPollRPS, 20),
		ReplicatorProcessorMaxPollInterval:                    dc.GetDurationProperty(dynamicconfig.ReplicatorProcessorMaxPollInterval, 1*time.Minute),
		ReplicatorProcessorMaxPollIntervalJitterCoefficient:   dc.GetFloat64Property(dynamicconfig.ReplicatorProcessorMaxPollIntervalJitterCoefficient, 0.15),
		ReplicatorProcessorUpdateAckInterval:                  dc.GetDurationProperty(dynamicconfig.ReplicatorProcessorUpdateAckInterval, 5*time.Second),
		ReplicatorProcessorUpdateAckIntervalJitterCoefficient: dc.GetFloat64Property(dynamicconfig.ReplicatorProcessorUpdateAckIntervalJitterCoefficient, 0.15),
		ReplicatorProcessorMaxRedispatchQueueSize:             dc.GetIntProperty(dynamicconfig.ReplicatorProcessorMaxRedispatchQueueSize, 10000),
		ReplicatorProcessorEnablePriorityTaskProcessor:        dc.GetBoolProperty(dynamicconfig.ReplicatorProcessorEnablePriorityTaskProcessor, false),
		ReplicatorProcessorFetchTasksBatchSize:                dc.GetIntPropertyFilteredByShardID(dynamicconfig.ReplicatorTaskBatchSize, 25),
		ReplicatorUpperLatency:                                dc.GetDurationProperty(dynamicconfig.ReplicatorUpperLatency, 40*time.Second),

		ExecutionMgrNumConns:       dc.GetIntProperty(dynamicconfig.ExecutionMgrNumConns, 50),
		HistoryMgrNumConns:         dc.GetIntProperty(dynamicconfig.HistoryMgrNumConns, 50),
		MaximumBufferedEventsBatch: dc.GetIntProperty(dynamicconfig.MaximumBufferedEventsBatch, 100),
		// 10K signals should big enough given workflow execution has 200K history lengh limit. It needs to be non-zero to protect continueAsNew from infinit loop
		MaximumSignalsPerExecution:      dc.GetIntPropertyFilteredByDomain(dynamicconfig.MaximumSignalsPerExecution, 10000),
		ShardUpdateMinInterval:          dc.GetDurationProperty(dynamicconfig.ShardUpdateMinInterval, 5*time.Minute),
		ShardSyncMinInterval:            dc.GetDurationProperty(dynamicconfig.ShardSyncMinInterval, 5*time.Minute),
		ShardSyncTimerJitterCoefficient: dc.GetFloat64Property(dynamicconfig.TransferProcessorMaxPollIntervalJitterCoefficient, 0.15),

		// history client: client/history/client.go set the client timeout 30s
		LongPollExpirationInterval:          dc.GetDurationPropertyFilteredByDomain(dynamicconfig.HistoryLongPollExpirationInterval, time.Second*20),
		EventEncodingType:                   dc.GetStringPropertyFilteredByDomain(dynamicconfig.DefaultEventEncoding, string(common.EncodingTypeThriftRW)),
		EnableParentClosePolicy:             dc.GetBoolPropertyFilteredByDomain(dynamicconfig.EnableParentClosePolicy, true),
		NumParentClosePolicySystemWorkflows: dc.GetIntProperty(dynamicconfig.NumParentClosePolicySystemWorkflows, 10),
		EnableParentClosePolicyWorker:       dc.GetBoolProperty(dynamicconfig.EnableParentClosePolicyWorker, true),
		ParentClosePolicyThreshold:          dc.GetIntPropertyFilteredByDomain(dynamicconfig.ParentClosePolicyThreshold, 10),

		NumArchiveSystemWorkflows:       dc.GetIntProperty(dynamicconfig.NumArchiveSystemWorkflows, 1000),
		ArchiveRequestRPS:               dc.GetIntProperty(dynamicconfig.ArchiveRequestRPS, 300), // should be much smaller than frontend RPS
		AllowArchivingIncompleteHistory: dc.GetBoolProperty(dynamicconfig.AllowArchivingIncompleteHistory, false),

		BlobSizeLimitError:     dc.GetIntPropertyFilteredByDomain(dynamicconfig.BlobSizeLimitError, 2*1024*1024),
		BlobSizeLimitWarn:      dc.GetIntPropertyFilteredByDomain(dynamicconfig.BlobSizeLimitWarn, 512*1024),
		HistorySizeLimitError:  dc.GetIntPropertyFilteredByDomain(dynamicconfig.HistorySizeLimitError, 200*1024*1024),
		HistorySizeLimitWarn:   dc.GetIntPropertyFilteredByDomain(dynamicconfig.HistorySizeLimitWarn, 50*1024*1024),
		HistoryCountLimitError: dc.GetIntPropertyFilteredByDomain(dynamicconfig.HistoryCountLimitError, 200*1024),
		HistoryCountLimitWarn:  dc.GetIntPropertyFilteredByDomain(dynamicconfig.HistoryCountLimitWarn, 50*1024),

		ThrottledLogRPS:   dc.GetIntProperty(dynamicconfig.HistoryThrottledLogRPS, 4),
		EnableStickyQuery: dc.GetBoolPropertyFilteredByDomain(dynamicconfig.EnableStickyQuery, true),

		ValidSearchAttributes:                    dc.GetMapProperty(dynamicconfig.ValidSearchAttributes, definition.GetDefaultIndexedKeys()),
		SearchAttributesNumberOfKeysLimit:        dc.GetIntPropertyFilteredByDomain(dynamicconfig.SearchAttributesNumberOfKeysLimit, 100),
		SearchAttributesSizeOfValueLimit:         dc.GetIntPropertyFilteredByDomain(dynamicconfig.SearchAttributesSizeOfValueLimit, 2*1024),
		SearchAttributesTotalSizeLimit:           dc.GetIntPropertyFilteredByDomain(dynamicconfig.SearchAttributesTotalSizeLimit, 40*1024),
		StickyTTL:                                dc.GetDurationPropertyFilteredByDomain(dynamicconfig.StickyTTL, time.Hour*24*365),
		DecisionHeartbeatTimeout:                 dc.GetDurationPropertyFilteredByDomain(dynamicconfig.DecisionHeartbeatTimeout, time.Minute*30),
		DecisionRetryCriticalAttempts:            dc.GetIntProperty(dynamicconfig.DecisionRetryCriticalAttempts, 10), // about 30m
		DecisionRetryMaxAttempts:                 dc.GetIntPropertyFilteredByDomain(dynamicconfig.DecisionRetryMaxAttempts, 1000),
		NormalDecisionScheduleToStartMaxAttempts: dc.GetIntPropertyFilteredByDomain(dynamicconfig.NormalDecisionScheduleToStartMaxAttempts, 0),
		NormalDecisionScheduleToStartTimeout:     dc.GetDurationPropertyFilteredByDomain(dynamicconfig.NormalDecisionScheduleToStartTimeout, time.Minute*5),

		ReplicationTaskFetcherParallelism:                  dc.GetIntProperty(dynamicconfig.ReplicationTaskFetcherParallelism, 1),
		ReplicationTaskFetcherAggregationInterval:          dc.GetDurationProperty(dynamicconfig.ReplicationTaskFetcherAggregationInterval, 2*time.Second),
		ReplicationTaskFetcherTimerJitterCoefficient:       dc.GetFloat64Property(dynamicconfig.ReplicationTaskFetcherTimerJitterCoefficient, 0.15),
		ReplicationTaskFetcherErrorRetryWait:               dc.GetDurationProperty(dynamicconfig.ReplicationTaskFetcherErrorRetryWait, time.Second),
		ReplicationTaskFetcherServiceBusyWait:              dc.GetDurationProperty(dynamicconfig.ReplicationTaskFetcherServiceBusyWait, 60*time.Second),
		ReplicationTaskProcessorErrorRetryWait:             dc.GetDurationPropertyFilteredByShardID(dynamicconfig.ReplicationTaskProcessorErrorRetryWait, 50*time.Millisecond),
		ReplicationTaskProcessorErrorRetryMaxAttempts:      dc.GetIntPropertyFilteredByShardID(dynamicconfig.ReplicationTaskProcessorErrorRetryMaxAttempts, 10),
		ReplicationTaskProcessorErrorSecondRetryWait:       dc.GetDurationPropertyFilteredByShardID(dynamicconfig.ReplicationTaskProcessorErrorSecondRetryWait, 5*time.Second),
		ReplicationTaskProcessorErrorSecondRetryMaxWait:    dc.GetDurationPropertyFilteredByShardID(dynamicconfig.ReplicationTaskProcessorErrorSecondRetryMaxWait, 30*time.Second),
		ReplicationTaskProcessorErrorSecondRetryExpiration: dc.GetDurationPropertyFilteredByShardID(dynamicconfig.ReplicationTaskProcessorErrorSecondRetryExpiration, 5*time.Minute),
		ReplicationTaskProcessorNoTaskRetryWait:            dc.GetDurationPropertyFilteredByShardID(dynamicconfig.ReplicationTaskProcessorNoTaskInitialWait, 2*time.Second),
		ReplicationTaskProcessorCleanupInterval:            dc.GetDurationPropertyFilteredByShardID(dynamicconfig.ReplicationTaskProcessorCleanupInterval, 1*time.Minute),
		ReplicationTaskProcessorCleanupJitterCoefficient:   dc.GetFloat64PropertyFilteredByShardID(dynamicconfig.ReplicationTaskProcessorCleanupJitterCoefficient, 0.15),
		ReplicationTaskProcessorReadHistoryBatchSize:       dc.GetIntProperty(dynamicconfig.ReplicationTaskProcessorReadHistoryBatchSize, 5),
		ReplicationTaskProcessorStartWait:                  dc.GetDurationPropertyFilteredByShardID(dynamicconfig.ReplicationTaskProcessorStartWait, 5*time.Second),
		ReplicationTaskProcessorStartWaitJitterCoefficient: dc.GetFloat64PropertyFilteredByShardID(dynamicconfig.ReplicationTaskProcessorStartWaitJitterCoefficient, 0.9),
		ReplicationTaskProcessorHostQPS:                    dc.GetFloat64Property(dynamicconfig.ReplicationTaskProcessorHostQPS, 1500),
		ReplicationTaskProcessorShardQPS:                   dc.GetFloat64Property(dynamicconfig.ReplicationTaskProcessorShardQPS, 5),
		ReplicationTaskGenerationQPS:                       dc.GetFloat64Property(dynamicconfig.ReplicationTaskGenerationQPS, 100),
		EnableReplicationTaskGeneration:                    dc.GetBoolPropertyFilteredByDomainIDAndWorkflowID(dynamicconfig.EnableReplicationTaskGeneration, true),

		EnableConsistentQuery:                 dc.GetBoolProperty(dynamicconfig.EnableConsistentQuery, true),
		EnableConsistentQueryByDomain:         dc.GetBoolPropertyFilteredByDomain(dynamicconfig.EnableConsistentQueryByDomain, false),
		EnableCrossClusterOperations:          dc.GetBoolPropertyFilteredByDomain(dynamicconfig.EnableCrossClusterOperations, false),
		MaxBufferedQueryCount:                 dc.GetIntProperty(dynamicconfig.MaxBufferedQueryCount, 1),
		MutableStateChecksumGenProbability:    dc.GetIntPropertyFilteredByDomain(dynamicconfig.MutableStateChecksumGenProbability, 0),
		MutableStateChecksumVerifyProbability: dc.GetIntPropertyFilteredByDomain(dynamicconfig.MutableStateChecksumVerifyProbability, 0),
		MutableStateChecksumInvalidateBefore:  dc.GetFloat64Property(dynamicconfig.MutableStateChecksumInvalidateBefore, 0),

		NotifyFailoverMarkerInterval:               dc.GetDurationProperty(dynamicconfig.NotifyFailoverMarkerInterval, 5*time.Second),
		NotifyFailoverMarkerTimerJitterCoefficient: dc.GetFloat64Property(dynamicconfig.NotifyFailoverMarkerTimerJitterCoefficient, 0.15),
		EnableGracefulFailover:                     dc.GetBoolProperty(dynamicconfig.EnableGracefulFailover, true),

		EnableActivityLocalDispatchByDomain: dc.GetBoolPropertyFilteredByDomain(dynamicconfig.EnableActivityLocalDispatchByDomain, false),

		ActivityMaxScheduleToStartTimeoutForRetry: dc.GetDurationPropertyFilteredByDomain(dynamicconfig.ActivityMaxScheduleToStartTimeoutForRetry, 30*time.Minute),

		EnableDebugMode:             dc.GetBoolProperty(dynamicconfig.EnableDebugMode, false)(),
		EnableTaskInfoLogByDomainID: dc.GetBoolPropertyFilteredByDomainID(dynamicconfig.HistoryEnableTaskInfoLogByDomainID, false),
	}

	return cfg
}

// NewForTest create new history service config for test
func NewForTest() *Config {
	return NewForTestByShardNumber(1)
}

// NewForTestByShardNumber create new history service config for test
func NewForTestByShardNumber(shardNumber int) *Config {
	dc := dynamicconfig.NewNopCollection()
	config := New(dc, shardNumber, config.StoreTypeCassandra, false)
	// reduce the duration of long poll to increase test speed
	config.LongPollExpirationInterval = dc.GetDurationPropertyFilteredByDomain(dynamicconfig.HistoryLongPollExpirationInterval, 10*time.Second)
	config.EnableConsistentQueryByDomain = dc.GetBoolPropertyFilteredByDomain(dynamicconfig.EnableConsistentQueryByDomain, true)
	config.ReplicationTaskProcessorHostQPS = dc.GetFloat64Property(dynamicconfig.ReplicationTaskProcessorHostQPS, 10000)
	config.ReplicationTaskProcessorShardQPS = dc.GetFloat64Property(dynamicconfig.ReplicationTaskProcessorShardQPS, 10000)
	config.ReplicationTaskProcessorStartWait = dc.GetDurationPropertyFilteredByShardID(dynamicconfig.ReplicationTaskProcessorStartWait, time.Nanosecond)
	config.EnableActivityLocalDispatchByDomain = dc.GetBoolPropertyFilteredByDomain(dynamicconfig.EnableActivityLocalDispatchByDomain, true)
	config.EnableCrossClusterOperations = dc.GetBoolPropertyFilteredByDomain(dynamicconfig.EnableCrossClusterOperations, true)
	config.NormalDecisionScheduleToStartMaxAttempts = dc.GetIntPropertyFilteredByDomain(dynamicconfig.NormalDecisionScheduleToStartMaxAttempts, 3)
	return config
}

// GetShardID return the corresponding shard ID for a given workflow ID
func (config *Config) GetShardID(workflowID string) int {
	return common.WorkflowIDToHistoryShard(workflowID, config.NumberOfShards)
}
