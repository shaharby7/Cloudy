docker build -t shaharshahar/fakeprovider:latest -f ./build/fakeprovider.Dockerfile .
docker push shaharshahar/fakeprovider:latest
minikube update-context --profile fakeprovider
kubectl rollout restart deployment/fakeprovider-api -n runspace