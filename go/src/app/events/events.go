package events

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

// ProcessEvent is the event dispatched to notify
// something happened to the monitored process
type ProcessEvent struct {
	Name    string
	Data    interface{}
	Message string
}

const MEMORY = "MEM"
const NETWORK = "NET"
const CPU = "CPU"
const UPTIME = "UPTIME"
const STARTED = "STARTED"
const DONE = "DONE"
const FAILED = "FAILED"
const LOG = "LOG"

func CPUEEvent(percent float64) ProcessEvent {
	cpuAsString := fmt.Sprint(percent)
	return ProcessEvent{Name: CPU, Data: cpuAsString, Message: "CPU %"}
}

func MemoryEvent(memoryStats *process.MemoryInfoStat) ProcessEvent {
	memAsString, _ := json.Marshal(memoryStats)
	return ProcessEvent{Name: MEMORY, Data: string(memAsString), Message: "Memory stats"}
}

func UptimeEvent(uptime time.Duration) ProcessEvent {
	uptimeSeconds, _ := json.Marshal(uptime.Seconds())
	return ProcessEvent{Name: UPTIME, Data: string(uptimeSeconds), Message: "Uptime"}
}

func NetworkEvent(netStats net.IOCountersStat) ProcessEvent {
	netAsString, _ := json.Marshal(netStats)
	return ProcessEvent{Name: NETWORK, Data: string(netAsString), Message: "Network stats"}
}

func StartedProcessEvent(pid int) ProcessEvent {
	return ProcessEvent{Name: STARTED, Data: pid}
}

func SuccessEvent(output []byte) ProcessEvent {
	return ProcessEvent{Name: DONE, Data: "Success", Message: string(output)}
}

func FailedEvent(err error) ProcessEvent {
	return ProcessEvent{Name: FAILED, Data: "Failed", Message: err.Error()}
}

func LogEvent(message string) ProcessEvent {
	return ProcessEvent{Name: LOG, Data: "", Message: message}
}

func LogEventf(message string, args interface{}) ProcessEvent {
	msg := fmt.Sprintf(message, args)
	return ProcessEvent{Name: LOG, Data: "", Message: msg}
}

func LogEventData(message string, extraData interface{}) ProcessEvent {
	data, _ := json.Marshal(extraData)
	return ProcessEvent{Name: LOG, Data: data, Message: message}
}
