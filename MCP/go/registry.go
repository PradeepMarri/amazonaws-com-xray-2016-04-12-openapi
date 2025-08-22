package main

import (
	"github.com/aws-x-ray/mcp-server/config"
	"github.com/aws-x-ray/mcp-server/models"
	tools_tracegraph "github.com/aws-x-ray/mcp-server/tools/tracegraph"
	tools_deletesamplingrule "github.com/aws-x-ray/mcp-server/tools/deletesamplingrule"
	tools_tracesegments "github.com/aws-x-ray/mcp-server/tools/tracesegments"
	tools_creategroup "github.com/aws-x-ray/mcp-server/tools/creategroup"
	tools_encryptionconfig "github.com/aws-x-ray/mcp-server/tools/encryptionconfig"
	tools_deleteresourcepolicy "github.com/aws-x-ray/mcp-server/tools/deleteresourcepolicy"
	tools_insightimpactgraph "github.com/aws-x-ray/mcp-server/tools/insightimpactgraph"
	tools_insight "github.com/aws-x-ray/mcp-server/tools/insight"
	tools_tagresource "github.com/aws-x-ray/mcp-server/tools/tagresource"
	tools_insightsummaries "github.com/aws-x-ray/mcp-server/tools/insightsummaries"
	tools_deletegroup "github.com/aws-x-ray/mcp-server/tools/deletegroup"
	tools_groups "github.com/aws-x-ray/mcp-server/tools/groups"
	tools_putresourcepolicy "github.com/aws-x-ray/mcp-server/tools/putresourcepolicy"
	tools_createsamplingrule "github.com/aws-x-ray/mcp-server/tools/createsamplingrule"
	tools_putencryptionconfig "github.com/aws-x-ray/mcp-server/tools/putencryptionconfig"
	tools_listresourcepolicies "github.com/aws-x-ray/mcp-server/tools/listresourcepolicies"
	tools_getsamplingrules "github.com/aws-x-ray/mcp-server/tools/getsamplingrules"
	tools_servicegraph "github.com/aws-x-ray/mcp-server/tools/servicegraph"
	tools_traces "github.com/aws-x-ray/mcp-server/tools/traces"
	tools_tracesummaries "github.com/aws-x-ray/mcp-server/tools/tracesummaries"
	tools_telemetryrecords "github.com/aws-x-ray/mcp-server/tools/telemetryrecords"
	tools_listtagsforresource "github.com/aws-x-ray/mcp-server/tools/listtagsforresource"
	tools_timeseriesservicestatistics "github.com/aws-x-ray/mcp-server/tools/timeseriesservicestatistics"
	tools_updategroup "github.com/aws-x-ray/mcp-server/tools/updategroup"
	tools_updatesamplingrule "github.com/aws-x-ray/mcp-server/tools/updatesamplingrule"
	tools_insightevents "github.com/aws-x-ray/mcp-server/tools/insightevents"
	tools_samplingstatisticsummaries "github.com/aws-x-ray/mcp-server/tools/samplingstatisticsummaries"
	tools_samplingtargets "github.com/aws-x-ray/mcp-server/tools/samplingtargets"
	tools_untagresource "github.com/aws-x-ray/mcp-server/tools/untagresource"
	tools_getgroup "github.com/aws-x-ray/mcp-server/tools/getgroup"
)

func GetAll(cfg *config.APIConfig) []models.Tool {
	return []models.Tool{
		tools_tracegraph.CreateGettracegraphTool(cfg),
		tools_deletesamplingrule.CreateDeletesamplingruleTool(cfg),
		tools_tracesegments.CreatePuttracesegmentsTool(cfg),
		tools_creategroup.CreateCreategroupTool(cfg),
		tools_encryptionconfig.CreateGetencryptionconfigTool(cfg),
		tools_deleteresourcepolicy.CreateDeleteresourcepolicyTool(cfg),
		tools_insightimpactgraph.CreateGetinsightimpactgraphTool(cfg),
		tools_insight.CreateGetinsightTool(cfg),
		tools_tagresource.CreateTagresourceTool(cfg),
		tools_insightsummaries.CreateGetinsightsummariesTool(cfg),
		tools_deletegroup.CreateDeletegroupTool(cfg),
		tools_groups.CreateGetgroupsTool(cfg),
		tools_putresourcepolicy.CreatePutresourcepolicyTool(cfg),
		tools_createsamplingrule.CreateCreatesamplingruleTool(cfg),
		tools_putencryptionconfig.CreatePutencryptionconfigTool(cfg),
		tools_listresourcepolicies.CreateListresourcepoliciesTool(cfg),
		tools_getsamplingrules.CreateGetsamplingrulesTool(cfg),
		tools_servicegraph.CreateGetservicegraphTool(cfg),
		tools_traces.CreateBatchgettracesTool(cfg),
		tools_tracesummaries.CreateGettracesummariesTool(cfg),
		tools_telemetryrecords.CreatePuttelemetryrecordsTool(cfg),
		tools_listtagsforresource.CreateListtagsforresourceTool(cfg),
		tools_timeseriesservicestatistics.CreateGettimeseriesservicestatisticsTool(cfg),
		tools_updategroup.CreateUpdategroupTool(cfg),
		tools_updatesamplingrule.CreateUpdatesamplingruleTool(cfg),
		tools_insightevents.CreateGetinsighteventsTool(cfg),
		tools_samplingstatisticsummaries.CreateGetsamplingstatisticsummariesTool(cfg),
		tools_samplingtargets.CreateGetsamplingtargetsTool(cfg),
		tools_untagresource.CreateUntagresourceTool(cfg),
		tools_getgroup.CreateGetgroupTool(cfg),
	}
}
