## 题目
    Determine if a Sudoku is valid, according to: Sudoku Puzzles - The Rules.

    The Sudoku board could be partially filled, where empty cells are filled with the character '.'. 

    A partially filled sudoku which is valid.

    Note: A valid Sudoku board (partially filled) is not necessarily solvable. Only the filled cells need to be validated.


    数独的规则很简单：

    横着的9个数，不能重复
    竖着的9个数，不能重复
    3×3小方块中的9个数，也不能重复

    那么解题思路，就是依次检查这三个方面有没有重复。