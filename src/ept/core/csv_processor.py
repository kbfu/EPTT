# coding=utf-8
import csv


def write(file_name, *args):
    with open(file_name, 'a') as f:
        '''
        name, resp, headers, status_code, elapsed_time, start_time,
        end_time, exception
        '''
        writer = csv.writer(f)
        writer.writerow([args[0], args[1], args[2], args[3], args[4],
                         args[5], args[6], args[7]])
