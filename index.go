// Copyright (c) KwanJunWen
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package estemplate

// Index index module created per index and control all aspects related to an index.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules.html
// for details.
type Index struct {
	// static settings
	numberOfShards                *int
	shardCheckOnStartup           string
	codec                         string
	routingPartitionSize          *int
	loadFixedBitsetFiltersEagerly *bool

	// dynamic settings
	numberOfReplicas           *int
	autoExpandReplicas         string
	searchIdleAfter            string
	refreshInterval            string
	maxResultWindow            *int
	maxInnerResultWindow       *int
	maxRescoreWindow           *int
	maxDocvalueFieldsSearch    *int
	maxScriptFields            *int
	maxNGramDiff               *int
	maxShingleDiff             *int
	blocksReadOnly             *bool
	blocksReadOnlyAllowDelete  *bool
	blocksRead                 *bool
	blocksWrite                *bool
	blocksMetadata             *bool
	maxRefreshListeners        *int
	analyzeMaxTokenCount       *int
	highlightMaxAnalyzedOffset *int
	maxTermsCount              *int
	maxRegexLength             *int
	routingAllocationEnable    string
	routingRebalanceEnable     string
	gcDeletes                  string
	defaultPipeline            string
	finalPipeline              string

	// analysis
	analysis *Analysis

	// index shard allocation
	routingAllocation                   []*RoutingAllocation
	unassignedNodeLeftDelayedTimeout    string
	priority                            *int
	routingAllocationTotalShardsPerNode *int

	// mapping
	mappings                    *Mappings
	mappingTotalFieldsLimit     *int
	mappingDepthLimit           *int
	mappingNestedFieldsLimit    *int
	mappingNestedObjectsLimit   *int
	mappingFieldNameLengthLimit *int

	// merging
	mergeSchedulerMaxThreadCount *int

	// similarity
	defaultSimilarity Similarity
	similarity        []Similarity

	// slowlog
	searchSlowlogThreshold []*SlowlogThreshold
	searchSlowlogLevel     string

	indexingSlowlogThreshold []*SlowlogThreshold
	indexingSlowlogLevel     string
	indexingSlowlogSource    string
	indexingSlowlogReformat  *bool

	// store
	storeType    string
	storePreload []string

	// translog
	translogSyncInterval       string
	translogDurability         string
	translogFlushThresholdSize string
	translogRetentionSize      string
	translogRetentionAge       string

	// history retention
	softDeletesEnabled              *bool
	softDeletesRetentionLeasePeriod string

	// index sorting
	sortField   []string
	sortOrder   string
	sortMode    string
	sortMissing string

	// index lifecycle management
	lifecycleName                 string
	lifecycleRolloverAlias        string
	lifecycleParseOriginationDate *bool
	lifecycleOriginationDate      *int
}

// NewIndex initializes a new Index.
func NewIndex() *Index {
	return &Index{}
}

// * <-- Static Settings -->
// Static settings can only be set at index creation time or on a closed index.
// ! Changing static or dynamic index settings on a closed index could result in incorrect settings that are
// impossible to rectify without deleting and recreating the index.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules.html#_static_index_settings
// for details.

// NumberOfShards sets the number of primary shards that the index should have.
// Can only be set at index creation time. It cannot be changed on a closed index.
// Number of shards are limit to 1024 per index. This limitation is a safety limit
// to prevent accidental creation of indices that can destabilize a cluster due to
// resource allocation. The limit can be modified by specifying
// `export ES_JAVA_OPTS="-Des.index.max_number_of_shards=128"` system property
// on every node that is part of the cluster.
// Defaults to 1.
func (i *Index) NumberOfShards(numberOfShards int) *Index {
	i.numberOfShards = &numberOfShards
	return i
}

// ShardCheckOnStartup sets whether or not shards should be checked for corruption
// before opening. When corruption is detected, it will prevent the shard from being
// opened.
// Can be set to the following values:
// "false" - Don't check for corruption when opening a shard.
// "checksum" - Check for physical corruption.
// "true" - Check for both physical and logical corruption. This is much more expensive
// 					in terms of CPU and memory usage
// Defaults to "false".
func (i *Index) ShardCheckOnStartup(shardCheckOnStartup string) *Index {
	i.shardCheckOnStartup = shardCheckOnStartup
	return i
}

// Codec sets the compression type for the index.
// Can be set to the following values:
// "default" - LZ4
// "best_compression" - DEFLATE, higher cmopression ratio, at the expense of slower
// 											stored fields performance.
// Defaults to "default".
func (i *Index) Codec(codec string) *Index {
	i.codec = codec
	return i
}

// RoutingPartitionSize sets the number of shards a custom routing value can go to.
// Can only be set at index creation time.This value must be less than the
// `number_of_shards` parameter unless the `number_of_shards` value is also 1.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/current/mapping-routing-field.html#routing-index-partition
// for details.
// Defaults to 1.
func (i *Index) RoutingPartitionSize(routingPartitionSize int) *Index {
	i.routingPartitionSize = &routingPartitionSize
	return i
}

// LoadFixedBitsetFiltersEagerly sets whether cache filters are pre-loaded for nested queries.
// Defaults to true.
func (i *Index) LoadFixedBitsetFiltersEagerly(loadFixedBitsetFiltersEagerly bool) *Index {
	i.loadFixedBitsetFiltersEagerly = &loadFixedBitsetFiltersEagerly
	return i
}

// * <-- Dynamic Settings -->
// Dynamic settings can be changed on a live index using the update-index-settings API.
// ! Changing static or dynamic index settings on a closed index could result in incorrect settings that are
// impossible to rectify without deleting and recreating the index.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules.html#dynamic-index-settings
// for details.

// NumberOfReplicas sets the number of replicas each primary shard has.
// Defaults to 1.
func (i *Index) NumberOfReplicas(numberOfReplicas int) *Index {
	i.numberOfReplicas = &numberOfReplicas
	return i
}

