"""
Tree data structure as Datamodel for application
"""
from abc import abstractmethod
from enum import Enum
from typing import Any, Optional

from pydantic import BaseModel


class DefaultReads(Enum):
    IN_ORDER = 'InOrder'
    PRE_ORDER = 'PreOrder'
    POST_ORDER = 'PostOrder'


class Node(BaseModel):
    model_config = {"arbitrary_types_allowed": True}

    value: Any
    left: Optional['Node'] = None
    right: Optional['Node'] = None


Node.model_rebuild()


class ABCTreeModel(BaseModel):
    model_config = {"arbitrary_types_allowed": True}
    root: Optional['Node'] = None


# ABCTree drops ABC inheritance; abstract methods still work via NotImplementedError
class ABCTree(ABCTreeModel):

    @abstractmethod
    def insert(self, value: Any) -> None:
        ...

    @abstractmethod
    def search(self, value: Any) -> bool:
        ...

    @abstractmethod
    def delete(self, value: Any) -> None:
        ...

    def update(self, old_value: Any, new_value: Any) -> bool:
        if self.search(old_value):
            self.delete(old_value)
            self.insert(new_value)
            return True
        return False

    def inorder(self, node: Optional[Node], result: list = None) -> list:
        if result is None:
            result = []
        if node:
            self.inorder(node.left, result)
            result.append(node.value)
            self.inorder(node.right, result)
        return result

    def preorder(self, node: Optional[Node], result: list = None) -> list:
        if result is None:
            result = []
        if node:
            result.append(node.value)
            self.preorder(node.left, result)
            self.preorder(node.right, result)
        return result

    def postorder(self, node: Optional[Node], result: list = None) -> list:
        if result is None:
            result = []
        if node:
            self.postorder(node.left, result)
            self.postorder(node.right, result)
            result.append(node.value)
        return result


class BinarySearchTree(ABCTree):

    def insert(self, value: Any) -> None:
        new_node = Node(value=value)
        if self.root is None:
            self.root = new_node
            return
        current = self.root
        while True:
            if value < current.value:
                if current.left is None:
                    current.left = new_node
                    break
                current = current.left
            else:
                if current.right is None:
                    current.right = new_node
                    break
                current = current.right

    def search(self, value: Any) -> bool:
        current = self.root
        while current is not None:
            if value == current.value:
                return True
            elif value < current.value:
                current = current.left
            else:
                current = current.right
        return False

    def delete(self, value: Any) -> None:
        self.root = self._delete_recursive(self.root, value)

    def _delete_recursive(self, node: Optional[Node], value: Any) -> Optional[Node]:
        if node is None:
            return None
        if value < node.value:
            node.left = self._delete_recursive(node.left, value)
        elif value > node.value:
            node.right = self._delete_recursive(node.right, value)
        else:
            if node.left is None and node.right is None:
                return None
            if node.left is None:
                return node.right
            if node.right is None:
                return node.left
            successor = self._find_min(node.right)

            node = node.model_copy(update={"value": successor.value})
            node.right = self._delete_recursive(node.right, successor.value)
        return node

    def _find_min(self, node: Node) -> Node:
        while node.left:
            node = node.left
        return node

    def read(self, node: Optional[Node], order: DefaultReads) -> list:
        match order:
            case DefaultReads.IN_ORDER:
                return self.inorder(node)
            case DefaultReads.PRE_ORDER:
                return self.preorder(node)
            case DefaultReads.POST_ORDER:
                return self.postorder(node)
            case _:
                raise ValueError(f"Unknown traversal order: {order}")
