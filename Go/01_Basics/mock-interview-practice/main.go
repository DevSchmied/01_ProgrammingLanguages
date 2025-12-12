package main

import (
	"fmt"
	"sync"
)

// Task 3
var (
	mTask3 = make(map[string]string)
	mu     sync.Mutex
)

// Task 4
var mTask4 sync.Map

func main() {
	// Task 1. What will be printed to the console?
	// Explain how string iteration works in Go when using range,
	// especially with UTF-8 encoded characters (runes).

	fmt.Println("Task 1")

	str1 := "hello 🙂💪"
	str2 := "привет 🙂💪"

	for i, char := range str1 {
		fmt.Println(i, ": ", char)
	}

	fmt.Println()
	for i, char := range str1 {
		fmt.Printf("%d: %c\n", i, char)
	}

	fmt.Println()
	fmt.Println()

	for i, char := range str2 {
		fmt.Println(i, ": ", char)
	}

	fmt.Println()
	for i, char := range str2 {
		fmt.Printf("%d: %c\n", i, char)
	}

	// Task 2. Create a small concurrent program in Go that demonstrates a data race
	// when multiple goroutines access and modify shared state without proper synchronization.
	fmt.Println()
	fmt.Println("Task 2")

	m := map[int]*int{}
	val := 0
	m[0] = &val

	for i := 0; i < 30; i++ {
		go func(i int) {
			cur := m[0]
			*cur += i
		}(i)
	}

	fmt.Println(*m[0])

	/*
		Task 3.
		Design a concurrency-safe key–value cache in Go that allows multiple goroutines to access shared data.
		The cache should return an existing value for a given key if it is already present, or create and store a new value otherwise.
		Ensure that all operations are safe under concurrent access.
	*/
	fmt.Println()
	fmt.Println("Task 3")

	fmt.Println(GetOrCreate("hello", "world"))
	fmt.Println(GetOrCreate("hello", "world"))
	fmt.Println(Get("hello"))

	/*
			Task 4:
		Implement an additional thread-safe “get or create” function using sync.Map.
		The new function must provide the same behavior as the existing implementation, ensure correct concurrent access, and guarantee that a value is created only once per key.
	*/

	fmt.Println()
	fmt.Println("Task 4")

	fmt.Println(GetOrCreateSyncMap("hello", "world"))
	fmt.Println(GetOrCreateSyncMap("hello", "world"))
	fmt.Println(GetOrCreateSyncMap("hello", "you"))

}

// Task 3.
func GetOrCreate(key, value string) string {
	mu.Lock()
	defer mu.Unlock()
	if v, ok := mTask3[key]; ok {
		return v
	}

	mTask3[key] = value
	return value
}

func Get(key string) string {
	mu.Lock()
	defer mu.Unlock()
	if v, ok := mTask3[key]; ok {
		return v
	}
	return ""
}

// Task 4
func GetOrCreateSyncMap(key, value string) string {
	if v, found := mTask4.Load(key); found {
		str, ok := v.(string)
		if ok {
			return str
		}
		return ""
	}

	v, _ := mTask4.LoadOrStore(key, value)

	return v.(string)
}
