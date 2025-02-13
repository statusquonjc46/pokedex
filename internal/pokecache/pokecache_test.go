package pokecache

import (
	"fmt"
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	// Create a cache with some interval (like 5 seconds)
	cache := NewCache(5 * time.Second)

	cases := []struct {
		testKey string
		testVal []byte
	}{
		{
			testKey: "testKey",
			testVal: []byte("testValue"),
		},
		{
			testKey: "GO",
			testVal: []byte("EAG"),
		},
		{
			testKey: "BIRDS",
			testVal: []byte("ELS"),
		},
	}
	// Add some test data
	for _, c := range cases {
		cache.Add(c.testKey, c.testVal)
	}

	// Try to get the data back
	for _, x := range cases {
		//cache.Get(x.testKey)
		val, ok := cache.Get(x.testKey)
		if !ok {
			t.Error("Expected to find key, but didn't")
		}
		if string(val) != string(x.testVal) {
			t.Errorf("val: %s does not equal testval: %s", string(val), string(x.testVal))
		} else {
			success := fmt.Sprintf("val: %s does not equal testval: %s", string(val), string(x.testVal))
			fmt.Println(success)
		}
	}
}

func TestReapLoop(t *testing.T) {
	// Create a cache with a very short interval (maybe milliseconds)
	cache := NewCache(5 * time.Millisecond)
	_, ok := cache.Get("doesn't exist")
	if ok {
		t.Error("Expected false for non-existent key")
	}
	// Add an item
	testK := "testK"
	testV := []byte("testV")
	cache.Add(testK, testV)
	// Verify it exists
	//cache.Get(testK)
	val, ok := cache.Get(testK)
	if !ok {
		t.Error("Expected to find key, but didn't")
	}
	if string(val) != string(testV) {
		fmt.Sprintf("val: %s does not equal testval: %s", string(val), string(testV))
	}
	// Wait longer than the interval
	time.Sleep(1 * time.Second)
	// Verify it's gone
	_, ok = cache.Get(testK)
	if ok {
		t.Error("Expected to find key, but didn't")
	} else {
		fmt.Println("removed via reaploop")
	}
}
