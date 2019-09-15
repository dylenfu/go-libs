package disorder

import "testing"

/*
3. 统计数字
中文English
计算数字 k 在 0 到 n 中的出现的次数，k 可能是 0~9 的一个值。

Example
样例 1：

输入：
k = 1, n = 1
输出：
1
解释：
在 [0, 1] 中，我们发现 1 出现了 1 次 (1)。
样例 2：

输入：
k = 1, n = 12
输出：
5
解释：
在 [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12] 中，我们发现 1 出现了 5 次 (1, 10, 11, 12)(注意11中有两个1)。
*/

/*
解题思路
这种题其实是很明显的归纳法，这里我们先举个例子(后面括号展示数据，比如ab(a,b)，数量用cnt表示)，从个位，十位，百位等位数来分析

如果n是两位数, 比如 n=53,
1.个位分析, if k<3 e.g k=2, cnt=5+1; if k>3 e.g k=4, cnt=5; if k==3, cnt=5+1
2.十位分析, if k<5 e.g k=2, cnt+=1*10; if k>5 e.g k=6, cnt+=0; if k==5, cnt+= 3+1(这里有50，51，52，53)

如果n是三位数, 比如 n=453
1.个位分析, if k<3 e.g k=2, cnt=45+1; if k>3 e.g k=4, cnt=45; if k==3, cnt=45+1
2.十位分析, if k<5 e.g k=2, cnt+=(4+1)*10; if k>5 e.g k=6, cnt+=4*10; if k==5, cnt+=4*10+3+1([50,53])
3.百位分析, if k<4 e.g k=2, cnt+=1*100; if k>4 e.gk=6, cnt+=0; if k==4, cnt+=53+1([450, 453])

按位数来分析,比如以n=453, 从十位(5)的角度来看，5前面有4，后面有3
百位比十位大一级，在[0,99]或者[100, 199],[200,299],[300,399]这个范围内，必然有10个k，百位为4那么就会有4*10，
剩下[400,453],最多还会有10个数，十位为k，
如果k>5 e.g k=6,那么查询范围就在460以上，个数为0，
如果k<5 e.g k=3，那么[430,439]就有10个,
如果k==5，那么[450, 453]就有3+1个(包含0)

归纳一下，如果分析位i，为current(刚才的十位), 那么百位就是before，个位就是after
if k < current cnt+= (before+1) * (10^i)
else if k > current cnt += before*(10^i)
else if k == current cnt += before*(10^i) + after + 1

除此以外，还有两种情况:
k == 0, 0在[0, n]之间是逢十加一,
n == 0, 如果k也等于0返回1，否则为0
*/

/**
 * @param k: An integer
 * @param n: An integer
 * @return: An integer denote the count of digit k in 1..n
 */
func digitCounts(k int, n int) int {
	var (
		cnt, current, before, after int = 0, 0, 0, 0
		i                           int = 1
	)

	if k == 0 {
		return n/10 + 1
	}
	if n == 0 {
		return 0
	}

	for n/i > 0 {
		current = (n / i) % 10
		before = n / (i * 10)
		after = n - (n/i)*i

		if k < current {
			cnt += (before + 1) * i
		} else if k > current {
			cnt += before * i
		} else {
			cnt += before*i + after + 1
		}
		i *= 10
	}

	return cnt
}

// go test -v github.com/dylenfu/go-libs/disorder/lintcode -run TestDigitCounts
func TestDigitCounts(t *testing.T) {
	if x := digitCounts(0, 0); x != 1 {
		t.Log(x)
		t.Fatal("digitCounts(0, 0)!= 1")
	}
	if x := digitCounts(0, 19); x != 2 {
		t.Log(x)
		t.Fatal("digitCounts(0, 19) != 2")
	}
	if x := digitCounts(1, 12); x != 5 {
		t.Log(x)
		t.Fatal("digitCounts(1, 12) != 5")
	}
	if x := digitCounts(0, 9); x != 1 {
		t.Log(x)
		t.Fatal("digitCounts(0, 9) != 1")
	}
}
