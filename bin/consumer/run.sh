#!/bin/bash

systemctl stop consumer.service
rm -rf /tmp/consumer.sock
systemctl start consumer.service
systemctl status consumer.service