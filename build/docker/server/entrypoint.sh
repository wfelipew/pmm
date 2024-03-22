#!/bin/bash
set -o errexit

sed -i -e "s/pmm:.*/pmm:x:$UID:0:PMM Server:\/home\/pmm:\/bin\/false/" /etc/passwd

# init /srv if empty
DIST_FILE=/srv/pmm-distribution
if [ ! -f $DIST_FILE ]; then
    echo "File $DIST_FILE doesn't exist. Initialize /srv..."
    echo docker > $DIST_FILE
    mkdir -p /srv/{clickhouse,grafana,logs,postgres14,prometheus,nginx,victoriametrics}
    echo "Copying plugins and VERSION file"
    cp /usr/share/percona-dashboards/VERSION /srv/grafana/PERCONA_DASHBOARDS_VERSION
    cp -r /usr/share/percona-dashboards/panels/ /srv/grafana/plugins
    # chown -R pmm:root /srv/grafana
    # chown pmm:root /srv/{victoriametrics,prometheus,logs}
    # chown postgres:postgres /srv/postgres14
    echo "Generating self-signed certificates for nginx"
    bash /var/lib/cloud/scripts/per-boot/generate-ssl-certificate
    echo "Initializing Postgres"
    /usr/pgsql-14/bin/initdb -D /srv/postgres14
    echo "Enable pg_stat_statements extension"
    /usr/pgsql-14/bin/pg_ctl start -D /srv/postgres14
    psql postgres -c 'CREATE EXTENSION pg_stat_statements SCHEMA public'
    /usr/pgsql-14/bin/pg_ctl stop -D /srv/postgres14
fi

#/usr/pgsql-14/bin/pg_ctl start -D /srv/postgres14
#psql -U postgres -c 'CREATE USER pmm SUPERUSER'
#/usr/pgsql-14/bin/pg_ctl stop -D /srv/postgres14

# pmm-managed-init validates environment variables.
pmm-managed-init

# Start supervisor in foreground
exec supervisord -n -c /etc/supervisord.conf
