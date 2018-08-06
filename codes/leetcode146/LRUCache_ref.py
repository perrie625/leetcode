class LRUCache(object):
    
    def __init__(self, capacity):
        """
        :type capacity: int
        """
        self.dic = {}
        from collections import deque
        self.q = deque()
        self.capacity = capacity
        self.count = 0

    def get(self, key):
        """
        :type key: int
        :rtype: int
        """
        #print("get: {}".format(self.q))
        if key in self.dic:
            self.dic[key][2] = 0
            newelem = [key, self.dic[key][1], 1]
            self.dic[key] = newelem
            self.q.append(newelem)
            return newelem[1]
        else:
            return -1

    def put(self, key, value):
        """
        :type key: int
        :type value: int
        :rtype: void
        """
        if key in self.dic:
            self.dic[key][2] = 0
            newelem = [key, value, 1]
            self.dic[key] = newelem
            self.q.append(newelem)
            return key
        if self.count == self.capacity:
            self.count-=1
            k, v, state = self.q.popleft()
            while state == 0 and self.q:
                k, v, state = self.q.popleft()
            if state == 0:
                return -1
            self.dic.pop(k)
        newelem = [key, value, 1]
        self.dic[key] = newelem
        self.q.append(newelem)
        self.count+=1