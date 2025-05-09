FROM amazoncorretto:17 AS build

WORKDIR /home/app

COPY . .

RUN chmod +x ./gradlew

RUN ./gradlew build --no-daemon --build-cache

FROM amazoncorretto:24-alpine

ARG JAR_FILE=build/libs/*.jar

COPY --from=build /home/app/$JAR_FILE /usr/local/lib/app.jar

EXPOSE 8080

ENTRYPOINT ["java", "-jar", "-Dfile.encoding=UTF-8", "-Dspring.profiles.active=production", "/usr/local/lib/app.jar"]
