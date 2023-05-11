from fractions import Fraction
import functools

@functools.cache
def get_prob(w, b):
    if w == 0 or b < 0:
        return Fraction(0, 1)
    if b == 0:
        return Fraction(1, 1)
    ans = Fraction(w, w+b)
    if b == 1:
        return ans
    ans += Fraction(b, w+b)*Fraction(b-1, w+b-1)*Fraction(b-2, w+b-2)*get_prob(w, b-3)
    ans += Fraction(b, w+b)*Fraction(b-1, w+b-1)*Fraction(w, w+b-2)*get_prob(w-1, b-2)
    return ans
w, b = tuple(map(int, input().split()))
#ans = get_prob(w, b)
dp = [[0 for i in range(w+1)] for j in range(b+1)]
for bi in range(3):
    for wi in range(w+1):
        dp[bi][wi] = get_prob(wi, bi)
for wi in range(2):
    for bi in range(b+1):
        dp[bi][wi] = get_prob(wi, bi)
for bi in range(3, b+1):
    for wi in range(2, w+1):
        dp[bi][wi] = Fraction(wi, wi+bi)
        dp[bi][wi] += Fraction(bi, wi+bi)*Fraction(bi-1, wi+bi-1)*Fraction(bi-2, wi+bi-2) * dp[bi-3][wi]
        dp[bi][wi] += Fraction(bi, wi+bi)*Fraction(bi-1, wi+bi-1)*Fraction(wi, wi+bi-2) * dp[bi-2][wi-1]
        
ans = dp[b][w]
print(ans.numerator/ans.denominator)
# for bi in range(b+1):
#     print([str(x) for x in dp[bi]])
print(ans)