// AutoExpandReplicas sets the number of replicas to be auto-expanded based on the number
// of data nodes in the cluster. Set to a dash delimited lower and upper bound (e.g. "0-5")
// or use "all" for the upper bound (e.g. "0-all").
//
// Defaults to "false" (disabled).
func (i *Index) AutoExpandReplicas(autoExpandReplicas string) *Index {
	i.autoExpandReplicas = autoExpandReplicas
	return i
}

// SearchIdleAfter sets how long a shard can not receive a search or get request until
// it's considered search idle.
// Defaults to "30s".
func (i *Index) SearchIdleAfter(searchIdleAfter string) *Index {
	i.searchIdleAfter = searchIdleAfter
	return i
}

// RefreshInterval sets how often to perform a refresh operation, which makes recent changes to
// the index visible to search. Can be set to "-1" to disable refresh. If this setting is not
// explicitly set, shards that haven't seen search traffic for at least `index.search.idle.after`
// seconds will not receive background refreshes until they receive a search request. Searches that
// hit an idle shard where a refresh is pending will wait for the next background refresh (within "1s").
// This behavior aims to automatically optimize bulk indexing in the default case when no searches are
// performed. In order to opt out of this behaviour an explicit value of "1s" should set as the refresh
// interval.
// Defaults to "1s".
func (i *Index) RefreshInterval(refreshInterval string) *Index {
	i.refreshInterval = refreshInterval
	return i
}

// MaxResultWindow sets the maximum value of `from + size` for searches to this index. Search requests
// take heap memory and time proportional to `from + size` and this limits that memory.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-scroll
// and https://www.elastic.co/guide/en/elasticsearch/reference/7.5/search-request-body.html#request-body-search-search-after
// for details.
// Defaults to 10000.
func (i *Index) MaxResultWindow(maxResultWindow int) *Index {
	i.maxResultWindow = &maxResultWindow
	return i
}

// MaxInnerResultWindow sets the maximum value of `from + size` for inner hits definition and top hits
// aggregations to this index. Inner hits and top hits aggregation take heap memory and time proportional
// to `from + size` and this limits that memory.
// Defaults to 100.
func (i *Index) MaxInnerResultWindow(maxInnerResultWindow int) *Index {
	i.maxInnerResultWindow = &maxInnerResultWindow
	return i
}

// MaxRescoreWindow sets the maximum value of `window_size` for rescore requests in searches of this index.
// Search requests take heap memory and time proportional to `max(window_size, from + size` and this limits
// that memory.
// Defaults to `index.max_result_window` which defaults to 10000.
func (i *Index) MaxRescoreWindow(maxRescoreWindow int) *Index {
	i.maxRescoreWindow = &maxRescoreWindow
	return i
}

// MaxDocvalueFieldsSearch sets the maximum number of `docvalue_fields` that are allowed in a query.
// Doc-value fields are costly since they might incur a per-field per-document seek.
// Defaults to 100.
func (i *Index) MaxDocvalueFieldsSearch(maxDocvalueFieldsSearch int) *Index {
	i.maxDocvalueFieldsSearch = &maxDocvalueFieldsSearch
	return i
}

// MaxScriptFields sets the maximum number of `script_fields` that are allowed in a query.
// Defaults to 32.
func (i *Index) MaxScriptFields(maxScriptFields int) *Index {
	i.maxScriptFields = &maxScriptFields
	return i
}

// MaxNGramDiff sets the maximum allowed difference between `min_gram` and `max_gram` for `NGramTokenizer`
// and `NGramTokenFilter`.
// Defaults to 1.
func (i *Index) MaxNGramDiff(maxNGramDiff int) *Index {
	i.maxNGramDiff = &maxNGramDiff
	return i
}

// MaxShingleDiff sets the maximum allowed difference between `max_shingle_size` and `min_shingle_size` for
// `ShingleTokenFilter`.
// Defaults to 3.
func (i *Index) MaxShingleDiff(maxShingleDiff int) *Index {
	i.maxShingleDiff = &maxShingleDiff
	return i
}

// BlocksReadOnly sets whether to make the index and index metadata read only.
// Set to false to allow writes and metadata changes.
func (i *Index) BlocksReadOnly(blocksReadOnly bool) *Index {
	i.blocksReadOnly = &blocksReadOnly
	return i
}

// BlocksReadOnlyAllowDelete sets whether to make the index and inexmetadata read only, but also allows deleting the
// index to free up resources. The disk-based shard allocator may add and remove this block automatically.
func (i *Index) BlocksReadOnlyAllowDelete(blocksReadOnlyAllowDelete bool) *Index {
	i.blocksReadOnlyAllowDelete = &blocksReadOnlyAllowDelete
	return i
}

// BlocksRead sets whether or not to disable read operations against the index.
func (i *Index) BlocksRead(blocksRead bool) *Index {
	i.blocksRead = &blocksRead
	return i
}

// BlocksWrite sets whether or not to disable data write operations against the index. Unlike `read_only`,
// this setting does not affect metadata. For instance, you can close an index with a `write` block, but not
// an index with a `read_only` block.
func (i *Index) BlocksWrite(blocksWrite bool) *Index {
	i.blocksWrite = &blocksWrite
	return i
}

// BlocksMetadata sets whether or not to disable index metadata reads and writes.
func (i *Index) BlocksMetadata(blocksMetadata bool) *Index {
	i.blocksMetadata = &blocksMetadata
	return i
}

// MaxRefreshListeners sets the maximum number of refresh listeners available on each shard of the index.
// These listeners are used to implement `refresh=wait_for`.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/docs-refresh.html
// for details.
func (i *Index) MaxRefreshListeners(maxRefreshListeners int) *Index {
	i.maxRefreshListeners = &maxRefreshListeners
	return i
}

// AnalyzeMaxTokenCount sets the maximum number of tokens that can be produced using `_analyze` API.
// Defaults to 10000.
func (i *Index) AnalyzeMaxTokenCount(analyzeMaxTokenCount int) *Index {
	i.analyzeMaxTokenCount = &analyzeMaxTokenCount
	return i
}

