# teamcore-project-46344

## Deployment GCP

### How to config envs

Create a `.env` file using his respective example.

`.env` is used to get envs when you run the project locally.

### In your source code directory, deploy from source using the following command:

`gcloud run deploy --set-env-vars "API_TOKEN=<TOKEN>"`

- If prompted to enable the API, Reply y to enable.

- When you are prompted for the source code location, press Enter to deploy the current folder.

- When you are prompted for the service name, press Enter to accept the default name, teamcore-project-46344.

- If you are prompted to enable the Artifact Registry API, respond by pressing y.

- When you are prompted for region: select the region of your choice, for example us-central1.

- You will be prompted to allow unauthenticated invocations: respond y .

- Then wait a few moments until the deployment is complete. On success, the command line displays the service URL.

- Visit your deployed service by opening the service URL in a web browser.

## Tests

Run tests using the following command:

`go test`

## Inspirations

#### Folders structure

[This structure allows you to separate the transport layer; whether you user HTTP, GRPC or whatever mechanism you want.](https://www.reddit.com/r/golang/comments/a35xfv/comment/eb4784e/?utm_source=share&utm_medium=web2x&context=3).
