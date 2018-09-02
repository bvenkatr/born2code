import time
from functools import wraps

def timethis(func):
	'''
	decorator that reports the execution time
	'''

	@wraps(func)
	def wrapper(*args, **kwargs):
		start = time.time()
		result = func(*args, **kwargs)
		end = time.time()
		print(func.__name__, end - start)
	return wrapper

@timethis
def countdown(n=10002222):
	'''
	Counts Down
	'''
	while n > 0:
		n -= 1

countdown(10)
