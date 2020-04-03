# 快速排序-QuickSort

## 基本思想 
快速排序使用分治思想排序。

快速排序是一种高效的排序算法，它基于将数据数组划分为较小的数组。将一个大型数组分为两个数组，其中一个数组的值小于指定值（例如，枢轴），基于该枢轴进行分区，另一个数组的值大于该枢轴值。

Quicksort对一个数组进行分区，然后递归调用两次以对两个结果子数组进行排序。该算法对于大型数据集非常有效，因为其平均和最坏情况下的复杂度分别为O(NlogN)和O(N^2).

## 排序过程
- 步骤1:选择具有索引的最高索引值
- 步骤2:将两个变量指向列表的左侧和右侧（不包括枢轴）
- 步骤3:左指向低索引
- 步骤4:指向高点
- 步骤5:当左侧的值小于枢轴右移时
- 步骤6:当右侧的值大于枢轴时向左移动
- 步骤7:如果步骤5和步骤6都不匹配，则左右互换
- 步骤8:如果左≥右，则它们相遇的点是新的枢轴

![QuickSort](../images/quick_sort_partition_animation.gif)

枢轴值将列表分为两部分。递归地，我们找到每个子列表的枢轴，直到所有列表仅包含一个元素

![QuickSort](../images/partitionA.png)

