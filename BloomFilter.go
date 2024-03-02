package main

import (
  "fmt"
  "math"
  "github.com/spaolacci/murmur3"
)

type BloomFilter struct{
  filter []bool
  filterSize int
}

func (bloomFilter *BloomFilter) CreateFilter(FPRate float64,NoOfKeys int){
  size := bloomFilter.GetFilterSize(FPRate,NoOfKeys)
  filter := make([]bool, size)
  for i := range filter {
      filter[i] = false
  }

  bloomFilter.filter = filter
  bloomFilter.filterSize = size
}

func (bloomFilter *BloomFilter) GetIndex(Item string) uint32{
  hash := murmur3.New32()
  hash.Write([]byte(Item))
  hashSum := hash.Sum32()
  result := hashSum % uint32(bloomFilter.filterSize)
  return result
}

func (bloomFilter *BloomFilter) GetFilterSize(FPRate float64,NoOfKeys int) int{
  size := -(float64(NoOfKeys) * math.Log(FPRate))/(math.Pow(math.Log(2),3))
  return int(math.Abs(size))
}

func (bloomFilter *BloomFilter) Insert(Item string) {
  index := bloomFilter.GetIndex(Item)
  bloomFilter.filter[index] = true
}

func (bloomFilter *BloomFilter) Check(Item string) bool{
  index := bloomFilter.GetIndex(Item)
  return bloomFilter.filter[index]
}

func main() {
  FPRate := 0.2
  NoOfKeys := 10
  filter := BloomFilter{}
  filter.CreateFilter(FPRate,NoOfKeys)
  fmt.Println(filter.Check("apple")) //False
  filter.Insert("apple")
  fmt.Println(filter.Check("apple")) //True
  fmt.Println(filter.Check("banana")) //False
}