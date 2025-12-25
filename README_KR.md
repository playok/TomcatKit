# TomcatKit

**[English](README.md)** | **[한국어](README_KR.md)** | **[日本語](README_JP.md)**

[![Go Version](https://img.shields.io/badge/Go-1.21+-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/playok/TomcatKit)](https://goreportcard.com/report/github.com/playok/TomcatKit)
[![Tomcat](https://img.shields.io/badge/Tomcat-9.0-F8DC75?style=flat&logo=apache-tomcat)](https://tomcat.apache.org/)
[![AI Generated](https://img.shields.io/badge/AI%20Generated-Claude%20Code-blueviolet?style=flat&logo=anthropic)](https://claude.ai/claude-code)

Apache Tomcat 9.0 설정 관리를 위한 CLI 기반 TUI(텍스트 사용자 인터페이스) 유틸리티입니다.

## 데모

![TomcatKit Demo](docs/assets/demo.gif)

## 주요 기능

- **대화형 TUI**: [tview](https://github.com/rivo/tview) 기반의 ncurses 스타일 터미널 인터페이스
- **포괄적인 설정**: Tomcat 9.0의 모든 주요 설정 영역 지원
- **자동 감지**: 환경 변수, 일반 경로, 실행 중인 프로세스에서 Tomcat 설치 자동 감지
- **안전한 편집**: 설정 파일 수정 전 자동 백업 생성
- **다중 인스턴스 지원**: 최근 사용한 Tomcat 인스턴스 기억
- **다국어 지원**: 영어, 한국어, 일본어 (F2로 전환)
- **컬러 UI**: 의미에 따른 직관적인 버튼 색상
  - 초록색: 저장, 추가, 적용
  - 빨간색: 삭제, 제거, 뒤로
  - 노란색: 취소
  - 파란색: 네비게이션 (컨텍스트, 파라미터)
- **상황별 도움말**: 각 설정 필드에 대한 속성 도움말 패널
- **실시간 XML 미리보기**: 설정 변경 사항 실시간 미리보기

## 지원 설정 모듈

| 모듈 | 상태 | 설명 |
|------|------|------|
| Server | 완료 | server.xml 핵심 설정 (Server, Service, Engine, Host) |
| Connector | 완료 | HTTP, AJP, SSL/TLS 커넥터 및 스레드 풀 |
| Security/Realm | 완료 | 인증 Realm 및 tomcat-users.xml 관리 |
| JNDI Resources | 완료 | DataSource, Mail Session, Environment, Resource Links |
| Virtual Hosts | 완료 | Host, Context, Parameters, Session Manager 설정 |
| Valves | 완료 | AccessLog, RemoteAddr, RemoteIp, ErrorReport, SSO, StuckThread 밸브 |
| Clustering | 완료 | 세션 복제, 멤버십, 인터셉터, 팜 디플로이어 |
| Logging | 완료 | JULI logging.properties, 파일 핸들러, 로거 |
| Context | 완료 | context.xml 설정, 리소스, 쿠키, 세션 매니저 |
| Web | 완료 | web.xml 서블릿, 필터, 세션, 보안 제약 |
| Quick Templates | 완료 | 가상 스레드, HTTPS, 커넥션 풀, Gzip, 보안 |

## 설치

### 소스에서 빌드

```bash
# 저장소 클론
git clone https://github.com/playok/tomcatkit.git
cd tomcatkit

# 빌드
make build

# 또는 go 직접 사용
go build -o bin/tomcatkit ./cmd/tomcatkit
```

### 요구 사항

- Go 1.21 이상
- Apache Tomcat 9.0 설치

## 사용법

### 기본 사용법

```bash
# 자동 감지로 실행
./bin/tomcatkit

# Tomcat 경로 직접 지정
./bin/tomcatkit -home /opt/tomcat -base /var/tomcat/instance1

# 버전 표시
./bin/tomcatkit -version

# 도움말 표시
./bin/tomcatkit -help
```

### CLI 옵션

| 옵션 | 설명 |
|------|------|
| `-home` | CATALINA_HOME 경로 (Tomcat 설치 디렉토리) |
| `-base` | CATALINA_BASE 경로 (인스턴스 디렉토리, 기본값: CATALINA_HOME) |
| `-version` | 버전 정보 표시 |
| `-help` | 도움말 표시 |

### 네비게이션

| 키 | 동작 |
|----|------|
| 방향키 | 메뉴 및 목록 탐색 |
| Enter | 항목 선택 또는 확인 |
| Escape | 한 단계 뒤로 |
| Tab | 폼 필드 간 이동 |
| F2 | 언어 전환 (EN/KR/JP) |
| q | 프로그램 종료 |
| Ctrl+C | 강제 종료 |

## 프로젝트 구조

```
tomcatkit/
├── cmd/
│   └── tomcatkit/
│       └── main.go           # 애플리케이션 진입점
├── internal/
│   ├── config/
│   │   ├── tomcat.go         # Tomcat 인스턴스 설정
│   │   ├── settings.go       # 애플리케이션 설정 저장
│   │   ├── server/           # server.xml 타입 및 작업
│   │   ├── connector/        # 커넥터 프로토콜 및 기본값
│   │   ├── realm/            # Realm 타입 및 tomcat-users.xml
│   │   ├── jndi/             # JNDI 리소스 타입 및 context.xml
│   │   ├── logging/          # 로깅 설정
│   │   └── web/              # web.xml 타입 및 작업
│   ├── detector/             # Tomcat 자동 감지
│   ├── i18n/                 # 국제화 (EN/KR/JP)
│   ├── parser/               # XML 파싱 유틸리티
│   └── tui/
│       ├── app.go            # 메인 TUI 애플리케이션
│       └── views/            # 설정 뷰
├── go.mod
├── go.sum
├── Makefile
└── README.md
```

## 설정 파일

TomcatKit이 관리하는 Tomcat 설정 파일:

- `$CATALINA_BASE/conf/server.xml` - 메인 서버 설정
- `$CATALINA_BASE/conf/tomcat-users.xml` - 사용자 및 역할 정의
- `$CATALINA_BASE/conf/context.xml` - 기본 컨텍스트 설정
- `$CATALINA_BASE/conf/web.xml` - 기본 웹 애플리케이션 설정
- `$CATALINA_BASE/conf/logging.properties` - JULI 로깅 설정

## 설정 저장 위치

애플리케이션 설정 저장 위치:
- Linux/macOS: `~/.config/tomcatkit/settings.json`
- Windows: `%APPDATA%\tomcatkit\settings.json`

저장되는 설정:
- 마지막 사용한 Tomcat 인스턴스
- 최근 인스턴스 경로
- 선호 언어

## 프로젝트 소개

이 프로젝트는 재미와 학습 목적으로 만든 **취미 프로젝트**입니다. AI 지원 개발을 탐구하고 Tomcat 관리자에게 유용한 도구를 제공하기 위해 만들어졌습니다.

### AI 생성

이 프로젝트는 **[Claude Code](https://claude.ai/claude-code)** (Anthropic의 Claude)를 사용하여 AI가 전적으로 생성했습니다.

- **AI 모델**: Claude Opus 4.5 (`claude-opus-4-5-20251101`)
- **개발 도구**: Claude Code CLI
- **인간의 역할**: 프로젝트 방향 설정, 요구사항 명세, 검토

이 저장소의 모든 코드, 문서, 설정은 AI 지원 개발을 통해 생성되었습니다. AI가 아키텍처 설계, 구현, 디버깅, 문서화를 담당하고 인간이 가이드와 검증을 제공했습니다.

> **참고**: 이 프로젝트는 개인 취미 프로젝트이며 Apache Software Foundation과 관련이 없습니다.

## 라이선스

MIT 라이선스 - 자세한 내용은 [LICENSE](LICENSE) 파일을 참조하세요.

## 기여

기여를 환영합니다! Pull Request를 자유롭게 제출해 주세요.

## 작성자

[playok](https://github.com/playok)

---

<p align="center">
  <sub>Claude Code의 AI 지원으로 제작됨</sub>
</p>
