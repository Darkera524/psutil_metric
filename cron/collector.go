package cron

import (
	"github.com/shirou/gopsutil/process"
	"fmt"
	//"github.com/open-falcon/common/model"
)

type ProcessInfo struct {
	pid int32
	cmdline string
	//excutablePath string
	//workingDirctory string
	CPUPercent float64
	MemPercent float32
	FileDescriptorNum int32
	ThreadNum int32
}

type MemInfo struct {

}

func Test(){
	fmt.Println("qweqe")
}

func Collect() {
	pids, err := process.Pids()


	if err != nil {
		//error handle
		fmt.Println("error:1")
		fmt.Println(err.Error())
		return
	}
	cpuInfoList,err := collectCPU(pids)
	if err != nil {
		//error handle
		fmt.Println("error:2")
		fmt.Println(err.Error())
		return
	}

	for i:=0;i<len(cpuInfoList);i++{
		//fmt.Print("pid:",cpuInfoList[i].pid,"cmdline:",cpuInfoList[i].cmdline ,"cpu:",cpuInfoList[i].CPUPercent,"mem:",cpuInfoList[i].MemPercent,"fdn:",cpuInfoList[i].FileDescriptorNum,"thread:",cpuInfoList[i].ThreadNum,"\n")
		fmt.Sprintf("pid=%d,cmdline=%s", cpuInfoList[i].pid, cpuInfoList[i].cmdline)
		}

}

func collectCPU(pids []int32) (CPUInfoList []*ProcessInfo,err error) {
	for _, pid := range pids{
		proc, err := process.NewProcess(pid)
		if err != nil {
			fmt.Println(err.Error())
			return CPUInfoList,err
		}
		var singleInfo *ProcessInfo

		CPUPercent,err := proc.CPUPercent()
		if err !=nil {
			return CPUInfoList,err
		}

		cmdline, err := proc.Cmdline()
		if err !=nil {
			return CPUInfoList,err
		}

		/*excutablePath, err := proc.Exe()
		if err !=nil {
			return CPUInfoList,err
		}*/

		/*workingDerectory, err := proc.Cwd()
		if err !=nil {
			return CPUInfoList,err
		}*/

		memPercent, err := proc.MemoryPercent()
		if err !=nil {
			return CPUInfoList,err
		}

		fileDescriptiorNum, err := proc.NumFDs()
		if err !=nil {
			return CPUInfoList,err
		}

		threadNum, err := proc.NumThreads()
		if err !=nil {
			return CPUInfoList,err
		}

		singleInfo = &ProcessInfo{
			pid:pid,
			cmdline:cmdline,
			//excutablePath:excutablePath,
			//workingDirctory:workingDerectory,
			CPUPercent:CPUPercent,
			MemPercent:memPercent,
			FileDescriptorNum:fileDescriptiorNum,
			ThreadNum:threadNum,
		}
		CPUInfoList = append(CPUInfoList, singleInfo)


	}


	return CPUInfoList,nil
}

/*func convirtProcessInfoToMetrics(procInfo []*ProcessInfo)(metrics []*model.MetricValue, err error){
	for i:=0;i<len(procInfo);i++{
		singleMetric := &model.MetricValue{
			Endpoint:  hostname,
			Metric:    "lvs.rip.active_conns",
			Value:     rip.ActiveConns,
			Timestamp: now,
			Step:      interval,
			Type:      "GAUGE",
			Tags:      tag,
		}
	}
}
*/