// HighlightMaxAnalyzedOffset sets the maximum number of characters that will be analyzed for a highlight request.
// This setting is only applicable when highlighting is requested on a text that was indexed without offsets or term
// vectors.
// Defaults to 1000000.
func (i *Index) HighlightMaxAnalyzedOffset(highlightMaxAnalyzedOffset int) *Index {
	i.highlightMaxAnalyzedOffset = &highlightMaxAnalyzedOffset
	return i
}

// MaxTermsCount sets the maximum number of terms that can be used in Terms Query.
// Defaults to 65536.
func (i *Index) MaxTermsCount(maxTermsCount int) *Index {
	i.maxTermsCount = &maxTermsCount
	return i
}

// MaxRegexLength sets the maximum length of regex that can be used in Regexp Query.
// Defaults to 1000.
func (i *Index) MaxRegexLength(maxRegexLength int) *Index {
	i.maxRegexLength = &maxRegexLength
	return i
}

// RoutingAllocationEnable sets the shard allocation for this index.
// Can be set to the following values:
// "all" - Allows shard allocation for all shards.
// "primaries" - Allow shard allocation only for primary shards.
// "new_primaries" - Allows shard allocation only for newly-created primary shards.
// "none" - No shard allocation is allowed.
// Defaults to "all".
func (i *Index) RoutingAllocationEnable(routingAllocationEnable string) *Index {
	i.routingAllocationEnable = routingAllocationEnable
	return i
}

// RoutingRebalanceEnable sets whether or not to enable shard rebalancing for this index.
// Can be set to the following values:
// "all" - Allows shard allocation for all shards.
// "primaries" - Allow shard allocation only for primary shards.
// "new_primaries" - Allows shard allocation only for newly-created primary shards.
// "none" - No shard allocation is allowed.
// Defaults to "all".
func (i *Index) RoutingRebalanceEnable(routingRebalanceEnable string) *Index {
	i.routingRebalanceEnable = routingRebalanceEnable
	return i
}

// GCDeletes sets the length of time that a deleted document's version number remains available for
// further versioned operations.
// Defaults to "60s".
func (i *Index) GCDeletes(gcDeletes string) *Index {
	i.gcDeletes = gcDeletes
	return i
}

// DefaultPipeline sets the default ingest node pipeline for this index. Index requests will fail if
// the default pipeline is set and the pipeline does not exist. The default may be overridden using the
// `pipeline` parameter. The special pipeline name "_none" indicates no ingest pipeline should be run.
func (i *Index) DefaultPipeline(defaultPipeline string) *Index {
	i.defaultPipeline = defaultPipeline
	return i
}

// FinalPipeline sets the final ingest node pipeline for this index. Index requests will fail if
// the default pipeline is set and the pipeline does not exist. The final pipeline always runs after the
// request pipeline (if specified) and the default pipeline (if it exists). The special pipeline name "_none"
// indicates no ingest pipeline should be run.
func (i *Index) FinalPipeline(finalPipeline string) *Index {
	i.finalPipeline = finalPipeline
	return i
}

// * <-- Analysis Settings -->
// Analysis settings define analyzers, tokenizers, token filters and character filters for this index.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-analysis.html
// for details.

// Analysis sets the settings to define analyzers, tokenizers, token filters and character filters.
func (i *Index) Analysis(analysis *Analysis) *Index {
	i.analysis = analysis
	return i
}

// * <-- Index Shard Allocation Settings -->
// Index shard allocation settings control over where, when, and how shards are allocated to nodes.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-allocation.html
// for details.

// RoutingAllocation sets shard allocation filtering for this index.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/shard-allocation-filtering.html
// for details.
func (i *Index) RoutingAllocation(routingAllocation ...*RoutingAllocation) *Index {
	i.routingAllocation = append(i.routingAllocation, routingAllocation...)
	return i
}

// UnassignedNodeLeftDelayedTimeout sets the delay timeout allocation when a node leaves.
// Defaults to "1m".
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/delayed-allocation.html
// for details.
func (i *Index) UnassignedNodeLeftDelayedTimeout(unassignedNodeLeftDelayedTimeout string) *Index {
	i.unassignedNodeLeftDelayedTimeout = unassignedNodeLeftDelayedTimeout
	return i
}

// Priority sets the order of priority for unallocated shards to be recovered. Indices are sorted into
// priority order as follows:
// - the optional `index.priority` setting (higher to lower)
// - the index creation date (higher before lower)
// - the index name (higher before lower)
// By default, newer indices will be recovered before older indices.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/recovery-prioritization.html
// for details.
func (i *Index) Priority(priority int) *Index {
	i.priority = &priority
	return i
}

// RoutingAllocationTotalShardsPerNode sets the maximum number of shards (replicas and primaries) that will
// be allocated to a single node.
// Defaults to unbounded.
func (i *Index) RoutingAllocationTotalShardsPerNode(routingAllocationTotalShardsPerNode int) *Index {
	i.routingAllocationTotalShardsPerNode = &routingAllocationTotalShardsPerNode
	return i
}

// * <-- Mappings Settings -->
// Mappings settings define an explicit fields mapping for this index.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-mapper.html
// for details.

// Mappings sets explicit mapping for this index. Disable dynamic mapping for this index if specified.
func (i *Index) Mappings(mappings *Mappings) *Index {
	i.mappings = mappings
	return i
}

// MappingTotalFieldsLimit sets the maximum number of fields in an index. Field and object mappings, as well
// as field aliases count towards this limit.
// Defaults to 1000.
func (i *Index) MappingTotalFieldsLimit(mappingTotalFieldsLimit int) *Index {
	i.mappingTotalFieldsLimit = &mappingTotalFieldsLimit
	return i
}

