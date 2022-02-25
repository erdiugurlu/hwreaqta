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
                    minorVersion = minorVersion + 1
                    version = "${majorVersion}.${minorVersion}"
                } else if (env.BRANCH_NAME == "main") {
                    minorVersion = minorVersion + 1
                    version = "${majorVersion}.${minorVersion}"
                } else {
                    version = "${majorVersion}.${minorVersion}-${BRANCH_NAME.replaceAll(/[-_\/]/, ".")}.${env.BUILD_ID}.SNAPSHOT"
                }
            }

            stage("Docker: build") {
                image = docker.build("${dockerRegistry}/${name}:${version}", ".")
            }

            stage("Docker: push") {
                docker.withRegistry('', 'dockerhub_user') {
                        image.push("${version}")
                    }
                docker.withRegistry('', 'dockerhub_user') {
                        image.push("latest")
                }
                echo("Docker Image pushed: ${dockerRegistry}/${name}:${version}")
            }

            stage("Version: tag") {
                if (env.BRANCH_NAME == "main") {
                    sh(script: """
                        echo ${version} > VERSION
                        git config --local user.email "ugurluerdi@gmail.com"
                        git config --local user.name "Erdi Uğurlu"
                        git checkout main
                        git add VERSION
                        git status
                        git commit --message="noticket: Update to version ${version}"
                        git tag --annotate --message="Tagging ${version}" "${version}"
                        git push --follow-tags origin main
                    """)
                }
            }
        }
    }
}