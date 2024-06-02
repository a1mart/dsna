package main

import (
    "hash/fnv"
    "math"
    "sync"
)

// BloomFilter struct
type BloomFilter struct {
    bitset        []bool
    numHashes     int
    expectedItems int
    size          int
    mutex         sync.RWMutex
}

// NewBloomFilter creates a new Bloom filter with the given expected number of items and false positive probability
func NewBloomFilter(expectedItems int, falsePositiveProbability float64) *BloomFilter {
    size := -1 * float64(expectedItems) * math.Log(falsePositiveProbability) / (math.Log(2) * math.Log(2))
    numHashes := int(math.Ceil((float64(size) / float64(expectedItems)) * math.Log(2)))
    return &BloomFilter{
        bitset:        make([]bool, int(size)),
        numHashes:     numHashes,
        expectedItems: expectedItems,
        size:          int(size),
    }
}

// Add inserts a new element into the Bloom filter
func (bf *BloomFilter) Add(element string) {
    bf.mutex.Lock()
    defer bf.mutex.Unlock()
    for _, i := range bf.getHashIndexes(element) {
        bf.bitset[i] = true
    }
}

// Contains checks if an element possibly exists in the Bloom filter
func (bf *BloomFilter) Contains(element string) bool {
    bf.mutex.RLock()
    defer bf.mutex.RUnlock()
    for _, i := range bf.getHashIndexes(element) {
        if !bf.bitset[i] {
            return false
        }
    }
    return true
}

// getHashIndexes calculates the hash indexes for an element
func (bf *BloomFilter) getHashIndexes(element string) []uint64 {
    indexes := make([]uint64, bf.numHashes)
    hash1 := fnv.New64a()
    hash1.Write([]byte(element))
    hashSum1 := hash1.Sum64()

    hash2 := fnv.New64a()
    hash2.Write([]byte(element))
    hashSum2 := hash2.Sum64()

    for i := 0; i < bf.numHashes; i++ {
        indexes[i] = (hashSum1 + uint64(i)*hashSum2) % uint64(bf.size)
    }

    return indexes
}

// BloomObject includes a Bloom filter address to check before accessing
type BloomObject struct {
    Data       interface{}
    BloomCheck *BloomFilter
}

// NewBloomObject creates a new BloomObject with the given data and Bloom filter
func NewBloomObject(data interface{}, bf *BloomFilter) *BloomObject {
    return &BloomObject{
        Data:       data,
        BloomCheck: bf,
    }
}

// AccessData checks if the data may exist in the Bloom filter before accessing
func (bo *BloomObject) AccessData(element string) interface{} {
    if bo.BloomCheck.Contains(element) {
        return bo.Data
    }
    return nil
}

func main() {
    // Example usage
    expectedItems := 100000
    falsePositiveProbability := 0.01 // 1% false positive probability
    bloomFilter := NewBloomFilter(expectedItems, falsePositiveProbability)

    // Adding some elements to the Bloom filter
    elements := []string{"apple", "banana", "orange", "grape", "melon"}
    for _, element := range elements {
        bloomFilter.Add(element)
    }

    // Creating a BloomObject with some data and the Bloom filter
    data := "Some important data"
    bloomObj := NewBloomObject(data, bloomFilter)

    bloomObj.BloomCheck.Add("tomato")

    // Accessing data by checking Bloom filter first
    elementToCheck := "tomato"
    if result := bloomObj.AccessData(elementToCheck); result != nil {
        println("Data exists:", result)
    } else {
        println("Data does not exist or may not exist in the Bloom filter.")
    }
}
