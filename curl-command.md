```
$ curl http://localhost:8000/noauth
Not Authorized!

$ curl http://localhost:8000/auth/info
{"message":"missing key in request header"}

$ curl http://localhost:8000/auth/info -H "Authorization:Bearer 208d97ff-0ab4-49b7-8eab-985ff27a11dc"

$ curl http://localhost:8000/auth/output -H "Authorization:Bearer 208d97ff-0ab4-49b7-8eab-985ff27a11dc" -o result.xlsx
```
