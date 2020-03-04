package main

import (
    "fmt"
)

func quickSort(values []int,left,right int){
    var i,j,temp int
    i, j = left, right

    if i < j {
        temp = values[i] // 保存基准数据值

        for i < j {
            // 第一个循环从右往左计算,找到第一个 不大于基准值的位置
            for j > i && values[j] >= temp{
                j--
            }
            // 这一步还是需要的，万一不需要替换
            if i < j {
                values[i] = values[j] // 让左侧是 <= temp的值 这个地方的j对应的值其实现在的数组中是双份，覆盖了基准值
                i++
            }
            
            // 从左往右寻找，比temp
            for i < j && values[i] <= temp{
                i++
            }
            
            if i < j {
                values[j] = values[i] // 这样j的值改变了，i对应的值在数组中出现了双份
                j--
            }
        }
        // 循环结束后，可以看到i和j已经相碰了
        values[i] = temp // 把基准值覆盖循环中出现的双份i的值

        // 左右开始分别快速排序
        quickSort(values, left, i - 1)
        quickSort(values, i+1, right)
    }
 
}

func main() {
    // 测试代码
    arr := []int{9, 8, 7, 6, 5, 1, 2, 3, 4, 0}
    quickSort(arr, 0, len(arr)-1)
    fmt.Println(arr)
}