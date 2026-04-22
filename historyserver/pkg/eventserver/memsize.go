package eventserver

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
)

const mapEntryOverhead = 11

func DeepSize(v any) uint64 {
	s := &sizer{visited: make(map[uintptr]struct{})}
	return uint64(s.sizeOf(reflect.ValueOf(v)))
}

type sizer struct {
	visited map[uintptr]struct{}
}

func (s *sizer) seen(p uintptr) bool {
	if p == 0 {
		return false
	}
	if _, ok := s.visited[p]; ok {
		return true
	}
	s.visited[p] = struct{}{}
	return false
}

// sizeOf returns bytes for v itself plus everything it indirectly references.
func (s *sizer) sizeOf(v reflect.Value) uintptr {
	if !v.IsValid() {
		return 0
	}
	switch v.Kind() {
	case reflect.Pointer:
		if v.IsNil() || s.seen(v.Pointer()) {
			return v.Type().Size()
		}
		return v.Type().Size() + s.sizeOf(v.Elem())

	case reflect.Interface:
		if v.IsNil() {
			return v.Type().Size()
		}
		return v.Type().Size() + s.sizeOf(v.Elem())

	case reflect.String:
		return v.Type().Size() + uintptr(v.Len())

	case reflect.Slice:
		if v.IsNil() || s.seen(v.Pointer()) {
			return v.Type().Size()
		}
		elem := v.Type().Elem()
		total := v.Type().Size() + uintptr(v.Cap())*elem.Size()
		for i := 0; i < v.Len(); i++ {
			total += s.sizeOf(v.Index(i)) - elem.Size()
		}
		return total

	case reflect.Array:
		elem := v.Type().Elem()
		total := v.Type().Size()
		for i := 0; i < v.Len(); i++ {
			total += s.sizeOf(v.Index(i)) - elem.Size()
		}
		return total

	case reflect.Map:
		if v.IsNil() {
			return v.Type().Size()
		}
		if ptr := v.UnsafePointer(); ptr != nil && s.seen(uintptr(ptr)) {
			return v.Type().Size()
		}
		kt := v.Type().Key()
		et := v.Type().Elem()
		total := v.Type().Size() + uintptr(v.Len())*(kt.Size()+et.Size()+mapEntryOverhead)
		iter := v.MapRange()
		for iter.Next() {
			total += s.sizeOf(iter.Key()) - kt.Size()
			total += s.sizeOf(iter.Value()) - et.Size()
		}
		return total

	case reflect.Struct:
		total := v.Type().Size()
		for i := 0; i < v.NumField(); i++ {
			ft := v.Field(i).Type()
			total += s.sizeOf(v.Field(i)) - ft.Size()
		}
		return total

	default:
		return v.Type().Size()
	}
}

// MapSizeReport summarizes per-map deep sizes inside an EventHandler.
type MapSizeReport struct {
	TaskBytes     uint64
	ActorBytes    uint64
	JobBytes      uint64
	NodeBytes     uint64
	LogEventBytes uint64
	TotalBytes    uint64
}

// MeasureMapSizes walks all five Cluster*Map fields and returns deep size estimates.
func (h *EventHandler) MeasureMapSizes() MapSizeReport {
	var r MapSizeReport

	h.ClusterTaskMap.RLock()
	r.TaskBytes = DeepSize(h.ClusterTaskMap)
	h.ClusterTaskMap.RUnlock()

	h.ClusterActorMap.RLock()
	r.ActorBytes = DeepSize(h.ClusterActorMap)
	h.ClusterActorMap.RUnlock()

	h.ClusterJobMap.RLock()
	r.JobBytes = DeepSize(h.ClusterJobMap)
	h.ClusterJobMap.RUnlock()

	h.ClusterNodeMap.RLock()
	r.NodeBytes = DeepSize(h.ClusterNodeMap)
	h.ClusterNodeMap.RUnlock()

	if h.ClusterLogEventMap != nil {
		r.LogEventBytes = DeepSize(h.ClusterLogEventMap)
	}

	r.TotalBytes = r.TaskBytes + r.ActorBytes + r.JobBytes + r.NodeBytes + r.LogEventBytes
	return r
}

// LogMapSizes measures and logs per-map sizes, total in-memory bytes, and ratio of in-memory bytes to raw bytes.
func (h *EventHandler) LogMapSizes(rawBytes int64) {
	r := h.MeasureMapSizes()

	ratio := ""
	if rawBytes > 0 {
		ratio = fmt.Sprintf(" ratio=%.2fx", float64(r.TotalBytes)/float64(rawBytes))
	}

	logrus.Infof(
		"[memsize] raw=%s in_mem=%s%s | task=%s actor=%s job=%s node=%s log_event=%s",
		humanBytes(uint64(rawBytes)), humanBytes(r.TotalBytes), ratio,
		humanBytes(r.TaskBytes),
		humanBytes(r.ActorBytes),
		humanBytes(r.JobBytes),
		humanBytes(r.NodeBytes),
		humanBytes(r.LogEventBytes),
	)
}

func humanBytes(b uint64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%dB", b)
	}
	div, exp := uint64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.2f%ciB", float64(b)/float64(div), "KMGTPE"[exp])
}
