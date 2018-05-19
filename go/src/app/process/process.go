package process

import (
	e "app/events"
	"os"
	"os/exec"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/shirou/gopsutil/process"
)

type MonitoredProcess struct {
	tick int16
	args []string
	send sender
}

type sender func(evt e.ProcessEvent)

func getSender(pEvents chan e.ProcessEvent) sender {
	return func(evt e.ProcessEvent) {
		pEvents <- evt
	}
}

func Init(tick int16, pEvents chan e.ProcessEvent, args []string) MonitoredProcess {
	send := getSender(pEvents)
	return MonitoredProcess{
		tick,
		args,
		send,
	}
}

// Start a new monitored process
func (mp MonitoredProcess) Start(startTrigger chan bool, wg *sync.WaitGroup) {
	<-startTrigger

	defer wg.Done()

	mp.send(e.LogEvent("Starting"))
	mp.send(e.LogEventf("Executing %v\n", mp.args))

	cmd := exec.Command(mp.args[0], mp.args[1:]...)
	cmd.Stdout = os.Stdout
	err := cmd.Start()

	if err != nil {
		mp.send(e.FailedEvent(err))
		return
	}

	pid := cmd.Process.Pid
	mp.send(e.LogEventf("Started subprocess %d", pid))
	mp.send(e.StartedProcessEvent(pid))

	proc, _ := process.NewProcess(int32(pid))

	go mp.forwardTerminatingSignals(cmd.Process)
	go mp.cpuMon(proc)
	go mp.memMon(proc)
	go mp.netMon(proc)

	mp.send(e.LogEventf("Waiting for subprocess %d", pid))
	err = cmd.Wait()

	if err == nil {
		output, _ := cmd.Output()
		mp.send(e.SuccessEvent(output))
	} else {
		mp.send(e.LogEventf("Error %v", err))
		mp.send(e.FailedEvent(err))
	}
}

func (mp MonitoredProcess) forwardTerminatingSignals(p *os.Process) {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	sig := <-sigs
	close(sigs)

	mp.send(e.LogEventf("Forwarding signal %v", sig))

	err := p.Signal(sig)
	if err != nil {
		mp.send(e.LogEventf("SIG %v", []interface{}{err.Error()}))
	}
}

func (mp MonitoredProcess) cpuMon(p *process.Process) {
	loopFor(time.Second*time.Duration(mp.tick), func() {
		cpu, _ := p.CPUPercent()
		mp.send(e.CPUEEvent(cpu))
	})
}

func (mp MonitoredProcess) memMon(p *process.Process) {
	loopFor(time.Second*time.Duration(mp.tick), func() {
		mem, _ := p.MemoryInfo()
		mp.send(e.MemoryEvent(mem))
	})
}

func (mp MonitoredProcess) netMon(p *process.Process) {
	loopFor(time.Second*time.Duration(mp.tick), func() {
		net, _ := p.NetIOCounters(false)
		mp.send(e.NetworkEvent(net[0]))
	})
}

func loopFor(t time.Duration, f func()) {
	func() {
		for {
			<-time.After(t)
			f()
		}
	}()
}
