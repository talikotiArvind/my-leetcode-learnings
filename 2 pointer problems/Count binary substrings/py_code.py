def countBinarySubstrings(s: str) -> int:
  prev, cur = 0, 1
      ans = 0
      for i in range(1, len(s)):
          if s[i] == s[i - 1]:
              cur += 1
          else:
              ans += min(prev, cur)  # close out the boundary between prev run and cur run
              prev = cur
              cur = 1
      ans += min(prev, cur)  # don't forget the final boundary
      return ans
