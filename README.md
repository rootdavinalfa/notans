
# Notans 
  
Link shortener written in Go. [**Experimental**]

## Build
You can build this project using make,this can be done with

    make build
 
After that, kindly check dist folder and run the executable from that folder.
  
## How to use
Because this application still in experimental, we can't provide the detailed documentation to use this application. But we can provide some how-to:

1. Disable middleware by commenting backend/routes/routes.go service.use(middle.AuthMiddle())
2. Create account on http://localhost:3000/service/user [POST]  you can refer the body with

     {  
      "Username": "davin",  
      "Password": "davin"  
    }

3. You can enable again the middleware to testing your account is able to login.
4. After login, you can obtain the token generated from system. Use that and insert to your http header on service/** endpoint with `Authorization: token`
5. You can create a new shortener link with endpoint http://localhost:3000/service/link [POST], just fill the request body with this example

    {  
      "OLink" : "https://github.com/"  
    }
 6. If success, the response will contain SLink, grab that and you can use the shortener with http://localhost:3000/slink/{SLink}'
 7. The request will be redirected to your original link (OLink) you has been put before

## Copyright
(C)2022, Davin Alfarizky Putra Basudewa
