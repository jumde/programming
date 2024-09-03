def fib(count):
    arr = [0, 1]
    for i in range(count):
        next = arr[-1] + arr[-2]
        arr.append(next)

    return arr

print(fib(5))