## 复杂度分析
- 时间复杂度    `O(N*logN)`
- 空间复杂度    `O(logN)`
## 代码模板
```c
procedure quickSort(left, right)

   if right-left <= 0
      return
   else     
      pivot = A[right]
      partition = partitionFunc(left, right, pivot)
      quickSort(left,partition-1)
      quickSort(partition+1,right)    
   end if		
   
end procedure

function partitionFunc(left, right, pivot)
   leftPointer = left
   rightPointer = right - 1

   while True do
      while A[++leftPointer] < pivot do
         //do-nothing            
      end while
		
      while rightPointer > 0 && A[--rightPointer] > pivot do
         //do-nothing         
      end while
		
      if leftPointer >= rightPointer
         break
      else                
         swap leftPointer,rightPointer
      end if
		
   end while 
	
   swap leftPointer,right
   return leftPointer
	
end function
```
## 代码实现
* C
```c
现场演示
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

void swap(int num1, int num2) {
   int temp = intArray[num1];
   intArray[num1] = intArray[num2];
   intArray[num2] = temp;
}

int partition(int left, int right, int pivot) {
   int leftPointer = left -1;
   int rightPointer = right;

   while(true) {
      while(intArray[++leftPointer] < pivot) {
         //do nothing
      }
		
      while(rightPointer > 0 && intArray[--rightPointer] > pivot) {
         //do nothing
      }

      if(leftPointer >= rightPointer) {
         break;
      } else {
         printf(" item swapped :%d,%d\n", intArray[leftPointer],intArray[rightPointer]);
         swap(leftPointer,rightPointer);
      }
   }
	
   printf(" pivot swapped :%d,%d\n", intArray[leftPointer],intArray[right]);
   swap(leftPointer,right);
   printf("Updated Array: "); 
   display();
   return leftPointer;
}

void quickSort(int left, int right) {
   if(right-left <= 0) {
      return;   
   } else {
      int pivot = intArray[right];
      int partitionPoint = partition(left, right, pivot);
      quickSort(left,partitionPoint-1);
      quickSort(partitionPoint+1,right);
   }        
}

int main() {
   printf("Input Array: ");
   display();
   printline(50);
   quickSort(0,MAX-1);
   printf("Output Array: ");
   display();
   printline(50);
}
```
* GO
```golang
func QiuckSort(arrays []int) {
	quickSortHelper(arrays, 0, len(arrays)-1)
}
func quickSortHelper(arrays []int, first, last int) {
	if last < first {
		return
	}

	splitpoint := partition(arrays, first, last)
	quickSortHelper(arrays, first, splitpoint-1)
	quickSortHelper(arrays, splitpoint+1, last)

}

func partition(arrays []int, first, last int) int {
	pivotvalue := arrays[first]

	leftMark := first + 1
	rightMark := last

	for {
		for rightMark >= leftMark && arrays[leftMark] <= pivotvalue {
			leftMark = leftMark + 1
		}
		for rightMark >= leftMark && arrays[rightMark] >= pivotvalue {
			rightMark = rightMark - 1
		}

		if rightMark < leftMark {
			break
		} else {
			arrays[leftMark], arrays[rightMark] = arrays[rightMark], arrays[leftMark]
		}	
	}
	
	arrays[first], arrays[rightMark] = arrays[rightMark], arrays[first]
	return rightMark
}
```
* Java
```java
public int[] quickSort(int[] arr, int left, int right) {
    if (left < right) {
        int partitionIndex = partition(arr, left, right);
        quickSort(arr, left, partitionIndex - 1);
        quickSort(arr, partitionIndex + 1, right);
    }
    return arr;
 }

public int partition(int[] arr, int left, int right) {
    // 设定基准值（pivot）
     int pivot = left;
    int index = pivot + 1;
    for (int i = index; i <= right; i++) {
        if (arr[i] < arr[pivot]) {
            swap(arr, i, index);
            index++;
        }
    }
    swap(arr, pivot, index - 1);
    return index - 1;
}

public void swap(int[] arr, int i, int j) {
    int temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}
```
* Python
```python
def quickSort(alist):
   quickSortHelper(alist, 0, len(alist)-1)

def quickSortHelper(alist, first, last):
   if first < last:

       splitpoint = partition(alist, first, last)

       quickSortHelper(alist, first, splitpoint-1)
       quickSortHelper(alist, splitpoint + 1, last)


def partition(alist, first, last):
   pivotvalue = alist[first]

   leftmark = first + 1
   rightmark = last

   done = False
   while not done:

       while leftmark <= rightmark and alist[leftmark] <= pivotvalue:
           leftmark = leftmark + 1

       while alist[rightmark] >= pivotvalue and rightmark >= leftmark:
           rightmark = rightmark -1

       if rightmark < leftmark:
           done = True
       else:
           alist[leftmark], alist[rightmark] = alist[rightmark],  alist[leftmark]

   alist[first], alist[rightmark] = alist[rightmark],  alist[first] 


   return rightmark

```
* JavaScript
```javascript
function quickSort(arr, left, right) {
    var len = arr.length,
        partitionIndex,
        left = typeof left != 'number' ? 0 : left,
        right = typeof right != 'number' ? len - 1 : right;

    if (left < right) {
        partitionIndex = partition(arr, left, right);
        quickSort(arr, left, partitionIndex-1);
        quickSort(arr, partitionIndex+1, right);
    }
    return arr;
}

function partition(arr, left ,right) {     // 分区操作
    var pivot = left,                      // 设定基准值（pivot）
        index = pivot + 1;
    for (var i = index; i <= right; i++) {
        if (arr[i] < arr[pivot]) {
            swap(arr, i, index);
            index++;
        }        
    }
    swap(arr, pivot, index - 1);
    return index-1;
}

function swap(arr, i, j) {
    var temp = arr[i];
    arr[i] = arr[j];
    arr[j] = temp;
}
function partition2(arr, low, high) {
  let pivot = arr[low];
  while (low < high) {
    while (low < high && arr[high] > pivot) {
      --high;
    }
    arr[low] = arr[high];
    while (low < high && arr[low] <= pivot) {
      ++low;
    }
    arr[high] = arr[low];
  }
  arr[low] = pivot;
  return low;
}

function quickSort2(arr, low, high) {
  if (low < high) {
    let pivot = partition2(arr, low, high);
    quickSort2(arr, low, pivot - 1);
    quickSort2(arr, pivot + 1, high);
  }
  return arr;
}
```
* Rust
```rust
```
