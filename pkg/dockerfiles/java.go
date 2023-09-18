package dockerfiles

const Java = `FROM adoptopenjdk:11-jre-hotspot

WORKDIR /app

COPY your-app.jar /app/your-app.jar

RUN groupadd -r yourapp && useradd -r -g yourapp yourapp
USER yourapp

CMD ["java", "-jar", "/app/your-app.jar"]`
