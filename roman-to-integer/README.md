## 题目
    Given a roman numeral, convert it to an integer.

    Input is guaranteed to be within the range from 1 to 3999.

## 解题思路
    此题，最关键的信息是

    右加左减，左减数字必须为一位，比如8写成VIII，而非IIX。

    解题思路

        从右往左处理字符串。
        当前字符代表的数字，小于右边字符的时候，总体减去当前字符代表的数字。
        否则，总体加上当前字符代表的数字。
