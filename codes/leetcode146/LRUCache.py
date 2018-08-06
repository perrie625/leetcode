class Node(object):
    def __init__(self, key, prev, next):
        self.key = key
        self.prev = prev
        self.next = next

class LRUCache(object):

    def __init__(self, capacity):
        """
        :type capacity: int
        """


        self.loop = Node("", None, None)
        self.loop.prev = self.loop
        self.loop.next = self.loop
        self.nodes = dict()
        self.capacity = capacity
        self.values = dict()
        

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        if key in self.values:
            node = self.nodes[key]
            self.delete_node(node)
            self.insert_node(node)
        return self.values.get(key, -1)
        
        

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: void
        """
        exists = key in self.values
        self.values[key] = value
        if exists:
            # 节点提前
            node = self.nodes[key]
            self.delete_node(node)
            self.insert_node(node)
            return
        
        node = Node(key, None, None)
        self.insert_node(node)
        self.nodes[key] = node
    
        while len(self.values) > self.capacity:
            self.diout_node()

    def delete_node(self, node):
        node.prev.next = node.next
        node.next.prev = node.prev
        node.next = None
        node.prev = None
        del self.nodes[node.key]

    def diout_node(self):
        node = self.loop.prev
        self.delete_node(node)
        del self.values[node.key]

    def insert_node(self, node):
        self.loop.next.prev = node
        node.next = self.loop.next
        self.loop.next = node
        node.prev = self.loop

        self.nodes[node.key] = node