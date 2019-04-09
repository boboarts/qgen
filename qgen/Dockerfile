###############################
# for Qgen CentOS
# Version 0.0.1
###############################
FROM centos
LABEL author = "Gary SHEN" \
      email = "boboarts@gamil.com" \
      version = "0.0.1" \
      descprtion = "The qgen program based on centos image"
RUN mkdir -p /etc/qgen
WORKDIR /etc/qgen
RUN curl -O https://raw.githubusercontent.com/boboarts/qgen/master/bin/centos/qgen
RUN curl -O https://raw.githubusercontent.com/boboarts/qgen/master/bin/centos/template.html
RUN chmod 755 qgen
CMD ["/etc/qgen/qgen"]
EXPOSE 9923
# docker build -t "boboarts/qgen" .
# docker run -d -p 9901:9923 --restart always --name qgen boboarts/qgen
# docker run -it --name qgen boboarts/qgen /bin/bash
# docker exec -it qgen /bin/bash
# docker commit qgen boboarts/qgen


###############################
# for Qgen Alpine
# Version 0.0.1
###############################
# download qgen programs first
curl -O https://raw.githubusercontent.com/boboarts/qgen/master/bin/centos/qgen
curl -O https://raw.githubusercontent.com/boboarts/qgen/master/bin/centos/problems.tmpl.html
curl -O https://raw.githubusercontent.com/boboarts/qgen/master/bin/centos/index.tmpl.html

FROM frolvlad/alpine-glibc
LABEL author = "Gary SHEN" \
      email = "boboarts@gamil.com" \
      version = "0.0.2" \
      descprtion = "The qgen program based on alpine image"
RUN mkdir -p /etc/qgen
WORKDIR /etc/qgen
ADD qgen qgen
ADD index.tmpl.html index.tmpl.html
ADD problems.tmpl.html problems.tmpl.html
RUN chmod 755 qgen
CMD ["/etc/qgen/qgen"]
EXPOSE 9923
# docker build -t "boboarts/qgen:alpine" .
# docker run -d -p 9901:9923 --restart always --name qgen boboarts/qgen:alpine
# docker commit qgen boboarts/qgen:alpine

# curl localhost:9923




