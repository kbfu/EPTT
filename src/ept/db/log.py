# coding=utf-8
from sqlalchemy import desc
from .db import DbModel
from ..core import date_util
from .decorator import close_db_connections


class Log(DbModel):
    @close_db_connections
    def add(self, name, resp, headers, elapsed_time, start_time, end_time, status_code,
            exception=''):
        data = self.Log(name=name, resp=resp, elapsed_time=elapsed_time,
                        start_time=start_time,
                        end_time=end_time, status_code=status_code,
                        headers=headers, exception=exception)
        self.session().add(data)
        self.session().commit()

    @close_db_connections
    def get_start_time(self, name):
        query = self.session().query(self.Log).filter_by(
            name=name).order_by(self.Log.start_time).first()
        return query.start_time

    @close_db_connections
    def get_global_end_time(self):
        query = self.session().query(self.Log).order_by(
            desc(self.Log.end_time)).first()
        return query.end_time

    @close_db_connections
    def get_names(self):
        query = self.session().query(self.Log.name).distinct(self.Log.name)
        return query

    @close_db_connections
    def get_end_times(self, name):
        query = self.session().query(self.Log.end_time).distinct(self.Log.end_time).filter_by(name=name)\
            .order_by(self.Log.end_time)
        return query

    @close_db_connections
    def get_tps(self, name, end_time):
        query = self.session().query(self.Log).filter(self.Log.end_time.like(end_time + '%'), self.Log.name == name)\
            .count()
        return query

    @close_db_connections
    def get_response_time(self, name):
        query = self.session().query(self.Log.elapsed_time).filter_by(name=name)
        return query

    @close_db_connections
    def get_status(self, name):
        query = self.session().query(self.Log.status).filter_by(name=name)
        return query

    @close_db_connections
    def delete_all(self):
        self.session().query(self.Log).delete()
        self.session().commit()

    @close_db_connections
    def get_count(self, name):
        query = self.session().query(self.Log).filter_by(name=name).count()
        return query
