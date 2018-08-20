'''
Progaram to calculate and print the longest subsequence of two strings
and its length
'''
def lcs(a,b):
    m=len(a)
    n=len(b)
    
    l=[[None]*(n+1) for i in range(m+1)]
    
    for i in range(m+1):
        for j in range(n+1):
            if (i==0 or j==0):
                l[i][j]=0
            elif(a[i-1]==b[j-1]):
                l[i][j]=l[i-1][j-1]+1 
            else:
                l[i][j]=max(l[i-1][j],l[i][j-1])
                
    print(l[m][n])
    index = l[m][n]
    i,j=m,n
    s=""
    while(i>0 and j>0):
        if(a[i-1]==b[j-1]):
            s=a[i-1]+s
            i-=1
            j-=1
        elif(l[i-1][j]>l[i][j-1]):
            i=-1
        else:
            j-=1
    print(s)

X = "AGGTAB"
Y = "GXTXAYB"
print("Length of LCS is ", lcs(X, Y))
 