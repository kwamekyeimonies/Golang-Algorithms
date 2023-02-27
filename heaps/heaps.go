package heaps

type MaxHeap struct{
	array []int
}

//Insert adds an element to the heap
func (h *MaxHeap) Insert(key int){
	h.array = append(h.array, key)
	
}

func (h *MaxHeap) maxHeafifyup(index int){

}

func Parent(i int)int{

}

func Left(i int)int{

}

func Right(i int) int{

}


func (h *MaxHeap) swap(i1,i2 int){
	h.array[i1],h.array[i2] = h.array[i2], h.array[i1]
}