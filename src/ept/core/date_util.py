# coding=utf-8
from datetime import datetime


def str2datetime(date_str):
    date = datetime.strptime(date_str, '%Y-%m-%d %H:%M:%S.%f')
    return date


def timestamp_now():
    return datetime.timestamp(datetime.now())
