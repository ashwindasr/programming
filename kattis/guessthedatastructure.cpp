#include <iostream>
#include <queue>
#include <stack>

using namespace std;


int main()
{
	stack<int> s;
	queue<int> q;
	priority_queue<int, vector<int>> pq;
	
	for(int i=0; i<1; i++)
	{
		int t;
		cin>>t;
		cout<<t;
		int a,b;
		for(int j=0;j<t;j++)
		{
			cin>>a>>b;
			if(a==1)
			{
				s.push(b);
				q.push(b);
				pq.push(b);
			}
			else
			{
				bool a1,b1,c1;
				if(s.top()==b)
					a1=true;
				if(q.front()==b)
					b1=true;
				if(pq.top()==b)
					c1=true;
					
				if(a1==true and b1==false and c1==false)
					cout<<"stack";
				else if(a1==false and b1==true and c1==false)
					cout<<"queue";
				else if(a1==false and b1==false and c1==true)
					cout<<"p queue";
				else
					cout<<"imp";
			}
		}
	}
	return 0;
}
