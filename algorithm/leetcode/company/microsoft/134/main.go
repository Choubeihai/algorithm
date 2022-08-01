package main

func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	var i = 0
	for i < n {
		var j = 0
		var leftGas = 0
		for j < n {
			k := (i + j) % n
			leftGas = leftGas + (gas[k] - cost[k])
			if leftGas < 0 {
				break
			}
			j++
		}
		if j == n {
			return i
		}
		i = i + j + 1
	}
	return -1
}
