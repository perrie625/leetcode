import functools

class Solution(object):
    def largestNumber(self, nums):
        """
        :type nums: List[int]
        :rtype: str
        """
        nums = [str(n) for n in nums]
        nums = sorted(nums, key=functools.cmp_to_key(self.compare))
        result = "".join(nums)
        if result[0] == "0":
            return "0"
        return result

    @staticmethod
    def compare(x, y):
        if x+y >= y+x:
            return 1
        return -1

