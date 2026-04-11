def isIsomorphic(s: str, t: str) -> bool:
        data_s, data_t = {}, {}
        for i in range(len(s)):
            if s[i] not in data_s:
                data_s[s[i]] = i
            if t[i] not in data_t:
                data_t[t[i]] = i

            if data_s[s[i]] != data_t[t[i]]:
                return False
        return True
    
