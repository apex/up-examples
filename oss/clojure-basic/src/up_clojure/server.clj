(ns up-clojure.server
  (:require [ring.adapter.jetty :as jetty])
  (:gen-class))

(defn handler [request]
  {:status 200
   :headers {"Content-Type" "text/html"}
   :body "Hello World from Clojure"})

(defn -main [& args]
  (let [port (or (some-> (System/getenv "PORT") Integer/parseInt)
                 3000)]
    (jetty/run-jetty handler {:port port})))
