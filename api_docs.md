#updoots™ API information

I will try to outline the basic way you should interact with the updoots™ backend.

The auth for this app is built on JWTs, but there is no login username/password.

Step 1 is to obtain a token by making the following request:

| HTTP  method | URL           | PAYLOAD |
|--------------|---------------|---------|
| POST         | /api/v1/token | { "lat": latitude(num), "lon", longitude(num), "userid": "phone_num"(string) }|

the server in one of the following ways:

1. `{"message": "Not within geofence", "reason": "invalid_location"}`
	
	if you are not wihtin the geofences and or you did not provide lat lon information.

2. `{"token": "v long signed jwt"}`

	if you are within the geofences and your userid has been registered
	
__NOTE:__ For the rest of the API calls you will make you have to make sure that the jwt is included in the header under "_Authorization_"


So now you are all authenticated, hooray! Time to get some posts and updoot things

Step 2 is to get some messages, right now there is no way to implement infinite scrolling i.e. this just gives you all of the active posts, this will change in the future

| HTTP  method | URL             | PAYLOAD |
|--------------|-----------------|---------|
| GET          | /api/v1/message | NONE    |

If you are authenticated properly the server will respond with the following object:

```
[ { "message":"dae drink beer?"(string), "id": 1234(num), "updoots": 237(num) }, 
{ "message":"this is the best app ever", "id": 6732(num), "updoots":-3(num) } ]
```
__NOTE:__ These will be returned in descending order for you, let the database do the heavy lifting!


Step 3 is to actually doot people!

| HTTP  method | URL                           | PAYLOAD |
|--------------|-------------------------------|---------|
| GET          | /api/v1/message/{id}/updoot   | NONE    |
| GET          | /api/v1/message/{id}/downdoot | NONE    |

If you are properly authenticated then the server should respond in the following manner to both requests:

`{"message":"success"}`

if you have already voted and are trying to vote again the server will respond:

`{"message":"success"}`

__NOTE:__ for non admin users this will delete their updoot from the total, admins will receive the same message but they will get an infinite number of updoots.

WOOOOOOOHOOOO get out there and doot™