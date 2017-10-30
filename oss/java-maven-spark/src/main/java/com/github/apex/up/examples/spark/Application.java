package com.github.apex.up.examples.spark;

import static java.lang.Integer.parseInt;
import static java.lang.System.getenv;
import static java.util.Optional.ofNullable;
import static spark.Spark.get;
import static spark.Spark.port;

public class Application {
  public static void main(String[] args) {
    String port = ofNullable(getenv("PORT")).orElse("3000");
    port(parseInt(port));
    get("/", (request, response) -> "Hello World");
  }
}
