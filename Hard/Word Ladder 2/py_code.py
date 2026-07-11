from collections import deque

class Solution:
    def findLadders(self, beginWord: str, endWord: str, wordList: list[str]) -> list[list[str]]:
        word_set = set(wordList)
        if endWord not in word_set:
            return []
            
        adj = {}
        steps = {beginWord: 1}
        queue = deque([beginWord])
        found = False
        
        while queue and not found:
            size = len(queue)
            for _ in range(size):
                curr = queue.popleft()
                if curr not in adj:
                    adj[curr] = []
                    
                for j in range(len(curr)):
                    for c in 'abcdefghijklmnopqrstuvwxyz':
                        next_word = curr[:j] + c + curr[j+1:]
                        if next_word in word_set:
                            if next_word not in steps or steps[next_word] == steps[curr] + 1:
                                adj[curr].append(next_word)
                                if next_word not in steps:
                                    steps[next_word] = steps[curr] + 1
                                    queue.append(next_word)
                            if next_word == endWord:
                                found = True
                                
        results = []
        if found:
            path = [beginWord]
            def dfs(word):
                if word == endWord:
                    results.append(list(path))
                    return
                for nxt in adj.get(word, []):
                    if steps.get(nxt) == steps[word] + 1:
                        path.append(nxt)
                        dfs(nxt)
                        path.pop()
            dfs(beginWord)
            
        return results
