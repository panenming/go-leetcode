## 题目
    Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

    For "(()", the longest valid parentheses substring is "()", which has length = 2.

    Another example is ")()())", where the longest valid parentheses substring is "()()", which has length = 4.

    1.记录每个符号的状态，(一律对应于0；)如果能够和前面的配上对，就记录为2，否则，记录为0

    ) ( ( ) ( ) ) ) ( ( ( ( ( ) ) ) ) (

    形成记录

    0 0 0 2 0 2 2 0 0 0 0 0 0 2 2 2 2 0

    2.从左往右检查record，如果record[i]==2，record[j]==0，且record[j+1:i]中没有0，则record[i]=1,record[j]=1，那么，上面的record就变成了

    0 1 1 1 1 1 1 0 0 1 1 1 1 1 1 1 1 0

    3.统计record中，最多的连续为1的次数，就是结果。