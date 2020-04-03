# 插入排序-InsertionSort
插入排序是一种简单的排序算法，可以一次构建一个最终排序的数组（或列表）。与大型高级算法（如quicksort，heapsort或merge sort）相比，它在大型列表上的效率要低得多

## 基本思想
插入排序这是基于比较的排序算法。在此，将始终维护一个子列表。例如，数组的下部保持被排序。要“插入”此排序子列表中的元素，必须找到其合适的位置，然后将其插入到该位置。因此名称为insert sort。

依次搜索该数组，然后将未排序的项目移动并插入到已排序的子列表中（在同一数组中）。

该算法不适用于大型数据集，因为其平均和最坏情况下的复杂度均为O(N^2).

插入排序是一种最简单直观的排序算法，它的工作原理是通过构建有序序列，对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
## 排序过程
- 步骤1:如果它是第一个元素，则它已经排序。 返回1;
- 第2步:选择下一个元素
- 步骤3:与排序的子列表中的所有元素进行比较
- 步骤4:移动已排序子列表中大于要排序的值的所有元素
- 步骤5:插入值
- 步骤6:重复直到列表被排序

![InsertSort](../images/insertion-sort.png)

![InsertSort](../images/insertionsort.png)


## 复杂度分析
- 时间复杂度    `O(N^2)`
- 空间复杂度    `O(1)`
## 代码模板
```c
procedure insertionSort( A : array of items )
   int holePosition
   int valueToInsert
	
   for i = 1 to length(A) inclusive do:
	
      /* select value to be inserted */
      valueToInsert = A[i]
      holePosition = i
      
      /*locate hole position for the element to be inserted */
		
      while holePosition > 0 and A[holePosition-1] > valueToInsert do:
         A[holePosition] = A[holePosition-1]
         holePosition = holePosition -1
      end while
		
      /* insert the number at hole position */
      A[holePosition] = valueToInsert
      
   end for
	
end procedure
```
## 代码实现
* C
```c
#include <stdio.h>
#include <stdbool.h>

#define MAX 7

int intArray[MAX] = {4,6,3,2,1,9,7};

void printline(int count) {
   int i;
	
   for(i = 0;i < count-1;i++) {
      printf("=");
   }
	
   printf("=\n");
}

void display() {
   int i;
   printf("[");
	
   // navigate through all items 
   for(i = 0;i < MAX;i++) {
      printf("%d ",intArray[i]);
   }
	
   printf("]\n");
}

void insertionSort() {

   int valueToInsert;
   int holePosition;
   int i;
  
   // loop through all numbers 
   for(i = 1; i < MAX; i++) { 
	
      // select a value to be inserted. 
      valueToInsert = intArray[i];
		
      // select the hole position where number is to be inserted 
      holePosition = i;
		
      // check if previous no. is larger than value to be inserted 
      while (holePosition > 0 && intArray[holePosition-1] > valueToInsert) {
         intArray[holePosition] = intArray[holePosition-1];
         holePosition--;
         printf(" item moved : %d\n" , intArray[holePosition]);
      }

      if(holePosition != i) {
         printf(" item inserted : %d, at position : %d\n" , valueToInsert,holePosition);
         // insert the number at hole position 
         intArray[holePosition] = valueToInsert;
      }

      printf("Iteration %d#:",i);
      display();
		
   }  
}

void main() {
   printf("Input Array: ");
   display();
   printline(50);
   insertionSort();
   printf("Output Array: ");
   display();
   printline(50);
}
```
* Java
```java
public int[] sort(int[] arr)  {

    int temp;

	for(int i = 0;i < lenth - 1; i++){
		for(int j = i + 1; j > 0; j--){
		    if(array[j] < array[j - 1]){
		        temp = array[j - 1];
		        array[j - 1] = array[j];
		        array[j] = temp;
		    }else{         //不需要交换
		        break;
		    }
		}
	}

    return arr;
}

```
* GO
```golang
func insertionSort(arr []int) []int {
    for i := range arr {
        preIndex := i - 1
        current := arr[i]
        
        for preIndex >= 0 && arr[preIndex] > current {
            arr[preIndex+1] = arr[preIndex]
            preIndex -= 1
        }
        
        arr[preIndex+1] = current
    }
    
        return arr
}
```
* Python
```python
def insertionSort(alist):
   for index in range(1,len(alist)):

     currentvalue = alist[index]
     position = index

     while position > 0 and alist[position - 1] > currentvalue:
         alist[position] = alist[position - 1]
         position = position - 1

     alist[position]=currentvalue
```
* JavaScript
```javascript
function insertionSort(arr) {
    var len = arr.length;
    var preIndex, current;
    for (var i = 1; i < len; i++) {
        preIndex = i - 1;
        current = arr[i];
        while(preIndex >= 0 && arr[preIndex] > current) {
            arr[preIndex+1] = arr[preIndex];
            preIndex--;
        }
        arr[preIndex+1] = current;
    }
    return arr;
}
```
* Rust
```rust
```