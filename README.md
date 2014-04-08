# feather label [![wercker status](https://app.wercker.com/status/5e2110ec6d03698e547ee1616befb7c9/s/ "wercker status")](https://app.wercker.com/project/bykey/5e2110ec6d03698e547ee1616befb7c9)

Support your favorite indie artists with your favorite indie currency via this all-new music distribution platform.

Web Client: [featherlabel.com](http://featherlabel.com)

Api Endpoint: [api.featherlabel.com](http://api.featherlabel.com)

## API

+ [Songs](#songs)
+ [Users](#users)


### Songs
##### Index 
cURL Example
```
curl -X GET http://api.featherlabel.com/song/
```
Response
```
[
  {
    "id": 9,
    "title": "Tom Sawyer",
    "artist": {
      "id": 1,
      "name": "Rush"
    }
  },
  {
    "id": 3,
    "title": "The Violent Bear It Away",
    "artist": {
      "id": 2,
      "name": Moby"
    }
  }
]
```

##### Retrieve 
cURL Example
```
curl -X GET http://api.featherlabel.com/song/3
```
Response
```
{
  "id": 3,
  "title": "The Violent Bear It Away",
  "artist": {
    "id": 1,
    "name": "Moby"
  },
  "url": "https://s3-us-west-2.amazonaws.com/media.jobgenius/songs/away.mp3?AWSAccessKeyId=AKIAJ2EASPWDMK6FOILA\u0026Expires=1396900314\u0026Signature=E8yr6E%2BMwwoeApW3%2FHhb2idENPA%3D",

}
```
Note: The authenticated url is not valid on its own and must be parsed.

##### Upload
cURL Example
```
```
Response
```
```

### Users
##### Index 
cURL Example
```
```
Response
```
```