// MappingDepthLimit sets the maximum depth for a field, which is measured as the number of inner objects.
// For instance, if all fields are defined at the root object level, then the depth is 1. If there is one
// object mapping, then the depth is 2, etc.
// Defaults to 20.
func (i *Index) MappingDepthLimit(mappingDepthLimit int) *Index {
	i.mappingDepthLimit = &mappingDepthLimit
	return i
}

// MappingNestedFieldsLimit sets the maximum number of distinct `nested` mappings in an index.
// Defaults to 50.
func (i *Index) MappingNestedFieldsLimit(mappingNestedFieldsLimit int) *Index {
	i.mappingNestedFieldsLimit = &mappingNestedFieldsLimit
	return i
}

// MappingNestedObjectsLimit sets the maximum number of `nested` JSON objects within a single document
// across all nested types.
// Defaults to 10000.
func (i *Index) MappingNestedObjectsLimit(mappingNestedObjectsLimit int) *Index {
	i.mappingNestedObjectsLimit = &mappingNestedObjectsLimit
	return i
}

// MappingFieldNameLengthLimit sets the maximum length of a field name. This setting isn't really something
// that addresses mappings explosion but might still be useful if you want to limit the field length. It
// usually shouldn't be necessary to set this etting. The default is okay unless a user starts to add a huge
// number of fields with really long names.
// Defaults to Long.MAX_VALUE (no limit).
func (i *Index) MappingFieldNameLengthLimit(mappingFieldNameLengthLimit int) *Index {
	i.mappingFieldNameLengthLimit = &mappingFieldNameLengthLimit
	return i
}

// * <-- Merging Settings -->
// Merging settings control over how shards are merged by the background merge process.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-merge.html
// for details.

// MergeSchedulerMaxThreadCount sets the maximum number of threads on a single shard that may be merging
// at once.
// Defaults to `Math.max(1, Math.min(4, Runtime.getRuntime().availableProcessors() / 2))`, which works well
// for a good solid-state-disk (SSD). If your index is on spinning platter drives instead, decrease this to 1.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-merge.html
// for details.
func (i *Index) MergeSchedulerMaxThreadCount(mergeSchedulerMaxThreadCount int) *Index {
	i.mergeSchedulerMaxThreadCount = &mergeSchedulerMaxThreadCount
	return i
}

// * <-- Similarities Settings -->

// DefaultSimilarity sets the default similarity settings.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-similarity.html#default-base
// for details.
func (i *Index) DefaultSimilarity(defaultSimilarity Similarity) *Index {
	i.defaultSimilarity = defaultSimilarity
	return i
}

// Similarities settings configure custom similarity settings to customize how search results are scored.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-similarity.html
// for details.

// Similarity sets custom similarity settings which configures how search results are scored for this index.
func (i *Index) Similarity(similarity ...Similarity) *Index {
	i.similarity = append(i.similarity, similarity...)
	return i
}

// * <-- Slowlog Settings -->
// Slowlog settings control over how slow queries and fetch requests are logged.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-slowlog.html
// for details.

// SearchSlowlogThreshold sets the shard level search slowlog.
func (i *Index) SearchSlowlogThreshold(searchSlowlogThreshold ...*SlowlogThreshold) *Index {
	i.searchSlowlogThreshold = append(i.searchSlowlogThreshold, searchSlowlogThreshold...)
	return i
}

// SearchSlowlogLevel sets the search slowlog level.
func (i *Index) SearchSlowlogLevel(searchSlowlogLevel string) *Index {
	i.searchSlowlogLevel = searchSlowlogLevel
	return i
}

// IndexingSlowlogThreshold sets the shard level indexing slowlog.
func (i *Index) IndexingSlowlogThreshold(indexingSlowlogThreshold ...*SlowlogThreshold) *Index {
	i.indexingSlowlogThreshold = append(i.indexingSlowlogThreshold, indexingSlowlogThreshold...)
	return i
}

// IndexingSlowlogLevel sets the indexing slowlog level.
func (i *Index) IndexingSlowlogLevel(indexingSlowlogLevel string) *Index {
	i.indexingSlowlogLevel = indexingSlowlogLevel
	return i
}

// IndexingSlowlogSource sets the indexing slowlog source. Setting this to "false" or "0" will skip logging the source
// entirely and setting it to "true" will log the entire source regardless of size.
// Defaults to "1000".
func (i *Index) IndexingSlowlogSource(indexingSlowlogSource string) *Index {
	i.indexingSlowlogSource = indexingSlowlogSource
	return i
}

// IndexingSlowlogReformat sets whether or not to preserve the original document format. Setting this to false to cause
// the source to be logged "as is" and can potentially span multiple log lines.
// Defaults to true.
func (i *Index) IndexingSlowlogReformat(indexingSlowlogReformat bool) *Index {
	i.indexingSlowlogReformat = &indexingSlowlogReformat
	return i
}

// * <-- Store Settings -->
// Store settings configure the type of filesystem used to access shard data.
// ! Expert settings.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-store.html
// for details.

// StoreType sets the file system storage types for this index which controls how index data is stored and
// accessed on disk.
// Can be set to the following values:
// "fs" - Default file system implementation. This will pick the best implementation depending on the operating
// 				enviroment, which is currently "hybridfs" on all supported systems but is subject to change.
// "simplefs" - The Simple FS type is a straightforward implementation of the file system storage (maps to Lucene
// 							SimpleFsDirectory) using a random access file. This implementation has poor concurrent performance
// 							(multiple threads will bottleneck). It is usually better to use the "niofs" when you need index persistence.
// "niofs" - The NIO FS type stores the shard index on the file system (maps to Lucune NIOFSDirectory) using NIO. It allows
// 					 multiple threads to read from the same file concurrently. It is not recommended on Windows because of a bug in
// 					 the SUN Java implementation.
// "mmapfs" - The MMap FS type stores the shard index on the file system (maps to Lucene MMapDirectory) by mapping a file into
// 						memory (mmap). Memory mapping uses up a portion of the virtual memory address space in your process equal to the
// 						size of the file being mapped. Before using this class, be sure you have allowed plenty of virtual address space.
// "hybridfs" - The "hybridfs" type is a hybrid of "niofs" and "mmapfs", which chooses the best file system type for each type
// 							of file based on the read access pattern. Currently only the Lucene term dictionary, norms and doc values files
// 							are memory mapped. All other files are opened using Lucene NIOFSDirectory. Similarly to "mmapfs" be sure you
// 							allowed plenty of virtual address space.
// ! This is an expert-only setting and may be removed in the future.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-store.html
// for details.
func (i *Index) StoreType(storeType string) *Index {
	i.storeType = storeType
	return i
}

