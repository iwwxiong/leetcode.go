# Substring with Concatenation of All Words

给定一个字符串 s ，给定 n 个单词 word，找出所有子串的开始下标，使得子串包含了给定的所有单词，顺序可以不对应。如果有重复的单词，比如有 [ " foo " , " foo " ] 那么子串也必须含有两个 " foo "，也就是说个数必须相同。

## 解题思路

1. 定义一个 HashMap mapV1，把 words 中间按照 key: 单词，value：出现次数处理；
2. 定义一个 HashMap mapV2，遍历字符串 s，如果遍历的单词存在与 mapV1，且在 mapV2 中出现次数小于等于 mapV1 中次数，则继续循环，否则跳出循环；
3. 当 mapV1 == mapV2 的时候则表示命中成功，记录下标；
