#!/bin/bash
# Flyway migration command

flyway -locations=filesystem:${PWD}/sql migrate
