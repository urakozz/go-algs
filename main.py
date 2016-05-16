

def primes_sieve(n):
    """ temporary odd direct sieve for primes < n """
    #sieve = list(range(3, n, 2))
    sieve = list(range(3, n, 2))
    l = len(sieve)
    for i in sieve:
        if i:
            f = (i*i-3) // 2
            if f >= l:
                break
            sieve[f::i] = [0] * -((f-l) // i)
    return sieve

T = input()
ans = []
for x in range(T):
    n = long(raw_input())
    sum = 2
    # print "To " + str(n)
    sieve = primes_sieve(n+1)
    for j in range(3, n+1, 2):
        if sieve[(j-2)//2]:
            # print j
            sum +=j

    # print "Sum " + str(sum)
    ans.append(sum)
for x in ans:
    print x