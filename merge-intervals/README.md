## 题目
    Given a collection of intervals, merge all overlapping intervals.

    For example,
    Given [1,3],[2,6],[8,10],[15,18],
    return [1,6],[8,10],[15,18].


    
    先对 intervals 进行排序，按照 Start 递增
    依次处理重叠的情况。
