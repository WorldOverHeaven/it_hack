FROM maven:3.8.4-openjdk-17  AS build
COPY . .
RUN mvn -f pom.xml clean package -T 4C

FROM eclipse-temurin:17-jdk
RUN mkdir crypt
COPY --from=build target/client-app-0.0.1-SNAPSHOT.jar /client-app-0.0.1-SNAPSHOT.jar
ENTRYPOINT ["java","-jar","/client-app-0.0.1-SNAPSHOT.jar"]
