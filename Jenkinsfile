
pipeline {
    agent any
    stages   {
        stage('Build') {
            agent { docker { image 'golang' }  }
            steps {
                // Create our project directory.
                // sh 'cd ${GOPATH}/src'
                // sh 'mkdir -p ${GOPATH}/src/go-kubernetes'
                // // Copy all files in our Jenkins workspace to our project directory.
                // sh 'cp -r ${WORKSPACE}/* ${GOPATH}/src/go-kubernetes'
                // Build the app.
                sh 'pwd'
                sh 'go build'     }
        }
        stage('Publish') {
                environment {
                    registryCredential = 'dockerhub' }
                    steps {
                        script { def appimage = docker.build registry + ":$BUILD_NUMBER"
                        docker.withRegistry( '', registryCredential ) {
                            appimage.push()
                            appimage.push('latest')
                        }
                        }
                    }
        }
    }
}
