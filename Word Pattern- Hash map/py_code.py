def wordPattern(pattern: str, s: str) -> bool:
      keys = list(pattern)
      values = s.split(" ")
      
      if len(keys) != len(values):
          return False
      
      output = {}
      used_values = set()  # track already-mapped words
      
      for i in range(len(pattern)):
          if keys[i] in output:
              if output[keys[i]] != values[i]:
                  return False
          else:
              if values[i] in used_values:  # word already used by another key
                  return False
              output[keys[i]] = values[i]
              used_values.add(values[i])
      
      return True
