class LRUCache:
    
    def __init__(self, capacity: int):
        self.cap = capacity
        self.nodeNum = 0
        self.doubleLink = DoubleLink(Node(), Node())
        self.dic = {}

    def get(self, key: int) -> int:
        if key not in self.dic:
            return -1
        else:
            self.doubleLink.moveTohead(self.dic[key])
            return self.dic[key].val

    def put(self, key: int, value: int) -> None:
        if key in self.dic:
            self.doubleLink.moveTohead(self.dic[key])
            self.dic[key].val = value
        else:
            newNode = Node(key, value)
            self.doubleLink.addNode(newNode)
            self.dic[key] = newNode
            self.nodeNum += 1
            if self.nodeNum > self.cap:
                toDeleteNode = self.doubleLink.tail.prev
                toDeleteNodeKey = toDeleteNode.key
                del self.dic[toDeleteNodeKey]
                self.doubleLink.delNode(toDeleteNode)
                self.nodeNum -= 1

class Node:
    def __init__(self, key=0, val=0):
        self.key = key
        self.val = val
        self.next = None
        self.prev = None

class DoubleLink:
    def __init__(self, head: Node, tail: Node):
        self.head = head
        self.tail = tail
        self.head.next = self.tail
        self.tail.prev = self.head
    
    def addNode(self, node: Node):
        node.next = self.head.next
        self.head.next.prev = node
        self.head.next = node
        node.prev = self.head
    
    def delNode(self, node: Node):
        node.next.prev = node.prev
        node.prev.next = node.next

    def moveTohead(self, node: Node):
        self.delNode(node)
        self.addNode(node)