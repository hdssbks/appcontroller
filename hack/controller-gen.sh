#!/bin/bash

# cd project root dir
cd /home/appcontroller
controller-gen crd paths=./... output:crd:dir=artifacts/crd