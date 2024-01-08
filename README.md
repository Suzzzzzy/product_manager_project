# 프로젝트명 
> product_manager_project

# 프로젝트 설명
상품을 등록 및 관리하여 가게를 운영하기 위한 백엔드 프로젝트 입니다.

## 개발
- 개발 언어: Golang:1.20
- web 프레임워크: [Gin](https://github.com/gin-gonic/gin)
- gORM
- docker-compose:3.8

## 폴더 구조
```bash
ad-server-project
├── src
    └── adapter
    └── domain
      └── model
    └── repository
    └── usecase
    └── main.go
├── initdb.d

```
- model: 개체의 구조와 메서드 정의합니다.
- repository: 데이터베이스와 연결, 데이터 처리를 담당합니다.
- usecase: 데이터를 가공, 비지니스 로직을 처리합니다.
- adapter: usecase의 output을 가져와 표시합니다.

## 데이터베이스
- MySQL:5.7
- 데이터 테이블 구조는 아래와 같습니다.

**users: 회원 정보를 저장**

| Key            | Datatype | Value                                  |
|----------------|----------|----------------------------------------|
| id             | INT      | 회원을 구분하는 고유 아이디                        |
| password       | VARCHAR  | 회원가입 시 입력한 비밀번호(해쉬값)                   |
| phone_number   | VARCHAR  | 회원의 휴대폰 번호                             |
| created_at     | DATETIME | 회원 가입 시각                               |
| deleted_at     | DATETIME | 회원 정보 삭제 시각                            |

**product: 상품 정보를 저장**

| Key             | Datatype | Value            |
|-----------------|----------|------------------|
| id              | INT      | 상품을 구분하는 고유 아이디  |
| category        | VARCHAR  | 카테고리             |
| price           | FLOAT    | 가격               |
| cost            | FLOAT    | 원가               |
| name            | VARCHAR  | 이름               |
| description     | VARCHAR  | 설명               |
| barcode         | VARCHAR  | 바코드              |
| expiration_date | DATETIME | 유통기한             |
| size            | VARCHAR  | 사이즈              |
| user_id         | INT      | 상품을 등록한 회원의 고유번호 |
| name_chosung    | VARCHAR  | 검색을 위한 상품 이름의 초성 |
| created_at      | DATETIME | 상품 등록 시각         |
| updated_at      | DATETIME | 상품 수정 시각         |
| deleted_at      | DATETIME | 상품 삭제 시각         |

# 프로젝트 실행 방법
도커가 설치되어있어야 합니다.

[Docker](https://www.docker.com/get-started) 설치 & 로그인 (tested on v4.3.0)

프로젝트를 다운 받아 해당 폴더로 위치합니다.
```bash
git clone https://github.com/Suzzzzzy/product_manager_project.git
```

DB를 세팅합니다.
```bash
docker-compose up
```
- 데이터베이스를 docker-compose 로 구성하면서, 필요한 리소스 데이터를 initdb sql 파일을 이용하여 import 합니다.
```bash
cd src
go run main.go
```
- 소스코드 폴더로 이동합니다.
- Go 프로그램을 실행합니다.
- docker-compose로 띄운 mysql에 연결되며 프로그램이 실행됩니다.

http://localhost:8080/ 혹은 http://0.0.0.0:8080/ 접속했을 때, "Hello world"가 출력된다면 서버가 정상적으로 실행된 것입니다.

## API 명세서
https://documenter.getpostman.com/view/19629582/2s9YsJCD8c

# 테스트 코드 실행
## usecase 테스트
`mockery`, `testify`를 사용하여 DB와의 영속성 보다는 비지니스 로직을 검증합니다.
- 데이터를 Mocking하여 가상의 데이터를 생성하고 비지니스 로직이 잘 작동하는지 확인합니다.
- `mockery`: 명령어를 사용하여 특정 인터페이스(repository_interface)에 대한 mock을 자동으로 생성합니다.
- `testify`: Go언어 테스트 라이브러리로, suite 패키지를 사용하여 여러가지의 테스트 케이스를 그룹화하였습니다.

usecase 테스트를 실행하고, coverage를 출력하는 명령어는 아래와 같습니다.
```bash
make td-usecase
```