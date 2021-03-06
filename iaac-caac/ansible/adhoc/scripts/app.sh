#!/bin/bash
#
# App server orchestration ad-hoc tasks.
set -e

# Configure Django on app server.
ansible app -b -m yum -a "name=python3-pip state=present"
ansible app -b -m pip -a "name=django<4 state=present" -e ansible_python_interpreter=/usr/bin/python3

# Check Django version.
ansible app -a "python3 -m django --version"

# Other commands from the book.
# ansible app -b -a "systemctl status chronyd"
ansible app -b -m group -a "name=admin state=present"
ansible app -b -m user -a "name=johndoe group=admin createhome=yes"
ansible app -b -m user -a "name=johndoe state=absent remove=yes"
ansible app -b -m package -a "name=git state=present"
