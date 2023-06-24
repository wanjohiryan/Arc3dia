#!/bin/sh
# Logger from this post http://www.cubicrace.com/2016/03/log-tracing-mechnism-for-shell-scripts.html

function INFO(){
    local function_name="${FUNCNAME[1]}"
    local msg="$1"
    timeAndDate=`date`
    echo "\033[01;34m[$timeAndDate] [INFO] [${0}] $msg\033[0m"
}

function DEBUG(){
    local function_name="${FUNCNAME[1]}"
    local msg="$1"
    timeAndDate=`date`
    echo "\033[01;33m[$timeAndDate] [DEBUG] [${0}] $msg\033[0m"
}

function ERROR(){
    local function_name="${FUNCNAME[1]}"
    local msg="$1"
    timeAndDate=`date`
    echo "\033[01;31m[$timeAndDate] [ERROR] $msg033[0m"
}