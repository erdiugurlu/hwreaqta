#!/usr/bin/env groovy

timestamps {
    timeout(time: 30, unit: "MINUTES") {
        node {
            properties([
                buildDiscarder(logRotator(numToKeepStr: "100", daysToKeepStr: "28")),
                disableConcurrentBuilds()
            ])

            def name = "reaqta-helloworld"
            def dockerRegistry = "docker.io/erdiugurlu"

            def version
            def image

            stage("Clean up workspace") {
                sh(script: "env | sort --unique")
                sh(script: "ls -la")
                deleteDir()
            }

            checkout scm

            stage("Version: update") {
                def versionFile = readFile(file: "VERSION")
                def splitVersion = versionFile.tokenize(".")
                def majorVersion = splitVersion.get(0) as Integer
                def minorVersion = splitVersion.get(1) as Integer

                if (env.BRANCH_NAME == "master") {
                    version = "${majorVersion}.${minorVersion}"
                } else if (env.BRANCH_NAME == "main") {
                    version = "${majorVersion}.${minorVersion}"
                } else {
                    version = "${majorVersion}.${minorVersion}-${BRANCH_NAME.replaceAll(/[-_\/]/, ".")}.${env.BUILD_ID}.SNAPSHOT"
                }
            }

            stage("Docker: build") {
                image = docker.build("${dockerRegistry}/${name}:${version}", ".")
            }

            stage("Docker: push") {
                docker.withRegistry('', 'docker_user') {
                        image.push("${version}")
                    }
                docker.withRegistry('', 'docker_user') {
                        image.push("latest")
                }
                echo("Docker Image pushed: ${dockerRegistry}/${name}:${version}")
            }
        }
    }
}