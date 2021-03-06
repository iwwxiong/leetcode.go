# 按照Z字形转换字符串

## 题目

给定一个字符串和长度，按照Z字形转换字符串后按顺序输出。

## 思路

给定一个二维数组，然后根据2*n-2的长度拆分字符串，然后遍历拆分字符串追加到二维数组中对应位置中。

比如给定字符串 `0123456789ABCDEFG`

按照长度 `3` 转换字符串

```text
0   4   8   C   G
1 3 5 7 9 B D F
2   6   A   E
```

按照长度 `4` 转换字符串

```text
0     6     C
1   5 7   B D
2 4   8 A   E G
3     9     F
```

按照长度 `4` 转换字符串

```text
0       8       G
1     7 9     F
2   6   A   E
3 5     B D
4       C
```

## 结果

目前解题思路时间复杂度和空间复杂度都相对比较大，后期找寻新的解题思路。
