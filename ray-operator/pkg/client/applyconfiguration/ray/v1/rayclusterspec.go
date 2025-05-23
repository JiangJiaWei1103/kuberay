// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

// RayClusterSpecApplyConfiguration represents a declarative configuration of the RayClusterSpec type for use
// with apply.
type RayClusterSpecApplyConfiguration struct {
	Suspend                  *bool                                       `json:"suspend,omitempty"`
	ManagedBy                *string                                     `json:"managedBy,omitempty"`
	AutoscalerOptions        *AutoscalerOptionsApplyConfiguration        `json:"autoscalerOptions,omitempty"`
	HeadServiceAnnotations   map[string]string                           `json:"headServiceAnnotations,omitempty"`
	EnableInTreeAutoscaling  *bool                                       `json:"enableInTreeAutoscaling,omitempty"`
	GcsFaultToleranceOptions *GcsFaultToleranceOptionsApplyConfiguration `json:"gcsFaultToleranceOptions,omitempty"`
	HeadGroupSpec            *HeadGroupSpecApplyConfiguration            `json:"headGroupSpec,omitempty"`
	RayVersion               *string                                     `json:"rayVersion,omitempty"`
	WorkerGroupSpecs         []WorkerGroupSpecApplyConfiguration         `json:"workerGroupSpecs,omitempty"`
}

// RayClusterSpecApplyConfiguration constructs a declarative configuration of the RayClusterSpec type for use with
// apply.
func RayClusterSpec() *RayClusterSpecApplyConfiguration {
	return &RayClusterSpecApplyConfiguration{}
}

// WithSuspend sets the Suspend field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Suspend field is set to the value of the last call.
func (b *RayClusterSpecApplyConfiguration) WithSuspend(value bool) *RayClusterSpecApplyConfiguration {
	b.Suspend = &value
	return b
}

// WithManagedBy sets the ManagedBy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ManagedBy field is set to the value of the last call.
func (b *RayClusterSpecApplyConfiguration) WithManagedBy(value string) *RayClusterSpecApplyConfiguration {
	b.ManagedBy = &value
	return b
}

// WithAutoscalerOptions sets the AutoscalerOptions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the AutoscalerOptions field is set to the value of the last call.
func (b *RayClusterSpecApplyConfiguration) WithAutoscalerOptions(value *AutoscalerOptionsApplyConfiguration) *RayClusterSpecApplyConfiguration {
	b.AutoscalerOptions = value
	return b
}

// WithHeadServiceAnnotations puts the entries into the HeadServiceAnnotations field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, the entries provided by each call will be put on the HeadServiceAnnotations field,
// overwriting an existing map entries in HeadServiceAnnotations field with the same key.
func (b *RayClusterSpecApplyConfiguration) WithHeadServiceAnnotations(entries map[string]string) *RayClusterSpecApplyConfiguration {
	if b.HeadServiceAnnotations == nil && len(entries) > 0 {
		b.HeadServiceAnnotations = make(map[string]string, len(entries))
	}
	for k, v := range entries {
		b.HeadServiceAnnotations[k] = v
	}
	return b
}

// WithEnableInTreeAutoscaling sets the EnableInTreeAutoscaling field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the EnableInTreeAutoscaling field is set to the value of the last call.
func (b *RayClusterSpecApplyConfiguration) WithEnableInTreeAutoscaling(value bool) *RayClusterSpecApplyConfiguration {
	b.EnableInTreeAutoscaling = &value
	return b
}

// WithGcsFaultToleranceOptions sets the GcsFaultToleranceOptions field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the GcsFaultToleranceOptions field is set to the value of the last call.
func (b *RayClusterSpecApplyConfiguration) WithGcsFaultToleranceOptions(value *GcsFaultToleranceOptionsApplyConfiguration) *RayClusterSpecApplyConfiguration {
	b.GcsFaultToleranceOptions = value
	return b
}

// WithHeadGroupSpec sets the HeadGroupSpec field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the HeadGroupSpec field is set to the value of the last call.
func (b *RayClusterSpecApplyConfiguration) WithHeadGroupSpec(value *HeadGroupSpecApplyConfiguration) *RayClusterSpecApplyConfiguration {
	b.HeadGroupSpec = value
	return b
}

// WithRayVersion sets the RayVersion field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the RayVersion field is set to the value of the last call.
func (b *RayClusterSpecApplyConfiguration) WithRayVersion(value string) *RayClusterSpecApplyConfiguration {
	b.RayVersion = &value
	return b
}

// WithWorkerGroupSpecs adds the given value to the WorkerGroupSpecs field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the WorkerGroupSpecs field.
func (b *RayClusterSpecApplyConfiguration) WithWorkerGroupSpecs(values ...*WorkerGroupSpecApplyConfiguration) *RayClusterSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithWorkerGroupSpecs")
		}
		b.WorkerGroupSpecs = append(b.WorkerGroupSpecs, *values[i])
	}
	return b
}