// StorePreload sets the files extensions in which all files whose extension is in the list will be preloaded upon opening. This
// be useful to improve search performance of an index, especially when host operating system is restarted, since this causes the
// file system cache to be trashed. However note that this may slow down the opening of indices, as they will only become available
// after data have been loaded into physical memory.
// A wildcard ("*") can be used in order to indicate that all files should be preloaded, however that it is generally not useful to
// load all files into memory, in particular those for stored fields and term vectors, so a better option might be to set it to
// ["nvd", "dvd", "tim", "doc", "dim"], which will preload norms, doc values, terms dictionaries, posting lists and points, which are
// the most important parts of the index for search and aggregations.
// Defaults to [].
// ! This is an expert-only setting and the details of which may change in the future.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/preload-data-to-file-system-cache.html
// for details.
func (i *Index) StorePreload(storePreload ...string) *Index {
	i.storePreload = append(i.storePreload, storePreload...)
	return i
}

// * <-- Translog Settings -->
// Translog settings control over the transaction log and background flush operations.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-translog.html
// for details.

// TranslogSyncInterval sets how often the translog is `fsync`ed to disk and committed, regardless of write operations.
// Values less than "100ms" are not allowed.
// Defaults to "5s".
func (i *Index) TranslogSyncInterval(translogSyncInterval string) *Index {
	i.translogSyncInterval = translogSyncInterval
	return i
}

// TranslogDurability sets whether or not to `fsync` and commit the translog after every index, delete, update, or bulk request.
// Can be set to the following values:
// "request" - `fsync` and commit after every request. In the event of hardware failure, all acknowledged writes will already
// 						 have been committed to disk.
// "async" - `fsync` and commit in the background every `sync_interval`. In the event of a failure, all acknowledged writes since
// 					  the last automatic commit will be discarded.
// Defaults to "request".
func (i *Index) TranslogDurability(translogDurability string) *Index {
	i.translogDurability = translogDurability
	return i
}

// TranslogFlushThresholdSize sets the maximum total size of the operations which are not yet safely persisted in Lucene
// (i.e., are not part of a Lucene commit point), to prevent recoveries from taking too long. Once the maximum size has been
// reached a flush will happen, generating a new Lucene commit point.
// Defaults to "512mb".
func (i *Index) TranslogFlushThresholdSize(translogFlushThresholdSize string) *Index {
	i.translogFlushThresholdSize = translogFlushThresholdSize
	return i
}

// TranslogRetentionSize sets the total size of translog files to keep for each shard. Keeping more translog files increases
// the chance of performing an operation based sync when recovering a replica. If the translog files are not sufficient, replica
// recovery will fall back to a file based sync. This setting is ignored, and should not be set, if `soft deletes` are enabled.
// Soft deletes are enabled by default in indices created in Elasticsearch versions `7.0.0` and later.
// Defaults to "512mb".
func (i *Index) TranslogRetentionSize(translogRetentionSize string) *Index {
	i.translogRetentionSize = translogRetentionSize
	return i
}

// TranslogRetentionAge sets the maximum duration for which translog files are kept by each shard. Keeping more translog files
// increases the chance of performing an operation based sync when recovering replicas. If the translog files are not sufficient,
// replica recovery will fall back to a file based sync. This setting is ignored, and should not be set, if `soft deletes` are enabled.
// Soft deletes are enabled by default in indices created in Elasticsearch versions `7.0.0` and later.
// Defaults to "12h".
func (i *Index) TranslogRetentionAge(translogRetentionAge string) *Index {
	i.translogRetentionAge = translogRetentionAge
	return i
}

// * <-- History Retention Settings -->
// History retention settings control over the retention of a history of operations in the index.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-history-retention.html
// for details.

// SoftDeletesEnabled sets whether or not soft deletes are enabled on the index. Soft deletes can only be configured at index creation
// and only on indices created on or after `6.5.0`.
// Defaults to true.
func (i *Index) SoftDeletesEnabled(softDeletesEnabled bool) *Index {
	i.softDeletesEnabled = &softDeletesEnabled
	return i
}

// SoftDeletesRetentionLeasePeriod sets the maximum length of time to retain a shard history retention lease before it expires and
// the history that is retains can be discarded.
// Defaults to "12h".
func (i *Index) SoftDeletesRetentionLeasePeriod(softDeletesRetentionLeasePeriod string) *Index {
	i.softDeletesRetentionLeasePeriod = softDeletesRetentionLeasePeriod
	return i
}

// * <-- Index Sorting Settings -->
// Index sorting settings configure how the Segments inside each Shard will be sorted. By default Lucene does not apply any sort.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/index-modules-index-sorting.html
// for details.

// SortField sets a list of fields used to sort the index. Only `boolean`, `numeric`, `date` and `keyword` fields with `doc_values`
// are allowed here.
func (i *Index) SortField(sortField ...string) *Index {
	i.sortField = append(i.sortField, sortField...)
	return i
}

// SortOrder sets the sort order to use for each field.
// Can be set to the following values:
// "asc" - Ascending order
// "desc" - Descending order
func (i *Index) SortOrder(sortOrder string) *Index {
	i.sortOrder = sortOrder
	return i
}

// SortMode sets the mode option that controls what value is picked to sort the document.
// Can be set to the following values:
// "min" - Pick the lowest value
// "max" - Pick the highest value
func (i *Index) SortMode(sortMode string) *Index {
	i.sortMode = sortMode
	return i
}

