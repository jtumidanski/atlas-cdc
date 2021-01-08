package com.atlas.cdc.event.producer;

import com.atlas.cdc.EventProducerRegistry;
import com.atlas.csrv.command.ShowDamageCharacterCommand;
import com.atlas.csrv.constant.CommandConstants;

public final class ShowCharacterDamageProducer {
   private ShowCharacterDamageProducer() {
   }

   public static void emit(byte skillId, int monsterId, int characterId, int mapId, int damage, int fake, byte direction,
                           boolean pgmr, int pgmr1, boolean isPg, int monsterUniqueId, int x, int y) {
      EventProducerRegistry.getInstance().send(CommandConstants.TOPIC_SHOW_DAMAGE_CHARACTER_COMMAND, characterId,
            new ShowDamageCharacterCommand(characterId, mapId, monsterId, monsterUniqueId, skillId, damage, fake, direction, x, y
                  , pgmr, pgmr1, isPg));
   }
}
