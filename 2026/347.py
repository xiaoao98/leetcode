class Solution:
    def topKFrequent(self, nums: list[int], k: int) -> list[int]:
        dicFreq = {}
        for num in nums:
            if num not in dicFreq:
                dicFreq[num] = 1
            else:
                dicFreq[num] += 1
        s = sorted(dicFreq.items(), key=lambda x: x[1], reverse=True)
        ret = []
        for i in range(k):
            ret.append(s[i][0])
        return ret