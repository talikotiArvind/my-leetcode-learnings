x = -121

def is_palindrome(num):
    """
    Docstring for is_palindrome
    
    :param num: Description
    """
    # Time complexity: O(n)
    # Space complexity: O(n)
    str_num = str(num)
    return str_num == str_num[::-1]


def is_palindrome_optimized(num):
    """
    Docstring for is_palindrome_optimized
    
    :param num: Description
    """
    # Time complexity: O(log(n))
    # Space complexity: O(1)
    if num< 0:
        return False
    
    original, reversed_num = num, 0
    while num != 0:
        digit = num % 10
        reversed_num = reversed_num * 10 + digit
        num //= 10
    
    return original == reversed_num