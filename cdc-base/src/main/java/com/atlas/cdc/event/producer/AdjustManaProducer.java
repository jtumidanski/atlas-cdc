package com.atlas.cdc.event.producer;

import com.atlas.cdc.EventProducerRegistry;
import com.atlas.cos.command.AdjustManaCommand;
import com.atlas.cos.constant.CommandConstants;

public final class AdjustManaProducer {
   private AdjustManaProducer() {
   }

   public static void emit(int characterId, int amount) {
      EventProducerRegistry.getInstance().send(CommandConstants.TOPIC_ADJUST_MANA, characterId,
            new AdjustManaCommand(characterId, amount));
   }
}
