import pandas as pd
print(pd.datetime.now().strftime("%Y-%m-%d"))

def isiterable(obj):
    try:
        iter(obj)
        return True
    except TypeError:
        return False

print(isiterable)
