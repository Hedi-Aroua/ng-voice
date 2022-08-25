pipeline {

    agent any

    environment {
        PATH = "/usr/locel/bin:${env.PATH}"
    }

    stages {

        stage("Build") {
            steps {
                sh "go version"
                sh "helm version"
                sh "docker version"
            }
        }
    }
}