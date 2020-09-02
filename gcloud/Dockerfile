FROM google/cloud-sdk

ARG HELM_VERSION=v3.2.4
ARG HELM_LOCATION="https://get.helm.sh"
ARG HELM_FILENAME="helm-${HELM_VERSION}-linux-amd64.tar.gz"
ARG HELM_SHA256="8eb56cbb7d0da6b73cd8884c6607982d0be8087027b8ded01d6b2759a72e34b1"
RUN curl --retry 3 --retry-connrefused -LO ${HELM_LOCATION}/${HELM_FILENAME} && \
    echo Verifying ${HELM_FILENAME}... && \
    sha256sum ${HELM_FILENAME} | grep -q "${HELM_SHA256}" && \
    echo Extracting ${HELM_FILENAME}... && \
    tar zxvf ${HELM_FILENAME} && mv /linux-amd64/helm /usr/local/bin/ && \
    rm ${HELM_FILENAME} && rm -r /linux-amd64

#    helm plugin install https://github.com/futuresimple/helm-secrets && \
#    helm plugin install https://github.com/hypnoglow/helm-s3.git && \
RUN helm plugin install https://github.com/databus23/helm-diff --version v3.1.3 && \
    helm plugin install https://github.com/aslafy-z/helm-git.git

ARG HELMFILE_VERSION=v0.126.2
ARG HELMFILE_LOCATION="https://github.com/roboll/helmfile/releases/download/${HELMFILE_VERSION}"
ARG HELMFILE_FILENAME="helmfile_linux_amd64"
ARG HELMFILE_SHA256="82c49932492df88707690b81c0032cb7591bb2e38388675a55530322e374cedd"
RUN curl --retry 3 --retry-connrefused -LO ${HELMFILE_LOCATION}/${HELMFILE_FILENAME} && \
    echo Verifying ${HELMFILE_FILENAME}... && \
    sha256sum ${HELMFILE_FILENAME} | grep -q "${HELMFILE_SHA256}" && \
    mv ${HELMFILE_FILENAME} /usr/local/bin/helmfile && \
    chmod +x /usr/local/bin/helmfile
