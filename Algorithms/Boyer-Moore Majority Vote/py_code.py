def majority_element(nums):
    """Return the element appearing > n/2 times (assumes such an element exists)."""
    candidate, count = None, 0
    for x in nums:
        if count == 0:
            candidate = x
        count += 1 if x == candidate else -1
    return candidate


def main():
    votes = [2, 2, 1, 1, 1, 2, 2]
    print(f"array    : {votes}")
    print(f"majority : {majority_element(votes)}")


if __name__ == "__main__":
    main()
