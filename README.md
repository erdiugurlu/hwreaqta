# Simple Golang Service

This repository holds code of the simple application as requested from ReaQta and Jenkinsfile of the  the container application.  

In order to build the code automatically, the following steps were done on Jenkins and the relevant GitHub repository.

 * First, an ssh key created on the EC2 server which Jenkins runs for pulling the codes from the repository
 * Deploy the public key on the github repository by following ´Settings -> Deploy keys´
 * In order to complete the Webhook integration on github, update Webhook setting by following ´Settings -> Webhook -> Manage webhook´
   * ´Payload URL´ is set http://<ec2-jenkins-server-ip>:8080/github-webhook/ 
   * ´Content type´ is set application/json
   * Select ´Just the push event´
 * The following configuration were done by creating Multi Branch Pipeline:
   * ´Repository HTTPS URL´ is https://github.com/erdiugurlu/hwreaqta
   * The relevant private key is selected on ´Checkout over SSH´
   * The other configurations are kept by default on the pipeline configuration.

## Jenkinsfile

Mainly, it includes three main steps like **Version Update, Docker Build and Docker Push**. I worked on main branch but in order to cover the other branch beside master/main, a custom name is given if the code is pushed on the application branch. 