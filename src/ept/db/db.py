# coding=utf-8
from sqlalchemy import create_engine, Column, Integer, String,\
    text, Text
from sqlalchemy.ext.declarative import declarative_base
from sqlalchemy.orm import sessionmaker, scoped_session


class DbModel(object):
    Base = declarative_base()
    metadata = Base.metadata

    engine = create_engine('sqlite:///db/test_log.db')
    # test_engine = create_engine('mysql://root:root@127.0.0.1:3306/test?charset=utf8'
    #                             , pool_size=20, max_overflow=100)
    session = scoped_session(sessionmaker(bind=engine))

    class Log(Base):
        __tablename__ = 'log'

        id = Column(Integer, primary_key=True)
        name = Column(String(100), nullable=False, server_default=text("''"))
        resp = Column(Text, nullable=False)
        headers = Column(Text, nullable=False)
        status_code = Column(Integer, nullable=False)
        elapsed_time = Column(String(50), nullable=False)
        start_time = Column(String(50), nullable=False)
        end_time = Column(String(50), nullable=False)
        exception = Column(Text, nullable=True)

    def create_db(self):
        self.metadata.create_all(self.engine)
