package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// GetServiceGraphRequest represents the GetServiceGraphRequest schema from the OpenAPI specification
type GetServiceGraphRequest struct {
	Endtime interface{} `json:"EndTime"`
	Grouparn interface{} `json:"GroupARN,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Starttime interface{} `json:"StartTime"`
}

// GetTimeSeriesServiceStatisticsRequest represents the GetTimeSeriesServiceStatisticsRequest schema from the OpenAPI specification
type GetTimeSeriesServiceStatisticsRequest struct {
	Entityselectorexpression interface{} `json:"EntitySelectorExpression,omitempty"`
	Forecaststatistics interface{} `json:"ForecastStatistics,omitempty"`
	Grouparn interface{} `json:"GroupARN,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Period interface{} `json:"Period,omitempty"`
	Starttime interface{} `json:"StartTime"`
	Endtime interface{} `json:"EndTime"`
}

// GetSamplingStatisticSummariesRequest represents the GetSamplingStatisticSummariesRequest schema from the OpenAPI specification
type GetSamplingStatisticSummariesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// InsightsConfiguration represents the InsightsConfiguration schema from the OpenAPI specification
type InsightsConfiguration struct {
	Insightsenabled interface{} `json:"InsightsEnabled,omitempty"`
	Notificationsenabled interface{} `json:"NotificationsEnabled,omitempty"`
}

// Service represents the Service schema from the OpenAPI specification
type Service struct {
	Referenceid interface{} `json:"ReferenceId,omitempty"`
	Responsetimehistogram interface{} `json:"ResponseTimeHistogram,omitempty"`
	Root interface{} `json:"Root,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Accountid interface{} `json:"AccountId,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	State interface{} `json:"State,omitempty"`
	Durationhistogram interface{} `json:"DurationHistogram,omitempty"`
	Names interface{} `json:"Names,omitempty"`
	Summarystatistics interface{} `json:"SummaryStatistics,omitempty"`
	Edges interface{} `json:"Edges,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// Alias represents the Alias schema from the OpenAPI specification
type Alias struct {
	Name interface{} `json:"Name,omitempty"`
	Names interface{} `json:"Names,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// GetSamplingStatisticSummariesResult represents the GetSamplingStatisticSummariesResult schema from the OpenAPI specification
type GetSamplingStatisticSummariesResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Samplingstatisticsummaries interface{} `json:"SamplingStatisticSummaries,omitempty"`
}

// DeleteGroupResult represents the DeleteGroupResult schema from the OpenAPI specification
type DeleteGroupResult struct {
}

// GetInsightRequest represents the GetInsightRequest schema from the OpenAPI specification
type GetInsightRequest struct {
	Insightid interface{} `json:"InsightId"`
}

// PutTraceSegmentsRequest represents the PutTraceSegmentsRequest schema from the OpenAPI specification
type PutTraceSegmentsRequest struct {
	Tracesegmentdocuments interface{} `json:"TraceSegmentDocuments"`
}

// BackendConnectionErrors represents the BackendConnectionErrors schema from the OpenAPI specification
type BackendConnectionErrors struct {
	Unknownhostcount interface{} `json:"UnknownHostCount,omitempty"`
	Connectionrefusedcount interface{} `json:"ConnectionRefusedCount,omitempty"`
	Httpcode4xxcount interface{} `json:"HTTPCode4XXCount,omitempty"`
	Httpcode5xxcount interface{} `json:"HTTPCode5XXCount,omitempty"`
	Othercount interface{} `json:"OtherCount,omitempty"`
	Timeoutcount interface{} `json:"TimeoutCount,omitempty"`
}

