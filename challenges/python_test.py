from my_functions import g

def f(x, y):
    return g(x + y)
def main():
    x=6
    y=7.5
    result=x+y
    print(g(result))
    print(result)
main()