project details

USE_FIRESTORE=true \
GOOGLE_CLOUD_PROJECT=ls-devx-int-np-e7d3 \
FIRESTORE_DATABASE=laena-quotes-api-db \
FIRESTORE_COLLECTION=quotes \
go run ./cmd/rest

USE_FIRESTORE=true \
GOOGLE_CLOUD_PROJECT=ls-devx-int-np-e7d3 \
FIRESTORE_DATABASE=laena-quotes-api-db \
FIRESTORE_COLLECTION=quotes \
go run ./cmd/graphql

GraphQL test (JSON body required):
curl -X POST http://localhost:8081/query \
	-H "Content-Type: application/json" \
	-d '{"query":"query { quotes { id text author } }"}'

gcloud auth login
gcloud auth application-default login


gcloud run services add-iam-policy-binding rest-api \
 --region=us-east4 \
 --member="allUsers" \
 --role="roles/run.invoker"