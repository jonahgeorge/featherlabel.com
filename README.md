# featherlabel.com [![wercker status](https://app.wercker.com/status/5e2110ec6d03698e547ee1616befb7c9/s/ "wercker status")](https://app.wercker.com/project/bykey/5e2110ec6d03698e547ee1616befb7c9)

Support your favorite indie artists with your favorite indie currency via this all-new music distribution platform.

## API
### Retrieve
cURL Example
```
curl -X GET https://api.featherlabel.com/v1/song/
```
Response
```
[
  {
    "id": 1,
    "title": "Tom Sawyer",
    "artist": "Rush"
  },
  {
    "id": 2,
    "title": "Get Lucky",
    "artist": "Daft Punk"
  },
  {
    "id": 3,
    "title": "The Violent Bear It Away",
    "artist": "Moby"
  }
]
```

