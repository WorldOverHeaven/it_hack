FROM maven AS build
COPY pom.xml ./pom.xml
RUN mvn -f pom.xml clean package -T 4C

FROM eclipse-temurin:17-jdk
COPY --from=build target/client-app-0.0.1-SNAPSHOT.jar /client-app-0.0.1-SNAPSHOT.jar
ENTRYPOINT ["java","-jar","/client-app-0.0.1-SNAPSHOT.jar"]
