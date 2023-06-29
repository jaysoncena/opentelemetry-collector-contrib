// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/pdata/pcommon"
	"go.opentelemetry.io/collector/pdata/pmetric"
	"go.opentelemetry.io/collector/receiver/receivertest"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

type testConfigCollection int

const (
	testSetDefault testConfigCollection = iota
	testSetAll
	testSetNone
)

func TestMetricsBuilder(t *testing.T) {
	tests := []struct {
		name      string
		configSet testConfigCollection
	}{
		{
			name:      "default",
			configSet: testSetDefault,
		},
		{
			name:      "all_set",
			configSet: testSetAll,
		},
		{
			name:      "none_set",
			configSet: testSetNone,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			start := pcommon.Timestamp(1_000_000_000)
			ts := pcommon.Timestamp(1_000_001_000)
			observedZapCore, observedLogs := observer.New(zap.WarnLevel)
			settings := receivertest.NewNopCreateSettings()
			settings.Logger = zap.New(observedZapCore)
			mb := NewMetricsBuilder(loadMetricsBuilderConfig(t, test.name), settings, WithStartTime(start))

			expectedWarnings := 0
			assert.Equal(t, expectedWarnings, observedLogs.Len())

			defaultMetricsCount := 0
			allMetricsCount := 0

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sNodeAllocatableCPUDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sNodeAllocatableEphemeralStorageDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sNodeAllocatableMemoryDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sNodeAllocatablePodsDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sNodeAllocatableStorageDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sNodeConditionDiskPressureDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sNodeConditionMemoryPressureDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sNodeConditionNetworkUnavailableDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sNodeConditionPidPressureDataPoint(ts, 1)

			defaultMetricsCount++
			allMetricsCount++
			mb.RecordK8sNodeConditionReadyDataPoint(ts, 1)

			metrics := mb.Emit(WithK8sNodeName("attr-val"), WithK8sNodeUID("attr-val"), WithOpencensusResourcetype("attr-val"))

			if test.configSet == testSetNone {
				assert.Equal(t, 0, metrics.ResourceMetrics().Len())
				return
			}

			assert.Equal(t, 1, metrics.ResourceMetrics().Len())
			rm := metrics.ResourceMetrics().At(0)
			attrCount := 0
			enabledAttrCount := 0
			attrVal, ok := rm.Resource().Attributes().Get("k8s.node.name")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.K8sNodeName.Enabled, ok)
			if mb.resourceAttributesConfig.K8sNodeName.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("k8s.node.uid")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.K8sNodeUID.Enabled, ok)
			if mb.resourceAttributesConfig.K8sNodeUID.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			attrVal, ok = rm.Resource().Attributes().Get("opencensus.resourcetype")
			attrCount++
			assert.Equal(t, mb.resourceAttributesConfig.OpencensusResourcetype.Enabled, ok)
			if mb.resourceAttributesConfig.OpencensusResourcetype.Enabled {
				enabledAttrCount++
				assert.EqualValues(t, "attr-val", attrVal.Str())
			}
			assert.Equal(t, enabledAttrCount, rm.Resource().Attributes().Len())
			assert.Equal(t, attrCount, 3)

			assert.Equal(t, 1, rm.ScopeMetrics().Len())
			ms := rm.ScopeMetrics().At(0).Metrics()
			if test.configSet == testSetDefault {
				assert.Equal(t, defaultMetricsCount, ms.Len())
			}
			if test.configSet == testSetAll {
				assert.Equal(t, allMetricsCount, ms.Len())
			}
			validatedMetrics := make(map[string]bool)
			for i := 0; i < ms.Len(); i++ {
				switch ms.At(i).Name() {
				case "k8s.node.allocatable_cpu":
					assert.False(t, validatedMetrics["k8s.node.allocatable_cpu"], "Found a duplicate in the metrics slice: k8s.node.allocatable_cpu")
					validatedMetrics["k8s.node.allocatable_cpu"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "How many CPU cores remaining that the node can allocate to pods", ms.At(i).Description())
					assert.Equal(t, "{cores}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeDouble, dp.ValueType())
					assert.Equal(t, float64(1), dp.DoubleValue())
				case "k8s.node.allocatable_ephemeral_storage":
					assert.False(t, validatedMetrics["k8s.node.allocatable_ephemeral_storage"], "Found a duplicate in the metrics slice: k8s.node.allocatable_ephemeral_storage")
					validatedMetrics["k8s.node.allocatable_ephemeral_storage"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "How many bytes of ephemeral storage remaining that the node can allocate to pods", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.node.allocatable_memory":
					assert.False(t, validatedMetrics["k8s.node.allocatable_memory"], "Found a duplicate in the metrics slice: k8s.node.allocatable_memory")
					validatedMetrics["k8s.node.allocatable_memory"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "How many bytes of RAM memory remaining that the node can allocate to pods", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.node.allocatable_pods":
					assert.False(t, validatedMetrics["k8s.node.allocatable_pods"], "Found a duplicate in the metrics slice: k8s.node.allocatable_pods")
					validatedMetrics["k8s.node.allocatable_pods"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "How many pods remaining the node can allocate", ms.At(i).Description())
					assert.Equal(t, "{pods}", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.node.allocatable_storage":
					assert.False(t, validatedMetrics["k8s.node.allocatable_storage"], "Found a duplicate in the metrics slice: k8s.node.allocatable_storage")
					validatedMetrics["k8s.node.allocatable_storage"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "How many bytes of storage remaining that the node can allocate to pods", ms.At(i).Description())
					assert.Equal(t, "By", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.node.condition_disk_pressure":
					assert.False(t, validatedMetrics["k8s.node.condition_disk_pressure"], "Found a duplicate in the metrics slice: k8s.node.condition_disk_pressure")
					validatedMetrics["k8s.node.condition_disk_pressure"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Whether this node is DiskPressure (1), not DiskPressure (0) or in an unknown state (-1)", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.node.condition_memory_pressure":
					assert.False(t, validatedMetrics["k8s.node.condition_memory_pressure"], "Found a duplicate in the metrics slice: k8s.node.condition_memory_pressure")
					validatedMetrics["k8s.node.condition_memory_pressure"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Whether this node is MemoryPressure (1), not MemoryPressure (0) or in an unknown state (-1)", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.node.condition_network_unavailable":
					assert.False(t, validatedMetrics["k8s.node.condition_network_unavailable"], "Found a duplicate in the metrics slice: k8s.node.condition_network_unavailable")
					validatedMetrics["k8s.node.condition_network_unavailable"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Whether this node is NetworkUnavailable (1), not NetworkUnavailable (0) or in an unknown state (-1)", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.node.condition_pid_pressure":
					assert.False(t, validatedMetrics["k8s.node.condition_pid_pressure"], "Found a duplicate in the metrics slice: k8s.node.condition_pid_pressure")
					validatedMetrics["k8s.node.condition_pid_pressure"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Whether this node is PidPressure (1), not PidPressure (0) or in an unknown state (-1)", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				case "k8s.node.condition_ready":
					assert.False(t, validatedMetrics["k8s.node.condition_ready"], "Found a duplicate in the metrics slice: k8s.node.condition_ready")
					validatedMetrics["k8s.node.condition_ready"] = true
					assert.Equal(t, pmetric.MetricTypeGauge, ms.At(i).Type())
					assert.Equal(t, 1, ms.At(i).Gauge().DataPoints().Len())
					assert.Equal(t, "Whether this node is Ready (1), not Ready (0) or in an unknown state (-1)", ms.At(i).Description())
					assert.Equal(t, "1", ms.At(i).Unit())
					dp := ms.At(i).Gauge().DataPoints().At(0)
					assert.Equal(t, start, dp.StartTimestamp())
					assert.Equal(t, ts, dp.Timestamp())
					assert.Equal(t, pmetric.NumberDataPointValueTypeInt, dp.ValueType())
					assert.Equal(t, int64(1), dp.IntValue())
				}
			}
		})
	}
}
