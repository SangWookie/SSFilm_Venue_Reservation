# Roadmap

- [ ] Implement full UI screen
    - [ ] Calendar
    - [ ] Reservation Information
    - [ ] Admin Dashboard
- [ ] Reactivity fixes
- [ ] Implement Animations
    - [ ] View transitions
- [ ] Implement Business logic
- [ ] Write API utils from actual api server
- [ ] Lastly check multi-browser support. (e.g: backdrop filter with blur does not work at webkit side)
- [ ] DISABLE SEO
- [ ] Check fail-safe

# Tech debts/Concerns

- 로그인제 없이 앱이 작성될 경우 다음과 같은 우려 사항 발생
    - 학번/이름/이메일 인증 없을 경우 잘못된 정보로 도배 발생
    - 예약 취소 불가능.
    - 자동 승인일 경우 더 꼬일 여지가 있음
- js 생태계에서 date 관리 매우 까다로움
    - `luxon` 라이브러리로 어느 정도 해결했다고 보는 중이지만, 브라우저별로 다르게 행동하는 concern이 존재.
- `camelCase`와 `snake_case` 중 하나로 정리해야만 함. 최대한 노력 중이지만 섞어 쓰는 문제가 발생할 여지가 존재함.
