package com.atlas.cdc.event.producer;

import com.atlas.cdc.EventProducerRegistry;
import com.atlas.cos.command.AdjustHealthCommand;
import com.atlas.cos.constant.CommandConstants;

public final class AdjustHealthProducer {
   private AdjustHealthProducer() {
   }

   public static void emit(int characterId, int amount) {
      EventProducerRegistry.getInstance().send(CommandConstants.TOPIC_ADJUST_HEALTH, characterId,
            new AdjustHealthCommand(characterId, amount));
   }
}
