def findContentChildren(g: List[int], s: List[int]) -> int:
        s.sort()
        g.sort()
        m,n = len(g), len(s)
        child, cookie = 0,0
        while child < m and cookie < n:
            if s[cookie] >= g[child]:
                child += 1
            cookie += 1
        return child
