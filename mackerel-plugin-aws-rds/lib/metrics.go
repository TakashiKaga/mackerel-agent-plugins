package mpawsrds

import (
	mp "github.com/mackerelio/go-mackerel-plugin"
)

func (p RDSPlugin) rdsMetrics() (metrics []string) {
	for _, v := range p.GraphDefinition() {
		for _, vv := range v.Metrics {
			metrics = append(metrics, vv.Name)
		}
	}
	return
}

func mergeGraphDefs(a, b map[string]mp.Graphs) map[string]mp.Graphs {
	for k, v := range b {
		a[k] = v
	}
	return a
}

func (p RDSPlugin) mySQLGraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		p.Prefix + ".BinLogDiskUsage": {
			Label: p.LabelPrefix + " BinLogDiskUsage",
			Unit:  "bytes",
			Metrics: []mp.Metrics{
				{Name: "BinLogDiskUsage", Label: "Usage"},
			},
		},
	}
}

func (p RDSPlugin) postgreSQLGraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		p.Prefix + ".MaximumUsedTransactionIDs": {
			Label: p.LabelPrefix + " Maximum Used Transaction IDs",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "MaximumUsedTransactionIDs", Label: "IDs"},
			},
		},
		p.Prefix + ".OldestReplicationSlotLag": {
			Label: p.LabelPrefix + " Oldest Replication Slot Lag",
			Unit:  "bytes",
			Metrics: []mp.Metrics{
				{Name: "OldestReplicationSlotLag", Label: "SlotLag"},
			},
		},
		p.Prefix + ".ReplicationSlotDiskUsage": {
			Label: p.LabelPrefix + " Replication Slot Disk Usage",
			Unit:  "bytes",
			Metrics: []mp.Metrics{
				{Name: "ReplicationSlotDiskUsage", Label: "SlotDiskUsage"},
			},
		},
		p.Prefix + ".TransactionLogsDiskUsage": {
			Label: p.LabelPrefix + " Transaction Logs Disk Usage",
			Unit:  "bytes",
			Metrics: []mp.Metrics{
				{Name: "TransactionLogsDiskUsage", Label: "DiskUsage"},
			},
		},
		p.Prefix + ".TransactionLogsGeneration": {
			Label: p.LabelPrefix + " Transaction Logs Generation",
			Unit:  "bytes/sec",
			Metrics: []mp.Metrics{
				{Name: "TransactionLogsGeneration", Label: "Generation"},
			},
		},
	}
}

func (p RDSPlugin) auroraGraphDefinition() map[string]mp.Graphs {
	return map[string]mp.Graphs{
		p.Prefix + ".Deadlocks": {
			Label: p.LabelPrefix + " Dead Locks",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "Deadlocks", Label: "Deadlocks"},
			},
		},
		p.Prefix + ".Transaction": {
			Label: p.LabelPrefix + " Transaction",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "ActiveTransactions", Label: "Active"},
				{Name: "BlockedTransactions", Label: "Blocked"},
			},
		},
		p.Prefix + ".EngineUptime": {
			Label: p.LabelPrefix + " Engine Uptime",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "EngineUptime", Label: "EngineUptime"},
			},
		},
		p.Prefix + ".Latency": {
			Label: p.LabelPrefix + " Latency in second",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "SelectThroughput", Label: "Select"},
				{Name: "InsertLatency", Label: "Insert"},
				{Name: "UpdateLatency", Label: "Update"},
				{Name: "DeleteLatency", Label: "Delete"},
				{Name: "CommitLatency", Label: "Commit"},
				{Name: "DDLLatency", Label: "DDL"},
				{Name: "DMLLatency", Label: "DML"},
			},
		},
		p.Prefix + ".Throughput": {
			Label: p.LabelPrefix + " Throughput",
			Unit:  "bytes/sec",
			Metrics: []mp.Metrics{
				{Name: "SelectThroughput", Label: "Select"},
				{Name: "InsertThroughput", Label: "Insert"},
				{Name: "UpdateThroughput", Label: "Update"},
				{Name: "DeleteThroughput", Label: "Delete"},
				{Name: "CommitThroughput", Label: "Commit"},
				{Name: "DDLThroughput", Label: "DDL"},
				{Name: "DMLThroughput", Label: "DML"},
			},
		},
		p.Prefix + ".Queries": {
			Label: p.LabelPrefix + " Queries",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "Queries", Label: "Queries"},
			},
		},
		p.Prefix + ".LoginFailures": {
			Label: p.LabelPrefix + " Login Failures",
			Unit:  "integer",
			Metrics: []mp.Metrics{
				{Name: "LoginFailures", Label: "LoginFailures"},
			},
		},
		p.Prefix + ".CacheHitRatio": {
			Label: p.LabelPrefix + " Cache Hit Ratio",
			Unit:  "percentage",
			Metrics: []mp.Metrics{
				{Name: "ResultSetCacheHitRatio", Label: "ResultSet"},
				{Name: "BufferCacheHitRatio", Label: "Buffer"},
			},
		},
		p.Prefix + ".AuroraBinlogReplicaLag": {
			Label: p.LabelPrefix + " Aurora Binlog ReplicaLag",
			Unit:  "integer",
			Metrics: []mp.Metrics{
				{Name: "AuroraBinlogReplicaLag", Label: "BinlogReplicaLag"},
			},
		},
		p.Prefix + ".AuroraReplicaLag": {
			Label: p.LabelPrefix + " Aurora ReplicaLag",
			Unit:  "float",
			Metrics: []mp.Metrics{
				{Name: "AuroraReplicaLagMaximum", Label: "Maximum"},
				{Name: "AuroraReplicaLagMinimum", Label: "Minimum"},
			},
		},
	}
}
