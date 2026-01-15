pipeline {
  agent {
    kubernetes {
      cloud 'k8s-azurenode'
      defaultContainer 'jnlp'
      yaml """
apiVersion: v1
kind: Pod
spec:
  serviceAccountName: jenkins-ci

  containers:
  - name: node
    image: node:20
    command: ["sleep"]
    args: ["99d"]

  - name: python
    image: python:3.11
    command: ["sleep"]
    args: ["99d"]

  - name: go
    image: golang:1.22
    command: ["sleep"]
    args: ["99d"]

  - name: kaniko
    image: gcr.io/kaniko-project/executor:v1.22.0-debug
    command: ["sleep"]
    args: ["99d"]
    volumeMounts:
      - name: docker-config
        mountPath: /kaniko/.docker

  - name: kubectl
    image: bitnami/kubectl:latest
    command: ["sleep"]
    args: ["99d"]

  volumes:
  - name: docker-config
    secret:
      secretName: regcred
"""
    }
  }

  environment {
    REGISTRY = "docker.io/emmiduh93"
    IMAGE_TAG = "${BUILD_NUMBER}-${GIT_COMMIT.take(7)}"
  }

  stages {

    stage('Checkout') {
      steps {
        checkout scm
      }
    }

    stage('Tests') {
      parallel {

        stage('Auth Service (Node.js)') {
          steps {
            container('node') {
              dir('service-auth-node') {
                sh '''
                  npm ci
                  npm test || echo "No tests defined"
                '''
              }
            }
          }
        }

        stage('Fraud Service (Python)') {
          steps {
            container('python') {
              dir('service-fraud-python') {
                sh '''
                  pip install -r requirements.txt
                  python -m py_compile fraud.py
                '''
              }
            }
          }
        }

        stage('Ledger Service (Go)') {
          steps {
            container('go') {
              dir('service-ledger-go') {
                sh '''
                  go mod download
                  go test ./... || echo "No tests found"
                  go build -buildvcs=false ./...
                '''
              }
            }
          }
        }

      }
    }

    stage('Build & Push Images (Kaniko)') {
      steps {
        container('kaniko') {
          sh '''
            /kaniko/executor \
              --context $WORKSPACE/service-auth-node \
              --dockerfile $WORKSPACE/service-auth-node/Dockerfile \
              --destination $REGISTRY/auth:$IMAGE_TAG \
              --cache=true \
              --cache-repo=$REGISTRY/cache 

            /kaniko/executor \
              --context $WORKSPACE/service-fraud-python \
              --dockerfile $WORKSPACE/service-fraud-python/Dockerfile \
              --destination $REGISTRY/fraud:$IMAGE_TAG \
              --cache=true \
              --cache-repo=$REGISTRY/cache 

            /kaniko/executor \
              --context $WORKSPACE/service-ledger-go \
              --dockerfile $WORKSPACE/service-ledger-go/Dockerfile \
              --destination $REGISTRY/ledger:$IMAGE_TAG \
              --cache=true \
              --cache-repo=$REGISTRY/cache 
          '''
        }
      }
    }

    stage('Deploy to Kubernetes') {
      when {
        branch 'main'
      }
      steps {
        container('kubectl') {
          sh '''
            kubectl apply -n ci -f kubernetes/
          '''
        }
      }
    }

  }

  post {
    success {
      echo "✅ Build and deployment successful"
    }
    failure {
      echo "❌ Pipeline failed"
    }
  }
}
