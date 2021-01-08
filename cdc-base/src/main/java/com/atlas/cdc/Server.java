package com.atlas.cdc;

import com.atlas.cdc.event.consumer.DamageCharacterCommandConsumer;
import com.atlas.kafka.consumer.SimpleEventConsumerBuilder;

public class Server {
   public static void main(String[] args) {
      //      URI uri = UriBuilder.host(RestService).uri();
      //      final HttpServer server = RestServerFactory.create(uri, "com.atlas.cdc.rest");
      SimpleEventConsumerBuilder.builder()
            .addConsumer(new DamageCharacterCommandConsumer())
            .initialize();
   }
}
