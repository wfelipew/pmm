FROM centos:7

ARG VERSION
ARG BUILD_DATE

LABEL org.opencontainers.image.created ${BUILD_DATE}
LABEL org.opencontainers.image.licenses AGPL-3.0
LABEL org.opencontainers.image.title Percona Monitoring and Management
LABEL org.opencontainers.image.vendor Percona
LABEL org.opencontainers.image.version ${VERSION}

EXPOSE 80 443

WORKDIR /opt

RUN yum -y install epel-release && yum -y install ansible

COPY RPMS /tmp/RPMS
COPY gitCommit /tmp/gitCommit

COPY ansible /opt/ansible
RUN cp -r /opt/ansible/roles /opt/ansible/pmm2-docker/roles
RUN ansible-playbook -vvv -i 'localhost,' -c local /opt/ansible/pmm2-docker/main.yml \
    && ansible-playbook -vvv -i 'localhost,' -c local /usr/share/pmm-update/ansible/playbook/tasks/update.yml \
    && ansible-playbook -vvv -i 'localhost,' -c local /opt/ansible/pmm2/post-build-actions.yml


RUN mkdir /.ansible/

RUN chown -R pmm.root /opt
RUN chown -R pmm.root /var
RUN chown -R pmm.root /srv
RUN chown -R pmm.root /usr/share
RUN chown -R pmm.root /etc
RUN chown -R pmm.root /usr/sbin
RUN chown -R pmm.root /run
RUN chown -R pmm.root /usr/local/percona/
RUN chown -R pmm.root /.ansible/
RUN chown -R pmm.root /home

RUN chmod -R g+rwx /opt
RUN chmod -R g+rwx /var
RUN chmod -R g+rwx /srv
RUN chmod -R g+rwx /usr/share
RUN chmod -R g+rwx /etc
RUN chmod -R g+rwx /usr/sbin
RUN chmod -R g+rwx /run
RUN chmod -R g+rwx /usr/local/percona
RUN chmod -R g+rwx /.ansible/
RUN chmod -R g+rwx /home

COPY entrypoint.sh /opt/entrypoint.sh
HEALTHCHECK --interval=3s --timeout=2s --start-period=10s --retries=3 CMD curl -f http://127.0.0.1/v1/readyz || exit 1
CMD ["/opt/entrypoint.sh"]
