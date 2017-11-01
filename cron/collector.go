package cron

import (
	"github.com/shirou/gopsutil/process"
	"fmt"
)

type CPUInfo struct {
	pid int32
	CPUPercent float64
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
		return
	}
	cpuInfoList,err := collectCPU(pids)
	if err != nil {
		//error handle
		fmt.Println("error:2")
		return
	}

	for i:=0;i<len(cpuInfoList);i++{
		fmt.Printf("%s:%d", cpuInfoList[i].pid, cpuInfoList[i].CPUPercent)
	}

}

func collectCPU(pids []int32) (CPUInfoList []*CPUInfo,err error) {
	for _, pid := range pids{
		proc, err := process.NewProcess(pid)
		if err != nil {
			fmt.Println(err.Error())
			return CPUInfoList,err
		}
		var singleInfo *CPUInfo

		CPUPercent,err := proc.CPUPercent()
		if err !=nil {
			fmt.Print(proc.Cmdline())
			fmt.Println(err.Error())
			continue
			//return CPUInfoList,err
		}

		singleInfo = &CPUInfo{
			pid:pid,
			CPUPercent:CPUPercent,

		}
		CPUInfoList = append(CPUInfoList, singleInfo)


	}


	return CPUInfoList,nil
}





