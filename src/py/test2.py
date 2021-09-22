a = 3
def x(x):
    print(x)
    global a
    a = 4
x(a)
print(a)