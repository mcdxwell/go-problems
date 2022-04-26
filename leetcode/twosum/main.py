class Solution:
    def twoSum(self, nums: list[int], target: int) -> list[int]:

        my_map = {}

        for i, num in enumerate(nums):

            if num in my_map:
                return [i, my_map[num]]

            my_map[target-num] = i


my_list = [5, 2, 1, -3, 3, 5, 10]
target = 8

x = Solution()
x.twoSum(my_list, target)

print(x.twoSum(my_list, target))


def fib(n):
    a, b = 0, 1

    for _ in range(n):
        a, b = b, a+b

    return a


print(fib(10))


