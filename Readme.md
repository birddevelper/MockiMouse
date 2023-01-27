# MockiMouse


<p align="center">
<img src="https://m-shaeri.ir/blog/wp-content/uploads/2023/01/mockimouse_mockserver.png"  height="200" >
</p>


Develop your UI without any concern about the backend. MockiMouse is a mock server that helps you make fake API to test or demo your frontend project.

## How to use
MockiMouse is easy to use, and easy to run. In a few lines of Yamel config file you can start serving requests from frontend. Let's start :

The below config is the simplest possible mock server to run. Two endpoints with single senarios without any conditional response. 

```yaml
MockServer :
 contextPath : /api
 port : 800
 endpoints :
  - name : My first endpoint
    path : /helloWorld
    method : GET
    scenarios :
     - description : no condition, always show same response
       response: Welcome to Hello wrold
  - name : My second endpoint
    path : /goodbye
    method : GET
    scenarios :
     - description : no condition, always show same goodbye
       response: goodbye
```