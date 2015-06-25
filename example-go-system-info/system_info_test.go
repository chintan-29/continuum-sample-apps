// Copyright 2014 Apcera Inc. All rights reserved.

package main

import (
	tt "github.com/apcera/util/testtool"
	"testing"
)

func TestCpuInfomation(t *testing.T) {
	tt.StartTest(t)
	defer tt.FinishTest(t)

	cpuMap := CpuInfo()
	tt.TestNotEqual(t, cpuMap["logicalCores"], "", "Number of logical cores in a system cannot be less than 1.")
	tt.TestNotEqual(t, cpuMap["physicalCores"], "", "Number of physical cores in a system cannot be less than 1.")
	tt.TestNotEqual(t, cpuMap["processorModel"], "", "Processor model must exist in the system.")
}

func TestMemoryInfomation(t *testing.T) {
	tt.StartTest(t)
	defer tt.FinishTest(t)

	memoryMap := MemoryInfo()
	tt.TestNotEqual(t, memoryMap["memoryTotal"], "", "Memory information of the system could not be retrieved.")
	tt.TestNotEqual(t, memoryMap["memoryFree"], "", "Cache information of the system could not be retrieved.")
	tt.TestNotEqual(t, memoryMap["buffers"], "", "Memory information of the system could not be retrieved.")
	tt.TestNotEqual(t, memoryMap["cache"], "", "Cache information of the system could not be retrieved.")
	tt.TestNotEqual(t, memoryMap["cacheInfo"], "", "Memory information of the system could not be retrieved.")
	tt.TestNotEqual(t, memoryMap["virtualizationInfo"], "", "Cache information of the system could not be retrieved.")
	tt.TestNotEqual(t, memoryMap["hypervisorInfo"], "", "Memory information of the system could not be retrieved.")
}

func TestOtherSystemInfomation(t *testing.T) {
	tt.StartTest(t)
	defer tt.FinishTest(t)

	otherSystemInfoMap := OtherSystemInfo()
	tt.TestNotEqual(t, otherSystemInfoMap["hostname"], "", "Hostname of the system could not be retrieved.")
	tt.TestNotEqual(t, otherSystemInfoMap["ipv4_address"], "", "IPv4 address of the system could not be retrieved.")
	tt.TestNotEqual(t, otherSystemInfoMap["os_pagesize"], "", "OS page size of the system could not be retrieved.")
	tt.TestNotEqual(t, otherSystemInfoMap["target_architecture"], "", "Target architecture of the system could not be retrieved.")
	tt.TestNotEqual(t, otherSystemInfoMap["target_os"], "", "Target OS of the system could not be retrieved.")
}
