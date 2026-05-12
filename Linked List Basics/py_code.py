
class Node:
    def __init__(self, data):
        self.data = data
        self.next = None
        
class LinkedList:
    def __init__(self):
        self.head = None
        
    #Insertion
    def append(self, data):
        new_node = Node(data)
        if not self.head:
            self.head = new_node
            return
        
        last_node = self.head
        while last_node.next:
            last_node = last_node.next
        last_node.next = new_node
        
    
    # Deletion
    def delete(self, value):
        current = self.head
        if current and current.data == value:
            self.head = current.next
            return
        previous = None
        while current and current.data != value:
            previous = current
            current = current.next
        
        if current:
            previous.next = current.next
            
    # Search
    def find(self, data):
        current = self.head
        while current:
            if current.data == data:
                return True
            current = current.next
        return False

    def print_list(self):
        current = self.head
        while current:
            print(current.data, end= " -> ")
            current = current.next
        print("None")
