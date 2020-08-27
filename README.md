# simple-go
Simple app(go) with docker build/push, deploy on kubernetes with Jenkins pipeline

## Environment
```
1 x Raspberry Pi3 1GB (192.168.1.118)
1 x Raspberry Pi4 4GB (192.168.1.181)
1 x Raspberry Pi4 8GB (192.168.1.185)
```

k3s running on raspberrypi cluster
* Master (Raspberry Pi4 4gb)
* Worker (Raspberry Pi4 8gb)

Jenkins
* Raspberry Pi3 1GB

Jenkins Agent
* Raspberry Pi4 8GB (Responsible for build docker images)

Docker registry 
* Raspberry Pi3 1GB

![Image Alt](https://raw.githubusercontent.com/nopp/simple-go/master/.img/jenkins.png)


Output from pipeline ;)

```
[Pipeline] Start of Pipeline (hide)
[Pipeline] node
Running on raspberrypi-8gb in /ansible/workspace/Simple-go
[Pipeline] {
[Pipeline] stage
[Pipeline] { (Pull Repository)
[Pipeline] git
The recommended git tool is: NONE
No credentials specified
Fetching changes from the remote Git repository
 > git rev-parse --is-inside-work-tree # timeout=10
 > git config remote.origin.url https://github.com/nopp/simple-go.git # timeout=10
Fetching upstream changes from https://github.com/nopp/simple-go.git
 > git --version # timeout=10
 > git --version # 'git version 2.20.1'
 > git fetch --tags --force --progress -- https://github.com/nopp/simple-go.git +refs/heads/*:refs/remotes/origin/* # timeout=10
Checking out Revision 4548d9280176710dc90340ad9deb3a0169ba47c5 (refs/remotes/origin/master)
Commit message: "Update app.go"
 > git rev-parse refs/remotes/origin/master^{commit} # timeout=10
 > git rev-parse refs/remotes/origin/origin/master^{commit} # timeout=10
 > git config core.sparsecheckout # timeout=10
 > git checkout -f 4548d9280176710dc90340ad9deb3a0169ba47c5 # timeout=10
 > git branch -a -v --no-abbrev # timeout=10
 > git branch -D master # timeout=10
 > git checkout -b master 4548d9280176710dc90340ad9deb3a0169ba47c5 # timeout=10
 > git rev-list --no-walk 4548d9280176710dc90340ad9deb3a0169ba47c5 # timeout=10
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (Git checkout tag)
[Pipeline] sh
+ git checkout 2.0
Note: checking out '2.0'.

You are in 'detached HEAD' state. You can look around, make experimental
changes and commit them, and you can discard any commits you make in this
state without impacting any branches by performing another checkout.

If you want to create a new branch to retain commits you create, you may
do so (now or later) by using -b with the checkout command again. Example:

  git checkout -b <new-branch-name>

HEAD is now at a3cdae5 Update app.go
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (Unit test)
[Pipeline] sh
+ go test -v
=== RUN   TestGetVersion
--- PASS: TestGetVersion (0.00s)
=== RUN   TestGetRoot
--- PASS: TestGetRoot (0.00s)
PASS
ok  	simple-go	0.020s
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (Build Docker Imge)
[Pipeline] sh
+ docker build -t registry.carlosmalucelli.com/simple-go:2.0 .
Sending build context to Docker daemon  613.4kB

Step 1/4 : FROM golang:1.14
 ---> d5a553a9a71e
Step 2/4 : COPY . .
 ---> e6d26d9d64b9
Step 3/4 : RUN unset GOPATH     && go build -o main .
 ---> Running in 02b06a970091
[91mgo: downloading github.com/gorilla/mux v1.8.0
[0mRemoving intermediate container 02b06a970091
 ---> df410463f2dc
Step 4/4 : ENTRYPOINT ["./main"]
 ---> Running in 1acb9274aec4
Removing intermediate container 1acb9274aec4
 ---> cef0bc47fe11
Successfully built cef0bc47fe11
Successfully tagged registry.carlosmalucelli.com/simple-go:2.0
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (Push Docker Imge)
[Pipeline] sh
+ docker push registry.carlosmalucelli.com/simple-go:2.0
The push refers to repository [registry.carlosmalucelli.com/simple-go]
ae1486513703: Preparing
cef7bd41fe79: Preparing
5bc3988dfdc4: Preparing
4d4ff09cedf8: Preparing
b90d1e309087: Preparing
5d41b2315360: Preparing
02beba293b78: Preparing
b8676a7dda39: Preparing
07aaf212e4f3: Preparing
5d41b2315360: Waiting
02beba293b78: Waiting
b8676a7dda39: Waiting
07aaf212e4f3: Waiting
4d4ff09cedf8: Layer already exists
b90d1e309087: Layer already exists
5bc3988dfdc4: Layer already exists
02beba293b78: Layer already exists
5d41b2315360: Layer already exists
b8676a7dda39: Layer already exists
07aaf212e4f3: Layer already exists
cef7bd41fe79: Pushed
ae1486513703: Pushed
2.0: digest: sha256:0ffc2ef4b23e014ae897ef6cefd77a0bb666e56607f2b50e07b7d97686f620e3 size: 2216
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (Kubernetes apply service)
[Pipeline] sh
+ kubectl apply -f k8s/service.yaml
service/simple-go-svc unchanged
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (Change tag version on deployment)
[Pipeline] sh
+ sed -i s/XXVERSIONXX/2.0/g k8s/deploy.yaml
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (Kubernetes apply deployment)
[Pipeline] sh
+ kubectl apply -f k8s/deploy.yaml
deployment.apps/simple-go configured
[Pipeline] }
[Pipeline] // stage
[Pipeline] stage
[Pipeline] { (Kubernetes rollout status)
[Pipeline] sh
+ kubectl rollout status deploy simple-go --timeout=300s
deployment "simple-go" successfully rolled out
[Pipeline] }
[Pipeline] // stage
[Pipeline] }
[Pipeline] // node
[Pipeline] End of Pipeline
Finished: SUCCESS
```
