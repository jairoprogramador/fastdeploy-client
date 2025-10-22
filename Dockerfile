FROM ubuntu:22.04@sha256:f9ff1df8e3e896d1c031de656e6b21ef91329419aba21e4a2029f0543e97243b

ARG MAVEN_VERSION="3.9.11"
ARG TERRAFORM_VERSION="1.13.3"
ARG KUBECTL_VERSION="1.34.1"
ARG KUBELOGIN_VERSION="0.1.4"
ARG FASTDEPLOY_VERSION="1.0.3"
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
    unzip \
    wget \
    git \
    # Java (OpenJDK 17)
    openjdk-17-jdk \
    && \
    # --- INSTALACIÓN DE MAVEN ---
    wget "https://dlcdn.apache.org/maven/maven-3/${MAVEN_VERSION}/binaries/apache-maven-${MAVEN_VERSION}-bin.tar.gz" -O /tmp/maven.tar.gz && \
    tar -xzf "/tmp/maven.tar.gz" -C /usr/share/ && \
    ln -s "/usr/share/apache-maven-${MAVEN_VERSION}" /usr/share/maven && \
    rm "/tmp/maven.tar.gz" \
    && \
    # --- INSTALACIÓN DE TERRAFORM ---
    wget "https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip" -O /tmp/terraform.zip && \
    unzip /tmp/terraform.zip -d /usr/local/bin/ && \
    rm "/tmp/terraform.zip" \
    && \
    # --- INSTALACIÓN DE KUBECTL ---
    wget "https://dl.k8s.io/release/v${KUBECTL_VERSION}/bin/linux/amd64/kubectl" -O /usr/local/bin/kubectl && \
    chmod +x /usr/local/bin/kubectl \
    && \
    # --- INSTALACIÓN DE FASTDEPLOY ---
    wget "https://github.com/jairoprogramador/fastdeploy/releases/download/v${FASTDEPLOY_VERSION}/fastdeploy_${FASTDEPLOY_VERSION}_linux_amd64.tar.gz" -O /tmp/fastdeploy.tar.gz && \
    tar -xzf "/tmp/fastdeploy.tar.gz" -C /tmp && \
    mv /tmp/fd /usr/local/bin/ && \
    chmod 755 /usr/local/bin/fd && \
    rm /tmp/fastdeploy.tar.gz \
    && \
    # --- LIMPIEZA ---
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

USER fastdeploy

RUN git config --global --add safe.directory '*'

WORKDIR /home/fastdeploy/app

ENTRYPOINT [ "fd" ]

CMD ["--version"]
