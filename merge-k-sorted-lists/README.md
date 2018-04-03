## 题目
    Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.


## 解决思路
    leetcode上最快的解决思路是 借鉴了归并排序的思想，让lists中临近的list先两两合并，再让新生成的lists中的临近list两两合并，直到合并完成。