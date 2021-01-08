FROM maven:3.6.3-openjdk-14-slim AS build

COPY settings.xml /usr/share/maven/conf/

COPY pom.xml pom.xml
COPY cdc-api/pom.xml cdc-api/pom.xml
COPY cdc-model/pom.xml cdc-model/pom.xml
COPY cdc-base/pom.xml cdc-base/pom.xml

RUN mvn dependency:go-offline package -B

COPY cdc-api/src cdc-api/src
COPY cdc-model/src cdc-model/src
COPY cdc-base/src cdc-base/src

RUN mvn install

FROM openjdk:14-ea-jdk-alpine
USER root

RUN mkdir service

COPY --from=build /cdc-base/target/ /service/

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.5.0/wait /wait

RUN chmod +x /wait

ENV JAVA_TOOL_OPTIONS -agentlib:jdwp=transport=dt_socket,server=y,suspend=n,address=*:5005

EXPOSE 5005

CMD /wait && java --enable-preview -jar /service/cdc-base-1.0-SNAPSHOT.jar -Xdebug