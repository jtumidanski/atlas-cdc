package com.atlas.cdc.event.producer;

import com.atlas.cdc.EventProducerRegistry;
import com.atlas.cos.command.AdjustMesoCommand;
import com.atlas.cos.constant.CommandConstants;

public final class AdjustMesoProducer {
   private AdjustMesoProducer() {
   }

   public static void emit(int characterId, int amount, boolean show) {
      EventProducerRegistry.getInstance().send(CommandConstants.TOPIC_ADJUST_MESO, characterId,
            new AdjustMesoCommand(characterId, amount, show));
   }
}
