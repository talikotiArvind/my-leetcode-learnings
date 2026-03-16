strs = ["dog","racecar","car"]
strs = ["flower","flow","flight"]
def longestCommonPrefix(strs):
    output = ""
    for i, chars in enumerate(zip(*strs)):
        if len(set(chars)) > 1:
            break
        output += chars[0]
    return output

print(longestCommonPrefix(strs))