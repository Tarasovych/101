Kubernetes setup:
1. [Install](https://skaffold.dev/docs/install/) Skaffold
2. [Install](https://kubernetes.io/docs/tasks/tools/install-kubectl/) kubectl
3. [Install](https://minikube.sigs.k8s.io/docs/start/) minikube
4. ```minikube start --kubernetes-version=v1.21.5```
5. ```minikube addons enable metrics-server```
6. ```skaffold dev```

Docker setup:
1. [Install](https://docs.docker.com/engine/install/) Docker Engine
2. [Install](https://docs.docker.com/compose/install/) Docker Compose
3. ```docker-compose --profile nginx up --scale golang=2 --build```
