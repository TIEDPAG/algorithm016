学习笔记

# Day 1
 - leetcode [两个数组的交集](https://github.com/TIEDPAG/leetcode-go/blob/master/t350/t350.go)

# Day 2
 - leetcode [滑动窗口的最大值](https://github.com/TIEDPAG/leetcode-go/blob/master/059/059.go)
 - HashMap
   - 底层数组容量设计为2的n次幂
     1. 扩容时使用左移即可
     2. 原本hash函数的取模变成了&运算， `n % length -> n & (length - 1)`
   - 哈希函数 使用了对象的hashcode的高十六位异或低十六位的结果，因为使用的hash取模方式换成了上面说的`&`的方式，直接使用hashcode则会让高位不参与运算，使用异或`^`的方式可以保证速度且保证hash结果较为均匀有效分布
   - 碰撞处理 对超过预设值长度的碰撞结果，将链表转化为红黑树，降低碰撞产生后的遍历时间复杂度
   - 内存分配 在首次插入元素的时候开始分配内存，使用最小容量
# Day3
  - leetcode [删除最外层括号](https://github.com/TIEDPAG/leetcode-go/blob/master/t1021/t1021.go)
# Day4
  - leetcode [FizzBuzz](https://github.com/TIEDPAG/leetcode-go/blob/master/t412/t412.go)
# Day5
  - leetcode 
    - [各位相加](https://github.com/TIEDPAG/leetcode-go/blob/master/t258/t258.go)
    - [K个一组翻转链表](https://github.com/TIEDPAG/leetcode-go/blob/master/t25/t25.go)
    - [有效的括号](https://github.com/TIEDPAG/leetcode-go/blob/master/t20/t20.go)
    - [最小栈](https://github.com/TIEDPAG/leetcode-go/blob/master/t155/t155.go)
    - [柱状图中最大的矩形](https://github.com/TIEDPAG/leetcode-go/blob/master/t84/t84.go)
    - [有效的字母异位词](https://github.com/TIEDPAG/leetcode-go/blob/master/t242/t242.go)
    - [字母异位词分组](https://github.com/TIEDPAG/leetcode-go/blob/master/t49/t49.go)
  - Go map
    - 源码实现分为两层
      1. map结构体，包含除一些描述性字段如key的长度外，主要包含哈希桶对象的指针；另外，有一个随机数结构，使得相同的key在不同的map里存储的位置不一致
        - 删除key时
      2. 哈希桶结构，结构体中只定义了哈希桶的哈希数据，实际上结合了unsafe包，使用指针偏移完成了一个链表结构
        - 哈希桶默认的长度为8，超过8则会触发overflow
        - 支持平滑升级
# Day6
  - leetcode
    - [二叉树的最大深度](https://github.com/TIEDPAG/leetcode-go/blob/master/t104/t104.go)
    - [二叉树的中序遍历](https://github.com/TIEDPAG/leetcode-go/blob/master/t94/t94.go)
    - [二叉树的前序遍历](https://github.com/TIEDPAG/leetcode-go/blob/master/t144/t144.go)
    - [N叉树的后序遍历](https://github.com/TIEDPAG/leetcode-go/blob/master/t590/t590.go)
    - [N叉树的前序遍历](https://github.com/TIEDPAG/leetcode-go/blob/master/t589/t589.go)
    - [N叉树的层序遍历](https://github.com/TIEDPAG/leetcode-go/blob/master/t429/t429.go)
    - [最小的k个数](https://github.com/TIEDPAG/leetcode-go/blob/master/o40/o40.go)
    - [丑数](https://github.com/TIEDPAG/leetcode-go/blob/master/o49/o49.go)
      