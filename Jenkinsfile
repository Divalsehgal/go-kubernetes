pipeline {
    agent any
    environment{
        IMAGE_URL_WITH_TAG = "divalsehgal/go-hello-world:1.0.0"
    }
    stages{
        stage('Build Docker Image'){
            steps{
                sh "docker build . -t ${IMAGE_URL_WITH_TAG}"
            }
        }
        stage('Docker Push'){
            steps{
                    sh "docker login -u divalsehgal -p getwellsoon"
                    sh "docker push ${IMAGE_URL_WITH_TAG}"
            }
        }
        stage('Pods Deploy'){
            steps{
                    sh "minikube start"
                    sh "kubectl apply -f k8s-deployment.yml"
            }
        }
        stage('Service Deploy'){
            steps{
                    sh "kubectl apply -f k8s-services.yml"
            }
        }
    }
}
