
// pipeline {
//     agent any
//     stages   {
//         stage('Build') {
//             agent { docker { image 'golang' }  }
//             steps {
//                 // Create our project directory.
//                 // sh 'cd ${GOPATH}/src'
//                 // sh 'mkdir -p ${GOPATH}/src/go-kubernetes'
//                 // // Copy all files in our Jenkins workspace to our project directory.
//                 // sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/go-kubernetes'
//                 // Build the app.
//                 sh 'pwd'
//                 sh 'go build'     }
//         }
//         stage('Publish') {
//                 environment {
//                     registryCredential = 'dockerhub' }
//                     steps {
//                         script { def appimage = docker.build registry + ":$BUILD_NUMBER"
//                         docker.withRegistry( '', registryCredential ) {
//                             appimage.push()
//                             appimage.push('latest')
//                         }
//                         }
//                     }
//         }
//     }
// }




pipeline {
    agent any
    environment{
        DOCKER_TAG = getDockerTag()
        IMAGE_URL_WITH_TAG = "divalsehgal/go-hello-world:1.0.0"
    }
    stages{
        stage('Build Docker Image'){
            steps{
                sh "docker build . -t ${IMAGE_URL_WITH_TAG}"
            }
        }
        stage('docker Push'){
            steps{
                withCredentials([string(credentialsId: 'dockerhub', variable: 'dockerhub')]) {
                    sh "docker login -u divalsehgal -p ${getwellsoon}"
                    sh "docker push ${IMAGE_URL_WITH_TAG}"
                }
            }
        }
       
    }
}

def getDockerTag(){
    def tag  = sh script: 'git rev-parse HEAD', returnStdout: true
    return tag
}