// SortMissing sets the missing parameter which specifies how docs which are missing the field should be treated.
// Can be set to the following values:
// "_last" - Documents without value for the field are sorted last
// "_first" - Documents without value for the field are sorted first
func (i *Index) SortMissing(sortMissing string) *Index {
	i.sortMissing = sortMissing
	return i
}

// * <-- Lifecycle Management Settings -->
// Lifecycle management settings specify the lifecycle policy and rollover alias for the index.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/ilm-settings.html
// for details.

// LifecycleName sets the name of the policy to use to manage the index.
func (i *Index) LifecycleName(lifecycleName string) *Index {
	i.lifecycleName = lifecycleName
	return i
}

// LifecycleRolloverAlias sets the index alias to update when the index rolls over. Specify when using a policy that contains
// a rollover action. WHen the index rolls over, the alias is updated to reflect that the index is no longer the write index.
//
// See https://www.elastic.co/guide/en/elasticsearch/reference/7.5/using-policies-rollover.html
// for details.
func (i *Index) LifecycleRolloverAlias(lifecycleRolloverAlias string) *Index {
	i.lifecycleRolloverAlias = lifecycleRolloverAlias
	return i
}

// LifecycleParseOriginationDate sets whether or not the origination date will be parsed from the index name. The index format must
// match the pattern `^,*-{date_format}-\\d+`, where the `date_Format` is `yyyy.MM.dd` and the trailing digits are optional (an index
// that was rolled over would normally match the full format eg. `logs-2016.10.31-000002`). If the index name doesn't match the pattern
// the index creation will fail.
func (i *Index) LifecycleParseOriginationDate(lifecycleParseOriginationDate bool) *Index {
	i.lifecycleParseOriginationDate = &lifecycleParseOriginationDate
	return i
}

// LifecycleOriginationDate sets whether or not the timestamp will be used to calculate the index age for its phase transitions. This
// allows the users to create an index containing old data and use the original creation date of the old data to calculate the index age.
func (i *Index) LifecycleOriginationDate(lifecycleOriginationDate int) *Index {
	i.lifecycleOriginationDate = &lifecycleOriginationDate
	return i
}

