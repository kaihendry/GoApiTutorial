#!/bin/bash
# assuming mariadb is set to localhost in your /etc/hosts
mysql -h mariadb -P 3306 -u root --password=secret rest_api_example
