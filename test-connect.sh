#!/bin/bash
# assuming mariadb is set to localhost in your /etc/hosts
result=$(echo 'show create database rest_api_example\G;' | mysql -h mariadb -P 3306 -u root --password=secret)

if ! echo "$result" | grep utf8mb4_unicode_ci
then
	echo FAIL, utf8mb4_unicode_ci not found
fi


result=$(echo 'show create table users\G;' | mysql -h mariadb -P 3306 -u root --password=secret rest_api_example)
if ! echo "$result" | grep utf8mb4_unicode_ci
then
	echo FAIL, utf8mb4_unicode_ci not found
fi


