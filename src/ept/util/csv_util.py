import csv
import os


def write(file_name, *args):
    with open(file_name, 'a') as f:
        '''
        name, resp, headers, status_code, elapsed_time, start_time,
        end_time, exception
        '''
        writer = csv.writer(f)
        writer.writerow([args[0], args[1], args[2], args[3], args[4],
                         args[5], args[6], args[7]])


def remove(file_name):
    try:
        os.remove(file_name)
    except FileNotFoundError:
        pass


def create(file_name):
    with open(file_name, 'w') as f:
        writer = csv.writer(f)
        writer.writerow(['name', 'resp', 'headers', 'status_code',
                         'elapsed_time', 'start_time', 'end_time', 'exception'])
