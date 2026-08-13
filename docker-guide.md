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

docker run --rm --name rest-api -p 8080:8080 rest-api:latest
us-east4-docker.pkg.dev/ls-devx-int-np-e7d3/laena-quotes-api

1. docker build graphql
2. push graphql into artifact reg
3. edit cloud run service graphql sa GUI
4. select container image to latest in artifact reg repo