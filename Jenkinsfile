// Define el agente donde se ejecutará el pipeline.
// 'linux-vps-agent' debe coincidir con la etiqueta que le diste a tu nodo Agente Jenkins en tu VPS.
pipeline {
    agent {
        label 'linux-vps-agent'
    }

    // Define variables de entorno que puedes necesitar.
    // DOCKER_HUB_CRED_ID es el ID de la credencial 'Username with password' que creaste en Jenkins
    // para acceder a Docker Hub.
    environment {
        APP_NAME = 'my-jeknins-app' // <-- ¡ACTUALIZA ESTO! Nombre de tu aplicación/contenedor
        DOCKER_IMAGE_NAME = 'codearl/my-jeknins-app' // <-- ¡ACTUALIZA ESTO! Usuario de Docker Hub + nombre de imagen
    }

    stages {
        stage('Clean Workspace') {
            steps {
                script {
                    echo "Cleaning workspace: ${WORKSPACE}"
                    // Limpiar el directorio de trabajo del agente antes de clonar
                    // Útil para asegurar un build limpio si hay artefactos de builds anteriores
                    cleanWs()
                }
            }
        }

        stage('Checkout Code') {
            steps {
                script {
                    echo "Checking out code from GitHub..."
                    // Jenkins clonará automáticamente el repositorio en el directorio de trabajo del agente
                    // si el SCM está configurado en el Job.
                    // Si tu repo es privado y necesitas credenciales específicas para clonar,
                    // puedes usar: git branch: 'main', credentialsId: 'your-github-cred-id'
                    checkout scm
                }
            }
        }

        stage('Debug Path') {
            steps {
                sh 'pwd'
                sh 'ls -F'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    echo "Building Docker image: ${DOCKER_IMAGE_NAME}:latest"
                    // Construye la imagen Docker usando el Dockerfile en la raíz del repositorio
                    sh "docker build -t ${DOCKER_IMAGE_NAME}:latest ."
                }
            }
        }

        stage('Test Docker Image (Optional)') {
            steps {
                script {
                    echo "Running tests against the Docker image (if applicable)..."
                    // Aquí puedes añadir pasos para ejecutar pruebas
                    // Por ejemplo, lanzar un contenedor temporal para pruebas:
                    // sh "docker run --rm ${DOCKER_IMAGE_NAME}:latest npm test"
                    // O ejecutar pruebas de integración si tienes un contenedor de pruebas separado.
                }
            }
        }

        // stage('Login to Docker Hub & Push Image') {
        //     steps {
        //         script {
        //             echo "Logging into Docker Hub and pushing image..."
        //             // Usa la credencial de Docker Hub configurada en Jenkins
        //             withCredentials([usernamePassword(credentialsId: "${DOCKER_HUB_CRED_ID}",
        //                                               usernameVariable: 'DOCKER_USER',
        //                                               passwordVariable: 'DOCKER_PASS')]) {
        //                 sh "echo ${DOCKER_PASS} | docker login -u ${DOCKER_USER} --password-stdin"
        //                 sh "docker push ${DOCKER_IMAGE_NAME}:latest"
        //             }
        //         }
        //     }
        // }

        stage('Deploy New Container') {
            steps {
                script {
                    echo "Deploying new container: ${APP_NAME}"
                    // Detener y eliminar el contenedor existente
                    // '|| true' evita que el pipeline falle si el contenedor no existe
                    sh "docker stop ${APP_NAME} || true"
                    sh "docker rm ${APP_NAME} || true"

                    // Eliminar la imagen vieja (opcional, para liberar espacio)
                    // sh "docker rmi ${DOCKER_IMAGE_NAME}:latest || true" # ¡Cuidado! Puede borrar la imagen recién empujada si no es otra etiqueta.
                    // Mejor: sh "docker rmi \$(docker images -q --filter 'dangling=true') || true" # Borra imágenes "dangling"

                    // Ejecutar el nuevo contenedor con la nueva imagen
                    // Ajusta los puertos y volúmenes según las necesidades de tu aplicación
                    sh "docker run -d --name ${APP_NAME} -p 8081:8081 ${DOCKER_IMAGE_NAME}:latest"
                }
            }
        }
    }

    // Acciones post-build (siempre se ejecutan)
    post {
        always {
            echo "Pipeline finished for ${APP_NAME}."
        }
        success {
            echo "Deployment successful for ${APP_NAME}!"
            // sh "curl -X POST -H 'Content-type: application/json' --data '{\"text\":\"Deployment to production successful!\"}' YOUR_SLACK_WEBHOOK_URL"
        }
        failure {
            echo "Deployment failed for ${APP_NAME}!"
            // sh "curl -X POST -H 'Content-type: application/json' --data '{\"text\":\"Deployment to production FAILED! Check Jenkins logs.\"}' YOUR_SLACK_WEBHOOK_URL"
        }
        unstable {
            echo "Pipeline was unstable (e.g., tests failed but build passed)."
        }
        changed {
            echo "Pipeline status changed since last run."
        }
    }
}