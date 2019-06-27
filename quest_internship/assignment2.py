arr = []
while(True):
    ch = int(input("\n1. Add Employee\n2. Display all employees\n3. Display employee by ID\n4. Update employee by ID\n5. Delete employee by ID\n6. Exit\n\nEnter your choice: "))
    if(ch==1):
        temp = []
        print("\nEnter details of employee")
        print("------------------------------")
        temp.append(input("Enter Employee ID: "))
        temp.append(input("Enter Employee name: "))
        temp.append(input("Enter Employee designation: "))
        temp.append(input("Enter date of joining: "))
        
        arr.append(temp)
        
    elif(ch==2):
        if(arr == []):
            print("Empty database")
        else:
            c = 1
            for i in arr:
                print("\nEmployee {} details".format(c))
                print("-------------------")
                print("Employee ID:",i[0])
                print("Employee name:",i[1])
                print("Employee designation:",i[2])
                print("Date of joining:",i[3])
                c+=1
    elif(ch==3):
        id = input("Enter employee ID: ")
        flag = 1
        for i in arr:
            if(i[0]==id):
                print("Employee ID:",i[0])
                print("Employee name:",i[1])
                print("Employee designation:",i[2])
                print("Date of joining:",i[3])
                flag = 0
                break
        if(flag):
            print("No employee with the specified ID")
        
    elif(ch==4):
        id = input("Enter employee ID: ")
        flag = 1
        for i in arr:
            if(i[0]==id):
                choice = int(input("1. Name\n2. Designation\n3. Date of joining\n\nEnter your choice: "))
                if(choice==1):
                    i[1] = input("Enter new employee name: ")
                elif(choice==2):
                    i[2] = input("Enter new designation: ")
                elif(choice==3):
                    i[3] = input("Enter new date of joining: ")
                else:
                    print("Invalid choice")
                flag = 0
                break
        if(flag):
            print("No employee with the specified ID")
    elif(ch==5):
        id = input("Enter employee ID: ")
        flag = 1
        for i in range(len(arr)):
            if(arr[i][0]==id):
                del(arr[i])
                print("Deleted successfully")
                flag = 0
                
                break
        if(flag):
            print("No employee with the specified ID")
    elif(ch==6):
        print("Exited successfully")
        break
    else:
        print("Invalid choice")
