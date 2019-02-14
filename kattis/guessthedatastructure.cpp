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
		int a,b;
		int a1,b1,c1;
		a1=true;
		b1=true;
		c1=true;
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
				if(s.top()==b and a1==true)
					a1=true;
				else
					a1=false;
				if(q.front()==b and b1==true)
					b1=true;
				else
					b1=false;
				if(pq.top()==b and c1==true)
					c1=true;
				else
					c1=false;
			}
		}		
				cout<<a1<<b1<<c1;
				if(a1==true)
					if(b1==true or c1==true)
						cout<<"more";
					else
						cout<<"stack";
				else if(b1==true)
					if(b1==true or c1==true)
						cout<<"more";
					else
						cout<<"queue";
				else if(c1==true)
					if(b1==true or c1==true)
						cout<<"more";
					else
						cout<<"pq";
				else
					cout<<"imp";
	}
		
	return 0;
}
