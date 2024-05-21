`minikube start`

`minikube kubectl -- create -f deploy.yaml`

`minikube kubectl -- apply -f deploy.yaml`

`minikube kubectl -- delete deployment my-hello-echo-app`

`minikube kubectl -- get deploy`

`minikube kubectl -- get pods`

`minikube kubectl -- port-forward deploy/hello-echo-app 15000:15000`