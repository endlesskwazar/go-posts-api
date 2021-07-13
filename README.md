# Run application

//TODO

# Run tests

```
docker-compose exec go go test -v go-cource-api/infrustructure/persistence
docker-compose exec go go test -v go-cource-api/interfaces/handlers
```

## Generating mocks

//TODO

# Generate API documentations

//TODO

# Known Issues

- Facebook Oauth may not return email(if no email, or it's not confirmed). 
  In this case application Sign In will fail
  
# TODO

- [ ] Fix docker development setup clashing with local go
- [ ] Figure out how to not use magic number for authenticated user in tests for id. And not to duplicate local const in every test
- [ ] Fix database migrations on every test(Migrate once and rollback inserted data after each test)
- [ ] Implement Twitter oauth(waiting for dev acc app verification)
- [ ] Implement Own errors with extended information. Implement Own error handler
- [ ] Implement refresh token and token expiration
- [ ] Try to use identity broker for all Auth e.g keycloak
- [ ] Make modular routing
- [x] Add content negotiation
- [ ] Add CI Dron tests pipeline
- [ ] Protect comment form non owner delete
- [ ] Extract test database config to envs