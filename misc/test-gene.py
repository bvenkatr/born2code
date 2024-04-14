"""
class yrange:
    def __init__(self, n):
        self.i = 0
        self.n = n

    def __iter__(self):
        return self

    def next(self):
        if self.i < self.n:
            i = self.i
            self.i += 1
            return i
        else:
            raise StopIteration()
"""

def yrange(n):
    i = 0
    while i < n:
        yield i
        i += 1

res = yrange(10)
print(res.next())
print(res.next())
print(res.next())
print(res.next())

