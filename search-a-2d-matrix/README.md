## 题目
    Write an efficient algorithm that searches for a value in an m x n matrix. This matrix has the following properties:

    Integers in each row are sorted from left to right.
    The first integer of each row is greater than the last integer of the previous row.

    For example, Consider the following matrix:

    [
        [1,   3,  5,  7],
        [10, 11, 16, 20],
        [23, 30, 34, 50]
    ]

    Given target = 3, return true.

    原有的作者的解法在查找行时没有采用二分查找法，我的方法在行查找的时候也加入了二分查找法，而且优化了在第一位就查找到元素的情况，执行速度有了部分提升。