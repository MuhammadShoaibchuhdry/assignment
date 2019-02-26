# Example of user service 
1. This is simple example of user api basic crud operation(UserSignUp, UserSignIn, DeleteUser, UpdateUser) with jwt authentication.
2. Mongodb used for database.
3. Govendor used for dependencies managment.

# To run just clone the repo and run ./start.sh in project directory for linux and make sure have mongodb running in same machine.  

For SignUp The payload will be:
        
    {
        "fullname": "test",
        "email": "test@gmail.com",
        "password": "1234"
    }

For Signin The payload will be:

    {
        "email": "test@gmail.com",
        "password": "1234"
    }
It will return json data of name, email and jwt token.

To Update The payload will be: 
    
    {
        "fullname": "update name",
        "phonenumber": "123456"
    }
  You need send jwt token in header other wise will return nill.

  For Delete user just need to call delete api of user service with jwt token in header.

  For more detaill of route can check user/user_routes.md


