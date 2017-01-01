package models

type CpuInfo struct {
	Idle float64
	Used float64
}

type DiskInfo struct {
	Free  float64
	Total float64
	In    float64
	Out   float64
}

type Memory struct {
	Free float64
	Buffer float64
	Cache float64
	Swap float64
}

type Network struct {
	In float64
	Out float64
}


