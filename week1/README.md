# week 1 homeworks

## 本周作业

### [66. 加一](https://leetcode-cn.com/problems/plus-one/)（Easy)


[le66.go](le66.go)

### 2. 合并两个有序链表（Easy）

[21. 合并两个有序链表](https://leetcode-cn.com/problems/merge-two-sorted-lists/)

[le21.go](le21.go)

### [641. 设计循环双端队列](https://leetcode-cn.com/problems/design-circular-deque/)（Medium)

[le641.go](le641.go)

### [85. 最大矩形](https://leetcode-cn.com/problems/maximal-rectangle/)（Hard)

[le85.go](le85.go)

## 实战例题

### [88. 合并有序数组](https://leetcode-cn.com/problems/merge-sorted-array/)（Easy）

[le88.go](le88.go)

### [26. 删除有序数组中的重复项](https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array/) (Easy)

### [283. 移动零](https://leetcode-cn.com/problems/move-zeroes/) (Easy)

[le283.go](le283.go)

## 链表

### [206. 反转链表](https://leetcode-cn.com/problems/reverse-linked-list/) (Easy)

[le206.go](le206.go)

### [25. K个一组翻转链表](https://leetcode-cn.com/problems/reverse-nodes-in-k-group/) （Hard)

[le25.go](le25.go)

### 邻值查找（Medium）

### [141. 环形链表](https://leetcode-cn.com/problems/linked-list-cycle/)（Medium

[le141.go](le141.go)
解题思路：快慢指针，快指针跳两格，慢指针跳一格，如果相遇就是有环否则无环。

### [142. 环形链表 II](https://leetcode-cn.com/problems/linked-list-cycle-ii/) (Medium)

[le142.go](le142.go)
解题思路：在141的基础上，快慢指针相遇后，让新的指针从头开始和慢指针同步前进，相遇的点就是产生环的点。也可以用hashmap记录次数。

## 第 2 课 栈、队列

### [20. 有效的括号](https://leetcode-cn.com/problems/valid-parentheses/) (Medium)

[le20.go](le20.go)

解题思路：使用栈记录左括号，遇到右括号，则栈顶必须和其匹配，然后出栈继续下一个。最后栈必须为空。

### [155. 最小栈](https://leetcode-cn.com/problems/min-stack/)（Medium)

[le155.go](le155.go)

解题思路：维护两个栈，一个值的栈，一个最小值的栈，大小相等。同时push，同时pop，top就是最小值。

### [150. 逆波兰表达式求值](https://leetcode-cn.com/problems/evaluate-reverse-polish-notation/)（Medium）

[le150.go](le150.go)

解题思路：["2","1","+","3","*"], 遇到数字就压栈，遇到运算符就出栈两个数字进行计算再把结果压栈。最后栈顶就是结果。

### [224. 基本计算器](https://leetcode-cn.com/problems/basic-calculator/) (Hard)

[le224.go](le224.go)

解题思路：先生成逆波兰表达式(tokens)，使用stack记录运算符。遇到数字就放入tokens，遇到运算符，如果栈顶的运算符优先级高于`等于`当前运算符就把栈顶的运算符出栈存入tokens，把当前运算符压栈。
遇到左括号就压栈，遇到右括号就一直出栈运算符并存到tokens，直到遇到左括号。
数字的正负号可以通过前边补零来实现，"(+-*/"之后都是需要补零的。")0~9"是不需要补零的。这样遇到"+-"号先看是不是需要补零，如果需要就压一个"0"到tokens里。

思路儿：这是一个二则运算（+-），可以简化。

## 单调栈、单调队列

### [84. 柱状图中最大的矩形](https://leetcode-cn.com/problems/largest-rectangle-in-histogram/)（Hard）

[le84.go](le84.go)

解题思路：单调栈的方法主要思路是通过计算矩形的面积替代一个一个柱子加起来算面积加速算法。也就是用乘法代替加法。具体到本题，构建一个高度单调增的栈，记录高度和宽度。当遇到一个比栈顶`低`/`等`的柱子时，出栈并计算是不是最大面积，知道栈顶的高度小于当前柱子高度，向栈顶存放一个当前高度，出栈所有矩形宽度和+1的宽度的矩形。也就是把之前所有高于当前高度的矩形按照当前高度计算合并成一个大矩形，因为之前的高的矩形已经参与过计算了，可以忽略他们的高于当前高度的部分。

在所有柱子后补一个0高度的柱子，这样之前的矩形就都会被计算。

### [239. 滑动窗口最大值](https://leetcode-cn.com/problems/sliding-window-maximum/)（Hard）

[le239.go](le239.go)

给你一个整数数组 nums，有一个大小为 k 的滑动窗口从数组的最左侧移动到数组的最右侧。你只可以看到在滑动窗口内的 k 个数字。滑动窗口每次只向右移动一位。

解题思路：使用一个长度为最大为k的队列记录候选最大值的索引。队列中的索引是安顺讯排列的，但遇到一个大于等于队尾的值时，从尾部删除所有小于等于当前值的数据索引，因为这样的数据已经不再是候选者。保证队列的左侧（头部）的索引和当前位置的差小于等于k，否则就不是候选者了，需要清除。整理后队列的最左侧（最大）idx的值就是当前窗口的最大值。

### [42. 接雨水](https://leetcode-cn.com/problems/trapping-rain-water/)（Hard）

[le42.go](le42.go)

解题思路：和最大矩形类似，使用单调栈记录右侧递减的高度，当遇到一个高于当前的柱子时向前进行计算，并把计算过的部分填平成一个大矩形方便下次计算（填平的部分刚刚已经计算过了）。

也可以使用双指针法，计算一遍leftMax[n],rightMax[n]。
