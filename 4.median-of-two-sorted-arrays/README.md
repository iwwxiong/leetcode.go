# 获取两个有序数组的中位数

## 题目

从两个有序数组中获取中位数，要求算法时间为 `O(log (m+n))` 。

## 思路

简单做法就是通过合并两个有序数组，然后直接计算中位数即可（需考虑合并后长度奇偶性）。本体解答为了考虑时间问题，当合并数据长度满足中位数长度的时候，立即中断循环，然后计算中位数即可。
