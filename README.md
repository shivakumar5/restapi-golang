# RESTful JSON API WITH GOLANG
This repo contains RESTful JSON API with Golang


Used Gorilla Mux package to implement a request router and dispatcher for matching incoming requests to their respective handler.

Import it with below URL:

```
github.com/gorilla/mux
```
### This API Contains 5 different endpoints:

- CreateEmployee
- GetAllEmployees
- GetAllEmployees/{id}
- DeleteEmployee/{id}
- UpdateEmployee/{id}

Note that, added {id} to endpoint path. This will represent id variable that we will be able to use when we wish to return only the employee record that features that exact key.

- **Employee Strcut**
```
// Employee ...
type Employee struct{
	ID string `json:"id"`
	Name string `json :"name"`
	Age int `json :"age"`
}
```

- **Creating New Employee :**  
The endpoint for getting single employee record is `/CreateEmployee`. When creating an employee, we get data from the user’s end. The user enters data which is in the form of http Request data. The request data is not is a human-readable format hence we use the package ioutil to convert it into a slice.
 Once it is converted to a slice, we unmarshal it to fit into our `employee` struct. Once it is successful, we append the new employee struct to the `Employees` slice.

- **Get All Employees :**  
The endpoint for getting single employee record is `/GetAllEmployees`. To get all the employees, we simply display the whole events slice. 

- **Get One Employee :** 
The endpoint for getting single employee record is `/GetAllEmployees/{id}` and uses a GET Method. Using Gorilla mux, we get the value of the “id” and use it to filter out a specific employee from the Employees slice. When the “id” resembles the one of an employee in the slice, we get its value and display it as a response to the user.

- **Update Employee :**  
The endpoint for updating an employee is `/UpdateEmployee/{id}` and uses a PUT Method. Using Gorilla mux, we get the value of the “id” and use it to filter out a specific employee from the events slice. When the “id” resembles the one of an employee in the slice, we get to alter the values of the Name and Age fields in the employee struct.
After that, we update the value of the struct in the Employees slice.

- **Delete Employee :**  
The endpoint for deleting an employee is `/DeleteEmployee/{id}` and uses a DELETE Method. Using Gorilla mux, we get the value of the “id” and use it to filter out a specific employee from the Employees slice. When the “id” resembles the one of an employee in the slice, we remove the employee in the slice.



