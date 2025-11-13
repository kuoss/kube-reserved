package calc

const (
	_ = 1 << (iota * 10)
	Ki
	Mi
	Gi
)

func CalculateCPUReservation(cpuMilli int64) int64 {
	reservation := int64(0)
	// 6% of first core
	if cpuMilli > 0 {
		reservation += 60
	}
	// + 1% of second core
	if cpuMilli > 1000 {
		reservation += 10
	}
	// + 0.5% each for core 3 and 4
	if cpuMilli > 2000 {
		reservation += (min(cpuMilli/1000, 4) - 2) * 5
	}
	// + 0.25% for the remaining CPU cores
	if cpuMilli > 4000 {
		reservation += (cpuMilli/1000 - 4) * 5 / 2
	}

	return reservation
}

func CalculateMemoryReservation(memory int64) int64 {
	reservation := int64(0)
	if memory < 1*Gi {
		reservation = 255 * Mi
	}
	// 25% of first 4 Gi
	if memory >= 1*Gi {
		reservation += min(memory, 4*Gi) / 4
	}
	// 20% for additional memory between 4Gi and 8Gi
	if memory >= 4*Gi {
		reservation += (min(memory, 8*Gi) - 4*Gi) / 5
	}
	// 10% for additional memory between 8Gi and 16Gi
	if memory >= 8*Gi {
		reservation += (min(memory, 16*Gi) - 8*Gi) / 10
	}
	// 6% for additional memory between 16Gi and 128Gi
	if memory >= 16*Gi {
		reservation += (min(memory, 128*Gi) - 16*Gi) / 100 * 6
	}
	// 2% of remaining memory
	if memory >= 128*Gi {
		reservation += (memory - 128*Gi) / 100 * 2
	}

	return reservation
}
