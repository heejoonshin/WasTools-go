#설정값을 Yaml로 읽어 오는 패키지


1. app port정보와 DB정보를 설정하는 yaml을 읽어온다.

2. 인증 정보를 읽어오는 정보를 yaml로 읽어온다.

    사용방법
    
    Oauth2폴더아래
    
    oauth2.yaml파일을 생성한후
    ~~~
    kakao:
      client:
        clientid : clientId
        endpoint :
          authurl : https://kauth.kakao.com/oauth/token
          tokenurl : https://kauth.kakao.com/oauth/authorize
        redirecturl : 개인이 설정한 redirecturl
      resource:
        sign : https://kapi.kakao.com/v1/user/signup
        userinfo : https://kapi.kakao.com/v2/user/me
    ~~~
    위와같은 방식으로 생성
    현재 kakao인증만 지원