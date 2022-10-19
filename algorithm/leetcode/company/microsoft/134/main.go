package main

func canCompleteCircuit(gas []int, cost []int) int {
	for i := 0; i < len(gas); {
		leftGas := 0
		j := 0
		for ; j < len(gas); j++ {
			k := (i + j) % len(gas)
			leftGas = leftGas + gas[k] - cost[k]
			if leftGas < 0 {
				break
			}
		}
		if j == len(gas) {
			return i
		} else {
			i = i + j + 1
		}
	}
	return -1
}
