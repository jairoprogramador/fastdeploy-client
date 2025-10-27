FROM ubuntu:22.04@sha256:f9ff1df8e3e896d1c031de656e6b21ef91329419aba21e4a2029f0543e97243b

ARG MAVEN_VERSION="3.9.11"
ARG TERRAFORM_VERSION="1.13.3"
ARG KUBECTL_VERSION="1.34.1"
ARG KUBELOGIN_VERSION="0.2.12"
ARG FASTDEPLOY_VERSION="1.0.3"
ARG DOCKER_VERSION="28.5.1"
ARG DEV_GID=1001

ENV DEBIAN_FRONTEND=noninteractive \
    JAVA_HOME=/usr/lib/jvm/java-17-openjdk-amd64 \
    MAVEN_HOME=/usr/share/maven \
    PATH="/usr/share/maven/bin:$PATH"

SHELL ["/bin/bash", "-o", "pipefail", "-c"]

RUN groupadd -o --system --gid "$DEV_GID" fastdeploy && \
    useradd --system --uid 1001 --gid fastdeploy --shell /bin/bash --create-home fastdeploy

RUN apt-get update && \
    apt-get install -y --no-install-recommends \
    curl \
    unzip \
    wget \
    git \
    # Java (OpenJDK 17) \
    openjdk-17-jdk \
    && \
    # --- INSTALACIÓN DE MAVEN --- \
    wget "https://dlcdn.apache.org/maven/maven-3/${MAVEN_VERSION}/binaries/apache-maven-${MAVEN_VERSION}-bin.tar.gz" -O /tmp/maven.tar.gz && \
    tar -xzf "/tmp/maven.tar.gz" -C /usr/share/ && \
    ln -s "/usr/share/apache-maven-${MAVEN_VERSION}" /usr/share/maven && \
    rm "/tmp/maven.tar.gz" && \
    # --- INSTALACIÓN DEL CLIENTE DOCKER --- \
    curl -fsSL "https://download.docker.com/linux/static/stable/x86_64/docker-${DOCKER_VERSION}.tgz" -o /tmp/docker.tgz && \
    tar -xzf /tmp/docker.tgz -C /tmp && \
    mv /tmp/docker/docker /usr/local/bin/ && \
    rm -rf /tmp/docker.tgz /tmp/docker && \
    # --- INSTALACIÓN DE TERRAFORM --- \
    wget "https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip" -O /tmp/terraform.zip && \
    unzip /tmp/terraform.zip -d /usr/local/bin/ && \
    rm "/tmp/terraform.zip" && \
    # --- INSTALACIÓN DE AZURE CLI --- \
    curl -sL https://aka.ms/InstallAzureCLIDeb | bash && \
    # --- INSTALACIÓN DE KUBELOGIN --- \
    wget "https://github.com/Azure/kubelogin/releases/download/v${KUBELOGIN_VERSION}/kubelogin-linux-amd64.zip" -O /tmp/kubelogin.zip && \
    unzip /tmp/kubelogin.zip -d /tmp && \
    mv /tmp/bin/linux_amd64/kubelogin /usr/local/bin && \
    rm -rf /tmp/kubelogin.zip /tmp/bin && \
    # --- INSTALACIÓN DE KUBECTL --- \
    wget "https://dl.k8s.io/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl" -O /usr/local/bin/kubectl && \
    chmod +x /usr/local/bin/kubectl && \
    # --- INSTALACIÓN DE FASTDEPLOY --- \
#    wget "https://github.com/jairoprogramador/fastdeploy/releases/download/v${FASTDEPLOY_VERSION}/fastdeploy_${FASTDEPLOY_VERSION}_linux_amd64.tar.gz" -O /tmp/fastdeploy.tar.gz && \
#    tar -xzf "/tmp/fastdeploy.tar.gz" -C /tmp && \
#    mv /tmp/fd /usr/local/bin/ && \
#    chmod 755 /usr/local/bin/fd && \
#    rm /tmp/fastdeploy.tar.gz \
#    && \
    # --- LIMPIEZA --- \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY fd /usr/local/bin/fd
RUN chmod +x /usr/local/bin/fd

USER fastdeploy

RUN git config --global --add safe.directory '*'

WORKDIR /home/fastdeploy/app

ENTRYPOINT [ "fd" ]

CMD ["--version"]
