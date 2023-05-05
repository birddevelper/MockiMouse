# MockiMouse


<p align="center">
<img src="https://raw.githubusercontent.com/birddevelper/MockiMouse/master/mockimouse_icon.png"  height="200" >
</p>


Develop your UI without any concern about the backend. MockiMouse is a mock server that helps you make dynamic fake API to test or demo your frontend project.

## How to use
MockiMouse is easy to use, and easy to run. In a few lines of YAML config file you can start serving requests from frontend. Let's start :

The below config is the simplest possible mock server to run. Two endpoints with single senarios without any conditional response. 

To learn more details read this article : [Mockimouse Mock server](https://mshaeri.com/blog/mockimouse-an-easy-to-use-mock-server-to-build-fake-dynamic-api/)

```yaml
MockServer :
 port : 800
 endpoints :
  - name : My first endpoint
    path : /helloWorld
    method : GET
    scenarios :
     - description : no condition, always show same response
       response: 
        - Welcome to Hello wrold
  - name : My second endpoint
    path : /goodbye
    method : GET
    scenarios :
     - description : no condition, always show same goodbye
       response: 
        - goodbye
```

Add unlimited scenarios for each endpoint and set multiple conditions for a scenario to trigger. For example for a login endpoint you can set two scenarios first for valid username and password and another scenario for invalid username and password :

```yaml

MockServer :
 contextPath : /api
 port : 800
 endpoints :
  - name : Login API
    path : /login
    accepts : application/json
    method : POST
    delay : 1000
    scenarios :
     - description : When credintial is valid
       condition :
         param :
            - name : username
              type : body
              operand : equal
              value : admin
            - name : password
              type : body
              operand : equal
              value : 1234
       response: 
         - file://helloWorld.json
       
     - description : When credintial is invalid
       condition :
          param :
            - name : username
              type : body
              operand : equal
              value : admin
            - name : password
              type : body
              operand : notEqual
              value : 1234
       response : 
         - file://invalidCredintial.json
       status : 200

```
Put your message file in **responses** folder beside the MockiMouse binary file and call them in response parameters in config file. The file can be json, xml or html.
## How to run

Run the server binary in any operating system and enjoy it :

Win OS:

```bash
c:\myFakeServer\mockimouse.exe 
```




<p align="center">
<img src="https://mshaeri.com/blog/wp-content/uploads/2023/01/mockimouse_mock_server_fake_api.jpg"  >
</p>
