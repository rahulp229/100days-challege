import 'package:flutter/material.dart';

void main() {
  runApp(MyApp());
}

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      title: 'Greeting',
      home: Scaffold(
        appBar: AppBar(
          title: Text('Greeting'),
        ),
        body: Center(
          child: Text('Hello World'),
        ),
      ),
    );
  }
}

Skillset:
Programming Languages: Golang, Java, C#, Python
Distributed Systems: EventHub, Redis, ElasticSearch, Kafka, Websockets
Distributed tracing - Jaeger, Monitoring - Prometheus, Grafana, Kibana
Cloud and Container Solutions: Azure, Kubernetes, Docker
Database: CosmosDB, Postgres, Oracle
CI Tools: CI/CD, GoCD, Pipeline as Code
Development Practices: Clean code, Test Driven Development (TDD), Multilayered testing, OOP, Refactoring, Pair Programming