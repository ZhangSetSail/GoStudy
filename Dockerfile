FROM openjdk:11-jre as java_builder

ARG MY_SERVER_FILE=tc-basic-data-app.jar
ADD ./target/${MY_SERVER_FILE} /app/jar/server.jar

FROM ubuntu:20.04

ENV TZ=Asia/Shanghai
RUN ln -sf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

RUN apt-get update && apt-get install -y wget
ENV TDENGINE_VERSION=2.6.0.34
RUN wget -c https://www.taosdata.com/assets-download/TDengine-client-${TDENGINE_VERSION}-Linux-x64.tar.gz
&& tar xvf TDengine-client-${TDENGINE_VERSION}-Linux-x64.tar.gz
&& cd TDengine-client-${TDENGINE_VERSION}
&& ./install_client.sh
&& cd ../
&& rm -rf TDengine-client-${TDENGINE_VERSION}-Linux-x64.tar.gz TDengine-client-${TDENGINE_VERSION}

ENV JAVA_OPTS=-Xms128m -Xmx384m -Djava.security.egd=file:/dev/./urandom

COPY --from=java_builder /app/jar/server.jar /app/jar/server.jar
COPY --from=java_builder /usr/local/openjdk-11 /usr/local/openjdk-11

EXPOSE 7803

CMD ["/usr/local/openjdk-11/bin/java", "$JAVA_OPTS", "-jar", "/app/jar/server.jar"]