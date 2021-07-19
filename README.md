# Run application

1. Clone repository
2. Copy and rename .env.example to .env .Modify env variables if needed 
3. Run:
```
docker-compose up --build
```
4. After that application will be available at http://localhost:8000

# Run tests

```
docker-compose exec go go test -v go-cource-api/infrustructure/persistence
docker-compose exec go go test -v go-cource-api/application/handlers
docker-compose exec go go test -v go-cource-api/application/lang
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

- Facebook Oauth may not return email(if no email, or it's not confirmed). In this case application Sign In will fail
- Application uses  codegangsta/gin for hot reload. It will automatically start watching files after docker-compose up.
  But there is a delay where gin actually will build application and start to serve it. So, need to wait until application will
  be built before start using it.
  
# TODO

- [ ] Fix docker development setup clashing with local go
- [ ] Fix database migrations on every test(Migrate once and rollback inserted data after each test)
- [ ] Implement Twitter oauth(waiting for dev acc app verification)
- [X] Implement Own error handler
- [ ] Add refresh token
- [ ] Try to use identity broker for all Auth e.g keycloak
- [X] Make modular routing
- [x] Add content negotiation
- [ ] Add CI Dron tests pipeline
- [x] Protect comment form non owner delete
- [ ] Extract test database config to envs
- [ ] Add tests for social oauth
- [ ] Pass url to templates(login, register). Don't hardcode it. Mb, pass config or use helper func