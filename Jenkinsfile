pipeline {

    agent any

    stages {

        stage("Build") {
            steps {
                withEnv(['PATH+EXTRA='/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:/usr/games:/usr/local/games:/snap/bin:/usr/local/go/bin:/usr/local/go/bin:/home/ng-voice/go/bin'])
                sh "go version"
                sh "helm version"
                sh "docker version"
            }
        }
    }
}