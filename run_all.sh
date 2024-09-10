#!/bin/bash

args=`find dataset -type f | xargs`

time bash go/concurrent-0/serial/run.sh $args
 
time bash go/concurrent-0/concurrent/run.sh $args
