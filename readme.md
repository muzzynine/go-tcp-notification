# go-tcp-notification


## Overview

#### 시작하기

이 프로젝트는 go로 작성되었으며 웹서버의 클라이언트 페이지 빌드를 원한다면 npm이 있어야한다.

	go get github.com/muzzynune/go-tcp-notification

rpc 소스코드를 빌드하여 라이브러리로 사용할 수 있도록 한다.

	cd rpc
	go build

웹서버가 제공하는 클라이언트 페이지를 위해 의존성 설치한다.

	cd web
	npm install
	
client, tcp, web에 접근하여 config파일에 관련값을 설정한다.
클라이언트의 connId는 웹 페이지에 접근하여 디바이스를 등록 한 후 입력한다.

클라이언트에서 제공하는 웹 페이지에 대한 코드는 web/public 에 위치해있으며, React(javascript)를 사용하여 작성되어있다.
