pipeline {

    agent any

    environment {
        DOCKERHUB_CREDENTIALS = credentials('heidist-dockerhub')
    }

    stages {

        stage('Build') {
            steps {
                withEnv(['PATH+EXTRA=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/usr/local/go/bin:/usr/local/go/bin:/home/ng-voice/go/bin']) {
                sh 'go version'
                sh 'go build'
                }
            }
        }

        stage('Docker Image') {
            steps {
                sh 'docker build -t heidist/casestudy:latest .'
                sh 'echo $DOCKERHUB_CREDENTIALS_PSW | docker login -u $DOCKERHUB_CREDENTIALS_USR --password-stdin'
                sh 'docker push heidist/casestudy:latest'
                }
            }
        

        stage('Deploy') {
            steps {
                sh 'helm version'
                sh 'helm delete case-chart'
                sh 'helm install case-chart casestudy/ --values casestudy/values.yaml'
                }
            }
    }
    
}