class Solution:
    def hasDuplicate(self, nums: List[int]) -> bool:
        seen = set()
        for num in nums:
            seen[num] += 1
        
        for val in seen.values:
            if val > 1:
                return True
        return false
        
