# -*- coding: utf-8 -*-

def recusion(n):
    if n<=1:
        return 1

    return recusion(n-1)*n
i=20
print(recusion(i))

