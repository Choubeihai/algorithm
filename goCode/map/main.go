package main

func main() {
	myMap := make(map[string]int, 53)
	myMap["key"] = 1
	print(myMap["key"])
	delete(myMap, "key")
}
