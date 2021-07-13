# Run application

//TODO

# Run tests

```
docker-compose exec go go test -v go-cource-api/infrustructure/persistence
docker-compose exec go go test -v go-cource-api/interfaces/handlers
```

## Generating mocks

Application uses gomock for mocks. See [gomock](https://github.com/golang/mock) for reference

# Generate API documentations

Docs can be viewed on http://localhost:8000/swagger/index.html

To regenerate docs execute:

```
swag init
```

# Known Issues

- Facebook Oauth may not return email(if no email, or it's not confirmed). 
  In this case application Sign In will fail
  
# TODO

- [ ] Fix docker development setup clashing with local go
- [ ] Fix database migrations on every test(Migrate once and rollback inserted data after each test)
- [ ] Implement Twitter oauth(waiting for dev acc app verification)
- [ ] Implement Own error handler
- [ ] Add refresh token
- [ ] Try to use identity broker for all Auth e.g keycloak
- [X] Make modular routing
- [x] Add content negotiation
- [ ] Add CI Dron tests pipeline
- [x] Protect comment form non owner delete
- [ ] Extract test database config to envs
- [ ] Add tests for social oauth