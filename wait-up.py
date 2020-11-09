from ping3 import ping
import sys
import time
import _thread
import os
import platform

ping_passed = False

help_msg = '''Usage: python3 wait-up.py {ip}

    requirements: ping3~=2.6.6


'''


def ping_with_shell(__ip__):
    lines = os.popen('ping -c 1 %s' % __ip__).read().split('\n')
    cursor_line = ''
    for line in lines:
        if line.find('bytes from') >= 0:
            cursor_line = line
            break
    if cursor_line.find('time=') > 0:
        tmp = cursor_line[cursor_line.find('time=') + 5:]
        return float(tmp[:tmp.find('ms')].strip())
    return False


def ping_with_raw_socks(__ip__):
    result = ping(__ip__)
    return result is not None


def ping_test(__ip__):
    if 'Darwin' == platform.system():
        try:
            return ping_with_shell(__ip__)
        except:
            return False
    try:
        return ping_with_raw_socks(__ip__)
    except:
        return False


def ping_thread(__ip__):
    while True:
        if ping_test(__ip__):
            global ping_passed
            ping_passed = True
            break


def main():
    if len(sys.argv) < 2:
        print(help_msg)
        exit(0)
    ip_address = sys.argv[1]
    print('Waiting for [%s] connecting ' % ip_address, end='')
    _thread.start_new_thread(ping_thread, (ip_address,))
    while not ping_passed:
        print('.', end='')
        sys.stdout.flush()
        time.sleep(1)
    print()
    sys.stdout.flush()


if __name__ == '__main__':
    main()
