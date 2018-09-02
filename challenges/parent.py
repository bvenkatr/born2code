class Parent:
     def __init__(self, name):
             self.name = name
     def spam(self):
             print("Parent.spam", self.name)
     def grok(self):
             print("Parent.grok")

class Child(Parent):
	def yow(self):
		print("Child.yow()")

