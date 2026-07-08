from collections import deque
from typing import List


class Employee:
    def __init__(self, id: int, importance: int, subordinates: List[int]):
        self.id = id
        self.importance = importance
        self.subordinates = subordinates


def getImportance(employees: List[Employee], id: int) -> int:
    emp_map = {emp.id: emp for emp in employees}
    total = 0
    queue = deque([id])
    while queue:
        curr_id = queue.popleft()
        emp = emp_map[curr_id]
        total += emp.importance
        queue.extend(emp.subordinates)
    return total


if __name__ == "__main__":
    employees = [Employee(1, 5, [2, 3]), Employee(2, 3, []), Employee(3, 3, [])]
    print(getImportance(employees, 1))  # 11
