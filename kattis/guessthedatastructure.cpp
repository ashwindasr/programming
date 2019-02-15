#include <iostream>
#include <queue>
#include <stack>

using namespace std;


int main()
{
	stack<int> s;
	queue<int> q;
	priority_queue<int, vector<int> > pq;
	
	for(int i=0; i<1; i++)
	{
		int t;
		cin>>t;
		int a,b;
		int a1,b1,c1;
		a1=true;
		b1=true;
		c1=true;
		int flag=0;

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
				
				flag=1;

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

				s.pop();
				q.pop();
				pq.pop();
			
			}
		}		

				if(a1==true)
					if(b1==true or c1==true)
						cout<<"not sure";
					else
						cout<<"stack";
				else if(b1==true)
					if(a1==true or c1==true)
						cout<<"not sure";
					else
						cout<<"queue";
				else if(c1==true)
					if(b1==true or a1==true)
						cout<<"not sure";
					else
						cout<<"priority queue";
				else
					cout<<"impossible";
			
				if(flag==0)
					cout<<"impossible";
	}
		
	return 0;
}
