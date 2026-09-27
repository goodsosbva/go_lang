# Go 학습 프로젝트

Go 문법을 작은 프로그램으로 확인하며 표준 라이브러리의 사용 방식과 코드 구성을 익히는 저장소입니다. 파일 처리와 동시성에서 시작해 HTTP 서버와 JSON 응답까지 다루고 있습니다.

## 현재 코드

### 파일 처리와 동시성

| 프로젝트 | 코드에서 다루는 내용 |
| --- | --- |
| `01-glob-file-listing` | `filepath.Glob` 결과를 슬라이스로 받아 파일 경로를 순회합니다. |
| `02-file-reading` | `os.Open`으로 파일을 열고 `bufio.Scanner`로 읽습니다. `defer`를 이용한 자원 정리도 다룹니다. |
| `03-sequential-word-search` | 검색 결과를 구조체로 표현하고, 파일 목록 수집과 파일별 검색을 함수로 나눕니다. |
| `04-concurrent-word-search` | 같은 파일 검색을 고루틴으로 실행하고 채널로 결과를 모읍니다. 순차 처리와 동시 처리의 구조 차이를 비교합니다. |

### HTTP 서버와 JSON

| 프로젝트 | 코드에서 다루는 내용 |
| --- | --- |
| `05-web-server/hello-server` | `net/http`의 핸들러 등록과 서버 시작을 다룹니다. |
| `05-web-server/have-query` | 요청 URL의 쿼리 값을 읽고 `strconv`으로 변환합니다. |
| `05-web-server/using_serve-mux` | `ServeMux`를 직접 만들고 경로별 핸들러를 등록합니다. |
| `05-web-server/file-server` | `http.FileServer`를 사용해 정적 파일을 제공합니다. |
| `05-web-server/https-server` | `ListenAndServeTLS`를 사용하는 예제입니다. 실행에는 인증서와 키 파일이 필요합니다. |
| `05-web-server/test-server` | 핸들러 구성을 `MakeWebHandler`로 분리하고, `httptest`로 서버를 실제로 띄우지 않고 응답을 검사합니다. |
| `05-web-server/json-server` | Go 구조체를 `encoding/json`으로 JSON 응답으로 변환하고, 테스트에서 응답을 다시 구조체로 읽습니다. |

이 예제들은 Go의 기본 패키지와 명시적인 함수·구조체·인터페이스를 사용해 요청 처리 흐름을 구성합니다. `http.Handler`를 서버 실행 코드와 분리하면 핸들러를 독립적으로 테스트할 수 있고, 고루틴과 채널을 사용하면 작업과 결과 전달을 코드에 드러낼 수 있습니다.

## 다음 단계

다음에는 JSON 요청을 읽고, HTTP 메서드와 상태 코드에 따라 생성·조회·수정·삭제를 처리하는 REST API를 만들 계획입니다. 이후 그 API를 사용하는 Todo 토이 프로젝트로 확장합니다.
