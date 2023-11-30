#!/bin/bash

systemctl stop appmetrica.service
rm -rf /tmp/appmetrica.sock
systemctl start appmetrica.service
systemctl status appmetrica.service