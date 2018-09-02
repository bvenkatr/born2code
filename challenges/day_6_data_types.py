#!/usr/local/bin/python

# Objects, strings, numbers, lists, dictionaries, tuples, and files


# files
f = open("/etc/passwd", "r")
print(f.readline()[:-1])
print(f.readline()[:-1])
f.close()