// Source returns the serializable JSON for the source builder.
func (i *Index) Source(includeName bool) (interface{}, error) {
	// {
	// 	"index": {
	// 		"number_of_shards": 1,
	// 		"shard.check_on_startup": "false",
	// 		"codec": "default",
	// 		"routing_partition_size": "1",
	// 		"load_fixed_bitset_filters_eagerly": true,
	// 		"number_of_replicas": 1,
	// 		"auto_expand_replicas": "false",
	// 		"search.idle.after": "30s",
	// 		"refresh_interval": "1s",
	// 		"max_result_window": "10000",
	// 		"max_inner_result_window": "100",
	// 		"max_rescore_window": "10000",
	// 		"max_docvalue_fields_search": "100",
	// 		"max_script_fields": "32",
	// 		"max_ngram_diff": "1",
	// 		"max_shingle_diff": "3",
	// 		"blocks.read_only": false,
	// 		"blocks.read_only_allow_delete": false,
	// 		"blocks.read": true,
	// 		"blocks.write": true,
	// 		"blocks.metadata": true,
	// 		"max_refresh_listeners": 1000,
	// 		"analyze.max_token_count": 10000,
	// 		"highlight.max_analyzed_offset": 1000000,
	// 		"max_terms_count": 65536,
	// 		"max_regex_length": 1000,
	// 		"routing.allocation.enable": "all",
	// 		"routing.rebalance.enable": "all",
	// 		"gc_deletes": "60s",
	// 		"default_pipeline": "_none",
	// 		"final_pipeline": "_none",
	// 		"analysis": {
	// 			"analyzer": {
	// 				"default": {
	// 					"type": "whitespace"
	// 				},
	// 				"custom": {
	// 					"type": "custom",
	// 					"tokenizer": "standard"
	// 				}
	// 			},
	// 			"normalizer": {
	// 				"custom_normalizer": {
	// 					"type": "custom",
	// 					"char_filter": ["quote"]
	// 				}
	// 			},
	// 			"filter": {
	// 				"custom_synonym": {
	// 					"type": "synonym",
	// 					"synonyms": ["i-pod, i pod => ipod", "universe, cosmos"]
	// 				}
	// 			},
	// 			"char_filter": {
	// 				"custom_mapping": {
	// 					"type": "mapping",
	// 					"mappings": ["٠ => 0", "١ => 1", "٢ => 2"]
	// 				}
	// 			}
	// 		}
	// 		"routing.allocation.include.size": "big",
	// 		"routing.allocation.require.size": "big,medium",
	// 		"routing.allocation.exclude.rack": "rack1"
	// 		"unassigned.node_left.delayed_timeout": "5m",
	// 		"priority": 10,
	// 		"routing.allocation.total_shards_per_node": 1,
	// 		"mappings": {
	// 			"dynamic_templates": [
	// 				{
	// 					"integers": {
	// 						"match_mapping_type": "long",
	// 						"mapping": {
	// 							"type": "integer"
	// 						}
	// 					}
	// 				}
	// 			],
	// 			"date_detection": false,
	// 			"dynamic_date_formats": ["strict_date_optional_time","yyyy/MM/dd HH:mm:ss Z||yyyy/MM/dd Z"],
	// 			"numeric_detection": false,
	// 			"_source": {
	// 				"enabled": true,
	// 				"includes": [
	// 					"*.count",
	// 					"meta.*"
	// 				],
	// 				"excludes": [
	// 					"meta.description",
	// 					"meta.other.*"
	// 				]
	// 			},
	// 			"_size": {
	// 				"enabled": true
	// 			},
	// 			"_field_names": {
	// 				"enabled": true
	// 			},
	// 			"_routing": {
	// 				"required": true
	// 			},
	// 			"_meta": {
	// 				"class": "MyApp::User",
	// 				"version": {
	// 					"min": "1.0",
	// 					"max": "1.3"
	// 				}
	// 			},
	// 			"properties": {
	// 				"field_name": {
	// 					"type": "text",
	// 					"analyzer": "standard"
	// 				}
	// 			}
	// 		}
	// 		"mapping.total_fields.limit": 1000,
	// 		"mapping.depth.limit": 20,
	// 		"mapping.nested_fields.limit": 50,
	// 		"mapping.nested_objects.limit": 10000,
	// 		"mapping.field_name_length.limit": 10000,
	// 		"merge.scheduler.max_thread_count": 1,
	// 		"similarity": {
	// 			"my_similarity": {
	// 				"type": "DFR",
	// 				"basic_model": "g",
	// 				"after_effect": "l",
	// 				"normalization": "h2",
	// 				"normalization.h2.c": "3.0"
	// 			}
	// 		},
	// 		"search.slowlog.threshold.query.warn": "10s",
	// 		"search.slowlog.threshold.fetch.trace": "5s",
	// 		"search.slowlog.level": "info",
	// 		"indexing.slowlog.threshold.index.info": "10s",
	// 		"indexing.slowlog.threshold.index.warn": "5s",
	// 		"indexing.slowlog.level": "warn",
	// 		"indexing.slowlog.source": "1000",
	// 		"indexing.slowlog.reformat": false,
	// 		"store.type": "niofs",
	// 		"store.preload": ["nvd", "dvd"],
	// 		"translog.sync_interval": "5s",
	// 		"translog.durability": "request",
	// 		"translog.flush_threshold_size": "512mb",
	// 		"translog.retention.size": "512mb",
	// 		"translog.retention.age": "12h",
	// 		"soft_deletes.enabled": true,
	// 		"soft_deletes.retention_lease.period": "12h",
	// 		"sort.field": "date",
	// 		"sort.order": "desc",
	// 		"sort.mode": "min",
	// 		"sort.missing": "_last",
	// 		"lifecycle.name": "lifecycle_name",
	// 		"lifecycle.rollover_alias": "lifecycle_alias",
	// 		"lifecycle.parse_origination_date": true,
	// 		"lifecycle.origination_date": 1579442569
	// 	}
	// }
	options := make(map[string]interface{})

	if i.numberOfShards != nil {
		options["number_of_shards"] = i.numberOfShards
	}
	if i.shardCheckOnStartup != "" {
		options["shard.check_on_startup"] = i.shardCheckOnStartup
	}
	if i.codec != "" {
		options["codec"] = i.codec
	}
	if i.routingPartitionSize != nil {
		options["routing_partition_size"] = i.routingPartitionSize
	}
	if i.loadFixedBitsetFiltersEagerly != nil {
		options["load_fixed_bitset_filters_eagerly"] = i.loadFixedBitsetFiltersEagerly
	}
	if i.numberOfReplicas != nil {
		options["number_of_replicas"] = i.numberOfReplicas
	}
	if i.autoExpandReplicas != "" {
		options["auto_expand_replicas"] = i.autoExpandReplicas
	}
	if i.searchIdleAfter != "" {
		options["search.idle.after"] = i.searchIdleAfter
	}
	if i.refreshInterval != "" {
		options["refresh_interval"] = i.refreshInterval
	}
	if i.maxResultWindow != nil {
		options["max_result_window"] = i.maxResultWindow
	}
	if i.maxInnerResultWindow != nil {
		options["max_inner_result_window"] = i.maxInnerResultWindow
	}
	if i.maxRescoreWindow != nil {
		options["max_rescore_window"] = i.maxRescoreWindow
	}
	if i.maxDocvalueFieldsSearch != nil {
		options["max_docvalue_fields_search"] = i.maxDocvalueFieldsSearch
	}
	if i.maxScriptFields != nil {
		options["max_script_fields"] = i.maxScriptFields
	}
	if i.maxNGramDiff != nil {
		options["max_ngram_diff"] = i.maxNGramDiff
	}
	if i.maxShingleDiff != nil {
		options["max_shingle_diff"] = i.maxShingleDiff
	}
	if i.blocksReadOnly != nil {
		options["blocks.read_only"] = i.blocksReadOnly
	}
	if i.blocksReadOnlyAllowDelete != nil {
		options["blocks.read_only_allow_delete"] = i.blocksReadOnlyAllowDelete
	}
	if i.blocksRead != nil {
		options["blocks.read"] = i.blocksRead
	}
	if i.blocksWrite != nil {
		options["blocks.write"] = i.blocksWrite
	}
	if i.blocksMetadata != nil {
		options["blocks.metadata"] = i.blocksMetadata
	}
	if i.maxRefreshListeners != nil {
		options["max_refresh_listeners"] = i.maxRefreshListeners
	}
	if i.analyzeMaxTokenCount != nil {
		options["analyze.max_token_count"] = i.analyzeMaxTokenCount
	}
	if i.highlightMaxAnalyzedOffset != nil {
		options["highlight.max_analyzed_offset"] = i.highlightMaxAnalyzedOffset
	}
	if i.maxTermsCount != nil {
		options["max_terms_count"] = i.maxTermsCount
	}
	if i.maxRegexLength != nil {
		options["max_regex_length"] = i.maxRegexLength
	}
	if i.routingAllocationEnable != "" {
		options["routing.allocation.enable"] = i.routingAllocationEnable
	}
	if i.routingRebalanceEnable != "" {
		options["routing.rebalance.enable"] = i.routingRebalanceEnable
	}
	if i.gcDeletes != "" {
		options["gc_deletes"] = i.gcDeletes
	}
	if i.defaultPipeline != "" {
		options["default_pipeline"] = i.defaultPipeline
	}
	if i.finalPipeline != "" {
		options["final_pipeline"] = i.finalPipeline
	}
	if i.analysis != nil {
		analysis, err := i.analysis.Source(false)
		if err != nil {
			return nil, err
		}
		options["analysis"] = analysis
	}
	if len(i.routingAllocation) > 0 {
		for _, a := range i.routingAllocation {
			allocation, err := a.Source(false)
			if err != nil {
				return nil, err
			}
			for k, v := range allocation.(map[string]interface{}) {
				options["routing."+k] = v
			}
		}
	}
	if i.unassignedNodeLeftDelayedTimeout != "" {
		options["unassigned.node_left.delayed_timeout"] = i.unassignedNodeLeftDelayedTimeout
	}
	if i.priority != nil {
		options["priority"] = i.priority
	}
	if i.routingAllocationTotalShardsPerNode != nil {
		options["routing.allocation.total_shards_per_node"] = i.routingAllocationTotalShardsPerNode
	}
	if i.mappings != nil {
		mappings, err := i.mappings.Source(false)
		if err != nil {
			return nil, err
		}
		options["mappings"] = mappings
	}
	if i.mappingTotalFieldsLimit != nil {
		options["mapping.total_fields.limit"] = i.mappingTotalFieldsLimit
	}
	if i.mappingDepthLimit != nil {
		options["mapping.depth.limit"] = i.mappingDepthLimit
	}
	if i.mappingNestedFieldsLimit != nil {
		options["mapping.nested_fields.limit"] = i.mappingNestedFieldsLimit
	}
	if i.mappingNestedObjectsLimit != nil {
		options["mapping.nested_objects.limit"] = i.mappingNestedObjectsLimit
	}
	if i.mappingFieldNameLengthLimit != nil {
		options["mapping.field_name_length.limit"] = i.mappingFieldNameLengthLimit
	}
	if i.mergeSchedulerMaxThreadCount != nil {
		options["merge.scheduler.max_thread_count"] = i.mergeSchedulerMaxThreadCount
	}
	similarity := make(map[string]interface{})
	if i.defaultSimilarity != nil {
		defaultSimilarity, err := i.defaultSimilarity.Source(false)
		if err != nil {
			return nil, err
		}
		similarity["default"] = defaultSimilarity
	}
	if len(i.similarity) > 0 {
		for _, s := range i.similarity {
			_similarity, err := s.Source(false)
			if err != nil {
				return nil, err
			}
			similarity[s.Name()] = _similarity
		}
	}
	if len(similarity) > 0 {
		options["similarity"] = similarity
	}
	if len(i.searchSlowlogThreshold) > 0 {
		for _, searchSlowlogThreshold := range i.searchSlowlogThreshold {
			searchSlowlogThreshold, err := searchSlowlogThreshold.Source(false)
			if err != nil {
				return nil, err
			}
			for k, v := range searchSlowlogThreshold.(map[string]interface{}) {
				options["search."+k] = v
			}
		}
	}
	if i.searchSlowlogLevel != "" {
		options["search.slowlog.level"] = i.searchSlowlogLevel
	}
	if len(i.indexingSlowlogThreshold) > 0 {
		for _, indexingSlowlogThreshold := range i.indexingSlowlogThreshold {
			indexingSlowlogThreshold, err := indexingSlowlogThreshold.Source(false)
			if err != nil {
				return nil, err
			}
			for k, v := range indexingSlowlogThreshold.(map[string]interface{}) {
				options["indexing."+k] = v
			}
		}
	}
	if i.indexingSlowlogLevel != "" {
		options["indexing.slowlog.level"] = i.indexingSlowlogLevel
	}
	if i.indexingSlowlogSource != "" {
		options["indexing.slowlog.source"] = i.indexingSlowlogSource
	}
	if i.indexingSlowlogReformat != nil {
		options["indexing.slowlog.reformat"] = i.indexingSlowlogReformat
	}
	if i.storeType != "" {
		options["store.type"] = i.storeType
	}
	if i.storePreload != nil {
		options["store.preload"] = i.storePreload
	}
	if i.translogSyncInterval != "" {
		options["translog.sync_interval"] = i.translogSyncInterval
	}
	if i.translogDurability != "" {
		options["translog.durability"] = i.translogDurability
	}
	if i.translogFlushThresholdSize != "" {
		options["translog.flush_threshold_size"] = i.translogFlushThresholdSize
	}
	if i.translogRetentionSize != "" {
		options["translog.retention.size"] = i.translogRetentionSize
	}
	if i.translogRetentionAge != "" {
		options["translog.retention.age"] = i.translogRetentionAge
	}
	if i.softDeletesEnabled != nil {
		options["soft_deletes.enabled"] = i.softDeletesEnabled
	}
	if i.softDeletesRetentionLeasePeriod != "" {
		options["soft_deletes.retention_lease.period"] = i.softDeletesRetentionLeasePeriod
	}
	if len(i.sortField) > 0 {
		var sortField interface{}
		switch {
		case len(i.sortField) > 1:
			sortField = i.sortField
			break
		case len(i.sortField) == 1:
			sortField = i.sortField[0]
			break
		default:
			sortField = ""
		}
		options["sort.field"] = sortField
	}
	if i.sortOrder != "" {
		options["sort.order"] = i.sortOrder
	}
	if i.sortMode != "" {
		options["sort.mode"] = i.sortMode
	}
	if i.sortMissing != "" {
		options["sort.missing"] = i.sortMissing
	}
	if i.lifecycleName != "" {
		options["lifecycle.name"] = i.lifecycleName
	}
	if i.lifecycleRolloverAlias != "" {
		options["lifecycle.rollover_alias"] = i.lifecycleRolloverAlias
	}
	if i.lifecycleParseOriginationDate != nil {
		options["lifecycle.parse_origination_date"] = i.lifecycleParseOriginationDate
	}
	if i.lifecycleOriginationDate != nil {
		options["lifecycle.origination_date"] = i.lifecycleOriginationDate
	}

	if !includeName {
		return options, nil
	}

	source := make(map[string]interface{})
	source["index"] = options
	return source, nil
}