// GetInsightSummariesRequest represents the GetInsightSummariesRequest schema from the OpenAPI specification
type GetInsightSummariesRequest struct {
	States interface{} `json:"States,omitempty"`
	Endtime interface{} `json:"EndTime"`
	Grouparn interface{} `json:"GroupARN,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Starttime interface{} `json:"StartTime"`
}

// ErrorRootCause represents the ErrorRootCause schema from the OpenAPI specification
type ErrorRootCause struct {
	Clientimpacting interface{} `json:"ClientImpacting,omitempty"`
	Services interface{} `json:"Services,omitempty"`
}

// FaultRootCauseService represents the FaultRootCauseService schema from the OpenAPI specification
type FaultRootCauseService struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Entitypath interface{} `json:"EntityPath,omitempty"`
	Inferred interface{} `json:"Inferred,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Names interface{} `json:"Names,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// ValueWithServiceIds represents the ValueWithServiceIds schema from the OpenAPI specification
type ValueWithServiceIds struct {
	Annotationvalue interface{} `json:"AnnotationValue,omitempty"`
	Serviceids interface{} `json:"ServiceIds,omitempty"`
}

// GetEncryptionConfigResult represents the GetEncryptionConfigResult schema from the OpenAPI specification
type GetEncryptionConfigResult struct {
	Encryptionconfig interface{} `json:"EncryptionConfig,omitempty"`
}

// SamplingStatisticSummary represents the SamplingStatisticSummary schema from the OpenAPI specification
type SamplingStatisticSummary struct {
	Borrowcount interface{} `json:"BorrowCount,omitempty"`
	Requestcount interface{} `json:"RequestCount,omitempty"`
	Rulename interface{} `json:"RuleName,omitempty"`
	Sampledcount interface{} `json:"SampledCount,omitempty"`
	Timestamp interface{} `json:"Timestamp,omitempty"`
}

// ResponseTimeRootCauseService represents the ResponseTimeRootCauseService schema from the OpenAPI specification
type ResponseTimeRootCauseService struct {
	Name interface{} `json:"Name,omitempty"`
	Names interface{} `json:"Names,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Accountid interface{} `json:"AccountId,omitempty"`
	Entitypath interface{} `json:"EntityPath,omitempty"`
	Inferred interface{} `json:"Inferred,omitempty"`
}

// Tag represents the Tag schema from the OpenAPI specification
type Tag struct {
	Key interface{} `json:"Key"`
	Value interface{} `json:"Value"`
}

// TagResourceRequest represents the TagResourceRequest schema from the OpenAPI specification
type TagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceARN"`
	Tags interface{} `json:"Tags"`
}

// DeleteSamplingRuleRequest represents the DeleteSamplingRuleRequest schema from the OpenAPI specification
type DeleteSamplingRuleRequest struct {
	Rulename interface{} `json:"RuleName,omitempty"`
	Rulearn interface{} `json:"RuleARN,omitempty"`
}

// ErrorRootCauseService represents the ErrorRootCauseService schema from the OpenAPI specification
type ErrorRootCauseService struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Entitypath interface{} `json:"EntityPath,omitempty"`
	Inferred interface{} `json:"Inferred,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Names interface{} `json:"Names,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// Insight represents the Insight schema from the OpenAPI specification
type Insight struct {
	Groupname interface{} `json:"GroupName,omitempty"`
	Insightid interface{} `json:"InsightId,omitempty"`
	Rootcauseserviceid ServiceId `json:"RootCauseServiceId,omitempty"` // <p/>
	Categories interface{} `json:"Categories,omitempty"`
	Rootcauseservicerequestimpactstatistics interface{} `json:"RootCauseServiceRequestImpactStatistics,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Grouparn interface{} `json:"GroupARN,omitempty"`
	Summary interface{} `json:"Summary,omitempty"`
	Topanomalousservices interface{} `json:"TopAnomalousServices,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	State interface{} `json:"State,omitempty"`
	Clientrequestimpactstatistics interface{} `json:"ClientRequestImpactStatistics,omitempty"`
}

// ForecastStatistics represents the ForecastStatistics schema from the OpenAPI specification
type ForecastStatistics struct {
	Faultcounthigh interface{} `json:"FaultCountHigh,omitempty"`
	Faultcountlow interface{} `json:"FaultCountLow,omitempty"`
}

// FaultRootCause represents the FaultRootCause schema from the OpenAPI specification
type FaultRootCause struct {
	Clientimpacting interface{} `json:"ClientImpacting,omitempty"`
	Services interface{} `json:"Services,omitempty"`
}

// GroupSummary represents the GroupSummary schema from the OpenAPI specification
type GroupSummary struct {
	Filterexpression interface{} `json:"FilterExpression,omitempty"`
	Grouparn interface{} `json:"GroupARN,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Insightsconfiguration interface{} `json:"InsightsConfiguration,omitempty"`
}

// DeleteResourcePolicyResult represents the DeleteResourcePolicyResult schema from the OpenAPI specification
type DeleteResourcePolicyResult struct {
}

// PutTelemetryRecordsRequest represents the PutTelemetryRecordsRequest schema from the OpenAPI specification
type PutTelemetryRecordsRequest struct {
	Ec2instanceid interface{} `json:"EC2InstanceId,omitempty"`
	Hostname interface{} `json:"Hostname,omitempty"`
	Resourcearn interface{} `json:"ResourceARN,omitempty"`
	Telemetryrecords interface{} `json:"TelemetryRecords"`
}

// ErrorRootCauseEntity represents the ErrorRootCauseEntity schema from the OpenAPI specification
type ErrorRootCauseEntity struct {
	Name interface{} `json:"Name,omitempty"`
	Remote interface{} `json:"Remote,omitempty"`
	Exceptions interface{} `json:"Exceptions,omitempty"`
}

// FaultRootCauseEntity represents the FaultRootCauseEntity schema from the OpenAPI specification
type FaultRootCauseEntity struct {
	Exceptions interface{} `json:"Exceptions,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Remote interface{} `json:"Remote,omitempty"`
}

// GetGroupRequest represents the GetGroupRequest schema from the OpenAPI specification
type GetGroupRequest struct {
	Grouparn interface{} `json:"GroupARN,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
}

// ResourceARNDetail represents the ResourceARNDetail schema from the OpenAPI specification
type ResourceARNDetail struct {
	Arn interface{} `json:"ARN,omitempty"`
}

// GetSamplingRulesRequest represents the GetSamplingRulesRequest schema from the OpenAPI specification
type GetSamplingRulesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// PutTraceSegmentsResult represents the PutTraceSegmentsResult schema from the OpenAPI specification
type PutTraceSegmentsResult struct {
	Unprocessedtracesegments interface{} `json:"UnprocessedTraceSegments,omitempty"`
}

// GetTraceSummariesRequest represents the GetTraceSummariesRequest schema from the OpenAPI specification
type GetTraceSummariesRequest struct {
	Samplingstrategy interface{} `json:"SamplingStrategy,omitempty"`
	Starttime interface{} `json:"StartTime"`
	Timerangetype interface{} `json:"TimeRangeType,omitempty"`
	Endtime interface{} `json:"EndTime"`
	Filterexpression interface{} `json:"FilterExpression,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Sampling interface{} `json:"Sampling,omitempty"`
}

// ResponseTimeRootCauseEntity represents the ResponseTimeRootCauseEntity schema from the OpenAPI specification
type ResponseTimeRootCauseEntity struct {
	Name interface{} `json:"Name,omitempty"`
	Remote interface{} `json:"Remote,omitempty"`
	Coverage interface{} `json:"Coverage,omitempty"`
}

// GetTraceSummariesResult represents the GetTraceSummariesResult schema from the OpenAPI specification
type GetTraceSummariesResult struct {
	Approximatetime interface{} `json:"ApproximateTime,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Tracesummaries interface{} `json:"TraceSummaries,omitempty"`
	Tracesprocessedcount interface{} `json:"TracesProcessedCount,omitempty"`
}

// SamplingTargetDocument represents the SamplingTargetDocument schema from the OpenAPI specification
type SamplingTargetDocument struct {
	Interval interface{} `json:"Interval,omitempty"`
	Reservoirquota interface{} `json:"ReservoirQuota,omitempty"`
	Reservoirquotattl interface{} `json:"ReservoirQuotaTTL,omitempty"`
	Rulename interface{} `json:"RuleName,omitempty"`
	Fixedrate interface{} `json:"FixedRate,omitempty"`
}

// AnnotationValue represents the AnnotationValue schema from the OpenAPI specification
type AnnotationValue struct {
	Booleanvalue interface{} `json:"BooleanValue,omitempty"`
	Numbervalue interface{} `json:"NumberValue,omitempty"`
	Stringvalue interface{} `json:"StringValue,omitempty"`
}

// ListResourcePoliciesRequest represents the ListResourcePoliciesRequest schema from the OpenAPI specification
type ListResourcePoliciesRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetSamplingTargetsRequest represents the GetSamplingTargetsRequest schema from the OpenAPI specification
type GetSamplingTargetsRequest struct {
	Samplingstatisticsdocuments interface{} `json:"SamplingStatisticsDocuments"`
}

// GetInsightEventsRequest represents the GetInsightEventsRequest schema from the OpenAPI specification
type GetInsightEventsRequest struct {
	Insightid interface{} `json:"InsightId"`
	Maxresults interface{} `json:"MaxResults,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// SamplingRule represents the SamplingRule schema from the OpenAPI specification
type SamplingRule struct {
	Resourcearn interface{} `json:"ResourceARN"`
	Rulearn interface{} `json:"RuleARN,omitempty"`
	Priority interface{} `json:"Priority"`
	Urlpath interface{} `json:"URLPath"`
	Version interface{} `json:"Version"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Fixedrate interface{} `json:"FixedRate"`
	Httpmethod interface{} `json:"HTTPMethod"`
	Host interface{} `json:"Host"`
	Servicename interface{} `json:"ServiceName"`
	Servicetype interface{} `json:"ServiceType"`
	Reservoirsize interface{} `json:"ReservoirSize"`
	Rulename interface{} `json:"RuleName,omitempty"`
}

// PutEncryptionConfigResult represents the PutEncryptionConfigResult schema from the OpenAPI specification
type PutEncryptionConfigResult struct {
	Encryptionconfig interface{} `json:"EncryptionConfig,omitempty"`
}

// InstanceIdDetail represents the InstanceIdDetail schema from the OpenAPI specification
type InstanceIdDetail struct {
	Id interface{} `json:"Id,omitempty"`
}

// ServiceId represents the ServiceId schema from the OpenAPI specification
type ServiceId struct {
	Accountid interface{} `json:"AccountId,omitempty"`
	Name interface{} `json:"Name,omitempty"`
	Names interface{} `json:"Names,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// CreateSamplingRuleResult represents the CreateSamplingRuleResult schema from the OpenAPI specification
type CreateSamplingRuleResult struct {
	Samplingrulerecord interface{} `json:"SamplingRuleRecord,omitempty"`
}

// UpdateSamplingRuleResult represents the UpdateSamplingRuleResult schema from the OpenAPI specification
type UpdateSamplingRuleResult struct {
	Samplingrulerecord interface{} `json:"SamplingRuleRecord,omitempty"`
}

// Trace represents the Trace schema from the OpenAPI specification
type Trace struct {
	Segments interface{} `json:"Segments,omitempty"`
	Duration interface{} `json:"Duration,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Limitexceeded interface{} `json:"LimitExceeded,omitempty"`
}

// UntagResourceRequest represents the UntagResourceRequest schema from the OpenAPI specification
type UntagResourceRequest struct {
	Resourcearn interface{} `json:"ResourceARN"`
	Tagkeys interface{} `json:"TagKeys"`
}

// ListResourcePoliciesResult represents the ListResourcePoliciesResult schema from the OpenAPI specification
type ListResourcePoliciesResult struct {
	Resourcepolicies interface{} `json:"ResourcePolicies,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// SamplingRuleUpdate represents the SamplingRuleUpdate schema from the OpenAPI specification
type SamplingRuleUpdate struct {
	Urlpath interface{} `json:"URLPath,omitempty"`
	Fixedrate interface{} `json:"FixedRate,omitempty"`
	Priority interface{} `json:"Priority,omitempty"`
	Rulename interface{} `json:"RuleName,omitempty"`
	Servicename interface{} `json:"ServiceName,omitempty"`
	Servicetype interface{} `json:"ServiceType,omitempty"`
	Httpmethod interface{} `json:"HTTPMethod,omitempty"`
	Resourcearn interface{} `json:"ResourceARN,omitempty"`
	Attributes interface{} `json:"Attributes,omitempty"`
	Host interface{} `json:"Host,omitempty"`
	Reservoirsize interface{} `json:"ReservoirSize,omitempty"`
	Rulearn interface{} `json:"RuleARN,omitempty"`
}

// TraceUser represents the TraceUser schema from the OpenAPI specification
type TraceUser struct {
	Serviceids interface{} `json:"ServiceIds,omitempty"`
	Username interface{} `json:"UserName,omitempty"`
}

// PutTelemetryRecordsResult represents the PutTelemetryRecordsResult schema from the OpenAPI specification
type PutTelemetryRecordsResult struct {
}

// PutResourcePolicyResult represents the PutResourcePolicyResult schema from the OpenAPI specification
type PutResourcePolicyResult struct {
	Resourcepolicy interface{} `json:"ResourcePolicy,omitempty"`
}

// ListTagsForResourceRequest represents the ListTagsForResourceRequest schema from the OpenAPI specification
type ListTagsForResourceRequest struct {
	Resourcearn interface{} `json:"ResourceARN"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// GetGroupResult represents the GetGroupResult schema from the OpenAPI specification
type GetGroupResult struct {
	Group interface{} `json:"Group,omitempty"`
}

// GetTimeSeriesServiceStatisticsResult represents the GetTimeSeriesServiceStatisticsResult schema from the OpenAPI specification
type GetTimeSeriesServiceStatisticsResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Timeseriesservicestatistics interface{} `json:"TimeSeriesServiceStatistics,omitempty"`
	Containsoldgroupversions interface{} `json:"ContainsOldGroupVersions,omitempty"`
}

// GetEncryptionConfigRequest represents the GetEncryptionConfigRequest schema from the OpenAPI specification
type GetEncryptionConfigRequest struct {
}

// SamplingRuleRecord represents the SamplingRuleRecord schema from the OpenAPI specification
type SamplingRuleRecord struct {
	Modifiedat interface{} `json:"ModifiedAt,omitempty"`
	Samplingrule interface{} `json:"SamplingRule,omitempty"`
	Createdat interface{} `json:"CreatedAt,omitempty"`
}

// HistogramEntry represents the HistogramEntry schema from the OpenAPI specification
type HistogramEntry struct {
	Count interface{} `json:"Count,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// DeleteResourcePolicyRequest represents the DeleteResourcePolicyRequest schema from the OpenAPI specification
type DeleteResourcePolicyRequest struct {
	Policyrevisionid interface{} `json:"PolicyRevisionId,omitempty"`
	Policyname interface{} `json:"PolicyName"`
}

// TimeSeriesServiceStatistics represents the TimeSeriesServiceStatistics schema from the OpenAPI specification
type TimeSeriesServiceStatistics struct {
	Servicesummarystatistics ServiceStatistics `json:"ServiceSummaryStatistics,omitempty"` // Response statistics for a service.
	Timestamp interface{} `json:"Timestamp,omitempty"`
	Edgesummarystatistics EdgeStatistics `json:"EdgeSummaryStatistics,omitempty"` // Response statistics for an edge.
	Responsetimehistogram interface{} `json:"ResponseTimeHistogram,omitempty"`
	Serviceforecaststatistics interface{} `json:"ServiceForecastStatistics,omitempty"`
}

// Annotations represents the Annotations schema from the OpenAPI specification
type Annotations struct {
}

// CreateGroupResult represents the CreateGroupResult schema from the OpenAPI specification
type CreateGroupResult struct {
	Group interface{} `json:"Group,omitempty"`
}

// TagResourceResponse represents the TagResourceResponse schema from the OpenAPI specification
type TagResourceResponse struct {
}

// BatchGetTracesRequest represents the BatchGetTracesRequest schema from the OpenAPI specification
type BatchGetTracesRequest struct {
	Traceids interface{} `json:"TraceIds"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// BatchGetTracesResult represents the BatchGetTracesResult schema from the OpenAPI specification
type BatchGetTracesResult struct {
	Unprocessedtraceids interface{} `json:"UnprocessedTraceIds,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Traces interface{} `json:"Traces,omitempty"`
}

// GetGroupsRequest represents the GetGroupsRequest schema from the OpenAPI specification
type GetGroupsRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// InsightSummary represents the InsightSummary schema from the OpenAPI specification
type InsightSummary struct {
	Rootcauseservicerequestimpactstatistics interface{} `json:"RootCauseServiceRequestImpactStatistics,omitempty"`
	Summary interface{} `json:"Summary,omitempty"`
	Categories interface{} `json:"Categories,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Lastupdatetime interface{} `json:"LastUpdateTime,omitempty"`
	Topanomalousservices interface{} `json:"TopAnomalousServices,omitempty"`
	Rootcauseserviceid ServiceId `json:"RootCauseServiceId,omitempty"` // <p/>
	Grouparn interface{} `json:"GroupARN,omitempty"`
	Insightid interface{} `json:"InsightId,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	State interface{} `json:"State,omitempty"`
	Clientrequestimpactstatistics interface{} `json:"ClientRequestImpactStatistics,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
}

// RootCauseException represents the RootCauseException schema from the OpenAPI specification
type RootCauseException struct {
	Message interface{} `json:"Message,omitempty"`
	Name interface{} `json:"Name,omitempty"`
}

// PutResourcePolicyRequest represents the PutResourcePolicyRequest schema from the OpenAPI specification
type PutResourcePolicyRequest struct {
	Bypasspolicylockoutcheck interface{} `json:"BypassPolicyLockoutCheck,omitempty"`
	Policydocument interface{} `json:"PolicyDocument"`
	Policyname interface{} `json:"PolicyName"`
	Policyrevisionid interface{} `json:"PolicyRevisionId,omitempty"`
}

// GetSamplingRulesResult represents the GetSamplingRulesResult schema from the OpenAPI specification
type GetSamplingRulesResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Samplingrulerecords interface{} `json:"SamplingRuleRecords,omitempty"`
}

// ResponseTimeRootCause represents the ResponseTimeRootCause schema from the OpenAPI specification
type ResponseTimeRootCause struct {
	Clientimpacting interface{} `json:"ClientImpacting,omitempty"`
	Services interface{} `json:"Services,omitempty"`
}

// EncryptionConfig represents the EncryptionConfig schema from the OpenAPI specification
type EncryptionConfig struct {
	Keyid interface{} `json:"KeyId,omitempty"`
	Status interface{} `json:"Status,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
}

// ServiceStatistics represents the ServiceStatistics schema from the OpenAPI specification
type ServiceStatistics struct {
	Faultstatistics interface{} `json:"FaultStatistics,omitempty"`
	Okcount interface{} `json:"OkCount,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
	Totalresponsetime interface{} `json:"TotalResponseTime,omitempty"`
	Errorstatistics interface{} `json:"ErrorStatistics,omitempty"`
}

// ErrorStatistics represents the ErrorStatistics schema from the OpenAPI specification
type ErrorStatistics struct {
	Othercount interface{} `json:"OtherCount,omitempty"`
	Throttlecount interface{} `json:"ThrottleCount,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
}

// AnomalousService represents the AnomalousService schema from the OpenAPI specification
type AnomalousService struct {
	Serviceid ServiceId `json:"ServiceId,omitempty"` // <p/>
}

// ResourcePolicy represents the ResourcePolicy schema from the OpenAPI specification
type ResourcePolicy struct {
	Policyrevisionid interface{} `json:"PolicyRevisionId,omitempty"`
	Lastupdatedtime interface{} `json:"LastUpdatedTime,omitempty"`
	Policydocument interface{} `json:"PolicyDocument,omitempty"`
	Policyname interface{} `json:"PolicyName,omitempty"`
}

// GetSamplingTargetsResult represents the GetSamplingTargetsResult schema from the OpenAPI specification
type GetSamplingTargetsResult struct {
	Unprocessedstatistics interface{} `json:"UnprocessedStatistics,omitempty"`
	Lastrulemodification interface{} `json:"LastRuleModification,omitempty"`
	Samplingtargetdocuments interface{} `json:"SamplingTargetDocuments,omitempty"`
}

// GetInsightEventsResult represents the GetInsightEventsResult schema from the OpenAPI specification
type GetInsightEventsResult struct {
	Insightevents interface{} `json:"InsightEvents,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// RequestImpactStatistics represents the RequestImpactStatistics schema from the OpenAPI specification
type RequestImpactStatistics struct {
	Faultcount interface{} `json:"FaultCount,omitempty"`
	Okcount interface{} `json:"OkCount,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
}

// Http represents the Http schema from the OpenAPI specification
type Http struct {
	Httpurl interface{} `json:"HttpURL,omitempty"`
	Useragent interface{} `json:"UserAgent,omitempty"`
	Clientip interface{} `json:"ClientIp,omitempty"`
	Httpmethod interface{} `json:"HttpMethod,omitempty"`
	Httpstatus interface{} `json:"HttpStatus,omitempty"`
}

// GetTraceGraphRequest represents the GetTraceGraphRequest schema from the OpenAPI specification
type GetTraceGraphRequest struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Traceids interface{} `json:"TraceIds"`
}

// GetServiceGraphResult represents the GetServiceGraphResult schema from the OpenAPI specification
type GetServiceGraphResult struct {
	Starttime interface{} `json:"StartTime,omitempty"`
	Containsoldgroupversions interface{} `json:"ContainsOldGroupVersions,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Services interface{} `json:"Services,omitempty"`
}

// CreateSamplingRuleRequest represents the CreateSamplingRuleRequest schema from the OpenAPI specification
type CreateSamplingRuleRequest struct {
	Samplingrule interface{} `json:"SamplingRule"`
	Tags interface{} `json:"Tags,omitempty"`
}

// GetGroupsResult represents the GetGroupsResult schema from the OpenAPI specification
type GetGroupsResult struct {
	Groups interface{} `json:"Groups,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// SamplingStrategy represents the SamplingStrategy schema from the OpenAPI specification
type SamplingStrategy struct {
	Name interface{} `json:"Name,omitempty"`
	Value interface{} `json:"Value,omitempty"`
}

// Group represents the Group schema from the OpenAPI specification
type Group struct {
	Grouparn interface{} `json:"GroupARN,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Insightsconfiguration interface{} `json:"InsightsConfiguration,omitempty"`
	Filterexpression interface{} `json:"FilterExpression,omitempty"`
}

// SamplingStatisticsDocument represents the SamplingStatisticsDocument schema from the OpenAPI specification
type SamplingStatisticsDocument struct {
	Clientid interface{} `json:"ClientID"`
	Requestcount interface{} `json:"RequestCount"`
	Rulename interface{} `json:"RuleName"`
	Sampledcount interface{} `json:"SampledCount"`
	Timestamp interface{} `json:"Timestamp"`
	Borrowcount interface{} `json:"BorrowCount,omitempty"`
}

// GetTraceGraphResult represents the GetTraceGraphResult schema from the OpenAPI specification
type GetTraceGraphResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Services interface{} `json:"Services,omitempty"`
}

// InsightImpactGraphService represents the InsightImpactGraphService schema from the OpenAPI specification
type InsightImpactGraphService struct {
	Name interface{} `json:"Name,omitempty"`
	Names interface{} `json:"Names,omitempty"`
	Referenceid interface{} `json:"ReferenceId,omitempty"`
	TypeField interface{} `json:"Type,omitempty"`
	Accountid interface{} `json:"AccountId,omitempty"`
	Edges interface{} `json:"Edges,omitempty"`
}

// Segment represents the Segment schema from the OpenAPI specification
type Segment struct {
	Document interface{} `json:"Document,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// UnprocessedStatistics represents the UnprocessedStatistics schema from the OpenAPI specification
type UnprocessedStatistics struct {
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Message interface{} `json:"Message,omitempty"`
	Rulename interface{} `json:"RuleName,omitempty"`
}

// FaultStatistics represents the FaultStatistics schema from the OpenAPI specification
type FaultStatistics struct {
	Othercount interface{} `json:"OtherCount,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
}

// GetInsightResult represents the GetInsightResult schema from the OpenAPI specification
type GetInsightResult struct {
	Insight interface{} `json:"Insight,omitempty"`
}

// AvailabilityZoneDetail represents the AvailabilityZoneDetail schema from the OpenAPI specification
type AvailabilityZoneDetail struct {
	Name interface{} `json:"Name,omitempty"`
}

// GetInsightImpactGraphResult represents the GetInsightImpactGraphResult schema from the OpenAPI specification
type GetInsightImpactGraphResult struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Servicegraphendtime interface{} `json:"ServiceGraphEndTime,omitempty"`
	Servicegraphstarttime interface{} `json:"ServiceGraphStartTime,omitempty"`
	Services interface{} `json:"Services,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
	Insightid interface{} `json:"InsightId,omitempty"`
}

// UpdateGroupRequest represents the UpdateGroupRequest schema from the OpenAPI specification
type UpdateGroupRequest struct {
	Filterexpression interface{} `json:"FilterExpression,omitempty"`
	Grouparn interface{} `json:"GroupARN,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
	Insightsconfiguration interface{} `json:"InsightsConfiguration,omitempty"`
}

// UpdateSamplingRuleRequest represents the UpdateSamplingRuleRequest schema from the OpenAPI specification
type UpdateSamplingRuleRequest struct {
	Samplingruleupdate interface{} `json:"SamplingRuleUpdate"`
}

// InsightEvent represents the InsightEvent schema from the OpenAPI specification
type InsightEvent struct {
	Topanomalousservices interface{} `json:"TopAnomalousServices,omitempty"`
	Clientrequestimpactstatistics interface{} `json:"ClientRequestImpactStatistics,omitempty"`
	Eventtime interface{} `json:"EventTime,omitempty"`
	Rootcauseservicerequestimpactstatistics interface{} `json:"RootCauseServiceRequestImpactStatistics,omitempty"`
	Summary interface{} `json:"Summary,omitempty"`
}

// ListTagsForResourceResponse represents the ListTagsForResourceResponse schema from the OpenAPI specification
type ListTagsForResourceResponse struct {
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
}

// TraceSummary represents the TraceSummary schema from the OpenAPI specification
type TraceSummary struct {
	Hasthrottle interface{} `json:"HasThrottle,omitempty"`
	Http interface{} `json:"Http,omitempty"`
	Id interface{} `json:"Id,omitempty"`
	Matchedeventtime interface{} `json:"MatchedEventTime,omitempty"`
	Serviceids interface{} `json:"ServiceIds,omitempty"`
	Annotations interface{} `json:"Annotations,omitempty"`
	Entrypoint interface{} `json:"EntryPoint,omitempty"`
	Errorrootcauses interface{} `json:"ErrorRootCauses,omitempty"`
	Resourcearns interface{} `json:"ResourceARNs,omitempty"`
	Duration interface{} `json:"Duration,omitempty"`
	Responsetime interface{} `json:"ResponseTime,omitempty"`
	Users interface{} `json:"Users,omitempty"`
	Availabilityzones interface{} `json:"AvailabilityZones,omitempty"`
	Faultrootcauses interface{} `json:"FaultRootCauses,omitempty"`
	Ispartial interface{} `json:"IsPartial,omitempty"`
	Haserror interface{} `json:"HasError,omitempty"`
	Instanceids interface{} `json:"InstanceIds,omitempty"`
	Revision interface{} `json:"Revision,omitempty"`
	Hasfault interface{} `json:"HasFault,omitempty"`
	Responsetimerootcauses interface{} `json:"ResponseTimeRootCauses,omitempty"`
}

// GetInsightSummariesResult represents the GetInsightSummariesResult schema from the OpenAPI specification
type GetInsightSummariesResult struct {
	Insightsummaries interface{} `json:"InsightSummaries,omitempty"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
}

// UpdateGroupResult represents the UpdateGroupResult schema from the OpenAPI specification
type UpdateGroupResult struct {
	Group interface{} `json:"Group,omitempty"`
}

// DeleteGroupRequest represents the DeleteGroupRequest schema from the OpenAPI specification
type DeleteGroupRequest struct {
	Grouparn interface{} `json:"GroupARN,omitempty"`
	Groupname interface{} `json:"GroupName,omitempty"`
}

// CreateGroupRequest represents the CreateGroupRequest schema from the OpenAPI specification
type CreateGroupRequest struct {
	Groupname interface{} `json:"GroupName"`
	Insightsconfiguration interface{} `json:"InsightsConfiguration,omitempty"`
	Tags interface{} `json:"Tags,omitempty"`
	Filterexpression interface{} `json:"FilterExpression,omitempty"`
}

// EdgeStatistics represents the EdgeStatistics schema from the OpenAPI specification
type EdgeStatistics struct {
	Okcount interface{} `json:"OkCount,omitempty"`
	Totalcount interface{} `json:"TotalCount,omitempty"`
	Totalresponsetime interface{} `json:"TotalResponseTime,omitempty"`
	Errorstatistics interface{} `json:"ErrorStatistics,omitempty"`
	Faultstatistics interface{} `json:"FaultStatistics,omitempty"`
}

// AttributeMap represents the AttributeMap schema from the OpenAPI specification
type AttributeMap struct {
}

// DeleteSamplingRuleResult represents the DeleteSamplingRuleResult schema from the OpenAPI specification
type DeleteSamplingRuleResult struct {
	Samplingrulerecord interface{} `json:"SamplingRuleRecord,omitempty"`
}

// Edge represents the Edge schema from the OpenAPI specification
type Edge struct {
	Receivedeventagehistogram interface{} `json:"ReceivedEventAgeHistogram,omitempty"`
	Referenceid interface{} `json:"ReferenceId,omitempty"`
	Responsetimehistogram interface{} `json:"ResponseTimeHistogram,omitempty"`
	Starttime interface{} `json:"StartTime,omitempty"`
	Summarystatistics interface{} `json:"SummaryStatistics,omitempty"`
	Aliases interface{} `json:"Aliases,omitempty"`
	Edgetype interface{} `json:"EdgeType,omitempty"`
	Endtime interface{} `json:"EndTime,omitempty"`
}

// InsightImpactGraphEdge represents the InsightImpactGraphEdge schema from the OpenAPI specification
type InsightImpactGraphEdge struct {
	Referenceid interface{} `json:"ReferenceId,omitempty"`
}

// UntagResourceResponse represents the UntagResourceResponse schema from the OpenAPI specification
type UntagResourceResponse struct {
}

// UnprocessedTraceSegment represents the UnprocessedTraceSegment schema from the OpenAPI specification
type UnprocessedTraceSegment struct {
	Message interface{} `json:"Message,omitempty"`
	Errorcode interface{} `json:"ErrorCode,omitempty"`
	Id interface{} `json:"Id,omitempty"`
}

// TelemetryRecord represents the TelemetryRecord schema from the OpenAPI specification
type TelemetryRecord struct {
	Timestamp interface{} `json:"Timestamp"`
	Backendconnectionerrors interface{} `json:"BackendConnectionErrors,omitempty"`
	Segmentsreceivedcount interface{} `json:"SegmentsReceivedCount,omitempty"`
	Segmentsrejectedcount interface{} `json:"SegmentsRejectedCount,omitempty"`
	Segmentssentcount interface{} `json:"SegmentsSentCount,omitempty"`
	Segmentsspillovercount interface{} `json:"SegmentsSpilloverCount,omitempty"`
}

// PutEncryptionConfigRequest represents the PutEncryptionConfigRequest schema from the OpenAPI specification
type PutEncryptionConfigRequest struct {
	Keyid interface{} `json:"KeyId,omitempty"`
	TypeField interface{} `json:"Type"`
}

// GetInsightImpactGraphRequest represents the GetInsightImpactGraphRequest schema from the OpenAPI specification
type GetInsightImpactGraphRequest struct {
	Endtime interface{} `json:"EndTime"`
	Insightid interface{} `json:"InsightId"`
	Nexttoken interface{} `json:"NextToken,omitempty"`
	Starttime interface{} `json:"StartTime"`
}
