
curl -v -X POST \
  http://localhost:9000/loginTest \
  -H 'content-type: application/json' \
  -d '{ "username": "milulululu","age":14 }'
  
curl -v -X POST \
  http://localhost:9000/loginTest \
  -H 'content-type: application/json' \
  -d '{"listinfo":[{ "username": "milu","age":18 },{ "username": "milu22","age":19 }]}'
  

  curl -v -X GET \
  http://localhost:9000/loginTest?username=mmlluu&age=15
  