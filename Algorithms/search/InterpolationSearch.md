# Interpolation Search

## 概述
插值搜索是二分搜索的改进变体。该搜索算法在所需值的探测位置上起作用。为了使该算法正常工作，数据收集应采用排序形式并平均分配。

与线性搜索相比，二分搜索具有时间复杂性的巨大优势。线性搜索的最坏情况复杂度为Ο（n），而二进制搜索的复杂度为Ο（log n）。
## 算法
由于这是对现有BST算法的一种改进，因此我们提到了使用位置探测来搜索“目标”数据值索引的步骤-
```
Step 1 − Start searching data from middle of the list.
Step 2 − If it is a match, return the index of the item, and exit.
Step 3 − If it is not a match, probe position.
Step 4 − Divide the list using probing formula and find the new midle.
Step 5 − If data is greater than middle, search in higher sub-list.
Step 6 − If data is smaller than middle, search in lower sub-list.
Step 7 − Repeat until match.
```
## 代码模板
```c
A → Array list
N → Size of A
X → Target Value

Procedure Interpolation_Search()

   Set Lo  →  0
   Set Mid → -1
   Set Hi  →  N-1

   While X does not match
   
      if Lo equals to Hi OR A[Lo] equals to A[Hi]
         EXIT: Failure, Target not found
      end if
      
      Set Mid = Lo + ((Hi - Lo) / (A[Hi] - A[Lo])) * (X - A[Lo]) 

      if A[Mid] = X
         EXIT: Success, Target found at Mid
      else 
         if A[Mid] < X
            Set Lo to Mid+1
         else if A[Mid] > X
            Set Hi to Mid-1
         end if
      end if
   End While

End Procedure
```
## 实现
* C
```c
#include<stdio.h>

#define MAX 10

// array of items on which linear search will be conducted. 
int list[MAX] = { 10, 14, 19, 26, 27, 31, 33, 35, 42, 44 };

int find(int data) {
   int lo = 0;
   int hi = MAX - 1;
   int mid = -1;
   int comparisons = 1;      
   int index = -1;

   while(lo <= hi) {
      printf("\nComparison %d  \n" , comparisons ) ;
      printf("lo : %d, list[%d] = %d\n", lo, lo, list[lo]);
      printf("hi : %d, list[%d] = %d\n", hi, hi, list[hi]);
      
      comparisons++;

      // probe the mid point 
      mid = lo + (((double)(hi - lo) / (list[hi] - list[lo])) * (data - list[lo]));
      printf("mid = %d\n",mid);

      // data found 
      if(list[mid] == data) {
         index = mid;
         break;
      } else {
         if(list[mid] < data) {
            // if data is larger, data is in upper half 
            lo = mid + 1;
         } else {
            // if data is smaller, data is in lower half 
            hi = mid - 1;
         }
      }               
   }
   
   printf("\nTotal comparisons made: %d", --comparisons);
   return index;
}

int main() {
   //find location of 33
   int location = find(33);

   // if element was found 
   if(location != -1)
      printf("\nElement found at location: %d" ,(location+1));
   else
      printf("Element not found.");
   
   return 0;
}

```