# 第三题:好数组

>
> 时间限制：1.000S 空间限制：256MB
>

## 题目描述

>
> 小红定义一个数组是好数组，当且仅当该数组中有且仅有一个元素和其他元素不同，剩余的所有元素相同。
> 
> 例如，[2,2,2,5,2,2] 是好数组。
> 
> 小红拿到了一个数组，她可以进行若干次操作，每次操作可以使得一个元素加 1 或者减 1。小红希望用尽可能少的操作次数使得该数组变成好数组，你能帮帮她吗?
>

## 输入描述

>
> 第一行输入一个正整数 n（2 <= n <= 10^5），代表数组的大小。
> 
> 第二行输入 n 个正整数 ai（1 <= ai <= 10^9），代表数组的元素。
>

## 输出描述

>
> 输一个整数，代表最小的操作次数。
>

## 输入示例

```
6
2 2 4 5 1 2
```


## 输出示例

```
3
```

## 提示信息
>
> 对第三个元素进行两次减 1 操作，对第五个元素进行一次加 1 操作即可。
> 