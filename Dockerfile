#version 0.0.1 for darknet
#可以from ubuntu:16.04 
#配置环境是python2 vim pip  tf pillow keras pillow scipy 等
#FROM wkwu001/ubuntu-force
FROM ubuntu:16.04
MAINTAINER WWK wenkang529@163.com

#change ubuntu source
RUN cp /etc/apt/sources.list /etc/apt/sources.list.back
RUN sed -i 's#http://archive.ubuntu.com#http://mirrors.aliyun.com#g' /etc/apt/sources.list
RUN apt-get update &&  apt-get install -y autoconf \
	automake \
	libtool \
	git \
	wget \
	vim

#opencv
RUN apt-get install -y build-essential

#install python
RUN apt-get install -y build-essential libssl-dev libevent-dev libjpeg-dev libxml2-dev libxslt-dev \
	python python-dev \
	python-pip

#python pip lib
RUN pip install numpy 
RUN pip install	pillow 

RUN pip install	h5py 
RUN pip install	pydot 
RUN pip install	scipy 
RUN pip install	runcython 

#tensorflow
RUN pip install tensorflow

#matplotlib 3.0 not support python2.X -python3.4 ,only py3.5+ are support.
#RUN pip install matplotlib 
RUN pip install keras 
RUN pip install sklearn

#COPY ./SSD-Tensorflow-master /home/
#https://blog.csdn.net/qq_21046135/article/details/75804917

