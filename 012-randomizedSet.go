package main

import (
	"math/rand"
)

type RandomizedSet struct {
	set        map[int]int
	setIndexes []int
}

func Constructor() RandomizedSet {
	randomizedSet := RandomizedSet{
		set:        make(map[int]int),
		setIndexes: make([]int, 0),
	}
	return randomizedSet
}

func (this *RandomizedSet) Insert(val int) bool {

	if _, exists := this.set[val]; exists {
		return false
	}

	size := len(this.setIndexes)
	this.setIndexes = append(this.setIndexes, val)

	this.set[val] = size

	return true
}

func (this *RandomizedSet) Remove(val int) bool {

	if i, exists := this.set[val]; exists {
		delete(this.set, val)

		lastVal := this.setIndexes[len(this.setIndexes)-1]
		this.setIndexes[i] = lastVal
		this.setIndexes = this.setIndexes[:len(this.setIndexes)-1]

		if _, exists := this.set[lastVal]; exists {
			this.set[lastVal] = i
		}

		return true
	}

	return false
}

func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.set))
	return this.setIndexes[index]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
