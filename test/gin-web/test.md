
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
  

  curl -v -X POST \
  http://localhost:8877/hippy_wns_buffer_log \
  -H 'content-type: application/json' \
  -d '{"ilogList":[{ "requestInfo": "milulululu" },{ "requestInfo": "222222" }]}'

    curl -v -X POST \
  http://localhost:8877/hippy_wns_buffer_log \
  -H 'content-type: application/json' \
  -d '{ "requestInfo": "33333333" }'

  curl -v -X POST \
  http://ilogs2.oa.com/ilogs_backend/post_data \
  -H 'content-type: application/json' \
  -d ’{"logid":"kg_hippy_speed_detection","docs":[{"Uin":0,"TimeStamp":"2020-12-08 13:13:34","ProjectName":"projectName_test_value","Platform":"platform_test_value","AppVersion":"app_version_test_value","JsbundleVersion":"jsbundle_version_test_value","Qua":"qua_test_value","Model":"model_test_value","NetworkType":"network_type_test_value","DownloadTime":0,"EngineTime":0,"LoadBundleTime":0,"CreateViewTime":0,"LoadJsTime":0,"FirstFrameTime":0,"FirstScreenTime":0,"ShowDataTime":0,"IsAssetFile":0,"LoadTotalTime":0,"ExtInfo":"{}"},{"Uin":0,"TimeStamp":"2020-12-08 13:13:34","ProjectName":"projectName_test_value","Platform":"platform_test_value","AppVersion":"app_version_test_value","JsbundleVersion":"jsbundle_version_test_value","Qua":"qua_test_value","Model":"model_test_value","NetworkType":"network_type_test_value","DownloadTime":0,"EngineTime":0,"LoadBundleTime":0,"CreateViewTime":0,"LoadJsTime":0,"FirstFrameTime":0,"FirstScreenTime":0,"ShowDataTime":0,"IsAssetFile":0,"LoadTotalTime":0,"ExtInfo":"{}"},{"Uin":0,"TimeStamp":"2020-12-08 13:13:34","ProjectName":"projectName_test_value","Platform":"platform_test_value","AppVersion":"app_version_test_value","JsbundleVersion":"jsbundle_version_test_value","Qua":"qua_test_value","Model":"model_test_value","NetworkType":"network_type_test_value","DownloadTime":0,"EngineTime":0,"LoadBundleTime":0,"CreateViewTime":0,"LoadJsTime":0,"FirstFrameTime":0,"FirstScreenTime":0,"ShowDataTime":0,"IsAssetFile":0,"LoadTotalTime":0,"ExtInfo":"{}"},{"Uin":0,"TimeStamp":"2020-12-08 13:13:34","ProjectName":"projectName_test_value","Platform":"platform_test_value","AppVersion":"app_version_test_value","JsbundleVersion":"jsbundle_version_test_value","Qua":"qua_test_value","Model":"model_test_value","NetworkType":"network_type_test_value","DownloadTime":0,"EngineTime":0,"LoadBundleTime":0,"CreateViewTime":0,"LoadJsTime":0,"FirstFrameTime":0,"FirstScreenTime":0,"ShowDataTime":0,"IsAssetFile":0,"LoadTotalTime":0,"ExtInfo":"{}"}]}'