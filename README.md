blockfind
=========

대용량 로그파일에서 특정 문자열이 들어있는 블록을 찾아주는 프로그램.

Install :
-------------

* install golang latest. (https://golang.org/dl/)
* download src (git clone https://github.com/jaksal/blockfind.git)
* goto src folder 
* go get && go build


Usage 
-----
```

dumpfind -sp="/\n" -path=./log -find="@[139558]"

//  -sp : 블록구분자. ( \n , \r\n , \t 포함해서 사용가능 )
//  -path : 검색할 디렉토리 ( 하위폴더 포함 )
//  -find : 검색할 문자열.
```


