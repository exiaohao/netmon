package task

import (
	"github.com/go-ping/ping"
	"log"
	"time"
)

type ICMPPingTaskParam struct {
	DstAddr			string
	Interval		time.Duration
	Timeout			time.Duration
	Count			int
	DataBytes		int
}

type ICMPPingTaskResult struct {

}

func ExecuteICMPPing(task *ICMPPingTaskParam) (*ping.Statistics, error) {
	log.Print("New pinger")
	pinger, err := ping.NewPinger(task.DstAddr)
	if err != nil {
		return nil, err
	}

	//if task.Count == 0 {
	//	pinger.Count = 5
	//} else {
	//	pinger.Count = task.Count
	//}
	//
	//if task.Interval == 0 {
	//	pinger.Interval = 200 * time.Microsecond
	//} else {
	//	pinger.Interval = task.Interval
	//}
	//
	//if task.Timeout == 0 {
	//	pinger.Timeout = 1 * time.Second
	//} else {
	//	pinger.Timeout = task.Timeout
	//}

	pinger.Count = 5
	pinger.Interval = 100 * time.Microsecond

	log.Print("Run pinger")
	err = pinger.Run()
	if err != nil {
		return nil, err
	}

	stat := pinger.Statistics()
	log.Printf("Show pinger stats: %x", stat)
	return stat, nil
}
