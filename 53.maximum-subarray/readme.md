# Maximum Subarray

动态规划思路

* 如果 dp [ i - 1 ] < 0，那么 dp [ i ] = nums [ i ]。
* 如果 dp [ i - 1 ] >= 0，那么 dp [ i ] = dp [ i - 1 ] + nums [ i ]
