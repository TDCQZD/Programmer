# Linear search

## 概述
线性搜索是一种非常简单的搜索算法。在这种类型的搜索中，对所有项目进行逐个搜索。检查每个项目，如果找到匹配项，则返回该特定项目，否则继续搜索直到数据收集结束。

顺序搜索：在计算机科学中，线性搜索或顺序搜索是一种用于在列表中查找特定值的方法，该列表将按顺序检查每个元素，直到找到所需的元素或列表用尽为止

## 算法
```
Linear Search ( Array A, Value x)

Step 1: Set i to 1
Step 2: if i > n then go to step 7
Step 3: if A[i] = x then go to step 6
Step 4: Set i to i + 1
Step 5: Go to Step 2
Step 6: Print Element x Found at index i and go to step 8
Step 7: Print element not found
Step 8: Exit
```
## 复杂度分析
- 平均时间复杂度  `O(n)`

## 代码模板（伪代码）
```
procedure linear_search (list, value)

   for each item in the list
      if match item == value
         return the item's location
      end if
   end for

end procedure
```

## 实现
* c
```
#include <stdio.h>

#define MAX 20

// array of items on which linear search will be conducted.
int intArray[MAX] = {1,2,3,4,6,7,9,11,12,14,15,16,17,19,33,34,43,45,55,66};

void printline(int count) {
   int i;
	
   for(i = 0;i <count-1;i++) {
      printf("=");
   }
	
   printf("=\n");
}

// this method makes a linear search. 
int find(int data) {

   int comparisons = 0;
   int index = -1;
   int i;

   // navigate through all items 
   for(i = 0;i<MAX;i++) {
	
      // count the comparisons made 
      comparisons++;
		
      // if data found, break the loop
      if(data == intArray[i]) {
         index = i;
         break;
      }
   }   
	
   printf("Total comparisons made: %d", comparisons);
   return index;
}

void display() {
   int i;
   printf("[");
	
   // navigate through all items 
   for(i = 0;i<MAX;i++) {
      printf("%d ",intArray[i]);
   }
	
   printf("]\n");
}

void main() {
   printf("Input Array: ");
   display();
   printline(50);
	
   //find location of 1
   int location = find(55);

   // if element was found 
   if(location != -1)
      printf("\nElement found at location: %d" ,(location+1));
   else
      printf("Element not found.");
}
```
* python
```
def Sequential_Search(dlist, item):

    pos = 0
    found = False
    
    while pos < len(dlist) and not found:
        if dlist[pos] == item:
            found = True
        else:
            pos = pos + 1
    
    return found, pos
```