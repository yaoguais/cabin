package main

import "fmt"
import "math/rand"
import "time"
import "sync"

func main() {
	length := 1000000
	arrSlice := make([]int, length)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		arrSlice[i] = rand.Intn(length)
	}
	
	arrSliceForSort := make([]int, length)
	arrSliceForSort2 := make([]int, length)
	for i := 0; i < length; i++ {
		arrSliceForSort[i] = arrSlice[i]
		arrSliceForSort2[i] = arrSlice[i]
	}
	
	fmt.Printf("length: %d\n", length);
	startNano := time.Now().UnixNano()
	merge_sort(arrSliceForSort, length)
	fmt.Printf("normal merge sort: \n%d(ns)\n", time.Now().UnixNano() - startNano);		

	startNano = time.Now().UnixNano()
	merge_sort2(arrSliceForSort2, length)
	fmt.Printf("routine merge sort: \n%d(ns)\n", time.Now().UnixNano() - startNano);		

	time.Sleep(5 * time.Second)
}

func merge_sort_routine(arr []int, arrCopy []int, left int, right int) {
	if left < right {
		mid := (left + right) / 2
		var wg sync.WaitGroup
		wg.Add(2)
		go func() {
			defer wg.Done()
			merge_sort_routine(arr, arrCopy, left, mid)
		}()
		go func() {
			defer wg.Done()
			merge_sort_routine(arr, arrCopy, mid+1, right)
		}()
		wg.Wait()
		// merge
		i := left
		for ; i <= right; i++ {
			arrCopy[i] = arr[i]
		}
		i = left
		j := mid + 1
		k := left
		for i <= mid && j <= right {
			if arrCopy[i] <= arrCopy[j] {
				arr[k] = arrCopy[i]
				i++
			} else {
				arr[k] = arrCopy[j]
				j++
			}
			k++
		}
		for i <= mid {
			arr[k] = arrCopy[i]
			i++
			k++
		}
		for j <= right {
			arr[k] = arrCopy[j]
			j++
			k++
		}
	}
}

func merge_sort_normal(arr []int, arrCopy []int, left int, right int) {
	if left < right {
		mid := (left + right) / 2
		merge_sort_normal(arr, arrCopy, left, mid)
		merge_sort_normal(arr, arrCopy, mid+1, right)
		// merge
		i := left
		for ; i <= right; i++ {
			arrCopy[i] = arr[i]
		}
		i = left
		j := mid + 1
		k := left
		for i <= mid && j <= right {
			if arrCopy[i] <= arrCopy[j] {
				arr[k] = arrCopy[i]
				i++
			} else {
				arr[k] = arrCopy[j]
				j++
			}
			k++
		}
		for i <= mid {
			arr[k] = arrCopy[i]
			i++
			k++
		}
		for j <= right {
			arr[k] = arrCopy[j]
			j++
			k++
		}
	}
}

func merge_sort(arr []int, length int) {
	arrCopy := make([]int, length)
	merge_sort_normal(arr, arrCopy, 0, length-1)
}

func merge_sort2(arr []int, length int) {
	arrCopy := make([]int, length)
	merge_sort_routine(arr, arrCopy, 0, length-1)
}
