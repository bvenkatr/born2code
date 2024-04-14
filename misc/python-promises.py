from promise import Promise
from threading import Timer

def twoArgs(resolve):
        return resolve("Resolved from twoArgs")

promise = Promise(
    lambda resolve, reject: t = Timer(5.0, twoArgs, (resolver) t.start()) 
)
promise.then(lambda data: print(data))
print("I am ")

