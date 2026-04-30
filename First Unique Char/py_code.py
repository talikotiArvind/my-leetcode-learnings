class Solution:
    def firstUniqChar(self, s: str) -> int:
        data = {char:s.count(char) for char in set(s)}
        for index, value in enumerate(s):
            if data[value] == 1:
                return index
        return -1
        
