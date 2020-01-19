// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

import (
	"encoding/json"
	"testing"
)

func TestIndexSerialization(t *testing.T) {
	tests := []struct {
		desc        string
		i           *Index
		includeName bool
		expected    string
	}{
		// #0
		{
			desc:        "Include Name with StaticSettings.",
			i:           NewIndex().NumberOfShards(1).ShardCheckOnStartup("false").Codec("default").RoutingPartitionSize(1).LoadFixedBitsetFiltersEagerly(true),
			includeName: true,
			expected:    `{"index":{"codec":"default","load_fixed_bitset_filters_eagerly":true,"number_of_shards":1,"routing_partition_size":1,"shard.check_on_startup":"false"}}`,
		},
		// #1
		{
			desc:        "Include Name with DynamicSettings.",
			i:           NewIndex().NumberOfReplicas(1).AutoExpandReplicas("false").SearchIdleAfter("30s").RefreshInterval("1s").MaxResultWindow(10000).MaxInnerResultWindow(100).MaxRescoreWindow(10000).MaxDocvalueFieldsSearch(100).MaxScriptFields(32).MaxNGramDiff(1).MaxShingleDiff(3).BlocksReadOnly(true).BlocksReadOnlyAllowDelete(true).BlocksRead(true).BlocksWrite(true).BlocksMetadata(true).MaxRefreshListeners(1000).AnalyzeMaxTokenCount(10000).HighlightMaxAnalyzedOffset(1000000).MaxTermsCount(65536).MaxRegexLength(1000).RoutingAllocationEnable("all").RoutingRebalanceEnable("all").GCDeletes("60s").DefaultPipeline("_none").FinalPipeline("_none"),
			includeName: true,
			expected:    `{"index":{"analyze.max_token_count":10000,"auto_expand_replicas":"false","blocks.metadata":true,"blocks.read":true,"blocks.read_only":true,"blocks.read_only_allow_delete":true,"blocks.write":true,"default_pipeline":"_none","final_pipeline":"_none","gc_deletes":"60s","highlight.max_analyzed_offset":1000000,"max_docvalue_fields_search":100,"max_inner_result_window":100,"max_ngram_diff":1,"max_refresh_listeners":1000,"max_regex_length":1000,"max_rescore_window":10000,"max_result_window":10000,"max_script_fields":32,"max_shingle_diff":3,"max_terms_count":65536,"number_of_replicas":1,"refresh_interval":"1s","routing.allocation.enable":"all","routing.rebalance.enable":"all","search.idle.after":"30s"}}`,
		},
		// #2
		{
			desc:        "Include Name with Analysis.",
			i:           NewIndex().Analysis(NewAnalysis().DefaultAnalyzer(NewAnalyzerWhitespace("")).Analyzer(NewAnalyzerCustom("custom_keyword", "standard"))),
			includeName: true,
			expected:    `{"index":{"analysis":{"analyzer":{"custom_keyword":{"tokenizer":"standard","type":"custom"},"default":{"type":"whitespace"}}}}}`,
		},
		// #3
		{
			desc:        "Include Name with IndexShardAllocation.",
			i:           NewIndex().RoutingAllocation(NewRoutingAllocation("include", "_name", "node_name"), NewRoutingAllocation("require", "_id", "id_1", "id_2")).UnassignedNodeLeftDelayedTimeout("10s").Priority(10).RoutingAllocationTotalShardsPerNode(1),
			includeName: true,
			expected:    `{"index":{"priority":10,"routing.allocation.include._name":"node_name","routing.allocation.require._id":"id_1,id_2","routing.allocation.total_shards_per_node":1,"unassigned.node_left.delayed_timeout":"10s"}}`,
		},
		// #4
		{
			desc:        "Include Name with Mappings.",
			i:           NewIndex().Mappings(NewMappings().DynamicTemplates(NewDynamicTemplate("template_1").MatchMappingType("string").Mapping(NewDatatypeText("").Analyzer("{name}")))).MappingTotalFieldsLimit(1000).MappingDepthLimit(5).MappingNestedFieldsLimit(5).MappingNestedObjectsLimit(5).MappingFieldNameLengthLimit(10000),
			includeName: true,
			expected:    `{"index":{"mapping.depth.limit":5,"mapping.field_name_length.limit":10000,"mapping.nested_fields.limit":5,"mapping.nested_objects.limit":5,"mapping.total_fields.limit":1000,"mappings":{"dynamic_templates":[{"template_1":{"mapping":{"analyzer":"{name}","type":"text"},"match_mapping_type":"string"}}]}}}`,
		},
		// #5
		{
			desc:        "Include Name with Merging.",
			i:           NewIndex().MergeSchedulerMaxThreadCount(100),
			includeName: true,
			expected:    `{"index":{"merge.scheduler.max_thread_count":100}}`,
		},
		// #6
		{
			desc:        "Include Name with Similarity.",
			i:           NewIndex().Similarity(NewSimilarityScripted("field_1").Script(NewScript("double idf = Math.log((field.docCount+1.0)/(term.docFreq+1.0)) + 1.0; return query.boost * idf;")), NewSimilarityIB("field_2").Distribution("ll").Lambda("df")).DefaultSimilarity(NewSimilarityBM25("").K1(1.2)),
			includeName: true,
			expected:    `{"index":{"similarity":{"default":{"k1":1.2,"type":"BM25"},"field_1":{"script":{"source":"double idf = Math.log((field.docCount+1.0)/(term.docFreq+1.0)) + 1.0; return query.boost * idf;"},"type":"scripted"},"field_2":{"distribution":"ll","lambda":"df","type":"IB"}}}}`,
		},
		// #7
		{
			desc:        "Include Name with Slowlog.",
			i:           NewIndex().SearchSlowlogThreshold(NewSlowlogThreshold("search", "query", "warn", "10s").Value("5s")).SearchSlowlogLevel("info").IndexingSlowlogThreshold(NewSlowlogThreshold("indexing", "index", "warn", "10s").Value("5s")).IndexingSlowlogLevel("warn").IndexingSlowlogSource("1000").IndexingSlowlogReformat(false),
			includeName: true,
			expected:    `{"index":{"indexing.slowlog.level":"warn","indexing.slowlog.reformat":false,"indexing.slowlog.source":"1000","indexing.slowlog.threshold.index.warn":"5s","search.slowlog.level":"info","search.slowlog.threshold.query.warn":"5s"}}`,
		},
		// #8
		{
			desc:        "Include Name with Store.",
			i:           NewIndex().StoreType("fs").StorePreload("nvd"),
			includeName: true,
			expected:    `{"index":{"store.preload":["nvd"],"store.type":"fs"}}`,
		},
		// #9
		{
			desc:        "Include Name with Translog.",
			i:           NewIndex().TranslogSyncInterval("5s").TranslogDurability("request").TranslogFlushThresholdSize("512mb").TranslogRetentionSize("512mb").TranslogRetentionAge("12h"),
			includeName: true,
			expected:    `{"index":{"translog.durability":"request","translog.flush_threshold_size":"512mb","translog.retention.age":"12h","translog.retention.size":"512mb","translog.sync_interval":"5s"}}`,
		},
		// #10
		{
			desc:        "Include Name with HistoryRetention.",
			i:           NewIndex().SoftDeletesEnabled(true).SoftDeletesRetentionLeasePeriod("12h"),
			includeName: true,
			expected:    `{"index":{"soft_deletes.enabled":true,"soft_deletes.retention_lease.period":"12h"}}`,
		},
		// #11
		{
			desc:        "Include Name with IndexSorting.",
			i:           NewIndex().SortField("date").SortOrder("desc").SortMode("min").SortMissing("_last"),
			includeName: true,
			expected:    `{"index":{"sort.field":"date","sort.missing":"_last","sort.mode":"min","sort.order":"desc"}}`,
		},
		// #12
		{
			desc:        "Include Name with IndexLifecycleManagement.",
			i:           NewIndex().LifecycleName("lifecycle_name").LifecycleRolloverAlias("lifecycle_alias").LifecycleParseOriginationDate(true).LifecycleOriginationDate(1579442569),
			includeName: true,
			expected:    `{"index":{"lifecycle.name":"lifecycle_name","lifecycle.origination_date":1579442569,"lifecycle.parse_origination_date":true,"lifecycle.rollover_alias":"lifecycle_alias"}}`,
		},
	}
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			src, err := test.i.Source(test.includeName)
			if err != nil {
				t.Fatal(err)
			}
			data, err := json.Marshal(src)
			if err != nil {
				t.Fatalf("marshaling to JSON failed: %v", err)
			}
			got := string(data)
			if got != test.expected {
				t.Errorf("expected\n%s\n,got:\n%s", test.expected, got)
			}
		})
	}
}
