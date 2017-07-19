# coding=utf-8
from setuptools import setup
from os.path import abspath, dirname, join

CURDIR = dirname(abspath(__file__))

with open(join(CURDIR, 'requirements.txt')) as f:
    REQUIREMENTS = f.read().splitlines()

setup(name='ept',
      version='1.0.2',
      description='ept',
      classifiers=[
          'Programming Language :: Python :: 3.6.2',
      ],
      packages=['ept.core', 'ept'],
      package_dir={'': 'src'},
      install_requires=REQUIREMENTS,
      scripts=['bin/ept_run'],
      include_package_data=True
      )
