package proc

import (
	nproc "github.com/niean/gotools/proc"
	"log"
)

// 监控
var (
	MonitorCronCnt            = nproc.NewSCounterQps("MonitorCronCnt")
	MonitorConcurrentErrorCnt = nproc.NewSCounterQps("MonitorConcurrentErrorCnt")
	MonitorAlarmMailCnt       = nproc.NewSCounterQps("MonitorAlarmMailCnt")
	MonitorAlarmSmsCnt        = nproc.NewSCounterQps("MonitorAlarmSmsCnt")
	MonitorAlarmCallbackCnt   = nproc.NewSCounterQps("MonitorAlarmCallbackCnt")
)

func Start() {
	log.Println("proc:Start, ok")
}

func GetAll() []interface{} {
	ret := make([]interface{}, 0)

	// monitor
	ret = append(ret, MonitorCronCnt.Get())
	ret = append(ret, MonitorConcurrentErrorCnt.Get())
	ret = append(ret, MonitorAlarmMailCnt.Get())
	ret = append(ret, MonitorAlarmSmsCnt.Get())
	ret = append(ret, MonitorAlarmCallbackCnt.Get())

	return ret
}
