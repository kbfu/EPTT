# coding=utf-8
def close_db_connections(func):
    def wrapper(*args, **kwargs):
        retval = func(*args, **kwargs)
        args[0].session().close()
        args[0].session.remove()
        args[0].engine.dispose()
        return retval
    return wrapper
