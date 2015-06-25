// Copyright 2012-2015 Apcera Inc. All rights reserved.

/*
This tool fetches the CPU, memory and some other information of the system that
it is running on and displays them to the user.
*/
package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

// This function runs a HTTP server that when visited returns the system information
// of the machine it is running on.
func main() {
	http.HandleFunc("/", SystemInfo)
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// SystemInfo is used to retrieve the CPU, memory, cache and other system information
// of the compute machine that it is running on.
func SystemInfo(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<!DOCTYPE html>")
	io.WriteString(w, "<html>")
	io.WriteString(w, "<body>")

	io.WriteString(w, "<div>")
	io.WriteString(w, `<h2><span style="color:blue">System Information</span></h2>`)
	io.WriteString(w, "<h3>CPU Information</h3>")
	cpuMap := CpuInfo()
	io.WriteString(w, `<ul style="list-style-type:none">`)
	io.WriteString(w, fmt.Sprintf("<li>Processor Model Name = %s</li><li></li>", cpuMap["processorModel"]))
	io.WriteString(w, fmt.Sprintf("<li>Total number of logical cores in the system = %s </li>", cpuMap["logicalCores"]))
	io.WriteString(w, fmt.Sprintf("<li>Total number of physical cores in the system = %s</li>", cpuMap["physicalCores"]))
	io.WriteString(w, "<li>")
	io.WriteString(w, `<span style="color:blue">Hyperthreading</span>: Please note that, if the number of logical cores is twice the number of `+
		`physical cores, then the cores are potentially hyperthreaded.`)
	io.WriteString(w, "</li>")
	io.WriteString(w, "</ul>")
	io.WriteString(w, "</div>")

	io.WriteString(w, "<div>")
	memoryMap := MemoryInfo()
	io.WriteString(w, "<h3>Memory Information</h3>")
	io.WriteString(w, `<ul style="list-style-type:none">`)
	io.WriteString(w, "<li>"+memoryMap["memoryTotal"]+"</li>")
	io.WriteString(w, "<li>"+memoryMap["memoryFree"]+"</li>")
	io.WriteString(w, "<li>"+memoryMap["buffers"]+"</li>")
	io.WriteString(w, "<li>"+memoryMap["cache"]+"</li>")
	io.WriteString(w, "</ul>")

	io.WriteString(w, "<h3>Cache Information</h3>")
	io.WriteString(w, `<ul style="list-style-type:none">`)
	io.WriteString(w, "<li>"+memoryMap["cacheInfo"]+"</li>")
	io.WriteString(w, "<li>"+memoryMap["virtualizationInfo"]+"</li>")
	io.WriteString(w, "<li>"+memoryMap["hypervisorInfo"]+"</li>")
	io.WriteString(w, "</ul>")
	io.WriteString(w, "</div>")

	io.WriteString(w, "<div>")
	io.WriteString(w, "<h3>Other System Information</h3>")
	io.WriteString(w, `<ul style="list-style-type:none">`)
	otherInfoMap := OtherSystemInfo()
	io.WriteString(w, fmt.Sprintf("<li>Hostname = %s</li>", otherInfoMap["hostname"]))
	io.WriteString(w, fmt.Sprintf("<li>IPv4 Address = %s</li>", otherInfoMap["ipv4_address"]))
	io.WriteString(w, fmt.Sprintf("<li>System's memory page size = %s</li>", otherInfoMap["os_pagesize"]))
	io.WriteString(w, fmt.Sprintf("<li>Target Architecture = %s</li>", otherInfoMap["target_architecture"]))
	io.WriteString(w, fmt.Sprintf("<li>Target Operating System = %s</li>", otherInfoMap["target_os"]))
	io.WriteString(w, "</ul>")
	io.WriteString(w, "</div>")

	io.WriteString(w, "</body>")
	io.WriteString(w, "</html>")
}

// CpuInfo retrieves the CPU information of the system like number of physical and logical cores
// in the system and processor model information.
func CpuInfo() map[string]string {
	cpuMap := make(map[string]string)

	// Read the CPU information
	// Total number of logical cores in the system
	numLogicalCores := strconv.Itoa(runtime.NumCPU())
	if numLogicalCores != "" {
		cpuMap["logicalCores"] = numLogicalCores
	} else {
		cpuMap["logicalCores"] = "Error: The number of logical cores for the current system could not be retrieved."
	}

	// Total number of physical cores in the system
	cmdNumProcessors := exec.Command("/bin/sh", "-c", "cat /proc/cpuinfo | grep processor | wc -l")
	outputNumProcessors, err := cmdNumProcessors.Output()
	if err != nil {
		cpuMap["physicalCores"] = "Error: The number of physical cores for the current system could not be retrieved."
	} else {
		cpuMap["physicalCores"] = string(outputNumProcessors)
	}

	// Processor Model Name
	cmdProcessorModel := exec.Command("/bin/sh", "-c", "cat /proc/cpuinfo | grep 'model name' | uniq | awk '{print substr($0, index($0,$4))}'")
	outputProcessorModel, err := cmdProcessorModel.Output()
	if err != nil {
		cpuMap["processorModel"] = "Error: The processor model for the current system could not be retrieved."
	} else {
		cpuMap["processorModel"] = string(outputProcessorModel)
	}

	return cpuMap
}

// MemoryInfo retrieves all the memory information of the system like total memory, free
// memory, size of buffers, cache information, hypervisor and virtualization type.
func MemoryInfo() map[string]string {
	// Memory information
	memoryMap := make(map[string]string)

	// Total main memory
	cmdTotalMemory := exec.Command("/bin/sh", "-c", "cat /proc/meminfo | egrep -w 'MemTotal'")
	outputTotalMemory, err := cmdTotalMemory.Output()
	if err != nil {
		memoryMap["memoryTotal"] = "Error: The total memory for the current system could not be retrieved."
	} else {
		memoryMap["memoryTotal"] = string(outputTotalMemory)
	}

	// Total free memory
	cmdMemoryFree := exec.Command("/bin/sh", "-c", "cat /proc/meminfo | egrep -w 'MemFree'")
	outputMemoryFree, err := cmdMemoryFree.Output()
	if err != nil {
		memoryMap["memoryFree"] = "Error: The free memory for the current system could not be retrieved."
	} else {
		memoryMap["memoryFree"] = string(outputMemoryFree)
	}

	// Buffers
	cmdBuffers := exec.Command("/bin/sh", "-c", "cat /proc/meminfo | egrep -w 'Buffers'")
	outputBuffers, err := cmdBuffers.Output()
	if err != nil {
		memoryMap["buffers"] = "Error: The buffer information for the current system could not be retrieved."
	} else {
		memoryMap["buffers"] = string(outputBuffers)
	}

	// Cache
	cmdCache := exec.Command("/bin/sh", "-c", "cat /proc/meminfo | egrep -w 'Cached'")
	outputCache, err := cmdCache.Output()
	if err != nil {
		memoryMap["cache"] = "Error: The cached memory information for the current system could not be retrieved."
	} else {
		memoryMap["cache"] = string(outputCache)
	}

	cmdCacheInfo := exec.Command("/bin/sh", "-c", "lscpu | egrep -w 'cache'")
	outputCacheInfo, err := cmdCacheInfo.Output()
	if err != nil {
		memoryMap["cacheInfo"] = "Error: The cache information for the current system could not be retrieved."
	} else {
		memoryMap["cacheInfo"] = string(outputCacheInfo)
	}

	cmdVirtualizationInfo := exec.Command("/bin/sh", "-c", "lscpu | egrep -w 'Virtualization'")
	outputVirtualizationInfo, err := cmdVirtualizationInfo.Output()
	if err != nil {
		memoryMap["virtualizationInfo"] = "Error: The virtualization information for the current system could not be retrieved."
	} else {
		memoryMap["virtualizationInfo"] = string(outputVirtualizationInfo)
	}

	cmdHypervisorInfo := exec.Command("/bin/sh", "-c", "lscpu | egrep -w 'Hypervisor'")
	outputHypervisorInfo, err := cmdHypervisorInfo.Output()
	if err != nil {
		memoryMap["hypervisorInfo"] = "Error: The hypervisor information for the current system could not be retrieved."
	} else {
		memoryMap["hypervisorInfo"] = string(outputHypervisorInfo)
	}

	return memoryMap
}

// OtherSystemInfo retrieves information from the system like hostname, IPv4 address, pagesize,
// target architecture and target operating system.
func OtherSystemInfo() map[string]string {
	otherInfoMap := make(map[string]string)

	// Hostname
	hostname, err := os.Hostname()
	if err != nil {
		otherInfoMap["hostname"] = "Error: The hostname for the current system could not be retrieved."
	} else {
		otherInfoMap["hostname"] = hostname
	}

	// IP address
	addresses, err := net.LookupHost(hostname)
	if err != nil {
		otherInfoMap["ipv4_address"] = "Error: The IPv4 address for the current system could not be retrieved."
	} else {
		for _, address := range addresses {
			ipv4_address := net.ParseIP(address).To4()
			if ipv4_address != nil {
				otherInfoMap["ipv4_address"] = address
			}
		}
	}

	otherInfoMap["os_pagesize"] = strconv.Itoa(os.Getpagesize())
	otherInfoMap["target_architecture"] = runtime.GOARCH
	otherInfoMap["target_os"] = runtime.GOOS
	return otherInfoMap
}
