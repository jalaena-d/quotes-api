docker build --no-cache -f Dockerfile.rest \
  -t rest-api \
  -t us-east4-docker.pkg.dev/ls-devx-int-np-e7d3/laena-quotes-api/rest-api:latest \
  .

```
docker build --no-cache -f Dockerfile.rest -t rest-api -t us-east4-docker.pkg.dev/ls-devx-int-np-e7d3/laena-quotes-api/rest-api:latest .
```
```
docker build --no-cache -f Dockerfile -t graphql-api -t us-east4-docker.pkg.dev/ls-devx-int-np-e7d3/laena-quotes-api/graphql-api:latest .
```