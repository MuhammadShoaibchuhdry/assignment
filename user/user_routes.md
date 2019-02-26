# assignment/user

User API.

## Routes

<details>
<summary>`/delete`</summary>

- [assignment/user/vendor/github.com/go-chi/cors.(*Cors).Handler-fm](/app/routes/routes.go#L27)
- **/delete**
	- _DELETE_
		- [Verifier.func1](/vendor/github.com/go-chi/jwtauth/jwtauth.go#L71)
		- [Authenticator](/vendor/github.com/go-chi/jwtauth/jwtauth.go#L162)
		- [assignment/user/app/repository.User.DeleteUserHandler-fm](/app/routes/routes.go#L38)

</details>
<details>
<summary>`/login`</summary>

- [assignment/user/vendor/github.com/go-chi/cors.(*Cors).Handler-fm](/app/routes/routes.go#L27)
- **/login**
	- _POST_
		- [assignment/user/app/repository.User.GetUserHandler-fm](/app/routes/routes.go#L31)

</details>
<details>
<summary>`/register`</summary>

- [assignment/user/vendor/github.com/go-chi/cors.(*Cors).Handler-fm](/app/routes/routes.go#L27)
- **/register**
	- _POST_
		- [assignment/user/app/repository.User.CreateUserHandler-fm](/app/routes/routes.go#L30)

</details>
<details>
<summary>`/updateuserprofile`</summary>

- [assignment/user/vendor/github.com/go-chi/cors.(*Cors).Handler-fm](/app/routes/routes.go#L27)
- **/updateuserprofile**
	- _PUT_
		- [Verifier.func1](/vendor/github.com/go-chi/jwtauth/jwtauth.go#L71)
		- [Authenticator](/vendor/github.com/go-chi/jwtauth/jwtauth.go#L162)
		- [assignment/user/app/repository.User.UpdateUserHandler-fm](/app/routes/routes.go#L37)

</details>

Total # of routes: 4
