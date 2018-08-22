# -*- coding: utf-8 -*-
'''
Program is to calculate the LOngest Increasing Subsequence (lis)

'''

def lis(arr):
    n = len(arr)
    l = [1]*n
    
    for i in range(1,n):
        for j in range(0,i):
            if((arr[i]>arr[j]) and (l[i]<l[j]+1)):
                l[i]=l[j]+1
    maximum=0           
    for i in l:
        maximum = max(maximum,i)
        
    return maximum

arr = [3, 10, 2, 1, 20]
print("Length of lis is", lis(arr))