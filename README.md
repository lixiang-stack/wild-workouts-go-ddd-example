# Wild Workouts

*The idea for this series, is to apply DDD by refactoring. This process is in progress! Please check articles, to know the current progress.*

Wild Workouts is an example project that we created to show how to build Go applications that are **easy to develop, maintain, and fun to work with, especially in the long term!**

No application is perfect from the beginning. With over a dozen coming articles, we will uncover what issues you can find in the current implementation. We will also show how to fix these issues and achieve clean implementation by refactoring.

### Articles

1. [**Too modern Go application? Building a serverless application with Google Cloud Run and Firebase**](https://threedots.tech/post/serverless-cloud-run-firebase-modern-go-application/?utm_source=github.com)
2. [**A complete Terraform setup of a serverless application on Google Cloud Run and Firebase**](https://threedots.tech/post/complete-setup-of-serverless-application/?utm_source=github.com)
3. [**Robust gRPC communication on Google Cloud Run (but not only!)**](https://threedots.tech/post/robust-grpc-google-cloud-run/?utm_source=github.com)
4. [**You should not build your own authentication. Let Firebase do it for you.**](https://threedots.tech/post/firebase-cloud-run-authentication/?utm_source=github.com)
5. [**Business Applications in Go: Things to know about DRY**](https://threedots.tech/post/things-to-know-about-dry/?utm_source=github.com)
6. [**When microservices in Go are not enough: introduction to DDD Lite**](https://threedots.tech/post/ddd-lite-in-go-introduction/?utm_source=github.com)
7. [**Repository pattern: painless way to simplify your Go service logic**](https://threedots.tech/post/repository-pattern-in-go/?utm_source=github.com)
8. *More articles are on the way!*

### Directories

- [api](api/) OpenAPI and gRPC definitions
- [docker](docker/) Dockerfiles
- [internal](internal/) application code
- [scripts](scripts/) deployment and development scripts
- [terraform](terraform/) - infrastructure defintion
- [web](web/) - frontend JavaScript code

### Live Demo

The example application is available at [https://threedotslabs-wildworkouts.web.app/](https://threedotslabs-wildworkouts.web.app/).

### Code Generation

The project uses `//go:generate` directives to generate OpenAPI types/servers and gRPC code. To regenerate all code:

```bash
# Install required tools first (one-time setup)
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.3.6

# Generate all code (run from each module directory)
cd internal/trainer && go generate ./... && cd -
cd internal/trainings && go generate ./... && cd -
cd internal/users && go generate ./... && cd -
cd internal/common && go generate ./... && cd -

# Or simply use make
make generate
```

You can also view all available make targets with `make help`.

### Refactoring
#### v1：with DDD Lite
1. 优化trainer.proto：使用业务语义定义rpc服务接口。
2. 新增trainer/domain：按照领域建模Hour对象，封装业务规则，并定义业务行为repository（存储层接口）。
3. 新增hour_repository.go：特定类型database实现存储层接口。
4. 更新grpc.go和http.go：只做流程编排，不设计具体逻辑（不展开细节），细节由domain和repository负责。
5. 更新trainings：修改调用trainer的rpc接口。
6. 其他：DTO

#### v2：with Repository pattern
1. 新增FactoryConfig：
	作为领域对象hour的可变配置项，hour.go内部全局的校验收敛到FactoryConfig对象内部。FactoryConfig内部包装New*Hour，对外统一使用FactoryConfig提供的NewHour接口。
2. 新增多存储库，repository的多种实现。

#### v3：add unit test
1. test tables
2. assert functions
3. parallel execution
4. black-box testing
5. 使用方式: make test或make test -- -v

### Running locally

```go
> docker-compose up

# ...

web_1             |  INFO  Starting development server...
web_1             |  DONE  Compiled successfully in 6315ms11:18:26 AM
web_1             |
web_1             |
web_1             |   App running at:
web_1             |   - Local:   http://localhost:8080/
web_1             |
web_1             |   It seems you are running Vue CLI inside a container.
web_1             |   Access the dev server via http://localhost:<your container's external mapped port>/
web_1             |
web_1             |   Note that the development build is not optimized.
web_1             |   To create a production build, run yarn build.
```

### Google Cloud Deployment

```go
> cd terraform/
> make

Fill all required parameters:
	project [current: wild-workouts project]:       # <----- put your Wild Workouts Google Cloud project name here (it will be created) 
	user [current: email@gmail.com]:                # <----- put your Google (Gmail, G-suite etc.) e-mail here
	billing_account [current: My billing account]:  # <----- your billing account name, can be found here https://console.cloud.google.com/billing
	region [current: europe-west1]: 
	firebase_location [current: europe-west]: 

# it may take a couple of minutes...

The setup is almost done!

Now you need to enable Email/Password provider in the Firebase console.
To do this, visit https://console.firebase.google.com/u/0/project/[your-project]/authentication/providers

You can also downgrade the subscription plan to Spark (it's set to Blaze by default).
The Spark plan is completely free and has all features needed for running this project.

Congratulations! Your project should be available at: https://[your-project].web.app

If it's not, check if the build finished successfully: https://console.cloud.google.com/cloud-build/builds?project=[your-project]

If you need help, feel free to contact us at https://threedots.tech
```

### Troubleshooting

#### 1. `oapi-codegen: Command not found`

The project requires `oapi-codegen` v1.3.6 to generate OpenAPI code. Install it with:

```bash
go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.3.6
```

Make sure `$GOPATH/bin` is in your `PATH`:

```bash
export PATH=$PATH:$(go env GOPATH)/bin
```

You can add the above line to your `~/.bashrc` or `~/.zshrc` to make it permanent.

#### 2. `go.mod` outdated (`updates to go.mod needed`)

If you see errors like `updates to go.mod needed; to update it: go mod tidy`, run `go mod tidy` for each sub-module:

```bash
cd internal/common && go mod tidy && cd -
cd internal/trainer && go mod tidy && cd -
cd internal/trainings && go mod tidy && cd -
cd internal/users && go mod tidy && cd -
```

### Screenshots

![Wild Workouts login](https://threedots.tech/media/serverless-cloud-run-firebase-modern-go-app/login.png "Logo Title Text 1")
![Wild Workouts trainer's schedule](https://threedots.tech/media/serverless-cloud-run-firebase-modern-go-app/schedule.png "Logo Title Text 1")
![Wild Workouts schedule training](https://threedots.tech/media/serverless-cloud-run-firebase-modern-go-app/new-training.png "Logo Title Text 1")